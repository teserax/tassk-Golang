package main

import "testing"

func TestFilter(t *testing.T) {

	tests := []Options{
		{MinAge: 1, MaxAge: 15, Group: "333", Diagnos: "test"},
		{2, 25, "qq", "est"},
		{3, 45, "ee", "top"},
		{4, 65, "00", "test"},
	}
}
