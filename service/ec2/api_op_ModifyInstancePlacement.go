// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyInstancePlacement = "ModifyInstancePlacement"

// ModifyInstancePlacementRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the placement attributes for a specified instance. You can do the
// following:
//
//    * Modify the affinity between an instance and a Dedicated Host (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html).
//    When affinity is set to host and the instance is not associated with a
//    specific Dedicated Host, the next time the instance is launched, it is
//    automatically associated with the host on which it lands. If the instance
//    is restarted or rebooted, this relationship persists.
//
//    * Change the Dedicated Host with which an instance is associated.
//
//    * Change the instance tenancy of an instance from host to dedicated, or
//    from dedicated to host.
//
//    * Move an instance to or from a placement group (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
//
// At least one attribute for affinity, host ID, tenancy, or placement group
// name must be specified in the request. Affinity and tenancy can be modified
// in the same request.
//
// To modify the host ID, tenancy, placement group, or partition for an instance,
// the instance must be in the stopped state.
//
//    // Example sending a request using ModifyInstancePlacementRequest.
//    req := client.ModifyInstancePlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstancePlacement
func (c *Client) ModifyInstancePlacementRequest(input *types.ModifyInstancePlacementInput) ModifyInstancePlacementRequest {
	op := &aws.Operation{
		Name:       opModifyInstancePlacement,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyInstancePlacementInput{}
	}

	req := c.newRequest(op, input, &types.ModifyInstancePlacementOutput{})
	return ModifyInstancePlacementRequest{Request: req, Input: input, Copy: c.ModifyInstancePlacementRequest}
}

// ModifyInstancePlacementRequest is the request type for the
// ModifyInstancePlacement API operation.
type ModifyInstancePlacementRequest struct {
	*aws.Request
	Input *types.ModifyInstancePlacementInput
	Copy  func(*types.ModifyInstancePlacementInput) ModifyInstancePlacementRequest
}

// Send marshals and sends the ModifyInstancePlacement API request.
func (r ModifyInstancePlacementRequest) Send(ctx context.Context) (*ModifyInstancePlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyInstancePlacementResponse{
		ModifyInstancePlacementOutput: r.Request.Data.(*types.ModifyInstancePlacementOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyInstancePlacementResponse is the response type for the
// ModifyInstancePlacement API operation.
type ModifyInstancePlacementResponse struct {
	*types.ModifyInstancePlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyInstancePlacement request.
func (r *ModifyInstancePlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
