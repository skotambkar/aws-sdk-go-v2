// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

const opDescribeIdentityUsage = "DescribeIdentityUsage"

// DescribeIdentityUsageRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Gets usage information for an identity, including number of datasets and
// data usage.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
//
//    // Example sending a request using DescribeIdentityUsageRequest.
//    req := client.DescribeIdentityUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeIdentityUsage
func (c *Client) DescribeIdentityUsageRequest(input *types.DescribeIdentityUsageInput) DescribeIdentityUsageRequest {
	op := &aws.Operation{
		Name:       opDescribeIdentityUsage,
		HTTPMethod: "GET",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}",
	}

	if input == nil {
		input = &types.DescribeIdentityUsageInput{}
	}

	req := c.newRequest(op, input, &types.DescribeIdentityUsageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeIdentityUsageMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeIdentityUsageRequest{Request: req, Input: input, Copy: c.DescribeIdentityUsageRequest}
}

// DescribeIdentityUsageRequest is the request type for the
// DescribeIdentityUsage API operation.
type DescribeIdentityUsageRequest struct {
	*aws.Request
	Input *types.DescribeIdentityUsageInput
	Copy  func(*types.DescribeIdentityUsageInput) DescribeIdentityUsageRequest
}

// Send marshals and sends the DescribeIdentityUsage API request.
func (r DescribeIdentityUsageRequest) Send(ctx context.Context) (*DescribeIdentityUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIdentityUsageResponse{
		DescribeIdentityUsageOutput: r.Request.Data.(*types.DescribeIdentityUsageOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIdentityUsageResponse is the response type for the
// DescribeIdentityUsage API operation.
type DescribeIdentityUsageResponse struct {
	*types.DescribeIdentityUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIdentityUsage request.
func (r *DescribeIdentityUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
