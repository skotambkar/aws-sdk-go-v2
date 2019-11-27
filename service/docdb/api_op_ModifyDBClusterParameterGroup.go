// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
)

const opModifyDBClusterParameterGroup = "ModifyDBClusterParameterGroup"

// ModifyDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Modifies the parameters of a DB cluster parameter group. To modify more than
// one parameter, submit a list of the following: ParameterName, ParameterValue,
// and ApplyMethod. A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot or maintenance window before the change can take
// effect.
//
// After you create a DB cluster parameter group, you should wait at least 5
// minutes before creating your first DB cluster that uses that DB cluster parameter
// group as the default parameter group. This allows Amazon DocumentDB to fully
// complete the create action before the parameter group is used as the default
// for a new DB cluster. This step is especially important for parameters that
// are critical when creating the default database for a DB cluster, such as
// the character set for the default database defined by the character_set_database
// parameter.
//
//    // Example sending a request using ModifyDBClusterParameterGroupRequest.
//    req := client.ModifyDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/ModifyDBClusterParameterGroup
func (c *Client) ModifyDBClusterParameterGroupRequest(input *types.ModifyDBClusterParameterGroupInput) ModifyDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opModifyDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &types.ModifyDBClusterParameterGroupOutput{})
	return ModifyDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.ModifyDBClusterParameterGroupRequest}
}

// ModifyDBClusterParameterGroupRequest is the request type for the
// ModifyDBClusterParameterGroup API operation.
type ModifyDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *types.ModifyDBClusterParameterGroupInput
	Copy  func(*types.ModifyDBClusterParameterGroupInput) ModifyDBClusterParameterGroupRequest
}

// Send marshals and sends the ModifyDBClusterParameterGroup API request.
func (r ModifyDBClusterParameterGroupRequest) Send(ctx context.Context) (*ModifyDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBClusterParameterGroupResponse{
		ModifyDBClusterParameterGroupOutput: r.Request.Data.(*types.ModifyDBClusterParameterGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBClusterParameterGroupResponse is the response type for the
// ModifyDBClusterParameterGroup API operation.
type ModifyDBClusterParameterGroupResponse struct {
	*types.ModifyDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBClusterParameterGroup request.
func (r *ModifyDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
