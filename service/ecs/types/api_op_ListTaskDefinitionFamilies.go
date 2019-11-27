// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type ListTaskDefinitionFamiliesInput struct {
	_ struct{} `type:"structure"`

	// The familyPrefix is a string that is used to filter the results of ListTaskDefinitionFamilies.
	// If you specify a familyPrefix, only task definition family names that begin
	// with the familyPrefix string are returned.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition family results returned by ListTaskDefinitionFamilies
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitionFamilies request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListTaskDefinitionFamilies returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitionFamilies
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The task definition family status with which to filter the ListTaskDefinitionFamilies
	// results. By default, both ACTIVE and INACTIVE task definition families are
	// listed. If this parameter is set to ACTIVE, only task definition families
	// that have an ACTIVE task definition revision are returned. If this parameter
	// is set to INACTIVE, only task definition families that do not have any ACTIVE
	// task definition revisions are returned. If you paginate the resulting output,
	// be sure to keep the status value constant in each subsequent request.
	Status enums.TaskDefinitionFamilyStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesInput) String() string {
	return awsutil.Prettify(s)
}

type ListTaskDefinitionFamiliesOutput struct {
	_ struct{} `type:"structure"`

	// The list of task definition family names that match the ListTaskDefinitionFamilies
	// request.
	Families []string `locationName:"families" type:"list"`

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults,
	// this value can be used to retrieve the next page of results. This value is
	// null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesOutput) String() string {
	return awsutil.Prettify(s)
}
