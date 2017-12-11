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
	"github.com/tpisntgod/Agenda/entity/meeting"
	"github.com/tpisntgod/Agenda/entity/user"
)

// mqCmd represents the mq command
var mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "to quit a meeting",
	Long: `to quit a meeting whose title is provided by you.
	For example:

./app mq -ttitle`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mq called")
		title, _ := cmd.Flags().GetString("title")
		if !user.IsLogin() {
			fmt.Println("Please login first!")
			os.Exit(1)
		}
		if title == "" {
			fmt.Println("title can not be blank!")
			os.Exit(2)
		}
		if err := meeting.QuitMeeting(title); err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	},
}

func init() {
	RootCmd.AddCommand(mqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mqCmd.Flags().StringP("title", "t", "", "the title of the meeting you wanna quit")
}
