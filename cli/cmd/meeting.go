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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bilibiliChangKai/Agenda-CS/cli/network/cookie"
	"github.com/spf13/cobra"
)

// mcCmd represents the mc command
var mcCmd = &cobra.Command{
	Use:   "mc",
	Short: "to create a new meeting",
	Long: `to create a new meeting with title, participator,starttime and endtime.
	 For example:

	Agenda mc -ttest -pPeter -pMarry -s"2017-10-28 09:30:00" -e"2017-10-28 10:30:00"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mc called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("parti")
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		if title == "" {
			fmt.Println("title can not be blank")
			return
		}
		if stime == "" {
			fmt.Println("stime can not be blank")
			return
		}
		if etime == "" {
			fmt.Println("etime can not be blank")
			return
		}
		t1, _ := time.Parse("2006-01-02 15:04:05", stime)
		t2, _ := time.Parse("2006-01-02 15:04:05", etime)
		if !CheckStarttimelessthanEndtime(t1, t2) {
			fmt.Println("start time should be less than end time")
			return
		}
		//data := url.Values{"title": {title}, "participators": participators, "stime": {stime}, "etime": {etime}}
		newMeeting, err := json.Marshal(struct {
			Title        string
			Host         string
			Participator []string
			StartTime    string
			EndTime      string
		}{
			Title:        title,
			Host:         title,
			Participator: participators,
			StartTime:    stime,
			EndTime:      etime})
		CheckPanic(err)
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://127.0.0.1:8080/v1/meetings", strings.NewReader(string(newMeeting)))
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		//res, err := http.PostForm("http://127.0.0.1:8080/v1/meetings", data)
		//res, _, err := goreq.New().Post("http://127.0.0.1:8080/v1/meetings").ContentType("json").SendStruct(data).End()
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		fmt.Println("create meeting" + title + "successfully")
	},
}

// apCmd represents the ap command
var apCmd = &cobra.Command{
	Use:   "ap",
	Short: "to add some participators to a meeting",
	Long: `to add some participators to a meeting with
	the title of the meeting and the name of the new participators.
	 For example:

	Agenda ap -ttitle -pPeter -pMarry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ap called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("parti")
		if title == "" {
			fmt.Println("please input title")
			return
		}
		destination := "http://127.0.0.1:8080/v1/meeting/" + title + "/adding-participators"
		//names := url.Values{"participators": participators}
		data, err := json.Marshal(struct {
			Participator []string
		}{
			Participator: participators})
		CheckPanic(err)
		client := &http.Client{}
		req, err := http.NewRequest("PATCH", destination, strings.NewReader(string(data)))
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		fmt.Println("meeting:" + title + " add participators successfully")
	},
}

// dpCmd represents the dp command
var dpCmd = &cobra.Command{
	Use:   "dp",
	Short: "to delete some participators to a meeting",
	Long: `to delete some participators to a meeting with
	the title of the meeting and the name of the new participators.
	 For example:

	./dpp dp -ttitle -pPeter -pMarry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dp called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("parti")
		if title == "" {
			fmt.Println("please input title")
			return
		}
		client := &http.Client{}
		destination := "http://127.0.0.1:8080/v1/meeting/" + title + "/deleting-participators"
		data, err := json.Marshal(struct {
			Participator []string
		}{
			Participator: participators})
		req, err := http.NewRequest("PATCH", destination, strings.NewReader(string(data)))
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
	},
}

// mccCmd represents the mcc command
var mccCmd = &cobra.Command{
	Use:   "mcc",
	Short: "to cancel a meeting that you sponsor",
	Long: `provide a title then that meeting will be canceled(Login first)
	For example:

	Agenda mcc -ttitle`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mcc called")
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("please input the title!")
			return
		}
		client := &http.Client{}
		destination := "http://127.0.0.1:8080/v1/users/cancel-a-meeting/" + title
		req, err := http.NewRequest("DELETE", destination, nil)
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
	},
}

// mclrCmd represents the mclr command
var mclrCmd = &cobra.Command{
	Use:   "mclr",
	Short: "to clear all meetings that you sponsor",
	Long: `to clear all meetings that you sponsor(Login first)
	For example:

	Agenda mclr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mclr called")
		client := &http.Client{}
		destination := "http://127.0.0.1:8080/v1/users/cancel-all-meeting"
		req, err := http.NewRequest("DELETE", destination, nil)
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
	},
}

