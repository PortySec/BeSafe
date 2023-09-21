package passgate

import "testing"

func TestPassgate(t *testing.T) {

	tests := []struct {
		Password string
		ShouldBe bool
	}{
		{"123456789", false},
		{"ALOOOOOaaaa888@@@", true},
		{"8888@@@@AAA", false},
		{"8888!!!!@@@@LLLLAaaaaaaa", true},
	}

	for _, test := range tests {
		res := IsValid(test.Password)
		if res != test.ShouldBe {
			t.Errorf("Test Faild for %q, it has to be %v but its ended in %v", test.Password, test.ShouldBe, res)
		}
	}

}
