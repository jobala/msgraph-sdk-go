package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number property
    number iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The significance property
    significance iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["significance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignificance(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetNumber gets the number property value. The number property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) GetNumber()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.number
}
// GetSignificance gets the significance property value. The significance property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) GetSignificance()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.significance
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("significance", m.GetSignificance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumber sets the number property value. The number property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) SetNumber(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.number = value
}
// SetSignificance sets the significance property value. The significance property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PrecisePostRequestBody) SetSignificance(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.significance = value
}