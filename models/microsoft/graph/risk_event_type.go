package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type RiskEventType int

const (
    UNLIKELYTRAVEL_RISKEVENTTYPE RiskEventType = iota
    ANONYMIZEDIPADDRESS_RISKEVENTTYPE
    MALICIOUSIPADDRESS_RISKEVENTTYPE
    UNFAMILIARFEATURES_RISKEVENTTYPE
    MALWAREINFECTEDIPADDRESS_RISKEVENTTYPE
    SUSPICIOUSIPADDRESS_RISKEVENTTYPE
    LEAKEDCREDENTIALS_RISKEVENTTYPE
    INVESTIGATIONSTHREATINTELLIGENCE_RISKEVENTTYPE
    GENERIC_RISKEVENTTYPE
    ADMINCONFIRMEDUSERCOMPROMISED_RISKEVENTTYPE
    MCASIMPOSSIBLETRAVEL_RISKEVENTTYPE
    MCASSUSPICIOUSINBOXMANIPULATIONRULES_RISKEVENTTYPE
    INVESTIGATIONSTHREATINTELLIGENCESIGNINLINKED_RISKEVENTTYPE
    MALICIOUSIPADDRESSVALIDCREDENTIALSBLOCKEDIP_RISKEVENTTYPE
    UNKNOWNFUTUREVALUE_RISKEVENTTYPE
)

func (i RiskEventType) String() string {
    return []string{"UNLIKELYTRAVEL", "ANONYMIZEDIPADDRESS", "MALICIOUSIPADDRESS", "UNFAMILIARFEATURES", "MALWAREINFECTEDIPADDRESS", "SUSPICIOUSIPADDRESS", "LEAKEDCREDENTIALS", "INVESTIGATIONSTHREATINTELLIGENCE", "GENERIC", "ADMINCONFIRMEDUSERCOMPROMISED", "MCASIMPOSSIBLETRAVEL", "MCASSUSPICIOUSINBOXMANIPULATIONRULES", "INVESTIGATIONSTHREATINTELLIGENCESIGNINLINKED", "MALICIOUSIPADDRESSVALIDCREDENTIALSBLOCKEDIP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRiskEventType(v string) (interface{}, error) {
    result := UNLIKELYTRAVEL_RISKEVENTTYPE
    switch strings.ToUpper(v) {
        case "UNLIKELYTRAVEL":
            result = UNLIKELYTRAVEL_RISKEVENTTYPE
        case "ANONYMIZEDIPADDRESS":
            result = ANONYMIZEDIPADDRESS_RISKEVENTTYPE
        case "MALICIOUSIPADDRESS":
            result = MALICIOUSIPADDRESS_RISKEVENTTYPE
        case "UNFAMILIARFEATURES":
            result = UNFAMILIARFEATURES_RISKEVENTTYPE
        case "MALWAREINFECTEDIPADDRESS":
            result = MALWAREINFECTEDIPADDRESS_RISKEVENTTYPE
        case "SUSPICIOUSIPADDRESS":
            result = SUSPICIOUSIPADDRESS_RISKEVENTTYPE
        case "LEAKEDCREDENTIALS":
            result = LEAKEDCREDENTIALS_RISKEVENTTYPE
        case "INVESTIGATIONSTHREATINTELLIGENCE":
            result = INVESTIGATIONSTHREATINTELLIGENCE_RISKEVENTTYPE
        case "GENERIC":
            result = GENERIC_RISKEVENTTYPE
        case "ADMINCONFIRMEDUSERCOMPROMISED":
            result = ADMINCONFIRMEDUSERCOMPROMISED_RISKEVENTTYPE
        case "MCASIMPOSSIBLETRAVEL":
            result = MCASIMPOSSIBLETRAVEL_RISKEVENTTYPE
        case "MCASSUSPICIOUSINBOXMANIPULATIONRULES":
            result = MCASSUSPICIOUSINBOXMANIPULATIONRULES_RISKEVENTTYPE
        case "INVESTIGATIONSTHREATINTELLIGENCESIGNINLINKED":
            result = INVESTIGATIONSTHREATINTELLIGENCESIGNINLINKED_RISKEVENTTYPE
        case "MALICIOUSIPADDRESSVALIDCREDENTIALSBLOCKEDIP":
            result = MALICIOUSIPADDRESSVALIDCREDENTIALSBLOCKEDIP_RISKEVENTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RISKEVENTTYPE
        default:
            return 0, errors.New("Unknown RiskEventType value: " + v)
    }
    return &result, nil
}
func SerializeRiskEventType(values []RiskEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
