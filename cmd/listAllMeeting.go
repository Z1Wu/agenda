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

// listAllMeetingCmd represents the listAllMeeting command
var listAllMeetingCmd = &cobra.Command{
	Use:   "listAllMeeting",
	Short: "List all meetings the sponsor created",
	Long:  `List all meetings the sponsor created`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listAllMeeting called")
		debugLog := log.New(logFile, "[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.AgendaStart() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in\n")
		}
		mm := entity.ListAllSponsorMeetings(entity.CurrentUser.Name)
		if len(mm)!=0 {
			debugLog.Println("List meeting successfully")
			fmt.Println("List meeting successfully\n")
			fmt.Println("Sponsor Title StartDate EndDate Participators")
			for i, m := range mm {
				fmt.Printf("%d. %s\n", i+1, m)
			}

		}
		entity.AgendaEnd()
	},
}

func init() {
	rootCmd.AddCommand(listAllMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
