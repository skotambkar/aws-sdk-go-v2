// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteClusterInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster to delete.
	//
	// Cluster is a required field
	Cluster *string `locationName:"cluster" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteClusterInput"}

	if s.Cluster == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cluster"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteClusterOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the deleted cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s DeleteClusterOutput) String() string {
	return awsutil.Prettify(s)
}