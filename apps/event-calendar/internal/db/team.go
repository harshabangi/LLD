package db

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/event-calendar/internal/models"
)

type teamStorage struct {
	teams map[string]map[string]models.User
}

type TeamDB interface {
	CreateTeam(teamName string, users []models.User)
	GetTeam(teamName string) (map[string]models.User, error)
	EnsureTeam(teamName string) error
	GetUsers(teamName string) []models.User
	IsUserInTeam(teamName, userName string) bool
}

func NewTeamDB() TeamDB {
	return &teamStorage{
		teams: make(map[string]map[string]models.User),
	}
}

func (t *teamStorage) CreateTeam(teamName string, users []models.User) {
	userMap := make(map[string]models.User)

	for _, v := range users {
		userMap[v.Name] = v
	}
	t.teams[teamName] = userMap
}

func (t *teamStorage) GetTeam(teamName string) (map[string]models.User, error) {
	teamUsers, ok := t.teams[teamName]
	if !ok {
		return nil, fmt.Errorf("unknown team: %s", teamName)
	}

	return teamUsers, nil
}

func (t *teamStorage) EnsureTeam(teamName string) error {
	_, ok := t.teams[teamName]
	if !ok {
		return fmt.Errorf("unknown team: %s", teamName)
	}

	return nil
}

func (t *teamStorage) IsUserInTeam(teamName, userName string) bool {
	// assuming teamName always exists
	_, ok := t.teams[teamName][userName]
	return ok
}

func (t *teamStorage) GetUsers(teamName string) []models.User {
	var users []models.User

	for _, v := range t.teams[teamName] {
		users = append(users, v)
	}
	return users
}
