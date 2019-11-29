// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeVpcEndpointServiceConfigurations = "DescribeVpcEndpointServiceConfigurations"

// DescribeVpcEndpointServiceConfigurationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the VPC endpoint service configurations in your account (your services).
//
//    // Example sending a request using DescribeVpcEndpointServiceConfigurationsRequest.
//    req := client.DescribeVpcEndpointServiceConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcEndpointServiceConfigurations
func (c *Client) DescribeVpcEndpointServiceConfigurationsRequest(input *types.DescribeVpcEndpointServiceConfigurationsInput) DescribeVpcEndpointServiceConfigurationsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcEndpointServiceConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeVpcEndpointServiceConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeVpcEndpointServiceConfigurationsOutput{})
	return DescribeVpcEndpointServiceConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeVpcEndpointServiceConfigurationsRequest}
}

// DescribeVpcEndpointServiceConfigurationsRequest is the request type for the
// DescribeVpcEndpointServiceConfigurations API operation.
type DescribeVpcEndpointServiceConfigurationsRequest struct {
	*aws.Request
	Input *types.DescribeVpcEndpointServiceConfigurationsInput
	Copy  func(*types.DescribeVpcEndpointServiceConfigurationsInput) DescribeVpcEndpointServiceConfigurationsRequest
}

// Send marshals and sends the DescribeVpcEndpointServiceConfigurations API request.
func (r DescribeVpcEndpointServiceConfigurationsRequest) Send(ctx context.Context) (*DescribeVpcEndpointServiceConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcEndpointServiceConfigurationsResponse{
		DescribeVpcEndpointServiceConfigurationsOutput: r.Request.Data.(*types.DescribeVpcEndpointServiceConfigurationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVpcEndpointServiceConfigurationsRequestPaginator returns a paginator for DescribeVpcEndpointServiceConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVpcEndpointServiceConfigurationsRequest(input)
//   p := ec2.NewDescribeVpcEndpointServiceConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVpcEndpointServiceConfigurationsPaginator(req DescribeVpcEndpointServiceConfigurationsRequest) DescribeVpcEndpointServiceConfigurationsPaginator {
	return DescribeVpcEndpointServiceConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeVpcEndpointServiceConfigurationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeVpcEndpointServiceConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVpcEndpointServiceConfigurationsPaginator struct {
	aws.Pager
}

func (p *DescribeVpcEndpointServiceConfigurationsPaginator) CurrentPage() *types.DescribeVpcEndpointServiceConfigurationsOutput {
	return p.Pager.CurrentPage().(*types.DescribeVpcEndpointServiceConfigurationsOutput)
}

// DescribeVpcEndpointServiceConfigurationsResponse is the response type for the
// DescribeVpcEndpointServiceConfigurations API operation.
type DescribeVpcEndpointServiceConfigurationsResponse struct {
	*types.DescribeVpcEndpointServiceConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcEndpointServiceConfigurations request.
func (r *DescribeVpcEndpointServiceConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
