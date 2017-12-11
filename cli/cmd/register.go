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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "creating a new agenda account",
	Long: `Command register is used to create a new user account. 
	You need to provide a username, a password, an email and a phone num. 
	For example:

./app register -uABB -p123 -e123@163.com -n13579`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phonenum")
		if username == "" {
			fmt.Println("username can not be blank.")
			os.Exit(1)
		}
		if user.IsRegisteredUser(username) {
			fmt.Println("this username has existed.")
			os.Exit(2)
		}
		if password == "" {
			fmt.Println("password can not be blank.")
			os.Exit(3)
		}
		if email == "" {
			fmt.Println("email can not be blank.")
			os.Exit(4)
		}
		if phone == "" {
			fmt.Println("phone number can not be blank.")
			os.Exit(5)
		}
		if err := user.RegisterUser(username, password, email, phone); err != nil {
			fmt.Println(err)
			os.Exit(6)
		}
		fmt.Println("a new account is registered named by " + username)

	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("user", "u", "", "your username which should be unique")
	registerCmd.Flags().StringP("password", "p", "", "your password ")
	registerCmd.Flags().StringP("email", "e", "", "your email")
	registerCmd.Flags().StringP("phonenum", "n", "", "the number of your telephone")
}
