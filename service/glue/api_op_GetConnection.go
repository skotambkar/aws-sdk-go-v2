// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opGetConnection = "GetConnection"

// GetConnectionRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a connection definition from the Data Catalog.
//
//    // Example sending a request using GetConnectionRequest.
//    req := client.GetConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetConnection
func (c *Client) GetConnectionRequest(input *types.GetConnectionInput) GetConnectionRequest {
	op := &aws.Operation{
		Name:       opGetConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetConnectionInput{}
	}

	req := c.newRequest(op, input, &types.GetConnectionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetConnectionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetConnectionRequest{Request: req, Input: input, Copy: c.GetConnectionRequest}
}

// GetConnectionRequest is the request type for the
// GetConnection API operation.
type GetConnectionRequest struct {
	*aws.Request
	Input *types.GetConnectionInput
	Copy  func(*types.GetConnectionInput) GetConnectionRequest
}

// Send marshals and sends the GetConnection API request.
func (r GetConnectionRequest) Send(ctx context.Context) (*GetConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionResponse{
		GetConnectionOutput: r.Request.Data.(*types.GetConnectionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectionResponse is the response type for the
// GetConnection API operation.
type GetConnectionResponse struct {
	*types.GetConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnection request.
func (r *GetConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
