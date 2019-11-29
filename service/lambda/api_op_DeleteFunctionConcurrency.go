// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opDeleteFunctionConcurrency = "DeleteFunctionConcurrency"

// DeleteFunctionConcurrencyRequest returns a request value for making API operation for
// AWS Lambda.
//
// Removes a concurrent execution limit from a function.
//
//    // Example sending a request using DeleteFunctionConcurrencyRequest.
//    req := client.DeleteFunctionConcurrencyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteFunctionConcurrency
func (c *Client) DeleteFunctionConcurrencyRequest(input *types.DeleteFunctionConcurrencyInput) DeleteFunctionConcurrencyRequest {
	op := &aws.Operation{
		Name:       opDeleteFunctionConcurrency,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2017-10-31/functions/{FunctionName}/concurrency",
	}

	if input == nil {
		input = &types.DeleteFunctionConcurrencyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteFunctionConcurrencyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteFunctionConcurrencyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteFunctionConcurrencyRequest{Request: req, Input: input, Copy: c.DeleteFunctionConcurrencyRequest}
}

// DeleteFunctionConcurrencyRequest is the request type for the
// DeleteFunctionConcurrency API operation.
type DeleteFunctionConcurrencyRequest struct {
	*aws.Request
	Input *types.DeleteFunctionConcurrencyInput
	Copy  func(*types.DeleteFunctionConcurrencyInput) DeleteFunctionConcurrencyRequest
}

// Send marshals and sends the DeleteFunctionConcurrency API request.
func (r DeleteFunctionConcurrencyRequest) Send(ctx context.Context) (*DeleteFunctionConcurrencyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFunctionConcurrencyResponse{
		DeleteFunctionConcurrencyOutput: r.Request.Data.(*types.DeleteFunctionConcurrencyOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFunctionConcurrencyResponse is the response type for the
// DeleteFunctionConcurrency API operation.
type DeleteFunctionConcurrencyResponse struct {
	*types.DeleteFunctionConcurrencyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFunctionConcurrency request.
func (r *DeleteFunctionConcurrencyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
