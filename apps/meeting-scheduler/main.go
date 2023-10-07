package main

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

type User struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber string
}

type Users []User

func (u *Users) GetEmails() []string {
	emails := make([]string, len(*u))
	for i, v := range *u {
		emails[i] = v.Email
	}
	return emails
}

type MeetingRoom struct {
	ID       string
	Capacity int
}

type Meeting struct {
	ID            string
	MeetingRoomID string
	Agenda        string
	StartTime     *time.Time
	EndTime       *time.Time
	Users         Users
	CreatedBy     string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type MeetingManager struct {
	Users          Users
	MeetingRooms   []MeetingRoom
	MeetingHistory []Meeting
}

func NewMeetingManager() *MeetingManager {
	return &MeetingManager{}
}

func (m *MeetingManager) AddMeetingRoom(mr MeetingRoom) {
	m.MeetingRooms = append(m.MeetingRooms, mr)
}

func (m *MeetingManager) AddUser(u User) {
	m.Users = append(m.Users, u)
}

func (m *MeetingManager) GetUsersByEmails(emails []string) []User {
	userNamesMap := toMap(emails)
	var res []User

	for _, u := range m.Users {
		if _, ok := userNamesMap[u.Email]; ok {
			res = append(res, u)
		}
	}
	return res
}

func (m *MeetingManager) CreateMeeting(mr MeetingReq) {
	if err := mr.Validate(); err != nil {
		log.Printf("invalid request for creating meeting: %s", err.Error())
		return
	}

	id := uuid.New().String()
	now := time.Now()

	meetingRoomID := m.GetAvailableMeetingRoom(mr.StartTime, mr.EndTime, len(mr.UserEmails))
	if meetingRoomID == "" {
		log.Printf("no meeting room available")
		return
	}

	meeting := Meeting{
		ID:            id,
		MeetingRoomID: meetingRoomID,
		Agenda:        mr.Agenda,
		StartTime:     mr.StartTime,
		EndTime:       mr.EndTime,
		Users:         m.GetUsersByEmails(mr.UserEmails),
		CreatedBy:     mr.CreatedBy,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	m.MeetingHistory = append(m.MeetingHistory, meeting)
	meeting.SendEmailInviteToUsers()
}

func (m *MeetingManager) GetAvailableMeetingRoom(startTime, endTime *time.Time, requiredCapacity int) string {
	var (
		meetingRoomIDs            = m.GetMeetingRoomsAvailableForCapacity(requiredCapacity)
		meetingRoomIDsMap         = toMap(meetingRoomIDs)
		overlappingMeetingRoomIDs []string
		result                    string
	)

	for _, mh := range m.MeetingHistory {
		if _, ok := meetingRoomIDsMap[mh.MeetingRoomID]; ok && timeOverlap(mh.StartTime, mh.EndTime, startTime, endTime) {
			overlappingMeetingRoomIDs = append(overlappingMeetingRoomIDs, mh.MeetingRoomID)
		}
	}

	for _, o := range overlappingMeetingRoomIDs {
		delete(meetingRoomIDsMap, o)
	}

	for k := range meetingRoomIDsMap {
		result = k
		break
	}
	return result
}

func timeOverlap(st, et, rst, ret *time.Time) bool {
	return (rst.After(*st) && rst.Before(*et)) ||
		(ret.After(*st) && ret.Before(*et))
}

func toMap(s []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

func (m *MeetingManager) GetMeetingRoomsAvailableForCapacity(requiredCapacity int) []string {
	var meetingRoomIDs []string

	for _, mr := range m.MeetingRooms {
		if mr.Capacity >= requiredCapacity {
			meetingRoomIDs = append(meetingRoomIDs, mr.ID)
		}
	}
	return meetingRoomIDs
}

func (m *Meeting) SendEmailInviteToUsers() {
	log.Printf("Sending email to users: [%s]", strings.Join(m.Users.GetEmails(), ", "))
}

type MeetingReq struct {
	Agenda     string
	CreatedBy  string
	StartTime  *time.Time
	EndTime    *time.Time
	UserEmails []string
}

func (mr *MeetingReq) Validate() error {
	if len(mr.UserEmails) == 0 {
		return errors.New("required user emails")
	}
	if mr.StartTime == nil || mr.EndTime == nil {
		return errors.New("required start time and end time")
	}
	return nil
}

func main() {
	mm := NewMeetingManager()

	// add users
	mm.AddUser(User{ID: "1001", Name: "x", Email: "x@gmail.com"})
	mm.AddUser(User{ID: "1002", Name: "y", Email: "y@gmail.com"})
	mm.AddUser(User{ID: "1003", Name: "z", Email: "z@gmail.com"})
	mm.AddUser(User{ID: "1004", Name: "a", Email: "a@gmail.com"})

	// add meeting rooms
	mm.AddMeetingRoom(MeetingRoom{ID: "221", Capacity: 2})
	mm.AddMeetingRoom(MeetingRoom{ID: "222", Capacity: 1})
	mm.AddMeetingRoom(MeetingRoom{ID: "223", Capacity: 1})
	mm.AddMeetingRoom(MeetingRoom{ID: "224", Capacity: 1})
	mm.AddMeetingRoom(MeetingRoom{ID: "225", Capacity: 1})

	now := time.Now()
	t1 := now
	t2 := now.Add(5 * time.Second)
	t3 := now.Add(2 * time.Second)
	t4 := now.Add(4 * time.Second)

	mm.CreateMeeting(MeetingReq{
		Agenda:     "discussion - 1",
		CreatedBy:  "x@gmail.com",
		StartTime:  &t1,
		EndTime:    &t2,
		UserEmails: []string{"x@gmail.com", "y@gmail.com"},
	})

	mm.CreateMeeting(MeetingReq{
		Agenda:     "discussion - 2",
		CreatedBy:  "z@gmail.com",
		StartTime:  &t3,
		EndTime:    &t4,
		UserEmails: []string{"z@gmail.com", "a@gmail.com"},
	})
}
