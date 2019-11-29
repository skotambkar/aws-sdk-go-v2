// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/enums"
)

// A MediaPackage VOD Asset resource.
type AssetShallow struct {
	_ struct{} `type:"structure"`

	// The ARN of the Asset.
	Arn *string `locationName:"arn" type:"string"`

	// The unique identifier for the Asset.
	Id *string `locationName:"id" type:"string"`

	// The ID of the PackagingGroup for the Asset.
	PackagingGroupId *string `locationName:"packagingGroupId" type:"string"`

	// The resource ID to include in SPEKE key requests.
	ResourceId *string `locationName:"resourceId" type:"string"`

	// ARN of the source object in S3.
	SourceArn *string `locationName:"sourceArn" type:"string"`

	// The IAM role ARN used to access the source S3 bucket.
	SourceRoleArn *string `locationName:"sourceRoleArn" type:"string"`
}

// String returns the string representation
func (s AssetShallow) String() string {
	return awsutil.Prettify(s)
}

// A CMAF encryption configuration.
type CmafEncryption struct {
	_ struct{} `type:"structure"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s CmafEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CmafEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CmafEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A CMAF packaging configuration.
type CmafPackage struct {
	_ struct{} `type:"structure"`

	// A CMAF encryption configuration.
	Encryption *CmafEncryption `locationName:"encryption" type:"structure"`

	// A list of HLS manifest configurations.
	//
	// HlsManifests is a required field
	HlsManifests []HlsManifest `locationName:"hlsManifests" type:"list" required:"true"`

	// Duration (in seconds) of each fragment. Actual fragments will berounded to
	// the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`
}

