package main

import (
	"fmt"

	"six_week/pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2004"
	dbname   = "univer"
)

func GetStudents(db *gorm.DB, departmentID int) []pkg.Student {
	var students []pkg.Student

	db.Where("department_id = ?", departmentID).Find(&students)

	return students
}

func GetCoursesByInstructor(db *gorm.DB, instructorID int) []pkg.Course {
	var courses []pkg.Course

	db.Where("instructor_id = ?", instructorID).Find(&courses)

	return courses
}

func GetEnrollmentsForStudent(db *gorm.DB, studentID int) []pkg.Enrollment {
	var enrollments []pkg.Enrollment

	db.Where("student_id = ?", studentID).Find(&enrollments)

	return enrollments
}

func GetTotalStudentsInDepartment(db *gorm.DB, departmentID uint) (int, error) {
	var count int64
	err := db.Model(&pkg.Student{}).Where("department_id = ?", departmentID).Count(&count).Error
	return int(count), err
}

type CourseWithEnrollmentCount struct {
	pkg.Course
	EnrolledStudentsCount int
}

func GetCoursesWithEnrollmentCount(db *gorm.DB) ([]CourseWithEnrollmentCount, error) {
	var coursesWithCount []CourseWithEnrollmentCount
	err := db.Model(&pkg.Course{}).
		Select("courses.*, COUNT(enrollments.id) as enrolled_students_count").
		Joins("LEFT JOIN enrollments ON courses.id = enrollments.course_id").
		Group("courses.id").
		Scan(&coursesWithCount).Error
	return coursesWithCount, err
}

func main() {
	dbs := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dbs), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&pkg.Department{}, &pkg.Student{}, &pkg.Instructor{}, &pkg.Course{}, &pkg.Enrollment{})

	// Soft delete a student
	err = pkg.SoftDeleteStudent(db, 2)
	if err != nil {
		fmt.Println("Failed to soft delete student:", err)
	}

	students := GetStudents(db, 1)

	for _, student := range students {
		fmt.Printf("Student ID: %d, Name: %s %s, Phone: %s, DeletedAt: %v\n",
			student.ID, student.First_Name, student.Last_Name, student.Phone, student.DeletedAt)
	}

	coursesWithCount, err := GetCoursesWithEnrollmentCount(db)
	if err != nil {
		fmt.Println("Error retrieving courses with enrollment count:", err)
		return
	}

	fmt.Println("List of Courses with Enrolled Students:")
	for _, course := range coursesWithCount {
		fmt.Printf("Course ID: %d, Name: %s, Duration: %d, Enrolled Students: %d\n",
			course.ID, course.Name, course.Duration, course.EnrolledStudentsCount)
	}

	x := []int{1, 2, 3}

	for num := range x {
		fmt.Print(num)
	}

}

func funcna() (int, bool) {
	return 0, true
}
