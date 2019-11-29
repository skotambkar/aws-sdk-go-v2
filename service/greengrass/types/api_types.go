// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/enums"
)

// Information about a bulk deployment. You cannot start a new bulk deployment
// while another one is still running or in a non-terminal state.
type BulkDeployment struct {
	_ struct{} `type:"structure"`

	// The ARN of the bulk deployment.
	BulkDeploymentArn *string `type:"string"`

	// The ID of the bulk deployment.
	BulkDeploymentId *string `type:"string"`

	// The time, in ISO format, when the deployment was created.
	CreatedAt *string `type:"string"`
}

// String returns the string representation
func (s BulkDeployment) String() string {
	return awsutil.Prettify(s)
}

// Relevant metrics on input records processed during bulk deployment.
type BulkDeploymentMetrics struct {
	_ struct{} `type:"structure"`

	// The total number of records that returned a non-retryable error. For example,
	// this can occur if a group record from the input file uses an invalid format
	// or specifies a nonexistent group version, or if the execution role doesn't
	// grant permission to deploy a group or group version.
	InvalidInputRecords *int64 `type:"integer"`

	// The total number of group records from the input file that have been processed
	// so far, or attempted.
	RecordsProcessed *int64 `type:"integer"`

	// The total number of deployment attempts that returned a retryable error.
	// For example, a retry is triggered if the attempt to deploy a group returns
	// a throttling error. ''StartBulkDeployment'' retries a group deployment up
	// to five times.
	RetryAttempts *int64 `type:"integer"`
}

// String returns the string representation
func (s BulkDeploymentMetrics) String() string {
	return awsutil.Prettify(s)
}

// Information about an individual group deployment in a bulk deployment operation.
type BulkDeploymentResult struct {
	_ struct{} `type:"structure"`

	// The time, in ISO format, when the deployment was created.
	CreatedAt *string `type:"string"`

	// The ARN of the group deployment.
	DeploymentArn *string `type:"string"`

	// The ID of the group deployment.
	DeploymentId *string `type:"string"`

	// The current status of the group deployment: ''InProgress'', ''Building'',
	// ''Success'', or ''Failure''.
	DeploymentStatus *string `type:"string"`

	// The type of the deployment.
	DeploymentType enums.DeploymentType `type:"string" enum:"true"`

	// Details about the error.
	ErrorDetails []ErrorDetail `type:"list"`

	// The error message for a failed deployment
	ErrorMessage *string `type:"string"`

	// The ARN of the Greengrass group.
	GroupArn *string `type:"string"`
}

// String returns the string representation
func (s BulkDeploymentResult) String() string {
	return awsutil.Prettify(s)
}

// Information about a Greengrass core's connectivity.
type ConnectivityInfo struct {
	_ struct{} `type:"structure"`

	// The endpoint for the Greengrass core. Can be an IP address or DNS.
	HostAddress *string `type:"string"`

	// The ID of the connectivity information.
	Id *string `type:"string"`

	// Metadata for this endpoint.
	Metadata *string `type:"string"`

	// The port of the Greengrass core. Usually 8883.
	PortNumber *int64 `type:"integer"`
}

// String returns the string representation
func (s ConnectivityInfo) String() string {
	return awsutil.Prettify(s)
}

// Information about a connector. Connectors run on the Greengrass core and
// contain built-in integration with local infrastructure, device protocols,
// AWS, and other cloud services.
type Connector struct {
	_ struct{} `type:"structure"`

	// The ARN of the connector.
	//
	// ConnectorArn is a required field
	ConnectorArn *string `type:"string" required:"true"`

	// A descriptive or arbitrary ID for the connector. This value must be unique
	// within the connector definition version. Max length is 128 characters with
	// pattern [a-zA-Z0-9:_-]+.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// The parameters or configuration that the connector uses.
	Parameters map[string]string `type:"map"`
}

// String returns the string representation
func (s Connector) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Connector) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Connector"}

	if s.ConnectorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectorArn"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the connector definition version, which is a container
// for connectors.
type ConnectorDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// A list of references to connectors in this version, with their corresponding
	// configuration settings.
	Connectors []Connector `type:"list"`
}

