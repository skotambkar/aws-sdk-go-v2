// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
)

const opCreateLocationEfs = "CreateLocationEfs"

// CreateLocationEfsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Creates an endpoint for an Amazon EFS file system.
//
//    // Example sending a request using CreateLocationEfsRequest.
//    req := client.CreateLocationEfsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateLocationEfs
func (c *Client) CreateLocationEfsRequest(input *types.CreateLocationEfsInput) CreateLocationEfsRequest {
	op := &aws.Operation{
		Name:       opCreateLocationEfs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateLocationEfsInput{}
	}

	req := c.newRequest(op, input, &types.CreateLocationEfsOutput{})
	return CreateLocationEfsRequest{Request: req, Input: input, Copy: c.CreateLocationEfsRequest}
}

// CreateLocationEfsRequest is the request type for the
// CreateLocationEfs API operation.
type CreateLocationEfsRequest struct {
	*aws.Request
	Input *types.CreateLocationEfsInput
	Copy  func(*types.CreateLocationEfsInput) CreateLocationEfsRequest
}

// Send marshals and sends the CreateLocationEfs API request.
func (r CreateLocationEfsRequest) Send(ctx context.Context) (*CreateLocationEfsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLocationEfsResponse{
		CreateLocationEfsOutput: r.Request.Data.(*types.CreateLocationEfsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLocationEfsResponse is the response type for the
// CreateLocationEfs API operation.
type CreateLocationEfsResponse struct {
	*types.CreateLocationEfsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLocationEfs request.
func (r *CreateLocationEfsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
