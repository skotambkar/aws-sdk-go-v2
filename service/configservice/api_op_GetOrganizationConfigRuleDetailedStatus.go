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

// Returns detailed status for each member account within an organization for a
// given organization config rule. Only a master account and a delegated
// administrator account can call this API. When calling this API with a delegated
// administrator, you must ensure AWS Organizations ListDelegatedAdministrator
// permissions are added.
func (c *Client) GetOrganizationConfigRuleDetailedStatus(ctx context.Context, params *GetOrganizationConfigRuleDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConfigRuleDetailedStatusOutput, error) {
	if params == nil {
		params = &GetOrganizationConfigRuleDetailedStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOrganizationConfigRuleDetailedStatus", params, optFns, addOperationGetOrganizationConfigRuleDetailedStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOrganizationConfigRuleDetailedStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOrganizationConfigRuleDetailedStatusInput struct {

	// The name of organization config rule for which you want status details for
	// member accounts.
	//
	// This member is required.
	OrganizationConfigRuleName *string

	// A StatusDetailFilters object.
	Filters *types.StatusDetailFilters

	// The maximum number of OrganizationConfigRuleDetailedStatus returned on each
	// page. If you do not specify a number, AWS Config uses the default. The default
	// is 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type GetOrganizationConfigRuleDetailedStatusOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of MemberAccountStatus objects.
	OrganizationConfigRuleDetailedStatus []types.MemberAccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOrganizationConfigRuleDetailedStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOrganizationConfigRuleDetailedStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOrganizationConfigRuleDetailedStatus{}, middleware.After)
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
	if err = addOpGetOrganizationConfigRuleDetailedStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOrganizationConfigRuleDetailedStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOrganizationConfigRuleDetailedStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetOrganizationConfigRuleDetailedStatus",
	}
}
