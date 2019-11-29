// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/enums"
)

// Request for the CreateSoftwareUpdateJob API.
type CreateSoftwareUpdateJobInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// The IAM Role that Greengrass will use to create pre-signed URLs pointing
	// towards the update artifact.
	//
	// S3UrlSignerRole is a required field
	S3UrlSignerRole *string `type:"string" required:"true"`

	// The piece of software on the Greengrass core that will be updated.
	//
	// SoftwareToUpdate is a required field
	SoftwareToUpdate enums.SoftwareToUpdate `type:"string" required:"true" enum:"true"`

	// The minimum level of log statements that should be logged by the OTA Agent
	// during an update.
	UpdateAgentLogLevel enums.UpdateAgentLogLevel `type:"string" enum:"true"`

	// The ARNs of the targets (IoT things or IoT thing groups) that this update
	// will be applied to.
	//
	// UpdateTargets is a required field
	UpdateTargets []string `type:"list" required:"true"`

	// The architecture of the cores which are the targets of an update.
	//
	// UpdateTargetsArchitecture is a required field
	UpdateTargetsArchitecture enums.UpdateTargetsArchitecture `type:"string" required:"true" enum:"true"`

	// The operating system of the cores which are the targets of an update.
	//
	// UpdateTargetsOperatingSystem is a required field
	UpdateTargetsOperatingSystem enums.UpdateTargetsOperatingSystem `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateSoftwareUpdateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSoftwareUpdateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSoftwareUpdateJobInput"}

	if s.S3UrlSignerRole == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3UrlSignerRole"))
	}
	if len(s.SoftwareToUpdate) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SoftwareToUpdate"))
	}

	if s.UpdateTargets == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdateTargets"))
	}
	if len(s.UpdateTargetsArchitecture) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("UpdateTargetsArchitecture"))
	}
	if len(s.UpdateTargetsOperatingSystem) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("UpdateTargetsOperatingSystem"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSoftwareUpdateJobOutput struct {
	_ struct{} `type:"structure"`

	// The IoT Job ARN corresponding to this update.
	IotJobArn *string `type:"string"`

	// The IoT Job Id corresponding to this update.
	IotJobId *string `type:"string"`

	// The software version installed on the device or devices after the update.
	PlatformSoftwareVersion *string `type:"string"`
}

// String returns the string representation
func (s CreateSoftwareUpdateJobOutput) String() string {
	return awsutil.Prettify(s)
}