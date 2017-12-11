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

	"github.com/spf13/cobra"
	"github.com/tpisntgod/Agenda/entity/user"
)

// usrSchCmd represents the usrSch command
var usrSchCmd = &cobra.Command{
	Use:   "usrSch",
	Short: "listing all the users",
	Long: `It will list the username, email, phoneNumber of all the accouts.
	Be sure  your have logined before using this command.
	For example:

./app usrSch.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("usrSch called")
		if !user.IsLogin() {
			fmt.Println("please login first!")
			os.Exit(1)
		}
		if err := user.ListUsers(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

	},
}

func init() {
	RootCmd.AddCommand(usrSchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usrSchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usrSchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//usrSchCmd.Flags().StringP("user", "u", "", "the username of the account you want to search")
}
