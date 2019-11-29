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

const opPutDestinationPolicy = "PutDestinationPolicy"

// PutDestinationPolicyRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Creates or updates an access policy associated with an existing destination.
// An access policy is an IAM policy document (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html)
// that is used to authorize claims to register a subscription filter against
// a given destination.
//
//    // Example sending a request using PutDestinationPolicyRequest.
//    req := client.PutDestinationPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutDestinationPolicy
func (c *Client) PutDestinationPolicyRequest(input *types.PutDestinationPolicyInput) PutDestinationPolicyRequest {
	op := &aws.Operation{
		Name:       opPutDestinationPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutDestinationPolicyInput{}
	}

	req := c.newRequest(op, input, &types.PutDestinationPolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutDestinationPolicyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutDestinationPolicyRequest{Request: req, Input: input, Copy: c.PutDestinationPolicyRequest}
}

// PutDestinationPolicyRequest is the request type for the
// PutDestinationPolicy API operation.
type PutDestinationPolicyRequest struct {
	*aws.Request
	Input *types.PutDestinationPolicyInput
	Copy  func(*types.PutDestinationPolicyInput) PutDestinationPolicyRequest
}

// Send marshals and sends the PutDestinationPolicy API request.
func (r PutDestinationPolicyRequest) Send(ctx context.Context) (*PutDestinationPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDestinationPolicyResponse{
		PutDestinationPolicyOutput: r.Request.Data.(*types.PutDestinationPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDestinationPolicyResponse is the response type for the
// PutDestinationPolicy API operation.
type PutDestinationPolicyResponse struct {
	*types.PutDestinationPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDestinationPolicy request.
func (r *PutDestinationPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
