package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AppCatalogs struct {
    Entity
    // 
    teamsApps []TeamsApp;
}
// Instantiates a new appCatalogs and sets the default values.
func NewAppCatalogs()(*AppCatalogs) {
    m := &AppCatalogs{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the teamsApps property value. 
func (m *AppCatalogs) GetTeamsApps()([]TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApps
    }
}
// The deserialization information for the current model
func (m *AppCatalogs) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["teamsApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamsApp))
            }
            m.SetTeamsApps(res)
        }
        return nil
    }
    return res
}
func (m *AppCatalogs) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppCatalogs) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTeamsApps()))
        for i, v := range m.GetTeamsApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("teamsApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the teamsApps property value. 
// Parameters:
//  - value : Value to set for the teamsApps property.
func (m *AppCatalogs) SetTeamsApps(value []TeamsApp)() {
    m.teamsApps = value
}
