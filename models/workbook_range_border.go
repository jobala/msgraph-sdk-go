package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeBorder 
type WorkbookRangeBorder struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewWorkbookRangeBorder instantiates a new workbookRangeBorder and sets the default values.
func NewWorkbookRangeBorder()(*WorkbookRangeBorder) {
    m := &WorkbookRangeBorder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeBorderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookRangeBorderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeBorder(), nil
}
// GetColor gets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange').
func (m *WorkbookRangeBorder) GetColor()(*string) {
    val, err := m.GetBackingStore().Get("color")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeBorder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["sideIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSideIndex(val)
        }
        return nil
    }
    res["style"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStyle(val)
        }
        return nil
    }
    res["weight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetSideIndex gets the sideIndex property value. Constant value that indicates the specific side of the border. The possible values are: EdgeTop, EdgeBottom, EdgeLeft, EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only.
func (m *WorkbookRangeBorder) GetSideIndex()(*string) {
    val, err := m.GetBackingStore().Get("sideIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStyle gets the style property value. One of the constants of line style specifying the line style for the border. The possible values are: None, Continuous, Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot.
func (m *WorkbookRangeBorder) GetStyle()(*string) {
    val, err := m.GetBackingStore().Get("style")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWeight gets the weight property value. Specifies the weight of the border around a range. The possible values are: Hairline, Thin, Medium, Thick.
func (m *WorkbookRangeBorder) GetWeight()(*string) {
    val, err := m.GetBackingStore().Get("weight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookRangeBorder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sideIndex", m.GetSideIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("style", m.GetStyle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColor sets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange').
func (m *WorkbookRangeBorder) SetColor(value *string)() {
    err := m.GetBackingStore().Set("color", value)
    if err != nil {
        panic(err)
    }
}
// SetSideIndex sets the sideIndex property value. Constant value that indicates the specific side of the border. The possible values are: EdgeTop, EdgeBottom, EdgeLeft, EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only.
func (m *WorkbookRangeBorder) SetSideIndex(value *string)() {
    err := m.GetBackingStore().Set("sideIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetStyle sets the style property value. One of the constants of line style specifying the line style for the border. The possible values are: None, Continuous, Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot.
func (m *WorkbookRangeBorder) SetStyle(value *string)() {
    err := m.GetBackingStore().Set("style", value)
    if err != nil {
        panic(err)
    }
}
// SetWeight sets the weight property value. Specifies the weight of the border around a range. The possible values are: Hairline, Thin, Medium, Thick.
func (m *WorkbookRangeBorder) SetWeight(value *string)() {
    err := m.GetBackingStore().Set("weight", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookRangeBorderable 
type WorkbookRangeBorderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetSideIndex()(*string)
    GetStyle()(*string)
    GetWeight()(*string)
    SetColor(value *string)()
    SetSideIndex(value *string)()
    SetStyle(value *string)()
    SetWeight(value *string)()
}
