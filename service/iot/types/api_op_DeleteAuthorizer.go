// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// The name of the authorizer to delete.
	//
	// AuthorizerName is a required field
	AuthorizerName *string `location:"uri" locationName:"authorizerName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAuthorizerInput"}

	if s.AuthorizerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerName"))
	}
	if s.AuthorizerName != nil && len(*s.AuthorizerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAuthorizerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}