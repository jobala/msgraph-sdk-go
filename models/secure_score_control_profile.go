package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureScoreControlProfile provides operations to manage the collection of agreement entities.
type SecureScoreControlProfile struct {
    Entity
    // Control action type (Config, Review, Behavior).
    actionType *string
    // URL to where the control can be actioned.
    actionUrl *string
    // GUID string for tenant ID.
    azureTenantId *string
    // The complianceInformation property
    complianceInformation []ComplianceInformationable
    // Control action category (Identity, Data, Device, Apps, Infrastructure).
    controlCategory *string
    // The controlStateUpdates property
    controlStateUpdates []SecureScoreControlStateUpdateable
    // Flag to indicate if a control is depreciated.
    deprecated *bool
    // Resource cost of implemmentating control (low, moderate, high).
    implementationCost *string
    // Time at which the control profile entity was last modified. The Timestamp type represents date and time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // max attainable score for the control.
    maxScore *float64
    // Microsoft's stack ranking of control.
    rank *int32
    // Description of what the control will help remediate.
    remediation *string
    // Description of the impact on users of the remediation.
    remediationImpact *string
    // Service that owns the control (Exchange, Sharepoint, Azure AD).
    service *string
    // List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
    threats []string
    // The tier property
    tier *string
    // Title of the control.
    title *string
    // The userImpact property
    userImpact *string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
}
// NewSecureScoreControlProfile instantiates a new secureScoreControlProfile and sets the default values.
func NewSecureScoreControlProfile()(*SecureScoreControlProfile) {
    m := &SecureScoreControlProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecureScoreControlProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoreControlProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureScoreControlProfile(), nil
}
// GetActionType gets the actionType property value. Control action type (Config, Review, Behavior).
func (m *SecureScoreControlProfile) GetActionType()(*string) {
    return m.actionType
}
// GetActionUrl gets the actionUrl property value. URL to where the control can be actioned.
func (m *SecureScoreControlProfile) GetActionUrl()(*string) {
    return m.actionUrl
}
// GetAzureTenantId gets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScoreControlProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetComplianceInformation gets the complianceInformation property value. The complianceInformation property
func (m *SecureScoreControlProfile) GetComplianceInformation()([]ComplianceInformationable) {
    return m.complianceInformation
}
// GetControlCategory gets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *SecureScoreControlProfile) GetControlCategory()(*string) {
    return m.controlCategory
}
// GetControlStateUpdates gets the controlStateUpdates property value. The controlStateUpdates property
func (m *SecureScoreControlProfile) GetControlStateUpdates()([]SecureScoreControlStateUpdateable) {
    return m.controlStateUpdates
}
// GetDeprecated gets the deprecated property value. Flag to indicate if a control is depreciated.
func (m *SecureScoreControlProfile) GetDeprecated()(*bool) {
    return m.deprecated
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScoreControlProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActionType)
    res["actionUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActionUrl)
    res["azureTenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureTenantId)
    res["complianceInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateComplianceInformationFromDiscriminatorValue , m.SetComplianceInformation)
    res["controlCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetControlCategory)
    res["controlStateUpdates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecureScoreControlStateUpdateFromDiscriminatorValue , m.SetControlStateUpdates)
    res["deprecated"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeprecated)
    res["implementationCost"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetImplementationCost)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["maxScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetMaxScore)
    res["rank"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRank)
    res["remediation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemediation)
    res["remediationImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemediationImpact)
    res["service"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetService)
    res["threats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetThreats)
    res["tier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTier)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTitle)
    res["userImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserImpact)
    res["vendorInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue , m.SetVendorInformation)
    return res
}
// GetImplementationCost gets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
func (m *SecureScoreControlProfile) GetImplementationCost()(*string) {
    return m.implementationCost
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
func (m *SecureScoreControlProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMaxScore gets the maxScore property value. max attainable score for the control.
func (m *SecureScoreControlProfile) GetMaxScore()(*float64) {
    return m.maxScore
}
// GetRank gets the rank property value. Microsoft's stack ranking of control.
func (m *SecureScoreControlProfile) GetRank()(*int32) {
    return m.rank
}
// GetRemediation gets the remediation property value. Description of what the control will help remediate.
func (m *SecureScoreControlProfile) GetRemediation()(*string) {
    return m.remediation
}
// GetRemediationImpact gets the remediationImpact property value. Description of the impact on users of the remediation.
func (m *SecureScoreControlProfile) GetRemediationImpact()(*string) {
    return m.remediationImpact
}
// GetService gets the service property value. Service that owns the control (Exchange, Sharepoint, Azure AD).
func (m *SecureScoreControlProfile) GetService()(*string) {
    return m.service
}
// GetThreats gets the threats property value. List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
func (m *SecureScoreControlProfile) GetThreats()([]string) {
    return m.threats
}
// GetTier gets the tier property value. The tier property
func (m *SecureScoreControlProfile) GetTier()(*string) {
    return m.tier
}
// GetTitle gets the title property value. Title of the control.
func (m *SecureScoreControlProfile) GetTitle()(*string) {
    return m.title
}
// GetUserImpact gets the userImpact property value. The userImpact property
func (m *SecureScoreControlProfile) GetUserImpact()(*string) {
    return m.userImpact
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *SecureScoreControlProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// Serialize serializes information the current object
func (m *SecureScoreControlProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionType", m.GetActionType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("actionUrl", m.GetActionUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComplianceInformation())
        err = writer.WriteCollectionOfObjectValues("complianceInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("controlCategory", m.GetControlCategory())
        if err != nil {
            return err
        }
    }
    if m.GetControlStateUpdates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetControlStateUpdates())
        err = writer.WriteCollectionOfObjectValues("controlStateUpdates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deprecated", m.GetDeprecated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("implementationCost", m.GetImplementationCost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediation", m.GetRemediation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediationImpact", m.GetRemediationImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetThreats() != nil {
        err = writer.WriteCollectionOfStringValues("threats", m.GetThreats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tier", m.GetTier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userImpact", m.GetUserImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionType sets the actionType property value. Control action type (Config, Review, Behavior).
func (m *SecureScoreControlProfile) SetActionType(value *string)() {
    m.actionType = value
}
// SetActionUrl sets the actionUrl property value. URL to where the control can be actioned.
func (m *SecureScoreControlProfile) SetActionUrl(value *string)() {
    m.actionUrl = value
}
// SetAzureTenantId sets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScoreControlProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetComplianceInformation sets the complianceInformation property value. The complianceInformation property
func (m *SecureScoreControlProfile) SetComplianceInformation(value []ComplianceInformationable)() {
    m.complianceInformation = value
}
// SetControlCategory sets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *SecureScoreControlProfile) SetControlCategory(value *string)() {
    m.controlCategory = value
}
// SetControlStateUpdates sets the controlStateUpdates property value. The controlStateUpdates property
func (m *SecureScoreControlProfile) SetControlStateUpdates(value []SecureScoreControlStateUpdateable)() {
    m.controlStateUpdates = value
}
// SetDeprecated sets the deprecated property value. Flag to indicate if a control is depreciated.
func (m *SecureScoreControlProfile) SetDeprecated(value *bool)() {
    m.deprecated = value
}
// SetImplementationCost sets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
func (m *SecureScoreControlProfile) SetImplementationCost(value *string)() {
    m.implementationCost = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
func (m *SecureScoreControlProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMaxScore sets the maxScore property value. max attainable score for the control.
func (m *SecureScoreControlProfile) SetMaxScore(value *float64)() {
    m.maxScore = value
}
// SetRank sets the rank property value. Microsoft's stack ranking of control.
func (m *SecureScoreControlProfile) SetRank(value *int32)() {
    m.rank = value
}
// SetRemediation sets the remediation property value. Description of what the control will help remediate.
func (m *SecureScoreControlProfile) SetRemediation(value *string)() {
    m.remediation = value
}
// SetRemediationImpact sets the remediationImpact property value. Description of the impact on users of the remediation.
func (m *SecureScoreControlProfile) SetRemediationImpact(value *string)() {
    m.remediationImpact = value
}
// SetService sets the service property value. Service that owns the control (Exchange, Sharepoint, Azure AD).
func (m *SecureScoreControlProfile) SetService(value *string)() {
    m.service = value
}
// SetThreats sets the threats property value. List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
func (m *SecureScoreControlProfile) SetThreats(value []string)() {
    m.threats = value
}
// SetTier sets the tier property value. The tier property
func (m *SecureScoreControlProfile) SetTier(value *string)() {
    m.tier = value
}
// SetTitle sets the title property value. Title of the control.
func (m *SecureScoreControlProfile) SetTitle(value *string)() {
    m.title = value
}
// SetUserImpact sets the userImpact property value. The userImpact property
func (m *SecureScoreControlProfile) SetUserImpact(value *string)() {
    m.userImpact = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *SecureScoreControlProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
