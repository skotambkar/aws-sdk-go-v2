// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opDeleteSegment = "DeleteSegment"

// DeleteSegmentRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes a segment from an application.
//
//    // Example sending a request using DeleteSegmentRequest.
//    req := client.DeleteSegmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteSegment
func (c *Client) DeleteSegmentRequest(input *types.DeleteSegmentInput) DeleteSegmentRequest {
	op := &aws.Operation{
		Name:       opDeleteSegment,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/segments/{segment-id}",
	}

	if input == nil {
		input = &types.DeleteSegmentInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSegmentOutput{})
	return DeleteSegmentRequest{Request: req, Input: input, Copy: c.DeleteSegmentRequest}
}

// DeleteSegmentRequest is the request type for the
// DeleteSegment API operation.
type DeleteSegmentRequest struct {
	*aws.Request
	Input *types.DeleteSegmentInput
	Copy  func(*types.DeleteSegmentInput) DeleteSegmentRequest
}

// Send marshals and sends the DeleteSegment API request.
func (r DeleteSegmentRequest) Send(ctx context.Context) (*DeleteSegmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSegmentResponse{
		DeleteSegmentOutput: r.Request.Data.(*types.DeleteSegmentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSegmentResponse is the response type for the
// DeleteSegment API operation.
type DeleteSegmentResponse struct {
	*types.DeleteSegmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSegment request.
func (r *DeleteSegmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
