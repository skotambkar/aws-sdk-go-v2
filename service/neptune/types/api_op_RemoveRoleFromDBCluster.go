// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveRoleFromDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster to disassociate the IAM role from.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB
	// cluster, for example arn:aws:iam::123456789012:role/NeptuneAccessRole.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveRoleFromDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveRoleFromDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveRoleFromDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveRoleFromDBClusterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveRoleFromDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}