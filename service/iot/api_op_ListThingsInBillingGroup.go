// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListThingsInBillingGroup = "ListThingsInBillingGroup"

// ListThingsInBillingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the things you have added to the given billing group.
//
//    // Example sending a request using ListThingsInBillingGroupRequest.
//    req := client.ListThingsInBillingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListThingsInBillingGroupRequest(input *types.ListThingsInBillingGroupInput) ListThingsInBillingGroupRequest {
	op := &aws.Operation{
		Name:       opListThingsInBillingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/billing-groups/{billingGroupName}/things",
	}

	if input == nil {
		input = &types.ListThingsInBillingGroupInput{}
	}

	req := c.newRequest(op, input, &types.ListThingsInBillingGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListThingsInBillingGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return ListThingsInBillingGroupRequest{Request: req, Input: input, Copy: c.ListThingsInBillingGroupRequest}
}

// ListThingsInBillingGroupRequest is the request type for the
// ListThingsInBillingGroup API operation.
type ListThingsInBillingGroupRequest struct {
	*aws.Request
	Input *types.ListThingsInBillingGroupInput
	Copy  func(*types.ListThingsInBillingGroupInput) ListThingsInBillingGroupRequest
}

// Send marshals and sends the ListThingsInBillingGroup API request.
func (r ListThingsInBillingGroupRequest) Send(ctx context.Context) (*ListThingsInBillingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThingsInBillingGroupResponse{
		ListThingsInBillingGroupOutput: r.Request.Data.(*types.ListThingsInBillingGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThingsInBillingGroupResponse is the response type for the
// ListThingsInBillingGroup API operation.
type ListThingsInBillingGroupResponse struct {
	*types.ListThingsInBillingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThingsInBillingGroup request.
func (r *ListThingsInBillingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
