// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

const opDeleteTargetGroup = "DeleteTargetGroup"

// DeleteTargetGroupRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Deletes the specified target group.
//
// You can delete a target group if it is not referenced by any actions. Deleting
// a target group also deletes any associated health checks.
//
//    // Example sending a request using DeleteTargetGroupRequest.
//    req := client.DeleteTargetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DeleteTargetGroup
func (c *Client) DeleteTargetGroupRequest(input *types.DeleteTargetGroupInput) DeleteTargetGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteTargetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTargetGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTargetGroupOutput{})
	return DeleteTargetGroupRequest{Request: req, Input: input, Copy: c.DeleteTargetGroupRequest}
}

// DeleteTargetGroupRequest is the request type for the
// DeleteTargetGroup API operation.
type DeleteTargetGroupRequest struct {
	*aws.Request
	Input *types.DeleteTargetGroupInput
	Copy  func(*types.DeleteTargetGroupInput) DeleteTargetGroupRequest
}

// Send marshals and sends the DeleteTargetGroup API request.
func (r DeleteTargetGroupRequest) Send(ctx context.Context) (*DeleteTargetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTargetGroupResponse{
		DeleteTargetGroupOutput: r.Request.Data.(*types.DeleteTargetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTargetGroupResponse is the response type for the
// DeleteTargetGroup API operation.
type DeleteTargetGroupResponse struct {
	*types.DeleteTargetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTargetGroup request.
func (r *DeleteTargetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
