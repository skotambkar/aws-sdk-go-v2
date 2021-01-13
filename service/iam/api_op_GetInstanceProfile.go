// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Retrieves information about the specified instance profile, including the
// instance profile's path, GUID, ARN, and role. For more information about
// instance profiles, see About Instance Profiles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html) in
// the IAM User Guide.
func (c *Client) GetInstanceProfile(ctx context.Context, params *GetInstanceProfileInput, optFns ...func(*Options)) (*GetInstanceProfileOutput, error) {
	if params == nil {
		params = &GetInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceProfile", params, optFns, addOperationGetInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceProfileInput struct {

	// The name of the instance profile to get information about. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	InstanceProfileName *string
}

// Contains the response to a successful GetInstanceProfile request.
type GetInstanceProfileOutput struct {

	// A structure containing details about the instance profile.
	//
	// This member is required.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetInstanceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceProfile(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetInstanceProfileAPIClient is a client that implements the GetInstanceProfile
// operation.
type GetInstanceProfileAPIClient interface {
	GetInstanceProfile(context.Context, *GetInstanceProfileInput, ...func(*Options)) (*GetInstanceProfileOutput, error)
}

var _ GetInstanceProfileAPIClient = (*Client)(nil)

// InstanceProfileExistsWaiterOptions are waiter options for
// InstanceProfileExistsWaiter
type InstanceProfileExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InstanceProfileExistsWaiter will use default minimum delay of 1 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, InstanceProfileExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetInstanceProfileInput, *GetInstanceProfileOutput, error) (bool, error)
}

// InstanceProfileExistsWaiter defines the waiters for InstanceProfileExists
type InstanceProfileExistsWaiter struct {
	client GetInstanceProfileAPIClient

	options InstanceProfileExistsWaiterOptions
}

// NewInstanceProfileExistsWaiter constructs a InstanceProfileExistsWaiter.
func NewInstanceProfileExistsWaiter(client GetInstanceProfileAPIClient, optFns ...func(*InstanceProfileExistsWaiterOptions)) *InstanceProfileExistsWaiter {
	options := InstanceProfileExistsWaiterOptions{}
	options.MinDelay = 1 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = instanceProfileExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InstanceProfileExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InstanceProfileExists waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *InstanceProfileExistsWaiter) Wait(ctx context.Context, params *GetInstanceProfileInput, maxWaitDur time.Duration, optFns ...func(*InstanceProfileExistsWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetInstanceProfile(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for InstanceProfileExists waiter")
}

func instanceProfileExistsStateRetryable(ctx context.Context, input *GetInstanceProfileInput, output *GetInstanceProfileOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var errorType *types.NoSuchEntityException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetInstanceProfile",
	}
}
