// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/enums"
)

type ListConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// A valid configuration identified by Application Discovery Service.
	//
	// ConfigurationType is a required field
	ConfigurationType enums.ConfigurationItemType `locationName:"configurationType" type:"string" required:"true" enum:"true"`

	// You can filter the request using various logical operators and a key-value
	// format. For example:
	//
	// {"key": "serverType", "value": "webServer"}
	//
	// For a complete list of filter options and guidance about using them with
	// this action, see Querying Discovered Configuration Items (http://docs.aws.amazon.com/application-discovery/latest/APIReference/discovery-api-queries.html#ListConfigurations).
	Filters []Filter `locationName:"filters" type:"list"`

	// The total number of items to return. The maximum value is 100.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// Token to retrieve the next set of results. For example, if a previous call
	// to ListConfigurations returned 100 items, but you set ListConfigurationsRequest$maxResults
	// to 10, you received a set of 10 results along with a token. Use that token
	// in this query to get the next set of 10.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Certain filter criteria return output that can be sorted in ascending or
	// descending order. For a list of output characteristics for each filter, see
	// Using the ListConfigurations Action (http://docs.aws.amazon.com/application-discovery/latest/APIReference/discovery-api-queries.html#ListConfigurations).
	OrderBy []OrderByElement `locationName:"orderBy" type:"list"`
}

// String returns the string representation
func (s ListConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListConfigurationsInput"}
	if len(s.ConfigurationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationType"))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OrderBy != nil {
		for i, v := range s.OrderBy {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OrderBy", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// Returns configuration details, including the configuration ID, attribute
	// names, and attribute values.
	Configurations []map[string]string `locationName:"configurations" type:"list"`

	// Token to retrieve the next set of results. For example, if your call to ListConfigurations
	// returned 100 items, but you set ListConfigurationsRequest$maxResults to 10,
	// you received a set of 10 results along with this token. Use this token in
	// the next query to retrieve the next set of 10.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}
