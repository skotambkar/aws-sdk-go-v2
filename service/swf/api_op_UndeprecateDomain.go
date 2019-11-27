// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

const opUndeprecateDomain = "UndeprecateDomain"

// UndeprecateDomainRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Undeprecates a previously deprecated domain. After a domain has been undeprecated
// it can be used to create new workflow executions or register new types.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using UndeprecateDomainRequest.
//    req := client.UndeprecateDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UndeprecateDomainRequest(input *types.UndeprecateDomainInput) UndeprecateDomainRequest {
	op := &aws.Operation{
		Name:       opUndeprecateDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UndeprecateDomainInput{}
	}

	req := c.newRequest(op, input, &types.UndeprecateDomainOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UndeprecateDomainRequest{Request: req, Input: input, Copy: c.UndeprecateDomainRequest}
}

// UndeprecateDomainRequest is the request type for the
// UndeprecateDomain API operation.
type UndeprecateDomainRequest struct {
	*aws.Request
	Input *types.UndeprecateDomainInput
	Copy  func(*types.UndeprecateDomainInput) UndeprecateDomainRequest
}

// Send marshals and sends the UndeprecateDomain API request.
func (r UndeprecateDomainRequest) Send(ctx context.Context) (*UndeprecateDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UndeprecateDomainResponse{
		UndeprecateDomainOutput: r.Request.Data.(*types.UndeprecateDomainOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UndeprecateDomainResponse is the response type for the
// UndeprecateDomain API operation.
type UndeprecateDomainResponse struct {
	*types.UndeprecateDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UndeprecateDomain request.
func (r *UndeprecateDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
