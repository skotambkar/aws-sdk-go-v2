// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a virtual service within a service mesh. A virtual service is an
// abstraction of a real service that is provided by a virtual node directly or
// indirectly by means of a virtual router. Dependent services call your virtual
// service by its virtualServiceName, and those requests are routed to the virtual
// node or virtual router that is specified as the provider for the virtual
// service. For more information about virtual services, see Virtual services
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html).
func (c *Client) CreateVirtualService(ctx context.Context, params *CreateVirtualServiceInput, optFns ...func(*Options)) (*CreateVirtualServiceOutput, error) {
	if params == nil {
		params = &CreateVirtualServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualService", params, optFns, addOperationCreateVirtualServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateVirtualServiceInput struct {

	// The name of the service mesh to create the virtual service in.
	//
	// This member is required.
	MeshName *string

	// The virtual service specification to apply.
	//
	// This member is required.
	Spec *types.VirtualServiceSpec

	// The name to use for the virtual service.
	//
	// This member is required.
	VirtualServiceName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then the account that you specify must share the mesh with your account
	// before you can create the resource in the service mesh. For more information
	// about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// Optional metadata that you can apply to the virtual service to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef
}

//
type CreateVirtualServiceOutput struct {

	// The full description of your virtual service following the create call.
	//
	// This member is required.
	VirtualService *types.VirtualServiceData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVirtualServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVirtualService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVirtualService{}, middleware.After)
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
	if err = addOpCreateVirtualServiceValidationMiddleware(stack); err != nil {
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
