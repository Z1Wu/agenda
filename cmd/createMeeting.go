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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "creat -m [MeetingName] -s [StartTime] -e [EndTime] -p [Participators]",
	Short: "Create a meeting",
	Long: `To create a new meeting with:
	[MeetingName] the name of the meeting
	[Participator] the Participator of the meeting,the Participator can only attend one meeting during one meeting time
	[StartTime] the StartTime of the meeting
	[EndTime] the EndTime of the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createMeeting called")
		debugLog := log.New(logFile, "[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		isLogin := entity.AgendaStart()
		defer entity.AgendaEnd()
		if isLogin == false {
			debugLog.Println("Fail, please log in")
			// fmt.Println("Fail, please log in")
			log.Fatal("Fail, please log in")
		}
		m, _ := cmd.Flags().GetString("MeetingName")
		p, _ := cmd.Flags().GetStringSlice("Participators")
		sd, _ := cmd.Flags().GetString("StartTime")
		ed, _ := cmd.Flags().GetString("EndTime")
		if entity.CreateMeeting(entity.CurrentUser.Name, m, sd, ed, p) {
			debugLog.Println("Create meeting successfully")
			fmt.Println("Create meeting successfully")
		} else {
			debugLog.Println("Fail to create meeting")
			fmt.Println("Fail to create meeting")
		}
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createMeetingCmd.Flags().StringP("MeetingName", "m", "", "meeting name")
	createMeetingCmd.Flags().StringP("StartTime", "s", "", "meeting's startDate")
	createMeetingCmd.Flags().StringP("EndTime", "e", "", "meeting's endDate")
	createMeetingCmd.Flags().StringSliceP("Participators", "p", []string{}, "meeting's participator")
}
