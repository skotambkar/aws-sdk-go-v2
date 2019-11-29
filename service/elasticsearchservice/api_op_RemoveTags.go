// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

const opRemoveTags = "RemoveTags"

// RemoveTagsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Removes the specified set of tags from the specified Elasticsearch domain.
//
//    // Example sending a request using RemoveTagsRequest.
//    req := client.RemoveTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RemoveTagsRequest(input *types.RemoveTagsInput) RemoveTagsRequest {
	op := &aws.Operation{
		Name:       opRemoveTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/tags-removal",
	}

	if input == nil {
		input = &types.RemoveTagsInput{}
	}

	req := c.newRequest(op, input, &types.RemoveTagsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.RemoveTagsMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveTagsRequest{Request: req, Input: input, Copy: c.RemoveTagsRequest}
}

// RemoveTagsRequest is the request type for the
// RemoveTags API operation.
type RemoveTagsRequest struct {
	*aws.Request
	Input *types.RemoveTagsInput
	Copy  func(*types.RemoveTagsInput) RemoveTagsRequest
}

// Send marshals and sends the RemoveTags API request.
func (r RemoveTagsRequest) Send(ctx context.Context) (*RemoveTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsResponse{
		RemoveTagsOutput: r.Request.Data.(*types.RemoveTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsResponse is the response type for the
// RemoveTags API operation.
type RemoveTagsResponse struct {
	*types.RemoveTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTags request.
func (r *RemoveTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
