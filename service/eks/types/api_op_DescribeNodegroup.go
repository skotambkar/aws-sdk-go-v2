// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeNodegroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster associated with the node group.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The name of the node group to describe.
	//
	// NodegroupName is a required field
	NodegroupName *string `location:"uri" locationName:"nodegroupName" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeNodegroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNodegroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNodegroupInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodegroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodegroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeNodegroupOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your node group.
	Nodegroup *Nodegroup `locationName:"nodegroup" type:"structure"`
}

// String returns the string representation
func (s DescribeNodegroupOutput) String() string {
	return awsutil.Prettify(s)
}