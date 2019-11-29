// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDisassociateUserStackInput struct {
	_ struct{} `type:"structure"`

	// The list of UserStackAssociation objects.
	//
	// UserStackAssociations is a required field
	UserStackAssociations []UserStackAssociation `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateUserStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDisassociateUserStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDisassociateUserStackInput"}

	if s.UserStackAssociations == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserStackAssociations"))
	}
	if s.UserStackAssociations != nil && len(s.UserStackAssociations) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserStackAssociations", 1))
	}
	if s.UserStackAssociations != nil {
		for i, v := range s.UserStackAssociations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserStackAssociations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDisassociateUserStackOutput struct {
	_ struct{} `type:"structure"`

	// The list of UserStackAssociationError objects.
	Errors []UserStackAssociationError `locationName:"errors" type:"list"`
}

// String returns the string representation
func (s BatchDisassociateUserStackOutput) String() string {
	return awsutil.Prettify(s)
}
