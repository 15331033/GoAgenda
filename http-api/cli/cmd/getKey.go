package cmd

import (
	"GoAgenda/http-api/cli/operation"
	"GoAgenda/http-api/cli/entity"
	"github.com/spf13/cobra"
)

var getKeyCmd = &cobra.Command{
	Use:   "gk",
	Short: "To get the key you use",
	Long:  `To get the key you use, the following parameters should be avaliable,u/username is your username,p/password is your password`,
	Run: func(cmd *cobra.Command, args []string) {
		//s := entity.GetStorage()
		//userList := s.QueryUser(getAll)
		username, err1 := cmd.Flags().GetString("username")
		password, err2 := cmd.Flags().GetString("password")
		
		outPutErr([]error{err1, err2})

		ifNull("key","id","oldPassword",username, password, "email")
		key := operation.GetKey()

		entity.Info.Println("your key infomation is ")
		entity.Info.Println(key)
	},
}

func init() {
	RootCmd.AddCommand(getKeyCmd)
	getKeyCmd.Flags().StringP("username", "u", "", "the username,it should not be empty")
	getKeyCmd.Flags().StringP("password", "p", "", "the password,it should not be empty")
}