// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opUpdateJourney = "UpdateJourney"

// UpdateJourneyRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the configuration and other settings for a journey.
//
//    // Example sending a request using UpdateJourneyRequest.
//    req := client.UpdateJourneyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateJourney
func (c *Client) UpdateJourneyRequest(input *types.UpdateJourneyInput) UpdateJourneyRequest {
	op := &aws.Operation{
		Name:       opUpdateJourney,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/journeys/{journey-id}",
	}

	if input == nil {
		input = &types.UpdateJourneyInput{}
	}

	req := c.newRequest(op, input, &types.UpdateJourneyOutput{})
	return UpdateJourneyRequest{Request: req, Input: input, Copy: c.UpdateJourneyRequest}
}

// UpdateJourneyRequest is the request type for the
// UpdateJourney API operation.
type UpdateJourneyRequest struct {
	*aws.Request
	Input *types.UpdateJourneyInput
	Copy  func(*types.UpdateJourneyInput) UpdateJourneyRequest
}

// Send marshals and sends the UpdateJourney API request.
func (r UpdateJourneyRequest) Send(ctx context.Context) (*UpdateJourneyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJourneyResponse{
		UpdateJourneyOutput: r.Request.Data.(*types.UpdateJourneyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJourneyResponse is the response type for the
// UpdateJourney API operation.
type UpdateJourneyResponse struct {
	*types.UpdateJourneyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJourney request.
func (r *UpdateJourneyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
