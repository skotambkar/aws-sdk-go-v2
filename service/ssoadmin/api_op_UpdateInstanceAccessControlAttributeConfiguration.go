// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the AWS SSO identity store attributes to use with the AWS SSO instance
// for attributes-based access control (ABAC). When using an external identity
// provider as an identity source, you can pass attributes through the SAML
// assertion as an alternative to configuring attributes from the AWS SSO identity
// store. If a SAML assertion passes any of these attributes, AWS SSO will replace
// the attribute value with the value from the AWS SSO identity store. For more
// information about ABAC, see Attribute-Based Access Control in the AWS SSO User
// Guide.
func (c *Client) UpdateInstanceAccessControlAttributeConfiguration(ctx context.Context, params *UpdateInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*UpdateInstanceAccessControlAttributeConfigurationOutput, error) {
	if params == nil {
		params = &UpdateInstanceAccessControlAttributeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInstanceAccessControlAttributeConfiguration", params, optFns, addOperationUpdateInstanceAccessControlAttributeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInstanceAccessControlAttributeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInstanceAccessControlAttributeConfigurationInput struct {

	// Updates the attributes for your ABAC configuration.
	//
	// This member is required.
	InstanceAccessControlAttributeConfiguration *types.InstanceAccessControlAttributeConfiguration

	// The ARN of the SSO instance under which the operation will be executed.
	//
	// This member is required.
	InstanceArn *string
}

type UpdateInstanceAccessControlAttributeConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInstanceAccessControlAttributeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInstanceAccessControlAttributeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInstanceAccessControlAttributeConfiguration{}, middleware.After)
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
	if err = addOpUpdateInstanceAccessControlAttributeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInstanceAccessControlAttributeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInstanceAccessControlAttributeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "UpdateInstanceAccessControlAttributeConfiguration",
	}
}
