package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementMessagesUnfavoriteResponse 
// Deprecated: This class is obsolete. Use unfavoritePostResponse instead.
type ServiceAnnouncementMessagesUnfavoriteResponse struct {
    ServiceAnnouncementMessagesUnfavoritePostResponse
}
// NewServiceAnnouncementMessagesUnfavoriteResponse instantiates a new ServiceAnnouncementMessagesUnfavoriteResponse and sets the default values.
func NewServiceAnnouncementMessagesUnfavoriteResponse()(*ServiceAnnouncementMessagesUnfavoriteResponse) {
    m := &ServiceAnnouncementMessagesUnfavoriteResponse{
        ServiceAnnouncementMessagesUnfavoritePostResponse: *NewServiceAnnouncementMessagesUnfavoritePostResponse(),
    }
    return m
}
// CreateServiceAnnouncementMessagesUnfavoriteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementMessagesUnfavoriteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementMessagesUnfavoriteResponse(), nil
}
// ServiceAnnouncementMessagesUnfavoriteResponseable 
// Deprecated: This class is obsolete. Use unfavoritePostResponse instead.
type ServiceAnnouncementMessagesUnfavoriteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAnnouncementMessagesUnfavoritePostResponseable
}
