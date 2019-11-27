// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the subnet group to delete.
	//
	// SubnetGroupName is a required field
	SubnetGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSubnetGroupInput"}

	if s.SubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// A user-specified message for this action (i.e., a reason for deleting the
	// subnet group).
	DeletionMessage *string `type:"string"`
}

// String returns the string representation
func (s DeleteSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}
