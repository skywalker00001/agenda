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

// exitMeetingCmd represents the exitMeeting command
var exitMeetingCmd = &cobra.Command{
	Use:   "exitm",
	Short: "Exit a meeting which current users participated",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("exitMeeting called")
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("You have not set the title yet, please do it first!")
		}

		instance := entity.GetStorage()
		if instance.GetCurUser().GetName() == "" {
			fmt.Println("You have not already logged in, please log in first!")
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
			fmt.Println("Exit meeting successfully!")
		} else {
			fmt.Println("You have not participated in this meeting!")
		}
	},
}

func init() {
	rootCmd.AddCommand(exitMeetingCmd)
	exitMeetingCmd.Flags().StringP("title", "t", "", "exit meeting")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
