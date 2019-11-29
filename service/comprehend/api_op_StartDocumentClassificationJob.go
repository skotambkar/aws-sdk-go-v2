// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opStartDocumentClassificationJob = "StartDocumentClassificationJob"

// StartDocumentClassificationJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Starts an asynchronous document classification job. Use the operation to
// track the progress of the job.
//
//    // Example sending a request using StartDocumentClassificationJobRequest.
//    req := client.StartDocumentClassificationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StartDocumentClassificationJob
func (c *Client) StartDocumentClassificationJobRequest(input *types.StartDocumentClassificationJobInput) StartDocumentClassificationJobRequest {
	op := &aws.Operation{
		Name:       opStartDocumentClassificationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartDocumentClassificationJobInput{}
	}

	req := c.newRequest(op, input, &types.StartDocumentClassificationJobOutput{})
	return StartDocumentClassificationJobRequest{Request: req, Input: input, Copy: c.StartDocumentClassificationJobRequest}
}

// StartDocumentClassificationJobRequest is the request type for the
// StartDocumentClassificationJob API operation.
type StartDocumentClassificationJobRequest struct {
	*aws.Request
	Input *types.StartDocumentClassificationJobInput
	Copy  func(*types.StartDocumentClassificationJobInput) StartDocumentClassificationJobRequest
}

// Send marshals and sends the StartDocumentClassificationJob API request.
func (r StartDocumentClassificationJobRequest) Send(ctx context.Context) (*StartDocumentClassificationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDocumentClassificationJobResponse{
		StartDocumentClassificationJobOutput: r.Request.Data.(*types.StartDocumentClassificationJobOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDocumentClassificationJobResponse is the response type for the
// StartDocumentClassificationJob API operation.
type StartDocumentClassificationJobResponse struct {
	*types.StartDocumentClassificationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDocumentClassificationJob request.
func (r *StartDocumentClassificationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
