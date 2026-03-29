# ModifyApAdvancedQosOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryEnable2g** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 2.4GHz. | [optional] 
**DeliveryEnable5g** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in the entire band of 5GHz. If the device supports 5GHz but does not support frequency division, you need to enter this field. | [optional] 
**DeliveryEnable5g1** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 5GHz-1. 5GHZ-1 is the first part of the 5GHz band. This field is required if the device supports 5GHz and frequency splitting. | [optional] 
**DeliveryEnable5g2** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 5GHz-2. 5GHZ-2 is the second part of the 5GHz band. This field is required if the device supports 5GHz and frequency splitting. | [optional] 
**DeliveryEnable6g** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 6GHz. | [optional] 

## Methods

### NewModifyApAdvancedQosOpenApiVO

`func NewModifyApAdvancedQosOpenApiVO() *ModifyApAdvancedQosOpenApiVO`

NewModifyApAdvancedQosOpenApiVO instantiates a new ModifyApAdvancedQosOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyApAdvancedQosOpenApiVOWithDefaults

`func NewModifyApAdvancedQosOpenApiVOWithDefaults() *ModifyApAdvancedQosOpenApiVO`

NewModifyApAdvancedQosOpenApiVOWithDefaults instantiates a new ModifyApAdvancedQosOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryEnable2g

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable2g() bool`

GetDeliveryEnable2g returns the DeliveryEnable2g field if non-nil, zero value otherwise.

### GetDeliveryEnable2gOk

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable2gOk() (*bool, bool)`

GetDeliveryEnable2gOk returns a tuple with the DeliveryEnable2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable2g

`func (o *ModifyApAdvancedQosOpenApiVO) SetDeliveryEnable2g(v bool)`

SetDeliveryEnable2g sets DeliveryEnable2g field to given value.

### HasDeliveryEnable2g

`func (o *ModifyApAdvancedQosOpenApiVO) HasDeliveryEnable2g() bool`

HasDeliveryEnable2g returns a boolean if a field has been set.

### GetDeliveryEnable5g

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable5g() bool`

GetDeliveryEnable5g returns the DeliveryEnable5g field if non-nil, zero value otherwise.

### GetDeliveryEnable5gOk

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable5gOk() (*bool, bool)`

GetDeliveryEnable5gOk returns a tuple with the DeliveryEnable5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable5g

`func (o *ModifyApAdvancedQosOpenApiVO) SetDeliveryEnable5g(v bool)`

SetDeliveryEnable5g sets DeliveryEnable5g field to given value.

### HasDeliveryEnable5g

`func (o *ModifyApAdvancedQosOpenApiVO) HasDeliveryEnable5g() bool`

HasDeliveryEnable5g returns a boolean if a field has been set.

### GetDeliveryEnable5g1

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable5g1() bool`

GetDeliveryEnable5g1 returns the DeliveryEnable5g1 field if non-nil, zero value otherwise.

### GetDeliveryEnable5g1Ok

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable5g1Ok() (*bool, bool)`

GetDeliveryEnable5g1Ok returns a tuple with the DeliveryEnable5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable5g1

`func (o *ModifyApAdvancedQosOpenApiVO) SetDeliveryEnable5g1(v bool)`

SetDeliveryEnable5g1 sets DeliveryEnable5g1 field to given value.

### HasDeliveryEnable5g1

`func (o *ModifyApAdvancedQosOpenApiVO) HasDeliveryEnable5g1() bool`

HasDeliveryEnable5g1 returns a boolean if a field has been set.

### GetDeliveryEnable5g2

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable5g2() bool`

GetDeliveryEnable5g2 returns the DeliveryEnable5g2 field if non-nil, zero value otherwise.

### GetDeliveryEnable5g2Ok

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable5g2Ok() (*bool, bool)`

GetDeliveryEnable5g2Ok returns a tuple with the DeliveryEnable5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable5g2

`func (o *ModifyApAdvancedQosOpenApiVO) SetDeliveryEnable5g2(v bool)`

SetDeliveryEnable5g2 sets DeliveryEnable5g2 field to given value.

### HasDeliveryEnable5g2

`func (o *ModifyApAdvancedQosOpenApiVO) HasDeliveryEnable5g2() bool`

HasDeliveryEnable5g2 returns a boolean if a field has been set.

### GetDeliveryEnable6g

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable6g() bool`

GetDeliveryEnable6g returns the DeliveryEnable6g field if non-nil, zero value otherwise.

### GetDeliveryEnable6gOk

`func (o *ModifyApAdvancedQosOpenApiVO) GetDeliveryEnable6gOk() (*bool, bool)`

GetDeliveryEnable6gOk returns a tuple with the DeliveryEnable6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable6g

`func (o *ModifyApAdvancedQosOpenApiVO) SetDeliveryEnable6g(v bool)`

SetDeliveryEnable6g sets DeliveryEnable6g field to given value.

### HasDeliveryEnable6g

`func (o *ModifyApAdvancedQosOpenApiVO) HasDeliveryEnable6g() bool`

HasDeliveryEnable6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


