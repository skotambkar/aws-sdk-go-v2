// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the configuration of an existing bot alias.
func (c *Client) UpdateBotAlias(ctx context.Context, params *UpdateBotAliasInput, optFns ...func(*Options)) (*UpdateBotAliasOutput, error) {
	if params == nil {
		params = &UpdateBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBotAlias", params, optFns, c.addOperationUpdateBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBotAliasInput struct {

	// The unique identifier of the bot alias.
	//
	// This member is required.
	BotAliasId *string

	// The new name to assign to the bot alias.
	//
	// This member is required.
	BotAliasName *string

	// The identifier of the bot with the updated alias.
	//
	// This member is required.
	BotId *string

	// The new Lambda functions to use in each locale for the bot alias.
	BotAliasLocaleSettings map[string]types.BotAliasLocaleSettings

	// The new bot version to assign to the bot alias.
	BotVersion *string

	// The new settings for storing conversation logs in Amazon CloudWatch Logs and
	// Amazon S3 buckets.
	ConversationLogSettings *types.ConversationLogSettings

	// The new description to assign to the bot alias.
	Description *string

	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment
	// of user utterances.
	SentimentAnalysisSettings *types.SentimentAnalysisSettings
}

type UpdateBotAliasOutput struct {

	// The identifier of the updated bot alias.
	BotAliasId *string

	// The updated Lambda functions to use in each locale for the bot alias.
	BotAliasLocaleSettings map[string]types.BotAliasLocaleSettings

	// The updated name of the bot alias.
	BotAliasName *string

	// The current status of the bot alias. When the status is Available the alias is
	// ready for use.
	BotAliasStatus types.BotAliasStatus

	// The identifier of the bot with the updated alias.
	BotId *string

	// The updated version of the bot that the alias points to.
	BotVersion *string

	// The updated settings for storing conversation logs in Amazon CloudWatch Logs and
	// Amazon S3 buckets.
	ConversationLogSettings *types.ConversationLogSettings

	// A timestamp of the date and time that the bot was created.
	CreationDateTime *time.Time

	// The updated description of the bot alias.
	Description *string

	// A timestamp of the date and time that the bot was last updated.
	LastUpdatedDateTime *time.Time

	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment
	// of user utterances.
	SentimentAnalysisSettings *types.SentimentAnalysisSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBotAlias{}, middleware.After)
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
	if err = addOpUpdateBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBotAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "UpdateBotAlias",
	}
}
