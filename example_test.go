package upid

import "fmt"

func ExampleEncode() {
	fmt.Println(Encode("user-id", "profile-id"))

	// Output:
	// user-id_profile-id
}

func ExampleDecode() {
	userID, profileID := Decode("user-id_profile-id")
	fmt.Printf("user ID: %s\nprofile ID: %s\n", userID, profileID)

	// Output:
	// user ID: user-id
	// profile ID: profile-id
}
