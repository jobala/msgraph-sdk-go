package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1c05393909ced8c970ac8182fc4a6c5cc419883d0d1475ce0773c5cf783ba989 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/listitem"
    i353c71cf4e7f00f79bca7c43c481c2b7fccfaf2bc763be7f85e457bcdb2a832b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/content"
    i35ec8d39bca25c3225ac281fac7242810cb642a0a53e14a875e33c70b99c171e "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/analytics"
    i6332f0883e6ab0e66763410aa3cab286bc78b7b61f88fefedb1be3fb930832df "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/thumbnails"
    i6a155f272c5391daf639268076f7e78fb66358ef654a4f57db99ce00a32f3c11 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children"
    i8586e46c09b38f90b06f86eb42fd10319ce6acb9451aefe28f39853cfb5719b9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/versions"
    ia05f97a343c7dbf32d5d42d3b714d8d4e180ebdebfee2e83c7a95ec1ed6e9e3e "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/permissions"
    ic72fcde28010dbd679eab351c33b57107adc23c1cc48e905dd109128ec3bc4db "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/subscriptions"
    i162b49d374c91a27c7d63a8c7103a36e56411b6b3e84fc64ecbf5d828ae8a8af "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/versions/item"
    i3d94da0dddfc16aca3f8f4350e9b0d9d0a38bc1ffb67f012ae9cb026450ac2a1 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/subscriptions/item"
    i896e3d0df31741011967f273d16af8fb1abe2a50bd0bb2f742e0aead26932993 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/thumbnails/item"
    i8ca583e1745c18b27ce93c86c752cf4c03f78604e1970006d1886875c01639c7 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/permissions/item"
    ic7d208ecab0cc1c68790901ba2ba97554ff0a0a86b46da4f6d1dba8ab740987f "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
}
// RootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i35ec8d39bca25c3225ac281fac7242810cb642a0a53e14a875e33c70b99c171e.AnalyticsRequestBuilder) {
    return i35ec8d39bca25c3225ac281fac7242810cb642a0a53e14a875e33c70b99c171e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i6a155f272c5391daf639268076f7e78fb66358ef654a4f57db99ce00a32f3c11.ChildrenRequestBuilder) {
    return i6a155f272c5391daf639268076f7e78fb66358ef654a4f57db99ce00a32f3c11.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ic7d208ecab0cc1c68790901ba2ba97554ff0a0a86b46da4f6d1dba8ab740987f.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ic7d208ecab0cc1c68790901ba2ba97554ff0a0a86b46da4f6d1dba8ab740987f.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *RootRequestBuilder) Content()(*i353c71cf4e7f00f79bca7c43c481c2b7fccfaf2bc763be7f85e457bcdb2a832b.ContentRequestBuilder) {
    return i353c71cf4e7f00f79bca7c43c481c2b7fccfaf2bc763be7f85e457bcdb2a832b.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drives
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for drives
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property root in drives
func (m *RootRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in drives
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property root for drives
func (m *RootRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property root for drives
func (m *RootRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the root folder of the drive. Read-only.
func (m *RootRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*i1c05393909ced8c970ac8182fc4a6c5cc419883d0d1475ce0773c5cf783ba989.ListItemRequestBuilder) {
    return i1c05393909ced8c970ac8182fc4a6c5cc419883d0d1475ce0773c5cf783ba989.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drives
func (m *RootRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property root in drives
func (m *RootRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Permissions the permissions property
func (m *RootRequestBuilder) Permissions()(*ia05f97a343c7dbf32d5d42d3b714d8d4e180ebdebfee2e83c7a95ec1ed6e9e3e.PermissionsRequestBuilder) {
    return ia05f97a343c7dbf32d5d42d3b714d8d4e180ebdebfee2e83c7a95ec1ed6e9e3e.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i8ca583e1745c18b27ce93c86c752cf4c03f78604e1970006d1886875c01639c7.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i8ca583e1745c18b27ce93c86c752cf4c03f78604e1970006d1886875c01639c7.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*ic72fcde28010dbd679eab351c33b57107adc23c1cc48e905dd109128ec3bc4db.SubscriptionsRequestBuilder) {
    return ic72fcde28010dbd679eab351c33b57107adc23c1cc48e905dd109128ec3bc4db.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i3d94da0dddfc16aca3f8f4350e9b0d9d0a38bc1ffb67f012ae9cb026450ac2a1.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i3d94da0dddfc16aca3f8f4350e9b0d9d0a38bc1ffb67f012ae9cb026450ac2a1.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i6332f0883e6ab0e66763410aa3cab286bc78b7b61f88fefedb1be3fb930832df.ThumbnailsRequestBuilder) {
    return i6332f0883e6ab0e66763410aa3cab286bc78b7b61f88fefedb1be3fb930832df.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i896e3d0df31741011967f273d16af8fb1abe2a50bd0bb2f742e0aead26932993.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i896e3d0df31741011967f273d16af8fb1abe2a50bd0bb2f742e0aead26932993.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i8586e46c09b38f90b06f86eb42fd10319ce6acb9451aefe28f39853cfb5719b9.VersionsRequestBuilder) {
    return i8586e46c09b38f90b06f86eb42fd10319ce6acb9451aefe28f39853cfb5719b9.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i162b49d374c91a27c7d63a8c7103a36e56411b6b3e84fc64ecbf5d828ae8a8af.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i162b49d374c91a27c7d63a8c7103a36e56411b6b3e84fc64ecbf5d828ae8a8af.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
