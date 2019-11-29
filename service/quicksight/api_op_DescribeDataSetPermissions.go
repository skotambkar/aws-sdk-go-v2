// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeDataSetPermissions = "DescribeDataSetPermissions"

// DescribeDataSetPermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes the permissions on a dataset.
//
// The permissions resource is arn:aws:quicksight:region:aws-account-id:dataset/data-set-id
//
// CLI syntax:
//
// aws quicksight describe-data-set-permissions \
//
// --aws-account-id=111122223333 \
//
// --data-set-id=unique-data-set-id \
//
//    // Example sending a request using DescribeDataSetPermissionsRequest.
//    req := client.DescribeDataSetPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeDataSetPermissions
func (c *Client) DescribeDataSetPermissionsRequest(input *types.DescribeDataSetPermissionsInput) DescribeDataSetPermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeDataSetPermissions,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}/permissions",
	}

	if input == nil {
		input = &types.DescribeDataSetPermissionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDataSetPermissionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeDataSetPermissionsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDataSetPermissionsRequest{Request: req, Input: input, Copy: c.DescribeDataSetPermissionsRequest}
}

// DescribeDataSetPermissionsRequest is the request type for the
// DescribeDataSetPermissions API operation.
type DescribeDataSetPermissionsRequest struct {
	*aws.Request
	Input *types.DescribeDataSetPermissionsInput
	Copy  func(*types.DescribeDataSetPermissionsInput) DescribeDataSetPermissionsRequest
}

// Send marshals and sends the DescribeDataSetPermissions API request.
func (r DescribeDataSetPermissionsRequest) Send(ctx context.Context) (*DescribeDataSetPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDataSetPermissionsResponse{
		DescribeDataSetPermissionsOutput: r.Request.Data.(*types.DescribeDataSetPermissionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDataSetPermissionsResponse is the response type for the
// DescribeDataSetPermissions API operation.
type DescribeDataSetPermissionsResponse struct {
	*types.DescribeDataSetPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataSetPermissions request.
func (r *DescribeDataSetPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
