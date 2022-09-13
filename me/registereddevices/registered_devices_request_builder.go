package registereddevices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0038fbe769ee74c442a37b3f1aa7c1f69e95198e07d90c1b3cc63b03e661a509 "github.com/microsoftgraph/msgraph-sdk-go/me/registereddevices/count"
    i319f64e4aebb40c6f2e1f6fc23df6af58292586d1bf495814fddf15918be7117 "github.com/microsoftgraph/msgraph-sdk-go/me/registereddevices/device"
    i55e406c2335376386e0a9cfc801f8b8b57bc5ab242aed29c62fa441f5b8f7d8e "github.com/microsoftgraph/msgraph-sdk-go/me/registereddevices/endpoint"
    i7b023c3a80ba448d71c19f1c176e9bb957a06e4ba7f7e4f0282e48a92b6350b3 "github.com/microsoftgraph/msgraph-sdk-go/me/registereddevices/approleassignment"
)

// RegisteredDevicesRequestBuilder provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
type RegisteredDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RegisteredDevicesRequestBuilderGetQueryParameters devices that are registered for the user. Read-only. Nullable. Supports $expand.
type RegisteredDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// RegisteredDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RegisteredDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RegisteredDevicesRequestBuilderGetQueryParameters
}
// AppRoleAssignment the appRoleAssignment property
func (m *RegisteredDevicesRequestBuilder) AppRoleAssignment()(*i7b023c3a80ba448d71c19f1c176e9bb957a06e4ba7f7e4f0282e48a92b6350b3.AppRoleAssignmentRequestBuilder) {
    return i7b023c3a80ba448d71c19f1c176e9bb957a06e4ba7f7e4f0282e48a92b6350b3.NewAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRegisteredDevicesRequestBuilderInternal instantiates a new RegisteredDevicesRequestBuilder and sets the default values.
func NewRegisteredDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredDevicesRequestBuilder) {
    m := &RegisteredDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/registeredDevices{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegisteredDevicesRequestBuilder instantiates a new RegisteredDevicesRequestBuilder and sets the default values.
func NewRegisteredDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegisteredDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RegisteredDevicesRequestBuilder) Count()(*i0038fbe769ee74c442a37b3f1aa7c1f69e95198e07d90c1b3cc63b03e661a509.CountRequestBuilder) {
    return i0038fbe769ee74c442a37b3f1aa7c1f69e95198e07d90c1b3cc63b03e661a509.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *RegisteredDevicesRequestBuilder) Device()(*i319f64e4aebb40c6f2e1f6fc23df6af58292586d1bf495814fddf15918be7117.DeviceRequestBuilder) {
    return i319f64e4aebb40c6f2e1f6fc23df6af58292586d1bf495814fddf15918be7117.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoint the endpoint property
func (m *RegisteredDevicesRequestBuilder) Endpoint()(*i55e406c2335376386e0a9cfc801f8b8b57bc5ab242aed29c62fa441f5b8f7d8e.EndpointRequestBuilder) {
    return i55e406c2335376386e0a9cfc801f8b8b57bc5ab242aed29c62fa441f5b8f7d8e.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
