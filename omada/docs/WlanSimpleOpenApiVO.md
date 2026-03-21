# WlanSimpleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidList** | Pointer to [**[]SsidSimpleOpenApiVO**](SsidSimpleOpenApiVO.md) | SSID list that support MAC-Based authentication | [optional] 
**WlanId** | Pointer to **string** | ID of WLAN | [optional] 
**WlanName** | Pointer to **string** | Name of WLAN | [optional] 

## Methods

### NewWlanSimpleOpenApiVO

`func NewWlanSimpleOpenApiVO() *WlanSimpleOpenApiVO`

NewWlanSimpleOpenApiVO instantiates a new WlanSimpleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanSimpleOpenApiVOWithDefaults

`func NewWlanSimpleOpenApiVOWithDefaults() *WlanSimpleOpenApiVO`

NewWlanSimpleOpenApiVOWithDefaults instantiates a new WlanSimpleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidList

`func (o *WlanSimpleOpenApiVO) GetSsidList() []SsidSimpleOpenApiVO`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *WlanSimpleOpenApiVO) GetSsidListOk() (*[]SsidSimpleOpenApiVO, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *WlanSimpleOpenApiVO) SetSsidList(v []SsidSimpleOpenApiVO)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *WlanSimpleOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.

### GetWlanId

`func (o *WlanSimpleOpenApiVO) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *WlanSimpleOpenApiVO) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *WlanSimpleOpenApiVO) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *WlanSimpleOpenApiVO) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.

### GetWlanName

`func (o *WlanSimpleOpenApiVO) GetWlanName() string`

GetWlanName returns the WlanName field if non-nil, zero value otherwise.

### GetWlanNameOk

`func (o *WlanSimpleOpenApiVO) GetWlanNameOk() (*string, bool)`

GetWlanNameOk returns a tuple with the WlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanName

`func (o *WlanSimpleOpenApiVO) SetWlanName(v string)`

SetWlanName sets WlanName field to given value.

### HasWlanName

`func (o *WlanSimpleOpenApiVO) HasWlanName() bool`

HasWlanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


