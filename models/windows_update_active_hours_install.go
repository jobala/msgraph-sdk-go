package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateActiveHoursInstall 
type WindowsUpdateActiveHoursInstall struct {
    WindowsUpdateInstallScheduleType
    // Active Hours End
    activeHoursEnd *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Active Hours Start
    activeHoursStart *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewWindowsUpdateActiveHoursInstall instantiates a new WindowsUpdateActiveHoursInstall and sets the default values.
func NewWindowsUpdateActiveHoursInstall()(*WindowsUpdateActiveHoursInstall) {
    m := &WindowsUpdateActiveHoursInstall{
        WindowsUpdateInstallScheduleType: *NewWindowsUpdateInstallScheduleType(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdateActiveHoursInstall"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsUpdateActiveHoursInstallFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateActiveHoursInstallFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateActiveHoursInstall(), nil
}
// GetActiveHoursEnd gets the activeHoursEnd property value. Active Hours End
func (m *WindowsUpdateActiveHoursInstall) GetActiveHoursEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.activeHoursEnd
}
// GetActiveHoursStart gets the activeHoursStart property value. Active Hours Start
func (m *WindowsUpdateActiveHoursInstall) GetActiveHoursStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.activeHoursStart
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateActiveHoursInstall) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateInstallScheduleType.GetFieldDeserializers()
    res["activeHoursEnd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveHoursEnd(val)
        }
        return nil
    }
    res["activeHoursStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveHoursStart(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsUpdateActiveHoursInstall) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateInstallScheduleType.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeOnlyValue("activeHoursEnd", m.GetActiveHoursEnd())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("activeHoursStart", m.GetActiveHoursStart())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveHoursEnd sets the activeHoursEnd property value. Active Hours End
func (m *WindowsUpdateActiveHoursInstall) SetActiveHoursEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.activeHoursEnd = value
}
// SetActiveHoursStart sets the activeHoursStart property value. Active Hours Start
func (m *WindowsUpdateActiveHoursInstall) SetActiveHoursStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.activeHoursStart = value
}
