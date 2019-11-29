// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeDataSet = "DescribeDataSet"

// DescribeDataSetRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes a dataset.
//
// CLI syntax:
//
// aws quicksight describe-data-set \
//
// --aws-account-id=111111111111 \
//
// --data-set-id=unique-data-set-id
//
//    // Example sending a request using DescribeDataSetRequest.
//    req := client.DescribeDataSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeDataSet
func (c *Client) DescribeDataSetRequest(input *types.DescribeDataSetInput) DescribeDataSetRequest {
	op := &aws.Operation{
		Name:       opDescribeDataSet,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}",
	}

	if input == nil {
		input = &types.DescribeDataSetInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDataSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeDataSetMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDataSetRequest{Request: req, Input: input, Copy: c.DescribeDataSetRequest}
}

// DescribeDataSetRequest is the request type for the
// DescribeDataSet API operation.
type DescribeDataSetRequest struct {
	*aws.Request
	Input *types.DescribeDataSetInput
	Copy  func(*types.DescribeDataSetInput) DescribeDataSetRequest
}

// Send marshals and sends the DescribeDataSet API request.
func (r DescribeDataSetRequest) Send(ctx context.Context) (*DescribeDataSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDataSetResponse{
		DescribeDataSetOutput: r.Request.Data.(*types.DescribeDataSetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDataSetResponse is the response type for the
// DescribeDataSet API operation.
type DescribeDataSetResponse struct {
	*types.DescribeDataSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataSet request.
func (r *DescribeDataSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
