// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type FailoverDBClusterInput struct {
	_ struct{} `type:"structure"`

	// A DB cluster identifier to force a failover for. This parameter is not case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the instance to promote to the primary instance.
	//
	// You must specify the instance identifier for an Aurora Replica in the DB
	// cluster. For example, mydbcluster-replica1.
	TargetDBInstanceIdentifier *string `type:"string"`
}

// String returns the string representation
func (s FailoverDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FailoverDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FailoverDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type FailoverDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s FailoverDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}
