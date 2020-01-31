package lib_test

import (
	"testing"

	"github.com/warrensbox/jscheck/lib"
)

func TestvalidJSON(t *testing.T) {

	str := `{"page": 1, "fruits": ["apple", "peach"]}` //no error

	valid := lib.IsJSON(str)
	if valid {
		t.Logf("Value boolean is %v [expected]", valid)
	} else {
		t.Errorf("Value boolean is %v [unexpected]", valid)
	}

	str = `{"page": 1 "fruits": ["apple", "peach"]}` //has error
	valid = lib.IsJSON(str)
	if !valid {
		t.Logf("Value boolean is %v [expected]", valid)
	} else {
		t.Errorf("Value boolean is %v [unexpected]", valid)
	}
}

func TestInvalidJSON(t *testing.T) {

	str := `{"page": 1 "fruits": ["apple", "peach"]}` //has error
	valid := lib.IsJSON(str)
	if !valid {
		t.Logf("Value boolean is %v [expected]", valid)
	} else {
		t.Errorf("Value boolean is %v [unexpected]", valid)
	}
}
