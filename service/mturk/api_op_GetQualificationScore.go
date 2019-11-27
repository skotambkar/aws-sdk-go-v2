// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opGetQualificationScore = "GetQualificationScore"

// GetQualificationScoreRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The GetQualificationScore operation returns the value of a Worker's Qualification
// for a given Qualification type.
//
// To get a Worker's Qualification, you must know the Worker's ID. The Worker's
// ID is included in the assignment data returned by the ListAssignmentsForHIT
// operation.
//
// Only the owner of a Qualification type can query the value of a Worker's
// Qualification of that type.
//
//    // Example sending a request using GetQualificationScoreRequest.
//    req := client.GetQualificationScoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/GetQualificationScore
func (c *Client) GetQualificationScoreRequest(input *types.GetQualificationScoreInput) GetQualificationScoreRequest {
	op := &aws.Operation{
		Name:       opGetQualificationScore,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetQualificationScoreInput{}
	}

	req := c.newRequest(op, input, &types.GetQualificationScoreOutput{})
	return GetQualificationScoreRequest{Request: req, Input: input, Copy: c.GetQualificationScoreRequest}
}

// GetQualificationScoreRequest is the request type for the
// GetQualificationScore API operation.
type GetQualificationScoreRequest struct {
	*aws.Request
	Input *types.GetQualificationScoreInput
	Copy  func(*types.GetQualificationScoreInput) GetQualificationScoreRequest
}

// Send marshals and sends the GetQualificationScore API request.
func (r GetQualificationScoreRequest) Send(ctx context.Context) (*GetQualificationScoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetQualificationScoreResponse{
		GetQualificationScoreOutput: r.Request.Data.(*types.GetQualificationScoreOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetQualificationScoreResponse is the response type for the
// GetQualificationScore API operation.
type GetQualificationScoreResponse struct {
	*types.GetQualificationScoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetQualificationScore request.
func (r *GetQualificationScoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
