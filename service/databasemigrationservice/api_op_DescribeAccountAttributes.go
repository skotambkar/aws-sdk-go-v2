// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribeAccountAttributes = "DescribeAccountAttributes"

// DescribeAccountAttributesRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Lists all of the AWS DMS attributes for a customer account. These attributes
// include AWS DMS quotas for the account and a unique account identifier in
// a particular DMS region. DMS quotas include a list of resource quotas supported
// by the account, such as the number of replication instances allowed. The
// description for each resource quota, includes the quota name, current usage
// toward that quota, and the quota's maximum value. DMS uses the unique account
// identifier to name each artifact used by DMS in the given region.
//
// This command does not take any parameters.
//
//    // Example sending a request using DescribeAccountAttributesRequest.
//    req := client.DescribeAccountAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeAccountAttributes
func (c *Client) DescribeAccountAttributesRequest(input *types.DescribeAccountAttributesInput) DescribeAccountAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAccountAttributesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAccountAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeAccountAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeAccountAttributesRequest{Request: req, Input: input, Copy: c.DescribeAccountAttributesRequest}
}

// DescribeAccountAttributesRequest is the request type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesRequest struct {
	*aws.Request
	Input *types.DescribeAccountAttributesInput
	Copy  func(*types.DescribeAccountAttributesInput) DescribeAccountAttributesRequest
}

// Send marshals and sends the DescribeAccountAttributes API request.
func (r DescribeAccountAttributesRequest) Send(ctx context.Context) (*DescribeAccountAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountAttributesResponse{
		DescribeAccountAttributesOutput: r.Request.Data.(*types.DescribeAccountAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountAttributesResponse is the response type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesResponse struct {
	*types.DescribeAccountAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountAttributes request.
func (r *DescribeAccountAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
