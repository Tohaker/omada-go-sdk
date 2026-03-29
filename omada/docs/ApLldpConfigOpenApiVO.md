# ApLldpConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LldpEnable** | Pointer to **int32** | Parameter [lldpEnable] should be a value as follows: 0:off; 1:on; 2:Use Site Settings. | [optional] 
**SupportLldp** | Pointer to **bool** | Whether the lldp function is supported. | [optional] 

## Methods

### NewApLldpConfigOpenApiVO

`func NewApLldpConfigOpenApiVO() *ApLldpConfigOpenApiVO`

NewApLldpConfigOpenApiVO instantiates a new ApLldpConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApLldpConfigOpenApiVOWithDefaults

`func NewApLldpConfigOpenApiVOWithDefaults() *ApLldpConfigOpenApiVO`

NewApLldpConfigOpenApiVOWithDefaults instantiates a new ApLldpConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldpEnable

`func (o *ApLldpConfigOpenApiVO) GetLldpEnable() int32`

GetLldpEnable returns the LldpEnable field if non-nil, zero value otherwise.

### GetLldpEnableOk

`func (o *ApLldpConfigOpenApiVO) GetLldpEnableOk() (*int32, bool)`

GetLldpEnableOk returns a tuple with the LldpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnable

`func (o *ApLldpConfigOpenApiVO) SetLldpEnable(v int32)`

SetLldpEnable sets LldpEnable field to given value.

### HasLldpEnable

`func (o *ApLldpConfigOpenApiVO) HasLldpEnable() bool`

HasLldpEnable returns a boolean if a field has been set.

### GetSupportLldp

`func (o *ApLldpConfigOpenApiVO) GetSupportLldp() bool`

GetSupportLldp returns the SupportLldp field if non-nil, zero value otherwise.

### GetSupportLldpOk

`func (o *ApLldpConfigOpenApiVO) GetSupportLldpOk() (*bool, bool)`

GetSupportLldpOk returns a tuple with the SupportLldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLldp

`func (o *ApLldpConfigOpenApiVO) SetSupportLldp(v bool)`

SetSupportLldp sets SupportLldp field to given value.

### HasSupportLldp

`func (o *ApLldpConfigOpenApiVO) HasSupportLldp() bool`

HasSupportLldp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


