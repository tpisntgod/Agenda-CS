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
	"fmt"
	"os"

	"github.com/tpisntgod/Agenda/entity/user"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login your agenda account",
	Long: `To login your account with your correct username and password.
	 For example:

./app login -uabb -p123`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		if user.IsLogin() {
			fmt.Println("You have already logined.Please logout first.")
			os.Exit(1)
		}
		if username == "" {
			fmt.Println("please input your username")
			os.Exit(2)
		}
		if password == "" {
			fmt.Println("please input your password")
			os.Exit(3)
		}
		if !user.IsRegisteredUser(username) {
			fmt.Println(username + " is not registered.")
			os.Exit(4)
		}
		if err := user.LoginUser(username, password); err != nil {
			//fmt.Println("the password is wrong.")
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println(username + " is logined")
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("user", "u", "", "the username of the account you want to login")
	loginCmd.Flags().StringP("password", "p", "", "password of your account")
}
