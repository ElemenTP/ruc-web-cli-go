package rucweb

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const BASE_URL string = "https://go.ruc.edu.cn"

var (
	HEADERS map[string]string
	client  http.Client
)

func HttpGet(turl string, params, headers map[string]string) (string, error) {
	furl := turl
	if params != nil {
		u, err := url.Parse(turl)
		if err != nil {
			return "", err
		}
		uparams := u.Query()
		for k, v := range params {
			uparams.Add(k, v)
		}
		u.RawQuery = uparams.Encode()
		furl = u.String()
	}
	req, err := http.NewRequest(http.MethodGet, furl, nil)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// func http_post(turl string, params, headers, data map[string]string) (string, error) {
// 	furl := turl
// 	if params != nil {
// 		u, err := url.Parse(turl)
// 		if err != nil {
// 			return "", err
// 		}
// 		uparams := u.Query()
// 		for k, v := range params {
// 			uparams.Add(k, v)
// 		}
// 		u.RawQuery = uparams.Encode()
// 		furl = u.String()
// 	}
// 	req, err := http.NewRequest(http.MethodPost, furl, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	for k, v := range headers {
// 		req.Header.Add(k, v)
// 	}
// 	req.ParseForm()
// 	for k, v := range data {
// 		req.PostForm.Add(k, v)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(body), nil
// }

func Login(username, password string) {
	ip := requestIP()
	log.Printf(`IP is %s\n`, ip)

	token := requestToken(username, ip)
	log.Printf(`Token is %s\n`, token)

	params := buildLoginParams(username, password, ip, token)
	j := requestLogin(params)
	if j["ecode"].(int) == 0 {
		log.Println("login requested.")
		if j["suc_msg"].(string) == "login_ok" {
			log.Println("login success!")
			log.Println("ip is", j["online_ip"])
			log.Println("username is", j["username"])
			log.Println("real name is", j["real_name"])
		} else if j["suc_msg"].(string) == "ip_already_online_error" {
			log.Fatalln("login failed.", "server reported this ip is already online, cannot login twice.")
		}
	} else {
		jsonstr, err := json.Marshal(j)
		if err != nil {
			jsonstr = []byte(err.Error())
		}
		log.Fatalln("server denied login request, response json as follows:", string(jsonstr))
	}
}

func requestLogin(params map[string]string) map[string]interface{} {
	turl := BASE_URL + "/cgi-bin/srun_portal"

	text, err := HttpGet(turl, params, HEADERS)
	if err != nil {
		log.Fatalln(err)
	}
	index := strings.IndexByte(text, '(')
	text = text[index+1 : len(text)-1]
	j := make(map[string]interface{})
	err = json.Unmarshal([]byte(text), &j)
	if err != nil {
		log.Fatalln(err)
	}
	return j
}

func buildLoginParams(username, password, ip, token string) map[string]string {
	const c_acid, c_enc_ver, c_n, c_type string = "1", "srun_bx1", "200", "1"
	info := map[string]string{
		"username": username,
		"password": password,
		"ip":       ip,
		"acid":     c_acid,
		"enc_ver":  c_enc_ver,
	}
	flatten_info, err := json.Marshal(info)
	if err != nil {
		log.Fatalln(err)
	}

	encoded_info_prefixed := `{SRBX1}` + GetBase64(GetxEncode(string(flatten_info), token))

	password_md5, err := GetMD5(password, token)
	if err != nil {
		log.Fatalln(err)
	}
	chksum_segments := []string{
		token,
		username,
		token,
		password_md5,
		token,
		c_acid,
		token,
		ip,
		token,
		c_n,
		token,
		c_type,
		token,
		encoded_info_prefixed,
	}
	chksum, err := GetSHA1(strings.Join(chksum_segments, ""))
	if err != nil {
		log.Fatalln(err)
	}
	time_nounce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	res := map[string]string{
		"callback":     "jQuery11240645308969735664_" + time_nounce,
		"action":       "login",
		"username":     username,
		"password":     `{MD5}` + password_md5,
		"ac_id":        c_acid,
		"ip":           ip,
		"chksum":       chksum,
		"info":         encoded_info_prefixed,
		"n":            c_n,
		"type":         c_type,
		"os":           "Windows 10",
		"name":         "Windows",
		"double_stack": "0",
		"_":            time_nounce,
	}

	return res
}

func requestIP() string {
	text, err := HttpGet(BASE_URL+"/", nil, HEADERS)
	if err != nil {
		log.Fatalln(err)
	}
	regex, err := regexp.Compile(`var CONFIG = (?s:{.*})`)
	if err != nil {
		log.Fatalln(err)
	}
	configSubmatch := regex.FindAllStringSubmatch(text, -1)
	if len(configSubmatch) == 0 {
		log.Fatalln(errors.New("no config find in text"))
	}
	config := configSubmatch[0][len(configSubmatch[0])-1]

	regex, err = regexp.Compile(`ip.*?:.*?"(.*?)"`)
	if err != nil {
		log.Fatalln(err)
	}
	ipSubmatch := regex.FindAllStringSubmatch(config, -1)
	if len(ipSubmatch) == 0 {
		log.Fatalln(errors.New("no ip find in config"))
	}
	ip := ipSubmatch[0][len(ipSubmatch[0])-1]
	return ip
}

func requestToken(username string, ip string) string {
	time_nounce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	turl := BASE_URL + "/cgi-bin/get_challenge"
	params := map[string]string{
		"callback": "jQuery112406382209524580216_" + time_nounce,
		"username": username,
		"ip":       ip,
		"_":        time_nounce,
	}
	text, err := HttpGet(turl, params, HEADERS)
	if err != nil {
		log.Fatalln(err)
	}
	index := strings.IndexByte(text, '(')
	text = text[index+1 : len(text)-1]
	j := make(map[string]interface{})
	err = json.Unmarshal([]byte(text), &j)
	if err != nil {
		log.Fatalln(err)
	}
	return j["challenge"].(string)
}

func init() {
	HEADERS = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36"}
	client = *http.DefaultClient
}
