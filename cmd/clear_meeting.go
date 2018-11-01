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

// clearMeetingCmd represents the clearMeeting command
var clearMeetingCmd = &cobra.Command{
	Use:   "clm",
	Short: "Clear all meetings which current user created",
	Long: `Clear all meetings which current user created`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := entity.NewLogger("[clm]")
		logger.Println("You are calling clm")

		instance := entity.GetStorage()
		curU := instance.GetCurUser()

		if curU.GetName() == "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}

		mfilter := func(m *entity.Meeting) bool {
			return curU.GetName() == m.GetSponsor()
		}

		if instance.DeleteMeeting(mfilter) > 0 {
			logger.Println("Delete successfully!")
		} else {
			logger.Println("ERROR: The User don't sponsor any meeting")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearMeetingCmd)
}
