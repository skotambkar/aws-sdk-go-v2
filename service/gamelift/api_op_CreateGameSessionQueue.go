// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a placement queue that processes requests for new game sessions. A queue
// uses FleetIQ algorithms to determine the best placement locations and find an
// available game server there, then prompts the game server process to start a new
// game session. A game session queue is configured with a set of destinations
// (GameLift fleets or aliases), which determine the locations where the queue can
// place new game sessions. These destinations can span multiple fleet types (Spot
// and On-Demand), instance types, and AWS Regions. If the queue includes
// multi-location fleets, the queue is able to place game sessions in all of a
// fleet's remote locations. You can opt to filter out individual locations if
// needed. The queue configuration also determines how FleetIQ selects the best
// available placement for a new game session. Before searching for an available
// game server, FleetIQ first prioritizes the queue's destinations and locations,
// with the best placement locations on top. You can set up the queue to use the
// FleetIQ default prioritization or provide an alternate set of priorities. To
// create a new queue, provide a name, timeout value, and a list of destinations.
// Optionally, specify a sort configuration and/or a filter, and define a set of
// latency cap policies. You can also include the ARN for an Amazon Simple
// Notification Service (SNS) topic to receive notifications of game session
// placement activity. Notifications using SNS or CloudWatch events is the
// preferred way to track placement activity. If successful, a new GameSessionQueue
// object is returned with an assigned queue ARN. New game session requests, which
// are submitted to the queue with StartGameSessionPlacement or StartMatchmaking,
// reference a queue's name or ARN. Learn more  Design a game session queue
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-design.html)
// Create a game session queue
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-creating.html)
// Related actions CreateGameSessionQueue | DescribeGameSessionQueues |
// UpdateGameSessionQueue | DeleteGameSessionQueue | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateGameSessionQueue(ctx context.Context, params *CreateGameSessionQueueInput, optFns ...func(*Options)) (*CreateGameSessionQueueOutput, error) {
	if params == nil {
		params = &CreateGameSessionQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGameSessionQueue", params, optFns, c.addOperationCreateGameSessionQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGameSessionQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type CreateGameSessionQueueInput struct {

	// A descriptive label that is associated with game session queue. Queue names must
	// be unique within each Region.
	//
	// This member is required.
	Name *string

	// Information to be added to all events that are related to this game session
	// queue.
	CustomEventData *string

	// A list of fleets and/or fleet aliases that can be used to fulfill game session
	// placement requests in the queue. Destinations are identified by either a fleet
	// ARN or a fleet alias ARN, and are listed in order of placement preference.
	Destinations []types.GameSessionQueueDestination

	// A list of locations where a queue is allowed to place new game sessions.
	// Locations are specified in the form of AWS Region codes, such as us-west-2. If
	// this parameter is not set, game sessions can be placed in any queue location.
	FilterConfiguration *types.FilterConfiguration

	// An SNS topic ARN that is set up to receive game session placement notifications.
	// See  Setting up notifications for game session placement
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queue-notification.html).
	NotificationTarget *string

	// A set of policies that act as a sliding cap on player latency. FleetIQ works to
	// deliver low latency for most players in a game session. These policies ensure
	// that no individual player can be placed into a game with unreasonably high
	// latency. Use multiple policies to gradually relax latency requirements a step at
	// a time. Multiple policies are applied based on their maximum allowed latency,
	// starting with the lowest value.
	PlayerLatencyPolicies []types.PlayerLatencyPolicy

	// Custom settings to use when prioritizing destinations and locations for game
	// session placements. This configuration replaces the FleetIQ default
	// prioritization process. Priority types that are not explicitly named will be
	// automatically applied at the end of the prioritization process.
	PriorityConfiguration *types.PriorityConfiguration

	// A list of labels to assign to the new game session queue resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources are useful for resource
	// management, access management and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []types.Tag

	// The maximum time, in seconds, that a new game session placement request remains
	// in the queue. When a request exceeds this time, the game session placement
	// changes to a TIMED_OUT status.
	TimeoutInSeconds *int32
}

// Represents the returned data in response to a request operation.
type CreateGameSessionQueueOutput struct {

	// An object that describes the newly created game session queue.
	GameSessionQueue *types.GameSessionQueue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateGameSessionQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGameSessionQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGameSessionQueue{}, middleware.After)
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
	if err = addOpCreateGameSessionQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGameSessionQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGameSessionQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateGameSessionQueue",
	}
}
