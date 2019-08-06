/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test replaceSpace
 */

package main

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
