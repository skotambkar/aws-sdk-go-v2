// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opDeleteDBClusterParameterGroup = "DeleteDBClusterParameterGroup"

// DeleteDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Deletes a specified DB cluster parameter group. The DB cluster parameter
// group to be deleted can't be associated with any DB clusters.
//
//    // Example sending a request using DeleteDBClusterParameterGroupRequest.
//    req := client.DeleteDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DeleteDBClusterParameterGroup
func (c *Client) DeleteDBClusterParameterGroupRequest(input *types.DeleteDBClusterParameterGroupInput) DeleteDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDBClusterParameterGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteDBClusterParameterGroupRequest}
}

// DeleteDBClusterParameterGroupRequest is the request type for the
// DeleteDBClusterParameterGroup API operation.
type DeleteDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *types.DeleteDBClusterParameterGroupInput
	Copy  func(*types.DeleteDBClusterParameterGroupInput) DeleteDBClusterParameterGroupRequest
}

// Send marshals and sends the DeleteDBClusterParameterGroup API request.
func (r DeleteDBClusterParameterGroupRequest) Send(ctx context.Context) (*DeleteDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBClusterParameterGroupResponse{
		DeleteDBClusterParameterGroupOutput: r.Request.Data.(*types.DeleteDBClusterParameterGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBClusterParameterGroupResponse is the response type for the
// DeleteDBClusterParameterGroup API operation.
type DeleteDBClusterParameterGroupResponse struct {
	*types.DeleteDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBClusterParameterGroup request.
func (r *DeleteDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
