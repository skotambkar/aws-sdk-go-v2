// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opMergeBranchesBySquash = "MergeBranchesBySquash"

// MergeBranchesBySquashRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Merges two branches using the squash merge strategy.
//
//    // Example sending a request using MergeBranchesBySquashRequest.
//    req := client.MergeBranchesBySquashRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/MergeBranchesBySquash
func (c *Client) MergeBranchesBySquashRequest(input *types.MergeBranchesBySquashInput) MergeBranchesBySquashRequest {
	op := &aws.Operation{
		Name:       opMergeBranchesBySquash,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.MergeBranchesBySquashInput{}
	}

	req := c.newRequest(op, input, &types.MergeBranchesBySquashOutput{})
	return MergeBranchesBySquashRequest{Request: req, Input: input, Copy: c.MergeBranchesBySquashRequest}
}

// MergeBranchesBySquashRequest is the request type for the
// MergeBranchesBySquash API operation.
type MergeBranchesBySquashRequest struct {
	*aws.Request
	Input *types.MergeBranchesBySquashInput
	Copy  func(*types.MergeBranchesBySquashInput) MergeBranchesBySquashRequest
}

// Send marshals and sends the MergeBranchesBySquash API request.
func (r MergeBranchesBySquashRequest) Send(ctx context.Context) (*MergeBranchesBySquashResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MergeBranchesBySquashResponse{
		MergeBranchesBySquashOutput: r.Request.Data.(*types.MergeBranchesBySquashOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MergeBranchesBySquashResponse is the response type for the
// MergeBranchesBySquash API operation.
type MergeBranchesBySquashResponse struct {
	*types.MergeBranchesBySquashOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MergeBranchesBySquash request.
func (r *MergeBranchesBySquashResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
