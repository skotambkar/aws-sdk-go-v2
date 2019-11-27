// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opCreateLoggerDefinition = "CreateLoggerDefinition"

// CreateLoggerDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a logger definition. You may provide the initial version of the logger
// definition now or use ''CreateLoggerDefinitionVersion'' at a later time.
//
//    // Example sending a request using CreateLoggerDefinitionRequest.
//    req := client.CreateLoggerDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateLoggerDefinition
func (c *Client) CreateLoggerDefinitionRequest(input *types.CreateLoggerDefinitionInput) CreateLoggerDefinitionRequest {
	op := &aws.Operation{
		Name:       opCreateLoggerDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/loggers",
	}

	if input == nil {
		input = &types.CreateLoggerDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.CreateLoggerDefinitionOutput{})
	return CreateLoggerDefinitionRequest{Request: req, Input: input, Copy: c.CreateLoggerDefinitionRequest}
}

// CreateLoggerDefinitionRequest is the request type for the
// CreateLoggerDefinition API operation.
type CreateLoggerDefinitionRequest struct {
	*aws.Request
	Input *types.CreateLoggerDefinitionInput
	Copy  func(*types.CreateLoggerDefinitionInput) CreateLoggerDefinitionRequest
}

// Send marshals and sends the CreateLoggerDefinition API request.
func (r CreateLoggerDefinitionRequest) Send(ctx context.Context) (*CreateLoggerDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoggerDefinitionResponse{
		CreateLoggerDefinitionOutput: r.Request.Data.(*types.CreateLoggerDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoggerDefinitionResponse is the response type for the
// CreateLoggerDefinition API operation.
type CreateLoggerDefinitionResponse struct {
	*types.CreateLoggerDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoggerDefinition request.
func (r *CreateLoggerDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
