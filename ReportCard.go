package main

import "fmt"

// Student struct represents student information
type Student struct {
	Name       string
	RollNumber int
	Paper1     int
	Paper2     int
	Paper3     int
}

// ReportCard struct represents the report card
type ReportCard struct {
	Student
	TotalMarks         int
	TotalMarksPerPaper int
	Percentage         float64
}

func main() {
	student := Student{
		Name:       "Mansi Mishra",
		RollNumber: 7,
		Paper1:     99,
		Paper2:     98,
		Paper3:     97,
	}

	totalMarksPerPaper := 100
	totalMarks := student.Paper1 + student.Paper2 + student.Paper3
	percentage := float64(totalMarks) / 3.0

	reportCard := ReportCard{
		Student:            student,
		TotalMarks:         totalMarks,
		TotalMarksPerPaper: totalMarksPerPaper,
		Percentage:         percentage,
	}

	printReportCard(reportCard)
}

func printReportCard(reportCard ReportCard) {
	fmt.Println("\n----------------- Report Card -----------------\n")
	fmt.Printf("Name: %s \tRoll Number: %d\n\n", reportCard.Name, reportCard.RollNumber)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Subject \t\tObtained \tTotal\n")
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Paper 1  \t\t   %d \t\t %d\n", reportCard.Paper1, reportCard.TotalMarksPerPaper)
	fmt.Printf("Paper 2  \t\t   %d \t\t %d\n", reportCard.Paper2, reportCard.TotalMarksPerPaper)
	fmt.Printf("Paper 3  \t\t   %d \t\t %d\n", reportCard.Paper3, reportCard.TotalMarksPerPaper)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Total   \t\t   %d \t\t %d\n", reportCard.TotalMarks, 3*reportCard.TotalMarksPerPaper)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("\tPercentage \t\t %.2f%% \n", reportCard.Percentage)
	fmt.Println("-----------------------------------------------")
	if reportCard.Paper1 > 40 && reportCard.Paper2 > 40 && reportCard.Paper3 > 40 {
		fmt.Println("-----------------  Passed !!  -----------------")
	} else {
		fmt.Println("-----------------  Failed !!  -----------------")
	}
}
