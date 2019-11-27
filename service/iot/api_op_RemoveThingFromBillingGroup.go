// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opRemoveThingFromBillingGroup = "RemoveThingFromBillingGroup"

// RemoveThingFromBillingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Removes the given thing from the billing group.
//
//    // Example sending a request using RemoveThingFromBillingGroupRequest.
//    req := client.RemoveThingFromBillingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RemoveThingFromBillingGroupRequest(input *types.RemoveThingFromBillingGroupInput) RemoveThingFromBillingGroupRequest {
	op := &aws.Operation{
		Name:       opRemoveThingFromBillingGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/billing-groups/removeThingFromBillingGroup",
	}

	if input == nil {
		input = &types.RemoveThingFromBillingGroupInput{}
	}

	req := c.newRequest(op, input, &types.RemoveThingFromBillingGroupOutput{})
	return RemoveThingFromBillingGroupRequest{Request: req, Input: input, Copy: c.RemoveThingFromBillingGroupRequest}
}

// RemoveThingFromBillingGroupRequest is the request type for the
// RemoveThingFromBillingGroup API operation.
type RemoveThingFromBillingGroupRequest struct {
	*aws.Request
	Input *types.RemoveThingFromBillingGroupInput
	Copy  func(*types.RemoveThingFromBillingGroupInput) RemoveThingFromBillingGroupRequest
}

// Send marshals and sends the RemoveThingFromBillingGroup API request.
func (r RemoveThingFromBillingGroupRequest) Send(ctx context.Context) (*RemoveThingFromBillingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveThingFromBillingGroupResponse{
		RemoveThingFromBillingGroupOutput: r.Request.Data.(*types.RemoveThingFromBillingGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveThingFromBillingGroupResponse is the response type for the
// RemoveThingFromBillingGroup API operation.
type RemoveThingFromBillingGroupResponse struct {
	*types.RemoveThingFromBillingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveThingFromBillingGroup request.
func (r *RemoveThingFromBillingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
