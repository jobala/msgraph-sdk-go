package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1b9f51a74b0ef03db9eeb60264b1432cd3a3e3f24085bbbbb757fa729fc1a0c5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/targetedmanagedappprotection/assign"
    i6c699c723b64aa9f1ef26cdc2a9ccfb6a343bcf4db6b69ab2fa231c7e70755d0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/targetedmanagedappprotection/targetapps"
)

// Builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i1b9f51a74b0ef03db9eeb60264b1432cd3a3e3f24085bbbbb757fa729fc1a0c5.AssignRequestBuilder) {
    return i1b9f51a74b0ef03db9eeb60264b1432cd3a3e3f24085bbbbb757fa729fc1a0c5.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/trending/{trending_id}/resource/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i6c699c723b64aa9f1ef26cdc2a9ccfb6a343bcf4db6b69ab2fa231c7e70755d0.TargetAppsRequestBuilder) {
    return i6c699c723b64aa9f1ef26cdc2a9ccfb6a343bcf4db6b69ab2fa231c7e70755d0.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
