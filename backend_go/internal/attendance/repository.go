package attendance

type Attendance struct {
	ID        int
	StudentID int
	Date      string
	Status    string
	CreatedAt string
	UpdatedAt string
}

type AttendancePostgresRepository interface {
	AddAttendance(studentID int, status string) error
	GetAttendanceForStudent(studentID int) ([]Attendance, error)
	UpdateAttendance(studentID int, status string) error
	DeleteAttendance(studentID int, attendanceDate string) error
}
