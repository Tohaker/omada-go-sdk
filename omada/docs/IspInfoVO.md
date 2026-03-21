# IspInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IspArr** | Pointer to [**[]IspVO**](IspVO.md) | Isp info detail. | [optional] 
**Type** | Pointer to **int32** | Isp type, should be a value as follows:0 : single ISP1 : multiple ISPs | [optional] 

## Methods

### NewIspInfoVO

`func NewIspInfoVO() *IspInfoVO`

NewIspInfoVO instantiates a new IspInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspInfoVOWithDefaults

`func NewIspInfoVOWithDefaults() *IspInfoVO`

NewIspInfoVOWithDefaults instantiates a new IspInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIspArr

`func (o *IspInfoVO) GetIspArr() []IspVO`

GetIspArr returns the IspArr field if non-nil, zero value otherwise.

### GetIspArrOk

`func (o *IspInfoVO) GetIspArrOk() (*[]IspVO, bool)`

GetIspArrOk returns a tuple with the IspArr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspArr

`func (o *IspInfoVO) SetIspArr(v []IspVO)`

SetIspArr sets IspArr field to given value.

### HasIspArr

`func (o *IspInfoVO) HasIspArr() bool`

HasIspArr returns a boolean if a field has been set.

### GetType

`func (o *IspInfoVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IspInfoVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IspInfoVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *IspInfoVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


