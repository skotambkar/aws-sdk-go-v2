// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opDescribeSolutionVersion = "DescribeSolutionVersion"

// DescribeSolutionVersionRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes a specific version of a solution. For more information on solutions,
// see CreateSolution.
//
//    // Example sending a request using DescribeSolutionVersionRequest.
//    req := client.DescribeSolutionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSolutionVersion
func (c *Client) DescribeSolutionVersionRequest(input *types.DescribeSolutionVersionInput) DescribeSolutionVersionRequest {
	op := &aws.Operation{
		Name:       opDescribeSolutionVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeSolutionVersionInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSolutionVersionOutput{})
	return DescribeSolutionVersionRequest{Request: req, Input: input, Copy: c.DescribeSolutionVersionRequest}
}

// DescribeSolutionVersionRequest is the request type for the
// DescribeSolutionVersion API operation.
type DescribeSolutionVersionRequest struct {
	*aws.Request
	Input *types.DescribeSolutionVersionInput
	Copy  func(*types.DescribeSolutionVersionInput) DescribeSolutionVersionRequest
}

// Send marshals and sends the DescribeSolutionVersion API request.
func (r DescribeSolutionVersionRequest) Send(ctx context.Context) (*DescribeSolutionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSolutionVersionResponse{
		DescribeSolutionVersionOutput: r.Request.Data.(*types.DescribeSolutionVersionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSolutionVersionResponse is the response type for the
// DescribeSolutionVersion API operation.
type DescribeSolutionVersionResponse struct {
	*types.DescribeSolutionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSolutionVersion request.
func (r *DescribeSolutionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
