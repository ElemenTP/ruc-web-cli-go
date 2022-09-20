/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	username string
	password string
	secrets  string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login campus network",
	Long: `Either call "ruc-web-cli login --username YOUR_USERNAME --password YOUR_PASSWORD"
or "ruc-web-cli login --secrets SECRETS_JSON_FILENAME"
or "ruc-web-cli login", which will use local secrets.json to login.`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		if len(username) != 0 && len(password) != 0 {
			log.Println("reading username and password from command line arguments")
		} else {
			log.Printf("reading username and password from json \"%s\".\n", secrets)
			jsonFile, err := os.ReadFile(secrets)
			if err != nil {
				log.Fatalln(err)
			}
			j := make(map[string]interface{})
			err = json.Unmarshal(jsonFile, &j)
			if err != nil {
				log.Fatalln(err)
			}
			for k, v := range j {
				if k == "username" {
					username = Convert2String(v)
				} else if k == "password" {
					password = Convert2String(v)
				} else {
					continue
				}
			}
		}
		log.Printf("received username = \"%s\" and password = \"%s\"", username, password)
		if len(username) == 0 || len(password) == 0 {
			log.Fatal("username or password is empty")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "username of your campus network account")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "password of your campus network account")
	loginCmd.Flags().StringVarP(&secrets, "secrets", "s", "secrets.json", "path of a json file contains your username and password")
}

func Convert2String(v interface{}) string {
	switch record := v.(type) {
	case string:
		return record
	default:
		log.Fatalln("username and password in json file should be string")
		return ""
	}
}
