// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetRelationalDatabaseLogStreams = "GetRelationalDatabaseLogStreams"

// GetRelationalDatabaseLogStreamsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns a list of available log streams for a specific database in Amazon
// Lightsail.
//
//    // Example sending a request using GetRelationalDatabaseLogStreamsRequest.
//    req := client.GetRelationalDatabaseLogStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseLogStreams
func (c *Client) GetRelationalDatabaseLogStreamsRequest(input *types.GetRelationalDatabaseLogStreamsInput) GetRelationalDatabaseLogStreamsRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseLogStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRelationalDatabaseLogStreamsInput{}
	}

	req := c.newRequest(op, input, &types.GetRelationalDatabaseLogStreamsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRelationalDatabaseLogStreamsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRelationalDatabaseLogStreamsRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseLogStreamsRequest}
}

// GetRelationalDatabaseLogStreamsRequest is the request type for the
// GetRelationalDatabaseLogStreams API operation.
type GetRelationalDatabaseLogStreamsRequest struct {
	*aws.Request
	Input *types.GetRelationalDatabaseLogStreamsInput
	Copy  func(*types.GetRelationalDatabaseLogStreamsInput) GetRelationalDatabaseLogStreamsRequest
}

// Send marshals and sends the GetRelationalDatabaseLogStreams API request.
func (r GetRelationalDatabaseLogStreamsRequest) Send(ctx context.Context) (*GetRelationalDatabaseLogStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseLogStreamsResponse{
		GetRelationalDatabaseLogStreamsOutput: r.Request.Data.(*types.GetRelationalDatabaseLogStreamsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseLogStreamsResponse is the response type for the
// GetRelationalDatabaseLogStreams API operation.
type GetRelationalDatabaseLogStreamsResponse struct {
	*types.GetRelationalDatabaseLogStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseLogStreams request.
func (r *GetRelationalDatabaseLogStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
