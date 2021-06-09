// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves detailed information about the health of instances in your AWS Elastic
// Beanstalk. This operation requires enhanced health reporting
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced.html).
func (c *Client) DescribeInstancesHealth(ctx context.Context, params *DescribeInstancesHealthInput, optFns ...func(*Options)) (*DescribeInstancesHealthOutput, error) {
	if params == nil {
		params = &DescribeInstancesHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstancesHealth", params, optFns, c.addOperationDescribeInstancesHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancesHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Parameters for a call to DescribeInstancesHealth.
type DescribeInstancesHealthInput struct {

	// Specifies the response elements you wish to receive. To retrieve all attributes,
	// set to All. If no attribute names are specified, returns a list of instances.
	AttributeNames []types.InstancesHealthAttribute

	// Specify the AWS Elastic Beanstalk environment by ID.
	EnvironmentId *string

	// Specify the AWS Elastic Beanstalk environment by name.
	EnvironmentName *string

	// Specify the pagination token returned by a previous call.
	NextToken *string
}

// Detailed health information about the Amazon EC2 instances in an AWS Elastic
// Beanstalk environment.
type DescribeInstancesHealthOutput struct {

	// Detailed health information about each instance. The output differs slightly
	// between Linux and Windows environments. There is a difference in the members
	// that are supported under the  type.
	InstanceHealthList []types.SingleInstanceHealth

	// Pagination token for the next page of results, if available.
	NextToken *string

	// The date and time that the health information was retrieved.
	RefreshedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeInstancesHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstancesHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstancesHealth{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstancesHealth(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstancesHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeInstancesHealth",
	}
}
