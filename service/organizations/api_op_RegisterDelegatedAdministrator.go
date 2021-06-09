// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the specified member account to administer the Organizations features of
// the specified AWS service. It grants read-only access to AWS Organizations
// service data. The account still requires IAM permissions to access and
// administer the AWS service. You can run this action only for AWS services that
// support this feature. For a current list of services that support it, see the
// column Supports Delegated Administrator in the table at AWS Services that you
// can use with AWS Organizations
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's management account.
func (c *Client) RegisterDelegatedAdministrator(ctx context.Context, params *RegisterDelegatedAdministratorInput, optFns ...func(*Options)) (*RegisterDelegatedAdministratorOutput, error) {
	if params == nil {
		params = &RegisterDelegatedAdministratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDelegatedAdministrator", params, optFns, c.addOperationRegisterDelegatedAdministratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDelegatedAdministratorInput struct {

	// The account ID number of the member account in the organization to register as a
	// delegated administrator.
	//
	// This member is required.
	AccountId *string

	// The service principal of the AWS service for which you want to make the member
	// account a delegated administrator.
	//
	// This member is required.
	ServicePrincipal *string
}

type RegisterDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationRegisterDelegatedAdministratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterDelegatedAdministrator{}, middleware.After)
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
	if err = addOpRegisterDelegatedAdministratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDelegatedAdministrator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterDelegatedAdministrator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "RegisterDelegatedAdministrator",
	}
}
