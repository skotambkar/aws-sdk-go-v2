// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opGetMergeOptions = "GetMergeOptions"

// GetMergeOptionsRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about the merge options available for merging two specified
// branches. For details about why a merge option is not available, use GetMergeConflicts
// or DescribeMergeConflicts.
//
//    // Example sending a request using GetMergeOptionsRequest.
//    req := client.GetMergeOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetMergeOptions
func (c *Client) GetMergeOptionsRequest(input *types.GetMergeOptionsInput) GetMergeOptionsRequest {
	op := &aws.Operation{
		Name:       opGetMergeOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetMergeOptionsInput{}
	}

	req := c.newRequest(op, input, &types.GetMergeOptionsOutput{})
	return GetMergeOptionsRequest{Request: req, Input: input, Copy: c.GetMergeOptionsRequest}
}

// GetMergeOptionsRequest is the request type for the
// GetMergeOptions API operation.
type GetMergeOptionsRequest struct {
	*aws.Request
	Input *types.GetMergeOptionsInput
	Copy  func(*types.GetMergeOptionsInput) GetMergeOptionsRequest
}

// Send marshals and sends the GetMergeOptions API request.
func (r GetMergeOptionsRequest) Send(ctx context.Context) (*GetMergeOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMergeOptionsResponse{
		GetMergeOptionsOutput: r.Request.Data.(*types.GetMergeOptionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMergeOptionsResponse is the response type for the
// GetMergeOptions API operation.
type GetMergeOptionsResponse struct {
	*types.GetMergeOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMergeOptions request.
func (r *GetMergeOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
