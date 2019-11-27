// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

const opPutApprovalResult = "PutApprovalResult"

// PutApprovalResultRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Provides the response to a manual approval request to AWS CodePipeline. Valid
// responses include Approved and Rejected.
//
//    // Example sending a request using PutApprovalResultRequest.
//    req := client.PutApprovalResultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutApprovalResult
func (c *Client) PutApprovalResultRequest(input *types.PutApprovalResultInput) PutApprovalResultRequest {
	op := &aws.Operation{
		Name:       opPutApprovalResult,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutApprovalResultInput{}
	}

	req := c.newRequest(op, input, &types.PutApprovalResultOutput{})
	return PutApprovalResultRequest{Request: req, Input: input, Copy: c.PutApprovalResultRequest}
}

// PutApprovalResultRequest is the request type for the
// PutApprovalResult API operation.
type PutApprovalResultRequest struct {
	*aws.Request
	Input *types.PutApprovalResultInput
	Copy  func(*types.PutApprovalResultInput) PutApprovalResultRequest
}

// Send marshals and sends the PutApprovalResult API request.
func (r PutApprovalResultRequest) Send(ctx context.Context) (*PutApprovalResultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutApprovalResultResponse{
		PutApprovalResultOutput: r.Request.Data.(*types.PutApprovalResultOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutApprovalResultResponse is the response type for the
// PutApprovalResult API operation.
type PutApprovalResultResponse struct {
	*types.PutApprovalResultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutApprovalResult request.
func (r *PutApprovalResultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
