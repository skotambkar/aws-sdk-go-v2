// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opDescribeAccountLimits = "DescribeAccountLimits"

// DescribeAccountLimitsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the current Elastic Load Balancing resource limits for your AWS
// account.
//
// For more information, see Limits for Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using DescribeAccountLimitsRequest.
//    req := client.DescribeAccountLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeAccountLimits
func (c *Client) DescribeAccountLimitsRequest(input *types.DescribeAccountLimitsInput) DescribeAccountLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAccountLimitsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAccountLimitsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeAccountLimitsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeAccountLimitsRequest{Request: req, Input: input, Copy: c.DescribeAccountLimitsRequest}
}

// DescribeAccountLimitsRequest is the request type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsRequest struct {
	*aws.Request
	Input *types.DescribeAccountLimitsInput
	Copy  func(*types.DescribeAccountLimitsInput) DescribeAccountLimitsRequest
}

// Send marshals and sends the DescribeAccountLimits API request.
func (r DescribeAccountLimitsRequest) Send(ctx context.Context) (*DescribeAccountLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountLimitsResponse{
		DescribeAccountLimitsOutput: r.Request.Data.(*types.DescribeAccountLimitsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountLimitsResponse is the response type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsResponse struct {
	*types.DescribeAccountLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountLimits request.
func (r *DescribeAccountLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
