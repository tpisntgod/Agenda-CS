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

// msCmd represents the ms command
var msCmd = &cobra.Command{
	Use:   "ms",
	Short: "to search meetings",
	Long: `to search those meetings in the time slot you provide
	For example:

./app ms -s"2017-10-28 09:30" -e"2017-10-28 10:30"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ms called")
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		if !user.IsLogin() {
			fmt.Println("Please login first!")
			os.Exit(1)
		}
		if stime == "" {
			fmt.Println("starttime can not be blank.The format is 2017-01-01 09:00")
			os.Exit(2)
		}
		if etime == "" {
			fmt.Println("endtime can not be blank.The format is 2017-01-01 09:00")
			os.Exit(3)
		}
		t1, _ := time.Parse("2006-01-02 15:04:05", stime)
		t2, _ := time.Parse("2006-01-02 15:04:05", etime)
		if !meeting.CheckStarttimelessthanEndtime(t1, t2) {
			fmt.Println("start time should be less than end time")
			os.Exit(4)
		}
		if err := meeting.QueryMeeting(t1, t2); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
	},
}

func init() {
	RootCmd.AddCommand(msCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// msCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// msCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	msCmd.Flags().StringP("stime", "s", "", "time when the meeting will begin")
	msCmd.Flags().StringP("etime", "e", "", "time when the meeting will end")
}
