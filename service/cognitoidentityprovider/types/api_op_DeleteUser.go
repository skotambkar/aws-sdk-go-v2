// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to delete a user.
type DeleteUserInput struct {
	_ struct{} `type:"structure"`

	// The access token from a request to delete a user.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DeleteUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserOutput) String() string {
	return awsutil.Prettify(s)
}