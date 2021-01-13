// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a component. Components are software that run on AWS IoT Greengrass core
// devices. After you develop and test a component on your core device, you can use
// this operation to upload your component to AWS IoT Greengrass. Then, you can
// deploy the component to other core devices. You can use this operation to do the
// following:
//
// * Create components from recipes Create a component from a recipe,
// which is a file that defines the component's metadata, parameters, dependencies,
// lifecycle, artifacts, and platform capability. For more information, see AWS IoT
// Greengrass component recipe reference
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/component-recipe-reference.html)
// in the AWS IoT Greengrass V2 Developer Guide. To create a component from a
// recipe, specify inlineRecipe when you call this operation.
//
// * Create components
// from Lambda functions Create a component from an AWS Lambda function that runs
// on AWS IoT Greengrass. This creates a recipe and artifacts from the Lambda
// function's deployment package. You can use this operation to migrate Lambda
// functions from AWS IoT Greengrass V1 to AWS IoT Greengrass V2. This function
// only accepts Lambda functions that use the following runtimes:
//
// * Python 2.7 –
// python2.7
//
// * Python 3.7 – python3.7
//
// * Python 3.8 – python3.8
//
// * Java 8 –
// java8
//
// * Node.js 10 – nodejs10.x
//
// * Node.js 12 – nodejs12.x
//
// To create a
// component from a Lambda function, specify lambdaFunction when you call this
// operation.
func (c *Client) CreateComponentVersion(ctx context.Context, params *CreateComponentVersionInput, optFns ...func(*Options)) (*CreateComponentVersionOutput, error) {
	if params == nil {
		params = &CreateComponentVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComponentVersion", params, optFns, addOperationCreateComponentVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComponentVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentVersionInput struct {

	// The recipe to use to create the component. The recipe defines the component's
	// metadata, parameters, dependencies, lifecycle, artifacts, and platform
	// compatibility. You must specify either inlineRecipe or lambdaFunction.
	InlineRecipe []byte

	// The parameters to create a component from a Lambda function. You must specify
	// either inlineRecipe or lambdaFunction.
	LambdaFunction *types.LambdaFunctionRecipeSource

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see Tag your resources
	// (https://docs.aws.amazon.com/greengrass/v2/tag-resources.html) in the AWS IoT
	// Greengrass V2 Developer Guide.
	Tags map[string]string
}

type CreateComponentVersionOutput struct {

	// The name of the component.
	//
	// This member is required.
	ComponentName *string

	// The version of the component.
	//
	// This member is required.
	ComponentVersion *string

	// The time at which the component was created, expressed in ISO 8601 format.
	//
	// This member is required.
	CreationTimestamp *time.Time

	// The status of the component version in AWS IoT Greengrass V2. This status is
	// different from the status of the component on a core device.
	//
	// This member is required.
	Status *types.CloudComponentStatus

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the component version.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateComponentVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComponentVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComponentVersion{}, middleware.After)
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
	if err = addOpCreateComponentVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponentVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateComponentVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateComponentVersion",
	}
}
