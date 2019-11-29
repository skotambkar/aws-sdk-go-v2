// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opEstimateTemplateCost = "EstimateTemplateCost"

// EstimateTemplateCostRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns the estimated monthly cost of a template. The return value is an
// AWS Simple Monthly Calculator URL with a query string that describes the
// resources required to run the template.
//
//    // Example sending a request using EstimateTemplateCostRequest.
//    req := client.EstimateTemplateCostRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/EstimateTemplateCost
func (c *Client) EstimateTemplateCostRequest(input *types.EstimateTemplateCostInput) EstimateTemplateCostRequest {
	op := &aws.Operation{
		Name:       opEstimateTemplateCost,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.EstimateTemplateCostInput{}
	}

	req := c.newRequest(op, input, &types.EstimateTemplateCostOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.EstimateTemplateCostMarshaler{Input: input}.GetNamedBuildHandler())

	return EstimateTemplateCostRequest{Request: req, Input: input, Copy: c.EstimateTemplateCostRequest}
}

// EstimateTemplateCostRequest is the request type for the
// EstimateTemplateCost API operation.
type EstimateTemplateCostRequest struct {
	*aws.Request
	Input *types.EstimateTemplateCostInput
	Copy  func(*types.EstimateTemplateCostInput) EstimateTemplateCostRequest
}

// Send marshals and sends the EstimateTemplateCost API request.
func (r EstimateTemplateCostRequest) Send(ctx context.Context) (*EstimateTemplateCostResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EstimateTemplateCostResponse{
		EstimateTemplateCostOutput: r.Request.Data.(*types.EstimateTemplateCostOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EstimateTemplateCostResponse is the response type for the
// EstimateTemplateCost API operation.
type EstimateTemplateCostResponse struct {
	*types.EstimateTemplateCostOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EstimateTemplateCost request.
func (r *EstimateTemplateCostResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
