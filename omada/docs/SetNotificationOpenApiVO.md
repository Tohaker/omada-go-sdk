# SetNotificationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmailSetting** | [**AlertEmailSettingVO**](AlertEmailSettingVO.md) |  | 
**Email** | [**NotificationConfigurationOpenApiVO**](NotificationConfigurationOpenApiVO.md) |  | 

## Methods

### NewSetNotificationOpenApiVO

`func NewSetNotificationOpenApiVO(alertEmailSetting AlertEmailSettingVO, email NotificationConfigurationOpenApiVO, ) *SetNotificationOpenApiVO`

NewSetNotificationOpenApiVO instantiates a new SetNotificationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNotificationOpenApiVOWithDefaults

`func NewSetNotificationOpenApiVOWithDefaults() *SetNotificationOpenApiVO`

NewSetNotificationOpenApiVOWithDefaults instantiates a new SetNotificationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetEmail

`func (o *SetNotificationOpenApiVO) GetEmail() NotificationConfigurationOpenApiVO`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SetNotificationOpenApiVO) GetEmailOk() (*NotificationConfigurationOpenApiVO, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SetNotificationOpenApiVO) SetEmail(v NotificationConfigurationOpenApiVO)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


