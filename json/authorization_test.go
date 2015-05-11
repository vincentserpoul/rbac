package json

import (
	"log"
	"reflect"
	"testing"
)

func getTestAppRbacConf() RbacConf {
	var AppRbacConf RbacConf
	var err error

	AppRbacConf.AppUserRole, err = LoadUserRoleFromFile("")
	if err != nil {
		log.Fatal(err)
	}

	AppRbacConf.AppRoleActions, err = LoadRoleActionsFromFile("")
	if err != nil {
		log.Fatal(err)
	}

	return AppRbacConf

}

func TestGetRoleFromUserID(t *testing.T) {
	cases := []struct {
		userID string
		want   string
	}{
		{"userID1", "user"},
		{"userID2", "admin"},
		{"randomuser", ""},
	}

	AppRbacConf := getTestAppRbacConf()

	for _, c := range cases {
		got, _ := AppRbacConf.getRoleFromUserID(c.userID)
		if got != c.want {
			t.Errorf("getRoleFromUserID(%q) == %q, want %q", c.userID, got, c.want)
		}
	}
}

func TestGetAvailableActionsFromRole(t *testing.T) {

	cases := []struct {
		role string
		want map[string]bool
	}{
		{
			"user",
			map[string]bool{"GET_/": true, "POST_/": false},
		},
		{
			"admin",
			map[string]bool{"GET_/": true, "POST_/": true},
		},
		{
			"randomrole",
			map[string]bool{},
		},
	}

	AppRbacConf := getTestAppRbacConf()

	for _, c := range cases {
		got, _ := AppRbacConf.getAvailableActionsFromRole(c.role)
		if reflect.DeepEqual(got, c.want) {
			t.Errorf("getAvailableActionsFromRole(%q) == %v, want %v", c.role, got, c.want)
		}
	}
}

func TestIsUserAuthorized(t *testing.T) {
	cases := []struct {
		userID, action string
		want           bool
	}{
		{"userID1", "GET_/", true},
		{"userID1", "POST_/", false},
		{"userID2", "GET_/", true},
		{"userID2", "POST_/", true},
		{"randomuser", "POST_/", false},
	}

	AppRbacConf := getTestAppRbacConf()

	for _, c := range cases {
		got, _ := AppRbacConf.IsUserAuthorized(c.userID, c.action)
		if got != c.want {
			t.Errorf("IsUserAuthorized(%q, %q) == %t, want %t", c.userID, c.action, got, c.want)
		}
	}
}
