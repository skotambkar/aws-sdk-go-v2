// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyReplicationSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// A description for the replication instance subnet group.
	ReplicationSubnetGroupDescription *string `type:"string"`

	// The name of the replication instance subnet group.
	//
	// ReplicationSubnetGroupIdentifier is a required field
	ReplicationSubnetGroupIdentifier *string `type:"string" required:"true"`

	// A list of subnet IDs.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyReplicationSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReplicationSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyReplicationSubnetGroupInput"}

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

type ModifyReplicationSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// The modified replication subnet group.
	ReplicationSubnetGroup *ReplicationSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s ModifyReplicationSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}
