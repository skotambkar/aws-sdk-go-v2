// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ChangeStatus string

// Enum values for ChangeStatus
const (
	ChangeStatusPreparing ChangeStatus = "PREPARING"
	ChangeStatusApplying  ChangeStatus = "APPLYING"
	ChangeStatusSucceeded ChangeStatus = "SUCCEEDED"
	ChangeStatusCancelled ChangeStatus = "CANCELLED"
	ChangeStatusFailed    ChangeStatus = "FAILED"
)

func (enum ChangeStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

func (enum SortOrder) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortOrder) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}