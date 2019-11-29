// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opUpdateFunctionCode = "UpdateFunctionCode"

// UpdateFunctionCodeRequest returns a request value for making API operation for
// AWS Lambda.
//
// Updates a Lambda function's code.
//
// The function's code is locked when you publish a version. You can't modify
// the code of a published version, only the unpublished version.
//
//    // Example sending a request using UpdateFunctionCodeRequest.
//    req := client.UpdateFunctionCodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UpdateFunctionCode
func (c *Client) UpdateFunctionCodeRequest(input *types.UpdateFunctionCodeInput) UpdateFunctionCodeRequest {
	op := &aws.Operation{
		Name:       opUpdateFunctionCode,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/code",
	}

	if input == nil {
		input = &types.UpdateFunctionCodeInput{}
	}

	req := c.newRequest(op, input, &types.UpdateFunctionCodeOutput{})
	return UpdateFunctionCodeRequest{Request: req, Input: input, Copy: c.UpdateFunctionCodeRequest}
}

// UpdateFunctionCodeRequest is the request type for the
// UpdateFunctionCode API operation.
type UpdateFunctionCodeRequest struct {
	*aws.Request
	Input *types.UpdateFunctionCodeInput
	Copy  func(*types.UpdateFunctionCodeInput) UpdateFunctionCodeRequest
}

// Send marshals and sends the UpdateFunctionCode API request.
func (r UpdateFunctionCodeRequest) Send(ctx context.Context) (*UpdateFunctionCodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFunctionCodeResponse{
		UpdateFunctionCodeOutput: r.Request.Data.(*types.UpdateFunctionCodeOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFunctionCodeResponse is the response type for the
// UpdateFunctionCode API operation.
type UpdateFunctionCodeResponse struct {
	*types.UpdateFunctionCodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFunctionCode request.
func (r *UpdateFunctionCodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
