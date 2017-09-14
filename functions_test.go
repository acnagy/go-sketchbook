package main

import "testing"

func TestCheckCI(t *testing.T) {

	if checkCI("Travis CI") != true {
		t.Fail()
	}
}
