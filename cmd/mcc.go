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

	"github.com/tpisntgod/Agenda/entity/meeting"
	"github.com/tpisntgod/Agenda/entity/user"

	"github.com/spf13/cobra"
)

// mccCmd represents the mcc command
var mccCmd = &cobra.Command{
	Use:   "mcc",
	Short: "to cancel a meeting that you sponsor",
	Long: `provide a title then that meeting will be canceled(Login first)
For example:

./app mcc -ttitle`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mcc called")
		title, _ := cmd.Flags().GetString("title")
		if !user.IsLogin() {
			fmt.Println("Please login first!")
			os.Exit(1)
		}
		if title == "" {
			fmt.Println("please input the title!")
			os.Exit(2)
		}
		if err := meeting.CancelMeeting(title); err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	},
}

func init() {
	RootCmd.AddCommand(mccCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mccCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mccCmd.Flags().StringP("title", "t", "", "title of the meeting you wanna delete")
}
