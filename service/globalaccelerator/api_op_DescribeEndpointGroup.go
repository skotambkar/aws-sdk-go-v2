// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
)

const opDescribeEndpointGroup = "DescribeEndpointGroup"

// DescribeEndpointGroupRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Describe an endpoint group.
//
//    // Example sending a request using DescribeEndpointGroupRequest.
//    req := client.DescribeEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/DescribeEndpointGroup
func (c *Client) DescribeEndpointGroupRequest(input *types.DescribeEndpointGroupInput) DescribeEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpointGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEndpointGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeEndpointGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeEndpointGroupRequest{Request: req, Input: input, Copy: c.DescribeEndpointGroupRequest}
}

// DescribeEndpointGroupRequest is the request type for the
// DescribeEndpointGroup API operation.
type DescribeEndpointGroupRequest struct {
	*aws.Request
	Input *types.DescribeEndpointGroupInput
	Copy  func(*types.DescribeEndpointGroupInput) DescribeEndpointGroupRequest
}

// Send marshals and sends the DescribeEndpointGroup API request.
func (r DescribeEndpointGroupRequest) Send(ctx context.Context) (*DescribeEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointGroupResponse{
		DescribeEndpointGroupOutput: r.Request.Data.(*types.DescribeEndpointGroupOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEndpointGroupResponse is the response type for the
// DescribeEndpointGroup API operation.
type DescribeEndpointGroupResponse struct {
	*types.DescribeEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpointGroup request.
func (r *DescribeEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
