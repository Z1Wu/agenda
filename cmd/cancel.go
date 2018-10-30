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

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel -m [MeetingName]",
	Short: "Cancel a meeting named [MeetingName]",
	Long:  `You can cancel one meeting with the name [MeetingName]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cancel called")
		debugLog := log.New(logFile, "[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.AgendaStart() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}
		n, _ := cmd.Flags().GetString("MeetingName")

		if entity.DeleteMeeting(entity.CurrentUser.Name, n) {
			debugLog.Println("Cancel meeting successfully")
			fmt.Println("Cancel meeting successfully")
		} else {
			debugLog.Println("Fail to Cancel meeting")
			fmt.Println("Fail to Cancel meeting")
		}
		entity.AgendaEnd()
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)
	cancelCmd.Flags().StringP("MeetingName", "m", "", "meeting name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
