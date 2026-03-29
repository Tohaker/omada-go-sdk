# SetNotificationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAlerts** | [**NotificationConfigurationOpenApiVO**](NotificationConfigurationOpenApiVO.md) |  | 
**AlertEmailSetting** | [**AlertEmailSettingVO**](AlertEmailSettingVO.md) |  | 

## Methods

### NewSetNotificationOpenApiVO

`func NewSetNotificationOpenApiVO(emailAlerts NotificationConfigurationOpenApiVO, alertEmailSetting AlertEmailSettingVO, ) *SetNotificationOpenApiVO`

NewSetNotificationOpenApiVO instantiates a new SetNotificationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNotificationOpenApiVOWithDefaults

`func NewSetNotificationOpenApiVOWithDefaults() *SetNotificationOpenApiVO`

NewSetNotificationOpenApiVOWithDefaults instantiates a new SetNotificationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAlerts

`func (o *SetNotificationOpenApiVO) GetEmailAlerts() NotificationConfigurationOpenApiVO`

GetEmailAlerts returns the EmailAlerts field if non-nil, zero value otherwise.

### GetEmailAlertsOk

`func (o *SetNotificationOpenApiVO) GetEmailAlertsOk() (*NotificationConfigurationOpenApiVO, bool)`

GetEmailAlertsOk returns a tuple with the EmailAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlerts

`func (o *SetNotificationOpenApiVO) SetEmailAlerts(v NotificationConfigurationOpenApiVO)`

SetEmailAlerts sets EmailAlerts field to given value.


### GetAlertEmailSetting

`func (o *SetNotificationOpenApiVO) GetAlertEmailSetting() AlertEmailSettingVO`

GetAlertEmailSetting returns the AlertEmailSetting field if non-nil, zero value otherwise.

### GetAlertEmailSettingOk

`func (o *SetNotificationOpenApiVO) GetAlertEmailSettingOk() (*AlertEmailSettingVO, bool)`

GetAlertEmailSettingOk returns a tuple with the AlertEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailSetting

`func (o *SetNotificationOpenApiVO) SetAlertEmailSetting(v AlertEmailSettingVO)`

SetAlertEmailSetting sets AlertEmailSetting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


