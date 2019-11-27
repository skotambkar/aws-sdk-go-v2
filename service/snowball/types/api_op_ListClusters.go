// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListClustersInput struct {
	_ struct{} `type:"structure"`

	// The number of ClusterListEntry objects to return.
	MaxResults *int64 `type:"integer"`

	// HTTP requests are stateless. To identify what object comes "next" in the
	// list of ClusterListEntry objects, you have the option of specifying NextToken
	// as the starting point for your returned list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListClustersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListClustersInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	// Each ClusterListEntry object contains a cluster's state, a cluster's ID,
	// and other important status information.
	ClusterListEntries []ClusterListEntry `type:"list"`

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next ClusterListEntry call, your list of returned clusters
	// will start from this point in the array.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}
