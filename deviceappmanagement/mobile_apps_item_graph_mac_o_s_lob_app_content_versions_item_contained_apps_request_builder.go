package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
type MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetQueryParameters the collection of contained apps in a MobileLobApp acting as a package.
type MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetQueryParameters struct {
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
// MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileContainedAppId provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) ByMobileContainedAppId(mobileContainedAppId string)(*MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileContainedAppId != "" {
        urlTplParams["mobileContainedApp%2Did"] = mobileContainedAppId
    }
    return NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderInternal instantiates a new ContainedAppsRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) {
    m := &MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.macOSLobApp/contentVersions/{mobileAppContent%2Did}/containedApps{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder instantiates a new ContainedAppsRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) Count()(*MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsCountRequestBuilder) {
    return NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of contained apps in a MobileLobApp acting as a package.
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileContainedAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppCollectionResponseable), nil
}
// Post create new navigation property to containedApps for deviceAppManagement
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, requestConfiguration *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileContainedAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable), nil
}
// ToGetRequestInformation the collection of contained apps in a MobileLobApp acting as a package.
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to containedApps for deviceAppManagement
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, requestConfiguration *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder) {
    return NewMobileAppsItemGraphMacOSLobAppContentVersionsItemContainedAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
