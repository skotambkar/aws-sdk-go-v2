// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opImportKeyPair = "ImportKeyPair"

// ImportKeyPairRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Imports a public SSH key from a specific key pair.
//
//    // Example sending a request using ImportKeyPairRequest.
//    req := client.ImportKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/ImportKeyPair
func (c *Client) ImportKeyPairRequest(input *types.ImportKeyPairInput) ImportKeyPairRequest {
	op := &aws.Operation{
		Name:       opImportKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ImportKeyPairInput{}
	}

	req := c.newRequest(op, input, &types.ImportKeyPairOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ImportKeyPairMarshaler{Input: input}.GetNamedBuildHandler())

	return ImportKeyPairRequest{Request: req, Input: input, Copy: c.ImportKeyPairRequest}
}

// ImportKeyPairRequest is the request type for the
// ImportKeyPair API operation.
type ImportKeyPairRequest struct {
	*aws.Request
	Input *types.ImportKeyPairInput
	Copy  func(*types.ImportKeyPairInput) ImportKeyPairRequest
}

// Send marshals and sends the ImportKeyPair API request.
func (r ImportKeyPairRequest) Send(ctx context.Context) (*ImportKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportKeyPairResponse{
		ImportKeyPairOutput: r.Request.Data.(*types.ImportKeyPairOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportKeyPairResponse is the response type for the
// ImportKeyPair API operation.
type ImportKeyPairResponse struct {
	*types.ImportKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportKeyPair request.
func (r *ImportKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
