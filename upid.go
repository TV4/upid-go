// Package upid encodes and decodes UPIDs. A UPID is a canonical string that
// combines a user ID and a profile ID.
//
//     <upid> ::= <user-id> | <user-id> "_" <profile-id>
package upid

import (
	"fmt"
	"strings"
)

// Encode encodes a user ID and profile ID into a UPID.
func Encode(userID, profileID string) string {
	if profileID == "" {
		return userID
	}
	return fmt.Sprintf("%s_%s", userID, profileID)
}

// Decode decides a UPID into a user ID and a profile ID.
func Decode(upid string) (userID, profileID string) {
	idx := strings.LastIndex(upid, "_")
	if idx == -1 {
		return upid, ""
	}
	return upid[:idx], upid[idx+1:]
}
