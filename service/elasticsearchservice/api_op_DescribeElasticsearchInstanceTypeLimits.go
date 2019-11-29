// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

const opDescribeElasticsearchInstanceTypeLimits = "DescribeElasticsearchInstanceTypeLimits"

// DescribeElasticsearchInstanceTypeLimitsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Describe Elasticsearch Limits for a given InstanceType and ElasticsearchVersion.
// When modifying existing Domain, specify the DomainName to know what Limits
// are supported for modifying.
//
//    // Example sending a request using DescribeElasticsearchInstanceTypeLimitsRequest.
//    req := client.DescribeElasticsearchInstanceTypeLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeElasticsearchInstanceTypeLimitsRequest(input *types.DescribeElasticsearchInstanceTypeLimitsInput) DescribeElasticsearchInstanceTypeLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeElasticsearchInstanceTypeLimits,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/instanceTypeLimits/{ElasticsearchVersion}/{InstanceType}",
	}

	if input == nil {
		input = &types.DescribeElasticsearchInstanceTypeLimitsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeElasticsearchInstanceTypeLimitsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeElasticsearchInstanceTypeLimitsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeElasticsearchInstanceTypeLimitsRequest{Request: req, Input: input, Copy: c.DescribeElasticsearchInstanceTypeLimitsRequest}
}

// DescribeElasticsearchInstanceTypeLimitsRequest is the request type for the
// DescribeElasticsearchInstanceTypeLimits API operation.
type DescribeElasticsearchInstanceTypeLimitsRequest struct {
	*aws.Request
	Input *types.DescribeElasticsearchInstanceTypeLimitsInput
	Copy  func(*types.DescribeElasticsearchInstanceTypeLimitsInput) DescribeElasticsearchInstanceTypeLimitsRequest
}

// Send marshals and sends the DescribeElasticsearchInstanceTypeLimits API request.
func (r DescribeElasticsearchInstanceTypeLimitsRequest) Send(ctx context.Context) (*DescribeElasticsearchInstanceTypeLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeElasticsearchInstanceTypeLimitsResponse{
		DescribeElasticsearchInstanceTypeLimitsOutput: r.Request.Data.(*types.DescribeElasticsearchInstanceTypeLimitsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeElasticsearchInstanceTypeLimitsResponse is the response type for the
// DescribeElasticsearchInstanceTypeLimits API operation.
type DescribeElasticsearchInstanceTypeLimitsResponse struct {
	*types.DescribeElasticsearchInstanceTypeLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeElasticsearchInstanceTypeLimits request.
func (r *DescribeElasticsearchInstanceTypeLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
