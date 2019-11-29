// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalizeruntime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/types"
)

const opGetRecommendations = "GetRecommendations"

// GetRecommendationsRequest returns a request value for making API operation for
// Amazon Personalize Runtime.
//
// Returns a list of recommended items. The required input depends on the recipe
// type used to create the solution backing the campaign, as follows:
//
//    * RELATED_ITEMS - itemId required, userId not used
//
//    * USER_PERSONALIZATION - itemId optional, userId required
//
// Campaigns that are backed by a solution created using a recipe of type PERSONALIZED_RANKING
// use the API.
//
//    // Example sending a request using GetRecommendationsRequest.
//    req := client.GetRecommendationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-runtime-2018-05-22/GetRecommendations
func (c *Client) GetRecommendationsRequest(input *types.GetRecommendationsInput) GetRecommendationsRequest {
	op := &aws.Operation{
		Name:       opGetRecommendations,
		HTTPMethod: "POST",
		HTTPPath:   "/recommendations",
	}

	if input == nil {
		input = &types.GetRecommendationsInput{}
	}

	req := c.newRequest(op, input, &types.GetRecommendationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetRecommendationsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRecommendationsRequest{Request: req, Input: input, Copy: c.GetRecommendationsRequest}
}

// GetRecommendationsRequest is the request type for the
// GetRecommendations API operation.
type GetRecommendationsRequest struct {
	*aws.Request
	Input *types.GetRecommendationsInput
	Copy  func(*types.GetRecommendationsInput) GetRecommendationsRequest
}

// Send marshals and sends the GetRecommendations API request.
func (r GetRecommendationsRequest) Send(ctx context.Context) (*GetRecommendationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRecommendationsResponse{
		GetRecommendationsOutput: r.Request.Data.(*types.GetRecommendationsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRecommendationsResponse is the response type for the
// GetRecommendations API operation.
type GetRecommendationsResponse struct {
	*types.GetRecommendationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRecommendations request.
func (r *GetRecommendationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
