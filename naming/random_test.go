package naming

import (
	"testing"
)

func TestGiveMeAName(t *testing.T) {
	for i := 0; i < 10; i++ {
		name := GiveMeAName()
		found := false
		for _, v := range names {
			if v == name {
				found = true
			}
		}
		if !found {
			t.FailNow()
		}
	}
}
