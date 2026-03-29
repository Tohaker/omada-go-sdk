# NotificationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAlerts** | Pointer to [**NotificationConfigurationOpenApiVO**](NotificationConfigurationOpenApiVO.md) |  | [optional] 
**AlertEmailSetting** | Pointer to [**AlertEmailSettingVO**](AlertEmailSettingVO.md) |  | [optional] 
**Resource** | Pointer to **int32** | The incident notifiction setting creation resource, such as: 0: new created, 1: from template, 2: override. | [optional] 

## Methods

### NewNotificationOpenApiVO

`func NewNotificationOpenApiVO() *NotificationOpenApiVO`

NewNotificationOpenApiVO instantiates a new NotificationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationOpenApiVOWithDefaults

`func NewNotificationOpenApiVOWithDefaults() *NotificationOpenApiVO`

NewNotificationOpenApiVOWithDefaults instantiates a new NotificationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAlerts

`func (o *NotificationOpenApiVO) GetEmailAlerts() NotificationConfigurationOpenApiVO`

GetEmailAlerts returns the EmailAlerts field if non-nil, zero value otherwise.

### GetEmailAlertsOk

`func (o *NotificationOpenApiVO) GetEmailAlertsOk() (*NotificationConfigurationOpenApiVO, bool)`

GetEmailAlertsOk returns a tuple with the EmailAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlerts

`func (o *NotificationOpenApiVO) SetEmailAlerts(v NotificationConfigurationOpenApiVO)`

SetEmailAlerts sets EmailAlerts field to given value.

### HasEmailAlerts

`func (o *NotificationOpenApiVO) HasEmailAlerts() bool`

HasEmailAlerts returns a boolean if a field has been set.

### GetAlertEmailSetting

`func (o *NotificationOpenApiVO) GetAlertEmailSetting() AlertEmailSettingVO`

GetAlertEmailSetting returns the AlertEmailSetting field if non-nil, zero value otherwise.

### GetAlertEmailSettingOk

`func (o *NotificationOpenApiVO) GetAlertEmailSettingOk() (*AlertEmailSettingVO, bool)`

GetAlertEmailSettingOk returns a tuple with the AlertEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailSetting

`func (o *NotificationOpenApiVO) SetAlertEmailSetting(v AlertEmailSettingVO)`

SetAlertEmailSetting sets AlertEmailSetting field to given value.

### HasAlertEmailSetting

`func (o *NotificationOpenApiVO) HasAlertEmailSetting() bool`

HasAlertEmailSetting returns a boolean if a field has been set.

### GetResource

`func (o *NotificationOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *NotificationOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *NotificationOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *NotificationOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


