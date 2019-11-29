// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateScriptInput struct {
	_ struct{} `type:"structure"`

	// Descriptive label that is associated with a script. Script names do not need
	// to be unique. You can use UpdateScript to change this value later.
	Name *string `min:"1" type:"string"`

	// Location of the Amazon S3 bucket where a zipped file containing your Realtime
	// scripts is stored. The storage location must specify the Amazon S3 bucket
	// name, the zip file name (the "key"), and a role ARN that allows Amazon GameLift
	// to access the Amazon S3 storage location. The S3 bucket must be in the same
	// region where you want to create a new script. By default, Amazon GameLift
	// uploads the latest version of the zip file; if you have S3 object versioning
	// turned on, you can use the ObjectVersion parameter to specify an earlier
	// version.
	StorageLocation *S3Location `type:"structure"`

	// Version that is associated with a build or script. Version strings do not
	// need to be unique. You can use UpdateScript to change this value later.
	Version *string `min:"1" type:"string"`

	// Data object containing your Realtime scripts and dependencies as a zip file.
	// The zip file can have one or multiple files. Maximum size of a zip file is
	// 5 MB.
	//
	// When using the AWS CLI tool to create a script, this parameter is set to
	// the zip file name. It must be prepended with the string "fileb://" to indicate
	// that the file data is a binary object. For example: --zip-file fileb://myRealtimeScript.zip.
	//
	// ZipFile is automatically base64 encoded/decoded by the SDK.
	ZipFile []byte `type:"blob"`
}

// String returns the string representation
func (s CreateScriptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateScriptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateScriptInput"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.StorageLocation != nil {
		if err := s.StorageLocation.Validate(); err != nil {
			invalidParams.AddNested("StorageLocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateScriptOutput struct {
	_ struct{} `type:"structure"`

	// The newly created script record with a unique script ID. The new script's
	// storage location reflects an Amazon S3 location: (1) If the script was uploaded
	// from an S3 bucket under your account, the storage location reflects the information
	// that was provided in the CreateScript request; (2) If the script file was
	// uploaded from a local zip file, the storage location reflects an S3 location
	// controls by the Amazon GameLift service.
	Script *Script `type:"structure"`
}

// String returns the string representation
func (s CreateScriptOutput) String() string {
	return awsutil.Prettify(s)
}