// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
)

const opDescribeScalingParameters = "DescribeScalingParameters"

// DescribeScalingParametersRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Gets the scaling parameters configured for a domain. A domain's scaling parameters
// specify the desired search instance type and replication count. For more
// information, see Configuring Scaling Options (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DescribeScalingParametersRequest.
//    req := client.DescribeScalingParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeScalingParametersRequest(input *types.DescribeScalingParametersInput) DescribeScalingParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeScalingParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeScalingParametersInput{}
	}

	req := c.newRequest(op, input, &types.DescribeScalingParametersOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeScalingParametersMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeScalingParametersRequest{Request: req, Input: input, Copy: c.DescribeScalingParametersRequest}
}

// DescribeScalingParametersRequest is the request type for the
// DescribeScalingParameters API operation.
type DescribeScalingParametersRequest struct {
	*aws.Request
	Input *types.DescribeScalingParametersInput
	Copy  func(*types.DescribeScalingParametersInput) DescribeScalingParametersRequest
}

// Send marshals and sends the DescribeScalingParameters API request.
func (r DescribeScalingParametersRequest) Send(ctx context.Context) (*DescribeScalingParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScalingParametersResponse{
		DescribeScalingParametersOutput: r.Request.Data.(*types.DescribeScalingParametersOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScalingParametersResponse is the response type for the
// DescribeScalingParameters API operation.
type DescribeScalingParametersResponse struct {
	*types.DescribeScalingParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScalingParameters request.
func (r *DescribeScalingParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
