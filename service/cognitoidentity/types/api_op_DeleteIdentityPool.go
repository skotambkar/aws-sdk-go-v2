// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the DeleteIdentityPool action.
type DeleteIdentityPoolInput struct {
	_ struct{} `type:"structure"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIdentityPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIdentityPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIdentityPoolInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteIdentityPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIdentityPoolOutput) String() string {
	return awsutil.Prettify(s)
}
