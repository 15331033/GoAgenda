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
	"GoAgenda/cli/entity"
	"GoAgenda/cli/operation"
	"github.com/spf13/cobra"
	"strconv"
	"os"
)


// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "To register an account",
	Long:  `To regist your account,the following parameters should be avaliable,k/key is your key,u/user is your username,p/password is your password,e/email is your email`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err1 := cmd.Flags().GetString("key")
		username, err2 := cmd.Flags().GetString("user")
		password, err3 := cmd.Flags().GetString("password")
		email, err4 := cmd.Flags().GetString("email")
		//phone, err4 := cmd.Flags().GetString("phone")
		
		outPutErr([]error{err1, err2, err3, err4})

		ifNull(key, "id", "oldPassword", username, password, email)
		//ifExist(username)
		//var tempUser entity.User
		tempUser := operation.CreateUser(key, username, password, email)
		//tempUser.SetUser(username, password, email, phone)
		//s := entity.GetStorage()
		//s.CreatUser(tempUser)

		entity.Info.Println("Register Success!")
		entity.Info.Println("Your information is\nuid: " + strconv.Itoa(tempUser.UID) + "\nusername: " + tempUser.UserName + "\npassword: " + tempUser.Password + "\nemail: " + tempUser.Email + "\ncreatedTime: " + tempUser.Created.Format("2006-01-02 15:04:05"))
	},
}

func getAll(u entity.User) bool { return true }

func switcher(u *entity.User) { u.SetPassword("laji") }

func ifNull(key, id, oldPassword, username, password, email string) {
	i := 0
	err := ""
	if key == "" {
		err += " key"
		i += 1
	}
	if id == "" {
		err += " id"
		i += 1
	}
	if username == "" {
		err += " name"
		i += 1
	}
	if password == "" {
		err += " password"
		i += 1
	}
	if email == "" {
		err += " email"
		i += 1
	}
	if i != 0 {
		entity.Error.Println("Fail! Lack of" + err)
		os.Exit(2)
	}
}

func ifExist(u string) {
	s := entity.GetStorage()
	userList := s.QueryUser(getAll)
	for _, v := range userList {
		if (&v).GetUsername() == u {
			entity.Error.Println("Register Fail!")
			entity.Error.Println("Your name " + u + " is already exist.")
			os.Exit(3)
		}
	}
}

func outPutErr(errs []error) {
	for _, value := range errs {
		if value != nil {
			entity.Error.Println(value)
			entity.Error.Println("Register Fail!")
			os.Exit(1)
		}
	}
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("key", "k", "", "the key,it should not be empty")
	registerCmd.Flags().StringP("user", "u", "", "the user,it should not be empty")
	registerCmd.Flags().StringP("password", "p", "", "the password,it should not be empty")
	registerCmd.Flags().StringP("email", "e", "", "the email,it should not be empty")
	//registerCmd.Flags().StringP("phone", "t", "", "your phone,it should not be empyt")
	// Here you will define your flags and configuration settings.
}