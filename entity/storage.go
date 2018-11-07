package entity

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

// Storage an entity used to read user list and meeting list
// from disk into memory
type Storage struct {
	userList    []User
	meetingList []Meeting
	curUser     User
}

type uFilter func(*User) bool
type uSwitcher func(*User)
type mFilter func(*Meeting) bool
type mSwitcher func(*Meeting)

var instance *Storage
var mu sync.Mutex

// GetStorage *
func GetStorage() *Storage {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Storage{}
			readFromFile()
		}
	}
	return instance
}

// read all users and meetings into instance
// read current user from curUser.txt
func readFromFile() {
	ufile, _ := os.OpenFile("./data/User.json", os.O_CREATE|os.O_RDONLY, 0666)
	defer ufile.Close()
	dec := json.NewDecoder(ufile)
	if err := dec.Decode(&(instance.userList)); err != nil && err != io.EOF {
		log.Fatal(err)
	}

	mfile, _ := os.OpenFile("./data/Meeting.json", os.O_CREATE|os.O_RDONLY, 0666)
	defer mfile.Close()
	dec = json.NewDecoder(mfile)
	if err := dec.Decode(&(instance.meetingList)); err != nil && err != io.EOF {
		log.Fatal(err)
	}

	cufile, _ := os.OpenFile("./data/curUser.txt", os.O_CREATE|os.O_RDONLY, 0666)
	defer cufile.Close()
	dec = json.NewDecoder(cufile)
	if err := dec.Decode(&(instance.curUser)); err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

// write updated users' and meetings' information as json into file
// write current user into curUser.txt
func writeToFile() {

	ufile, _ := os.OpenFile("./data/User.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer ufile.Close()
	enc := json.NewEncoder(ufile)
	err := enc.Encode(instance.userList)
	if err != nil {
		log.Println("Error in encoding userList")
	}

	mfile, _ := os.OpenFile("./data/Meeting.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer mfile.Close()
	enc = json.NewEncoder(mfile)
	err = enc.Encode(instance.meetingList)
	if err != nil {
		log.Println("Error in encoding meetingList")
	}

	cufile, _ := os.OpenFile("./data/curUser.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer cufile.Close()
	enc = json.NewEncoder(cufile)
	err = enc.Encode(instance.curUser)
	if err != nil {
		log.Println("Error in encoding currUser")
	}
}

// CreateUser *
func (s *Storage) CreateUser(newUser User) {
	s.userList = append(s.userList, newUser)
	writeToFile()
}

// QueryUser *
func (s *Storage) QueryUser(filter uFilter) []User {
	userQuery := []User{}
	for _, user := range s.userList {
		if filter(&user) {
			userQuery = append(userQuery, user)
		}
	}
	return userQuery
}

// UpdateUser *
func (s *Storage) UpdateUser(filter uFilter, switcher uSwitcher) int {
	count := 0
	for i := 0; i < len(s.userList); i++ {
		if filter(&(s.userList[i])) {
			switcher(&(s.userList[i]))
			count++
		}
	}
	writeToFile()
	return count
}

//DeleteUser *
func (s *Storage) DeleteUser(filter uFilter) int {
	count := 0
	for i := 0; i < len(s.userList); i++ {
		user := s.userList[i]
		if filter(&user) {
			s.userList = append(s.userList[:i], s.userList[i+1:]...)
			i--
			count++
		}
	}
	writeToFile()
	return count
}

// CreateMeeting *
func (s *Storage) CreateMeeting(newMeeting Meeting) {
	s.meetingList = append(s.meetingList, newMeeting)
	writeToFile()
}

// QueryMeeting *
func (s *Storage) QueryMeeting(filter mFilter) []Meeting {
	mQuery := []Meeting{}
	for _, meeting := range s.meetingList {
		if filter(&meeting) {
			mQuery = append(mQuery, meeting)
		}
	}
	return mQuery
}

// UpdateMeeting *
func (s *Storage) UpdateMeeting(filter mFilter, switcher mSwitcher) int {
	count := 0
	for i := 0; i < len(s.meetingList); i++ {
		if filter(&(s.meetingList[i])) {
			switcher(&(s.meetingList[i]))
			count++
		}
	}
	writeToFile()
	return count
}

// DeleteMeeting *
func (s *Storage) DeleteMeeting(filter mFilter) int {
	count := 0
	for i := 0; i < len(s.meetingList); i++ {
		meeting := s.meetingList[i]
		if filter(&meeting) {
			s.meetingList = append(s.meetingList[:i], s.meetingList[i+1:]...)
			i--
			count++
		}
	}
	writeToFile()
	return count
}

// GetCurUser *
func (s Storage) GetCurUser() User {
	return s.curUser
}

// SetCurUser *
func (s *Storage) SetCurUser(u User) {
	s.curUser.Assign(u)
	writeToFile()
}

func init() {
	GetStorage()
}
