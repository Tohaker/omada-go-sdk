# NotificationConfigurationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventCategory** | **map[string]bool** | For the values of event category Key, refer to section 5.7.1.1 of the Open API Access. | 
**EventLevel** | **map[string]bool** | The incident level of event, such as: critical, error, warning, info. | 
**EventObjectType** | **map[string]bool** | The incident object of event, such as: gateway, switch, ap, wiredClient, wirelessClient. | 

## Methods

### NewNotificationConfigurationOpenApiVO

`func NewNotificationConfigurationOpenApiVO(eventCategory map[string]bool, eventLevel map[string]bool, eventObjectType map[string]bool, ) *NotificationConfigurationOpenApiVO`

NewNotificationConfigurationOpenApiVO instantiates a new NotificationConfigurationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigurationOpenApiVOWithDefaults

`func NewNotificationConfigurationOpenApiVOWithDefaults() *NotificationConfigurationOpenApiVO`

NewNotificationConfigurationOpenApiVOWithDefaults instantiates a new NotificationConfigurationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventCategory

`func (o *NotificationConfigurationOpenApiVO) GetEventCategory() map[string]bool`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *NotificationConfigurationOpenApiVO) GetEventCategoryOk() (*map[string]bool, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *NotificationConfigurationOpenApiVO) SetEventCategory(v map[string]bool)`

SetEventCategory sets EventCategory field to given value.


### GetEventLevel

`func (o *NotificationConfigurationOpenApiVO) GetEventLevel() map[string]bool`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *NotificationConfigurationOpenApiVO) GetEventLevelOk() (*map[string]bool, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *NotificationConfigurationOpenApiVO) SetEventLevel(v map[string]bool)`

SetEventLevel sets EventLevel field to given value.


### GetEventObjectType

`func (o *NotificationConfigurationOpenApiVO) GetEventObjectType() map[string]bool`

GetEventObjectType returns the EventObjectType field if non-nil, zero value otherwise.

### GetEventObjectTypeOk

`func (o *NotificationConfigurationOpenApiVO) GetEventObjectTypeOk() (*map[string]bool, bool)`

GetEventObjectTypeOk returns a tuple with the EventObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventObjectType

`func (o *NotificationConfigurationOpenApiVO) SetEventObjectType(v map[string]bool)`

SetEventObjectType sets EventObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


