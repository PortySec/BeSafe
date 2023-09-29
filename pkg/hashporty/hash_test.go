package hashporty

import "testing"

func TestHasher(t *testing.T) {

	testCases := []struct {
		Plain     string
		Hash      string
		Algorithm string
	}{
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", "sha256"},
		{"hello", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043", "sha512"},
		{"hello", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", "sha1"},
	}
	for _, test := range testCases {
		testResult, err := HashPlainText(test.Algorithm, test.Plain)
		if err != nil {
			t.Error("hash function error", err)
		}
		if testResult != test.Hash {
			t.Errorf("Test failed for %v plain value and %v algorithm with follwoing result %v", test.Plain, test.Algorithm, testResult)
		}

	}
}
