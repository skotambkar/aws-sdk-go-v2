// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
)

const opSetVisibleToAllUsers = "SetVisibleToAllUsers"

// SetVisibleToAllUsersRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// This member will be deprecated.
//
// Sets whether all AWS Identity and Access Management (IAM) users under your
// account can access the specified clusters (job flows). This action works
// on running clusters. You can also set the visibility of a cluster when you
// launch it using the VisibleToAllUsers parameter of RunJobFlow. The SetVisibleToAllUsers
// action can be called only by an IAM user who created the cluster or the AWS
// account that owns the cluster.
//
//    // Example sending a request using SetVisibleToAllUsersRequest.
//    req := client.SetVisibleToAllUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/SetVisibleToAllUsers
func (c *Client) SetVisibleToAllUsersRequest(input *types.SetVisibleToAllUsersInput) SetVisibleToAllUsersRequest {
	op := &aws.Operation{
		Name:       opSetVisibleToAllUsers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetVisibleToAllUsersInput{}
	}

	req := c.newRequest(op, input, &types.SetVisibleToAllUsersOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetVisibleToAllUsersRequest{Request: req, Input: input, Copy: c.SetVisibleToAllUsersRequest}
}

// SetVisibleToAllUsersRequest is the request type for the
// SetVisibleToAllUsers API operation.
type SetVisibleToAllUsersRequest struct {
	*aws.Request
	Input *types.SetVisibleToAllUsersInput
	Copy  func(*types.SetVisibleToAllUsersInput) SetVisibleToAllUsersRequest
}

// Send marshals and sends the SetVisibleToAllUsers API request.
func (r SetVisibleToAllUsersRequest) Send(ctx context.Context) (*SetVisibleToAllUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetVisibleToAllUsersResponse{
		SetVisibleToAllUsersOutput: r.Request.Data.(*types.SetVisibleToAllUsersOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetVisibleToAllUsersResponse is the response type for the
// SetVisibleToAllUsers API operation.
type SetVisibleToAllUsersResponse struct {
	*types.SetVisibleToAllUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetVisibleToAllUsers request.
func (r *SetVisibleToAllUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
