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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "cm -t=[title] -s=[starttime] -e=[endtime] [participators]",
	Short: "Create a meeting with title, startTime, endTime and participators",
	Long: `Create a meeting with title, startTime, endTime and participators:
1. Make sure you have logged in first, and when you create meeting, you will be sponsor
2. 'Title' is unique of all meetings, if this title has been used, please change another one
3. Start time and end time should be in format with 'XXXX-XX-XX/XX:XX'
4. Participators cannot be empty
5. If sponsor's time or participator's time is conflict with their former meetings, 
the meeting will fail to create`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		startTimeS, _ := cmd.Flags().GetString("starttime")
		endTimeS, _ := cmd.Flags().GetString("endtime")
		ptcpt := cmd.Flags().Args()

		logger := entity.NewLogger("[cm]")
		logger.Println("You are calling cm -t="+title+" -s="+startTimeS+" -e="+endTimeS, ptcpt)

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}
		if title == "" {
			logger.Println("ERROR: You do not enter title, please input again!")
			return
		}
		if startTimeS == "" {
			logger.Println("ERROR: You do not enter start time, please input again!")
			return
		}
		if endTimeS == "" {
			logger.Println("ERROR: You do not enter end time, please input again!")
			return
		}
		if len(ptcpt) == 0 {
			logger.Println("ERROR: No participator!")
			return
		}

		sponsor := instance.GetCurUser().GetName()
		startTime := entity.StringToDate(startTimeS)
		endTime := entity.StringToDate(endTimeS)

		if !startTime.IsValid() {
			logger.Println("ERROR: Invalid start time!")
			return
		}
		if !endTime.IsValid() {
			logger.Println("ERROR: Invalid end time!")
			return
		}
		if startTime.IsGreaterThanEqual(endTime) {
			logger.Println("ERROR: Start time cannot be later or equal than end time!")
			return
		}

		mfilter1 := func(m *entity.Meeting) bool {
			return m.GetTitle() == title
		}
		if len(instance.QueryMeeting(mfilter1)) > 0 {
			logger.Println("ERROR: Duplicate title, please change it!")
			return
		}

		for _, p := range ptcpt {
			if p == sponsor {
				logger.Println("ERROR: Sponsor cannot be participator!")
				return
			}
		}

		for i := 0; i < len(ptcpt); i++ {
			for j := i + 1; j < len(ptcpt); j++ {
				if ptcpt[i] == ptcpt[j] {
					logger.Println("ERROR: Duplicate participators!")
					return
				}
			}
		}

		for _, p := range ptcpt {
			ufilter1 := func(u *entity.User) bool {
				return u.GetName() == p
			}
			if len(instance.QueryUser(ufilter1)) == 0 {
				logger.Println("ERROR: There is at least one non-existent participator!")
				return
			}
		}

		mfilter2 := func(m *entity.Meeting) bool {
			if m.GetSponsor() != sponsor && !m.IsParticipator(sponsor) {
				return false
			}
			if startTime.IsGreaterThanEqual(m.GetEndTime()) ||
				endTime.IsLessThanEqual(m.GetStartTime()) {
				return false
			}
			return true
		}
		if len(instance.QueryMeeting(mfilter2)) > 0 {
			logger.Println("ERROR: Sponsor's time conflict!")
			return
		}

		for _, p := range ptcpt {
			mfilter3 := func(m *entity.Meeting) bool {
				if m.GetSponsor() != p && !m.IsParticipator(p) {
					return false
				}
				if startTime.IsGreaterThanEqual(m.GetEndTime()) ||
					endTime.IsLessThanEqual(m.GetStartTime()) {
					return false
				}
				return true
			}
			if len(instance.QueryMeeting(mfilter3)) > 0 {
				logger.Println("ERROR: Participator's time conflict!")
				return
			}
		}
		logger.Println("Create meeting successfully!")
		instance.CreateMeeting(*entity.NewMeeting(sponsor, title, startTime, endTime, ptcpt))
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	createMeetingCmd.Flags().StringP("title", "t", "", "title of meeitng")
	createMeetingCmd.Flags().StringP("starttime", "s", "", "start time of meeting")
	createMeetingCmd.Flags().StringP("endtime", "e", "", "end time of meeting")
}
