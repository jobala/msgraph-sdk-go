package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse 
// Deprecated: This class is obsolete. Use imageWithWidthGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponseable 
// Deprecated: This class is obsolete. Use imageWithWidthGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
