// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

const opStartQueryExecution = "StartQueryExecution"

// StartQueryExecutionRequest returns a request value for making API operation for
// Amazon Athena.
//
// Runs the SQL query statements contained in the Query. Requires you to have
// access to the workgroup in which the query ran.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using StartQueryExecutionRequest.
//    req := client.StartQueryExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/StartQueryExecution
func (c *Client) StartQueryExecutionRequest(input *types.StartQueryExecutionInput) StartQueryExecutionRequest {
	op := &aws.Operation{
		Name:       opStartQueryExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartQueryExecutionInput{}
	}

	req := c.newRequest(op, input, &types.StartQueryExecutionOutput{})
	return StartQueryExecutionRequest{Request: req, Input: input, Copy: c.StartQueryExecutionRequest}
}

// StartQueryExecutionRequest is the request type for the
// StartQueryExecution API operation.
type StartQueryExecutionRequest struct {
	*aws.Request
	Input *types.StartQueryExecutionInput
	Copy  func(*types.StartQueryExecutionInput) StartQueryExecutionRequest
}

// Send marshals and sends the StartQueryExecution API request.
func (r StartQueryExecutionRequest) Send(ctx context.Context) (*StartQueryExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartQueryExecutionResponse{
		StartQueryExecutionOutput: r.Request.Data.(*types.StartQueryExecutionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartQueryExecutionResponse is the response type for the
// StartQueryExecution API operation.
type StartQueryExecutionResponse struct {
	*types.StartQueryExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartQueryExecution request.
func (r *StartQueryExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
