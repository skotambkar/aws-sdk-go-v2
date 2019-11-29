// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type CreateNotebookInstanceInput struct {
	_ struct{} `type:"structure"`

	// A list of Elastic Inference (EI) instance types to associate with this notebook
	// instance. Currently, only one instance type can be associated with a notebook
	// instance. For more information, see Using Elastic Inference in Amazon SageMaker
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html).
	AcceleratorTypes []enums.NotebookInstanceAcceleratorType `type:"list"`

	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in
	// your account, or the URL of Git repositories in AWS CodeCommit (codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. These repositories are cloned at the same
	// level as the default repository of your notebook instance. For more information,
	// see Associating Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	AdditionalCodeRepositories []string `type:"list"`

	// A Git repository to associate with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in AWS CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. When you open a notebook instance, it opens
	// in the directory that contains this repository. For more information, see
	// Associating Git Repositories with Amazon SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	DefaultCodeRepository *string `min:"1" type:"string"`

	// Sets whether Amazon SageMaker provides internet access to the notebook instance.
	// If you set this to Disabled this notebook instance will be able to access
	// resources only in your VPC, and will not be able to connect to Amazon SageMaker
	// training and endpoint services unless your configure a NAT Gateway in your
	// VPC.
	//
	// For more information, see Notebook Instances Are Internet-Enabled by Default
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access).
	// You can set the value of this parameter to Disabled only if you set a value
	// for the SubnetId parameter.
	DirectInternetAccess enums.DirectInternetAccess `type:"string" enum:"true"`

	// The type of ML compute instance to launch for the notebook instance.
	//
	// InstanceType is a required field
	InstanceType enums.InstanceType `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon
	// SageMaker uses to encrypt data on the storage volume attached to your notebook
	// instance. The KMS key you provide must be enabled. For information, see Enabling
	// and Disabling Keys (https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html)
	// in the AWS Key Management Service Developer Guide.
	KmsKeyId *string `type:"string"`

	// The name of a lifecycle configuration to associate with the notebook instance.
	// For information about lifestyle configurations, see Step 2.1: (Optional)
	// Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
	LifecycleConfigName *string `type:"string"`

	// The name of the new notebook instance.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`

	// When you send any requests to AWS resources from the notebook instance, Amazon
	// SageMaker assumes this role to perform tasks on your behalf. You must grant
	// this role necessary permissions so Amazon SageMaker can perform these tasks.
	// The policy must allow the Amazon SageMaker service principal (sagemaker.amazonaws.com)
	// permissionsto to assume this role. For more information, see Amazon SageMaker
	// Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).
	//
	// To be able to pass this role to Amazon SageMaker, the caller of this API
	// must have the iam:PassRole permission.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// Whether root access is enabled or disabled for users of the notebook instance.
	// The default value is Enabled.
	//
	// Lifecycle configurations need root access to be able to set up a notebook
	// instance. Because of this, lifecycle configurations associated with a notebook
	// instance always run with root access even if you disable root access for
	// users.
	RootAccess enums.RootAccess `type:"string" enum:"true"`

	// The VPC security group IDs, in the form sg-xxxxxxxx. The security groups
	// must be for the same VPC as specified in the subnet.
	SecurityGroupIds []string `type:"list"`

	// The ID of the subnet in a VPC to which you would like to have a connectivity
	// from your ML compute instance.
	SubnetId *string `type:"string"`

	// A list of tags to associate with the notebook instance. You can add tags
	// later by using the CreateTags API.
	Tags []Tag `type:"list"`

	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	// The default value is 5 GB.
	VolumeSizeInGB *int64 `min:"5" type:"integer"`
}

// String returns the string representation
func (s CreateNotebookInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNotebookInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNotebookInstanceInput"}
	if s.DefaultCodeRepository != nil && len(*s.DefaultCodeRepository) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultCodeRepository", 1))
	}
	if len(s.InstanceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("InstanceType"))
	}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.VolumeSizeInGB != nil && *s.VolumeSizeInGB < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("VolumeSizeInGB", 5))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNotebookInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the notebook instance.
	NotebookInstanceArn *string `type:"string"`
}

// String returns the string representation
func (s CreateNotebookInstanceOutput) String() string {
	return awsutil.Prettify(s)
}
