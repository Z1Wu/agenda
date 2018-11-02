package entity

// storage handler, read from file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type uFilter func(*User) bool

// type uSwitcher func(*User)
type mFilter func(*Meeting) bool

// type mSwitcher func(*Meeting)

var meetingList []Meeting
var userList []User

// CurrentUser is the login user, defined when login function invoked
var CurrentUser User

// var my_logger = log.

// initfor test

// ReadFromFile read the user and meeting from file
// 把数据从文件中读取进入内存，包括用户信息和文件
func ReadFromFile() {
	//读user
	file1, err1 := os.Open("./data/users.json")
	defer file1.Close()
	if err1 != nil {
		// my_logger.Fatal("Failed to read file user.json")
		fmt.Fprintf(os.Stderr, "Fail to open ")
	}
	dec1 := json.NewDecoder(file1)
	err1 = dec1.Decode(&userList)
	if err1 != nil {
		// my_logger.Fatal("Fail to Decode USERS LIST")
		log.Print("DECODE USERLIST, EMPTY USERLIST")
		// log.Fatal(err1)
	}
	//读Meeting
	file2, err2 := os.Open("./data/meetings.json")
	defer file2.Close()
	if err2 != nil {
		// my_logger.Fatal("Fail to Decode Metting list")
		fmt.Fprintf(os.Stderr, "Fail to open file meetings.json")
		log.Fatal(err2)
	}
	dec2 := json.NewDecoder(file2)
	// should pass the argument as a pointer
	err2 = dec2.Decode(&meetingList)
	if err2 != nil {
		log.Print("DECODE MEETING LIST, EMPTY MEETINGLIST")
		log.Fatal(err2)
	}

	// 用户的登陆状态
	return
}

// WriteToFile write all users and meetings to file
// 把用户信息，会议信息和当前用户的登陆状态写入文件中
// TODO
func WriteToFile() {
	file1, err1 := os.Create("./data/users.json")
	defer file1.Close()
	if err1 != nil {
		// my_logger.Fatal("Fail to create UserInfo")
		fmt.Fprintf(os.Stderr, "Fail to create UserInfo")
	}
	enc1 := json.NewEncoder(file1)
	if err1 := enc1.Encode(&userList); err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to encode")
	}
	//写Meeting
	file2, err2 := os.Create("./data/meetings.json")
	defer file2.Close()
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Fail to create MeetingInfo")
	}
	enc2 := json.NewEncoder(file2)
	if err2 := enc2.Encode(&meetingList); err2 != nil {
		fmt.Fprintf(os.Stderr, "Fail to encode")
	}

}

// 把当前登陆的用户名写入文件中
func CacheLoginState() bool {
	// os.Create("./data/cahce.json")
	cache, err := os.Create("data/cache.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(cache, CurrentUser.getName())
	// writer  := bufio.Writer()
	return true
}

// 恢复登陆用户状态
func ResotreLoginState() (username string) {
	username = ""
	cache, err := os.Open("data/cache.json")
	defer cache.Close()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(cache)

	username, err = reader.ReadString('\n')
	// remove the last newline character
	username = strings.TrimSuffix(username, "\n")

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

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

func deleteMeeting(filter mFilter) int {
	n := 0
	for i, m := range meetingList {
		if filter(&m) {
			meetingList[i] = meetingList[len(meetingList)-1-n]
			n++
		}
	}
	// 更改meeting list 中的内容
	meetingList = meetingList[:len(meetingList)-n]
	return n
}
