// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opDescribeDBClusterSnapshots = "DescribeDBClusterSnapshots"

// DescribeDBClusterSnapshotsRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Returns information about DB cluster snapshots. This API action supports
// pagination.
//
//    // Example sending a request using DescribeDBClusterSnapshotsRequest.
//    req := client.DescribeDBClusterSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBClusterSnapshots
func (c *Client) DescribeDBClusterSnapshotsRequest(input *types.DescribeDBClusterSnapshotsInput) DescribeDBClusterSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeDBClusterSnapshotsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDBClusterSnapshotsOutput{})
	return DescribeDBClusterSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeDBClusterSnapshotsRequest}
}

// DescribeDBClusterSnapshotsRequest is the request type for the
// DescribeDBClusterSnapshots API operation.
type DescribeDBClusterSnapshotsRequest struct {
	*aws.Request
	Input *types.DescribeDBClusterSnapshotsInput
	Copy  func(*types.DescribeDBClusterSnapshotsInput) DescribeDBClusterSnapshotsRequest
}

// Send marshals and sends the DescribeDBClusterSnapshots API request.
func (r DescribeDBClusterSnapshotsRequest) Send(ctx context.Context) (*DescribeDBClusterSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterSnapshotsResponse{
		DescribeDBClusterSnapshotsOutput: r.Request.Data.(*types.DescribeDBClusterSnapshotsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterSnapshotsResponse is the response type for the
// DescribeDBClusterSnapshots API operation.
type DescribeDBClusterSnapshotsResponse struct {
	*types.DescribeDBClusterSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterSnapshots request.
func (r *DescribeDBClusterSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
