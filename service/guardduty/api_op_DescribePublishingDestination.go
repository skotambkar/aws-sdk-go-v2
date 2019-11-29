// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
)

const opDescribePublishingDestination = "DescribePublishingDestination"

// DescribePublishingDestinationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Returns information about the publishing destination specified by the provided
// destinationId.
//
//    // Example sending a request using DescribePublishingDestinationRequest.
//    req := client.DescribePublishingDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DescribePublishingDestination
func (c *Client) DescribePublishingDestinationRequest(input *types.DescribePublishingDestinationInput) DescribePublishingDestinationRequest {
	op := &aws.Operation{
		Name:       opDescribePublishingDestination,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/publishingDestination/{destinationId}",
	}

	if input == nil {
		input = &types.DescribePublishingDestinationInput{}
	}

	req := c.newRequest(op, input, &types.DescribePublishingDestinationOutput{})
	return DescribePublishingDestinationRequest{Request: req, Input: input, Copy: c.DescribePublishingDestinationRequest}
}

// DescribePublishingDestinationRequest is the request type for the
// DescribePublishingDestination API operation.
type DescribePublishingDestinationRequest struct {
	*aws.Request
	Input *types.DescribePublishingDestinationInput
	Copy  func(*types.DescribePublishingDestinationInput) DescribePublishingDestinationRequest
}

// Send marshals and sends the DescribePublishingDestination API request.
func (r DescribePublishingDestinationRequest) Send(ctx context.Context) (*DescribePublishingDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePublishingDestinationResponse{
		DescribePublishingDestinationOutput: r.Request.Data.(*types.DescribePublishingDestinationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePublishingDestinationResponse is the response type for the
// DescribePublishingDestination API operation.
type DescribePublishingDestinationResponse struct {
	*types.DescribePublishingDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePublishingDestination request.
func (r *DescribePublishingDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
