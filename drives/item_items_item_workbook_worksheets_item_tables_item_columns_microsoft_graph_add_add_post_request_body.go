package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The index property
    index *int32
    // The name property
    name *string
    // The values property
    values iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The index property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) GetIndex()(*int32) {
    return m.index
}
// GetName gets the name property value. The name property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) GetName()(*string) {
    return m.name
}
// GetValues gets the values property value. The values property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) GetValues()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.values
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIndex sets the index property value. The index property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) SetIndex(value *int32)() {
    m.index = value
}
// SetName sets the name property value. The name property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetValues sets the values property value. The values property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsMicrosoftGraphAddAddPostRequestBody) SetValues(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.values = value
}
