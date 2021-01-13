// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates organization config rule for your entire organization evaluating
// whether your AWS resources comply with your desired configurations. Only a
// master account and a delegated administrator can create or update an
// organization config rule. When calling this API with a delegated administrator,
// you must ensure AWS Organizations ListDelegatedAdministrator permissions are
// added. This API enables organization service access through the
// EnableAWSServiceAccess action and creates a service linked role
// AWSServiceRoleForConfigMultiAccountSetup in the master or delegated
// administrator account of your organization. The service linked role is created
// only when the role does not exist in the caller account. AWS Config verifies the
// existence of role with GetRole action. To use this API with delegated
// administrator, register a delegated administrator by calling AWS Organization
// register-delegated-administrator for config-multiaccountsetup.amazonaws.com. You
// can use this action to create both custom AWS Config rules and AWS managed
// Config rules. If you are adding a new custom AWS Config rule, you must first
// create AWS Lambda function in the master account or a delegated administrator
// that the rule invokes to evaluate your resources. When you use the
// PutOrganizationConfigRule action to add the rule to AWS Config, you must specify
// the Amazon Resource Name (ARN) that AWS Lambda assigns to the function. If you
// are adding an AWS managed Config rule, specify the rule's identifier for the
// RuleIdentifier key. The maximum number of organization config rules that AWS
// Config supports is 150 and 3 delegated administrator per organization.
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in an
// organization. Specify either OrganizationCustomRuleMetadata or
// OrganizationManagedRuleMetadata.
func (c *Client) PutOrganizationConfigRule(ctx context.Context, params *PutOrganizationConfigRuleInput, optFns ...func(*Options)) (*PutOrganizationConfigRuleOutput, error) {
	if params == nil {
		params = &PutOrganizationConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutOrganizationConfigRule", params, optFns, addOperationPutOrganizationConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutOrganizationConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutOrganizationConfigRuleInput struct {

	// The name that you assign to an organization config rule.
	//
	// This member is required.
	OrganizationConfigRuleName *string

	// A comma-separated list of accounts that you want to exclude from an organization
	// config rule.
	ExcludedAccounts []string

	// An OrganizationCustomRuleMetadata object.
	OrganizationCustomRuleMetadata *types.OrganizationCustomRuleMetadata

	// An OrganizationManagedRuleMetadata object.
	OrganizationManagedRuleMetadata *types.OrganizationManagedRuleMetadata
}

type PutOrganizationConfigRuleOutput struct {

	// The Amazon Resource Name (ARN) of an organization config rule.
	OrganizationConfigRuleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutOrganizationConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutOrganizationConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutOrganizationConfigRule{}, middleware.After)
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
	if err = addOpPutOrganizationConfigRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutOrganizationConfigRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutOrganizationConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutOrganizationConfigRule",
	}
}
