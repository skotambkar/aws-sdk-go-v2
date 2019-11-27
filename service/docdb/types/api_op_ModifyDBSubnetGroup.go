// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to ModifyDBSubnetGroup.
type ModifyDBSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The description for the DB subnet group.
	DBSubnetGroupDescription *string `type:"string"`

	// The name for the DB subnet group. This value is stored as a lowercase string.
	// You can't modify the default subnet group.
	//
	// Constraints: Must match the name of an existing DBSubnetGroup. Must not be
	// default.
	//
	// Example: mySubnetgroup
	//
	// DBSubnetGroupName is a required field
	DBSubnetGroupName *string `type:"string" required:"true"`

	// The Amazon EC2 subnet IDs for the DB subnet group.
	//
	// SubnetIds is a required field
	SubnetIds []string `locationNameList:"SubnetIdentifier" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyDBSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBSubnetGroupInput"}

	if s.DBSubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSubnetGroupName"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a DB subnet group.
	DBSubnetGroup *DBSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s ModifyDBSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}
