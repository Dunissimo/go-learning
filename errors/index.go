package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CreateUserRequest is a request to create a new user.
type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

// the terms of the task state that it is possible to
// disregard the principle of single responsibility in this case
func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	decodedBody := CreateUserRequest{}

	err := json.Unmarshal(requestBody, &decodedBody)

	if err != nil {
		return CreateUserRequest{}, err
	}

	if len(decodedBody.Email) == 0 {
		return CreateUserRequest{}, errEmailRequired
	}

	if len(decodedBody.Password) == 0 {
		return CreateUserRequest{}, errPasswordRequired
	}

	if len(decodedBody.PasswordConfirmation) == 0 {
		return CreateUserRequest{}, errPasswordConfirmationRequired
	}

	if decodedBody.Password != decodedBody.PasswordConfirmation {
		return CreateUserRequest{}, errPasswordDoesNotMatch
	}

	return decodedBody, err
}

func main() {
	fmt.Println(DecodeAndValidateRequest([]byte("")))
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")))
	// fmt.Println(DecodeAndValidateRequest())
	// fmt.Println(DecodeAndValidateRequest())
	// fmt.Println(DecodeAndValidateRequest())
}
