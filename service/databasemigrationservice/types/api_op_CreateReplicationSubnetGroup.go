// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateReplicationSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The description for the subnet group.
	//
	// ReplicationSubnetGroupDescription is a required field
	ReplicationSubnetGroupDescription *string `type:"string" required:"true"`

	// The name for the replication subnet group. This value is stored as a lowercase
	// string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters, periods,
	// spaces, underscores, or hyphens. Must not be "default".
	//
	// Example: mySubnetgroup
	//
	// ReplicationSubnetGroupIdentifier is a required field
	ReplicationSubnetGroupIdentifier *string `type:"string" required:"true"`

	// One or more subnet IDs to be assigned to the subnet group.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`

	// One or more tags to be assigned to the subnet group.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateReplicationSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReplicationSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReplicationSubnetGroupInput"}

	if s.ReplicationSubnetGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationSubnetGroupDescription"))
	}

	if s.ReplicationSubnetGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationSubnetGroupIdentifier"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateReplicationSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// The replication subnet group that was created.
	ReplicationSubnetGroup *ReplicationSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s CreateReplicationSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}
