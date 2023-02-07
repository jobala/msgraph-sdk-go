package users

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody 
type ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The lastMessageReadDateTime property
    lastMessageReadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user property
    user iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkUserIdentityable
}
// NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody instantiates a new ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody and sets the default values.
func NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody()(*ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) {
    m := &ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastMessageReadDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastMessageReadDateTime(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetLastMessageReadDateTime gets the lastMessageReadDateTime property value. The lastMessageReadDateTime property
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) GetLastMessageReadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastMessageReadDateTime
}
// GetUser gets the user property value. The user property
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) GetUser()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkUserIdentityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastMessageReadDateTime", m.GetLastMessageReadDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLastMessageReadDateTime sets the lastMessageReadDateTime property value. The lastMessageReadDateTime property
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) SetLastMessageReadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastMessageReadDateTime = value
}
// SetUser sets the user property value. The user property
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBody) SetUser(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkUserIdentityable)() {
    m.user = value
}
