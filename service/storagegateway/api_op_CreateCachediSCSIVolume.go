// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opCreateCachediSCSIVolume = "CreateCachediSCSIVolume"

// CreateCachediSCSIVolumeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Creates a cached volume on a specified cached volume gateway. This operation
// is only supported in the cached volume gateway type.
//
// Cache storage must be allocated to the gateway before you can create a cached
// volume. Use the AddCache operation to add cache storage to a gateway.
//
// In the request, you must specify the gateway, size of the volume in bytes,
// the iSCSI target name, an IP address on which to expose the target, and a
// unique client token. In response, the gateway creates the volume and returns
// information about it. This information includes the volume Amazon Resource
// Name (ARN), its size, and the iSCSI target ARN that initiators can use to
// connect to the volume target.
//
// Optionally, you can provide the ARN for an existing volume as the SourceVolumeARN
// for this cached volume, which creates an exact copy of the existing volume’s
// latest recovery point. The VolumeSizeInBytes value must be equal to or larger
// than the size of the copied volume, in bytes.
//
//    // Example sending a request using CreateCachediSCSIVolumeRequest.
//    req := client.CreateCachediSCSIVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CreateCachediSCSIVolume
func (c *Client) CreateCachediSCSIVolumeRequest(input *types.CreateCachediSCSIVolumeInput) CreateCachediSCSIVolumeRequest {
	op := &aws.Operation{
		Name:       opCreateCachediSCSIVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateCachediSCSIVolumeInput{}
	}

	req := c.newRequest(op, input, &types.CreateCachediSCSIVolumeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateCachediSCSIVolumeMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateCachediSCSIVolumeRequest{Request: req, Input: input, Copy: c.CreateCachediSCSIVolumeRequest}
}

// CreateCachediSCSIVolumeRequest is the request type for the
// CreateCachediSCSIVolume API operation.
type CreateCachediSCSIVolumeRequest struct {
	*aws.Request
	Input *types.CreateCachediSCSIVolumeInput
	Copy  func(*types.CreateCachediSCSIVolumeInput) CreateCachediSCSIVolumeRequest
}

// Send marshals and sends the CreateCachediSCSIVolume API request.
func (r CreateCachediSCSIVolumeRequest) Send(ctx context.Context) (*CreateCachediSCSIVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCachediSCSIVolumeResponse{
		CreateCachediSCSIVolumeOutput: r.Request.Data.(*types.CreateCachediSCSIVolumeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCachediSCSIVolumeResponse is the response type for the
// CreateCachediSCSIVolume API operation.
type CreateCachediSCSIVolumeResponse struct {
	*types.CreateCachediSCSIVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCachediSCSIVolume request.
func (r *CreateCachediSCSIVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
