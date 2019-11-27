// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opUpdateStackInstances = "UpdateStackInstances"

// UpdateStackInstancesRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Updates the parameter values for stack instances for the specified accounts,
// within the specified regions. A stack instance refers to a stack in a specific
// account and region.
//
// You can only update stack instances in regions and accounts where they already
// exist; to create additional stack instances, use CreateStackInstances (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html).
//
// During stack set updates, any parameters overridden for a stack instance
// are not updated, but retain their overridden value.
//
// You can only update the parameter values that are specified in the stack
// set; to add or delete a parameter itself, use UpdateStackSet (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
// to update the stack set template. If you add a parameter to a template, before
// you can override the parameter value specified in the stack set you must
// first use UpdateStackSet (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
// to update all stack instances with the updated template and parameter value
// specified in the stack set. Once a stack instance has been updated with the
// new parameter, you can then override the parameter value using UpdateStackInstances.
//
//    // Example sending a request using UpdateStackInstancesRequest.
//    req := client.UpdateStackInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/UpdateStackInstances
func (c *Client) UpdateStackInstancesRequest(input *types.UpdateStackInstancesInput) UpdateStackInstancesRequest {
	op := &aws.Operation{
		Name:       opUpdateStackInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateStackInstancesInput{}
	}

	req := c.newRequest(op, input, &types.UpdateStackInstancesOutput{})
	return UpdateStackInstancesRequest{Request: req, Input: input, Copy: c.UpdateStackInstancesRequest}
}

// UpdateStackInstancesRequest is the request type for the
// UpdateStackInstances API operation.
type UpdateStackInstancesRequest struct {
	*aws.Request
	Input *types.UpdateStackInstancesInput
	Copy  func(*types.UpdateStackInstancesInput) UpdateStackInstancesRequest
}

// Send marshals and sends the UpdateStackInstances API request.
func (r UpdateStackInstancesRequest) Send(ctx context.Context) (*UpdateStackInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateStackInstancesResponse{
		UpdateStackInstancesOutput: r.Request.Data.(*types.UpdateStackInstancesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateStackInstancesResponse is the response type for the
// UpdateStackInstances API operation.
type UpdateStackInstancesResponse struct {
	*types.UpdateStackInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateStackInstances request.
func (r *UpdateStackInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
