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

	"github.com/spf13/cobra"

	"github.com/Z1Wu/agenda/entity"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new user",
	// TODO
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		// 无论是否已经登陆都可以进行注册。

		u, _ := cmd.Flags().GetString("username")
		k, _ := cmd.Flags().GetString("key")
		e, _ := cmd.Flags().GetString("email")
		p, _ := cmd.Flags().GetString("phone")

		// 任何命令都要在结束之前退出
		defer entity.AgendaEnd()
		if entity.UserRegister(u, k, e, p) {
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

	registerCmd.Flags().StringP("username", "u", "", "new user's username")
	registerCmd.Flags().StringP("key", "k", "", "new user's password")
	registerCmd.Flags().StringP("email", "e", "", "new user's email")
	registerCmd.Flags().StringP("phone", "p", "", "phone new user's phone")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
