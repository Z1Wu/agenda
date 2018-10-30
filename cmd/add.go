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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add -p [Participators] -m [MeetingName]",
	Short: "To add Participators of the meeting",
	Long: `Add [Participator] to the meeting with the title of [Title]:
	P.S: If the Participators cannot attend such meeting during the time, add will fail.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
	if entity.AgendaStart() == false {
		debugLog.Println("Fail,please log in")
		fmt.Println("Fail,please log in")
	}
	
	p, _ := cmd.Flags().GetStringSlice("Participators")
	n, _ := cmd.Flags().GetString("MeetingName")
	if entity.Addparticipator(n, p) {
		debugLog.Println("Add participators successfully")
		fmt.Println("Add participators successfully")
	} else {
		debugLog.Println("Fail to add participators")
		fmt.Println("Fail to add participators")
	}
	entity.AgendaEnd()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringSliceP("Participators", "p", []string{}, "meeting's participator")
	addCmd.Flags().StringP("MeetingName", "m", "", "meeting title")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
