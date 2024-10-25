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
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RoleID   int    `json:"role_id"`
	GroupID  int    `json:"group_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Role struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	GroupId int    `json:"group_id"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

type Schedule struct {
	ID         int    `json:"id"`
	GroupID    int    `json:"group_id"`
	DayOfWeek  int    `json:"day_of_week"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Subject    string `json:"subject"`
	TeacherID  int    `json:"teacher_id"`
	Location   string `json:"location"`
	Recurrence string `json:"recurrence"`
}
