// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opRejectQualificationRequest = "RejectQualificationRequest"

// RejectQualificationRequestRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The RejectQualificationRequest operation rejects a user's request for a Qualification.
//
// You can provide a text message explaining why the request was rejected. The
// Worker who made the request can see this message.
//
//    // Example sending a request using RejectQualificationRequestRequest.
//    req := client.RejectQualificationRequestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/RejectQualificationRequest
func (c *Client) RejectQualificationRequestRequest(input *types.RejectQualificationRequestInput) RejectQualificationRequestRequest {
	op := &aws.Operation{
		Name:       opRejectQualificationRequest,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RejectQualificationRequestInput{}
	}

	req := c.newRequest(op, input, &types.RejectQualificationRequestOutput{})
	return RejectQualificationRequestRequest{Request: req, Input: input, Copy: c.RejectQualificationRequestRequest}
}

// RejectQualificationRequestRequest is the request type for the
// RejectQualificationRequest API operation.
type RejectQualificationRequestRequest struct {
	*aws.Request
	Input *types.RejectQualificationRequestInput
	Copy  func(*types.RejectQualificationRequestInput) RejectQualificationRequestRequest
}

// Send marshals and sends the RejectQualificationRequest API request.
func (r RejectQualificationRequestRequest) Send(ctx context.Context) (*RejectQualificationRequestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectQualificationRequestResponse{
		RejectQualificationRequestOutput: r.Request.Data.(*types.RejectQualificationRequestOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectQualificationRequestResponse is the response type for the
// RejectQualificationRequest API operation.
type RejectQualificationRequestResponse struct {
	*types.RejectQualificationRequestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectQualificationRequest request.
func (r *RejectQualificationRequestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
