# IspResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isp** | Pointer to **string** | ISP name. | [optional] 
**IspNum** | Pointer to **int32** | The ID of the ISP. | [optional] 
**State** | Pointer to **bool** | Whether the ISP available or not. | [optional] 

## Methods

### NewIspResultOpenApiVO

`func NewIspResultOpenApiVO() *IspResultOpenApiVO`

NewIspResultOpenApiVO instantiates a new IspResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspResultOpenApiVOWithDefaults

`func NewIspResultOpenApiVOWithDefaults() *IspResultOpenApiVO`

NewIspResultOpenApiVOWithDefaults instantiates a new IspResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsp

`func (o *IspResultOpenApiVO) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *IspResultOpenApiVO) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *IspResultOpenApiVO) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *IspResultOpenApiVO) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIspNum

`func (o *IspResultOpenApiVO) GetIspNum() int32`

GetIspNum returns the IspNum field if non-nil, zero value otherwise.

### GetIspNumOk

`func (o *IspResultOpenApiVO) GetIspNumOk() (*int32, bool)`

GetIspNumOk returns a tuple with the IspNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspNum

`func (o *IspResultOpenApiVO) SetIspNum(v int32)`

SetIspNum sets IspNum field to given value.

### HasIspNum

`func (o *IspResultOpenApiVO) HasIspNum() bool`

HasIspNum returns a boolean if a field has been set.

### GetState

`func (o *IspResultOpenApiVO) GetState() bool`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IspResultOpenApiVO) GetStateOk() (*bool, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IspResultOpenApiVO) SetState(v bool)`

SetState sets State field to given value.

### HasState

`func (o *IspResultOpenApiVO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