// mqCmd represents the mq command
var mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "to quit a meeting",
	Long: `to quit a meeting whose title is provided by you.
	For example:

	Agenda mq -ttitle`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mq called")
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("title can not be blank!")
			return
		}
		client := &http.Client{}
		destination := "http://127.0.0.1:8080/v1/users/quit-meeting/" + title
		req, err := http.NewRequest("PATCH", destination, nil)
		CheckPanic(err)
		req.AddCookie(cookie.GetCookie())
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
	},
}

// msCmd represents the ms command
var msCmd = &cobra.Command{
	Use:   "ms",
	Short: "to search meetings",
	Long: `to search those meetings in the time slot you provide
	For example:

	Agenda ms -s"2017-10-28 09:30:00" -e"2017-10-28 10:30:00"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ms called")
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		if stime == "" {
			fmt.Println("starttime can not be blank.The format is 2017-01-01 09:00:00")
			return
		}
		if etime == "" {
			fmt.Println("endtime can not be blank.The format is 2017-01-01 09:00:00")
			return
		}
		t1, _ := time.Parse("2006-01-02 15:04:05", stime)
		t2, _ := time.Parse("2006-01-02 15:04:05", etime)
		if !CheckStarttimelessthanEndtime(t1, t2) {
			fmt.Println("start time should be less than end time")
			return
		}
		client := &http.Client{}
		starttime := strings.Split(stime, " ")
		endtime := strings.Split(etime, " ")
		destination := "http://127.0.0.1:8080/v1/users/query-meeting?starttime=" + starttime[0] + "%20" +
			starttime[1] + "&endtime=" + endtime[0] + "%20" + endtime[1]
		req, err := http.NewRequest("GET", destination, nil)
		req.AddCookie(cookie.GetCookie())
		res, err := client.Do(req)
		CheckPanic(err)
		defer res.Body.Close()
		DealWithResponse(res)
		body, err := ioutil.ReadAll(res.Body)
		CheckPanic(err)
		result := map[string]interface{}{}
		json.Unmarshal(body, &result)
		fmt.Print(result["Information"])
		/*
			body, err := ioutil.ReadAll(res.Body)
			CheckPanic(err)
			result := map[string]interface{}{}
			json.Unmarshal(body, &result)
			result2print, _ := json.MarshalIndent(result["Meetings"], "", "    ")
			fmt.Print(string(result2print) + "\n")*/
	},
}

func CheckStarttimelessthanEndtime(startTime time.Time, endTime time.Time) bool {
	if startTime.After(endTime) {
		return false
	}
	return true
}

func init() {
	//create a new meeting
	RootCmd.AddCommand(mcCmd)
	var strarr []string
	mcCmd.Flags().StringP("title", "t", "", "title of the meeting which should be unique")
	mcCmd.Flags().StringArrayP("parti", "p", strarr, "participators of the meeting ")
	mcCmd.Flags().StringP("stime", "s", "", "time when the meeting will begin")
	mcCmd.Flags().StringP("etime", "e", "", "time when the meeting will end")
	//add some new participators
	RootCmd.AddCommand(apCmd)
	apCmd.Flags().StringP("title", "t", "", "title of the meeting")
	apCmd.Flags().StringArrayP("parti", "p", strarr, "name of the participators you want to add ")
	//delete some new participators
	RootCmd.AddCommand(dpCmd)
	dpCmd.Flags().StringP("title", "t", "", "title of the meeting")
	dpCmd.Flags().StringArrayP("parti", "p", strarr, "name of the participators you want to delete ")
	//cancel a meeting
	RootCmd.AddCommand(mccCmd)
	mccCmd.Flags().StringP("title", "t", "", "title of the meeting you wanna delete")
	//clear all meetings
	RootCmd.AddCommand(mclrCmd)
	//quit a meeting
	RootCmd.AddCommand(mqCmd)
	mqCmd.Flags().StringP("title", "t", "", "the title of the meeting you wanna quit")
	//search meetings
	RootCmd.AddCommand(msCmd)
	msCmd.Flags().StringP("stime", "s", "", "time when the meeting will begin")
	msCmd.Flags().StringP("etime", "e", "", "time when the meeting will end")
}
