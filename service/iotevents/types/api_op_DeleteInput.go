// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteInputInput struct {
	_ struct{} `type:"structure"`

	// The name of the input to delete.
	//
	// InputName is a required field
	InputName *string `location:"uri" locationName:"inputName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInputInput"}

	if s.InputName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputName"))
	}
	if s.InputName != nil && len(*s.InputName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteInputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteInputOutput) String() string {
	return awsutil.Prettify(s)
}
