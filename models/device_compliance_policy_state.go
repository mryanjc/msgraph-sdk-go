package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyState device Compliance Policy State for a given device.
type DeviceCompliancePolicyState struct {
    Entity
    // The name of the policy for this policyBase
    displayName *string
    // Supported platform types for policies.
    platformType *PolicyPlatformType
    // Count of how many setting a policy holds
    settingCount *int32
    // The settingStates property
    settingStates []DeviceCompliancePolicySettingStateable
    // The state property
    state *ComplianceStatus
    // The version of the policy
    version *int32
}
// NewDeviceCompliancePolicyState instantiates a new deviceCompliancePolicyState and sets the default values.
func NewDeviceCompliancePolicyState()(*DeviceCompliancePolicyState) {
    m := &DeviceCompliancePolicyState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyState(), nil
}
// GetDisplayName gets the displayName property value. The name of the policy for this policyBase
func (m *DeviceCompliancePolicyState) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["platformType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePolicyPlatformType , m.SetPlatformType)
    res["settingCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSettingCount)
    res["settingStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateFromDiscriminatorValue , m.SetSettingStates)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceStatus , m.SetState)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *DeviceCompliancePolicyState) GetPlatformType()(*PolicyPlatformType) {
    return m.platformType
}
// GetSettingCount gets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceCompliancePolicyState) GetSettingCount()(*int32) {
    return m.settingCount
}
// GetSettingStates gets the settingStates property value. The settingStates property
func (m *DeviceCompliancePolicyState) GetSettingStates()([]DeviceCompliancePolicySettingStateable) {
    return m.settingStates
}
// GetState gets the state property value. The state property
func (m *DeviceCompliancePolicyState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetVersion gets the version property value. The version of the policy
func (m *DeviceCompliancePolicyState) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicyState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingCount", m.GetSettingCount())
        if err != nil {
            return err
        }
    }
    if m.GetSettingStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSettingStates())
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the policy for this policyBase
func (m *DeviceCompliancePolicyState) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *DeviceCompliancePolicyState) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// SetSettingCount sets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceCompliancePolicyState) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// SetSettingStates sets the settingStates property value. The settingStates property
func (m *DeviceCompliancePolicyState) SetSettingStates(value []DeviceCompliancePolicySettingStateable)() {
    m.settingStates = value
}
// SetState sets the state property value. The state property
func (m *DeviceCompliancePolicyState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetVersion sets the version property value. The version of the policy
func (m *DeviceCompliancePolicyState) SetVersion(value *int32)() {
    m.version = value
}
