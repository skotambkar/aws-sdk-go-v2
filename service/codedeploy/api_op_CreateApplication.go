// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opCreateApplication = "CreateApplication"

// CreateApplicationRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Creates an application.
//
//    // Example sending a request using CreateApplicationRequest.
//    req := client.CreateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateApplication
func (c *Client) CreateApplicationRequest(input *types.CreateApplicationInput) CreateApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateApplicationInput{}
	}

	req := c.newRequest(op, input, &types.CreateApplicationOutput{})
	return CreateApplicationRequest{Request: req, Input: input, Copy: c.CreateApplicationRequest}
}

// CreateApplicationRequest is the request type for the
// CreateApplication API operation.
type CreateApplicationRequest struct {
	*aws.Request
	Input *types.CreateApplicationInput
	Copy  func(*types.CreateApplicationInput) CreateApplicationRequest
}

// Send marshals and sends the CreateApplication API request.
func (r CreateApplicationRequest) Send(ctx context.Context) (*CreateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationResponse{
		CreateApplicationOutput: r.Request.Data.(*types.CreateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationResponse is the response type for the
// CreateApplication API operation.
type CreateApplicationResponse struct {
	*types.CreateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplication request.
func (r *CreateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
