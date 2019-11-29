// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opCreateJourney = "CreateJourney"

// CreateJourneyRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a journey for an application.
//
//    // Example sending a request using CreateJourneyRequest.
//    req := client.CreateJourneyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateJourney
func (c *Client) CreateJourneyRequest(input *types.CreateJourneyInput) CreateJourneyRequest {
	op := &aws.Operation{
		Name:       opCreateJourney,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/journeys",
	}

	if input == nil {
		input = &types.CreateJourneyInput{}
	}

	req := c.newRequest(op, input, &types.CreateJourneyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateJourneyMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateJourneyRequest{Request: req, Input: input, Copy: c.CreateJourneyRequest}
}

// CreateJourneyRequest is the request type for the
// CreateJourney API operation.
type CreateJourneyRequest struct {
	*aws.Request
	Input *types.CreateJourneyInput
	Copy  func(*types.CreateJourneyInput) CreateJourneyRequest
}

// Send marshals and sends the CreateJourney API request.
func (r CreateJourneyRequest) Send(ctx context.Context) (*CreateJourneyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJourneyResponse{
		CreateJourneyOutput: r.Request.Data.(*types.CreateJourneyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJourneyResponse is the response type for the
// CreateJourney API operation.
type CreateJourneyResponse struct {
	*types.CreateJourneyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJourney request.
func (r *CreateJourneyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
