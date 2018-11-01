package entity

import (
	"log"
)

// 处理用户相关逻辑
// AgendaStart invoked when start
func AgendaStart() bool {
	// 名字为空代表没有登陆
	CurrentUser.setName("")

	ReadFromFile()
	// fmt.Print("hello")
	// ReadCurrentUser()

	// 获取登陆信息
	curusername := ResotreLoginState()

	if curusername == "" {
		return false
	}

	// no error here
	CurrentUser = queryUser(func(u *User) bool {
		return u.getName() == curusername
	})[0]

	return true
}

// AgendaEnd invoked when quit
func AgendaEnd() {
	// 把磁盘文件读入
	WriteToFile()
	CacheLoginState()
}

// UserLogin if the username match password
//@param userName the username want to login
//@param password the password user enter
//@return if success, true will be returned
//登录命令不需要调用StartAgenda,但需要调用QuitAgenda来保存登录信息
func UserLogin(userName string, password string) bool {
	// ReadFromFile()
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

// UserLogout 用户登出
func UserLogout() bool {
	if CurrentUser.getName() != "" {
		return false
	} else {
		CurrentUser.setName("")
		return true
	}
}

// UserRegister
// @param userName new user's username
// @param password new user's password
// @param email new user's email
// @param phone new user's phone
// @return if success, true will be returned
func UserRegister(userName, password, email, phone string) bool {

	if userName == "" {
		log.Fatal("invalid user name")
	} else if password == "" {
		log.Fatal("invalid password")
	} else if email == "" {
		log.Fatal("invalid email")
	} else if phone == "" {
		log.Fatal("invalid phone")
	}
	filter := func(u *User) bool {
		return u.getName() == userName
	}
	ulist := queryUser(filter)

	// 检查合法性.

	// 查看是否出现重复
	if len(ulist) == 0 {
		createUser(User{userName, password, email, phone})
		return true
	}
	return false
}

// DeleteUser delete user
// @param userName user's username
// @param password user's password
// @return if success, true will be returned
// 删除用户， 如果有该用户创建的会议后者存在该用户的会议，则解散会议。
func DeleteUser(userName string, password string) bool {
	uf := func(u *User) bool {
		return (u.getName() == userName) && (u.getPassword() == password)
	}
	mf := func(m *Meeting) bool {
		if m.getSponsor() == userName {
			return true
		} else if m.isParticipator(userName) && len(m.Participators) == 1 {
			return true
		}
		return false
	}
	if deleteUser(uf) != 0 {
		deleteMeeting(mf)

		// 如果删除的是当前登陆的用户，需要登出
		if userName == CurrentUser.getName() {
			UserLogout()
		}
		return true
	}

	return false
}

//ListAllUsers list all users from storage
//@return a user list result
func ListAllUsers() []User {
	filter := func(u *User) bool {
		return true
	}
	return queryUser(filter)
}

// QueryUserByName return user with the specify name
func QueryUserByName(username string) []User {
	filter := func(u *User) bool {
		return u.getName() == username
	}
	return queryUser(filter)
}
