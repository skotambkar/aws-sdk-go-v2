// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLayerInput struct {
	_ struct{} `type:"structure"`

	// The layer ID.
	//
	// LayerId is a required field
	LayerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLayerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLayerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLayerInput"}

	if s.LayerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLayerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLayerOutput) String() string {
	return awsutil.Prettify(s)
}
