// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/support/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
)

const opDescribeTrustedAdvisorCheckResult = "DescribeTrustedAdvisorCheckResult"

// DescribeTrustedAdvisorCheckResultRequest returns a request value for making API operation for
// AWS Support.
//
// Returns the results of the Trusted Advisor check that has the specified check
// ID. Check IDs can be obtained by calling DescribeTrustedAdvisorChecks.
//
// The response contains a TrustedAdvisorCheckResult object, which contains
// these three objects:
//
//    * TrustedAdvisorCategorySpecificSummary
//
//    * TrustedAdvisorResourceDetail
//
//    * TrustedAdvisorResourcesSummary
//
// In addition, the response contains these fields:
//
//    * status. The alert status of the check: "ok" (green), "warning" (yellow),
//    "error" (red), or "not_available".
//
//    * timestamp. The time of the last refresh of the check.
//
//    * checkId. The unique identifier for the check.
//
//    // Example sending a request using DescribeTrustedAdvisorCheckResultRequest.
//    req := client.DescribeTrustedAdvisorCheckResultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorCheckResult
func (c *Client) DescribeTrustedAdvisorCheckResultRequest(input *types.DescribeTrustedAdvisorCheckResultInput) DescribeTrustedAdvisorCheckResultRequest {
	op := &aws.Operation{
		Name:       opDescribeTrustedAdvisorCheckResult,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeTrustedAdvisorCheckResultInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTrustedAdvisorCheckResultOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeTrustedAdvisorCheckResultMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeTrustedAdvisorCheckResultRequest{Request: req, Input: input, Copy: c.DescribeTrustedAdvisorCheckResultRequest}
}

// DescribeTrustedAdvisorCheckResultRequest is the request type for the
// DescribeTrustedAdvisorCheckResult API operation.
type DescribeTrustedAdvisorCheckResultRequest struct {
	*aws.Request
	Input *types.DescribeTrustedAdvisorCheckResultInput
	Copy  func(*types.DescribeTrustedAdvisorCheckResultInput) DescribeTrustedAdvisorCheckResultRequest
}

// Send marshals and sends the DescribeTrustedAdvisorCheckResult API request.
func (r DescribeTrustedAdvisorCheckResultRequest) Send(ctx context.Context) (*DescribeTrustedAdvisorCheckResultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrustedAdvisorCheckResultResponse{
		DescribeTrustedAdvisorCheckResultOutput: r.Request.Data.(*types.DescribeTrustedAdvisorCheckResultOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTrustedAdvisorCheckResultResponse is the response type for the
// DescribeTrustedAdvisorCheckResult API operation.
type DescribeTrustedAdvisorCheckResultResponse struct {
	*types.DescribeTrustedAdvisorCheckResultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrustedAdvisorCheckResult request.
func (r *DescribeTrustedAdvisorCheckResultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
