package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedMobileAppable 
type ManagedMobileAppable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetMobileAppIdentifier()(MobileAppIdentifierable)
    GetVersion()(*string)
    SetMobileAppIdentifier(value MobileAppIdentifierable)()
    SetVersion(value *string)()
}
