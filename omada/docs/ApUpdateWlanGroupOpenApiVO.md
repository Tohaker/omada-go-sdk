# ApUpdateWlanGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WlanGroupId** | **string** | The wlan group Id that the AP should switch to, must be in the same site as the AP and cannot be the current wlan group. | 

## Methods

### NewApUpdateWlanGroupOpenApiVO

`func NewApUpdateWlanGroupOpenApiVO(wlanGroupId string, ) *ApUpdateWlanGroupOpenApiVO`

NewApUpdateWlanGroupOpenApiVO instantiates a new ApUpdateWlanGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApUpdateWlanGroupOpenApiVOWithDefaults

`func NewApUpdateWlanGroupOpenApiVOWithDefaults() *ApUpdateWlanGroupOpenApiVO`

NewApUpdateWlanGroupOpenApiVOWithDefaults instantiates a new ApUpdateWlanGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWlanGroupId

`func (o *ApUpdateWlanGroupOpenApiVO) GetWlanGroupId() string`

GetWlanGroupId returns the WlanGroupId field if non-nil, zero value otherwise.

### GetWlanGroupIdOk

`func (o *ApUpdateWlanGroupOpenApiVO) GetWlanGroupIdOk() (*string, bool)`

GetWlanGroupIdOk returns a tuple with the WlanGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanGroupId

`func (o *ApUpdateWlanGroupOpenApiVO) SetWlanGroupId(v string)`

SetWlanGroupId sets WlanGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


