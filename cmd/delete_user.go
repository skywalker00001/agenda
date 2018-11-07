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

// deleteUserCmd represents the delu command
var deleteUserCmd = &cobra.Command{
	Use:   "delu",
	Short: "Delete your account if you have logined",
	Long: `Delete your account if you have logined:
1. Please make sure you have logged in first
2. You will delete the current user logged in. By the way, meetings which this user sponsored will be dissolved,
meetings which this user participated will delete its name from paticipators set, if paticipators count is 0 after deleting,
the meeting will also be dissolved`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := entity.NewLogger("delu")
		logger.Println("You are calling delu")

		instance := entity.GetStorage()
		curU := instance.GetCurUser()

		if curU.GetName() != "" {
			logger.Println("ERROR: You have not logged in yet, please log in first!")
			return
		}

		ufilter := func(u *entity.User) bool {
			return u.GetName() == curU.GetName()
		}

		mfilter1 := func(m *entity.Meeting) bool {
			return m.GetSponsor() == curU.GetName()
		}
		mfilter2 := func(m *entity.Meeting) bool {
			return m.IsParticipator(curU.GetName())
		}
		mswitcher := func(m *entity.Meeting) {
			m.RemoveParticipator(curU.GetName())
		}

		if instance.DeleteUser(ufilter) > 0 {
			instance.DeleteMeeting(mfilter1)
			if instance.UpdateMeeting(mfilter2, mswitcher) > 0 {
				mfilter3 := func(m *entity.Meeting) bool {
					return len(m.GetParticipators()) == 0
				}
				instance.DeleteMeeting(mfilter3)
			}
			instance.SetCurUser(*entity.NewUser("", "", "", ""))
			logger.Println("Delete successfully!")

		} else {
			logger.Println("ERROR: Fail to delete!")
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
