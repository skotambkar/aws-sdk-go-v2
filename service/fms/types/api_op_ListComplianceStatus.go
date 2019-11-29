// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListComplianceStatusInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of PolicyComplianceStatus objects that you want AWS
	// Firewall Manager to return for this request. If you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, the response includes
	// a NextToken value that you can use to get another batch of PolicyComplianceStatus
	// objects.
	MaxResults *int64 `min:"1" type:"integer"`

	// If you specify a value for MaxResults and you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, AWS Firewall Manager
	// returns a NextToken value in the response that allows you to list another
	// group of PolicyComplianceStatus objects. For the second and subsequent ListComplianceStatus
	// requests, specify the value of NextToken from the previous response to get
	// information about another batch of PolicyComplianceStatus objects.
	NextToken *string `min:"1" type:"string"`

	// The ID of the AWS Firewall Manager policy that you want the details for.
	//
	// PolicyId is a required field
	PolicyId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s ListComplianceStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListComplianceStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListComplianceStatusInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}
	if s.PolicyId != nil && len(*s.PolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListComplianceStatusOutput struct {
	_ struct{} `type:"structure"`

	// If you have more PolicyComplianceStatus objects than the number that you
	// specified for MaxResults in the request, the response includes a NextToken
	// value. To list more PolicyComplianceStatus objects, submit another ListComplianceStatus
	// request, and specify the NextToken value from the response in the NextToken
	// value in the next request.
	NextToken *string `min:"1" type:"string"`

	// An array of PolicyComplianceStatus objects.
	PolicyComplianceStatusList []PolicyComplianceStatus `type:"list"`
}

// String returns the string representation
func (s ListComplianceStatusOutput) String() string {
	return awsutil.Prettify(s)
}