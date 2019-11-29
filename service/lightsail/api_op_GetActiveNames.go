// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetActiveNames = "GetActiveNames"

// GetActiveNamesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the names of all active (not deleted) resources.
//
//    // Example sending a request using GetActiveNamesRequest.
//    req := client.GetActiveNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetActiveNames
func (c *Client) GetActiveNamesRequest(input *types.GetActiveNamesInput) GetActiveNamesRequest {
	op := &aws.Operation{
		Name:       opGetActiveNames,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetActiveNamesInput{}
	}

	req := c.newRequest(op, input, &types.GetActiveNamesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetActiveNamesMarshaler{Input: input}.GetNamedBuildHandler())

	return GetActiveNamesRequest{Request: req, Input: input, Copy: c.GetActiveNamesRequest}
}

// GetActiveNamesRequest is the request type for the
// GetActiveNames API operation.
type GetActiveNamesRequest struct {
	*aws.Request
	Input *types.GetActiveNamesInput
	Copy  func(*types.GetActiveNamesInput) GetActiveNamesRequest
}

// Send marshals and sends the GetActiveNames API request.
func (r GetActiveNamesRequest) Send(ctx context.Context) (*GetActiveNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetActiveNamesResponse{
		GetActiveNamesOutput: r.Request.Data.(*types.GetActiveNamesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetActiveNamesResponse is the response type for the
// GetActiveNames API operation.
type GetActiveNamesResponse struct {
	*types.GetActiveNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetActiveNames request.
func (r *GetActiveNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
