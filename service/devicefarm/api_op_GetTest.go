// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opGetTest = "GetTest"

// GetTestRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about a test.
//
//    // Example sending a request using GetTestRequest.
//    req := client.GetTestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetTest
func (c *Client) GetTestRequest(input *types.GetTestInput) GetTestRequest {
	op := &aws.Operation{
		Name:       opGetTest,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetTestInput{}
	}

	req := c.newRequest(op, input, &types.GetTestOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetTestMarshaler{Input: input}.GetNamedBuildHandler())

	return GetTestRequest{Request: req, Input: input, Copy: c.GetTestRequest}
}

// GetTestRequest is the request type for the
// GetTest API operation.
type GetTestRequest struct {
	*aws.Request
	Input *types.GetTestInput
	Copy  func(*types.GetTestInput) GetTestRequest
}

// Send marshals and sends the GetTest API request.
func (r GetTestRequest) Send(ctx context.Context) (*GetTestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTestResponse{
		GetTestOutput: r.Request.Data.(*types.GetTestOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTestResponse is the response type for the
// GetTest API operation.
type GetTestResponse struct {
	*types.GetTestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTest request.
func (r *GetTestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
