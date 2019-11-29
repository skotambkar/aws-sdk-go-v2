// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opRequestSpotInstances = "RequestSpotInstances"

// RequestSpotInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a Spot Instance request.
//
// For more information, see Spot Instance Requests (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html)
// in the Amazon EC2 User Guide for Linux Instances.
//
//    // Example sending a request using RequestSpotInstancesRequest.
//    req := client.RequestSpotInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotInstances
func (c *Client) RequestSpotInstancesRequest(input *types.RequestSpotInstancesInput) RequestSpotInstancesRequest {
	op := &aws.Operation{
		Name:       opRequestSpotInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RequestSpotInstancesInput{}
	}

	req := c.newRequest(op, input, &types.RequestSpotInstancesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.RequestSpotInstancesMarshaler{Input: input}.GetNamedBuildHandler())

	return RequestSpotInstancesRequest{Request: req, Input: input, Copy: c.RequestSpotInstancesRequest}
}

// RequestSpotInstancesRequest is the request type for the
// RequestSpotInstances API operation.
type RequestSpotInstancesRequest struct {
	*aws.Request
	Input *types.RequestSpotInstancesInput
	Copy  func(*types.RequestSpotInstancesInput) RequestSpotInstancesRequest
}

// Send marshals and sends the RequestSpotInstances API request.
func (r RequestSpotInstancesRequest) Send(ctx context.Context) (*RequestSpotInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestSpotInstancesResponse{
		RequestSpotInstancesOutput: r.Request.Data.(*types.RequestSpotInstancesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestSpotInstancesResponse is the response type for the
// RequestSpotInstances API operation.
type RequestSpotInstancesResponse struct {
	*types.RequestSpotInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestSpotInstances request.
func (r *RequestSpotInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
