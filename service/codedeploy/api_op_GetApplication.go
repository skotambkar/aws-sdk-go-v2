// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opGetApplication = "GetApplication"

// GetApplicationRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about an application.
//
//    // Example sending a request using GetApplicationRequest.
//    req := client.GetApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetApplication
func (c *Client) GetApplicationRequest(input *types.GetApplicationInput) GetApplicationRequest {
	op := &aws.Operation{
		Name:       opGetApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetApplicationInput{}
	}

	req := c.newRequest(op, input, &types.GetApplicationOutput{})
	return GetApplicationRequest{Request: req, Input: input, Copy: c.GetApplicationRequest}
}

// GetApplicationRequest is the request type for the
// GetApplication API operation.
type GetApplicationRequest struct {
	*aws.Request
	Input *types.GetApplicationInput
	Copy  func(*types.GetApplicationInput) GetApplicationRequest
}

// Send marshals and sends the GetApplication API request.
func (r GetApplicationRequest) Send(ctx context.Context) (*GetApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationResponse{
		GetApplicationOutput: r.Request.Data.(*types.GetApplicationOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationResponse is the response type for the
// GetApplication API operation.
type GetApplicationResponse struct {
	*types.GetApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplication request.
func (r *GetApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
