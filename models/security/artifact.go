package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Artifact 
type Artifact struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // The OdataType property
    OdataType *string
}
// NewArtifact instantiates a new artifact and sets the default values.
func NewArtifact()(*Artifact) {
    m := &Artifact{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateArtifactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateArtifactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.host":
                        return NewHost(), nil
                    case "#microsoft.graph.security.hostComponent":
                        return NewHostComponent(), nil
                    case "#microsoft.graph.security.hostCookie":
                        return NewHostCookie(), nil
                    case "#microsoft.graph.security.hostname":
                        return NewHostname(), nil
                    case "#microsoft.graph.security.hostTracker":
                        return NewHostTracker(), nil
                    case "#microsoft.graph.security.ipAddress":
                        return NewIpAddress(), nil
                    case "#microsoft.graph.security.passiveDnsRecord":
                        return NewPassiveDnsRecord(), nil
                    case "#microsoft.graph.security.unclassifiedArtifact":
                        return NewUnclassifiedArtifact(), nil
                }
            }
        }
    }
    return NewArtifact(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Artifact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Artifact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// Artifactable 
type Artifactable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
