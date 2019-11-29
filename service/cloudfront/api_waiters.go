// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

// WaitUntilDistributionDeployed uses the CloudFront API operation
// GetDistribution to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDistributionDeployed(ctx context.Context, input *types.GetDistributionInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDistributionDeployed",
		MaxAttempts: 35,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Distribution.Status",
				Expected: "Deployed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *types.GetDistributionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetDistributionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInvalidationCompleted uses the CloudFront API operation
// GetInvalidation to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilInvalidationCompleted(ctx context.Context, input *types.GetInvalidationInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInvalidationCompleted",
		MaxAttempts: 30,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Invalidation.Status",
				Expected: "Completed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *types.GetInvalidationInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetInvalidationRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilStreamingDistributionDeployed uses the CloudFront API operation
// GetStreamingDistribution to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilStreamingDistributionDeployed(ctx context.Context, input *types.GetStreamingDistributionInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilStreamingDistributionDeployed",
		MaxAttempts: 25,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "StreamingDistribution.Status",
				Expected: "Deployed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *types.GetStreamingDistributionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetStreamingDistributionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
