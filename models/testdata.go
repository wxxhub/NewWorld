package models

import (
	"fmt"
)

// TestData .
var TestData []Message

// GetTestData .
func GetTestData() []Message {
	TestData := make([]Message, 100, 100)

	for i := 0; i < 100; i++ {
		TestData[i].Name = fmt.Sprintf("Uset%d", i)
		TestData[i].Text = fmt.Sprintf("Text%d", i)
		TestData[i].Image = fmt.Sprintf("Image%d", i)
		commits := make([]string, 2)
		commits[0] = fmt.Sprintf("Commit%d-1", i)
		commits[1] = fmt.Sprintf("Commit%d-2", i)
		TestData[i].Commit = commits
		TestData[i].Praise = i
	}

	return TestData
}
