package model

const (
	MONDAY    = "Monday"
	TUESDAY   = "Tuesday"
	WEDNESDAY = "Wednesday"
	THURSDAY  = "Thursday"
	FRIDAY    = "Friday"
	SATURDAY  = "Saturday"
	SUNDAY    = "Sunday"
)

type UserAuth struct {
	Login    string
	Password string
}

type User struct {
	ID       int
	Name     string
	RoleID   int
	GroupID  int
	Login    string
	Password string
}

type Role struct {
	ID    int
	Value string
}

type Note struct {
	Id      int
	Title   string
	Body    string
	GroupId int
}

type Group struct {
	Id   int
	Name string
	Body string
}

type Schedule struct {
	Id        int
	GroupId   int
	DayOfWeek string
	Subject   string
	Teacher   string
	allowNull string
}

type Student struct {
	Id      int
	Name    string
	GroupId int
	Role    string
}
