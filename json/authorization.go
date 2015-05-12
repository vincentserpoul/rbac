// Package json to be able to read rbac config from files
package json

import (
	"errors"
	"strings"
)

// UserRole contains association between user and role, one to one
type UserRole struct {
	UserID string
	Role   string
}

// RoleActions contains association between role and Actions, one to many
type RoleActions struct {
	Role    string
	Actions []Action
}

// Action contains a string, label, representing the wanted action
type Action struct {
	Label string
}

// RbacConf complete struct to work with rbac
type RbacConf struct {
	AppUserRole    []UserRole
	AppRoleActions []RoleActions
}

// IsUserAuthorized Tells if the user is authorized for this action or not
func (AppRbacConf RbacConf) IsUserAuthorized(userID string, action string) (bool, error) {
	// by default all actions are not authorized
	isAuthorized := false

	var err error

	role, err := AppRbacConf.getRoleFromUserID(userID)

	if err != nil || role == "" {
		return false, err
	}

	availableActions, err := AppRbacConf.getAvailableActionsFromRole(role)

	if err != nil {
		return false, err
	}

	isAuthorized = IsActionWithinAvailableActions(action, availableActions)

	if !isAuthorized {
		err = errors.New(action + " not allowed for the user " + userID + " with his current role, " + role)
	}

	return isAuthorized, err
}

// IsActionWithinAvailableActions tells you if an action is within a set of actions
func IsActionWithinAvailableActions(action string, availableActions []Action) bool {
	isWithin := false
	for _, availableAction := range availableActions {
		// In case the available action finishes with a *, allowing all subactions
		if strings.LastIndex(availableAction.Label, "*") == len(availableAction.Label)-1 {
			actionWOStar := strings.TrimSuffix(availableAction.Label, "*")
			if strings.Index(action, actionWOStar) == 0 {
				isWithin = true
				break
			}
		} else if action == availableAction.Label {
			isWithin = true
			break
		}
	}
	return isWithin
}

// getRoleFromUserID gets a role according to userId, if more than one, only one will be returned
func (AppRbacConf *RbacConf) getRoleFromUserID(userID string) (string, error) {

	var role string
	var err error

	for _, userRole := range AppRbacConf.AppUserRole {
		if userRole.UserID == userID {
			role = userRole.Role
			break

		}
	}

	if role == "" {
		err = errors.New("No existing role configured for " + userID)
	}

	return role, err
}

// getAvailableActionsFromRole gets a list of available actions according to the role, if defined more than once, only the first one will be returned
func (AppRbacConf *RbacConf) getAvailableActionsFromRole(role string) ([]Action, error) {
	availableActions := []Action{}

	var err error

	for _, roleActions := range AppRbacConf.AppRoleActions {
		if roleActions.Role == role {
			availableActions = roleActions.Actions
		}
	}

	if len(availableActions) == 0 {
		err = errors.New("No action configured for this role: " + role)
	}

	return availableActions, err

}
