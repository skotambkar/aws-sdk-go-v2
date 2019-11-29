// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListNodegroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster that you would like to list node groups
	// in.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The maximum number of node group results returned by ListNodegroups in paginated
	// output. When you use this parameter, ListNodegroups returns only maxResults
	// results in a single page along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another ListNodegroups
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListNodegroups returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListNodegroups request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListNodegroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNodegroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNodegroupsInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListNodegroupsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListNodegroups request. When the
	// results of a ListNodegroups request exceed maxResults, you can use this value
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of all of the node groups associated with the specified cluster.
	Nodegroups []string `locationName:"nodegroups" type:"list"`
}

// String returns the string representation
func (s ListNodegroupsOutput) String() string {
	return awsutil.Prettify(s)
}
