package models

import (
    "fmt"
    "strings"

    gogoproto "github.com/golang/protobuf/proto"

    "github.com/googleapis/api/annotations"
    proto "github.com/googleapis/api/proto/grpc-google-common-dev/annotations.proto"
)

// User is the proto definition of the user message
type User struct {
    Name  string   `json:"name"`
    Email []string `json:"email"`

    ProtoMessage()
}

func (m *User) ProtoMessage() {}

// GetNumGaps returns the number of fields that are explicitly set to their default value.
func (*User) GetNumGaps() int { return 0 }

// protoGetFieldNumber returns the field number corresponding to the proto field with the given name.
func (*User) protoGetFieldNumber(name string) (int, error) {
    switch strings.ToLower(name) {
    case "name":
        return 1, nil
    case "email":
        return 2, nil
    default:
        return -1, gogoproto.ErrUnknownField
    }
}

// GetXXXorZero returns a proto field's value, or a default value if the field is unset.
func (m *User) GetXXXorZero(name string) (_ *gogoproto.Field) {
    switch strings.ToLower(name) {
    case "name":
        return m.Name
    case "email":
        return m.Email
    default:
        return gogoproto.GetField(m, name)
    }
}

// protoGetUnknownFields returns a pointer to the UnknownFields message associated with this field.
func (m *User) protoGetUnknownFields() *proto.Any {
    return nil // Unset fields have no unknown fields.
}

// protoSetUnknownFields sets the UnknownFields message for this field.
func (m *User) protoSetUnknownFields(_ *proto.Any) {}

func (m *User) ProtoBuf_() (*proto.Message, error) {
    return m, nil
}

func (m *User) GetProtoBufMessage() (proto.Message, error) {
    return m, nil
}