// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opGetSampledRequests = "GetSampledRequests"

// GetSampledRequestsRequest returns a request value for making API operation for
// AWS WAF.
//
// Gets detailed information about a specified number of requests--a sample--that
// AWS WAF randomly selects from among the first 5,000 requests that your AWS
// resource received during a time range that you choose. You can specify a
// sample size of up to 500 requests, and you can specify any time range in
// the previous three hours.
//
// GetSampledRequests returns a time range, which is usually the time range
// that you specified. However, if your resource (such as a CloudFront distribution)
// received 5,000 requests before the specified time range elapsed, GetSampledRequests
// returns an updated time range. This new time range indicates the actual period
// during which AWS WAF selected the requests in the sample.
//
//    // Example sending a request using GetSampledRequestsRequest.
//    req := client.GetSampledRequestsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetSampledRequests
func (c *Client) GetSampledRequestsRequest(input *types.GetSampledRequestsInput) GetSampledRequestsRequest {
	op := &aws.Operation{
		Name:       opGetSampledRequests,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetSampledRequestsInput{}
	}

	req := c.newRequest(op, input, &types.GetSampledRequestsOutput{})
	return GetSampledRequestsRequest{Request: req, Input: input, Copy: c.GetSampledRequestsRequest}
}

// GetSampledRequestsRequest is the request type for the
// GetSampledRequests API operation.
type GetSampledRequestsRequest struct {
	*aws.Request
	Input *types.GetSampledRequestsInput
	Copy  func(*types.GetSampledRequestsInput) GetSampledRequestsRequest
}

// Send marshals and sends the GetSampledRequests API request.
func (r GetSampledRequestsRequest) Send(ctx context.Context) (*GetSampledRequestsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSampledRequestsResponse{
		GetSampledRequestsOutput: r.Request.Data.(*types.GetSampledRequestsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSampledRequestsResponse is the response type for the
// GetSampledRequests API operation.
type GetSampledRequestsResponse struct {
	*types.GetSampledRequestsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSampledRequests request.
func (r *GetSampledRequestsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
