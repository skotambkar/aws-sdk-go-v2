// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
)

const opCreateApplicationVersion = "CreateApplicationVersion"

// CreateApplicationVersionRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Creates an application version.
//
//    // Example sending a request using CreateApplicationVersionRequest.
//    req := client.CreateApplicationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateApplicationVersion
func (c *Client) CreateApplicationVersionRequest(input *types.CreateApplicationVersionInput) CreateApplicationVersionRequest {
	op := &aws.Operation{
		Name:       opCreateApplicationVersion,
		HTTPMethod: "PUT",
		HTTPPath:   "/applications/{applicationId}/versions/{semanticVersion}",
	}

	if input == nil {
		input = &types.CreateApplicationVersionInput{}
	}

	req := c.newRequest(op, input, &types.CreateApplicationVersionOutput{})
	return CreateApplicationVersionRequest{Request: req, Input: input, Copy: c.CreateApplicationVersionRequest}
}

// CreateApplicationVersionRequest is the request type for the
// CreateApplicationVersion API operation.
type CreateApplicationVersionRequest struct {
	*aws.Request
	Input *types.CreateApplicationVersionInput
	Copy  func(*types.CreateApplicationVersionInput) CreateApplicationVersionRequest
}

// Send marshals and sends the CreateApplicationVersion API request.
func (r CreateApplicationVersionRequest) Send(ctx context.Context) (*CreateApplicationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationVersionResponse{
		CreateApplicationVersionOutput: r.Request.Data.(*types.CreateApplicationVersionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationVersionResponse is the response type for the
// CreateApplicationVersion API operation.
type CreateApplicationVersionResponse struct {
	*types.CreateApplicationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplicationVersion request.
func (r *CreateApplicationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
