// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies one or more tags to the specified resources. Note the following:
//
// * Not
// all resources can have tags. For a list of services that support tagging, see
// this list
// (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
// *
// Each resource can have up to 50 tags. For other limits, see Tag Naming and Usage
// Conventions
// (http://docs.aws.amazon.com/general/latest/gr/aws_tagging.html#tag-conventions)
// in the AWS General Reference.
//
// * You can only tag resources that are located in
// the specified Region for the AWS account.
//
// * To add tags to a resource, you need
// the necessary permissions for the service that the resource belongs to as well
// as permissions for adding tags. For more information, see this list
// (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
// Do
// not store personally identifiable information (PII) or other confidential or
// sensitive information in tags. We use tags to provide you with billing and
// administration services. Tags are not intended to be used for private or
// sensitive data.
func (c *Client) TagResources(ctx context.Context, params *TagResourcesInput, optFns ...func(*Options)) (*TagResourcesOutput, error) {
	if params == nil {
		params = &TagResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagResources", params, optFns, addOperationTagResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourcesInput struct {

	// A list of ARNs. An ARN (Amazon Resource Name) uniquely identifies a resource.
	// For more information, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	ResourceARNList []string

	// The tags that you want to add to the specified resources. A tag consists of a
	// key and a value that you define.
	//
	// This member is required.
	Tags map[string]string
}

type TagResourcesOutput struct {

	// A map containing a key-value pair for each failed item that couldn't be tagged.
	// The key is the ARN of the failed resource. The value is a FailureInfo object
	// that contains an error code, a status code, and an error message. If there are
	// no errors, the FailedResourcesMap is empty.
	FailedResourcesMap map[string]types.FailureInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagResources{}, middleware.After)
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
	if err = addOpTagResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "TagResources",
	}
}
