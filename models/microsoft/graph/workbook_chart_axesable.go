package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxesable 
type WorkbookChartAxesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCategoryAxis()(WorkbookChartAxisable)
    GetSeriesAxis()(WorkbookChartAxisable)
    GetValueAxis()(WorkbookChartAxisable)
    SetCategoryAxis(value WorkbookChartAxisable)()
    SetSeriesAxis(value WorkbookChartAxisable)()
    SetValueAxis(value WorkbookChartAxisable)()
}
