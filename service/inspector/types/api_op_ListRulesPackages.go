// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRulesPackagesInput struct {
	_ struct{} `type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListRulesPackages action. Subsequent
	// calls to the action fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListRulesPackagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRulesPackagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRulesPackagesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRulesPackagesOutput struct {
	_ struct{} `type:"structure"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The list of ARNs that specifies the rules packages returned by the action.
	//
	// RulesPackageArns is a required field
	RulesPackageArns []string `locationName:"rulesPackageArns" type:"list" required:"true"`
}

// String returns the string representation
func (s ListRulesPackagesOutput) String() string {
	return awsutil.Prettify(s)
}