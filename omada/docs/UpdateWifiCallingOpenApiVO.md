# UpdateWifiCallingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiCallingEnable** | **bool** | SSID Wi-Fi Calling global config status. True: enable, false: disable. | 
**WifiCallingId** | Pointer to **string** | This field represents Wi-Fi Calling Profile ID. Wi-Fi Calling Profile can be created using Create a new wifi calling profile. | [optional] 

## Methods

### NewUpdateWifiCallingOpenApiVO

`func NewUpdateWifiCallingOpenApiVO(wifiCallingEnable bool, ) *UpdateWifiCallingOpenApiVO`

NewUpdateWifiCallingOpenApiVO instantiates a new UpdateWifiCallingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWifiCallingOpenApiVOWithDefaults

`func NewUpdateWifiCallingOpenApiVOWithDefaults() *UpdateWifiCallingOpenApiVO`

NewUpdateWifiCallingOpenApiVOWithDefaults instantiates a new UpdateWifiCallingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiCallingEnable

`func (o *UpdateWifiCallingOpenApiVO) GetWifiCallingEnable() bool`

GetWifiCallingEnable returns the WifiCallingEnable field if non-nil, zero value otherwise.

### GetWifiCallingEnableOk

`func (o *UpdateWifiCallingOpenApiVO) GetWifiCallingEnableOk() (*bool, bool)`

GetWifiCallingEnableOk returns a tuple with the WifiCallingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiCallingEnable

`func (o *UpdateWifiCallingOpenApiVO) SetWifiCallingEnable(v bool)`

SetWifiCallingEnable sets WifiCallingEnable field to given value.


### GetWifiCallingId

`func (o *UpdateWifiCallingOpenApiVO) GetWifiCallingId() string`

GetWifiCallingId returns the WifiCallingId field if non-nil, zero value otherwise.

### GetWifiCallingIdOk

`func (o *UpdateWifiCallingOpenApiVO) GetWifiCallingIdOk() (*string, bool)`

GetWifiCallingIdOk returns a tuple with the WifiCallingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiCallingId

`func (o *UpdateWifiCallingOpenApiVO) SetWifiCallingId(v string)`

SetWifiCallingId sets WifiCallingId field to given value.

### HasWifiCallingId

`func (o *UpdateWifiCallingOpenApiVO) HasWifiCallingId() bool`

HasWifiCallingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


