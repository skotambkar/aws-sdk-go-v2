// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
)

const opCreatePipeline = "CreatePipeline"

// CreatePipelineRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Creates a new, empty pipeline. Use PutPipelineDefinition to populate the
// pipeline.
//
//    // Example sending a request using CreatePipelineRequest.
//    req := client.CreatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/CreatePipeline
func (c *Client) CreatePipelineRequest(input *types.CreatePipelineInput) CreatePipelineRequest {
	op := &aws.Operation{
		Name:       opCreatePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePipelineInput{}
	}

	req := c.newRequest(op, input, &types.CreatePipelineOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreatePipelineMarshaler{Input: input}.GetNamedBuildHandler())

	return CreatePipelineRequest{Request: req, Input: input, Copy: c.CreatePipelineRequest}
}

// CreatePipelineRequest is the request type for the
// CreatePipeline API operation.
type CreatePipelineRequest struct {
	*aws.Request
	Input *types.CreatePipelineInput
	Copy  func(*types.CreatePipelineInput) CreatePipelineRequest
}

// Send marshals and sends the CreatePipeline API request.
func (r CreatePipelineRequest) Send(ctx context.Context) (*CreatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePipelineResponse{
		CreatePipelineOutput: r.Request.Data.(*types.CreatePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePipelineResponse is the response type for the
// CreatePipeline API operation.
type CreatePipelineResponse struct {
	*types.CreatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePipeline request.
func (r *CreatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
