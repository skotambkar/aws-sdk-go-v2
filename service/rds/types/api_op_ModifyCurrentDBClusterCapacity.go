// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyCurrentDBClusterCapacityInput struct {
	_ struct{} `type:"structure"`

	// The DB cluster capacity.
	//
	// When you change the capacity of a paused Aurora Serverless DB cluster, it
	// automatically resumes.
	//
	// Constraints:
	//
	//    * For Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64,
	//    128, and 256.
	//
	//    * For Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64,
	//    192, and 384.
	Capacity *int64 `type:"integer"`

	// The DB cluster identifier for the cluster being modified. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DB cluster.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The amount of time, in seconds, that Aurora Serverless tries to find a scaling
	// point to perform seamless scaling before enforcing the timeout action. The
	// default is 300.
	//
	//    * Value must be from 10 through 600.
	SecondsBeforeTimeout *int64 `type:"integer"`

	// The action to take when the timeout is reached, either ForceApplyCapacityChange
	// or RollbackCapacityChange.
	//
	// ForceApplyCapacityChange, the default, sets the capacity to the specified
	// value as soon as possible.
	//
	// RollbackCapacityChange ignores the capacity change if a scaling point isn't
	// found in the timeout period.
	TimeoutAction *string `type:"string"`
}

// String returns the string representation
func (s ModifyCurrentDBClusterCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyCurrentDBClusterCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyCurrentDBClusterCapacityInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyCurrentDBClusterCapacityOutput struct {
	_ struct{} `type:"structure"`

	// The current capacity of the DB cluster.
	CurrentCapacity *int64 `type:"integer"`

	// A user-supplied DB cluster identifier. This identifier is the unique key
	// that identifies a DB cluster.
	DBClusterIdentifier *string `type:"string"`

	// A value that specifies the capacity that the DB cluster scales to next.
	PendingCapacity *int64 `type:"integer"`

	// The number of seconds before a call to ModifyCurrentDBClusterCapacity times
	// out.
	SecondsBeforeTimeout *int64 `type:"integer"`

	// The timeout action of a call to ModifyCurrentDBClusterCapacity, either ForceApplyCapacityChange
	// or RollbackCapacityChange.
	TimeoutAction *string `type:"string"`
}

// String returns the string representation
func (s ModifyCurrentDBClusterCapacityOutput) String() string {
	return awsutil.Prettify(s)
}