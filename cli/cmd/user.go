// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/smallnest/goreq"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login your agenda account",
	Long: `To login your account with your correct username and password.
	 For example:

	Agenda login -uabb -p123`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		if username == "" {
			fmt.Println("please input your username")
			return
		}
		if password == "" {
			fmt.Println("please input your password")
			return
		}
		data := make(url.Values)
		data["username"] = []string{username}
		data["password"] = []string{password}

		res, err := http.PostForm("http://127.0.0.1:8080/v1/user/login", data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer res.Body.Close()
		DealWithResponse(res)

		fmt.Println(username + " is logined")
	},
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "creating a new agenda account",
	Long: `Command register is used to create a new user account. 
	You need to provide a username, a password, an email and a phone num. 
	For example:

	Agenda register -uABB -p123 -e123@163.com -n13579`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phonenum")
		if username == "" {
			fmt.Println("username can not be blank.")
			return
		}
		if password == "" {
			fmt.Println("password can not be blank.")
			return
		}
		if email == "" {
			fmt.Println("email can not be blank.")
			return
		}
		if phone == "" {
			fmt.Println("phone number can not be blank.")
			return
		}
		/*if err := user.RegisterUser(username, password, email, phone); err != nil {
			fmt.Println(err)
			os.Exit(6)
		} it is examined by server */
		data := url.Values{"username": {username}, "password": {password}, "email": {email}, "phone": {phone}}
		res, err := http.PostForm("http://127.0.0.1:8080/v1/users", data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer res.Body.Close()
		DealWithResponse(res)
		fmt.Println("a new account is registered named by " + username)

	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "to logout your agenda accout",
	Long: `Be sure that you have logined before you wanna logout.
	 For example:

	Agenda logout`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		res, err := http.Get("http://127.0.0.1:8080/v1/user/logout")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		DealWithResponse(res)
		fmt.Println("logout successfully")
	},
}

var usrDelCmd = &cobra.Command{
	Use:   "usrDel",
	Short: "to delete the current accout",
	Long: `Be sure that you have logined before the operation of deleting.
	 For example:

	Agenda usrDel`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usrDel called")
		res, _, err := goreq.New().Delete("http://127.0.0.1:8080/v1/users").SendRawString("delete current user and logout").End()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		DealWithResponse(res)
		fmt.Println("user is canceled successfully.")
	},
}

var usrSchCmd = &cobra.Command{
	Use:   "usrSch",
	Short: "listing all the users",
	Long: `It will list the username, email, phoneNumber of all the accouts.
	Be sure  your have logined before using this command.
	For example:

	Agenda usrSch `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usrSch called")
		res, err := http.Get("http://127.0.0.1:8080/v1/users")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		DealWithResponse(res)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := map[string]interface{}{}
		json.Unmarshal(body, &result)
		result2print, _ := json.MarshalIndent(result["Items"], "", "    ")
		fmt.Print(string(result2print) + "\n")
	},
}

func init() {
	//login
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("user", "u", "", "the username of the account you want to login")
	loginCmd.Flags().StringP("password", "p", "", "password of your account")
	//register a new user
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "", "your username which should be unique")
	registerCmd.Flags().StringP("password", "p", "", "your password ")
	registerCmd.Flags().StringP("email", "e", "", "your email")
	registerCmd.Flags().StringP("phonenum", "n", "", "the number of your telephone")
	//logout
	RootCmd.AddCommand(logoutCmd)
	//delete current user
	RootCmd.AddCommand(usrDelCmd)
	//list all users
	RootCmd.AddCommand(usrSchCmd)
}
