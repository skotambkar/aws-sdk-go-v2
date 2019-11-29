// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opDescribeResize = "DescribeResize"

// DescribeResizeRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns information about the last resize operation for the specified cluster.
// If no resize operation has ever been initiated for the specified cluster,
// a HTTP 404 error is returned. If a resize operation was initiated and completed,
// the status of the resize remains as SUCCEEDED until the next resize.
//
// A resize operation can be requested using ModifyCluster and specifying a
// different number or type of nodes for the cluster.
//
//    // Example sending a request using DescribeResizeRequest.
//    req := client.DescribeResizeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeResize
func (c *Client) DescribeResizeRequest(input *types.DescribeResizeInput) DescribeResizeRequest {
	op := &aws.Operation{
		Name:       opDescribeResize,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeResizeInput{}
	}

	req := c.newRequest(op, input, &types.DescribeResizeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeResizeMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeResizeRequest{Request: req, Input: input, Copy: c.DescribeResizeRequest}
}

// DescribeResizeRequest is the request type for the
// DescribeResize API operation.
type DescribeResizeRequest struct {
	*aws.Request
	Input *types.DescribeResizeInput
	Copy  func(*types.DescribeResizeInput) DescribeResizeRequest
}

// Send marshals and sends the DescribeResize API request.
func (r DescribeResizeRequest) Send(ctx context.Context) (*DescribeResizeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeResizeResponse{
		DescribeResizeOutput: r.Request.Data.(*types.DescribeResizeOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeResizeResponse is the response type for the
// DescribeResize API operation.
type DescribeResizeResponse struct {
	*types.DescribeResizeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeResize request.
func (r *DescribeResizeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
