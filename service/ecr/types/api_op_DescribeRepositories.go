// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRepositoriesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of repository results returned by DescribeRepositories
	// in paginated output. When this parameter is used, DescribeRepositories only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeRepositories request with the returned nextToken value. This
	// value can be between 1 and 1000. If this parameter is not used, then DescribeRepositories
	// returns up to 100 results and a nextToken value, if applicable. This option
	// cannot be used when you specify repositories with repositoryNames.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeRepositories
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return. This option cannot be used when you specify repositories with
	// repositoryNames.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The AWS account ID associated with the registry that contains the repositories
	// to be described. If you do not specify a registry, the default registry is
	// assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// A list of repositories to describe. If this parameter is omitted, then all
	// repositories in a registry are described.
	RepositoryNames []string `locationName:"repositoryNames" min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeRepositoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRepositoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRepositoriesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.RepositoryNames != nil && len(s.RepositoryNames) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryNames", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRepositoriesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future DescribeRepositories request.
	// When the results of a DescribeRepositories request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of repository objects corresponding to valid repositories.
	Repositories []Repository `locationName:"repositories" type:"list"`
}

// String returns the string representation
func (s DescribeRepositoriesOutput) String() string {
	return awsutil.Prettify(s)
}