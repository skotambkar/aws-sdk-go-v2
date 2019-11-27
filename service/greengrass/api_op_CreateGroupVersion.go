// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opCreateGroupVersion = "CreateGroupVersion"

// CreateGroupVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a version of a group which has already been defined.
//
//    // Example sending a request using CreateGroupVersionRequest.
//    req := client.CreateGroupVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateGroupVersion
func (c *Client) CreateGroupVersionRequest(input *types.CreateGroupVersionInput) CreateGroupVersionRequest {
	op := &aws.Operation{
		Name:       opCreateGroupVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/groups/{GroupId}/versions",
	}

	if input == nil {
		input = &types.CreateGroupVersionInput{}
	}

	req := c.newRequest(op, input, &types.CreateGroupVersionOutput{})
	return CreateGroupVersionRequest{Request: req, Input: input, Copy: c.CreateGroupVersionRequest}
}

// CreateGroupVersionRequest is the request type for the
// CreateGroupVersion API operation.
type CreateGroupVersionRequest struct {
	*aws.Request
	Input *types.CreateGroupVersionInput
	Copy  func(*types.CreateGroupVersionInput) CreateGroupVersionRequest
}

// Send marshals and sends the CreateGroupVersion API request.
func (r CreateGroupVersionRequest) Send(ctx context.Context) (*CreateGroupVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupVersionResponse{
		CreateGroupVersionOutput: r.Request.Data.(*types.CreateGroupVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupVersionResponse is the response type for the
// CreateGroupVersion API operation.
type CreateGroupVersionResponse struct {
	*types.CreateGroupVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroupVersion request.
func (r *CreateGroupVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
