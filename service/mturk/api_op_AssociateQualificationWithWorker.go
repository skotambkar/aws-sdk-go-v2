// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opAssociateQualificationWithWorker = "AssociateQualificationWithWorker"

// AssociateQualificationWithWorkerRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The AssociateQualificationWithWorker operation gives a Worker a Qualification.
// AssociateQualificationWithWorker does not require that the Worker submit
// a Qualification request. It gives the Qualification directly to the Worker.
//
// You can only assign a Qualification of a Qualification type that you created
// (using the CreateQualificationType operation).
//
// Note: AssociateQualificationWithWorker does not affect any pending Qualification
// requests for the Qualification by the Worker. If you assign a Qualification
// to a Worker, then later grant a Qualification request made by the Worker,
// the granting of the request may modify the Qualification score. To resolve
// a pending Qualification request without affecting the Qualification the Worker
// already has, reject the request with the RejectQualificationRequest operation.
//
//    // Example sending a request using AssociateQualificationWithWorkerRequest.
//    req := client.AssociateQualificationWithWorkerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/AssociateQualificationWithWorker
func (c *Client) AssociateQualificationWithWorkerRequest(input *types.AssociateQualificationWithWorkerInput) AssociateQualificationWithWorkerRequest {
	op := &aws.Operation{
		Name:       opAssociateQualificationWithWorker,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateQualificationWithWorkerInput{}
	}

	req := c.newRequest(op, input, &types.AssociateQualificationWithWorkerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateQualificationWithWorkerMarshaler{Input: input}.GetNamedBuildHandler())

	return AssociateQualificationWithWorkerRequest{Request: req, Input: input, Copy: c.AssociateQualificationWithWorkerRequest}
}

// AssociateQualificationWithWorkerRequest is the request type for the
// AssociateQualificationWithWorker API operation.
type AssociateQualificationWithWorkerRequest struct {
	*aws.Request
	Input *types.AssociateQualificationWithWorkerInput
	Copy  func(*types.AssociateQualificationWithWorkerInput) AssociateQualificationWithWorkerRequest
}

// Send marshals and sends the AssociateQualificationWithWorker API request.
func (r AssociateQualificationWithWorkerRequest) Send(ctx context.Context) (*AssociateQualificationWithWorkerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateQualificationWithWorkerResponse{
		AssociateQualificationWithWorkerOutput: r.Request.Data.(*types.AssociateQualificationWithWorkerOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateQualificationWithWorkerResponse is the response type for the
// AssociateQualificationWithWorker API operation.
type AssociateQualificationWithWorkerResponse struct {
	*types.AssociateQualificationWithWorkerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateQualificationWithWorker request.
func (r *AssociateQualificationWithWorkerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
