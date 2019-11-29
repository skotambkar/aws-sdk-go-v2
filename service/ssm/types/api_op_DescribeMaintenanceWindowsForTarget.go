// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type DescribeMaintenanceWindowsForTargetInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The type of resource you want to retrieve information about. For example,
	// "INSTANCE".
	//
	// ResourceType is a required field
	ResourceType enums.MaintenanceWindowResourceType `type:"string" required:"true" enum:"true"`

	// The instance ID or key/value pair to retrieve information about.
	//
	// Targets is a required field
	Targets []Target `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowsForTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMaintenanceWindowsForTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMaintenanceWindowsForTargetInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeMaintenanceWindowsForTargetOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. (You use this token in the
	// next call.)
	NextToken *string `type:"string"`

	// Information about the maintenance window targets and tasks an instance is
	// associated with.
	WindowIdentities []MaintenanceWindowIdentityForTarget `type:"list"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowsForTargetOutput) String() string {
	return awsutil.Prettify(s)
}
