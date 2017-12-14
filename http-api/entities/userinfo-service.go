package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

func (*UserInfoAtomicService) CreatUser(u *UserInfo) *UserInfo {
	session := myEngine.NewSession()
	defer session.Close()

	err := session.Begin()
	checkErr(err)

	_, err = session.Insert(u)
	checkErr(err)
	if err == nil {
		session.Commit()
		return u
	} else {
		session.Rollback()
	}
	return nil
}
func (*UserInfoAtomicService) UpdateUser(u *UserInfo) *UserInfo  {
	session := myEngine.NewSession()
	defer session.Close()

	err := session.Begin()
	checkErr(err)

	_, err = session.Where("UID = ?", u.UID).Update(u)
	checkErr(err)
	if err == nil {
		session.Commit()
		return u
	} else {
		session.Rollback()
	}
	return nil
}
func (*UserInfoAtomicService) DeleteUser(id int) error {
	session := myEngine.NewSession()
	defer session.Close()

	err := session.Begin()
	checkErr(err)

	_, err = session.Exec("delete from UserInfo where UID = ?", id)
	if err != nil {
		session.Rollback()
		return err
	} else {
		err = session.Commit()
		return err
	}
}
// FindAll .
func (*UserInfoAtomicService) FindAllUser() []UserInfo {
	var users []UserInfo
	err := myEngine.Find(&users)
	checkErr(err)
	return users
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var user UserInfo
	_, err := myEngine.Id(id).Get(&user)
	checkErr(err)
	return &user
}