// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a VPC link.
func (c *Client) CreateVpcLink(ctx context.Context, params *CreateVpcLinkInput, optFns ...func(*Options)) (*CreateVpcLinkOutput, error) {
	if params == nil {
		params = &CreateVpcLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcLink", params, optFns, c.addOperationCreateVpcLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a VPC link
type CreateVpcLinkInput struct {

	// The name of the VPC link.
	//
	// This member is required.
	Name *string

	// A list of subnet IDs to include in the VPC link.
	//
	// This member is required.
	SubnetIds []string

	// A list of security group IDs for the VPC link.
	SecurityGroupIds []string

	// A list of tags.
	Tags map[string]string
}

type CreateVpcLinkOutput struct {

	// The timestamp when the VPC link was created.
	CreatedDate *time.Time

	// The name of the VPC link.
	Name *string

	// A list of security group IDs for the VPC link.
	SecurityGroupIds []string

	// A list of subnet IDs to include in the VPC link.
	SubnetIds []string

	// Tags for the VPC link.
	Tags map[string]string

	// The ID of the VPC link.
	VpcLinkId *string

	// The status of the VPC link.
	VpcLinkStatus types.VpcLinkStatus

	// A message summarizing the cause of the status of the VPC link.
	VpcLinkStatusMessage *string

	// The version of the VPC link.
	VpcLinkVersion types.VpcLinkVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateVpcLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVpcLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVpcLink{}, middleware.After)
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
	if err = addOpCreateVpcLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateVpcLink",
	}
}
