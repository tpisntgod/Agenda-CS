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

// apCmd represents the ap command
var apCmd = &cobra.Command{
	Use:   "ap",
	Short: "to add some participators to a meeting",
	Long: `to add some participators to a meeting with 
	the title of the meeting and the name of the new participators.
	 For example:

./app ap -ttitle -pPeter -pMarry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ap called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("parti")
		if !user.IsLogin() {
			fmt.Println("you have not logined yet")
			os.Exit(1)
		}
		if title == "" {
			fmt.Println("please input title")
			os.Exit(2)
		}
		isUsersUnregistered := 0
		for i := 0; i < len(participators); i++ {
			if !user.IsRegisteredUser(participators[i]) {
				isUsersUnregistered = 1
				fmt.Println(participators[i] + " is not registered")
			}
		}
		if isUsersUnregistered == 1 {
			os.Exit(3)
		}
		if err := meeting.AddMeetingParticipators(title, participators); err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	},
}

func init() {
	RootCmd.AddCommand(apCmd)
	var strarr []string
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	apCmd.Flags().StringP("title", "t", "", "title of the meeting")
	apCmd.Flags().StringArrayP("parti", "p", strarr, "name of the participators you want to add ")
}
