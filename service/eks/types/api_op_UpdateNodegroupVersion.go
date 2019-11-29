// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateNodegroupVersionInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// The name of the Amazon EKS cluster that is associated with the managed node
	// group to update.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// Force the update if the existing node group's pods are unable to be drained
	// due to a pod disruption budget issue. If a previous update fails because
	// pods could not be drained, you can force the update after it fails to terminate
	// the old node regardless of whether or not any pods are running on the node.
	Force *bool `locationName:"force" type:"boolean"`

	// The name of the managed node group to update.
	//
	// NodegroupName is a required field
	NodegroupName *string `location:"uri" locationName:"nodegroupName" type:"string" required:"true"`

	// The AMI version of the Amazon EKS-optimized AMI to use for the update. By
	// default, the latest available AMI version for the node group's Kubernetes
	// version is used. For more information, see Amazon EKS-Optimized Linux AMI
	// Versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide.
	ReleaseVersion *string `locationName:"releaseVersion" type:"string"`

	// The Kubernetes version to update to. If no version is specified, then the
	// Kubernetes version of the node group does not change. You can specify the
	// Kubernetes version of the cluster to update the node group to the latest
	// AMI version of the cluster's Kubernetes version.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s UpdateNodegroupVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNodegroupVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNodegroupVersionInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodegroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodegroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNodegroupVersionOutput struct {
	_ struct{} `type:"structure"`

	// An object representing an asynchronous update.
	Update *Update `locationName:"update" type:"structure"`
}

// String returns the string representation
func (s UpdateNodegroupVersionOutput) String() string {
	return awsutil.Prettify(s)
}