// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workmail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opPutMailboxPermissions = "PutMailboxPermissions"

// PutMailboxPermissionsRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Sets permissions for a user, group, or resource. This replaces any pre-existing
// permissions.
//
//    // Example sending a request using PutMailboxPermissionsRequest.
//    req := client.PutMailboxPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/PutMailboxPermissions
func (c *Client) PutMailboxPermissionsRequest(input *types.PutMailboxPermissionsInput) PutMailboxPermissionsRequest {
	op := &aws.Operation{
		Name:       opPutMailboxPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutMailboxPermissionsInput{}
	}

	req := c.newRequest(op, input, &types.PutMailboxPermissionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutMailboxPermissionsMarshaler{Input: input}.GetNamedBuildHandler())

	return PutMailboxPermissionsRequest{Request: req, Input: input, Copy: c.PutMailboxPermissionsRequest}
}

// PutMailboxPermissionsRequest is the request type for the
// PutMailboxPermissions API operation.
type PutMailboxPermissionsRequest struct {
	*aws.Request
	Input *types.PutMailboxPermissionsInput
	Copy  func(*types.PutMailboxPermissionsInput) PutMailboxPermissionsRequest
}

// Send marshals and sends the PutMailboxPermissions API request.
func (r PutMailboxPermissionsRequest) Send(ctx context.Context) (*PutMailboxPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutMailboxPermissionsResponse{
		PutMailboxPermissionsOutput: r.Request.Data.(*types.PutMailboxPermissionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutMailboxPermissionsResponse is the response type for the
// PutMailboxPermissions API operation.
type PutMailboxPermissionsResponse struct {
	*types.PutMailboxPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutMailboxPermissions request.
func (r *PutMailboxPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
