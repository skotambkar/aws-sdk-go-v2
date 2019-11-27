// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListNodesInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListNodesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNodesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNodesInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about nodes in the cluster.
type ListNodesOutput struct {
	_ struct{} `type:"structure"`

	// The paginated results marker. When the result of a ListNodes operation is
	// truncated, the call returns NextToken in the response. To get another batch
	// of nodes, provide this token in your next request.
	NextToken *string `locationName:"nextToken" type:"string"`

	// List containing a NodeInfo object.
	NodeInfoList []NodeInfo `locationName:"nodeInfoList" type:"list"`
}

// String returns the string representation
func (s ListNodesOutput) String() string {
	return awsutil.Prettify(s)
}
