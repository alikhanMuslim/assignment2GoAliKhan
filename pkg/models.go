package pkg

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	First_Name   string
	Last_Name    string
	Phone        string
	DepartmentID int
	Course       []Course `gorm:"many2many:student_courses;"`
	Enrollment   []Enrollment
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Course struct {
	gorm.Model
	Duration     int
	Name         string
	Enrollment   []Enrollment
	InstructorID int
}

type Department struct {
	gorm.Model
	Name       string
	Location   string
	Instructor []Instructor
	Student    []Student
}

type Instructor struct {
	gorm.Model
	First_Name   string
	Last_Name    string
	Phone        string
	DepartmentID int
	Courses      []Course
}

type Enrollment struct {
	gorm.Model
	Grade     int
	CourseID  int
	StudentID int
}
