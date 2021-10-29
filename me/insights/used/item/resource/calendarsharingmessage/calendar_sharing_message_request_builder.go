package calendarsharingmessage

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2721ccc1481760ff258835b632d1b9662c687785bdb8c05732db15682548aa29 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/calendarsharingmessage/accept"
)

// Builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.calendarSharingMessage
type CalendarSharingMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *CalendarSharingMessageRequestBuilder) Accept()(*i2721ccc1481760ff258835b632d1b9662c687785bdb8c05732db15682548aa29.AcceptRequestBuilder) {
    return i2721ccc1481760ff258835b632d1b9662c687785bdb8c05732db15682548aa29.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/used/{usedInsight_id}/resource/microsoft.graph.calendarSharingMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarSharingMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarSharingMessageRequestBuilderInternal(urlParams, requestAdapter)
}
