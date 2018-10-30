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
	"log"

	"github.com/7cthunder/agenda/entity"

	"github.com/spf13/cobra"
)

// deleteMeetingUserCmd represents the deleteMeetingUser command
var deleteMeetingUserCmd = &cobra.Command{
	Use:   "delmu",
	Short: "Delete meeting members from the meeting which current user created",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators := cmd.Flags().Args()

		storage := entity.GetStorage()

		if storage.GetCurUser().GetName() == "" {
			log.Println("You have not logged in yet, please log in first!")
			return
		}

		if title == "" {
			log.Println("You do not enter a title for meeting, please enter it!")
			return
		}

		mfilter := func(m *entity.Meeting) bool {
			return m.GetTitle() == title
		}

		meetings := storage.QueryMeeting(mfilter)
		if len(meetings) == 0 {
			log.Println("This meeting is not existed, please enter a correct title!")
			return
		}

		for _, p := range participators {
			isInMeeting := false
			for _, u := range meetings[0].GetParticipators() {
				if u == p {
					isInMeeting = true
				}
			}
			if !isInMeeting {
				log.Println(p, "is not in this meeting, please check the participators list of this meeting!")
				return
			}
			if p == meetings[0].GetSponsor() {
				log.Println("You can't delete yourself from this meeting for you're the sponsor for it!\n If you want to delete this meeting, please use command 'delm'!")
				return
			}
		}

		storage.UpdateMeeting(
			mfilter,
			func(m *entity.Meeting) {
				for _, u := range participators {
					m.RemoveParticipator(u)
				}
			})

		meetings = storage.QueryMeeting(mfilter)

		if len(meetings[0].GetParticipators()) == 0 {
			storage.DeleteMeeting(func(m *entity.Meeting) bool {
				return m.GetTitle() == title
			})
		}

		log.Println("You have successfully removed your designated participant from this meeting!")

	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingUserCmd)

	deleteMeetingUserCmd.Flags().StringP("title", "t", "", "title of meeting")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteMeetingUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteMeetingUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
