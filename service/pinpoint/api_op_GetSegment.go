// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetSegment = "GetSegment"

// GetSegmentRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the configuration, dimension, and other settings
// for a specific segment that's associated with an application.
//
//    // Example sending a request using GetSegmentRequest.
//    req := client.GetSegmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegment
func (c *Client) GetSegmentRequest(input *types.GetSegmentInput) GetSegmentRequest {
	op := &aws.Operation{
		Name:       opGetSegment,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/segments/{segment-id}",
	}

	if input == nil {
		input = &types.GetSegmentInput{}
	}

	req := c.newRequest(op, input, &types.GetSegmentOutput{})
	return GetSegmentRequest{Request: req, Input: input, Copy: c.GetSegmentRequest}
}

// GetSegmentRequest is the request type for the
// GetSegment API operation.
type GetSegmentRequest struct {
	*aws.Request
	Input *types.GetSegmentInput
	Copy  func(*types.GetSegmentInput) GetSegmentRequest
}

// Send marshals and sends the GetSegment API request.
func (r GetSegmentRequest) Send(ctx context.Context) (*GetSegmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSegmentResponse{
		GetSegmentOutput: r.Request.Data.(*types.GetSegmentOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSegmentResponse is the response type for the
// GetSegment API operation.
type GetSegmentResponse struct {
	*types.GetSegmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSegment request.
func (r *GetSegmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
