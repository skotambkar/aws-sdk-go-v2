// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opCreateMitigationAction = "CreateMitigationAction"

// CreateMitigationActionRequest returns a request value for making API operation for
// AWS IoT.
//
// Defines an action that can be applied to audit findings by using StartAuditMitigationActionsTask.
// Each mitigation action can apply only one type of change.
//
//    // Example sending a request using CreateMitigationActionRequest.
//    req := client.CreateMitigationActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateMitigationActionRequest(input *types.CreateMitigationActionInput) CreateMitigationActionRequest {
	op := &aws.Operation{
		Name:       opCreateMitigationAction,
		HTTPMethod: "POST",
		HTTPPath:   "/mitigationactions/actions/{actionName}",
	}

	if input == nil {
		input = &types.CreateMitigationActionInput{}
	}

	req := c.newRequest(op, input, &types.CreateMitigationActionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateMitigationActionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateMitigationActionRequest{Request: req, Input: input, Copy: c.CreateMitigationActionRequest}
}

// CreateMitigationActionRequest is the request type for the
// CreateMitigationAction API operation.
type CreateMitigationActionRequest struct {
	*aws.Request
	Input *types.CreateMitigationActionInput
	Copy  func(*types.CreateMitigationActionInput) CreateMitigationActionRequest
}

// Send marshals and sends the CreateMitigationAction API request.
func (r CreateMitigationActionRequest) Send(ctx context.Context) (*CreateMitigationActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMitigationActionResponse{
		CreateMitigationActionOutput: r.Request.Data.(*types.CreateMitigationActionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMitigationActionResponse is the response type for the
// CreateMitigationAction API operation.
type CreateMitigationActionResponse struct {
	*types.CreateMitigationActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMitigationAction request.
func (r *CreateMitigationActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
