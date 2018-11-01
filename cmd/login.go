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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login -u=[username] -p=[password]",
	Short: "Login with your username and password",
	Long: `Login with your username and password:
1. If you have logged in, you should log out first
2. Make sure the username entered has been registered before, or if you do not have an account, you should register first
3. Make sure you enter the right username and password, or you will fail to login`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		logger := entity.NewLogger("[login]")
		logger.Println("You are calling login -u=" + username + " -p=" + password)

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() != "" {
			logger.Println("ERROR: You have already logged in, please log out first!")
			return
		}

		if username == "" {
			logger.Println("ERROR: You do not enter username, please input again!")
			return
		}
		if password == "" {
			logger.Println("ERROR: You do not enter password, please input again!")
			return
		}

		filter := func(u *entity.User) bool {
			return u.GetName() == username && u.GetPassword() == password
		}

		ulist := instance.QueryUser(filter)

		if len(ulist) == 0 {
			logger.Println("ERROR: Wrong username or password!")
		} else {
			instance.SetCurUser(ulist[0])
			logger.Println("Log in successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "login name")
	loginCmd.Flags().StringP("password", "p", "", "login password")
}
