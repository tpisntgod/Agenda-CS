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
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bilibiliChangKai/Agenda-CS/cli/network/cookie"
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
		if cookie.ExistCookie() {
			fmt.Println("please logout first")
			return
		}

		user, err := json.Marshal(struct {
			Name     string
			Password string
		}{
			Name:     username,
			Password: hashFunc(password)})
		CheckPanic(err)
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://127.0.0.1:8080/v1/user/login", strings.NewReader(string(user)))
		CheckPanic(err)
		//res, err := http.Post("http://127.0.0.1:8080/v1/user/login", bytes.NewBuffer(user))
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		CurCookies := res.Cookies()
		cookie.WriteCookie(CurCookies[0])
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
		// data := url.Values{"username": username, "password": {hashFunc(password)}, "email": {email}, "phone": {phone}}
		newUser, err := json.Marshal(struct {
			Name     string
			Password string
			Email    string
			Phone    string
		}{
			Name:     username,
			Password: hashFunc(password),
			Email:    email,
			Phone:    phone})
		CheckPanic(err)
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://127.0.0.1:8080/v1/users", strings.NewReader(string(newUser)))
		CheckPanic(err)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)
		CheckPanic(err)
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
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://127.0.0.1:8080/v1/user/logout", nil)
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		//res, err := http.Get("http://127.0.0.1:8080/v1/user/logout")
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		cookie.DeleteCookie()
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
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", "http://127.0.0.1:8080/v1/users", nil)
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		cookie.DeleteCookie()
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
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://127.0.0.1:8080/v1/users", nil)
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		body, err := ioutil.ReadAll(res.Body)
		CheckPanic(err)
		result := map[string]interface{}{}
		json.Unmarshal(body, &result)
		result2print, _ := json.MarshalIndent(result["Items"], "", "    ")
		fmt.Print(string(result2print) + "\n")
	},
}

//hash the password
func hashFunc(hashString string) string {
	h := md5.New()
	h.Write([]byte(hashString))
	cipheStr := h.Sum(nil)
	return hex.EncodeToString(cipheStr)
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
