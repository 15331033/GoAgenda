package test

import (
	"testing"
	"Agenda/123/cli/entity"
)

func TestIsParticipator(t *testing.T) {
	cases := []struct{
		in string
		want bool
	}{
		{"admin",true},
		{"",false},
	}
	participator := []string{"admin"}
	var testMeeting entity.Meeting
	testMeeting.SetParticipator(participator)
	for _, c := range cases {
		got := testMeeting.IsParticipator(c.in)
		if got != c.want {
			t.Errorf("IsParticipator(%v) == %v, want %v",c.in, got, c.want)
		}
	}
}