// String returns the string representation
func (s CmafPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CmafPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CmafPackage"}

	if s.HlsManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("HlsManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
type DashEncryption struct {
	_ struct{} `type:"structure"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s DashEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DashEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DashEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A DASH manifest configuration.
type DashManifest struct {
	_ struct{} `type:"structure"`

	// An optional string to include in the name of the manifest.
	ManifestName *string `locationName:"manifestName" type:"string"`

	// Minimum duration (in seconds) that a player will buffer media before starting
	// the presentation.
	MinBufferTimeSeconds *int64 `locationName:"minBufferTimeSeconds" type:"integer"`

	// The Dynamic Adaptive Streaming over HTTP (DASH) profile type. When set to
	// "HBBTV_1_5", HbbTV 1.5 compliant output is enabled.
	Profile enums.Profile `locationName:"profile" type:"string" enum:"true"`

	// A StreamSelection configuration.
	StreamSelection *StreamSelection `locationName:"streamSelection" type:"structure"`
}

// String returns the string representation
func (s DashManifest) String() string {
	return awsutil.Prettify(s)
}

// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
type DashPackage struct {
	_ struct{} `type:"structure"`

	// A list of DASH manifest configurations.
	//
	// DashManifests is a required field
	DashManifests []DashManifest `locationName:"dashManifests" type:"list" required:"true"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
	Encryption *DashEncryption `locationName:"encryption" type:"structure"`

	// Duration (in seconds) of each segment. Actual segments will berounded to
	// the nearest multiple of the source segment duration.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`
}

// String returns the string representation
func (s DashPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DashPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DashPackage"}

	if s.DashManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The endpoint URL used to access an Asset using one PackagingConfiguration.
type EgressEndpoint struct {
	_ struct{} `type:"structure"`

	// The ID of the PackagingConfiguration being applied to the Asset.
	PackagingConfigurationId *string `locationName:"packagingConfigurationId" type:"string"`

	// The URL of the parent manifest for the repackaged Asset.
	Url *string `locationName:"url" type:"string"`
}

// String returns the string representation
func (s EgressEndpoint) String() string {
	return awsutil.Prettify(s)
}

// An HTTP Live Streaming (HLS) encryption configuration.
type HlsEncryption struct {
	_ struct{} `type:"structure"`

	// A constant initialization vector for encryption (optional).When not specified
	// the initialization vector will be periodically rotated.
	ConstantInitializationVector *string `locationName:"constantInitializationVector" type:"string"`

	// The encryption method to use.
	EncryptionMethod enums.EncryptionMethod `locationName:"encryptionMethod" type:"string" enum:"true"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s HlsEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HlsEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HlsEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An HTTP Live Streaming (HLS) manifest configuration.
type HlsManifest struct {
	_ struct{} `type:"structure"`

	// This setting controls how ad markers are included in the packaged OriginEndpoint."NONE"
	// will omit all SCTE-35 ad markers from the output."PASSTHROUGH" causes the
	// manifest to contain a copy of the SCTE-35 admarkers (comments) taken directly
	// from the input HTTP Live Streaming (HLS) manifest."SCTE35_ENHANCED" generates
	// ad markers and blackout tags based on SCTE-35messages in the input source.
	AdMarkers enums.AdMarkers `locationName:"adMarkers" type:"string" enum:"true"`

	// When enabled, an I-Frame only stream will be included in the output.
	IncludeIframeOnlyStream *bool `locationName:"includeIframeOnlyStream" type:"boolean"`

	// An optional string to include in the name of the manifest.
	ManifestName *string `locationName:"manifestName" type:"string"`

	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME taginserted
	// into manifests. Additionally, when an interval is specifiedID3Timed Metadata
	// messages will be generated every 5 seconds using theingest time of the content.If
	// the interval is not specified, or set to 0, thenno EXT-X-PROGRAM-DATE-TIME
	// tags will be inserted into manifests and noID3Timed Metadata messages will
	// be generated. Note that irrespectiveof this parameter, if any ID3 Timed Metadata
	// is found in HTTP Live Streaming (HLS) input,it will be passed through to
	// HLS output.
	ProgramDateTimeIntervalSeconds *int64 `locationName:"programDateTimeIntervalSeconds" type:"integer"`

	// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
	RepeatExtXKey *bool `locationName:"repeatExtXKey" type:"boolean"`

	// A StreamSelection configuration.
	StreamSelection *StreamSelection `locationName:"streamSelection" type:"structure"`
}

// String returns the string representation
func (s HlsManifest) String() string {
	return awsutil.Prettify(s)
}

// An HTTP Live Streaming (HLS) packaging configuration.
type HlsPackage struct {
	_ struct{} `type:"structure"`

	// An HTTP Live Streaming (HLS) encryption configuration.
	Encryption *HlsEncryption `locationName:"encryption" type:"structure"`

	// A list of HLS manifest configurations.
	//
	// HlsManifests is a required field
	HlsManifests []HlsManifest `locationName:"hlsManifests" type:"list" required:"true"`

	// Duration (in seconds) of each fragment. Actual fragments will berounded to
	// the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`

	// When enabled, audio streams will be placed in rendition groups in the output.
	UseAudioRenditionGroup *bool `locationName:"useAudioRenditionGroup" type:"boolean"`
}

// String returns the string representation
func (s HlsPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HlsPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HlsPackage"}

	if s.HlsManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("HlsManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A Microsoft Smooth Streaming (MSS) encryption configuration.
type MssEncryption struct {
	_ struct{} `type:"structure"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s MssEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MssEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MssEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A Microsoft Smooth Streaming (MSS) manifest configuration.
type MssManifest struct {
	_ struct{} `type:"structure"`

	// An optional string to include in the name of the manifest.
	ManifestName *string `locationName:"manifestName" type:"string"`

	// A StreamSelection configuration.
	StreamSelection *StreamSelection `locationName:"streamSelection" type:"structure"`
}

// String returns the string representation
func (s MssManifest) String() string {
	return awsutil.Prettify(s)
}

// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
type MssPackage struct {
	_ struct{} `type:"structure"`

	// A Microsoft Smooth Streaming (MSS) encryption configuration.
	Encryption *MssEncryption `locationName:"encryption" type:"structure"`

	// A list of MSS manifest configurations.
	//
	// MssManifests is a required field
	MssManifests []MssManifest `locationName:"mssManifests" type:"list" required:"true"`

	// The duration (in seconds) of each segment.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`
}

// String returns the string representation
func (s MssPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MssPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MssPackage"}

	if s.MssManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("MssManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A MediaPackage VOD PackagingConfiguration resource.
type PackagingConfiguration struct {
	_ struct{} `type:"structure"`

	// The ARN of the PackagingConfiguration.
	Arn *string `locationName:"arn" type:"string"`

	// A CMAF packaging configuration.
	CmafPackage *CmafPackage `locationName:"cmafPackage" type:"structure"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage `locationName:"dashPackage" type:"structure"`

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage `locationName:"hlsPackage" type:"structure"`

	// The ID of the PackagingConfiguration.
	Id *string `locationName:"id" type:"string"`

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *MssPackage `locationName:"mssPackage" type:"structure"`

	// The ID of a PackagingGroup.
	PackagingGroupId *string `locationName:"packagingGroupId" type:"string"`
}

// String returns the string representation
func (s PackagingConfiguration) String() string {
	return awsutil.Prettify(s)
}

// A MediaPackage VOD PackagingGroup resource.
type PackagingGroup struct {
	_ struct{} `type:"structure"`

	// The ARN of the PackagingGroup.
	Arn *string `locationName:"arn" type:"string"`

	// The ID of the PackagingGroup.
	Id *string `locationName:"id" type:"string"`
}

// String returns the string representation
func (s PackagingGroup) String() string {
	return awsutil.Prettify(s)
}

// A configuration for accessing an external Secure Packager and Encoder Key
// Exchange (SPEKE) service that will provide encryption keys.
type SpekeKeyProvider struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) of an IAM role that AWS ElementalMediaPackage
	// will assume when accessing the key provider service.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" type:"string" required:"true"`

	// The system IDs to include in key requests.
	//
	// SystemIds is a required field
	SystemIds []string `locationName:"systemIds" type:"list" required:"true"`

	// The URL of the external key provider service.
	//
	// Url is a required field
	Url *string `locationName:"url" type:"string" required:"true"`
}

// String returns the string representation
func (s SpekeKeyProvider) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SpekeKeyProvider) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SpekeKeyProvider"}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if s.SystemIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SystemIds"))
	}

	if s.Url == nil {
		invalidParams.Add(aws.NewErrParamRequired("Url"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A StreamSelection configuration.
type StreamSelection struct {
	_ struct{} `type:"structure"`

	// The maximum video bitrate (bps) to include in output.
	MaxVideoBitsPerSecond *int64 `locationName:"maxVideoBitsPerSecond" type:"integer"`

	// The minimum video bitrate (bps) to include in output.
	MinVideoBitsPerSecond *int64 `locationName:"minVideoBitsPerSecond" type:"integer"`

	// A directive that determines the order of streams in the output.
	StreamOrder enums.StreamOrder `locationName:"streamOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s StreamSelection) String() string {
	return awsutil.Prettify(s)
}