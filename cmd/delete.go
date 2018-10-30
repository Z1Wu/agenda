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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete -n [UserName] -c [PassWord]",
	Short: "Delete a user",
	Long: `You can delete your account from the meeting management system' database.
	P.S: After your deleting, you won't be able to log in  with this username and password again any more.`,
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.AgendaStart() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}

		if entity.DeleteUser(entity.CurrentUser.Name, entity.CurrentUser.Password) {
			debugLog.Println("Delete this account successfully")
			fmt.Println("Delete this account successfully")
		} else {
			debugLog.Println("Fail to delete this account")
			fmt.Println("Fail to delete this account")
		}
		entity.AgendaEnd()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP("username", "n", "", "deleted user's username")
	deleteCmd.Flags().StringP("password", "c", "", "deleted user's password")
}
