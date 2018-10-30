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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout the meeting system",
	Long: `After logging out the system, you can only register a new user or log in.
	register -n [UserName] -pass [PassWord] -e [Email]
	login -n [UserName] -pass [PassWord]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		debugLog := log.New(logFile, "[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.AgendaStart() == true {
			entity.CurrentUser.InitUser("", "", "", "")
			debugLog.Println("Log out successfully")
			fmt.Println("Log out successfully")
		} else {
			// debugLog := log.New(logFile, "[Result]", log.Ldate|log.Ltime|log.Lshortfile)
			debugLog.Println("You already logout")
			fmt.Println("You already logout")
		}
		defer entity.AgendaEnd()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
