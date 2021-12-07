package columnsbeforewithcount

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ColumnsBeforeWithCountRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
type ColumnsBeforeWithCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ColumnsBeforeWithCountRequestBuilderGetOptions options for Get
type ColumnsBeforeWithCountRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ColumnsBeforeWithCountResponse union type wrapper for classes workbookRange
type ColumnsBeforeWithCountResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type workbookRange
    workbookRange *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange;
}
// NewColumnsBeforeWithCountResponse instantiates a new columnsBeforeWithCountResponse and sets the default values.
func NewColumnsBeforeWithCountResponse()(*ColumnsBeforeWithCountResponse) {
    m := &ColumnsBeforeWithCountResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnsBeforeWithCountResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetWorkbookRange gets the workbookRange property value. Union type representation for type workbookRange
func (m *ColumnsBeforeWithCountResponse) GetWorkbookRange()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange) {
    if m == nil {
        return nil
    } else {
        return m.workbookRange
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnsBeforeWithCountResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["workbookRange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbookRange(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange))
        }
        return nil
    }
    return res
}
func (m *ColumnsBeforeWithCountResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ColumnsBeforeWithCountResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("workbookRange", m.GetWorkbookRange())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnsBeforeWithCountResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetWorkbookRange sets the workbookRange property value. Union type representation for type workbookRange
func (m *ColumnsBeforeWithCountResponse) SetWorkbookRange(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange)() {
    if m != nil {
        m.workbookRange = value
    }
}
// NewColumnsBeforeWithCountRequestBuilderInternal instantiates a new ColumnsBeforeWithCountRequestBuilder and sets the default values.
func NewColumnsBeforeWithCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, count *int32)(*ColumnsBeforeWithCountRequestBuilder) {
    m := &ColumnsBeforeWithCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.workbookRange/microsoft.graph.columnsBefore(count={count})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if count != nil {
        urlTplParams["count"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*count), 10)
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewColumnsBeforeWithCountRequestBuilder instantiates a new ColumnsBeforeWithCountRequestBuilder and sets the default values.
func NewColumnsBeforeWithCountRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ColumnsBeforeWithCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewColumnsBeforeWithCountRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function columnsBefore
func (m *ColumnsBeforeWithCountRequestBuilder) CreateGetRequestInformation(options *ColumnsBeforeWithCountRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function columnsBefore
func (m *ColumnsBeforeWithCountRequestBuilder) Get(options *ColumnsBeforeWithCountRequestBuilderGetOptions)(*ColumnsBeforeWithCountResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnsBeforeWithCountResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ColumnsBeforeWithCountResponse), nil
}
