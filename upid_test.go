package upid

import "testing"

func TestEncode(t *testing.T) {
	for n, tc := range []struct {
		userID    string
		profileID string
		wantUPID  string
	}{
		{
			userID:    "user-id",
			profileID: "",
			wantUPID:  "user-id",
		},
		{
			userID:    "user-id",
			profileID: "profile-id",
			wantUPID:  "user-id_profile-id",
		},
		{
			userID:    "user-id_with-separator",
			profileID: "profile-id",
			wantUPID:  "user-id_with-separator_profile-id",
		},
		{
			userID:    "user-id_with_many_separators",
			profileID: "profile-id",
			wantUPID:  "user-id_with_many_separators_profile-id",
		},
		{
			userID:    "",
			profileID: "profile-id",
			wantUPID:  "_profile-id",
		},
		{
			userID:    "",
			profileID: "",
			wantUPID:  "",
		},
	} {
		gotUPID := Encode(tc.userID, tc.profileID)

		if gotUPID != tc.wantUPID {
			t.Errorf(
				"[%d] Encode(%q, %q) = %q, want %q",
				n, tc.userID, tc.profileID, gotUPID, tc.wantUPID,
			)
		}
	}
}

func TestDecode(t *testing.T) {
	for n, tc := range []struct {
		upid          string
		wantUserID    string
		wantProfileID string
	}{
		{
			upid:          "user-id",
			wantUserID:    "user-id",
			wantProfileID: "",
		},
		{
			upid:          "user-id_profile-id",
			wantUserID:    "user-id",
			wantProfileID: "profile-id",
		},
		{
			upid:          "user-id_with-separator_profile-id",
			wantUserID:    "user-id_with-separator",
			wantProfileID: "profile-id",
		},
		{
			upid:          "user-id_with_many_separators_profile-id",
			wantUserID:    "user-id_with_many_separators",
			wantProfileID: "profile-id",
		},
		{
			upid:          "_profile-id",
			wantUserID:    "",
			wantProfileID: "profile-id",
		},
		{
			upid:          "",
			wantUserID:    "",
			wantProfileID: "",
		},
	} {
		gotUserID, gotProfileID := Decode(tc.upid)

		if gotUserID != tc.wantUserID || gotProfileID != tc.wantProfileID {
			t.Errorf(
				"[%d] Decode(%q) = %q, %q; want %q, %q",
				n, tc.upid, gotUserID, gotProfileID, tc.wantUserID, tc.wantProfileID,
			)
		}
	}
}
