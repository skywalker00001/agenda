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
	Use:   "qym",
	Short: "Query meetings start from [startTime] to [endTime]",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("queryMeeting called")
		startTime, _ := cmd.Flags().GetString("stime")
		endTime, _ := cmd.Flags().GetString("etime")

		if startTime == "" || endTime == "" {
			fmt.Println("You have not set the start time or end time of the meeting yet, please do it first!")
			return
		}

		instance := entity.GetStorage()

		if instance.GetCurUser().GetName() == "" {
			fmt.Println("You have not logged in yet, please log in first!")
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

		fmt.Println("Sponsor Title StartTime EndTime")

		for i, meeting := range mlist {
			fmt.Printf("Meeting%d: %s %s %s %s\n", i+1, meeting.GetSponsor(), meeting.GetTitle(), startTime, endTime)
			fmt.Printf("Meeting%d-Participators:", i+1)
			participators := meeting.GetParticipators()
			for j := 0; j < len(participators); j++ {
				fmt.Printf("%s", participators[j])
				if j != len(participators)-1 {
					fmt.Printf(", ")
				} else {
					fmt.Printf("\n")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)
	queryMeetingCmd.Flags().StringP("stime", "s", "", "start time of the meeting")
	queryMeetingCmd.Flags().StringP("etime", "e", "", "end time of the meeting")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
