// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about all provisioned clusters if no cluster identifier is
// specified, or about a specific cache cluster if a cluster identifier is
// supplied. By default, abbreviated information about the clusters is returned.
// You can use the optional ShowCacheNodeInfo flag to retrieve detailed information
// about the cache nodes associated with the clusters. These details include the
// DNS address and port for the cache node endpoint. If the cluster is in the
// creating state, only cluster-level information is displayed until all of the
// nodes are successfully provisioned. If the cluster is in the deleting state,
// only cluster-level information is displayed. If cache nodes are currently being
// added to the cluster, node endpoint information and creation time for the
// additional nodes are not displayed until they are completely provisioned. When
// the cluster state is available, the cluster is ready for use. If cache nodes are
// currently being removed from the cluster, no endpoint information for the
// removed nodes is displayed.
func (c *Client) DescribeCacheClusters(ctx context.Context, params *DescribeCacheClustersInput, optFns ...func(*Options)) (*DescribeCacheClustersOutput, error) {
	if params == nil {
		params = &DescribeCacheClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheClusters", params, optFns, c.addOperationDescribeCacheClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheClusters operation.
type DescribeCacheClustersInput struct {

	// The user-supplied cluster identifier. If this parameter is specified, only
	// information about that specific cluster is returned. This parameter isn't case
	// sensitive.
	CacheClusterId *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	// An optional flag that can be included in the DescribeCacheCluster request to
	// show only nodes (API/CLI: clusters) that are not members of a replication group.
	// In practice, this mean Memcached and single node Redis clusters.
	ShowCacheClustersNotInReplicationGroups *bool

	// An optional flag that can be included in the DescribeCacheCluster request to
	// retrieve information about the individual cache nodes.
	ShowCacheNodeInfo *bool
}

// Represents the output of a DescribeCacheClusters operation.
type DescribeCacheClustersOutput struct {

	// A list of clusters. Each item in the list contains detailed information about
	// one cluster.
	CacheClusters []types.CacheCluster

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeCacheClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheClusters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheClusters(options.Region), middleware.Before); err != nil {
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

// DescribeCacheClustersAPIClient is a client that implements the
// DescribeCacheClusters operation.
type DescribeCacheClustersAPIClient interface {
	DescribeCacheClusters(context.Context, *DescribeCacheClustersInput, ...func(*Options)) (*DescribeCacheClustersOutput, error)
}

var _ DescribeCacheClustersAPIClient = (*Client)(nil)

// DescribeCacheClustersPaginatorOptions is the paginator options for
// DescribeCacheClusters
type DescribeCacheClustersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCacheClustersPaginator is a paginator for DescribeCacheClusters
type DescribeCacheClustersPaginator struct {
	options   DescribeCacheClustersPaginatorOptions
	client    DescribeCacheClustersAPIClient
	params    *DescribeCacheClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeCacheClustersPaginator returns a new DescribeCacheClustersPaginator
func NewDescribeCacheClustersPaginator(client DescribeCacheClustersAPIClient, params *DescribeCacheClustersInput, optFns ...func(*DescribeCacheClustersPaginatorOptions)) *DescribeCacheClustersPaginator {
	if params == nil {
		params = &DescribeCacheClustersInput{}
	}

	options := DescribeCacheClustersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCacheClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCacheClustersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeCacheClusters page.
func (p *DescribeCacheClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCacheClustersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeCacheClusters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// CacheClusterAvailableWaiterOptions are waiter options for
// CacheClusterAvailableWaiter
type CacheClusterAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// CacheClusterAvailableWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, CacheClusterAvailableWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeCacheClustersInput, *DescribeCacheClustersOutput, error) (bool, error)
}

// CacheClusterAvailableWaiter defines the waiters for CacheClusterAvailable
type CacheClusterAvailableWaiter struct {
	client DescribeCacheClustersAPIClient

	options CacheClusterAvailableWaiterOptions
}

// NewCacheClusterAvailableWaiter constructs a CacheClusterAvailableWaiter.
func NewCacheClusterAvailableWaiter(client DescribeCacheClustersAPIClient, optFns ...func(*CacheClusterAvailableWaiterOptions)) *CacheClusterAvailableWaiter {
	options := CacheClusterAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = cacheClusterAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CacheClusterAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CacheClusterAvailable waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *CacheClusterAvailableWaiter) Wait(ctx context.Context, params *DescribeCacheClustersInput, maxWaitDur time.Duration, optFns ...func(*CacheClusterAvailableWaiterOptions)) error {
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

		out, err := w.client.DescribeCacheClusters(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for CacheClusterAvailable waiter")
}

func cacheClusterAvailableStateRetryable(ctx context.Context, input *DescribeCacheClustersInput, output *DescribeCacheClustersOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("CacheClusters[].CacheClusterStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("CacheClusters[].CacheClusterStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("CacheClusters[].CacheClusterStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("CacheClusters[].CacheClusterStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-network"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("CacheClusters[].CacheClusterStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "restore-failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeCacheClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheClusters",
	}
}
