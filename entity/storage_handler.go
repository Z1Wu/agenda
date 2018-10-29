package entity

// storage handler, read from file

type uFilter func(*User) bool

// type uSwitcher func(*User)
type mFilter func(*Meeting) bool

// type mSwitcher func(*Meeting)

var meetingList []Meeting
var userList []User

// CurrentUser is the login user, defined when login function invoked
var CurrentUser User

// ReadFromFile read the user and meeting from file
// TODO
func ReadFromFile() {
	return
}

// WriteToFile write all users and meetings to file
// TODO
func WriteToFile() {
	return
}

func createUser(tUser User) {
	userList = append(userList, tUser)
}

func createMeeting(tMeeting Meeting) {
	meetingList = append(meetingList, tMeeting)
}

func queryUser(filter uFilter) []User {
	var dy []User
	for _, u := range userList {
		if filter(&u) {
			dy = append(dy, u)
		}
	}
	return dy
}

// func updateUser(filter uFilter, switcher uSwitcher) int {
// 	n := 0
// 	for _, u := range userList {
// 		if filter(&u) {
// 			switcher(&u)
// 			n++
// 		}
// 	}
// 	return n
// }

func deleteUser(filter uFilter) int {
	n := 0
	for i, u := range userList {
		if filter(&u) {
			userList[i] = userList[len(userList)-1-n]
			n++
		}
	}
	userList = userList[:len(userList)-n]
	return n
}

func queryMeeting(filter mFilter) []Meeting {
	var dy []Meeting
	for _, m := range meetingList {
		if filter(&m) {
			dy = append(dy, m)
		}
	}
	return dy
}

// func updateMeeting(filter mFilter, switcher mSwitcher) int {
// 	n := 0
// 	for _, m := range meetingList {
// 		if filter(&m) {
// 			switcher(&m)
// 			n++
// 		}
// 	}
// 	return n
// }

func deleteMeeting(filter mFilter) int {
	n := 0
	for i, m := range meetingList {
		if filter(&m) {
			meetingList[i] = meetingList[len(meetingList)-1-n]
			n++
		}
	}
	meetingList = meetingList[:len(meetingList)-n]
	return n
}
