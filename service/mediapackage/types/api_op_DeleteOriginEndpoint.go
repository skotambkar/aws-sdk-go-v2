// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteOriginEndpointInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOriginEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOriginEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOriginEndpointInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOriginEndpointOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOriginEndpointOutput) String() string {
	return awsutil.Prettify(s)
}