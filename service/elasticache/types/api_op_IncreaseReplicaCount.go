// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type IncreaseReplicaCountInput struct {
	_ struct{} `type:"structure"`

	// If True, the number of replica nodes is increased immediately. ApplyImmediately=False
	// is not currently supported.
	//
	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	// The number of read replica nodes you want at the completion of this operation.
	// For Redis (cluster mode disabled) replication groups, this is the number
	// of replica nodes in the replication group. For Redis (cluster mode enabled)
	// replication groups, this is the number of replica nodes in each of the replication
	// group's node groups.
	NewReplicaCount *int64 `type:"integer"`

	// A list of ConfigureShard objects that can be used to configure each shard
	// in a Redis (cluster mode enabled) replication group. The ConfigureShard has
	// three members: NewReplicaCount, NodeGroupId, and PreferredAvailabilityZones.
	ReplicaConfiguration []ConfigureShard `locationNameList:"ConfigureShard" type:"list"`

	// The id of the replication group to which you want to add replica nodes.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s IncreaseReplicaCountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IncreaseReplicaCountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "IncreaseReplicaCountInput"}

	if s.ApplyImmediately == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyImmediately"))
	}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}
	if s.ReplicaConfiguration != nil {
		for i, v := range s.ReplicaConfiguration {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ReplicaConfiguration", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IncreaseReplicaCountOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s IncreaseReplicaCountOutput) String() string {
	return awsutil.Prettify(s)
}