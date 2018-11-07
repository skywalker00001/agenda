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
	"github.com/7cthunder/agenda/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -u=[username] -p=[password] -e=[email] -t=[phone]",
	Short: "Register a new account with username, password, email and phone",
	Long: `Register a new account with username, password, email and phone:
1. If the username you enter has been registered, you should change another username`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		logger := entity.NewLogger("[register]")
		logger.Println("You are calling register -u=" + username + " -p=" + password + " -e=" + email + " -t=" + phone)

		instance := entity.GetStorage()

		if username == "" {
			logger.Println("ERROR: You do not enter username, please input again!")
			return
		}
		filter := func(u *entity.User) bool {
			return u.GetName() == username
		}
		ulist := instance.QueryUser(filter)
		if len(ulist) > 0 {
			logger.Println("ERROR: Duplicate username, please change another one!")
			return
		}
		if password == "" {
			logger.Println("ERROR: You do not enter password, please input again!")
			return
		}
		if email == "" {
			logger.Println("ERROR: You do not enter email, please input again!")
			return
		}
		if phone == "" {
			logger.Println("ERROR: You do not enter phone, please input again!")
			return
		}

		instance.CreateUser(*entity.NewUser(username, password, email, phone))
		logger.Println("Register new user successfully!")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("username", "u", "", "username message")
	registerCmd.Flags().StringP("password", "p", "", "password message")
	registerCmd.Flags().StringP("email", "e", "", "email message")
	registerCmd.Flags().StringP("phone", "t", "", "phone message")

}
