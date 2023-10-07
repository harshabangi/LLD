package db

type DB struct {
	User  UserDB
	Team  TeamDB
	Event EventDB
}

func NewDB() *DB {
	return &DB{
		User:  NewUserDB(),
		Team:  NewTeamDB(),
		Event: NewEventDB(),
	}
}
