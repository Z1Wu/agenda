// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"log"

	entity "github.com/Z1Wu/agenda/entity"
	"github.com/spf13/cobra"
)

// quitCmd represents the quit command
var quitCmd = &cobra.Command{
	Use:   "quit -m [MeetingName]",
	Short: "quit the meeting with the name [MeetingName]",
	Long: `you can quit the meeting with the name of [MeetingName]:
	P.S: if there is no participators in this meeting,the meeting will be deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quit called")
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.AgendaStart() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}
		m, _ := cmd.Flags().GetString("Title")

		if entity.QuitMeeting(m) {
			debugLog.Println("Quit meeting successfully")
			fmt.Println("Quit meeting successfully")
		} else {
			debugLog.Println("Fail to quit meeting")
			fmt.Println("不存在该会议或者该会议不是本用户创建")
			fmt.Println("Fail to quit meeting")
		}
		entity.AgendaEnd()
	},
}

func init() {
	rootCmd.AddCommand(quitCmd)
	quitCmd.Flags().StringP("MeetingName", "m", "", "meeting name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
