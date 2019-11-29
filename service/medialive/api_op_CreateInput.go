// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/medialive/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
)

const opCreateInput = "CreateInput"

// CreateInputRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Create an input
//
//    // Example sending a request using CreateInputRequest.
//    req := client.CreateInputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateInput
func (c *Client) CreateInputRequest(input *types.CreateInputInput) CreateInputRequest {
	op := &aws.Operation{
		Name:       opCreateInput,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/inputs",
	}

	if input == nil {
		input = &types.CreateInputInput{}
	}

	req := c.newRequest(op, input, &types.CreateInputOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateInputMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateInputRequest{Request: req, Input: input, Copy: c.CreateInputRequest}
}

// CreateInputRequest is the request type for the
// CreateInput API operation.
type CreateInputRequest struct {
	*aws.Request
	Input *types.CreateInputInput
	Copy  func(*types.CreateInputInput) CreateInputRequest
}

// Send marshals and sends the CreateInput API request.
func (r CreateInputRequest) Send(ctx context.Context) (*CreateInputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInputResponse{
		CreateInputOutput: r.Request.Data.(*types.CreateInputOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInputResponse is the response type for the
// CreateInput API operation.
type CreateInputResponse struct {
	*types.CreateInputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInput request.
func (r *CreateInputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
