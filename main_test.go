package main

import (
	"fmt"
	"six_week/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const testDBName = "test_db"

func setupTestDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost user=postgres password=2004 dbname=%s port=5432 sslmode=disable", testDBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&pkg.Department{}, &pkg.Student{}, &pkg.Instructor{}, &pkg.Course{}, &pkg.Enrollment{})

	return db, nil
}
func TestCreateStudent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}

	departmentName := "MATH"
	err = pkg.CreateDepartment(db, departmentName)
	assert.NoError(t, err)

	var department pkg.Department
	err = db.Where("name = ?", departmentName).First(&department).Error
	assert.NoError(t, err)

	firstName := "John"
	lastName := "Doe"
	phone := "123456789"

	err = pkg.CreateStudent(db, firstName, lastName, phone, 1)
	assert.NoError(t, err)

	err = pkg.CreateStudent(db, "Alice", "Smith", "987654321", 1)
	assert.NoError(t, err)

	err = pkg.CreateStudent(db, "Bob", "Johnson", "555555555", 999)
	assert.Error(t, err)

}
