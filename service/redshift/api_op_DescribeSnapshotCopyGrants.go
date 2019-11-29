// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opDescribeSnapshotCopyGrants = "DescribeSnapshotCopyGrants"

// DescribeSnapshotCopyGrantsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns a list of snapshot copy grants owned by the AWS account in the destination
// region.
//
// For more information about managing snapshot copy grants, go to Amazon Redshift
// Database Encryption (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using DescribeSnapshotCopyGrantsRequest.
//    req := client.DescribeSnapshotCopyGrantsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeSnapshotCopyGrants
func (c *Client) DescribeSnapshotCopyGrantsRequest(input *types.DescribeSnapshotCopyGrantsInput) DescribeSnapshotCopyGrantsRequest {
	op := &aws.Operation{
		Name:       opDescribeSnapshotCopyGrants,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeSnapshotCopyGrantsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSnapshotCopyGrantsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeSnapshotCopyGrantsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeSnapshotCopyGrantsRequest{Request: req, Input: input, Copy: c.DescribeSnapshotCopyGrantsRequest}
}

// DescribeSnapshotCopyGrantsRequest is the request type for the
// DescribeSnapshotCopyGrants API operation.
type DescribeSnapshotCopyGrantsRequest struct {
	*aws.Request
	Input *types.DescribeSnapshotCopyGrantsInput
	Copy  func(*types.DescribeSnapshotCopyGrantsInput) DescribeSnapshotCopyGrantsRequest
}

// Send marshals and sends the DescribeSnapshotCopyGrants API request.
func (r DescribeSnapshotCopyGrantsRequest) Send(ctx context.Context) (*DescribeSnapshotCopyGrantsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSnapshotCopyGrantsResponse{
		DescribeSnapshotCopyGrantsOutput: r.Request.Data.(*types.DescribeSnapshotCopyGrantsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSnapshotCopyGrantsResponse is the response type for the
// DescribeSnapshotCopyGrants API operation.
type DescribeSnapshotCopyGrantsResponse struct {
	*types.DescribeSnapshotCopyGrantsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSnapshotCopyGrants request.
func (r *DescribeSnapshotCopyGrantsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
