// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opCreateDiskSnapshot = "CreateDiskSnapshot"

// CreateDiskSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates a snapshot of a block storage disk. You can use snapshots for backups,
// to make copies of disks, and to save data before shutting down a Lightsail
// instance.
//
// You can take a snapshot of an attached disk that is in use; however, snapshots
// only capture data that has been written to your disk at the time the snapshot
// command is issued. This may exclude any data that has been cached by any
// applications or the operating system. If you can pause any file systems on
// the disk long enough to take a snapshot, your snapshot should be complete.
// Nevertheless, if you cannot pause all file writes to the disk, you should
// unmount the disk from within the Lightsail instance, issue the create disk
// snapshot command, and then remount the disk to ensure a consistent and complete
// snapshot. You may remount and use your disk while the snapshot status is
// pending.
//
// You can also use this operation to create a snapshot of an instance's system
// volume. You might want to do this, for example, to recover data from the
// system volume of a botched instance or to create a backup of the system volume
// like you would for a block storage disk. To create a snapshot of a system
// volume, just define the instance name parameter when issuing the snapshot
// command, and a snapshot of the defined instance's system volume will be created.
// After the snapshot is available, you can create a block storage disk from
// the snapshot and attach it to a running instance to access the data on the
// disk.
//
// The create disk snapshot operation supports tag-based access control via
// request tags. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateDiskSnapshotRequest.
//    req := client.CreateDiskSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateDiskSnapshot
func (c *Client) CreateDiskSnapshotRequest(input *types.CreateDiskSnapshotInput) CreateDiskSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateDiskSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDiskSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.CreateDiskSnapshotOutput{})
	return CreateDiskSnapshotRequest{Request: req, Input: input, Copy: c.CreateDiskSnapshotRequest}
}

// CreateDiskSnapshotRequest is the request type for the
// CreateDiskSnapshot API operation.
type CreateDiskSnapshotRequest struct {
	*aws.Request
	Input *types.CreateDiskSnapshotInput
	Copy  func(*types.CreateDiskSnapshotInput) CreateDiskSnapshotRequest
}

// Send marshals and sends the CreateDiskSnapshot API request.
func (r CreateDiskSnapshotRequest) Send(ctx context.Context) (*CreateDiskSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDiskSnapshotResponse{
		CreateDiskSnapshotOutput: r.Request.Data.(*types.CreateDiskSnapshotOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDiskSnapshotResponse is the response type for the
// CreateDiskSnapshot API operation.
type CreateDiskSnapshotResponse struct {
	*types.CreateDiskSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDiskSnapshot request.
func (r *CreateDiskSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
