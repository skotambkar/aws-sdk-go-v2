// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartAssociationsOnceInput struct {
	_ struct{} `type:"structure"`

	// The association IDs that you want to run immediately and only one time.
	//
	// AssociationIds is a required field
	AssociationIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s StartAssociationsOnceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartAssociationsOnceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartAssociationsOnceInput"}

	if s.AssociationIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationIds"))
	}
	if s.AssociationIds != nil && len(s.AssociationIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssociationIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartAssociationsOnceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartAssociationsOnceOutput) String() string {
	return awsutil.Prettify(s)
}