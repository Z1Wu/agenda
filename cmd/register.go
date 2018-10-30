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
	"log"

	"github.com/spf13/cobra"

	"github.com/Z1Wu/agenda/entity"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "regist -n [UserName] -c [PassWord] -e [Email] -t [Phone]",
	Short: "register a new user",
	// TODO
	Long: `To create a news user for the meeting agenda system. You have to come up with a username, your password, your email, your phone.
	[UserNme] is new user's name
	[PassWord] is new user's password
	[Email] is new user's email
	[Phone] is new user's phone.`,

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("register called")
		// 无论是否已经登陆都可以进行注册。
		entity.ReadFromFile()
		n, _ := cmd.Flags().GetString("username")
		c, _ := cmd.Flags().GetString("password")
		e, _ := cmd.Flags().GetString("email")
		t, _ := cmd.Flags().GetString("phone")

		// 任何命令都要在结束之前退出
		defer entity.AgendaEnd()
		if entity.UserRegister(n, c, e, t) {
			log.Print("Successful Rergister")
		} else {
			log.Fatal("Register fail")
		}

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	registerCmd.Flags().StringP("username", "n", "", "new user's username")
	registerCmd.Flags().StringP("password", "c", "", "new user's password")
	registerCmd.Flags().StringP("email", "e", "", "new user's email")
	registerCmd.Flags().StringP("phone", "t", "", "phone new user's phone")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
