package main

import "testing"

func TestGetHeading(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "Test get heading",
			inputHTML: "<html><body><h2>Please Work</h2><main><p>functional</p></main></html></body>",
			expected:  "Please Work",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected header: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraph(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "Test get first paragraph",
			inputHTML: "<html><body><h2>Please Work</h2><main><p>functional</p></main></html></body>",
			expected:  "functional",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected paragraph: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
