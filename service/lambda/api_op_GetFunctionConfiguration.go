// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opGetFunctionConfiguration = "GetFunctionConfiguration"

// GetFunctionConfigurationRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns the version-specific settings of a Lambda function or version. The
// output includes only options that can vary between versions of a function.
// To modify these settings, use UpdateFunctionConfiguration.
//
// To get all of a function's details, including function-level settings, use
// GetFunction.
//
//    // Example sending a request using GetFunctionConfigurationRequest.
//    req := client.GetFunctionConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunctionConfiguration
func (c *Client) GetFunctionConfigurationRequest(input *types.GetFunctionConfigurationInput) GetFunctionConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetFunctionConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/configuration",
	}

	if input == nil {
		input = &types.GetFunctionConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.GetFunctionConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetFunctionConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return GetFunctionConfigurationRequest{Request: req, Input: input, Copy: c.GetFunctionConfigurationRequest}
}

// GetFunctionConfigurationRequest is the request type for the
// GetFunctionConfiguration API operation.
type GetFunctionConfigurationRequest struct {
	*aws.Request
	Input *types.GetFunctionConfigurationInput
	Copy  func(*types.GetFunctionConfigurationInput) GetFunctionConfigurationRequest
}

// Send marshals and sends the GetFunctionConfiguration API request.
func (r GetFunctionConfigurationRequest) Send(ctx context.Context) (*GetFunctionConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFunctionConfigurationResponse{
		GetFunctionConfigurationOutput: r.Request.Data.(*types.GetFunctionConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFunctionConfigurationResponse is the response type for the
// GetFunctionConfiguration API operation.
type GetFunctionConfigurationResponse struct {
	*types.GetFunctionConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFunctionConfiguration request.
func (r *GetFunctionConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
