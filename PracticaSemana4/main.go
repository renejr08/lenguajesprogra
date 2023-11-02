package main

import (
	"fmt"
	"time"
)

type Aula int

const (
	Aula101 Aula = iota
	Aula102
	LabComputacion
)

type Schedule struct {
	DayOfWeek time.Weekday
	StartTime int // Use an integer to represent hours (7 to 16 for 7am to 4pm)
	EndTime   int
}

type Student struct {
	Name  string
	Score float64
}

type Course struct {
	Name     string
	Schedule Schedule
	Room     Aula
	Students []Student
}

var courses []Course

// Function to add a course ensuring there is no schedule conflict
func AddCourse(newCourse Course) error {
	for _, course := range courses {
		if course.Schedule.DayOfWeek == newCourse.Schedule.DayOfWeek &&
			((newCourse.Schedule.StartTime >= course.Schedule.StartTime && newCourse.Schedule.StartTime < course.Schedule.EndTime) ||
				(newCourse.Schedule.EndTime > course.Schedule.StartTime && newCourse.Schedule.EndTime <= course.Schedule.EndTime)) &&
			course.Room == newCourse.Room {
			return fmt.Errorf("schedule conflict detected for course: %s", newCourse.Name)
		}
	}
	courses = append(courses, newCourse)
	return nil
}

// Function to add a student to a course
func AddStudentToCourse(courseName string, student Student) error {
	for i, course := range courses {
		if course.Name == courseName {
			courses[i].Students = append(courses[i].Students, student)
			return nil
		}
	}
	return fmt.Errorf("course not found: %s", courseName)
}

// Function to calculate statistics of a course
func CalculateCourseStatistics(courseName string) (int, int, float64, float64, float64) {
	var approved, failed int
	var sum, approvedSum, failedSum float64

	for _, course := range courses {
		if course.Name == courseName {
			for _, student := range course.Students {
				sum += student.Score
				if student.Score >= 60 {
					approved++
					approvedSum += student.Score
				} else {
					failed++
					failedSum += student.Score
				}
			}
			break
		}
	}

	average := sum / float64(len(courses))
	approvedAverage := approvedSum / float64(approved)
	failedAverage := failedSum / float64(failed)

	return approved, failed, average, approvedAverage, failedAverage
}

func main() {
	// Hardcoded course and student data as per instructions

	// Adding courses
	coursesData := []Course{
		{
			Name: "Programming 101",
			Schedule: Schedule{
				DayOfWeek: time.Monday,
				StartTime: 7,
				EndTime:   11,
			},
			Room: Aula101,
		},
		// Add more courses as needed
	}

	for _, course := range coursesData {
		err := AddCourse(course)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Adding students
	studentsData := []struct {
		CourseName string
		Student    Student
	}{
		{
			CourseName: "Programming 101",
			Student: Student{
				Name:  "John Doe",
				Score: 85.0,
			},
		},
		// Add more students as needed
	}

	for _, data := range studentsData {
		err := AddStudentToCourse(data.CourseName, data.Student)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Calculate statistics for a course
	approved, failed, average, approvedAverage, failedAverage := CalculateCourseStatistics("Programming 101")
	fmt.Printf("Approved: %d, Failed: %d, Average: %.2f, Approved Average: %.2f, Failed Average: %.2f\n", approved, failed, average, approvedAverage, failedAverage)
}
