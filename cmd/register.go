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

	"github.com/skywalker00001/agenda/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an account with username, password, email and phone",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		instance := entity.GetStorage()
		filter := func(u *entity.User) bool {
			return u.GetName() == username
		}

		ulist := instance.QueryUser(filter)
		if len(ulist) == 0 {
			instance.CreateUser(*entity.NewUser(username, password, email, phone))
			fmt.Println("Register new user successfully!")
		} else {
			fmt.Println("Duplicate username, please change another one!")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	// Here you will define your flags and configuration settings.
	registerCmd.Flags().StringP("username", "u", "Anonymous", "username message")
	registerCmd.Flags().StringP("password", "p", "null", "password message")
	registerCmd.Flags().StringP("email", "e", "null", "email message")
	registerCmd.Flags().StringP("phone", "t", "null", "phone message")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
