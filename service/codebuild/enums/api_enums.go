// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ArtifactNamespace string

// Enum values for ArtifactNamespace
const (
	ArtifactNamespaceNone    ArtifactNamespace = "NONE"
	ArtifactNamespaceBuildId ArtifactNamespace = "BUILD_ID"
)

func (enum ArtifactNamespace) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArtifactNamespace) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ArtifactPackaging string

// Enum values for ArtifactPackaging
const (
	ArtifactPackagingNone ArtifactPackaging = "NONE"
	ArtifactPackagingZip  ArtifactPackaging = "ZIP"
)

func (enum ArtifactPackaging) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArtifactPackaging) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ArtifactsType string

// Enum values for ArtifactsType
const (
	ArtifactsTypeCodepipeline ArtifactsType = "CODEPIPELINE"
	ArtifactsTypeS3           ArtifactsType = "S3"
	ArtifactsTypeNoArtifacts  ArtifactsType = "NO_ARTIFACTS"
)

func (enum ArtifactsType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArtifactsType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuthType string

// Enum values for AuthType
const (
	AuthTypeOauth               AuthType = "OAUTH"
	AuthTypeBasicAuth           AuthType = "BASIC_AUTH"
	AuthTypePersonalAccessToken AuthType = "PERSONAL_ACCESS_TOKEN"
)

func (enum AuthType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuthType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BuildPhaseType string

// Enum values for BuildPhaseType
const (
	BuildPhaseTypeSubmitted       BuildPhaseType = "SUBMITTED"
	BuildPhaseTypeQueued          BuildPhaseType = "QUEUED"
	BuildPhaseTypeProvisioning    BuildPhaseType = "PROVISIONING"
	BuildPhaseTypeDownloadSource  BuildPhaseType = "DOWNLOAD_SOURCE"
	BuildPhaseTypeInstall         BuildPhaseType = "INSTALL"
	BuildPhaseTypePreBuild        BuildPhaseType = "PRE_BUILD"
	BuildPhaseTypeBuild           BuildPhaseType = "BUILD"
	BuildPhaseTypePostBuild       BuildPhaseType = "POST_BUILD"
	BuildPhaseTypeUploadArtifacts BuildPhaseType = "UPLOAD_ARTIFACTS"
	BuildPhaseTypeFinalizing      BuildPhaseType = "FINALIZING"
	BuildPhaseTypeCompleted       BuildPhaseType = "COMPLETED"
)

func (enum BuildPhaseType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BuildPhaseType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CacheMode string

// Enum values for CacheMode
const (
	CacheModeLocalDockerLayerCache CacheMode = "LOCAL_DOCKER_LAYER_CACHE"
	CacheModeLocalSourceCache      CacheMode = "LOCAL_SOURCE_CACHE"
	CacheModeLocalCustomCache      CacheMode = "LOCAL_CUSTOM_CACHE"
)

func (enum CacheMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CacheMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CacheType string

// Enum values for CacheType
const (
	CacheTypeNoCache CacheType = "NO_CACHE"
	CacheTypeS3      CacheType = "S3"
	CacheTypeLocal   CacheType = "LOCAL"
)

func (enum CacheType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CacheType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComputeType string

// Enum values for ComputeType
const (
	ComputeTypeBuildGeneral1Small  ComputeType = "BUILD_GENERAL1_SMALL"
	ComputeTypeBuildGeneral1Medium ComputeType = "BUILD_GENERAL1_MEDIUM"
	ComputeTypeBuildGeneral1Large  ComputeType = "BUILD_GENERAL1_LARGE"
)

func (enum ComputeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComputeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CredentialProviderType string

// Enum values for CredentialProviderType
const (
	CredentialProviderTypeSecretsManager CredentialProviderType = "SECRETS_MANAGER"
)

func (enum CredentialProviderType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CredentialProviderType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EnvironmentType string

// Enum values for EnvironmentType
const (
	EnvironmentTypeWindowsContainer EnvironmentType = "WINDOWS_CONTAINER"
	EnvironmentTypeLinuxContainer   EnvironmentType = "LINUX_CONTAINER"
)

func (enum EnvironmentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EnvironmentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EnvironmentVariableType string

// Enum values for EnvironmentVariableType
const (
	EnvironmentVariableTypePlaintext      EnvironmentVariableType = "PLAINTEXT"
	EnvironmentVariableTypeParameterStore EnvironmentVariableType = "PARAMETER_STORE"
)

func (enum EnvironmentVariableType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EnvironmentVariableType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ImagePullCredentialsType string

// Enum values for ImagePullCredentialsType
const (
	ImagePullCredentialsTypeCodebuild   ImagePullCredentialsType = "CODEBUILD"
	ImagePullCredentialsTypeServiceRole ImagePullCredentialsType = "SERVICE_ROLE"
)

func (enum ImagePullCredentialsType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ImagePullCredentialsType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LanguageType string

// Enum values for LanguageType
const (
	LanguageTypeJava    LanguageType = "JAVA"
	LanguageTypePython  LanguageType = "PYTHON"
	LanguageTypeNodeJs  LanguageType = "NODE_JS"
	LanguageTypeRuby    LanguageType = "RUBY"
	LanguageTypeGolang  LanguageType = "GOLANG"
	LanguageTypeDocker  LanguageType = "DOCKER"
	LanguageTypeAndroid LanguageType = "ANDROID"
	LanguageTypeDotnet  LanguageType = "DOTNET"
	LanguageTypeBase    LanguageType = "BASE"
	LanguageTypePhp     LanguageType = "PHP"
)

func (enum LanguageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LanguageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogsConfigStatusType string

// Enum values for LogsConfigStatusType
const (
	LogsConfigStatusTypeEnabled  LogsConfigStatusType = "ENABLED"
	LogsConfigStatusTypeDisabled LogsConfigStatusType = "DISABLED"
)

func (enum LogsConfigStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogsConfigStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PlatformType string

// Enum values for PlatformType
const (
	PlatformTypeDebian        PlatformType = "DEBIAN"
	PlatformTypeAmazonLinux   PlatformType = "AMAZON_LINUX"
	PlatformTypeUbuntu        PlatformType = "UBUNTU"
	PlatformTypeWindowsServer PlatformType = "WINDOWS_SERVER"
)

func (enum PlatformType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PlatformType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProjectSortByType string

// Enum values for ProjectSortByType
const (
	ProjectSortByTypeName             ProjectSortByType = "NAME"
	ProjectSortByTypeCreatedTime      ProjectSortByType = "CREATED_TIME"
	ProjectSortByTypeLastModifiedTime ProjectSortByType = "LAST_MODIFIED_TIME"
)

func (enum ProjectSortByType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProjectSortByType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServerType string

// Enum values for ServerType
const (
	ServerTypeGithub           ServerType = "GITHUB"
	ServerTypeBitbucket        ServerType = "BITBUCKET"
	ServerTypeGithubEnterprise ServerType = "GITHUB_ENTERPRISE"
)

func (enum ServerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortOrderType string

// Enum values for SortOrderType
const (
	SortOrderTypeAscending  SortOrderType = "ASCENDING"
	SortOrderTypeDescending SortOrderType = "DESCENDING"
)

func (enum SortOrderType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortOrderType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SourceAuthType string

// Enum values for SourceAuthType
const (
	SourceAuthTypeOauth SourceAuthType = "OAUTH"
)

func (enum SourceAuthType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SourceAuthType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeCodecommit       SourceType = "CODECOMMIT"
	SourceTypeCodepipeline     SourceType = "CODEPIPELINE"
	SourceTypeGithub           SourceType = "GITHUB"
	SourceTypeS3               SourceType = "S3"
	SourceTypeBitbucket        SourceType = "BITBUCKET"
	SourceTypeGithubEnterprise SourceType = "GITHUB_ENTERPRISE"
	SourceTypeNoSource         SourceType = "NO_SOURCE"
)

func (enum SourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StatusType string

// Enum values for StatusType
const (
	StatusTypeSucceeded  StatusType = "SUCCEEDED"
	StatusTypeFailed     StatusType = "FAILED"
	StatusTypeFault      StatusType = "FAULT"
	StatusTypeTimedOut   StatusType = "TIMED_OUT"
	StatusTypeInProgress StatusType = "IN_PROGRESS"
	StatusTypeStopped    StatusType = "STOPPED"
)

func (enum StatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WebhookFilterType string

// Enum values for WebhookFilterType
const (
	WebhookFilterTypeEvent          WebhookFilterType = "EVENT"
	WebhookFilterTypeBaseRef        WebhookFilterType = "BASE_REF"
	WebhookFilterTypeHeadRef        WebhookFilterType = "HEAD_REF"
	WebhookFilterTypeActorAccountId WebhookFilterType = "ACTOR_ACCOUNT_ID"
	WebhookFilterTypeFilePath       WebhookFilterType = "FILE_PATH"
)

func (enum WebhookFilterType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WebhookFilterType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
