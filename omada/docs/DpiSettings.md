# DpiSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dpi** | **bool** | Enable dpi. true:enable / false:disable | 
**LoggingTraffic** | **bool** | Enable logging traffic. true:enable / false:disable | 

## Methods

### NewDpiSettings

`func NewDpiSettings(dpi bool, loggingTraffic bool, ) *DpiSettings`

NewDpiSettings instantiates a new DpiSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpiSettingsWithDefaults

`func NewDpiSettingsWithDefaults() *DpiSettings`

NewDpiSettingsWithDefaults instantiates a new DpiSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpi

`func (o *DpiSettings) GetDpi() bool`

GetDpi returns the Dpi field if non-nil, zero value otherwise.

### GetDpiOk

`func (o *DpiSettings) GetDpiOk() (*bool, bool)`

GetDpiOk returns a tuple with the Dpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpi

`func (o *DpiSettings) SetDpi(v bool)`

SetDpi sets Dpi field to given value.


### GetLoggingTraffic

`func (o *DpiSettings) GetLoggingTraffic() bool`

GetLoggingTraffic returns the LoggingTraffic field if non-nil, zero value otherwise.

### GetLoggingTrafficOk

`func (o *DpiSettings) GetLoggingTrafficOk() (*bool, bool)`

GetLoggingTrafficOk returns a tuple with the LoggingTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingTraffic

`func (o *DpiSettings) SetLoggingTraffic(v bool)`

SetLoggingTraffic sets LoggingTraffic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


