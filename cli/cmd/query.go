// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"github.com/caijh23/GoAgenda/cli/operation"
	"github.com/caijh23/GoAgenda/cli/entity"
	"github.com/spf13/cobra"
	"strconv"
)

// registerCmd represents the register command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "To query all account",
	Long:  `To query all account and show the username, email and phone,the following parameters should be avaliable,k/key is your key`,
	Run: func(cmd *cobra.Command, args []string) {
		//s := entity.GetStorage()
		//userList := s.QueryUser(getAll)
		key, err := cmd.Flags().GetString("key")

		outPutErr([]error{err})

		userList := operation.ListAllUser(key)
		for _, v := range userList {
			entity.Info.Println("uid: " + strconv.Itoa(v.UID) + " username: " + v.UserName + " password: " + v.Password + " email: " + v.Email + " createdTime: " + v.Created.Format("2006-01-02 15:04:05"))
		}
	},
}

func init() {
	RootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringP("key", "k", "", "the key,it should not be empty")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}