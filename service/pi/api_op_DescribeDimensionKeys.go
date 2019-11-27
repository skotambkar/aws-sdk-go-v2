// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
)

const opDescribeDimensionKeys = "DescribeDimensionKeys"

// DescribeDimensionKeysRequest returns a request value for making API operation for
// AWS Performance Insights.
//
// For a specific time period, retrieve the top N dimension keys for a metric.
//
//    // Example sending a request using DescribeDimensionKeysRequest.
//    req := client.DescribeDimensionKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeys
func (c *Client) DescribeDimensionKeysRequest(input *types.DescribeDimensionKeysInput) DescribeDimensionKeysRequest {
	op := &aws.Operation{
		Name:       opDescribeDimensionKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeDimensionKeysInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDimensionKeysOutput{})
	return DescribeDimensionKeysRequest{Request: req, Input: input, Copy: c.DescribeDimensionKeysRequest}
}

// DescribeDimensionKeysRequest is the request type for the
// DescribeDimensionKeys API operation.
type DescribeDimensionKeysRequest struct {
	*aws.Request
	Input *types.DescribeDimensionKeysInput
	Copy  func(*types.DescribeDimensionKeysInput) DescribeDimensionKeysRequest
}

// Send marshals and sends the DescribeDimensionKeys API request.
func (r DescribeDimensionKeysRequest) Send(ctx context.Context) (*DescribeDimensionKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDimensionKeysResponse{
		DescribeDimensionKeysOutput: r.Request.Data.(*types.DescribeDimensionKeysOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDimensionKeysResponse is the response type for the
// DescribeDimensionKeys API operation.
type DescribeDimensionKeysResponse struct {
	*types.DescribeDimensionKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDimensionKeys request.
func (r *DescribeDimensionKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
