package models

import (
	"fmt"
)

// TestData .
var TestData []Message

// GetTestData .
func GetTestData() []Message {
	TestData := make([]Message, 50)

	for i := 0; i < 50; i++ {
		TestData[i].Text = fmt.Sprintf("Text%d", i)
		TestData[i].UserID = "wxx"
		TestData[i].UserName = "奔跑的兔"
		TestData[i].Image = fmt.Sprintf("Image%d", i)
		TestData[i].MessageID = "0"
		commits := make([]CommitInfo, 2)
		commits[0].UserID = "0"
		commits[1].UserID = "1"
		commits[0].Commit = fmt.Sprintf("Commit%d-1", i)
		commits[1].Commit = fmt.Sprintf("Commit%d-2", i)
		TestData[i].Commit = commits
		TestData[i].Praise = i
		TestData[i].Time = "2019-11-05 11:29:18"

		if i%2 == 0 {
			TestData[i].HavePraise = false
		} else {
			TestData[i].HavePraise = true
		}
	}

	return TestData
}
