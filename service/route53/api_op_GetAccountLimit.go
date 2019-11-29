// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opGetAccountLimit = "GetAccountLimit"

// GetAccountLimitRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the specified limit for the current account, for example, the maximum
// number of health checks that you can create using the account.
//
// For the default limit, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a
// case (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53).
//
// You can also view account limits in AWS Trusted Advisor. Sign in to the AWS
// Management Console and open the Trusted Advisor console at https://console.aws.amazon.com/trustedadvisor/
// (https://console.aws.amazon.com/trustedadvisor). Then choose Service limits
// in the navigation pane.
//
//    // Example sending a request using GetAccountLimitRequest.
//    req := client.GetAccountLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetAccountLimit
func (c *Client) GetAccountLimitRequest(input *types.GetAccountLimitInput) GetAccountLimitRequest {
	op := &aws.Operation{
		Name:       opGetAccountLimit,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/accountlimit/{Type}",
	}

	if input == nil {
		input = &types.GetAccountLimitInput{}
	}

	req := c.newRequest(op, input, &types.GetAccountLimitOutput{})
	return GetAccountLimitRequest{Request: req, Input: input, Copy: c.GetAccountLimitRequest}
}

// GetAccountLimitRequest is the request type for the
// GetAccountLimit API operation.
type GetAccountLimitRequest struct {
	*aws.Request
	Input *types.GetAccountLimitInput
	Copy  func(*types.GetAccountLimitInput) GetAccountLimitRequest
}

// Send marshals and sends the GetAccountLimit API request.
func (r GetAccountLimitRequest) Send(ctx context.Context) (*GetAccountLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountLimitResponse{
		GetAccountLimitOutput: r.Request.Data.(*types.GetAccountLimitOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountLimitResponse is the response type for the
// GetAccountLimit API operation.
type GetAccountLimitResponse struct {
	*types.GetAccountLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountLimit request.
func (r *GetAccountLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
