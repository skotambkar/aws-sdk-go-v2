// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opCreateFolder = "CreateFolder"

// CreateFolderRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Creates a folder with the specified name and parent folder.
//
//    // Example sending a request using CreateFolderRequest.
//    req := client.CreateFolderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateFolder
func (c *Client) CreateFolderRequest(input *types.CreateFolderInput) CreateFolderRequest {
	op := &aws.Operation{
		Name:       opCreateFolder,
		HTTPMethod: "POST",
		HTTPPath:   "/api/v1/folders",
	}

	if input == nil {
		input = &types.CreateFolderInput{}
	}

	req := c.newRequest(op, input, &types.CreateFolderOutput{})
	return CreateFolderRequest{Request: req, Input: input, Copy: c.CreateFolderRequest}
}

// CreateFolderRequest is the request type for the
// CreateFolder API operation.
type CreateFolderRequest struct {
	*aws.Request
	Input *types.CreateFolderInput
	Copy  func(*types.CreateFolderInput) CreateFolderRequest
}

// Send marshals and sends the CreateFolder API request.
func (r CreateFolderRequest) Send(ctx context.Context) (*CreateFolderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFolderResponse{
		CreateFolderOutput: r.Request.Data.(*types.CreateFolderOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFolderResponse is the response type for the
// CreateFolder API operation.
type CreateFolderResponse struct {
	*types.CreateFolderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFolder request.
func (r *CreateFolderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
