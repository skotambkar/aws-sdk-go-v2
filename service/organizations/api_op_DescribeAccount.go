// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opDescribeAccount = "DescribeAccount"

// DescribeAccountRequest returns a request value for making API operation for
// AWS Organizations.
//
// Retrieves AWS Organizations-related information about the specified account.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DescribeAccountRequest.
//    req := client.DescribeAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DescribeAccount
func (c *Client) DescribeAccountRequest(input *types.DescribeAccountInput) DescribeAccountRequest {
	op := &aws.Operation{
		Name:       opDescribeAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAccountInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAccountOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeAccountMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeAccountRequest{Request: req, Input: input, Copy: c.DescribeAccountRequest}
}

// DescribeAccountRequest is the request type for the
// DescribeAccount API operation.
type DescribeAccountRequest struct {
	*aws.Request
	Input *types.DescribeAccountInput
	Copy  func(*types.DescribeAccountInput) DescribeAccountRequest
}

// Send marshals and sends the DescribeAccount API request.
func (r DescribeAccountRequest) Send(ctx context.Context) (*DescribeAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountResponse{
		DescribeAccountOutput: r.Request.Data.(*types.DescribeAccountOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountResponse is the response type for the
// DescribeAccount API operation.
type DescribeAccountResponse struct {
	*types.DescribeAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccount request.
func (r *DescribeAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
