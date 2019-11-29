// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateClusterSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name for the subnet group. Amazon Redshift stores the value as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * Must contain no more than 255 alphanumeric characters or hyphens.
	//
	//    * Must not be "Default".
	//
	//    * Must be unique for all subnet groups that are created by your AWS account.
	//
	// Example: examplesubnetgroup
	//
	// ClusterSubnetGroupName is a required field
	ClusterSubnetGroupName *string `type:"string" required:"true"`

	// A description for the subnet group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a
	// single request.
	//
	// SubnetIds is a required field
	SubnetIds []string `locationNameList:"SubnetIdentifier" type:"list" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateClusterSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterSubnetGroupInput"}

	if s.ClusterSubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterSubnetGroupName"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateClusterSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Describes a subnet group.
	ClusterSubnetGroup *ClusterSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s CreateClusterSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}
