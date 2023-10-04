package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody 
type MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody instantiates a new MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody and sets the default values.
func NewMobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody()(*MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) {
    m := &MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fileEncryptionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFileEncryptionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileEncryptionInfo(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileEncryptionInfoable))
        }
        return nil
    }
    return res
}
// GetFileEncryptionInfo gets the fileEncryptionInfo property value. The fileEncryptionInfo property
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) GetFileEncryptionInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileEncryptionInfoable) {
    val, err := m.GetBackingStore().Get("fileEncryptionInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileEncryptionInfoable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("fileEncryptionInfo", m.GetFileEncryptionInfo())
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
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFileEncryptionInfo sets the fileEncryptionInfo property value. The fileEncryptionInfo property
func (m *MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBody) SetFileEncryptionInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileEncryptionInfoable)() {
    err := m.GetBackingStore().Set("fileEncryptionInfo", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBodyable 
type MobileAppsItemGraphIosLobAppContentVersionsItemFilesItemCommitPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFileEncryptionInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileEncryptionInfoable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFileEncryptionInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileEncryptionInfoable)()
}
