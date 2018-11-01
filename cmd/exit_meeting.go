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

// exitMeetingCmd represents the exitMeeting command
var exitMeetingCmd = &cobra.Command{
	Use:   "exitm -t=[title]",
	Short: "Exit a meeting which current users participated",
	Long: `Exit a meeting
	1. Make sure you enter a title for the meeting
	2. Make sure you have  participated in this meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := entity.NewLogger("[exitm]")
		
		title, _ := cmd.Flags().GetString("title")
		
		logger.Println("You are calling exitm -t=", title)
		
		if title == "" {
			logger.Println("ERROR: You have not set the title yet, please do it first!")
			return
		}

		instance := entity.GetStorage()
		if instance.GetCurUser().GetName() == "" {
			logger.Println("ERROR: You have not already logged in, please log in first!")
			return
		}

		filter1 := func(m *entity.Meeting) bool {
			if m.GetTitle() != title {
				return false
			}
			participators := m.GetParticipators()
			return m.GetSponsor() == instance.GetCurUser().GetName() || m.IsParticipator(instance.GetCurUser().GetName()) && len(participators) == 1
		}
		num1 := instance.DeleteMeeting(filter1)

		filter2 := func(m *entity.Meeting) bool {
			if m.GetTitle() != title {
				return false
			}
			return m.GetSponsor() == instance.GetCurUser().GetName() || m.IsParticipator(instance.GetCurUser().GetName())
		}
		switcher2 := func(m *entity.Meeting) {
			m.RemoveParticipator(instance.GetCurUser().GetName())
		}
		num2 := instance.UpdateMeeting(filter2, switcher2)

		if num1+num2 != 0 {
			logger.Println("Exit meeting successfully!")
		} else {
			logger.Println("ERROR: You have not participated in this meeting!")
		}
	},
}

func init() {
	rootCmd.AddCommand(exitMeetingCmd)
	exitMeetingCmd.Flags().StringP("title", "t", "", "exit meeting")
}
