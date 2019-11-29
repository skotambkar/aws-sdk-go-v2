// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

const opDeleteResourcePolicy = "DeleteResourcePolicy"

// DeleteResourcePolicyRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Deletes a resource policy from this account. This revokes the access of the
// identities in that policy to put log events to this account.
//
//    // Example sending a request using DeleteResourcePolicyRequest.
//    req := client.DeleteResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteResourcePolicy
func (c *Client) DeleteResourcePolicyRequest(input *types.DeleteResourcePolicyInput) DeleteResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteResourcePolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteResourcePolicyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteResourcePolicyRequest{Request: req, Input: input, Copy: c.DeleteResourcePolicyRequest}
}

// DeleteResourcePolicyRequest is the request type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyRequest struct {
	*aws.Request
	Input *types.DeleteResourcePolicyInput
	Copy  func(*types.DeleteResourcePolicyInput) DeleteResourcePolicyRequest
}

// Send marshals and sends the DeleteResourcePolicy API request.
func (r DeleteResourcePolicyRequest) Send(ctx context.Context) (*DeleteResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourcePolicyResponse{
		DeleteResourcePolicyOutput: r.Request.Data.(*types.DeleteResourcePolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourcePolicyResponse is the response type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyResponse struct {
	*types.DeleteResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourcePolicy request.
func (r *DeleteResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
