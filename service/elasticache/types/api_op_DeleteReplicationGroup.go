// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DeleteReplicationGroup operation.
type DeleteReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of a final node group (shard) snapshot. ElastiCache creates the
	// snapshot from the primary node in the cluster, rather than one of the replicas;
	// this is to ensure that it captures the freshest data. After the final snapshot
	// is taken, the replication group is immediately deleted.
	FinalSnapshotIdentifier *string `type:"string"`

	// The identifier for the cluster to be deleted. This parameter is not case
	// sensitive.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`

	// If set to true, all of the read replicas are deleted, but the primary node
	// is retained.
	RetainPrimaryCluster *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReplicationGroupInput"}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s DeleteReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}
