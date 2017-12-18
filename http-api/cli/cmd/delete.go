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
	"GoAgenda/http-api/cli/operation"
	"GoAgenda/http-api/cli/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete one account",
	Long:  `To delete one account,the following parameters should be avaliable,i/UID is the id you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		//s := entity.GetStorage()
		//userList := s.QueryUser(getAll)
		id, err := cmd.Flags().GetString("UID")

		outPutErr([]error{err})

		operation.DeleteUser(id)

		entity.Info.Println("Delete Success!")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("UID", "i", "", "the id,it should not be empty")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}