// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDescribeRemediationExecutionStatus = "DescribeRemediationExecutionStatus"

// DescribeRemediationExecutionStatusRequest returns a request value for making API operation for
// AWS Config.
//
// Provides a detailed view of a Remediation Execution for a set of resources
// including state, timestamps for when steps for the remediation execution
// occur, and any error messages for steps that have failed. When you specify
// the limit and the next token, you receive a paginated response.
//
//    // Example sending a request using DescribeRemediationExecutionStatusRequest.
//    req := client.DescribeRemediationExecutionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeRemediationExecutionStatus
func (c *Client) DescribeRemediationExecutionStatusRequest(input *types.DescribeRemediationExecutionStatusInput) DescribeRemediationExecutionStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeRemediationExecutionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeRemediationExecutionStatusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeRemediationExecutionStatusOutput{})
	return DescribeRemediationExecutionStatusRequest{Request: req, Input: input, Copy: c.DescribeRemediationExecutionStatusRequest}
}

// DescribeRemediationExecutionStatusRequest is the request type for the
// DescribeRemediationExecutionStatus API operation.
type DescribeRemediationExecutionStatusRequest struct {
	*aws.Request
	Input *types.DescribeRemediationExecutionStatusInput
	Copy  func(*types.DescribeRemediationExecutionStatusInput) DescribeRemediationExecutionStatusRequest
}

// Send marshals and sends the DescribeRemediationExecutionStatus API request.
func (r DescribeRemediationExecutionStatusRequest) Send(ctx context.Context) (*DescribeRemediationExecutionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRemediationExecutionStatusResponse{
		DescribeRemediationExecutionStatusOutput: r.Request.Data.(*types.DescribeRemediationExecutionStatusOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeRemediationExecutionStatusRequestPaginator returns a paginator for DescribeRemediationExecutionStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeRemediationExecutionStatusRequest(input)
//   p := configservice.NewDescribeRemediationExecutionStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeRemediationExecutionStatusPaginator(req DescribeRemediationExecutionStatusRequest) DescribeRemediationExecutionStatusPaginator {
	return DescribeRemediationExecutionStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeRemediationExecutionStatusInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeRemediationExecutionStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeRemediationExecutionStatusPaginator struct {
	aws.Pager
}

func (p *DescribeRemediationExecutionStatusPaginator) CurrentPage() *types.DescribeRemediationExecutionStatusOutput {
	return p.Pager.CurrentPage().(*types.DescribeRemediationExecutionStatusOutput)
}

// DescribeRemediationExecutionStatusResponse is the response type for the
// DescribeRemediationExecutionStatus API operation.
type DescribeRemediationExecutionStatusResponse struct {
	*types.DescribeRemediationExecutionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRemediationExecutionStatus request.
func (r *DescribeRemediationExecutionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
