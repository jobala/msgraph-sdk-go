package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceRole 
type AccessPackageResourceRole struct {
    Entity
}
// NewAccessPackageResourceRole instantiates a new accessPackageResourceRole and sets the default values.
func NewAccessPackageResourceRole()(*AccessPackageResourceRole) {
    m := &AccessPackageResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRole(), nil
}
// GetDescription gets the description property value. The description property
func (m *AccessPackageResourceRole) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AccessPackageResourceRole) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["originId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginId(val)
        }
        return nil
    }
    res["originSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginSystem(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AccessPackageResourceable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageResourceRole) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginId gets the originId property value. The originId property
func (m *AccessPackageResourceRole) GetOriginId()(*string) {
    val, err := m.GetBackingStore().Get("originId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginSystem gets the originSystem property value. The originSystem property
func (m *AccessPackageResourceRole) GetOriginSystem()(*string) {
    val, err := m.GetBackingStore().Get("originSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResource gets the resource property value. The resource property
func (m *AccessPackageResourceRole) GetResource()(AccessPackageResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description property
func (m *AccessPackageResourceRole) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AccessPackageResourceRole) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageResourceRole) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginId sets the originId property value. The originId property
func (m *AccessPackageResourceRole) SetOriginId(value *string)() {
    err := m.GetBackingStore().Set("originId", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginSystem sets the originSystem property value. The originSystem property
func (m *AccessPackageResourceRole) SetOriginSystem(value *string)() {
    err := m.GetBackingStore().Set("originSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. The resource property
func (m *AccessPackageResourceRole) SetResource(value AccessPackageResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageResourceRoleable 
type AccessPackageResourceRoleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetResource()(AccessPackageResourceable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetResource(value AccessPackageResourceable)()
}