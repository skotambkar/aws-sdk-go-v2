// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opGetLoggerDefinition = "GetLoggerDefinition"

// GetLoggerDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a logger definition.
//
//    // Example sending a request using GetLoggerDefinitionRequest.
//    req := client.GetLoggerDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetLoggerDefinition
func (c *Client) GetLoggerDefinitionRequest(input *types.GetLoggerDefinitionInput) GetLoggerDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetLoggerDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/loggers/{LoggerDefinitionId}",
	}

	if input == nil {
		input = &types.GetLoggerDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.GetLoggerDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetLoggerDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLoggerDefinitionRequest{Request: req, Input: input, Copy: c.GetLoggerDefinitionRequest}
}

// GetLoggerDefinitionRequest is the request type for the
// GetLoggerDefinition API operation.
type GetLoggerDefinitionRequest struct {
	*aws.Request
	Input *types.GetLoggerDefinitionInput
	Copy  func(*types.GetLoggerDefinitionInput) GetLoggerDefinitionRequest
}

// Send marshals and sends the GetLoggerDefinition API request.
func (r GetLoggerDefinitionRequest) Send(ctx context.Context) (*GetLoggerDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoggerDefinitionResponse{
		GetLoggerDefinitionOutput: r.Request.Data.(*types.GetLoggerDefinitionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoggerDefinitionResponse is the response type for the
// GetLoggerDefinition API operation.
type GetLoggerDefinitionResponse struct {
	*types.GetLoggerDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoggerDefinition request.
func (r *GetLoggerDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
