// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ResourceOwner string

// Enum values for ResourceOwner
const (
	ResourceOwnerSelf          ResourceOwner = "SELF"
	ResourceOwnerOtherAccounts ResourceOwner = "OTHER-ACCOUNTS"
)

func (enum ResourceOwner) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceOwner) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceShareAssociationStatus string

// Enum values for ResourceShareAssociationStatus
const (
	ResourceShareAssociationStatusAssociating    ResourceShareAssociationStatus = "ASSOCIATING"
	ResourceShareAssociationStatusAssociated     ResourceShareAssociationStatus = "ASSOCIATED"
	ResourceShareAssociationStatusFailed         ResourceShareAssociationStatus = "FAILED"
	ResourceShareAssociationStatusDisassociating ResourceShareAssociationStatus = "DISASSOCIATING"
	ResourceShareAssociationStatusDisassociated  ResourceShareAssociationStatus = "DISASSOCIATED"
)

func (enum ResourceShareAssociationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceShareAssociationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceShareAssociationType string

// Enum values for ResourceShareAssociationType
const (
	ResourceShareAssociationTypePrincipal ResourceShareAssociationType = "PRINCIPAL"
	ResourceShareAssociationTypeResource  ResourceShareAssociationType = "RESOURCE"
)

func (enum ResourceShareAssociationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceShareAssociationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceShareInvitationStatus string

// Enum values for ResourceShareInvitationStatus
const (
	ResourceShareInvitationStatusPending  ResourceShareInvitationStatus = "PENDING"
	ResourceShareInvitationStatusAccepted ResourceShareInvitationStatus = "ACCEPTED"
	ResourceShareInvitationStatusRejected ResourceShareInvitationStatus = "REJECTED"
	ResourceShareInvitationStatusExpired  ResourceShareInvitationStatus = "EXPIRED"
)

func (enum ResourceShareInvitationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceShareInvitationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceShareStatus string

// Enum values for ResourceShareStatus
const (
	ResourceShareStatusPending  ResourceShareStatus = "PENDING"
	ResourceShareStatusActive   ResourceShareStatus = "ACTIVE"
	ResourceShareStatusFailed   ResourceShareStatus = "FAILED"
	ResourceShareStatusDeleting ResourceShareStatus = "DELETING"
	ResourceShareStatusDeleted  ResourceShareStatus = "DELETED"
)

func (enum ResourceShareStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceShareStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusAvailable                 ResourceStatus = "AVAILABLE"
	ResourceStatusZonalResourceInaccessible ResourceStatus = "ZONAL_RESOURCE_INACCESSIBLE"
	ResourceStatusLimitExceeded             ResourceStatus = "LIMIT_EXCEEDED"
	ResourceStatusUnavailable               ResourceStatus = "UNAVAILABLE"
	ResourceStatusPending                   ResourceStatus = "PENDING"
)

func (enum ResourceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
