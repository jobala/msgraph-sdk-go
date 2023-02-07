package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder provides operations to call the upgrade method.
type ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderInternal instantiates a new MicrosoftGraphUpgradeRequestBuilder and sets the default values.
func NewChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder) {
    m := &ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}/installedApps/{teamsAppInstallation%2Did}/microsoft.graph.upgrade";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder instantiates a new MicrosoftGraphUpgradeRequestBuilder and sets the default values.
func NewChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upgrade an app installation within a chat.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-teamsappinstallation-upgrade?view=graph-rest-1.0
func (m *ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder) Post(ctx context.Context, requestConfiguration *ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation upgrade an app installation within a chat.
func (m *ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ChatsItemInstalledAppsItemMicrosoftGraphUpgradeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
