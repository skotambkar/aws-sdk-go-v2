// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates room membership details, such as the member role, for a room in an
// Amazon Chime Enterprise account. The member role designates whether the member
// is a chat room administrator or a general chat room member. The member role can
// be updated only for user IDs.
func (c *Client) UpdateRoomMembership(ctx context.Context, params *UpdateRoomMembershipInput, optFns ...func(*Options)) (*UpdateRoomMembershipOutput, error) {
	if params == nil {
		params = &UpdateRoomMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoomMembership", params, optFns, c.addOperationUpdateRoomMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoomMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoomMembershipInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The member ID.
	//
	// This member is required.
	MemberId *string

	// The room ID.
	//
	// This member is required.
	RoomId *string

	// The role of the member.
	Role types.RoomMembershipRole
}

type UpdateRoomMembershipOutput struct {

	// The room membership details.
	RoomMembership *types.RoomMembership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateRoomMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoomMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoomMembership{}, middleware.After)
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
	if err = addOpUpdateRoomMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoomMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoomMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateRoomMembership",
	}
}