// String returns the string representation
func (s ConnectorDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConnectorDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConnectorDefinitionVersion"}
	if s.Connectors != nil {
		for i, v := range s.Connectors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Connectors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a core.
type Core struct {
	_ struct{} `type:"structure"`

	// The ARN of the certificate associated with the core.
	//
	// CertificateArn is a required field
	CertificateArn *string `type:"string" required:"true"`

	// A descriptive or arbitrary ID for the core. This value must be unique within
	// the core definition version. Max length is 128 characters with pattern ''[a-zA-Z0-9:_-]+''.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// If true, the core's local shadow is automatically synced with the cloud.
	SyncShadow *bool `type:"boolean"`

	// The ARN of the thing which is the core.
	//
	// ThingArn is a required field
	ThingArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Core) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Core) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Core"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.ThingArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a core definition version.
type CoreDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// A list of cores in the core definition version.
	Cores []Core `type:"list"`
}

// String returns the string representation
func (s CoreDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CoreDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CoreDefinitionVersion"}
	if s.Cores != nil {
		for i, v := range s.Cores {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Cores", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a definition.
type DefinitionInformation struct {
	_ struct{} `type:"structure"`

	// The ARN of the definition.
	Arn *string `type:"string"`

	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string `type:"string"`

	// The ID of the definition.
	Id *string `type:"string"`

	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string `type:"string"`

	// The ID of the latest version associated with the definition.
	LatestVersion *string `type:"string"`

	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string `type:"string"`

	// The name of the definition.
	Name *string `type:"string"`

	// Tag(s) attached to the resource arn.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DefinitionInformation) String() string {
	return awsutil.Prettify(s)
}

// Information about a deployment.
type Deployment struct {
	_ struct{} `type:"structure"`

	// The time, in milliseconds since the epoch, when the deployment was created.
	CreatedAt *string `type:"string"`

	// The ARN of the deployment.
	DeploymentArn *string `type:"string"`

	// The ID of the deployment.
	DeploymentId *string `type:"string"`

	// The type of the deployment.
	DeploymentType enums.DeploymentType `type:"string" enum:"true"`

	// The ARN of the group for this deployment.
	GroupArn *string `type:"string"`
}

// String returns the string representation
func (s Deployment) String() string {
	return awsutil.Prettify(s)
}

// Information about a device.
type Device struct {
	_ struct{} `type:"structure"`

	// The ARN of the certificate associated with the device.
	//
	// CertificateArn is a required field
	CertificateArn *string `type:"string" required:"true"`

	// A descriptive or arbitrary ID for the device. This value must be unique within
	// the device definition version. Max length is 128 characters with pattern
	// ''[a-zA-Z0-9:_-]+''.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// If true, the device's local shadow will be automatically synced with the
	// cloud.
	SyncShadow *bool `type:"boolean"`

	// The thing ARN of the device.
	//
	// ThingArn is a required field
	ThingArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Device) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Device) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Device"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.ThingArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a device definition version.
type DeviceDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// A list of devices in the definition version.
	Devices []Device `type:"list"`
}

// String returns the string representation
func (s DeviceDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeviceDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeviceDefinitionVersion"}
	if s.Devices != nil {
		for i, v := range s.Devices {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Devices", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Details about the error.
type ErrorDetail struct {
	_ struct{} `type:"structure"`

	// A detailed error code.
	DetailedErrorCode *string `type:"string"`

	// A detailed error message.
	DetailedErrorMessage *string `type:"string"`
}

// String returns the string representation
func (s ErrorDetail) String() string {
	return awsutil.Prettify(s)
}

// Information about a Lambda function.
type Function struct {
	_ struct{} `type:"structure"`

	// The ARN of the Lambda function.
	FunctionArn *string `type:"string"`

	// The configuration of the Lambda function.
	FunctionConfiguration *FunctionConfiguration `type:"structure"`

	// A descriptive or arbitrary ID for the function. This value must be unique
	// within the function definition version. Max length is 128 characters with
	// pattern ''[a-zA-Z0-9:_-]+''.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Function) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Function) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Function"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.FunctionConfiguration != nil {
		if err := s.FunctionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("FunctionConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The configuration of the Lambda function.
type FunctionConfiguration struct {
	_ struct{} `type:"structure"`

	// The expected encoding type of the input payload for the function. The default
	// is ''json''.
	EncodingType enums.EncodingType `type:"string" enum:"true"`

	// The environment configuration of the function.
	Environment *FunctionConfigurationEnvironment `type:"structure"`

	// The execution arguments.
	ExecArgs *string `type:"string"`

	// The name of the function executable.
	Executable *string `type:"string"`

	// The memory size, in KB, which the function requires. This setting is not
	// applicable and should be cleared when you run the Lambda function without
	// containerization.
	MemorySize *int64 `type:"integer"`

	// True if the function is pinned. Pinned means the function is long-lived and
	// starts when the core starts.
	Pinned *bool `type:"boolean"`

	// The allowed function execution time, after which Lambda should terminate
	// the function. This timeout still applies to pinned Lambda functions for each
	// request.
	Timeout *int64 `type:"integer"`
}

// String returns the string representation
func (s FunctionConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FunctionConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FunctionConfiguration"}
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			invalidParams.AddNested("Environment", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The environment configuration of the function.
type FunctionConfigurationEnvironment struct {
	_ struct{} `type:"structure"`

	// If true, the Lambda function is allowed to access the host's /sys folder.
	// Use this when the Lambda function needs to read device information from /sys.
	// This setting applies only when you run the Lambda function in a Greengrass
	// container.
	AccessSysfs *bool `type:"boolean"`

	// Configuration related to executing the Lambda function
	Execution *FunctionExecutionConfig `type:"structure"`

	// A list of the resources, with their permissions, to which the Lambda function
	// will be granted access. A Lambda function can have at most 10 resources.
	// ResourceAccessPolicies apply only when you run the Lambda function in a Greengrass
	// container.
	ResourceAccessPolicies []ResourceAccessPolicy `type:"list"`

	// Environment variables for the Lambda function's configuration.
	Variables map[string]string `type:"map"`
}

// String returns the string representation
func (s FunctionConfigurationEnvironment) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FunctionConfigurationEnvironment) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FunctionConfigurationEnvironment"}
	if s.ResourceAccessPolicies != nil {
		for i, v := range s.ResourceAccessPolicies {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceAccessPolicies", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The default configuration that applies to all Lambda functions in the group.
// Individual Lambda functions can override these settings.
type FunctionDefaultConfig struct {
	_ struct{} `type:"structure"`

	// Configuration information that specifies how a Lambda function runs.
	Execution *FunctionDefaultExecutionConfig `type:"structure"`
}

// String returns the string representation
func (s FunctionDefaultConfig) String() string {
	return awsutil.Prettify(s)
}

// Configuration information that specifies how a Lambda function runs.
type FunctionDefaultExecutionConfig struct {
	_ struct{} `type:"structure"`

	// Specifies whether the Lambda function runs in a Greengrass container (default)
	// or without containerization. Unless your scenario requires that you run without
	// containerization, we recommend that you run in a Greengrass container. Omit
	// this value to run the Lambda function with the default containerization for
	// the group.
	IsolationMode enums.FunctionIsolationMode `type:"string" enum:"true"`

	// Specifies the user and group whose permissions are used when running the
	// Lambda function. You can specify one or both values to override the default
	// values. We recommend that you avoid running as root unless absolutely necessary
	// to minimize the risk of unintended changes or malicious attacks. To run as
	// root, you must set ''IsolationMode'' to ''NoContainer'' and update config.json
	// in ''greengrass-root/config'' to set ''allowFunctionsToRunAsRoot'' to ''yes''.
	RunAs *FunctionRunAsConfig `type:"structure"`
}

// String returns the string representation
func (s FunctionDefaultExecutionConfig) String() string {
	return awsutil.Prettify(s)
}

// Information about a function definition version.
type FunctionDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// The default configuration that applies to all Lambda functions in this function
	// definition version. Individual Lambda functions can override these settings.
	DefaultConfig *FunctionDefaultConfig `type:"structure"`

	// A list of Lambda functions in this function definition version.
	Functions []Function `type:"list"`
}

// String returns the string representation
func (s FunctionDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FunctionDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FunctionDefinitionVersion"}
	if s.Functions != nil {
		for i, v := range s.Functions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Functions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Configuration information that specifies how a Lambda function runs.
type FunctionExecutionConfig struct {
	_ struct{} `type:"structure"`

	// Specifies whether the Lambda function runs in a Greengrass container (default)
	// or without containerization. Unless your scenario requires that you run without
	// containerization, we recommend that you run in a Greengrass container. Omit
	// this value to run the Lambda function with the default containerization for
	// the group.
	IsolationMode enums.FunctionIsolationMode `type:"string" enum:"true"`

	// Specifies the user and group whose permissions are used when running the
	// Lambda function. You can specify one or both values to override the default
	// values. We recommend that you avoid running as root unless absolutely necessary
	// to minimize the risk of unintended changes or malicious attacks. To run as
	// root, you must set ''IsolationMode'' to ''NoContainer'' and update config.json
	// in ''greengrass-root/config'' to set ''allowFunctionsToRunAsRoot'' to ''yes''.
	RunAs *FunctionRunAsConfig `type:"structure"`
}

// String returns the string representation
func (s FunctionExecutionConfig) String() string {
	return awsutil.Prettify(s)
}

// Specifies the user and group whose permissions are used when running the
// Lambda function. You can specify one or both values to override the default
// values. We recommend that you avoid running as root unless absolutely necessary
// to minimize the risk of unintended changes or malicious attacks. To run as
// root, you must set ''IsolationMode'' to ''NoContainer'' and update config.json
// in ''greengrass-root/config'' to set ''allowFunctionsToRunAsRoot'' to ''yes''.
type FunctionRunAsConfig struct {
	_ struct{} `type:"structure"`

	// The group ID whose permissions are used to run a Lambda function.
	Gid *int64 `type:"integer"`

	// The user ID whose permissions are used to run a Lambda function.
	Uid *int64 `type:"integer"`
}

// String returns the string representation
func (s FunctionRunAsConfig) String() string {
	return awsutil.Prettify(s)
}

// Information about a certificate authority for a group.
type GroupCertificateAuthorityProperties struct {
	_ struct{} `type:"structure"`

	// The ARN of the certificate authority for the group.
	GroupCertificateAuthorityArn *string `type:"string"`

	// The ID of the certificate authority for the group.
	GroupCertificateAuthorityId *string `type:"string"`
}

// String returns the string representation
func (s GroupCertificateAuthorityProperties) String() string {
	return awsutil.Prettify(s)
}

// Information about a group.
type GroupInformation struct {
	_ struct{} `type:"structure"`

	// The ARN of the group.
	Arn *string `type:"string"`

	// The time, in milliseconds since the epoch, when the group was created.
	CreationTimestamp *string `type:"string"`

	// The ID of the group.
	Id *string `type:"string"`

	// The time, in milliseconds since the epoch, when the group was last updated.
	LastUpdatedTimestamp *string `type:"string"`

	// The ID of the latest version associated with the group.
	LatestVersion *string `type:"string"`

	// The ARN of the latest version associated with the group.
	LatestVersionArn *string `type:"string"`

	// The name of the group.
	Name *string `type:"string"`
}

// String returns the string representation
func (s GroupInformation) String() string {
	return awsutil.Prettify(s)
}

// Group owner related settings for local resources.
type GroupOwnerSetting struct {
	_ struct{} `type:"structure"`

	// If true, AWS IoT Greengrass automatically adds the specified Linux OS group
	// owner of the resource to the Lambda process privileges. Thus the Lambda process
	// will have the file access permissions of the added Linux group.
	AutoAddGroupOwner *bool `type:"boolean"`

	// The name of the Linux OS group whose privileges will be added to the Lambda
	// process. This field is optional.
	GroupOwner *string `type:"string"`
}

// String returns the string representation
func (s GroupOwnerSetting) String() string {
	return awsutil.Prettify(s)
}

// Information about a group version.
type GroupVersion struct {
	_ struct{} `type:"structure"`

	// The ARN of the connector definition version for this group.
	ConnectorDefinitionVersionArn *string `type:"string"`

	// The ARN of the core definition version for this group.
	CoreDefinitionVersionArn *string `type:"string"`

	// The ARN of the device definition version for this group.
	DeviceDefinitionVersionArn *string `type:"string"`

	// The ARN of the function definition version for this group.
	FunctionDefinitionVersionArn *string `type:"string"`

	// The ARN of the logger definition version for this group.
	LoggerDefinitionVersionArn *string `type:"string"`

	// The ARN of the resource definition version for this group.
	ResourceDefinitionVersionArn *string `type:"string"`

	// The ARN of the subscription definition version for this group.
	SubscriptionDefinitionVersionArn *string `type:"string"`
}

// String returns the string representation
func (s GroupVersion) String() string {
	return awsutil.Prettify(s)
}

// Attributes that define a local device resource.
type LocalDeviceResourceData struct {
	_ struct{} `type:"structure"`

	// Group/owner related settings for local resources.
	GroupOwnerSetting *GroupOwnerSetting `type:"structure"`

	// The local absolute path of the device resource. The source path for a device
	// resource can refer only to a character device or block device under ''/dev''.
	SourcePath *string `type:"string"`
}

// String returns the string representation
func (s LocalDeviceResourceData) String() string {
	return awsutil.Prettify(s)
}

// Attributes that define a local volume resource.
type LocalVolumeResourceData struct {
	_ struct{} `type:"structure"`

	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `type:"string"`

	// Allows you to configure additional group privileges for the Lambda process.
	// This field is optional.
	GroupOwnerSetting *GroupOwnerSetting `type:"structure"`

	// The local absolute path of the volume resource on the host. The source path
	// for a volume resource type cannot start with ''/sys''.
	SourcePath *string `type:"string"`
}

// String returns the string representation
func (s LocalVolumeResourceData) String() string {
	return awsutil.Prettify(s)
}

// Information about a logger
type Logger struct {
	_ struct{} `type:"structure"`

	// The component that will be subject to logging.
	//
	// Component is a required field
	Component enums.LoggerComponent `type:"string" required:"true" enum:"true"`

	// A descriptive or arbitrary ID for the logger. This value must be unique within
	// the logger definition version. Max length is 128 characters with pattern
	// ''[a-zA-Z0-9:_-]+''.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// The level of the logs.
	//
	// Level is a required field
	Level enums.LoggerLevel `type:"string" required:"true" enum:"true"`

	// The amount of file space, in KB, to use if the local file system is used
	// for logging purposes.
	Space *int64 `type:"integer"`

	// The type of log output which will be used.
	//
	// Type is a required field
	Type enums.LoggerType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Logger) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Logger) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Logger"}
	if len(s.Component) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Component"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if len(s.Level) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Level"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a logger definition version.
type LoggerDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// A list of loggers.
	Loggers []Logger `type:"list"`
}

// String returns the string representation
func (s LoggerDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LoggerDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LoggerDefinitionVersion"}
	if s.Loggers != nil {
		for i, v := range s.Loggers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Loggers", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a resource.
type Resource struct {
	_ struct{} `type:"structure"`

	// The resource ID, used to refer to a resource in the Lambda function configuration.
	// Max length is 128 characters with pattern ''[a-zA-Z0-9:_-]+''. This must
	// be unique within a Greengrass group.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// The descriptive resource name, which is displayed on the AWS IoT Greengrass
	// console. Max length 128 characters with pattern ''[a-zA-Z0-9:_-]+''. This
	// must be unique within a Greengrass group.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// A container of data for all resource types.
	//
	// ResourceDataContainer is a required field
	ResourceDataContainer *ResourceDataContainer `type:"structure" required:"true"`
}

// String returns the string representation
func (s Resource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Resource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Resource"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.ResourceDataContainer == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceDataContainer"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A policy used by the function to access a resource.
type ResourceAccessPolicy struct {
	_ struct{} `type:"structure"`

	// The permissions that the Lambda function has to the resource. Can be one
	// of ''rw'' (read/write) or ''ro'' (read-only).
	Permission enums.Permission `type:"string" enum:"true"`

	// The ID of the resource. (This ID is assigned to the resource when you create
	// the resource definiton.)
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResourceAccessPolicy) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResourceAccessPolicy) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResourceAccessPolicy"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A container for resource data. The container takes only one of the following
// supported resource data types: ''LocalDeviceResourceData'', ''LocalVolumeResourceData'',
// ''SageMakerMachineLearningModelResourceData'', ''S3MachineLearningModelResourceData'',
// ''SecretsManagerSecretResourceData''.
type ResourceDataContainer struct {
	_ struct{} `type:"structure"`

	// Attributes that define the local device resource.
	LocalDeviceResourceData *LocalDeviceResourceData `type:"structure"`

	// Attributes that define the local volume resource.
	LocalVolumeResourceData *LocalVolumeResourceData `type:"structure"`

	// Attributes that define an Amazon S3 machine learning resource.
	S3MachineLearningModelResourceData *S3MachineLearningModelResourceData `type:"structure"`

	// Attributes that define an Amazon SageMaker machine learning resource.
	SageMakerMachineLearningModelResourceData *SageMakerMachineLearningModelResourceData `type:"structure"`

	// Attributes that define a secret resource, which references a secret from
	// AWS Secrets Manager.
	SecretsManagerSecretResourceData *SecretsManagerSecretResourceData `type:"structure"`
}

// String returns the string representation
func (s ResourceDataContainer) String() string {
	return awsutil.Prettify(s)
}

// Information about a resource definition version.
type ResourceDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// A list of resources.
	Resources []Resource `type:"list"`
}

// String returns the string representation
func (s ResourceDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResourceDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResourceDefinitionVersion"}
	if s.Resources != nil {
		for i, v := range s.Resources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Resources", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Attributes that define an Amazon S3 machine learning resource.
type S3MachineLearningModelResourceData struct {
	_ struct{} `type:"structure"`

	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `type:"string"`

	// The URI of the source model in an S3 bucket. The model package must be in
	// tar.gz or .zip format.
	S3Uri *string `type:"string"`
}

// String returns the string representation
func (s S3MachineLearningModelResourceData) String() string {
	return awsutil.Prettify(s)
}

// Attributes that define an Amazon SageMaker machine learning resource.
type SageMakerMachineLearningModelResourceData struct {
	_ struct{} `type:"structure"`

	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `type:"string"`

	// The ARN of the Amazon SageMaker training job that represents the source model.
	SageMakerJobArn *string `type:"string"`
}

// String returns the string representation
func (s SageMakerMachineLearningModelResourceData) String() string {
	return awsutil.Prettify(s)
}

// Attributes that define a secret resource, which references a secret from
// AWS Secrets Manager. AWS IoT Greengrass stores a local, encrypted copy of
// the secret on the Greengrass core, where it can be securely accessed by connectors
// and Lambda functions.
type SecretsManagerSecretResourceData struct {
	_ struct{} `type:"structure"`

	// The ARN of the Secrets Manager secret to make available on the core. The
	// value of the secret's latest version (represented by the ''AWSCURRENT'' staging
	// label) is included by default.
	ARN *string `type:"string"`

	// Optional. The staging labels whose values you want to make available on the
	// core, in addition to ''AWSCURRENT''.
	AdditionalStagingLabelsToDownload []string `type:"list"`
}

// String returns the string representation
func (s SecretsManagerSecretResourceData) String() string {
	return awsutil.Prettify(s)
}

// Information about a subscription.
type Subscription struct {
	_ struct{} `type:"structure"`

	// A descriptive or arbitrary ID for the subscription. This value must be unique
	// within the subscription definition version. Max length is 128 characters
	// with pattern ''[a-zA-Z0-9:_-]+''.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// The source of the subscription. Can be a thing ARN, a Lambda function ARN,
	// a connector ARN, 'cloud' (which represents the AWS IoT cloud), or 'GGShadowService'.
	//
	// Source is a required field
	Source *string `type:"string" required:"true"`

	// The MQTT topic used to route the message.
	//
	// Subject is a required field
	Subject *string `type:"string" required:"true"`

	// Where the message is sent to. Can be a thing ARN, a Lambda function ARN,
	// a connector ARN, 'cloud' (which represents the AWS IoT cloud), or 'GGShadowService'.
	//
	// Target is a required field
	Target *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Subscription) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Subscription) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Subscription"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Source == nil {
		invalidParams.Add(aws.NewErrParamRequired("Source"))
	}

	if s.Subject == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subject"))
	}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a subscription definition version.
type SubscriptionDefinitionVersion struct {
	_ struct{} `type:"structure"`

	// A list of subscriptions.
	Subscriptions []Subscription `type:"list"`
}

// String returns the string representation
func (s SubscriptionDefinitionVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubscriptionDefinitionVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubscriptionDefinitionVersion"}
	if s.Subscriptions != nil {
		for i, v := range s.Subscriptions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Subscriptions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a version.
type VersionInformation struct {
	_ struct{} `type:"structure"`

	// The ARN of the version.
	Arn *string `type:"string"`

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string `type:"string"`

	// The ID of the parent definition that the version is associated with.
	Id *string `type:"string"`

	// The ID of the version.
	Version *string `type:"string"`
}

// String returns the string representation
func (s VersionInformation) String() string {
	return awsutil.Prettify(s)
}