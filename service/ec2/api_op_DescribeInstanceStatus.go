// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeInstanceStatus = "DescribeInstanceStatus"

// DescribeInstanceStatusRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the status of the specified instances or all of your instances.
// By default, only running instances are described, unless you specifically
// indicate to return the status of all instances.
//
// Instance status includes the following components:
//
//    * Status checks - Amazon EC2 performs status checks on running EC2 instances
//    to identify hardware and software issues. For more information, see Status
//    Checks for Your Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html)
//    and Troubleshooting Instances with Failed Status Checks (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstances.html)
//    in the Amazon Elastic Compute Cloud User Guide.
//
//    * Scheduled events - Amazon EC2 can schedule events (such as reboot, stop,
//    or terminate) for your instances related to hardware issues, software
//    updates, or system maintenance. For more information, see Scheduled Events
//    for Your Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-instances-status-check_sched.html)
//    in the Amazon Elastic Compute Cloud User Guide.
//
//    * Instance state - You can manage your instances from the moment you launch
//    them through their termination. For more information, see Instance Lifecycle
//    (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
//    in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeInstanceStatusRequest.
//    req := client.DescribeInstanceStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeInstanceStatus
func (c *Client) DescribeInstanceStatusRequest(input *types.DescribeInstanceStatusInput) DescribeInstanceStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceStatus,
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
		input = &types.DescribeInstanceStatusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeInstanceStatusOutput{})
	return DescribeInstanceStatusRequest{Request: req, Input: input, Copy: c.DescribeInstanceStatusRequest}
}

// DescribeInstanceStatusRequest is the request type for the
// DescribeInstanceStatus API operation.
type DescribeInstanceStatusRequest struct {
	*aws.Request
	Input *types.DescribeInstanceStatusInput
	Copy  func(*types.DescribeInstanceStatusInput) DescribeInstanceStatusRequest
}

// Send marshals and sends the DescribeInstanceStatus API request.
func (r DescribeInstanceStatusRequest) Send(ctx context.Context) (*DescribeInstanceStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceStatusResponse{
		DescribeInstanceStatusOutput: r.Request.Data.(*types.DescribeInstanceStatusOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInstanceStatusRequestPaginator returns a paginator for DescribeInstanceStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInstanceStatusRequest(input)
//   p := ec2.NewDescribeInstanceStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInstanceStatusPaginator(req DescribeInstanceStatusRequest) DescribeInstanceStatusPaginator {
	return DescribeInstanceStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeInstanceStatusInput
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

// DescribeInstanceStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInstanceStatusPaginator struct {
	aws.Pager
}

func (p *DescribeInstanceStatusPaginator) CurrentPage() *types.DescribeInstanceStatusOutput {
	return p.Pager.CurrentPage().(*types.DescribeInstanceStatusOutput)
}

// DescribeInstanceStatusResponse is the response type for the
// DescribeInstanceStatus API operation.
type DescribeInstanceStatusResponse struct {
	*types.DescribeInstanceStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceStatus request.
func (r *DescribeInstanceStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
