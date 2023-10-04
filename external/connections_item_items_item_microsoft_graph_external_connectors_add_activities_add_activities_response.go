package external

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse 
// Deprecated: This class is obsolete. Use addActivitiesPostResponse instead.
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse struct {
    ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostResponse
}
// NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse instantiates a new ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse and sets the default values.
func NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse()(*ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) {
    m := &ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse{
        ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostResponse: *NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostResponse(),
    }
    return m
}
// CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse(), nil
}
// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable 
// Deprecated: This class is obsolete. Use addActivitiesPostResponse instead.
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable interface {
    ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
