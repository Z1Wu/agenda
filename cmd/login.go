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
	"os"

	entity "github.com/Z1Wu/agenda/entity"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login -n [UserName] -c [PassWord]",
	Short: "Login to the meeting system.",
	Long: `Using your UserName and PassWord to login Agenda.
	P.S:If the PassWord is right,you can login Agenda and use it. If you forget the PassWord,you must register another one User`,

	PreRun: func(cmd *cobra.Command, args []string) {
		debugLog := log.New(logFile, "[Execute]", log.Ldate|log.Ltime|log.Lshortfile)
		debugLog.Printf("%v\n", os.Args[1:])
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		entity.AgendaStart()
		n, _ := cmd.Flags().GetString("username")
		c, _ := cmd.Flags().GetString("password")

		defer entity.AgendaEnd()
		if entity.UserLogin(n, c) {
			log.Print("Successfully Login")
		} else {
			log.Fatal("Login fail")
		}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "n", "", "logged user's username")
	loginCmd.Flags().StringP("password", "c", "", "logged user's password")
}
