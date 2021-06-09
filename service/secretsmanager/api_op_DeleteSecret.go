// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an entire secret and all of the versions. You can optionally include a
// recovery window during which you can restore the secret. If you don't specify a
// recovery window value, the operation defaults to 30 days. Secrets Manager
// attaches a DeletionDate stamp to the secret that specifies the end of the
// recovery window. At the end of the recovery window, Secrets Manager deletes the
// secret permanently. At any time before recovery window ends, you can use
// RestoreSecret to remove the DeletionDate and cancel the deletion of the secret.
// You cannot access the encrypted secret information in any secret scheduled for
// deletion. If you need to access that information, you must cancel the deletion
// with RestoreSecret and then retrieve the information.
//
// * There is no explicit
// operation to delete a version of a secret. Instead, remove all staging labels
// from the VersionStage field of a version. That marks the version as deprecated
// and allows Secrets Manager to delete it as needed. Versions without any staging
// labels do not show up in ListSecretVersionIds unless you specify
// IncludeDeprecated.
//
// * The permanent secret deletion at the end of the waiting
// period is performed as a background task with low priority. There is no
// guarantee of a specific time after the recovery window for the actual delete
// operation to occur.
//
// Minimum permissions To run this command, you must have the
// following permissions:
//
// * secretsmanager:DeleteSecret
//
// Related operations
//
// * To
// create a secret, use CreateSecret.
//
// * To cancel deletion of a version of a
// secret before the recovery window has expired, use RestoreSecret.
func (c *Client) DeleteSecret(ctx context.Context, params *DeleteSecretInput, optFns ...func(*Options)) (*DeleteSecretOutput, error) {
	if params == nil {
		params = &DeleteSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSecret", params, optFns, c.addOperationDeleteSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSecretInput struct {

	// Specifies the secret to delete. You can specify either the Amazon Resource Name
	// (ARN) or the friendly name of the secret. If you specify an ARN, we generally
	// recommend that you specify a complete ARN. You can specify a partial ARN too—for
	// example, if you don’t include the final hyphen and six random characters that
	// Secrets Manager adds at the end of the ARN when you created the secret. A
	// partial ARN match can work as long as it uniquely matches only one secret.
	// However, if your secret has a name that ends in a hyphen followed by six
	// characters (before Secrets Manager adds the hyphen and six characters to the
	// ARN) and you try to use that as a partial ARN, then those characters cause
	// Secrets Manager to assume that you’re specifying a complete ARN. This confusion
	// can cause unexpected results. To avoid this situation, we recommend that you
	// don’t create secret names ending with a hyphen followed by six characters. If
	// you specify an incomplete ARN without the random suffix, and instead provide the
	// 'friendly name', you must not include the random suffix. If you do include the
	// random suffix added by Secrets Manager, you receive either a
	// ResourceNotFoundException or an AccessDeniedException error, depending on your
	// permissions.
	//
	// This member is required.
	SecretId *string

	// (Optional) Specifies that the secret is to be deleted without any recovery
	// window. You can't use both this parameter and the RecoveryWindowInDays parameter
	// in the same API call. An asynchronous background process performs the actual
	// deletion, so there can be a short delay before the operation completes. If you
	// write code to delete and then immediately recreate a secret with the same name,
	// ensure that your code includes appropriate back off and retry logic. Use this
	// parameter with caution. This parameter causes the operation to skip the normal
	// waiting period before the permanent deletion that AWS would normally impose with
	// the RecoveryWindowInDays parameter. If you delete a secret with the
	// ForceDeleteWithouRecovery parameter, then you have no opportunity to recover the
	// secret. You lose the secret permanently. If you use this parameter and include a
	// previously deleted or nonexistent secret, the operation does not return the
	// error ResourceNotFoundException in order to correctly handle retries.
	ForceDeleteWithoutRecovery bool

	// (Optional) Specifies the number of days that Secrets Manager waits before
	// Secrets Manager can delete the secret. You can't use both this parameter and the
	// ForceDeleteWithoutRecovery parameter in the same API call. This value can range
	// from 7 to 30 days with a default value of 30.
	RecoveryWindowInDays int64
}

type DeleteSecretOutput struct {

	// The ARN of the secret that is now scheduled for deletion.
	ARN *string

	// The date and time after which this secret can be deleted by Secrets Manager and
	// can no longer be restored. This value is the date and time of the delete request
	// plus the number of days specified in RecoveryWindowInDays.
	DeletionDate *time.Time

	// The friendly name of the secret currently scheduled for deletion.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSecret{}, middleware.After)
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
	if err = addOpDeleteSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSecret(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "DeleteSecret",
	}
}
