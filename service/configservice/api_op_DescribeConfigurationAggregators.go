// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDescribeConfigurationAggregators = "DescribeConfigurationAggregators"

// DescribeConfigurationAggregatorsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the details of one or more configuration aggregators. If the configuration
// aggregator is not specified, this action returns the details for all the
// configuration aggregators associated with the account.
//
//    // Example sending a request using DescribeConfigurationAggregatorsRequest.
//    req := client.DescribeConfigurationAggregatorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationAggregators
func (c *Client) DescribeConfigurationAggregatorsRequest(input *types.DescribeConfigurationAggregatorsInput) DescribeConfigurationAggregatorsRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurationAggregators,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeConfigurationAggregatorsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeConfigurationAggregatorsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeConfigurationAggregatorsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeConfigurationAggregatorsRequest{Request: req, Input: input, Copy: c.DescribeConfigurationAggregatorsRequest}
}

// DescribeConfigurationAggregatorsRequest is the request type for the
// DescribeConfigurationAggregators API operation.
type DescribeConfigurationAggregatorsRequest struct {
	*aws.Request
	Input *types.DescribeConfigurationAggregatorsInput
	Copy  func(*types.DescribeConfigurationAggregatorsInput) DescribeConfigurationAggregatorsRequest
}

// Send marshals and sends the DescribeConfigurationAggregators API request.
func (r DescribeConfigurationAggregatorsRequest) Send(ctx context.Context) (*DescribeConfigurationAggregatorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationAggregatorsResponse{
		DescribeConfigurationAggregatorsOutput: r.Request.Data.(*types.DescribeConfigurationAggregatorsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationAggregatorsResponse is the response type for the
// DescribeConfigurationAggregators API operation.
type DescribeConfigurationAggregatorsResponse struct {
	*types.DescribeConfigurationAggregatorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurationAggregators request.
func (r *DescribeConfigurationAggregatorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
