// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentsInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of deployments for the requested groups.
	Deployments []Deployment `type:"list"`

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}
