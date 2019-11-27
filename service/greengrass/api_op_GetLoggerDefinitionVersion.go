// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opGetLoggerDefinitionVersion = "GetLoggerDefinitionVersion"

// GetLoggerDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a logger definition version.
//
//    // Example sending a request using GetLoggerDefinitionVersionRequest.
//    req := client.GetLoggerDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetLoggerDefinitionVersion
func (c *Client) GetLoggerDefinitionVersionRequest(input *types.GetLoggerDefinitionVersionInput) GetLoggerDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opGetLoggerDefinitionVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/loggers/{LoggerDefinitionId}/versions/{LoggerDefinitionVersionId}",
	}

	if input == nil {
		input = &types.GetLoggerDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &types.GetLoggerDefinitionVersionOutput{})
	return GetLoggerDefinitionVersionRequest{Request: req, Input: input, Copy: c.GetLoggerDefinitionVersionRequest}
}

// GetLoggerDefinitionVersionRequest is the request type for the
// GetLoggerDefinitionVersion API operation.
type GetLoggerDefinitionVersionRequest struct {
	*aws.Request
	Input *types.GetLoggerDefinitionVersionInput
	Copy  func(*types.GetLoggerDefinitionVersionInput) GetLoggerDefinitionVersionRequest
}

// Send marshals and sends the GetLoggerDefinitionVersion API request.
func (r GetLoggerDefinitionVersionRequest) Send(ctx context.Context) (*GetLoggerDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoggerDefinitionVersionResponse{
		GetLoggerDefinitionVersionOutput: r.Request.Data.(*types.GetLoggerDefinitionVersionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoggerDefinitionVersionResponse is the response type for the
// GetLoggerDefinitionVersion API operation.
type GetLoggerDefinitionVersionResponse struct {
	*types.GetLoggerDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoggerDefinitionVersion request.
func (r *GetLoggerDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
