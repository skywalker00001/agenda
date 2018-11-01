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

// addMeetingUserCmd represents the addMeetingUser command
var addMeetingUserCmd = &cobra.Command{
	Use:   "addmu -t=[title] [participators]",
	Short: "Add meeting members to the meeting which current user created",
	Long: `Add meeting members into the meeting:
	1. Make sure you have sponsored the meeting with the title
	2. Make sure the participators have not repeat and have not been in the meeting
	3. Make sure there aren't conflicts between  participators' time and meeting's time `,
	Run: func(cmd *cobra.Command, args []string) {
		logger := entity.NewLogger("[addmu]")

		instance := entity.GetStorage()
		curU := instance.GetCurUser()
		title, _ := cmd.Flags().GetString("title")
		participators := cmd.Flags().Args()
		
		logger.Println("You are calling addmu -t=" + title + " ", participators)
		
		if curU.GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}

		if title == "" {
			logger.Println("ERROR: Please input the title of the meeting you want to delete")
			return
		}

		if len(participators) == 0 {
			logger.Println("ERROR: You must add someone")
			return
		}

		for i := 0; i < len(participators); i++ {
			filter := func(u *entity.User) bool {
				return u.GetName() == participators[i]
			}
			if len(instance.QueryUser(filter)) == 0 {
				logger.Println("ERROR: " + participators[i] + " isn't existed")
				return
			}
		}

		filter1 := func(m *entity.Meeting) bool {
			return curU.GetName() == m.GetSponsor() && title == m.GetTitle()
		}
		meeting := instance.QueryMeeting(filter1)
		if len(meeting) == 0 {
			logger.Println("ERROR: You don't sponsor this meeting")
			return
		}

		for i := 0; i < len(participators); i++ {
			for j := i + 1; j < len(participators); j++ {
				if participators[i] == participators[j] {
					logger.Println("ERROR: The participators you add can't repeat")
					return
				}
			}
		}

		for _, p := range participators {
			if meeting[0].IsParticipator(p) {
				logger.Println("ERROR: " + p + " is in the meeting")
				return
			}
		}

		for _, p := range participators {
			if curU.GetName() == p {
				logger.Println("ERROR: You add yourself wrongly")
				return
			}
		}

		startTime := meeting[0].GetStartTime()
		endTime := meeting[0].GetEndTime()
		for _, p := range participators {
			filter2 := func(m *entity.Meeting) bool {
				mST := m.GetStartTime()
				mET := m.GetEndTime()
				if (m.IsParticipator(p) || m.GetSponsor() == p) &&
					((startTime.IsGreaterThanEqual(mST) && startTime.IsLess(mET)) ||
						(endTime.IsGreater(mST) && endTime.IsLessThanEqual(mET)) ||
						(startTime.IsLessThanEqual(mST) && endTime.IsGreaterThanEqual(mET))) {
					return true
				}
				return false
			}
			if len(instance.QueryMeeting(filter2)) > 0 {
				logger.Println("ERROR: There are conflicts between  " + p + "'s time and meeting's time ")
				return
			}
		}

		for _, p := range participators {
			mSwitch := func(m *entity.Meeting) {
				m.AddParticipator(p)
			}
			instance.UpdateMeeting(filter1, mSwitch)
		}
		logger.Println("addmu successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addMeetingUserCmd)

	// Here you will define your flags and configuration settings.
	addMeetingUserCmd.Flags().StringP("title", "t", "", "meeting title")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addMeetingUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addMeetingUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
