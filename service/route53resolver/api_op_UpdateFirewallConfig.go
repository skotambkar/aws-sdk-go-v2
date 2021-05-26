// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of the firewall behavior provided by DNS Firewall for
// a single Amazon virtual private cloud (VPC).
func (c *Client) UpdateFirewallConfig(ctx context.Context, params *UpdateFirewallConfigInput, optFns ...func(*Options)) (*UpdateFirewallConfigOutput, error) {
	if params == nil {
		params = &UpdateFirewallConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallConfig", params, optFns, c.addOperationUpdateFirewallConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallConfigInput struct {

	// Determines how Route 53 Resolver handles queries during failures, for example
	// when all traffic that is sent to DNS Firewall fails to receive a reply.
	//
	// * By
	// default, fail open is disabled, which means the failure mode is closed. This
	// approach favors security over availability. DNS Firewall blocks queries that it
	// is unable to evaluate properly.
	//
	// * If you enable this option, the failure mode
	// is open. This approach favors availability over security. DNS Firewall allows
	// queries to proceed if it is unable to properly evaluate them.
	//
	// This behavior is
	// only enforced for VPCs that have at least one DNS Firewall rule group
	// association.
	//
	// This member is required.
	FirewallFailOpen types.FirewallFailOpenStatus

	// The ID of the Amazon virtual private cloud (VPC) that the configuration is for.
	//
	// This member is required.
	ResourceId *string
}

type UpdateFirewallConfigOutput struct {

	// Configuration of the firewall behavior provided by DNS Firewall for a single
	// Amazon virtual private cloud (VPC).
	FirewallConfig *types.FirewallConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateFirewallConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFirewallConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFirewallConfig{}, middleware.After)
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
	if err = addOpUpdateFirewallConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "UpdateFirewallConfig",
	}
}
