// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetSegmentVersions = "GetSegmentVersions"

// GetSegmentVersionsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the configuration, dimension, and other settings
// for all versions of a specific segment that's associated with an application.
//
//    // Example sending a request using GetSegmentVersionsRequest.
//    req := client.GetSegmentVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegmentVersions
func (c *Client) GetSegmentVersionsRequest(input *types.GetSegmentVersionsInput) GetSegmentVersionsRequest {
	op := &aws.Operation{
		Name:       opGetSegmentVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/segments/{segment-id}/versions",
	}

	if input == nil {
		input = &types.GetSegmentVersionsInput{}
	}

	req := c.newRequest(op, input, &types.GetSegmentVersionsOutput{})
	return GetSegmentVersionsRequest{Request: req, Input: input, Copy: c.GetSegmentVersionsRequest}
}

// GetSegmentVersionsRequest is the request type for the
// GetSegmentVersions API operation.
type GetSegmentVersionsRequest struct {
	*aws.Request
	Input *types.GetSegmentVersionsInput
	Copy  func(*types.GetSegmentVersionsInput) GetSegmentVersionsRequest
}

// Send marshals and sends the GetSegmentVersions API request.
func (r GetSegmentVersionsRequest) Send(ctx context.Context) (*GetSegmentVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSegmentVersionsResponse{
		GetSegmentVersionsOutput: r.Request.Data.(*types.GetSegmentVersionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSegmentVersionsResponse is the response type for the
// GetSegmentVersions API operation.
type GetSegmentVersionsResponse struct {
	*types.GetSegmentVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSegmentVersions request.
func (r *GetSegmentVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
