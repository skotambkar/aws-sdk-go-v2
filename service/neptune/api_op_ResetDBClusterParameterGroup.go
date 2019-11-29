// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opResetDBClusterParameterGroup = "ResetDBClusterParameterGroup"

// ResetDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Modifies the parameters of a DB cluster parameter group to the default value.
// To reset specific parameters submit a list of the following: ParameterName
// and ApplyMethod. To reset the entire DB cluster parameter group, specify
// the DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When resetting the entire group, dynamic parameters are updated immediately
// and static parameters are set to pending-reboot to take effect on the next
// DB instance restart or RebootDBInstance request. You must call RebootDBInstance
// for every DB instance in your DB cluster that you want the updated static
// parameter to apply to.
//
//    // Example sending a request using ResetDBClusterParameterGroupRequest.
//    req := client.ResetDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/ResetDBClusterParameterGroup
func (c *Client) ResetDBClusterParameterGroupRequest(input *types.ResetDBClusterParameterGroupInput) ResetDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opResetDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResetDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &types.ResetDBClusterParameterGroupOutput{})
	return ResetDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.ResetDBClusterParameterGroupRequest}
}

// ResetDBClusterParameterGroupRequest is the request type for the
// ResetDBClusterParameterGroup API operation.
type ResetDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *types.ResetDBClusterParameterGroupInput
	Copy  func(*types.ResetDBClusterParameterGroupInput) ResetDBClusterParameterGroupRequest
}

// Send marshals and sends the ResetDBClusterParameterGroup API request.
func (r ResetDBClusterParameterGroupRequest) Send(ctx context.Context) (*ResetDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetDBClusterParameterGroupResponse{
		ResetDBClusterParameterGroupOutput: r.Request.Data.(*types.ResetDBClusterParameterGroupOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetDBClusterParameterGroupResponse is the response type for the
// ResetDBClusterParameterGroup API operation.
type ResetDBClusterParameterGroupResponse struct {
	*types.ResetDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetDBClusterParameterGroup request.
func (r *ResetDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
