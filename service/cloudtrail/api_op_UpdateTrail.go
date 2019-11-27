// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

const opUpdateTrail = "UpdateTrail"

// UpdateTrailRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Updates the settings that specify delivery of log files. Changes to a trail
// do not require stopping the CloudTrail service. Use this action to designate
// an existing bucket for log delivery. If the existing bucket has previously
// been a target for CloudTrail log files, an IAM policy exists for the bucket.
// UpdateTrail must be called from the region in which the trail was created;
// otherwise, an InvalidHomeRegionException is thrown.
//
//    // Example sending a request using UpdateTrailRequest.
//    req := client.UpdateTrailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/UpdateTrail
func (c *Client) UpdateTrailRequest(input *types.UpdateTrailInput) UpdateTrailRequest {
	op := &aws.Operation{
		Name:       opUpdateTrail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateTrailInput{}
	}

	req := c.newRequest(op, input, &types.UpdateTrailOutput{})
	return UpdateTrailRequest{Request: req, Input: input, Copy: c.UpdateTrailRequest}
}

// UpdateTrailRequest is the request type for the
// UpdateTrail API operation.
type UpdateTrailRequest struct {
	*aws.Request
	Input *types.UpdateTrailInput
	Copy  func(*types.UpdateTrailInput) UpdateTrailRequest
}

// Send marshals and sends the UpdateTrail API request.
func (r UpdateTrailRequest) Send(ctx context.Context) (*UpdateTrailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrailResponse{
		UpdateTrailOutput: r.Request.Data.(*types.UpdateTrailOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrailResponse is the response type for the
// UpdateTrail API operation.
type UpdateTrailResponse struct {
	*types.UpdateTrailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrail request.
func (r *UpdateTrailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
