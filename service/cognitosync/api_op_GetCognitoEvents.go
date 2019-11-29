// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

const opGetCognitoEvents = "GetCognitoEvents"

// GetCognitoEventsRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Gets the events and the corresponding Lambda functions associated with an
// identity pool.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
//    // Example sending a request using GetCognitoEventsRequest.
//    req := client.GetCognitoEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/GetCognitoEvents
func (c *Client) GetCognitoEventsRequest(input *types.GetCognitoEventsInput) GetCognitoEventsRequest {
	op := &aws.Operation{
		Name:       opGetCognitoEvents,
		HTTPMethod: "GET",
		HTTPPath:   "/identitypools/{IdentityPoolId}/events",
	}

	if input == nil {
		input = &types.GetCognitoEventsInput{}
	}

	req := c.newRequest(op, input, &types.GetCognitoEventsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetCognitoEventsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCognitoEventsRequest{Request: req, Input: input, Copy: c.GetCognitoEventsRequest}
}

// GetCognitoEventsRequest is the request type for the
// GetCognitoEvents API operation.
type GetCognitoEventsRequest struct {
	*aws.Request
	Input *types.GetCognitoEventsInput
	Copy  func(*types.GetCognitoEventsInput) GetCognitoEventsRequest
}

// Send marshals and sends the GetCognitoEvents API request.
func (r GetCognitoEventsRequest) Send(ctx context.Context) (*GetCognitoEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCognitoEventsResponse{
		GetCognitoEventsOutput: r.Request.Data.(*types.GetCognitoEventsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCognitoEventsResponse is the response type for the
// GetCognitoEvents API operation.
type GetCognitoEventsResponse struct {
	*types.GetCognitoEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCognitoEvents request.
func (r *GetCognitoEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
