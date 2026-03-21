# WlanGroupStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exceeded** | Pointer to **bool** | whether the number of WLAN groups exceeds the limit | [optional] 
**WlanGroupNum** | Pointer to **int64** | the number of wlan groups | [optional] 

## Methods

### NewWlanGroupStatusOpenApiVO

`func NewWlanGroupStatusOpenApiVO() *WlanGroupStatusOpenApiVO`

NewWlanGroupStatusOpenApiVO instantiates a new WlanGroupStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanGroupStatusOpenApiVOWithDefaults

`func NewWlanGroupStatusOpenApiVOWithDefaults() *WlanGroupStatusOpenApiVO`

NewWlanGroupStatusOpenApiVOWithDefaults instantiates a new WlanGroupStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceeded

`func (o *WlanGroupStatusOpenApiVO) GetExceeded() bool`

GetExceeded returns the Exceeded field if non-nil, zero value otherwise.

### GetExceededOk

`func (o *WlanGroupStatusOpenApiVO) GetExceededOk() (*bool, bool)`

GetExceededOk returns a tuple with the Exceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceeded

`func (o *WlanGroupStatusOpenApiVO) SetExceeded(v bool)`

SetExceeded sets Exceeded field to given value.

### HasExceeded

`func (o *WlanGroupStatusOpenApiVO) HasExceeded() bool`

HasExceeded returns a boolean if a field has been set.

### GetWlanGroupNum

`func (o *WlanGroupStatusOpenApiVO) GetWlanGroupNum() int64`

GetWlanGroupNum returns the WlanGroupNum field if non-nil, zero value otherwise.

### GetWlanGroupNumOk

`func (o *WlanGroupStatusOpenApiVO) GetWlanGroupNumOk() (*int64, bool)`

GetWlanGroupNumOk returns a tuple with the WlanGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanGroupNum

`func (o *WlanGroupStatusOpenApiVO) SetWlanGroupNum(v int64)`

SetWlanGroupNum sets WlanGroupNum field to given value.

### HasWlanGroupNum

`func (o *WlanGroupStatusOpenApiVO) HasWlanGroupNum() bool`

HasWlanGroupNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


