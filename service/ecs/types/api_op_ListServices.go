// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type ListServicesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the services to list. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The launch type for the services to list.
	LaunchType enums.LaunchType `locationName:"launchType" type:"string" enum:"true"`

	// The maximum number of service results returned by ListServices in paginated
	// output. When this parameter is used, ListServices only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListServices
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListServices returns up to 10 results
	// and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a ListServices request indicating that
	// more results are available to fulfill the request and further calls will
	// be needed. If maxResults was provided, it is possible the number of results
	// to be fewer than maxResults.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The scheduling strategy for services to list.
	SchedulingStrategy enums.SchedulingStrategy `locationName:"schedulingStrategy" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListServicesInput) String() string {
	return awsutil.Prettify(s)
}

type ListServicesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of full ARN entries for each service associated with the specified
	// cluster.
	ServiceArns []string `locationName:"serviceArns" type:"list"`
}

// String returns the string representation
func (s ListServicesOutput) String() string {
	return awsutil.Prettify(s)
}
