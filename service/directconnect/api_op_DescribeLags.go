// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opDescribeLags = "DescribeLags"

// DescribeLagsRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Describes all your link aggregation groups (LAG) or the specified LAG.
//
//    // Example sending a request using DescribeLagsRequest.
//    req := client.DescribeLagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeLags
func (c *Client) DescribeLagsRequest(input *types.DescribeLagsInput) DescribeLagsRequest {
	op := &aws.Operation{
		Name:       opDescribeLags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeLagsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLagsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeLagsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeLagsRequest{Request: req, Input: input, Copy: c.DescribeLagsRequest}
}

// DescribeLagsRequest is the request type for the
// DescribeLags API operation.
type DescribeLagsRequest struct {
	*aws.Request
	Input *types.DescribeLagsInput
	Copy  func(*types.DescribeLagsInput) DescribeLagsRequest
}

// Send marshals and sends the DescribeLags API request.
func (r DescribeLagsRequest) Send(ctx context.Context) (*DescribeLagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLagsResponse{
		DescribeLagsOutput: r.Request.Data.(*types.DescribeLagsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLagsResponse is the response type for the
// DescribeLags API operation.
type DescribeLagsResponse struct {
	*types.DescribeLagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLags request.
func (r *DescribeLagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
