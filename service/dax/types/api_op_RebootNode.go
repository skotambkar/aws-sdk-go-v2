// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RebootNodeInput struct {
	_ struct{} `type:"structure"`

	// The name of the DAX cluster containing the node to be rebooted.
	//
	// ClusterName is a required field
	ClusterName *string `type:"string" required:"true"`

	// The system-assigned ID of the node to be rebooted.
	//
	// NodeId is a required field
	NodeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebootNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootNodeInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebootNodeOutput struct {
	_ struct{} `type:"structure"`

	// A description of the DAX cluster after a node has been rebooted.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s RebootNodeOutput) String() string {
	return awsutil.Prettify(s)
}
