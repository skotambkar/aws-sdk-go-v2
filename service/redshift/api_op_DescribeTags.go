// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opDescribeTags = "DescribeTags"

// DescribeTagsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns a list of tags. You can return tags from a specific resource by specifying
// an ARN, or you can return all tags for a given type of resource, such as
// clusters, snapshots, and so on.
//
// The following are limitations for DescribeTags:
//
//    * You cannot specify an ARN and a resource-type value together in the
//    same request.
//
//    * You cannot use the MaxRecords and Marker parameters together with the
//    ARN parameter.
//
//    * The MaxRecords parameter can be a range from 10 to 50 results to return
//    in a request.
//
// If you specify both tag keys and tag values in the same request, Amazon Redshift
// returns all resources that match any combination of the specified keys and
// values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all resources that have any combination of
// those values are returned.
//
// If both tag keys and values are omitted from the request, resources are returned
// regardless of whether they have tag keys or values associated with them.
//
//    // Example sending a request using DescribeTagsRequest.
//    req := client.DescribeTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeTags
func (c *Client) DescribeTagsRequest(input *types.DescribeTagsInput) DescribeTagsRequest {
	op := &aws.Operation{
		Name:       opDescribeTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeTagsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTagsOutput{})
	return DescribeTagsRequest{Request: req, Input: input, Copy: c.DescribeTagsRequest}
}

// DescribeTagsRequest is the request type for the
// DescribeTags API operation.
type DescribeTagsRequest struct {
	*aws.Request
	Input *types.DescribeTagsInput
	Copy  func(*types.DescribeTagsInput) DescribeTagsRequest
}

// Send marshals and sends the DescribeTags API request.
func (r DescribeTagsRequest) Send(ctx context.Context) (*DescribeTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTagsResponse{
		DescribeTagsOutput: r.Request.Data.(*types.DescribeTagsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTagsResponse is the response type for the
// DescribeTags API operation.
type DescribeTagsResponse struct {
	*types.DescribeTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTags request.
func (r *DescribeTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
