package cmd

import (
	"Agenda/123/cli/operation"
	"Agenda/123/cli/entity"
	"github.com/spf13/cobra"
)

var upDateCmd = &cobra.Command{
	Use:   "update",
	Short: "To update the information",
	Long:  `To update the information, the following parameters should be avaliable, id/UID is your user id, u/username is your new username, p/password is your new password, e/email is your new email`,
	Run: func(cmd *cobra.Command, args []string) {
		//s := entity.GetStorage()
		//userList := s.QueryUser(getAll)
		id, err1 := cmd.Flags().GetString("UID")
		username, err2 := cmd.Flags().GetString("username")
		password, err3 := cmd.Flags().GetString("password")
		email, err4 := cmd.Flags().GetString("email")
		
		outPutErr([]error{err1, err2, err3, err4})

		ifNull("key", id, "oldPassword", username, password, email)

		updateUser := operation.UpdateUser(id,username,password,email)
		entity.Info.Println("Update Success!")
		entity.Info.Println("Update information is\nuid: " + string(updateUser.UID) + "\nusername: " + updateUser.UserName + "\npassword: " + updateUser.Password + "\nemail: " + updateUser.Email + "\ncreatedTime: " + updateUser.Created.Format("2006-01-02 15:04:05"))
	},
}

func init() {
	RootCmd.AddCommand(upDateCmd)
	upDateCmd.Flags().StringP("UID", "id", "", "the id,it should not be empty")
	upDateCmd.Flags().StringP("username", "u", "", "the username,it should not be empty")
	upDateCmd.Flags().StringP("password", "p", "", "the password,it should not be empty")
	upDateCmd.Flags().StringP("email", "e", "", "the email,it should not be empty")
}