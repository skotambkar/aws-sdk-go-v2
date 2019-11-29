// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

const opGetCompatibleElasticsearchVersions = "GetCompatibleElasticsearchVersions"

// GetCompatibleElasticsearchVersionsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Returns a list of upgrade compatible Elastisearch versions. You can optionally
// pass a DomainName to get all upgrade compatible Elasticsearch versions for
// that specific domain.
//
//    // Example sending a request using GetCompatibleElasticsearchVersionsRequest.
//    req := client.GetCompatibleElasticsearchVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetCompatibleElasticsearchVersionsRequest(input *types.GetCompatibleElasticsearchVersionsInput) GetCompatibleElasticsearchVersionsRequest {
	op := &aws.Operation{
		Name:       opGetCompatibleElasticsearchVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/compatibleVersions",
	}

	if input == nil {
		input = &types.GetCompatibleElasticsearchVersionsInput{}
	}

	req := c.newRequest(op, input, &types.GetCompatibleElasticsearchVersionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetCompatibleElasticsearchVersionsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCompatibleElasticsearchVersionsRequest{Request: req, Input: input, Copy: c.GetCompatibleElasticsearchVersionsRequest}
}

// GetCompatibleElasticsearchVersionsRequest is the request type for the
// GetCompatibleElasticsearchVersions API operation.
type GetCompatibleElasticsearchVersionsRequest struct {
	*aws.Request
	Input *types.GetCompatibleElasticsearchVersionsInput
	Copy  func(*types.GetCompatibleElasticsearchVersionsInput) GetCompatibleElasticsearchVersionsRequest
}

// Send marshals and sends the GetCompatibleElasticsearchVersions API request.
func (r GetCompatibleElasticsearchVersionsRequest) Send(ctx context.Context) (*GetCompatibleElasticsearchVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCompatibleElasticsearchVersionsResponse{
		GetCompatibleElasticsearchVersionsOutput: r.Request.Data.(*types.GetCompatibleElasticsearchVersionsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCompatibleElasticsearchVersionsResponse is the response type for the
// GetCompatibleElasticsearchVersions API operation.
type GetCompatibleElasticsearchVersionsResponse struct {
	*types.GetCompatibleElasticsearchVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCompatibleElasticsearchVersions request.
func (r *GetCompatibleElasticsearchVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
