package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// TriggerAndScopeBasedConditions 
type TriggerAndScopeBasedConditions struct {
    WorkflowExecutionConditions
}
// NewTriggerAndScopeBasedConditions instantiates a new triggerAndScopeBasedConditions and sets the default values.
func NewTriggerAndScopeBasedConditions()(*TriggerAndScopeBasedConditions) {
    m := &TriggerAndScopeBasedConditions{
        WorkflowExecutionConditions: *NewWorkflowExecutionConditions(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.triggerAndScopeBasedConditions"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggerAndScopeBasedConditions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TriggerAndScopeBasedConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionConditions.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectSetable))
        }
        return nil
    }
    res["trigger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowExecutionTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrigger(val.(WorkflowExecutionTriggerable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TriggerAndScopeBasedConditions) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScope gets the scope property value. Defines who the workflow runs for.
func (m *TriggerAndScopeBasedConditions) GetScope()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectSetable) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectSetable)
    }
    return nil
}
// GetTrigger gets the trigger property value. What triggers a workflow to run.
func (m *TriggerAndScopeBasedConditions) GetTrigger()(WorkflowExecutionTriggerable) {
    val, err := m.GetBackingStore().Get("trigger")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkflowExecutionTriggerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TriggerAndScopeBasedConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionConditions.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trigger", m.GetTrigger())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TriggerAndScopeBasedConditions) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetScope sets the scope property value. Defines who the workflow runs for.
func (m *TriggerAndScopeBasedConditions) SetScope(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectSetable)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetTrigger sets the trigger property value. What triggers a workflow to run.
func (m *TriggerAndScopeBasedConditions) SetTrigger(value WorkflowExecutionTriggerable)() {
    err := m.GetBackingStore().Set("trigger", value)
    if err != nil {
        panic(err)
    }
}
// TriggerAndScopeBasedConditionsable 
type TriggerAndScopeBasedConditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowExecutionConditionsable
    GetOdataType()(*string)
    GetScope()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectSetable)
    GetTrigger()(WorkflowExecutionTriggerable)
    SetOdataType(value *string)()
    SetScope(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectSetable)()
    SetTrigger(value WorkflowExecutionTriggerable)()
}
