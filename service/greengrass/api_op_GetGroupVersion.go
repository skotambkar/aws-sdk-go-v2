// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a group version.
func (c *Client) GetGroupVersion(ctx context.Context, params *GetGroupVersionInput, optFns ...func(*Options)) (*GetGroupVersionOutput, error) {
	if params == nil {
		params = &GetGroupVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupVersion", params, optFns, c.addOperationGetGroupVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupVersionInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string

	// The ID of the group version. This value maps to the ''Version'' property of the
	// corresponding ''VersionInformation'' object, which is returned by
	// ''ListGroupVersions'' requests. If the version is the last one that was
	// associated with a group, the value also maps to the ''LatestVersion'' property
	// of the corresponding ''GroupInformation'' object.
	//
	// This member is required.
	GroupVersionId *string
}

type GetGroupVersionOutput struct {

	// The ARN of the group version.
	Arn *string

	// The time, in milliseconds since the epoch, when the group version was created.
	CreationTimestamp *string

	// Information about the group version definition.
	Definition *types.GroupVersion

	// The ID of the group that the version is associated with.
	Id *string

	// The ID of the group version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetGroupVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGroupVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGroupVersion{}, middleware.After)
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
	if err = addOpGetGroupVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGroupVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetGroupVersion",
	}
}
