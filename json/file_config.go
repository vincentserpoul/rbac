// Package rbacjson let you use files as data storage for your rbac conf
package json

import (
	"encoding/json"
	"io/ioutil"
)

// LoadUserRoleFromFile loads config user role from json file, according to environment
func LoadUserRoleFromFile(fileName string) ([]UserRole, error) {
	var appUserRole []UserRole

	if fileName == "" {
		fileName = "default/userrole.json"
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []UserRole{}, err
	}
	if err := json.Unmarshal(file, &appUserRole); err != nil {
		return []UserRole{}, err
	}

	return appUserRole, nil

}

// LoadRoleActionsFromFile loads config role actions/rights from json file, according to environment
func LoadRoleActionsFromFile(fileName string) ([]RoleActions, error) {
	var appRoleActions []RoleActions

	if fileName == "" {
		fileName = "default/roleactions.json"
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []RoleActions{}, err
	}

	if err := json.Unmarshal(file, &appRoleActions); err != nil {
		return []RoleActions{}, err
	}

	return appRoleActions, nil
}
