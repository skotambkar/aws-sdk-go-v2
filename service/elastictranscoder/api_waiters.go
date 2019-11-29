// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
)

// WaitUntilJobComplete uses the Amazon Elastic Transcoder API operation
// ReadJob to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilJobComplete(ctx context.Context, input *types.ReadJobInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilJobComplete",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Job.Status",
				Expected: "Complete",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Job.Status",
				Expected: "Canceled",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Job.Status",
				Expected: "Error",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *types.ReadJobInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.ReadJobRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
