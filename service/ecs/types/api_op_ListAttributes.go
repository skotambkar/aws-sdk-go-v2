// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type ListAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute with which to filter the results.
	AttributeName *string `locationName:"attributeName" type:"string"`

	// The value of the attribute with which to filter results. You must also specify
	// an attribute name to use this parameter.
	AttributeValue *string `locationName:"attributeValue" type:"string"`

	// The short name or full Amazon Resource Name (ARN) of the cluster to list
	// attributes. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of cluster results returned by ListAttributes in paginated
	// output. When this parameter is used, ListAttributes only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListAttributes
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListAttributes returns up to 100
	// results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a ListAttributes request indicating that
	// more results are available to fulfill the request and further calls will
	// be needed. If maxResults was provided, it is possible the number of results
	// to be fewer than maxResults.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The type of the target with which to list attributes.
	//
	// TargetType is a required field
	TargetType enums.TargetType `locationName:"targetType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ListAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttributesInput"}
	if len(s.TargetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TargetType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A list of attribute objects that meet the criteria of the request.
	Attributes []Attribute `locationName:"attributes" type:"list"`

	// The nextToken value to include in a future ListAttributes request. When the
	// results of a ListAttributes request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAttributesOutput) String() string {
	return awsutil.Prettify(s)
}
