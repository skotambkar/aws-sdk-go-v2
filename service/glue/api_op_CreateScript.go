// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opCreateScript = "CreateScript"

// CreateScriptRequest returns a request value for making API operation for
// AWS Glue.
//
// Transforms a directed acyclic graph (DAG) into code.
//
//    // Example sending a request using CreateScriptRequest.
//    req := client.CreateScriptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateScript
func (c *Client) CreateScriptRequest(input *types.CreateScriptInput) CreateScriptRequest {
	op := &aws.Operation{
		Name:       opCreateScript,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateScriptInput{}
	}

	req := c.newRequest(op, input, &types.CreateScriptOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateScriptMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateScriptRequest{Request: req, Input: input, Copy: c.CreateScriptRequest}
}

// CreateScriptRequest is the request type for the
// CreateScript API operation.
type CreateScriptRequest struct {
	*aws.Request
	Input *types.CreateScriptInput
	Copy  func(*types.CreateScriptInput) CreateScriptRequest
}

// Send marshals and sends the CreateScript API request.
func (r CreateScriptRequest) Send(ctx context.Context) (*CreateScriptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateScriptResponse{
		CreateScriptOutput: r.Request.Data.(*types.CreateScriptOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateScriptResponse is the response type for the
// CreateScript API operation.
type CreateScriptResponse struct {
	*types.CreateScriptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateScript request.
func (r *CreateScriptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
