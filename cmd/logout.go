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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout your account if you have logined",
	Long: `Logout your account if you have logined:
1. If you did not logged in, you will fail to log out
2. If you log out, you can then only use 'register' and 'login' commands`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := entity.NewLogger("[logout]")
		logger.Println("You are calling logout")

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
		} else {
			instance.SetCurUser(*entity.NewUser("", "", "", ""))
			logger.Println("Log out successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

}
