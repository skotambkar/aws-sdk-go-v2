// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opRemoveAllResourcePermissions = "RemoveAllResourcePermissions"

// RemoveAllResourcePermissionsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Removes all the permissions from the specified resource.
//
//    // Example sending a request using RemoveAllResourcePermissionsRequest.
//    req := client.RemoveAllResourcePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/RemoveAllResourcePermissions
func (c *Client) RemoveAllResourcePermissionsRequest(input *types.RemoveAllResourcePermissionsInput) RemoveAllResourcePermissionsRequest {
	op := &aws.Operation{
		Name:       opRemoveAllResourcePermissions,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/resources/{ResourceId}/permissions",
	}

	if input == nil {
		input = &types.RemoveAllResourcePermissionsInput{}
	}

	req := c.newRequest(op, input, &types.RemoveAllResourcePermissionsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveAllResourcePermissionsRequest{Request: req, Input: input, Copy: c.RemoveAllResourcePermissionsRequest}
}

// RemoveAllResourcePermissionsRequest is the request type for the
// RemoveAllResourcePermissions API operation.
type RemoveAllResourcePermissionsRequest struct {
	*aws.Request
	Input *types.RemoveAllResourcePermissionsInput
	Copy  func(*types.RemoveAllResourcePermissionsInput) RemoveAllResourcePermissionsRequest
}

// Send marshals and sends the RemoveAllResourcePermissions API request.
func (r RemoveAllResourcePermissionsRequest) Send(ctx context.Context) (*RemoveAllResourcePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveAllResourcePermissionsResponse{
		RemoveAllResourcePermissionsOutput: r.Request.Data.(*types.RemoveAllResourcePermissionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveAllResourcePermissionsResponse is the response type for the
// RemoveAllResourcePermissions API operation.
type RemoveAllResourcePermissionsResponse struct {
	*types.RemoveAllResourcePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveAllResourcePermissions request.
func (r *RemoveAllResourcePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
