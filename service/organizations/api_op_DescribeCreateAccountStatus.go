// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opDescribeCreateAccountStatus = "DescribeCreateAccountStatus"

// DescribeCreateAccountStatusRequest returns a request value for making API operation for
// AWS Organizations.
//
// Retrieves the current status of an asynchronous request to create an account.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DescribeCreateAccountStatusRequest.
//    req := client.DescribeCreateAccountStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DescribeCreateAccountStatus
func (c *Client) DescribeCreateAccountStatusRequest(input *types.DescribeCreateAccountStatusInput) DescribeCreateAccountStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeCreateAccountStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeCreateAccountStatusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeCreateAccountStatusOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeCreateAccountStatusMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeCreateAccountStatusRequest{Request: req, Input: input, Copy: c.DescribeCreateAccountStatusRequest}
}

// DescribeCreateAccountStatusRequest is the request type for the
// DescribeCreateAccountStatus API operation.
type DescribeCreateAccountStatusRequest struct {
	*aws.Request
	Input *types.DescribeCreateAccountStatusInput
	Copy  func(*types.DescribeCreateAccountStatusInput) DescribeCreateAccountStatusRequest
}

// Send marshals and sends the DescribeCreateAccountStatus API request.
func (r DescribeCreateAccountStatusRequest) Send(ctx context.Context) (*DescribeCreateAccountStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCreateAccountStatusResponse{
		DescribeCreateAccountStatusOutput: r.Request.Data.(*types.DescribeCreateAccountStatusOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCreateAccountStatusResponse is the response type for the
// DescribeCreateAccountStatus API operation.
type DescribeCreateAccountStatusResponse struct {
	*types.DescribeCreateAccountStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCreateAccountStatus request.
func (r *DescribeCreateAccountStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
