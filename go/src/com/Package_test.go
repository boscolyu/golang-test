package bosco

import (
	"testing"
)

func TestBoscoFunction(test *testing.T) {
	var ret = BoscoFunction()
	if ret == false {
		test.Fail()
	}
}
