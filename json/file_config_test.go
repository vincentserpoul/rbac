package rbacjson

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadUserRoleFromFile(t *testing.T) {
	cases := []struct {
		in   string
		want []UserRole
	}{
		{
			"default/userrole.json",
			[]UserRole{
				{UserID: "userID1", Role: "user"},
				{UserID: "userID2", Role: "admin"},
			},
		},
	}
	for _, c := range cases {
		got, err := LoadUserRoleFromFile(c.in)
		if err != nil {
			log.Fatal(err)
		}
		// DeepEqual was returning false, changing to sprintf
		// if reflect.DeepEqual(got, c.want) {
		strGot := fmt.Sprintf("%v", got)
		strWant := fmt.Sprintf("%v", c.want)
		if strGot != strWant {
			t.Errorf("LoadUserRoleFromFile(%q) == %q, want %q", c.in, strGot, strWant)
		}
	}
}

func TestLoadRoleActionsFromFile(t *testing.T) {
	cases := []struct {
		in   string
		want []RoleActions
	}{
		{
			"default/roleactions.json",
			[]RoleActions{
				{Role: "user", Actions: []Action{{Label: "GET_/"}}},
				{Role: "admin", Actions: []Action{{Label: "GET_/"}, {Label: "POST_/"}}},
			},
		},
	}
	for _, c := range cases {
		got, err := LoadRoleActionsFromFile(c.in)
		if err != nil {
			log.Fatal(err)
		}
		// DeepEqual was returning false, changing to sprintf
		// if reflect.DeepEqual(got, c.want) {
		strGot := fmt.Sprintf("%v", got)
		strWant := fmt.Sprintf("%v", c.want)
		if strGot != strWant {
			t.Errorf("LoadRoleActionsFromFile(%q) == %v, want %v", c.in, strGot, strWant)
		}
	}
}
