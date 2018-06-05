// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// cars HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/cars/design -o
// $(GOPATH)/src/goa.design/goa/examples/cars

package client

import (
	carssvc "goa.design/goa/examples/cars/gen/cars"
)

// BuildLoginPayload builds the payload for the cars login endpoint from CLI
// flags.
func BuildLoginPayload(carsLoginUser string, carsLoginPassword string) (*carssvc.LoginPayload, error) {
	var user string
	{
		user = carsLoginUser
	}
	var password string
	{
		password = carsLoginPassword
	}
	payload := &carssvc.LoginPayload{
		User:     user,
		Password: password,
	}
	return payload, nil
}

// BuildListPayload builds the payload for the cars list endpoint from CLI
// flags.
func BuildListPayload(carsListStyle string, carsListToken string) (*carssvc.ListPayload, error) {
	var style string
	{
		style = carsListStyle
	}
	var token string
	{
		token = carsListToken
	}
	payload := &carssvc.ListPayload{
		Style: style,
		Token: token,
	}
	return payload, nil
}
