package leetcode

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	got := groupAnagrams([]string{"dog", "cat", "act", "big", "gib", "ibg"})
	want := [][]string{{"dog"}, {"cat", "act"}, {"big", "gib", "ibg"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %q \nwant %q", got, want)
	}
}
