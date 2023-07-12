package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkConversationIdentity 
type TeamworkConversationIdentity struct {
    Identity
}
// NewTeamworkConversationIdentity instantiates a new teamworkConversationIdentity and sets the default values.
func NewTeamworkConversationIdentity()(*TeamworkConversationIdentity) {
    m := &TeamworkConversationIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.teamworkConversationIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamworkConversationIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkConversationIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkConversationIdentity(), nil
}
// GetConversationIdentityType gets the conversationIdentityType property value. Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue.
func (m *TeamworkConversationIdentity) GetConversationIdentityType()(*TeamworkConversationIdentityType) {
    val, err := m.GetBackingStore().Get("conversationIdentityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamworkConversationIdentityType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConversationIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["conversationIdentityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkConversationIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationIdentityType(val.(*TeamworkConversationIdentityType))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkConversationIdentity) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkConversationIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConversationIdentityType() != nil {
        cast := (*m.GetConversationIdentityType()).String()
        err = writer.WriteStringValue("conversationIdentityType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConversationIdentityType sets the conversationIdentityType property value. Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue.
func (m *TeamworkConversationIdentity) SetConversationIdentityType(value *TeamworkConversationIdentityType)() {
    err := m.GetBackingStore().Set("conversationIdentityType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkConversationIdentity) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkConversationIdentityable 
type TeamworkConversationIdentityable interface {
    Identityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConversationIdentityType()(*TeamworkConversationIdentityType)
    GetOdataType()(*string)
    SetConversationIdentityType(value *TeamworkConversationIdentityType)()
    SetOdataType(value *string)()
}
