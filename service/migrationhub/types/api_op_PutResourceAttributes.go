// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutResourceAttributesInput struct {
	_ struct{} `type:"structure"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// Unique identifier that references the migration task.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`

	// Information about the resource that is being migrated. This data will be
	// used to map the task to a resource in the Application Discovery Service (ADS)'s
	// repository.
	//
	// Takes the object array of ResourceAttribute where the Type field is reserved
	// for the following values: IPV4_ADDRESS | IPV6_ADDRESS | MAC_ADDRESS | FQDN
	// | VM_MANAGER_ID | VM_MANAGED_OBJECT_REFERENCE | VM_NAME | VM_PATH | BIOS_ID
	// | MOTHERBOARD_SERIAL_NUMBER where the identifying value can be a string up
	// to 256 characters.
	//
	//    * If any "VM" related value is set for a ResourceAttribute object, it
	//    is required that VM_MANAGER_ID, as a minimum, is always set. If VM_MANAGER_ID
	//    is not set, then all "VM" fields will be discarded and "VM" fields will
	//    not be used for matching the migration task to a server in Application
	//    Discovery Service (ADS)'s repository. See the Example (https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#API_PutResourceAttributes_Examples)
	//    section below for a use case of specifying "VM" related values.
	//
	//    * If a server you are trying to match has multiple IP or MAC addresses,
	//    you should provide as many as you know in separate type/value pairs passed
	//    to the ResourceAttributeList parameter to maximize the chances of matching.
	//
	// ResourceAttributeList is a required field
	ResourceAttributeList []ResourceAttribute `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutResourceAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResourceAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResourceAttributesInput"}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}

	if s.ResourceAttributeList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceAttributeList"))
	}
	if s.ResourceAttributeList != nil && len(s.ResourceAttributeList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceAttributeList", 1))
	}
	if s.ResourceAttributeList != nil {
		for i, v := range s.ResourceAttributeList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceAttributeList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutResourceAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutResourceAttributesOutput) String() string {
	return awsutil.Prettify(s)
}
