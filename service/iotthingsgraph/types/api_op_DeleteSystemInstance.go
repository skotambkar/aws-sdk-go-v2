// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteSystemInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the system instance to be deleted.
	Id *string `locationName:"id" type:"string"`
}

// String returns the string representation
func (s DeleteSystemInstanceInput) String() string {
	return awsutil.Prettify(s)
}

type DeleteSystemInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSystemInstanceOutput) String() string {
	return awsutil.Prettify(s)
}