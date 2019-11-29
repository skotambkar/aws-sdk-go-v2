// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteVpcEndpoints = "DeleteVpcEndpoints"

// DeleteVpcEndpointsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes one or more specified VPC endpoints. Deleting a gateway endpoint
// also deletes the endpoint routes in the route tables that were associated
// with the endpoint. Deleting an interface endpoint deletes the endpoint network
// interfaces.
//
//    // Example sending a request using DeleteVpcEndpointsRequest.
//    req := client.DeleteVpcEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpcEndpoints
func (c *Client) DeleteVpcEndpointsRequest(input *types.DeleteVpcEndpointsInput) DeleteVpcEndpointsRequest {
	op := &aws.Operation{
		Name:       opDeleteVpcEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteVpcEndpointsInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVpcEndpointsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DeleteVpcEndpointsMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteVpcEndpointsRequest{Request: req, Input: input, Copy: c.DeleteVpcEndpointsRequest}
}

// DeleteVpcEndpointsRequest is the request type for the
// DeleteVpcEndpoints API operation.
type DeleteVpcEndpointsRequest struct {
	*aws.Request
	Input *types.DeleteVpcEndpointsInput
	Copy  func(*types.DeleteVpcEndpointsInput) DeleteVpcEndpointsRequest
}

// Send marshals and sends the DeleteVpcEndpoints API request.
func (r DeleteVpcEndpointsRequest) Send(ctx context.Context) (*DeleteVpcEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpcEndpointsResponse{
		DeleteVpcEndpointsOutput: r.Request.Data.(*types.DeleteVpcEndpointsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpcEndpointsResponse is the response type for the
// DeleteVpcEndpoints API operation.
type DeleteVpcEndpointsResponse struct {
	*types.DeleteVpcEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpcEndpoints request.
func (r *DeleteVpcEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
