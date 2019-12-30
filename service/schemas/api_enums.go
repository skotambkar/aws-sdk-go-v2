// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

type CodeGenerationStatus string

// Enum values for CodeGenerationStatus
const (
	CodeGenerationStatusCreateInProgress CodeGenerationStatus = "CREATE_IN_PROGRESS"
	CodeGenerationStatusCreateComplete   CodeGenerationStatus = "CREATE_COMPLETE"
	CodeGenerationStatusCreateFailed     CodeGenerationStatus = "CREATE_FAILED"
)

func (enum CodeGenerationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CodeGenerationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DiscovererState string

// Enum values for DiscovererState
const (
	DiscovererStateStarted DiscovererState = "STARTED"
	DiscovererStateStopped DiscovererState = "STOPPED"
)

func (enum DiscovererState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DiscovererState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Type string

// Enum values for Type
const (
	TypeOpenApi3 Type = "OpenApi3"
)

func (enum Type) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Type) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
