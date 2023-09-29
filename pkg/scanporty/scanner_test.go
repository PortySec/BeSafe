package scanporty

import "testing"

func TestScanner(t *testing.T) {
	testCases := []struct {
		Host     string
		Port     string
		Expected bool
	}{
		{"127.0.0.1", "80", false},
		{"google.com", "80", true},
	}

	for _, testCase := range testCases {
		testRes := Scanner(testCase.Host, testCase.Port)
		if testRes != testCase.Expected {
			t.Errorf("Test failed for %v host and %v port", testCase.Host, testCase.Port)
		}
	}
}
