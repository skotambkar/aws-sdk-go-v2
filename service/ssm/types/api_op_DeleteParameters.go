// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteParametersInput struct {
	_ struct{} `type:"structure"`

	// The names of the parameters to delete.
	//
	// Names is a required field
	Names []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteParametersInput"}

	if s.Names == nil {
		invalidParams.Add(aws.NewErrParamRequired("Names"))
	}
	if s.Names != nil && len(s.Names) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Names", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteParametersOutput struct {
	_ struct{} `type:"structure"`

	// The names of the deleted parameters.
	DeletedParameters []string `min:"1" type:"list"`

	// The names of parameters that weren't deleted because the parameters are not
	// valid.
	InvalidParameters []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DeleteParametersOutput) String() string {
	return awsutil.Prettify(s)
}