// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opStartUserImportJob = "StartUserImportJob"

// StartUserImportJobRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Starts the user import.
//
//    // Example sending a request using StartUserImportJobRequest.
//    req := client.StartUserImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/StartUserImportJob
func (c *Client) StartUserImportJobRequest(input *types.StartUserImportJobInput) StartUserImportJobRequest {
	op := &aws.Operation{
		Name:       opStartUserImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartUserImportJobInput{}
	}

	req := c.newRequest(op, input, &types.StartUserImportJobOutput{})
	return StartUserImportJobRequest{Request: req, Input: input, Copy: c.StartUserImportJobRequest}
}

// StartUserImportJobRequest is the request type for the
// StartUserImportJob API operation.
type StartUserImportJobRequest struct {
	*aws.Request
	Input *types.StartUserImportJobInput
	Copy  func(*types.StartUserImportJobInput) StartUserImportJobRequest
}

// Send marshals and sends the StartUserImportJob API request.
func (r StartUserImportJobRequest) Send(ctx context.Context) (*StartUserImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartUserImportJobResponse{
		StartUserImportJobOutput: r.Request.Data.(*types.StartUserImportJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartUserImportJobResponse is the response type for the
// StartUserImportJob API operation.
type StartUserImportJobResponse struct {
	*types.StartUserImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartUserImportJob request.
func (r *StartUserImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
