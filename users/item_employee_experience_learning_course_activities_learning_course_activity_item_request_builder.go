package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperienceUser entity.
type ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetQueryParameters get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user. This API is supported in the following national cloud deployments.
type ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetQueryParameters
}
// NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderInternal instantiates a new LearningCourseActivityItemRequestBuilder and sets the default values.
func NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) {
    m := &ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/employeeExperience/learningCourseActivities/{learningCourseActivity%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder instantiates a new LearningCourseActivityItemRequestBuilder and sets the default values.
func NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcourseactivity-get?view=graph-rest-1.0
func (m *ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningCourseActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable), nil
}
// ToGetRequestInformation get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user. This API is supported in the following national cloud deployments.
func (m *ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) WithUrl(rawUrl string)(*ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) {
    return NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
