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

// listAllUsersCmd represents the listAllUsers command
var listAllUsersCmd = &cobra.Command{
	Use:   "listAllUser",
	Short: "List all users' name",
	Long:  `You can query all the users's names who have registed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listAllUsers called")
		debugLog := log.New(logFile, "[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		defer entity.AgendaEnd()
		if entity.AgendaStart() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
			return
		}
		uu := entity.ListAllUsers()
		fmt.Println("Name Email Telephone")
		for i, u := range uu {
			fmt.Printf("%d. %s %s %s\n", i+1, u.Name, u.Email, u.Phone)
		}
	},
}

func init() {
	rootCmd.AddCommand(listAllUsersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllUsersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllUsersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
