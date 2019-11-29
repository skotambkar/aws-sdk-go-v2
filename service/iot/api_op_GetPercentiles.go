// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opGetPercentiles = "GetPercentiles"

// GetPercentilesRequest returns a request value for making API operation for
// AWS IoT.
//
// Groups the aggregated values that match the query into percentile groupings.
// The default percentile groupings are: 1,5,25,50,75,95,99, although you can
// specify your own when you call GetPercentiles. This function returns a value
// for each percentile group specified (or the default percentile groupings).
// The percentile group "1" contains the aggregated field value that occurs
// in approximately one percent of the values that match the query. The percentile
// group "5" contains the aggregated field value that occurs in approximately
// five percent of the values that match the query, and so on. The result is
// an approximation, the more values that match the query, the more accurate
// the percentile values.
//
//    // Example sending a request using GetPercentilesRequest.
//    req := client.GetPercentilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetPercentilesRequest(input *types.GetPercentilesInput) GetPercentilesRequest {
	op := &aws.Operation{
		Name:       opGetPercentiles,
		HTTPMethod: "POST",
		HTTPPath:   "/indices/percentiles",
	}

	if input == nil {
		input = &types.GetPercentilesInput{}
	}

	req := c.newRequest(op, input, &types.GetPercentilesOutput{})
	return GetPercentilesRequest{Request: req, Input: input, Copy: c.GetPercentilesRequest}
}

// GetPercentilesRequest is the request type for the
// GetPercentiles API operation.
type GetPercentilesRequest struct {
	*aws.Request
	Input *types.GetPercentilesInput
	Copy  func(*types.GetPercentilesInput) GetPercentilesRequest
}

// Send marshals and sends the GetPercentiles API request.
func (r GetPercentilesRequest) Send(ctx context.Context) (*GetPercentilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPercentilesResponse{
		GetPercentilesOutput: r.Request.Data.(*types.GetPercentilesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPercentilesResponse is the response type for the
// GetPercentiles API operation.
type GetPercentilesResponse struct {
	*types.GetPercentilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPercentiles request.
func (r *GetPercentilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
