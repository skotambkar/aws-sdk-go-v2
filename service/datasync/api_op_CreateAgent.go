// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates an AWS DataSync agent that you have deployed on your host. The
// activation process associates your agent with your account. In the activation
// process, you specify information such as the AWS Region that you want to
// activate the agent in. You activate the agent in the AWS Region where your
// target locations (in Amazon S3 or Amazon EFS) reside. Your tasks are created in
// this AWS Region. You can activate the agent in a VPC (virtual private cloud) or
// provide the agent access to a VPC endpoint so you can run tasks without going
// over the public internet. You can use an agent for more than one location. If a
// task uses multiple agents, all of them need to have status AVAILABLE for the
// task to run. If you use multiple agents for a source location, the status of all
// the agents must be AVAILABLE for the task to run. Agents are automatically
// updated by AWS on a regular basis, using a mechanism that ensures minimal
// interruption to your tasks.
func (c *Client) CreateAgent(ctx context.Context, params *CreateAgentInput, optFns ...func(*Options)) (*CreateAgentOutput, error) {
	if params == nil {
		params = &CreateAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAgent", params, optFns, c.addOperationCreateAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateAgentRequest
type CreateAgentInput struct {

	// Your agent activation key. You can get the activation key either by sending an
	// HTTP GET request with redirects that enable you to get the agent IP address
	// (port 80). Alternatively, you can get it from the AWS DataSync console. The
	// redirect URL returned in the response provides you the activation key for your
	// agent in the query string parameter activationKey. It might also include other
	// activation-related parameters; however, these are merely defaults. The arguments
	// you pass to this API call determine the actual configuration of your agent. For
	// more information, see Activating an Agent in the AWS DataSync User Guide.
	//
	// This member is required.
	ActivationKey *string

	// The name you configured for your agent. This value is a text reference that is
	// used to identify the agent in the console.
	AgentName *string

	// The ARNs of the security groups used to protect your data transfer task subnets.
	// See CreateAgentRequest$SubnetArns.
	SecurityGroupArns []string

	// The Amazon Resource Names (ARNs) of the subnets in which DataSync will create
	// elastic network interfaces for each data transfer task. The agent that runs a
	// task must be private. When you start a task that is associated with an agent
	// created in a VPC, or one that has access to an IP address in a VPC, then the
	// task is also private. In this case, DataSync creates four network interfaces for
	// each task in your subnet. For a data transfer to work, the agent must be able to
	// route to all these four network interfaces.
	SubnetArns []string

	// The key-value pair that represents the tag that you want to associate with the
	// agent. The value can be an empty string. This value helps you manage, filter,
	// and search for your agents. Valid characters for key and value are letters,
	// spaces, and numbers representable in UTF-8 format, and the following special
	// characters: + - = . _ : / @.
	Tags []types.TagListEntry

	// The ID of the VPC (virtual private cloud) endpoint that the agent has access to.
	// This is the client-side VPC endpoint, also called a PrivateLink. If you don't
	// have a PrivateLink VPC endpoint, see Creating a VPC Endpoint Service
	// Configuration
	// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html#create-endpoint-service)
	// in the Amazon VPC User Guide. VPC endpoint ID looks like this:
	// vpce-01234d5aff67890e1.
	VpcEndpointId *string
}

// CreateAgentResponse
type CreateAgentOutput struct {

	// The Amazon Resource Name (ARN) of the agent. Use the ListAgents operation to
	// return a list of agents for your account and AWS Region.
	AgentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAgent{}, middleware.After)
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
	if err = addOpCreateAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateAgent",
	}
}
