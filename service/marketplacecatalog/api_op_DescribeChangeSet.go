// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
)

const opDescribeChangeSet = "DescribeChangeSet"

// DescribeChangeSetRequest returns a request value for making API operation for
// AWS Marketplace Catalog Service.
//
// Provides information about a given change set.
//
//    // Example sending a request using DescribeChangeSetRequest.
//    req := client.DescribeChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplace-catalog-2018-09-17/DescribeChangeSet
func (c *Client) DescribeChangeSetRequest(input *types.DescribeChangeSetInput) DescribeChangeSetRequest {
	op := &aws.Operation{
		Name:       opDescribeChangeSet,
		HTTPMethod: "GET",
		HTTPPath:   "/DescribeChangeSet",
	}

	if input == nil {
		input = &types.DescribeChangeSetInput{}
	}

	req := c.newRequest(op, input, &types.DescribeChangeSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeChangeSetMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeChangeSetRequest{Request: req, Input: input, Copy: c.DescribeChangeSetRequest}
}

// DescribeChangeSetRequest is the request type for the
// DescribeChangeSet API operation.
type DescribeChangeSetRequest struct {
	*aws.Request
	Input *types.DescribeChangeSetInput
	Copy  func(*types.DescribeChangeSetInput) DescribeChangeSetRequest
}

// Send marshals and sends the DescribeChangeSet API request.
func (r DescribeChangeSetRequest) Send(ctx context.Context) (*DescribeChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChangeSetResponse{
		DescribeChangeSetOutput: r.Request.Data.(*types.DescribeChangeSetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChangeSetResponse is the response type for the
// DescribeChangeSet API operation.
type DescribeChangeSetResponse struct {
	*types.DescribeChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChangeSet request.
func (r *DescribeChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
