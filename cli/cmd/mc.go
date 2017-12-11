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
	"time"

	"github.com/tpisntgod/Agenda/entity/meeting"
	"github.com/tpisntgod/Agenda/entity/user"

	"github.com/spf13/cobra"
)

// mcCmd represents the mc command
var mcCmd = &cobra.Command{
	Use:   "mc",
	Short: "to create a new meeting",
	Long: `to create a new meeting with title, participator,starttime and endtime.
	 For example:

./app mc -ttest -pPeter -pMarry -s"2017-10-28 09:30:00" -e"2017-10-28 10:30:00"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mc called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("parti")
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		if !user.IsLogin() {
			fmt.Println("Login first!")
			os.Exit(1)
		}
		if title == "" {
			fmt.Println("title can not be blank")
			os.Exit(2)
		}
		if stime == "" {
			fmt.Println("stime can not be blank")
			os.Exit(3)
		}
		if etime == "" {
			fmt.Println("etime can not be blank")
			os.Exit(4)
		}
		t1, _ := time.Parse("2006-01-02 15:04:05", stime)
		t2, _ := time.Parse("2006-01-02 15:04:05", etime)
		if !meeting.CheckStarttimelessthanEndtime(t1, t2) {
			fmt.Println("start time should be less than end time")
			os.Exit(5)
		}
		isUsersUnregistered := 0
		for i := 0; i < len(participators); i++ {
			if !user.IsRegisteredUser(participators[i]) {
				isUsersUnregistered = 1
				fmt.Println(participators[i] + " is not registered")
			}
		}
		if isUsersUnregistered == 1 {
			os.Exit(6)
		}
		if err := meeting.CreateMeeting(title, participators, t1, t2); err != nil {
			fmt.Println(err)
			os.Exit(7)
		}
		//MeetingCreate(title, participators, stime, etime)
		fmt.Println(title + " created")

	},
}

func init() {
	RootCmd.AddCommand(mcCmd)
	var strarr []string

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mcCmd.Flags().StringP("title", "t", "", "title of the meeting which should be unique")
	mcCmd.Flags().StringArrayP("parti", "p", strarr, "participators of the meeting ")
	mcCmd.Flags().StringP("stime", "s", "", "time when the meeting will begin")
	mcCmd.Flags().StringP("etime", "e", "", "time when the meeting will end")
}
