// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering/types"
)

const opRegisterUsage = "RegisterUsage"

// RegisterUsageRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// Paid container software products sold through AWS Marketplace must integrate
// with the AWS Marketplace Metering Service and call the RegisterUsage operation
// for software entitlement and metering. Free and BYOL products for Amazon
// ECS or Amazon EKS aren't required to call RegisterUsage, but you may choose
// to do so if you would like to receive usage data in your seller reports.
// The sections below explain the behavior of RegisterUsage. RegisterUsage performs
// two primary functions: metering and entitlement.
//
//    * Entitlement: RegisterUsage allows you to verify that the customer running
//    your paid software is subscribed to your product on AWS Marketplace, enabling
//    you to guard against unauthorized use. Your container image that integrates
//    with RegisterUsage is only required to guard against unauthorized use
//    at container startup, as such a CustomerNotSubscribedException/PlatformNotSupportedException
//    will only be thrown on the initial call to RegisterUsage. Subsequent calls
//    from the same Amazon ECS task instance (e.g. task-id) or Amazon EKS pod
//    will not throw a CustomerNotSubscribedException, even if the customer
//    unsubscribes while the Amazon ECS task or Amazon EKS pod is still running.
//
//    * Metering: RegisterUsage meters software use per ECS task, per hour,
//    or per pod for Amazon EKS with usage prorated to the second. A minimum
//    of 1 minute of usage applies to tasks that are short lived. For example,
//    if a customer has a 10 node Amazon ECS or Amazon EKS cluster and a service
//    configured as a Daemon Set, then Amazon ECS or Amazon EKS will launch
//    a task on all 10 cluster nodes and the customer will be charged: (10 *
//    hourly_rate). Metering for software use is automatically handled by the
//    AWS Marketplace Metering Control Plane -- your software is not required
//    to perform any metering specific actions, other than call RegisterUsage
//    once for metering of software use to commence. The AWS Marketplace Metering
//    Control Plane will also continue to bill customers for running ECS tasks
//    and Amazon EKS pods, regardless of the customers subscription state, removing
//    the need for your software to perform entitlement checks at runtime.
//
//    // Example sending a request using RegisterUsageRequest.
//    req := client.RegisterUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/RegisterUsage
func (c *Client) RegisterUsageRequest(input *types.RegisterUsageInput) RegisterUsageRequest {
	op := &aws.Operation{
		Name:       opRegisterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterUsageInput{}
	}

	req := c.newRequest(op, input, &types.RegisterUsageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RegisterUsageMarshaler{Input: input}.GetNamedBuildHandler())

	return RegisterUsageRequest{Request: req, Input: input, Copy: c.RegisterUsageRequest}
}

// RegisterUsageRequest is the request type for the
// RegisterUsage API operation.
type RegisterUsageRequest struct {
	*aws.Request
	Input *types.RegisterUsageInput
	Copy  func(*types.RegisterUsageInput) RegisterUsageRequest
}

// Send marshals and sends the RegisterUsage API request.
func (r RegisterUsageRequest) Send(ctx context.Context) (*RegisterUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterUsageResponse{
		RegisterUsageOutput: r.Request.Data.(*types.RegisterUsageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterUsageResponse is the response type for the
// RegisterUsage API operation.
type RegisterUsageResponse struct {
	*types.RegisterUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterUsage request.
func (r *RegisterUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
