package pkg

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB, firstname, lastname, phone string, departmentID int) error {
	student := Student{First_Name: firstname, Last_Name: lastname, Phone: phone, DepartmentID: departmentID}

	return db.Create(&student).Error
}

func CreateCourse(db *gorm.DB, name string, duration int, instructorID int) error {
	course := Course{Name: name, Duration: duration, InstructorID: instructorID}

	return db.Create(&course).Error
}

func CreateDepartment(db *gorm.DB, name string) error {
	department := Department{Name: name}

	return db.Create(&department).Error
}

func CreateInstructor(db *gorm.DB, firstname, lastname, phone string, departmentID int) error {
	instructor := Instructor{First_Name: firstname, Last_Name: lastname, Phone: phone, DepartmentID: departmentID}

	return db.Create(&instructor).Error
}

func CreateEnrollment(db *gorm.DB, grade int, courseID int, studentID int) error {
	enrollment := Enrollment{Grade: grade, CourseID: courseID, StudentID: studentID}

	return db.Create(&enrollment).Error
}

func ReadStudent(db *gorm.DB, studentID uint) (Student, error) {
	var student Student
	err := db.First(&student, studentID).Error
	return student, err
}

func ReadCourse(db *gorm.DB, courseID uint) (Course, error) {
	var course Course
	err := db.First(&course, courseID).Error
	return course, err
}

func ReadDepartment(db *gorm.DB, departmentID uint) (Department, error) {
	var depart Department
	err := db.First(&depart, departmentID).Error
	return depart, err
}

func ReadInstructor(db *gorm.DB, instructorID uint) (Instructor, error) {
	var instructor Instructor
	err := db.First(&instructor, instructorID).Error
	return instructor, err
}

func ReadEnrollment(db *gorm.DB, enrollmentID uint) (Enrollment, error) {
	var enrollment Enrollment
	err := db.First(&enrollment, enrollmentID).Error
	return enrollment, err
}

func UpdateInstructor(db *gorm.DB, instructorID uint, firstName, lastName, phone string, departmentId int) error {
	instructor, err := ReadInstructor(db, instructorID)

	if err != nil {
		return err
	}

	instructor.First_Name = firstName
	instructor.Last_Name = lastName
	instructor.Phone = phone
	instructor.DepartmentID = departmentId

	return db.Save(&instructor).Error
}

func UpdateStudent(db *gorm.DB, studentID uint, firstName, lastName, phone string, departmentID int) error {
	student, err := ReadStudent(db, studentID)

	if err != nil {
		return err
	}

	student.First_Name = firstName
	student.Last_Name = lastName
	student.Phone = phone
	student.DepartmentID = departmentID

	return db.Save(&student).Error
}

func UpdateDepartment(db *gorm.DB, departmentID uint, name string) error {
	department, err := ReadDepartment(db, departmentID)

	if err != nil {
		return err
	}

	department.Name = name

	return db.Save(&department).Error
}

func UpdateCourse(db *gorm.DB, courseID uint, name string, duration int, instructorID int) error {
	course, err := ReadCourse(db, courseID)

	if err != nil {
		return err
	}

	course.Name = name
	course.Duration = duration
	course.InstructorID = instructorID

	return db.Save(&course).Error
}

func UpdateEnrollment(db *gorm.DB, enrollmentID uint, grade int, courseID, studentID uint) error {
	enrollment, err := ReadEnrollment(db, enrollmentID)

	if err != nil {
		return err
	}

	enrollment.Grade = grade
	enrollment.CourseID = int(courseID)
	enrollment.StudentID = int(studentID)
	return db.Save(&enrollment).Error
}

func DeleteStudent(db *gorm.DB, studentID uint) error {
	student, err := ReadStudent(db, studentID)

	if err != nil {
		return err
	}
	return db.Delete(&student).Error
}

func DeleteCourse(db *gorm.DB, courseID uint) error {
	course, err := ReadCourse(db, courseID)

	if err != nil {
		return err
	}
	return db.Delete(&course).Error
}

func DeleteDepartment(db *gorm.DB, departmentID uint) error {
	department, err := ReadDepartment(db, departmentID)

	if err != nil {
		return err
	}

	return db.Delete(&department).Error
}

func DeleteInstructor(db *gorm.DB, instructorID uint) error {
	instructor, err := ReadInstructor(db, instructorID)

	if err != nil {
		return err
	}

	return db.Delete(&instructor).Error
}

func DeleteEnrollment(db *gorm.DB, enrollmentID uint) error {
	enrollment, err := ReadEnrollment(db, enrollmentID)

	if err != nil {
		return err
	}

	return db.Delete(&enrollment).Error
}

func EnrollStudentInCourse(db *gorm.DB, studentID, courseID uint, grade int) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	enrollment := Enrollment{
		Grade:     grade,
		CourseID:  int(courseID),
		StudentID: int(studentID),
	}

	if err := tx.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to enroll student in course: %v", err)
	}

	if err := tx.Model(&Student{}).Where("id = ?", studentID).Association("Enrollment").Append(&enrollment); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update student's enrollment: %v", err)
	}

	if err := tx.Model(&Course{}).Where("id = ?", courseID).Association("Enrollment").Append(&enrollment); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update course's enrollment: %v", err)
	}

	return tx.Commit().Error
}

func (student *Student) BeforeCreate(tx *gorm.DB) (err error) {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	return nil
}

func (instructor *Instructor) BeforeUpdate(tx *gorm.DB) (err error) {
	instructor.UpdatedAt = time.Now()
	return nil
}

func SoftDeleteStudent(db *gorm.DB, studentID uint) error {
	return db.Model(&Student{}).Where("id = ?", studentID).Delete(&Student{}).Error
}
