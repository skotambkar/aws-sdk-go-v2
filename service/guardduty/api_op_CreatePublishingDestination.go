// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
)

const opCreatePublishingDestination = "CreatePublishingDestination"

// CreatePublishingDestinationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Creates a publishing destination to send findings to. The resource to send
// findings to must exist before you use this operation.
//
//    // Example sending a request using CreatePublishingDestinationRequest.
//    req := client.CreatePublishingDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreatePublishingDestination
func (c *Client) CreatePublishingDestinationRequest(input *types.CreatePublishingDestinationInput) CreatePublishingDestinationRequest {
	op := &aws.Operation{
		Name:       opCreatePublishingDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/publishingDestination",
	}

	if input == nil {
		input = &types.CreatePublishingDestinationInput{}
	}

	req := c.newRequest(op, input, &types.CreatePublishingDestinationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreatePublishingDestinationMarshaler{Input: input}.GetNamedBuildHandler())

	return CreatePublishingDestinationRequest{Request: req, Input: input, Copy: c.CreatePublishingDestinationRequest}
}

// CreatePublishingDestinationRequest is the request type for the
// CreatePublishingDestination API operation.
type CreatePublishingDestinationRequest struct {
	*aws.Request
	Input *types.CreatePublishingDestinationInput
	Copy  func(*types.CreatePublishingDestinationInput) CreatePublishingDestinationRequest
}

// Send marshals and sends the CreatePublishingDestination API request.
func (r CreatePublishingDestinationRequest) Send(ctx context.Context) (*CreatePublishingDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePublishingDestinationResponse{
		CreatePublishingDestinationOutput: r.Request.Data.(*types.CreatePublishingDestinationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePublishingDestinationResponse is the response type for the
// CreatePublishingDestination API operation.
type CreatePublishingDestinationResponse struct {
	*types.CreatePublishingDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePublishingDestination request.
func (r *CreatePublishingDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
