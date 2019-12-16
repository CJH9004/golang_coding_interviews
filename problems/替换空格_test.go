package problems

import "testing"

func TestReplaceSpace(t *testing.T) {
	tests := []struct {
		str string
		ret string
	}{
		{" as  asdfa   asdf asdf ", "%20as%20%20asdfa%20%20%20asdf%20asdf%20"},
		{"  ", "%20%20"},
		{"", ""},
	}

	for _, tt := range tests {
		ret := ReplaceSpace(tt.str)
		if ret != tt.ret {
			t.Errorf("Test ReplaceSpace(%v); got %v; expect %v",
				tt.str, ret, tt.ret)
			t.Fail()
		}
	}
}
