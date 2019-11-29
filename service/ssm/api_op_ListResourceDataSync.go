// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opListResourceDataSync = "ListResourceDataSync"

// ListResourceDataSyncRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Lists your resource data sync configurations. Includes information about
// the last time a sync attempted to start, the last sync status, and the last
// time a sync successfully completed.
//
// The number of sync configurations might be too large to return using a single
// call to ListResourceDataSync. You can limit the number of sync configurations
// returned by using the MaxResults parameter. To determine whether there are
// more sync configurations to list, check the value of NextToken in the output.
// If there are more sync configurations to list, you can request them by specifying
// the NextToken returned in the call to the parameter of a subsequent call.
//
//    // Example sending a request using ListResourceDataSyncRequest.
//    req := client.ListResourceDataSyncRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListResourceDataSync
func (c *Client) ListResourceDataSyncRequest(input *types.ListResourceDataSyncInput) ListResourceDataSyncRequest {
	op := &aws.Operation{
		Name:       opListResourceDataSync,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListResourceDataSyncInput{}
	}

	req := c.newRequest(op, input, &types.ListResourceDataSyncOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListResourceDataSyncMarshaler{Input: input}.GetNamedBuildHandler())

	return ListResourceDataSyncRequest{Request: req, Input: input, Copy: c.ListResourceDataSyncRequest}
}

// ListResourceDataSyncRequest is the request type for the
// ListResourceDataSync API operation.
type ListResourceDataSyncRequest struct {
	*aws.Request
	Input *types.ListResourceDataSyncInput
	Copy  func(*types.ListResourceDataSyncInput) ListResourceDataSyncRequest
}

// Send marshals and sends the ListResourceDataSync API request.
func (r ListResourceDataSyncRequest) Send(ctx context.Context) (*ListResourceDataSyncResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourceDataSyncResponse{
		ListResourceDataSyncOutput: r.Request.Data.(*types.ListResourceDataSyncOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResourceDataSyncResponse is the response type for the
// ListResourceDataSync API operation.
type ListResourceDataSyncResponse struct {
	*types.ListResourceDataSyncOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourceDataSync request.
func (r *ListResourceDataSyncResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
