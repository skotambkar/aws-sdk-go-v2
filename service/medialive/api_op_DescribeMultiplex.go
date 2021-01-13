// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets details about a multiplex.
func (c *Client) DescribeMultiplex(ctx context.Context, params *DescribeMultiplexInput, optFns ...func(*Options)) (*DescribeMultiplexOutput, error) {
	if params == nil {
		params = &DescribeMultiplexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMultiplex", params, optFns, addOperationDescribeMultiplexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeMultiplexRequest
type DescribeMultiplexInput struct {

	// The ID of the multiplex.
	//
	// This member is required.
	MultiplexId *string
}

// Placeholder documentation for DescribeMultiplexResponse
type DescribeMultiplexOutput struct {

	// The unique arn of the multiplex.
	Arn *string

	// A list of availability zones for the multiplex.
	AvailabilityZones []string

	// A list of the multiplex output destinations.
	Destinations []types.MultiplexOutputDestination

	// The unique id of the multiplex.
	Id *string

	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings

	// The name of the multiplex.
	Name *string

	// The number of currently healthy pipelines.
	PipelinesRunningCount int32

	// The number of programs in the multiplex.
	ProgramCount int32

	// The current state of the multiplex.
	State types.MultiplexState

	// A collection of key-value pairs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMultiplexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMultiplex{}, middleware.After)
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
	if err = addOpDescribeMultiplexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMultiplex(options.Region), middleware.Before); err != nil {
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

// DescribeMultiplexAPIClient is a client that implements the DescribeMultiplex
// operation.
type DescribeMultiplexAPIClient interface {
	DescribeMultiplex(context.Context, *DescribeMultiplexInput, ...func(*Options)) (*DescribeMultiplexOutput, error)
}

var _ DescribeMultiplexAPIClient = (*Client)(nil)

// MultiplexCreatedWaiterOptions are waiter options for MultiplexCreatedWaiter
type MultiplexCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// MultiplexCreatedWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, MultiplexCreatedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexCreatedWaiter defines the waiters for MultiplexCreated
type MultiplexCreatedWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexCreatedWaiterOptions
}

// NewMultiplexCreatedWaiter constructs a MultiplexCreatedWaiter.
func NewMultiplexCreatedWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexCreatedWaiterOptions)) *MultiplexCreatedWaiter {
	options := MultiplexCreatedWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexCreated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexCreatedWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexCreatedWaiterOptions)) error {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for MultiplexCreated waiter")
}

func multiplexCreatedStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "IDLE"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATING"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// MultiplexDeletedWaiterOptions are waiter options for MultiplexDeletedWaiter
type MultiplexDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// MultiplexDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, MultiplexDeletedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexDeletedWaiter defines the waiters for MultiplexDeleted
type MultiplexDeletedWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexDeletedWaiterOptions
}

// NewMultiplexDeletedWaiter constructs a MultiplexDeletedWaiter.
func NewMultiplexDeletedWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexDeletedWaiterOptions)) *MultiplexDeletedWaiter {
	options := MultiplexDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexDeletedWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexDeletedWaiterOptions)) error {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for MultiplexDeleted waiter")
}

func multiplexDeletedStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETED"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETING"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// MultiplexRunningWaiterOptions are waiter options for MultiplexRunningWaiter
type MultiplexRunningWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// MultiplexRunningWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, MultiplexRunningWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexRunningWaiter defines the waiters for MultiplexRunning
type MultiplexRunningWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexRunningWaiterOptions
}

// NewMultiplexRunningWaiter constructs a MultiplexRunningWaiter.
func NewMultiplexRunningWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexRunningWaiterOptions)) *MultiplexRunningWaiter {
	options := MultiplexRunningWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexRunningStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexRunningWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexRunning waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexRunningWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexRunningWaiterOptions)) error {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for MultiplexRunning waiter")
}

func multiplexRunningStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STARTING"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// MultiplexStoppedWaiterOptions are waiter options for MultiplexStoppedWaiter
type MultiplexStoppedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// MultiplexStoppedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, MultiplexStoppedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexStoppedWaiter defines the waiters for MultiplexStopped
type MultiplexStoppedWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexStoppedWaiterOptions
}

// NewMultiplexStoppedWaiter constructs a MultiplexStoppedWaiter.
func NewMultiplexStoppedWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexStoppedWaiterOptions)) *MultiplexStoppedWaiter {
	options := MultiplexStoppedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexStopped waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexStoppedWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexStoppedWaiterOptions)) error {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for MultiplexStopped waiter")
}

func multiplexStoppedStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "IDLE"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STOPPING"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeMultiplex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeMultiplex",
	}
}
