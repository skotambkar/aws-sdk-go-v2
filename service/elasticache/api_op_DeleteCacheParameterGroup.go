// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opDeleteCacheParameterGroup = "DeleteCacheParameterGroup"

// DeleteCacheParameterGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deletes the specified cache parameter group. You cannot delete a cache parameter
// group if it is associated with any cache clusters.
//
//    // Example sending a request using DeleteCacheParameterGroupRequest.
//    req := client.DeleteCacheParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheParameterGroup
func (c *Client) DeleteCacheParameterGroupRequest(input *types.DeleteCacheParameterGroupInput) DeleteCacheParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteCacheParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteCacheParameterGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteCacheParameterGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCacheParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteCacheParameterGroupRequest}
}

// DeleteCacheParameterGroupRequest is the request type for the
// DeleteCacheParameterGroup API operation.
type DeleteCacheParameterGroupRequest struct {
	*aws.Request
	Input *types.DeleteCacheParameterGroupInput
	Copy  func(*types.DeleteCacheParameterGroupInput) DeleteCacheParameterGroupRequest
}

// Send marshals and sends the DeleteCacheParameterGroup API request.
func (r DeleteCacheParameterGroupRequest) Send(ctx context.Context) (*DeleteCacheParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCacheParameterGroupResponse{
		DeleteCacheParameterGroupOutput: r.Request.Data.(*types.DeleteCacheParameterGroupOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCacheParameterGroupResponse is the response type for the
// DeleteCacheParameterGroup API operation.
type DeleteCacheParameterGroupResponse struct {
	*types.DeleteCacheParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCacheParameterGroup request.
func (r *DeleteCacheParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
