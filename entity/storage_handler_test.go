package entity

import (
	"fmt"
	"reflect"
	"testing"
)

var t_users = []User{
	{"a", "a", "a", "a"},
	{"b", "b", "b", "b"},
	{"c", "c", "c", "c"},
	{"aaa", "c", "c", "c"},
	{"bbb", "c", "c", "c"},
	{"ccc", "c", "c", "c"},
}
var t_meetings = []Meeting{
	{"a", "a_b", "2000-02-02/00:00", "2001-02-02/00:00", []string{"b"}},
	{"b", "b_c", "2002-02-02/00:00", "2003-02-02/00:00", []string{"c"}},
	{"c", "c", "2004-02-02/00:00", "2005-02-02/00:00", []string{}},
}

func init() {
	userList = nil
	meetingList = nil
}
func TestCreateUser(t *testing.T) {

	cases := []struct {
		in   User
		want []User
	}{
		{t_users[0], t_users[:1]},
		{t_users[1], t_users[:2]},
		{t_users[2], t_users[:3]},
	}
	for _, c := range cases {
		createUser(c.in)
		fmt.Printf("userList: %v\n", userList)
		if got := userList; !reflect.DeepEqual(got, c.want) {
			t.Errorf("CreateUser(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
func TestCreateMeeting(t *testing.T) {
	cases := []struct {
		in   Meeting
		want []Meeting
	}{
		{t_meetings[0], t_meetings[:1]},
		{t_meetings[1], t_meetings[:2]},
		{t_meetings[2], t_meetings[:3]},
	}
	for _, c := range cases {
		createMeeting(c.in)
		fmt.Printf("meetingList: %v\n", meetingList)
		if got := meetingList; !reflect.DeepEqual(got, c.want) {
			t.Errorf("CreateMeeting(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestQueryMeeting(t *testing.T) {
	cases := []struct {
		in   mFilter
		want []Meeting
	}{
		{func(m *Meeting) bool { return m.getSponsor() == "a" }, []Meeting{t_meetings[0]}},
		{func(m *Meeting) bool { return m.getSponsor() == "b" }, []Meeting{t_meetings[1]}},
		{func(m *Meeting) bool { return m.getSponsor() == "c" }, []Meeting{t_meetings[2]}},
	}
	for _, c := range cases {
		got := queryMeeting(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("queryMeeting(%v) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestQueryUser(t *testing.T) {
	cases := []struct {
		in   uFilter
		want []User
	}{
		{func(m *User) bool { return m.getName() == "a" }, []User{t_users[0]}},
		{func(m *User) bool { return m.getName() == "b" }, []User{t_users[1]}},
		{func(m *User) bool { return m.getName() == "c" }, []User{t_users[2]}},
	}
	for _, c := range cases {
		got := queryUser(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("queryMeeting(%v) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	cases := []struct {
		in   uFilter
		want []User
	}{
		{func(m *User) bool { return m.getName() == "a" }, []User{t_users[2], t_users[1]}},
		{func(m *User) bool { return m.getName() == "b" }, t_users[2:3]},
		{func(m *User) bool { return m.getName() == "c" }, []User{}},
	}
	for _, c := range cases {
		deleteUser(c.in)
		fmt.Printf("userList: %v\n", userList)
		if got := userList; !reflect.DeepEqual(got, c.want) {
			t.Errorf("DeleteUser(%v) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDeleteMeeting(t *testing.T) {
	cases := []struct {
		in   mFilter
		want []Meeting
	}{
		{func(m *Meeting) bool { return m.getSponsor() == "a" }, []Meeting{t_meetings[2], t_meetings[1]}},
		{func(m *Meeting) bool { return m.getSponsor() == "b" }, t_meetings[2:3]},
		{func(m *Meeting) bool { return m.getSponsor() == "c" }, []Meeting{}},
	}
	for _, c := range cases {
		deleteMeeting(c.in)
		fmt.Printf("meetingList: %v\n", meetingList)
		if got := meetingList; !reflect.DeepEqual(got, c.want) {
			t.Errorf("DeleteMeeting(%v) == %q, want %q", c.in, got, c.want)
		}
	}
}
