// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

const opGetNamedQuery = "GetNamedQuery"

// GetNamedQueryRequest returns a request value for making API operation for
// Amazon Athena.
//
// Returns information about a single query. Requires that you have access to
// the workgroup in which the query was saved.
//
//    // Example sending a request using GetNamedQueryRequest.
//    req := client.GetNamedQueryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/GetNamedQuery
func (c *Client) GetNamedQueryRequest(input *types.GetNamedQueryInput) GetNamedQueryRequest {
	op := &aws.Operation{
		Name:       opGetNamedQuery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetNamedQueryInput{}
	}

	req := c.newRequest(op, input, &types.GetNamedQueryOutput{})
	return GetNamedQueryRequest{Request: req, Input: input, Copy: c.GetNamedQueryRequest}
}

// GetNamedQueryRequest is the request type for the
// GetNamedQuery API operation.
type GetNamedQueryRequest struct {
	*aws.Request
	Input *types.GetNamedQueryInput
	Copy  func(*types.GetNamedQueryInput) GetNamedQueryRequest
}

// Send marshals and sends the GetNamedQuery API request.
func (r GetNamedQueryRequest) Send(ctx context.Context) (*GetNamedQueryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNamedQueryResponse{
		GetNamedQueryOutput: r.Request.Data.(*types.GetNamedQueryOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNamedQueryResponse is the response type for the
// GetNamedQuery API operation.
type GetNamedQueryResponse struct {
	*types.GetNamedQueryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNamedQuery request.
func (r *GetNamedQueryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
