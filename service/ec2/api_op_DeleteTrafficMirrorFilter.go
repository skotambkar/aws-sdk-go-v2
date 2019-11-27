// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteTrafficMirrorFilter = "DeleteTrafficMirrorFilter"

// DeleteTrafficMirrorFilterRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Traffic Mirror filter.
//
// You cannot delete a Traffic Mirror filter that is in use by a Traffic Mirror
// session.
//
//    // Example sending a request using DeleteTrafficMirrorFilterRequest.
//    req := client.DeleteTrafficMirrorFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorFilter
func (c *Client) DeleteTrafficMirrorFilterRequest(input *types.DeleteTrafficMirrorFilterInput) DeleteTrafficMirrorFilterRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficMirrorFilter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTrafficMirrorFilterInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTrafficMirrorFilterOutput{})
	return DeleteTrafficMirrorFilterRequest{Request: req, Input: input, Copy: c.DeleteTrafficMirrorFilterRequest}
}

// DeleteTrafficMirrorFilterRequest is the request type for the
// DeleteTrafficMirrorFilter API operation.
type DeleteTrafficMirrorFilterRequest struct {
	*aws.Request
	Input *types.DeleteTrafficMirrorFilterInput
	Copy  func(*types.DeleteTrafficMirrorFilterInput) DeleteTrafficMirrorFilterRequest
}

// Send marshals and sends the DeleteTrafficMirrorFilter API request.
func (r DeleteTrafficMirrorFilterRequest) Send(ctx context.Context) (*DeleteTrafficMirrorFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficMirrorFilterResponse{
		DeleteTrafficMirrorFilterOutput: r.Request.Data.(*types.DeleteTrafficMirrorFilterOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficMirrorFilterResponse is the response type for the
// DeleteTrafficMirrorFilter API operation.
type DeleteTrafficMirrorFilterResponse struct {
	*types.DeleteTrafficMirrorFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficMirrorFilter request.
func (r *DeleteTrafficMirrorFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
