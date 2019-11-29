// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type DescribeDocumentPermissionInput struct {
	_ struct{} `type:"structure"`

	// The name of the document for which you are the owner.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The permission type for the document. The permission type can be Share.
	//
	// PermissionType is a required field
	PermissionType enums.DocumentPermissionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DescribeDocumentPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDocumentPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDocumentPermissionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.PermissionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PermissionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDocumentPermissionOutput struct {
	_ struct{} `type:"structure"`

	// The account IDs that have permission to use this document. The ID can be
	// either an AWS account or All.
	AccountIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeDocumentPermissionOutput) String() string {
	return awsutil.Prettify(s)
}
