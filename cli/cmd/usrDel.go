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

// usrDelCmd represents the usrDel command
var usrDelCmd = &cobra.Command{
	Use:   "usrDel",
	Short: "to delete the current accout",
	Long: `Be sure that you have logined before the operation of deleting.
	 For example:

./app usrDel`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usrDel called")
		if !user.IsLogin() {
			fmt.Println("you have not logined yet.")
			os.Exit(1)
		}

		if err := user.DeleteUser(); err != nil {
			// fmt.Println("error happened.")
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("user is canceled successfully.")
	},
}

func init() {
	RootCmd.AddCommand(usrDelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usrDelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usrDelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
