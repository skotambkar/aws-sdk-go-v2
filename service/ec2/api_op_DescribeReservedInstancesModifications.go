// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeReservedInstancesModifications = "DescribeReservedInstancesModifications"

// DescribeReservedInstancesModificationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the modifications made to your Reserved Instances. If no parameter
// is specified, information about all your Reserved Instances modification
// requests is returned. If a modification ID is specified, only information
// about the specific modification is returned.
//
// For more information, see Modifying Reserved Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeReservedInstancesModificationsRequest.
//    req := client.DescribeReservedInstancesModificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeReservedInstancesModifications
func (c *Client) DescribeReservedInstancesModificationsRequest(input *types.DescribeReservedInstancesModificationsInput) DescribeReservedInstancesModificationsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedInstancesModifications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeReservedInstancesModificationsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeReservedInstancesModificationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribeReservedInstancesModificationsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeReservedInstancesModificationsRequest{Request: req, Input: input, Copy: c.DescribeReservedInstancesModificationsRequest}
}

// DescribeReservedInstancesModificationsRequest is the request type for the
// DescribeReservedInstancesModifications API operation.
type DescribeReservedInstancesModificationsRequest struct {
	*aws.Request
	Input *types.DescribeReservedInstancesModificationsInput
	Copy  func(*types.DescribeReservedInstancesModificationsInput) DescribeReservedInstancesModificationsRequest
}

// Send marshals and sends the DescribeReservedInstancesModifications API request.
func (r DescribeReservedInstancesModificationsRequest) Send(ctx context.Context) (*DescribeReservedInstancesModificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedInstancesModificationsResponse{
		DescribeReservedInstancesModificationsOutput: r.Request.Data.(*types.DescribeReservedInstancesModificationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedInstancesModificationsRequestPaginator returns a paginator for DescribeReservedInstancesModifications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedInstancesModificationsRequest(input)
//   p := ec2.NewDescribeReservedInstancesModificationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedInstancesModificationsPaginator(req DescribeReservedInstancesModificationsRequest) DescribeReservedInstancesModificationsPaginator {
	return DescribeReservedInstancesModificationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeReservedInstancesModificationsInput
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

// DescribeReservedInstancesModificationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedInstancesModificationsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedInstancesModificationsPaginator) CurrentPage() *types.DescribeReservedInstancesModificationsOutput {
	return p.Pager.CurrentPage().(*types.DescribeReservedInstancesModificationsOutput)
}

// DescribeReservedInstancesModificationsResponse is the response type for the
// DescribeReservedInstancesModifications API operation.
type DescribeReservedInstancesModificationsResponse struct {
	*types.DescribeReservedInstancesModificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedInstancesModifications request.
func (r *DescribeReservedInstancesModificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
