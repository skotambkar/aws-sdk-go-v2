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

// Creates an Amazon Lex conversational bot.
func (c *Client) CreateBot(ctx context.Context, params *CreateBotInput, optFns ...func(*Options)) (*CreateBotOutput, error) {
	if params == nil {
		params = &CreateBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBot", params, optFns, c.addOperationCreateBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBotInput struct {

	// The name of the bot. The bot name must be unique in the account that creates the
	// bot.
	//
	// This member is required.
	BotName *string

	// Provides information on additional privacy protections Amazon Lex should use
	// with the bot's data.
	//
	// This member is required.
	DataPrivacy *types.DataPrivacy

	// The time, in seconds, that Amazon Lex should keep information about a user's
	// conversation with the bot. A user interaction remains active for the amount of
	// time specified. If no conversation occurs during this time, the session expires
	// and Amazon Lex deletes any data provided before the timeout. You can specify
	// between 60 (1 minute) and 86,400 (24 hours) seconds.
	//
	// This member is required.
	IdleSessionTTLInSeconds *int32

	// The Amazon Resource Name (ARN) of an IAM role that has permission to access the
	// bot.
	//
	// This member is required.
	RoleArn *string

	// A list of tags to add to the bot. You can only add tags when you create a bot.
	// You can't use the UpdateBot operation to update tags. To update tags, use the
	// TagResource operation.
	BotTags map[string]string

	// A description of the bot. It appears in lists to help you identify a particular
	// bot.
	Description *string

	// A list of tags to add to the test alias for a bot. You can only add tags when
	// you create a bot. You can't use the UpdateAlias operation to update tags. To
	// update tags on the test alias, use the TagResource operation.
	TestBotAliasTags map[string]string
}

type CreateBotOutput struct {

	// A unique identifier for a particular bot. You use this to identify the bot when
	// you call other Amazon Lex API operations.
	BotId *string

	// The name specified for the bot.
	BotName *string

	// Shows the current status of the bot. The bot is first in the Creating status.
	// Once the bot is read for use, it changes to the Available status. After the bot
	// is created, you can use the Draft version of the bot.
	BotStatus types.BotStatus

	// A list of tags associated with the bot.
	BotTags map[string]string

	// A timestamp indicating the date and time that the bot was created.
	CreationDateTime *time.Time

	// The data privacy settings specified for the bot.
	DataPrivacy *types.DataPrivacy

	// The description specified for the bot.
	Description *string

	// The session idle time specified for the bot.
	IdleSessionTTLInSeconds *int32

	// The IAM role specified for the bot.
	RoleArn *string

	// A list of tags associated with the test alias for the bot.
	TestBotAliasTags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBot{}, middleware.After)
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
	if err = addOpCreateBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateBot",
	}
}
