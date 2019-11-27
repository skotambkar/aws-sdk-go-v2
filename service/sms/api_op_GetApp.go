// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opGetApp = "GetApp"

// GetAppRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Retrieve information about an application.
//
//    // Example sending a request using GetAppRequest.
//    req := client.GetAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/GetApp
func (c *Client) GetAppRequest(input *types.GetAppInput) GetAppRequest {
	op := &aws.Operation{
		Name:       opGetApp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetAppInput{}
	}

	req := c.newRequest(op, input, &types.GetAppOutput{})
	return GetAppRequest{Request: req, Input: input, Copy: c.GetAppRequest}
}

// GetAppRequest is the request type for the
// GetApp API operation.
type GetAppRequest struct {
	*aws.Request
	Input *types.GetAppInput
	Copy  func(*types.GetAppInput) GetAppRequest
}

// Send marshals and sends the GetApp API request.
func (r GetAppRequest) Send(ctx context.Context) (*GetAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAppResponse{
		GetAppOutput: r.Request.Data.(*types.GetAppOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAppResponse is the response type for the
// GetApp API operation.
type GetAppResponse struct {
	*types.GetAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApp request.
func (r *GetAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
