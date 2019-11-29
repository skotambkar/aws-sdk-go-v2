// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
)

const opUpdatePublishingDestination = "UpdatePublishingDestination"

// UpdatePublishingDestinationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Updates information about the publishing destination specified by the destinationId.
//
//    // Example sending a request using UpdatePublishingDestinationRequest.
//    req := client.UpdatePublishingDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UpdatePublishingDestination
func (c *Client) UpdatePublishingDestinationRequest(input *types.UpdatePublishingDestinationInput) UpdatePublishingDestinationRequest {
	op := &aws.Operation{
		Name:       opUpdatePublishingDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/publishingDestination/{destinationId}",
	}

	if input == nil {
		input = &types.UpdatePublishingDestinationInput{}
	}

	req := c.newRequest(op, input, &types.UpdatePublishingDestinationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdatePublishingDestinationMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdatePublishingDestinationRequest{Request: req, Input: input, Copy: c.UpdatePublishingDestinationRequest}
}

// UpdatePublishingDestinationRequest is the request type for the
// UpdatePublishingDestination API operation.
type UpdatePublishingDestinationRequest struct {
	*aws.Request
	Input *types.UpdatePublishingDestinationInput
	Copy  func(*types.UpdatePublishingDestinationInput) UpdatePublishingDestinationRequest
}

// Send marshals and sends the UpdatePublishingDestination API request.
func (r UpdatePublishingDestinationRequest) Send(ctx context.Context) (*UpdatePublishingDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePublishingDestinationResponse{
		UpdatePublishingDestinationOutput: r.Request.Data.(*types.UpdatePublishingDestinationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePublishingDestinationResponse is the response type for the
// UpdatePublishingDestination API operation.
type UpdatePublishingDestinationResponse struct {
	*types.UpdatePublishingDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePublishingDestination request.
func (r *UpdatePublishingDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
