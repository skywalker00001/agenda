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

// queryUserCmd represents the qyu command
var queryUserCmd = &cobra.Command{
	Use:   "qyu",
	Short: "Query all users if you have logined",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("queryUser called")

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() == "" {
			fmt.Println("You have not logged in yet, please log in first!")
			return
		}

		filter := func(u *entity.User) bool {
			return true
		}
		ulist := instance.QueryUser(filter)

		fmt.Println("Name Email Telephone")

		for i, user := range ulist {
			fmt.Printf("User%d: %s %s %s\n", i+1, user.Name, user.Email, user.Phone)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
