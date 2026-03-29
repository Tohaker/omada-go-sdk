# ApAvailableChannelOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApChannelDetailList** | Pointer to [**[]ApChannelDetailOpenApiVO**](ApChannelDetailOpenApiVO.md) | channels detail supported by device configuration. | [optional] 
**RadioId** | Pointer to **int32** | (Wireless) Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz-1; 2: 5GHz-2; 3: 6GHz. | [optional] 

## Methods

### NewApAvailableChannelOpenApiVO

`func NewApAvailableChannelOpenApiVO() *ApAvailableChannelOpenApiVO`

NewApAvailableChannelOpenApiVO instantiates a new ApAvailableChannelOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApAvailableChannelOpenApiVOWithDefaults

`func NewApAvailableChannelOpenApiVOWithDefaults() *ApAvailableChannelOpenApiVO`

NewApAvailableChannelOpenApiVOWithDefaults instantiates a new ApAvailableChannelOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApChannelDetailList

`func (o *ApAvailableChannelOpenApiVO) GetApChannelDetailList() []ApChannelDetailOpenApiVO`

GetApChannelDetailList returns the ApChannelDetailList field if non-nil, zero value otherwise.

### GetApChannelDetailListOk

`func (o *ApAvailableChannelOpenApiVO) GetApChannelDetailListOk() (*[]ApChannelDetailOpenApiVO, bool)`

GetApChannelDetailListOk returns a tuple with the ApChannelDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApChannelDetailList

`func (o *ApAvailableChannelOpenApiVO) SetApChannelDetailList(v []ApChannelDetailOpenApiVO)`

SetApChannelDetailList sets ApChannelDetailList field to given value.

### HasApChannelDetailList

`func (o *ApAvailableChannelOpenApiVO) HasApChannelDetailList() bool`

HasApChannelDetailList returns a boolean if a field has been set.

### GetRadioId

`func (o *ApAvailableChannelOpenApiVO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ApAvailableChannelOpenApiVO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ApAvailableChannelOpenApiVO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ApAvailableChannelOpenApiVO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


