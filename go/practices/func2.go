package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	ID   int
	Name string
}

type updateStatus struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	u := user{
		ID:   1432,
		Name: "Betty",
	}

	if us, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(us)
		return
	}

	fmt.Println("Updated user record for ID", u.ID)
}

func updateUser(u *user) (*updateStatus, error) {
	// response simulates a JSON response.
	response := `{"Modified":1, "Duration":0.005, "Success": true, "Message": "updated"}`

	// Unmarshal the json document into a value of
	// the userStats struct type.
	var us updateStatus
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update was Successful.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
