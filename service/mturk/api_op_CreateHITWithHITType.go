// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opCreateHITWithHITType = "CreateHITWithHITType"

// CreateHITWithHITTypeRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The CreateHITWithHITType operation creates a new Human Intelligence Task
// (HIT) using an existing HITTypeID generated by the CreateHITType operation.
//
// This is an alternative way to create HITs from the CreateHIT operation. This
// is the recommended best practice for Requesters who are creating large numbers
// of HITs.
//
// CreateHITWithHITType also supports several ways to provide question data:
// by providing a value for the Question parameter that fully specifies the
// contents of the HIT, or by providing a HitLayoutId and associated HitLayoutParameters.
//
// If a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see Amazon Mechanical Turk Pricing (https://requester.mturk.com/pricing).
//
//    // Example sending a request using CreateHITWithHITTypeRequest.
//    req := client.CreateHITWithHITTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/CreateHITWithHITType
func (c *Client) CreateHITWithHITTypeRequest(input *types.CreateHITWithHITTypeInput) CreateHITWithHITTypeRequest {
	op := &aws.Operation{
		Name:       opCreateHITWithHITType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateHITWithHITTypeInput{}
	}

	req := c.newRequest(op, input, &types.CreateHITWithHITTypeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateHITWithHITTypeMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateHITWithHITTypeRequest{Request: req, Input: input, Copy: c.CreateHITWithHITTypeRequest}
}

// CreateHITWithHITTypeRequest is the request type for the
// CreateHITWithHITType API operation.
type CreateHITWithHITTypeRequest struct {
	*aws.Request
	Input *types.CreateHITWithHITTypeInput
	Copy  func(*types.CreateHITWithHITTypeInput) CreateHITWithHITTypeRequest
}

// Send marshals and sends the CreateHITWithHITType API request.
func (r CreateHITWithHITTypeRequest) Send(ctx context.Context) (*CreateHITWithHITTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHITWithHITTypeResponse{
		CreateHITWithHITTypeOutput: r.Request.Data.(*types.CreateHITWithHITTypeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHITWithHITTypeResponse is the response type for the
// CreateHITWithHITType API operation.
type CreateHITWithHITTypeResponse struct {
	*types.CreateHITWithHITTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHITWithHITType request.
func (r *CreateHITWithHITTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
