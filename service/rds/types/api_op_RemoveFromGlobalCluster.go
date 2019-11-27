// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveFromGlobalClusterInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) identifying the cluster that was detached
	// from the Aurora global database cluster.
	DbClusterIdentifier *string `type:"string"`

	// The cluster identifier to detach from the Aurora global database cluster.
	GlobalClusterIdentifier *string `type:"string"`
}

// String returns the string representation
func (s RemoveFromGlobalClusterInput) String() string {
	return awsutil.Prettify(s)
}

type RemoveFromGlobalClusterOutput struct {
	_ struct{} `type:"structure"`

	// A data type representing an Aurora global database.
	GlobalCluster *GlobalCluster `type:"structure"`
}

// String returns the string representation
func (s RemoveFromGlobalClusterOutput) String() string {
	return awsutil.Prettify(s)
}
