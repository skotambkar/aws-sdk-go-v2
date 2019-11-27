// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CopyDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier or Amazon Resource Name (ARN) for the source DB cluster parameter
	// group. For information about creating an ARN, see Constructing an ARN for
	// Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// in the Amazon Aurora User Guide.
	//
	// Constraints:
	//
	//    * Must specify a valid DB cluster parameter group.
	//
	//    * If the source DB cluster parameter group is in the same AWS Region as
	//    the copy, specify a valid DB parameter group identifier, for example my-db-cluster-param-group,
	//    or a valid ARN.
	//
	//    * If the source DB parameter group is in a different AWS Region than the
	//    copy, specify a valid DB cluster parameter group ARN, for example arn:aws:rds:us-east-1:123456789012:cluster-pg:custom-cluster-group1.
	//
	// SourceDBClusterParameterGroupIdentifier is a required field
	SourceDBClusterParameterGroupIdentifier *string `type:"string" required:"true"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A description for the copied DB cluster parameter group.
	//
	// TargetDBClusterParameterGroupDescription is a required field
	TargetDBClusterParameterGroupDescription *string `type:"string" required:"true"`

	// The identifier for the copied DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Can't be null, empty, or blank
	//
	//    * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-cluster-param-group1
	//
	// TargetDBClusterParameterGroupIdentifier is a required field
	TargetDBClusterParameterGroupIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyDBClusterParameterGroupInput"}

	if s.SourceDBClusterParameterGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBClusterParameterGroupIdentifier"))
	}

	if s.TargetDBClusterParameterGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBClusterParameterGroupDescription"))
	}

	if s.TargetDBClusterParameterGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBClusterParameterGroupIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB cluster parameter group.
	//
	// This data type is used as a response element in the DescribeDBClusterParameterGroups
	// action.
	DBClusterParameterGroup *DBClusterParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CopyDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}
