// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opGetLoggingOptions = "GetLoggingOptions"

// GetLoggingOptionsRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets the logging options.
//
// NOTE: use of this command is not recommended. Use GetV2LoggingOptions instead.
//
//    // Example sending a request using GetLoggingOptionsRequest.
//    req := client.GetLoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetLoggingOptionsRequest(input *types.GetLoggingOptionsInput) GetLoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opGetLoggingOptions,
		HTTPMethod: "GET",
		HTTPPath:   "/loggingOptions",
	}

	if input == nil {
		input = &types.GetLoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &types.GetLoggingOptionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetLoggingOptionsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLoggingOptionsRequest{Request: req, Input: input, Copy: c.GetLoggingOptionsRequest}
}

// GetLoggingOptionsRequest is the request type for the
// GetLoggingOptions API operation.
type GetLoggingOptionsRequest struct {
	*aws.Request
	Input *types.GetLoggingOptionsInput
	Copy  func(*types.GetLoggingOptionsInput) GetLoggingOptionsRequest
}

// Send marshals and sends the GetLoggingOptions API request.
func (r GetLoggingOptionsRequest) Send(ctx context.Context) (*GetLoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoggingOptionsResponse{
		GetLoggingOptionsOutput: r.Request.Data.(*types.GetLoggingOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoggingOptionsResponse is the response type for the
// GetLoggingOptions API operation.
type GetLoggingOptionsResponse struct {
	*types.GetLoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoggingOptions request.
func (r *GetLoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
