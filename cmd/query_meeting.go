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

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
	Use:   "qym -s=[startTime] -e=[endTime]",
	Short: "Query meetings start from [startTime] to [endTime]",
	Long: `You can query meetings start from [startTime] to [endTime]
	1. A string of Date has format YYYY-MM-DD/HH:mm
	2. The endTime must later than startTime`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := entity.NewLogger("[qym]")
		

		startTime, _ := cmd.Flags().GetString("stime")
		endTime, _ := cmd.Flags().GetString("etime")

		logger.Println("You are calling qym -s=", startTime, " -e=", endTime)
		
		if startTime == "" || endTime == "" {
			logger.Println("ERROR: You have not set the start time or end time of the meeting yet, please do it first!")
			return
		}

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}

		filter := func(m *entity.Meeting) bool {
			if !entity.StringToDate(endTime).IsLess(m.GetStartTime()) && !entity.StringToDate(startTime).IsGreater(m.GetEndTime()) {
				if instance.GetCurUser().GetName() == m.GetSponsor() || m.IsParticipator(instance.GetCurUser().GetName()) {
					return true
				}
			}
			return false
		}
		mlist := instance.QueryMeeting(filter)
		
		s := fmt.Sprintf("ID        Sponsor         Title           StartTime            EndTime\n")

		for i, meeting := range mlist {
			s = s + fmt.Sprintf("Meeting%d: %-15s %-15s %-20s %s\n", i+1, meeting.GetSponsor(), meeting.GetTitle(), 
										entity.DateToString(meeting.GetStartTime()), entity.DateToString(meeting.GetEndTime()))
			s = s + fmt.Sprintf("Meeting%d-Participators:", i+1)
			participators := meeting.GetParticipators()
			for j := 0; j < len(participators); j++ {
				s = s + fmt.Sprintf("%s", participators[j])
				if j != len(participators)-1 {
					s = s + fmt.Sprintf(", ")
				} else {
					s = s + fmt.Sprintf("\n\n")
				}
			}
		}
		logger.Printf("The result is: \n%s", s)
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)
	queryMeetingCmd.Flags().StringP("stime", "s", "", "start time of the meeting")
	queryMeetingCmd.Flags().StringP("etime", "e", "", "end time of the meeting")
}
