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
	entity "github.com/7cthunder/agenda/entity"
	"github.com/spf13/cobra"
)

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "delm -t=[title]",
	Short: "Delete a meeting which current user created",
	Long: `You can delete a meetng you sponsor
1. Make sure you input the title of the meeting
2. Make sure you have sponsored the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := entity.NewLogger("[delm]")

		instance := entity.GetStorage()
		curU := instance.GetCurUser()
		title, _ := cmd.Flags().GetString("title")

		logger.Println("You are calling delm -t=", title)

		if curU.GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}

		if title == "" {
			logger.Println("ERROR: Please input the title of the meeting you want to delete")
			return
		}

		mfilter := func(m *entity.Meeting) bool {
			return curU.GetName() == m.GetSponsor() && title == m.GetTitle()
		}

		if instance.DeleteMeeting(mfilter) > 0 {
			logger.Println("Delete successfully!")
		} else {
			logger.Println("ERROR: you don't sponsor this meeting")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingCmd)
	deleteMeetingCmd.Flags().StringP("title", "t", "", "meeting title")
}
