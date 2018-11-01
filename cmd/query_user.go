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

	"github.com/7cthunder/agenda/entity"
	"github.com/spf13/cobra"
)

// queryUserCmd represents the qyu command
var queryUserCmd = &cobra.Command{
	Use:   "qyu",
	Short: "Query all users if you have logined",
	Long: `Query all users if you have logined:
1. This command can be used only after you have logged in
2. The result will show you all registered users with format 'Name Email Telephone'`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := entity.NewLogger("[qyu]")
		logger.Println("You are calling qyu")

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}

		filter := func(u *entity.User) bool {
			return true
		}
		ulist := instance.QueryUser(filter)

		s := fmt.Sprintf("ID        Name                 Email                Phone\n")

		for i, user := range ulist {
			s = s + fmt.Sprintf("User%d:    %-20s %-20s %-20s\n", i+1, user.Name, user.Email, user.Phone)
		}
		logger.Printf("The result is: \n%s", s)
	},
}

func init() {
	rootCmd.AddCommand(queryUserCmd)
}
