// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRemediationExecutionStatusInput struct {
	_ struct{} `type:"structure"`

	// A list of AWS Config rule names.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// The maximum number of RemediationExecutionStatuses returned on each page.
	// The default is maximum. If you specify 0, AWS Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	ResourceKeys []ResourceKey `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeRemediationExecutionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRemediationExecutionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRemediationExecutionStatusInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}
	if s.ResourceKeys != nil && len(s.ResourceKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceKeys", 1))
	}
	if s.ResourceKeys != nil {
		for i, v := range s.ResourceKeys {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceKeys", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRemediationExecutionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// Returns a list of remediation execution statuses objects.
	RemediationExecutionStatuses []RemediationExecutionStatus `type:"list"`
}

// String returns the string representation
func (s DescribeRemediationExecutionStatusOutput) String() string {
	return awsutil.Prettify(s)
}
