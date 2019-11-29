// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opRegisterImage = "RegisterImage"

// RegisterImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Registers an AMI. When you're creating an AMI, this is the final step you
// must complete before you can launch an instance from the AMI. For more information
// about creating AMIs, see Creating Your Own AMIs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// For Amazon EBS-backed instances, CreateImage creates and registers the AMI
// in a single request, so you don't have to register the AMI yourself.
//
// You can also use RegisterImage to create an Amazon EBS-backed Linux AMI from
// a snapshot of a root device volume. You specify the snapshot using the block
// device mapping. For more information, see Launching a Linux Instance from
// a Backup (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-launch-snapshot.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// You can't register an image where a secondary (non-root) snapshot has AWS
// Marketplace product codes.
//
// Windows and some Linux distributions, such as Red Hat Enterprise Linux (RHEL)
// and SUSE Linux Enterprise Server (SLES), use the EC2 billing product code
// associated with an AMI to verify the subscription status for package updates.
// To create a new AMI for operating systems that require a billing product
// code, do the following:
//
// Launch an instance from an existing AMI with that billing product code.
//
// Customize the instance.
//
// Create a new AMI from the instance using CreateImage to preserve the billing
// product code association.
//
// If you purchase a Reserved Instance to apply to an On-Demand Instance that
// was launched from an AMI with a billing product code, make sure that the
// Reserved Instance has the matching billing product code. If you purchase
// a Reserved Instance without the matching billing product code, the Reserved
// Instance will not be applied to the On-Demand Instance.
//
// If needed, you can deregister an AMI at any time. Any modifications you make
// to an AMI backed by an instance store volume invalidates its registration.
// If you make changes to an image, deregister the previous image and register
// the new image.
//
//    // Example sending a request using RegisterImageRequest.
//    req := client.RegisterImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RegisterImage
func (c *Client) RegisterImageRequest(input *types.RegisterImageInput) RegisterImageRequest {
	op := &aws.Operation{
		Name:       opRegisterImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterImageInput{}
	}

	req := c.newRequest(op, input, &types.RegisterImageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.RegisterImageMarshaler{Input: input}.GetNamedBuildHandler())

	return RegisterImageRequest{Request: req, Input: input, Copy: c.RegisterImageRequest}
}

// RegisterImageRequest is the request type for the
// RegisterImage API operation.
type RegisterImageRequest struct {
	*aws.Request
	Input *types.RegisterImageInput
	Copy  func(*types.RegisterImageInput) RegisterImageRequest
}

// Send marshals and sends the RegisterImage API request.
func (r RegisterImageRequest) Send(ctx context.Context) (*RegisterImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterImageResponse{
		RegisterImageOutput: r.Request.Data.(*types.RegisterImageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterImageResponse is the response type for the
// RegisterImage API operation.
type RegisterImageResponse struct {
	*types.RegisterImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterImage request.
func (r *RegisterImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
