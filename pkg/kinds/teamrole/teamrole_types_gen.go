// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     GoTypesJenny
//     LatestJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package teamrole

import (
	"time"
)

// TeamRole defines model for TeamRole.
type TeamRole struct {
	// Created indicates when the team role was created.
	Created time.Time `json:"created"`

	// Namespace aka tenant/org id.
	Namespace string `json:"namespace"`

	// Unique role uid.
	RoleUid string `json:"roleUid"`

	// Unique team name
	TeamName string `json:"teamName"`
}
