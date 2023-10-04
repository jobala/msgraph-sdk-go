package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPrimaryChannelMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamPrimaryChannelMessagesDeltaResponse struct {
    ItemTeamPrimaryChannelMessagesDeltaGetResponse
}
// NewItemTeamPrimaryChannelMessagesDeltaResponse instantiates a new ItemTeamPrimaryChannelMessagesDeltaResponse and sets the default values.
func NewItemTeamPrimaryChannelMessagesDeltaResponse()(*ItemTeamPrimaryChannelMessagesDeltaResponse) {
    m := &ItemTeamPrimaryChannelMessagesDeltaResponse{
        ItemTeamPrimaryChannelMessagesDeltaGetResponse: *NewItemTeamPrimaryChannelMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelMessagesDeltaResponse(), nil
}
// ItemTeamPrimaryChannelMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamPrimaryChannelMessagesDeltaResponseable interface {
    ItemTeamPrimaryChannelMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
