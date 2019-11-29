// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opDescribeDBClusterSnapshotAttributes = "DescribeDBClusterSnapshotAttributes"

// DescribeDBClusterSnapshotAttributesRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Returns a list of DB cluster snapshot attribute names and values for a manual
// DB cluster snapshot.
//
// When sharing snapshots with other AWS accounts, DescribeDBClusterSnapshotAttributes
// returns the restore attribute and a list of IDs for the AWS accounts that
// are authorized to copy or restore the manual DB cluster snapshot. If all
// is included in the list of values for the restore attribute, then the manual
// DB cluster snapshot is public and can be copied or restored by all AWS accounts.
//
// To add or remove access for an AWS account to copy or restore a manual DB
// cluster snapshot, or to make the manual DB cluster snapshot public or private,
// use the ModifyDBClusterSnapshotAttribute API action.
//
//    // Example sending a request using DescribeDBClusterSnapshotAttributesRequest.
//    req := client.DescribeDBClusterSnapshotAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBClusterSnapshotAttributes
func (c *Client) DescribeDBClusterSnapshotAttributesRequest(input *types.DescribeDBClusterSnapshotAttributesInput) DescribeDBClusterSnapshotAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterSnapshotAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeDBClusterSnapshotAttributesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDBClusterSnapshotAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeDBClusterSnapshotAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDBClusterSnapshotAttributesRequest{Request: req, Input: input, Copy: c.DescribeDBClusterSnapshotAttributesRequest}
}

// DescribeDBClusterSnapshotAttributesRequest is the request type for the
// DescribeDBClusterSnapshotAttributes API operation.
type DescribeDBClusterSnapshotAttributesRequest struct {
	*aws.Request
	Input *types.DescribeDBClusterSnapshotAttributesInput
	Copy  func(*types.DescribeDBClusterSnapshotAttributesInput) DescribeDBClusterSnapshotAttributesRequest
}

// Send marshals and sends the DescribeDBClusterSnapshotAttributes API request.
func (r DescribeDBClusterSnapshotAttributesRequest) Send(ctx context.Context) (*DescribeDBClusterSnapshotAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterSnapshotAttributesResponse{
		DescribeDBClusterSnapshotAttributesOutput: r.Request.Data.(*types.DescribeDBClusterSnapshotAttributesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterSnapshotAttributesResponse is the response type for the
// DescribeDBClusterSnapshotAttributes API operation.
type DescribeDBClusterSnapshotAttributesResponse struct {
	*types.DescribeDBClusterSnapshotAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterSnapshotAttributes request.
func (r *DescribeDBClusterSnapshotAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
