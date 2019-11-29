// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
)

const opListPlaybackConfigurations = "ListPlaybackConfigurations"

// ListPlaybackConfigurationsRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Returns a list of the playback configurations defined in AWS Elemental MediaTailor.
// You can specify a maximum number of configurations to return at a time. The
// default maximum is 50. Results are returned in pagefuls. If MediaTailor has
// more configurations than the specified maximum, it provides parameters in
// the response that you can use to retrieve the next pageful.
//
//    // Example sending a request using ListPlaybackConfigurationsRequest.
//    req := client.ListPlaybackConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurations
func (c *Client) ListPlaybackConfigurationsRequest(input *types.ListPlaybackConfigurationsInput) ListPlaybackConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListPlaybackConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfigurations",
	}

	if input == nil {
		input = &types.ListPlaybackConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.ListPlaybackConfigurationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListPlaybackConfigurationsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPlaybackConfigurationsRequest{Request: req, Input: input, Copy: c.ListPlaybackConfigurationsRequest}
}

// ListPlaybackConfigurationsRequest is the request type for the
// ListPlaybackConfigurations API operation.
type ListPlaybackConfigurationsRequest struct {
	*aws.Request
	Input *types.ListPlaybackConfigurationsInput
	Copy  func(*types.ListPlaybackConfigurationsInput) ListPlaybackConfigurationsRequest
}

// Send marshals and sends the ListPlaybackConfigurations API request.
func (r ListPlaybackConfigurationsRequest) Send(ctx context.Context) (*ListPlaybackConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPlaybackConfigurationsResponse{
		ListPlaybackConfigurationsOutput: r.Request.Data.(*types.ListPlaybackConfigurationsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPlaybackConfigurationsResponse is the response type for the
// ListPlaybackConfigurations API operation.
type ListPlaybackConfigurationsResponse struct {
	*types.ListPlaybackConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPlaybackConfigurations request.
func (r *ListPlaybackConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
