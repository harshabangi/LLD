package service

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/event-calendar/internal/db"
	"github.com/harshabangi/LLD/apps/event-calendar/internal/models"
)

type Service struct {
	db *db.DB
}

func NewService() *Service {
	return &Service{
		db: db.NewDB(),
	}
}

func (s *Service) createUser(info *models.User) error {
	if err := info.Validate(); err != nil {
		return err
	}
	s.db.User.CreateUser(*info)
	return nil
}

func (s *Service) creatTeam(teamName string, userNames []string) error {
	users := make([]models.User, len(userNames))

	for i, v := range userNames {
		user, err := s.db.User.GetUser(v)
		if err != nil {
			return err
		}
		if user.TeamName != "" && user.TeamName != teamName {
			return fmt.Errorf("user %s is already associated with another team %s", v, user.TeamName)
		}
		user.TeamName = teamName
		users[i] = *user
	}

	s.db.Team.CreateTeam(teamName, users)
	return nil
}

func (s *Service) creatEvent(event *models.Event) error {
	// validate event request body
	if err := event.Validate(); err != nil {
		return err
	}

	// ensure teams
	for _, v := range event.Teams {
		if err := s.db.Team.EnsureTeam(v); err != nil {
			return err
		}
	}

	// ensure users
	for _, userName := range event.Users {
		userInfo, err := s.db.User.GetUser(userName)
		if err != nil {
			return err
		}
		if event.StartTime.Sub(userInfo.StartWorkingTime) < 0 || event.EndTime.Sub(userInfo.EndWorkingTime) > 0 {
			return fmt.Errorf("event %s is out of working hours for user %s", event.Name, userName)
		}

		for _, teamName := range event.Teams {
			if s.db.Team.IsUserInTeam(teamName, userName) {
				return fmt.Errorf("user %s is already part of team %s", userName, teamName)
			}
		}
	}

	return nil
}
