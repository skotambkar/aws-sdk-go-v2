// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opPutEvaluations = "PutEvaluations"

// PutEvaluationsRequest returns a request value for making API operation for
// AWS Config.
//
// Used by an AWS Lambda function to deliver evaluation results to AWS Config.
// This action is required in every AWS Lambda function that is invoked by an
// AWS Config rule.
//
//    // Example sending a request using PutEvaluationsRequest.
//    req := client.PutEvaluationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutEvaluations
func (c *Client) PutEvaluationsRequest(input *types.PutEvaluationsInput) PutEvaluationsRequest {
	op := &aws.Operation{
		Name:       opPutEvaluations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutEvaluationsInput{}
	}

	req := c.newRequest(op, input, &types.PutEvaluationsOutput{})
	return PutEvaluationsRequest{Request: req, Input: input, Copy: c.PutEvaluationsRequest}
}

// PutEvaluationsRequest is the request type for the
// PutEvaluations API operation.
type PutEvaluationsRequest struct {
	*aws.Request
	Input *types.PutEvaluationsInput
	Copy  func(*types.PutEvaluationsInput) PutEvaluationsRequest
}

// Send marshals and sends the PutEvaluations API request.
func (r PutEvaluationsRequest) Send(ctx context.Context) (*PutEvaluationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEvaluationsResponse{
		PutEvaluationsOutput: r.Request.Data.(*types.PutEvaluationsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEvaluationsResponse is the response type for the
// PutEvaluations API operation.
type PutEvaluationsResponse struct {
	*types.PutEvaluationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEvaluations request.
func (r *PutEvaluationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
