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
		{
			name:      "Test no heading",
			inputHTML: "<html><body><main><p>functional</p></main></html></body>",
			expected:  "",
		},
		{
			name:      "Test get heading 6",
			inputHTML: "<html><body><h6>Please Work</h6><main><p>functional</p></main></html></body>",
			expected:  "Please Work",
		},
		{
			name:      "Test get fake heading",
			inputHTML: "<html><body><h7>Please Work</h7><main><p>functional</p></main></html></body>",
			expected:  "",
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
		{
			name:      "Test no paragraph",
			inputHTML: "<html><body><h2>Please Work</h2><main></main></html></body>",
			expected:  "",
		},
		{
			name:      "Test get first paragraph in main",
			inputHTML: "<html><body><h2>Please Work</h2><p>not functional</p><main><p>functional</p></main></html></body>",
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

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody []byte
	}{}
}
