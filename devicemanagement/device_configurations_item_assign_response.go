package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationsItemAssignResponse 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type DeviceConfigurationsItemAssignResponse struct {
    DeviceConfigurationsItemAssignPostResponse
}
// NewDeviceConfigurationsItemAssignResponse instantiates a new DeviceConfigurationsItemAssignResponse and sets the default values.
func NewDeviceConfigurationsItemAssignResponse()(*DeviceConfigurationsItemAssignResponse) {
    m := &DeviceConfigurationsItemAssignResponse{
        DeviceConfigurationsItemAssignPostResponse: *NewDeviceConfigurationsItemAssignPostResponse(),
    }
    return m
}
// CreateDeviceConfigurationsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsItemAssignResponse(), nil
}
// DeviceConfigurationsItemAssignResponseable 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type DeviceConfigurationsItemAssignResponseable interface {
    DeviceConfigurationsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
