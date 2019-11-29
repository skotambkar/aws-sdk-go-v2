// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const opAddPermission = "AddPermission"

// AddPermissionRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Adds a permission to a queue for a specific principal (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P).
// This allows sharing access to the queue.
//
// When you create a queue, you have full control access rights for the queue.
// Only you, the owner of the queue, can grant or deny permissions to the queue.
// For more information about these permissions, see Allow Developers to Write
// Messages to a Shared Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon Simple Queue Service Developer Guide.
//
//    * AddPermission generates a policy for you. You can use SetQueueAttributes
//    to upload your policy. For more information, see Using Custom Policies
//    with the Amazon SQS Access Policy Language (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html)
//    in the Amazon Simple Queue Service Developer Guide.
//
//    * An Amazon SQS policy can have a maximum of 7 actions.
//
//    * To remove the ability to change queue permissions, you must deny permission
//    to the AddPermission, RemovePermission, and SetQueueAttributes actions
//    in your IAM policy.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
//    // Example sending a request using AddPermissionRequest.
//    req := client.AddPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/AddPermission
func (c *Client) AddPermissionRequest(input *types.AddPermissionInput) AddPermissionRequest {
	op := &aws.Operation{
		Name:       opAddPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddPermissionInput{}
	}

	req := c.newRequest(op, input, &types.AddPermissionOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddPermissionRequest{Request: req, Input: input, Copy: c.AddPermissionRequest}
}

// AddPermissionRequest is the request type for the
// AddPermission API operation.
type AddPermissionRequest struct {
	*aws.Request
	Input *types.AddPermissionInput
	Copy  func(*types.AddPermissionInput) AddPermissionRequest
}

// Send marshals and sends the AddPermission API request.
func (r AddPermissionRequest) Send(ctx context.Context) (*AddPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddPermissionResponse{
		AddPermissionOutput: r.Request.Data.(*types.AddPermissionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddPermissionResponse is the response type for the
// AddPermission API operation.
type AddPermissionResponse struct {
	*types.AddPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddPermission request.
func (r *AddPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
