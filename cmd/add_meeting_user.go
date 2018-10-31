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

	entity "github.com/7cthunder/agenda/entity"
	"github.com/spf13/cobra"
)

// addMeetingUserCmd represents the addMeetingUser command
var addMeetingUserCmd = &cobra.Command{
	Use:   "addmu",
	Short: "Add meeting members to the meeting which current user created",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		instance := entity.GetStorage()
		curU := instance.GetCurUser()
		title, _ := cmd.Flags().GetString("title")
		participators := cmd.Flags().Args()
		if curU.GetName() == "" {
			fmt.Println("You have not logged in yet, please log in first!")
		}

		if title == "" {
			fmt.Println("Please input the title of the meeting you want to delete")
		}

		if len(participators) == 0 {
			fmt.Println("you must add someone")
			return
		}

		for i := 0; i < len(participators); i++ {
			filter := func(u *entity.User) bool {
				return u.GetName() == participators[i]
			}
			if len(instance.QueryUser(filter)) == 0 {
				fmt.Println(participators[i] + " isn't existed")
			}
		}

		filter1 := func(m *entity.Meeting) bool {
			return curU.GetName() == m.GetSponsor() && title == m.GetTitle()
		}
		meeting := instance.QueryMeeting(filter1)
		if len(meeting) == 0 {
			fmt.Println("you don't sponsor this meeting")
			return
		}

		for i := 0; i < len(participators); i++ {
			for j := i + 1; j < len(participators); j++ {
				if participators[i] == participators[j] {
					fmt.Println("The participators you add can't repeat")
					return
				}
			}
		}

		for _, p := range participators {
			if meeting[0].IsParticipator(p) {
				fmt.Println(p + " is in the meeting")
				return
			}
		}

		for _, p := range participators {
			if curU.GetName() == p {
				fmt.Println("you add yourself wrongly")
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
				fmt.Println("There are conflicts between  participators' time and meeting's time ", p)
				return
			}
		}

		for _, p := range participators {
			mSwitch := func(m *entity.Meeting) {
				m.AddParticipator(p)
			}
			instance.UpdateMeeting(filter1, mSwitch)
		}
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
