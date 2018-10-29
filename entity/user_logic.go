package entity

// 处理用户相关逻辑
// AgendaStart invoked when start
func AgendaStart() bool {
	ReadFromFile()
	// ReadCurrentUser()
	// if CurrentUser.Name == "" {
	// 	return false
	// }
	return true
}

// AgendaEnd invoked when quit
func AgendaEnd() {
	WriteToFile()
	// writeCurrentUser()
}

// UserLogin if the username match password
//@param userName the username want to login
//@param password the password user enter
//@return if success, true will be returned
//登录命令不需要调用StartAgenda,但需要调用QuitAgenda来保存登录信息
func UserLogin(userName string, password string) bool {
	ReadFromFile()
	if CurrentUser.Name != "" {
		return false
	}
	filter := func(u *User) bool {
		return u.getName() == userName && u.getPassword() == password
	}

	ulist := queryUser(filter)
	if len(ulist) == 0 {
		return false
	} else {
		//当前用户信息
		CurrentUser = ulist[0]
		return true
	}
}

/**
 * regist a user
 * @param userName new user's username
 * @param password new user's password
 * @param email new user's email
 * @param phone new user's phone
 * @return if success, true will be returned
 */
func UserRegister(userName, password, email, phone string) bool {
	filter := func(u *User) bool {
		return u.getName() == userName
	}
	ulist := queryUser(filter)

	if len(ulist) == 0 {
		createUser(User{userName, password, email, phone})
		return true
	} else {
		return false
	}
}

// DeleteUser delete user
// @param userName user's username
// @param password user's password
// @return if success, true will be returned
func DeleteUser(userName string, password string) bool {
	uf := func(u *User) bool {
		return (u.getName() == userName) && (u.getPassword() == password)
	}
	mf := func(m *Meeting) bool {
		return m.getSponsor() == userName || m.isParticipator(userName)
	}
	if deleteUser(uf) != 0 {
		deleteMeeting(mf)
		if userName == CurrentUser.Name {
			CurrentUser.InitUser("", "", "", "")
		}
		return true
	} else {
		return false
	}
}

//ListAllUsers list all users from storage
//@return a user list result
func ListAllUsers() []User {
	filter := func(u *User) bool {
		return true
	}
	return queryUser(filter)
}
