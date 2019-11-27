// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
)

const opDeleteAssessmentRun = "DeleteAssessmentRun"

// DeleteAssessmentRunRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Deletes the assessment run that is specified by the ARN of the assessment
// run.
//
//    // Example sending a request using DeleteAssessmentRunRequest.
//    req := client.DeleteAssessmentRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DeleteAssessmentRun
func (c *Client) DeleteAssessmentRunRequest(input *types.DeleteAssessmentRunInput) DeleteAssessmentRunRequest {
	op := &aws.Operation{
		Name:       opDeleteAssessmentRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAssessmentRunInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAssessmentRunOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAssessmentRunRequest{Request: req, Input: input, Copy: c.DeleteAssessmentRunRequest}
}

// DeleteAssessmentRunRequest is the request type for the
// DeleteAssessmentRun API operation.
type DeleteAssessmentRunRequest struct {
	*aws.Request
	Input *types.DeleteAssessmentRunInput
	Copy  func(*types.DeleteAssessmentRunInput) DeleteAssessmentRunRequest
}

// Send marshals and sends the DeleteAssessmentRun API request.
func (r DeleteAssessmentRunRequest) Send(ctx context.Context) (*DeleteAssessmentRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssessmentRunResponse{
		DeleteAssessmentRunOutput: r.Request.Data.(*types.DeleteAssessmentRunOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssessmentRunResponse is the response type for the
// DeleteAssessmentRun API operation.
type DeleteAssessmentRunResponse struct {
	*types.DeleteAssessmentRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAssessmentRun request.
func (r *DeleteAssessmentRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
