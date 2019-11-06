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
		TestData[i].Image = fmt.Sprintf("Image%d", i)
		commits := make([]CommitInfo, 2)
		commits[0].UserID = "0"
		commits[1].UserID = "1"
		commits[0].Commit = fmt.Sprintf("Commit%d-1", i)
		commits[1].Commit = fmt.Sprintf("Commit%d-2", i)
		TestData[i].Commit = commits
		TestData[i].Praise = i
	}

	return TestData
}
