# ApAdvancedQosOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryEnable2g** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 2.4GHz. | [optional] 
**DeliveryEnable5g** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 5GHz. The field [deliveryEnable5g] will be filled in when the device supports 5GHz. If the device supports frequency division, the field indicates 5GHZ-1.(5GHZ-1 is the second part of the 5GHz band) If the device does not support frequency division, the field indicates the entire band of 5GHz. | [optional] 
**DeliveryEnable5g2** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 5GHz-2.(5GHZ-2 is the second part of the 5GHz band) The field [deliveryEnable5g2] will be filled in when the device supports 5GHz and frequency splitting. | [optional] 
**DeliveryEnable6g** | Pointer to **bool** | Whether Unscheduled Automatic Power Save Delivery enabled in 6GHz. | [optional] 
**Support2g** | Pointer to **bool** | Whether the device supports 2.4GHz. | [optional] 
**Support5g** | Pointer to **bool** | Whether the device supports 5GHz. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the device supports 5GHz frequency splitting into two parts. | [optional] 
**Support6g** | Pointer to **bool** | Whether the device supports 6GHz. | [optional] 

## Methods

### NewApAdvancedQosOpenApiVO

`func NewApAdvancedQosOpenApiVO() *ApAdvancedQosOpenApiVO`

NewApAdvancedQosOpenApiVO instantiates a new ApAdvancedQosOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApAdvancedQosOpenApiVOWithDefaults

`func NewApAdvancedQosOpenApiVOWithDefaults() *ApAdvancedQosOpenApiVO`

NewApAdvancedQosOpenApiVOWithDefaults instantiates a new ApAdvancedQosOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryEnable2g

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable2g() bool`

GetDeliveryEnable2g returns the DeliveryEnable2g field if non-nil, zero value otherwise.

### GetDeliveryEnable2gOk

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable2gOk() (*bool, bool)`

GetDeliveryEnable2gOk returns a tuple with the DeliveryEnable2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable2g

`func (o *ApAdvancedQosOpenApiVO) SetDeliveryEnable2g(v bool)`

SetDeliveryEnable2g sets DeliveryEnable2g field to given value.

### HasDeliveryEnable2g

`func (o *ApAdvancedQosOpenApiVO) HasDeliveryEnable2g() bool`

HasDeliveryEnable2g returns a boolean if a field has been set.

### GetDeliveryEnable5g

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable5g() bool`

GetDeliveryEnable5g returns the DeliveryEnable5g field if non-nil, zero value otherwise.

### GetDeliveryEnable5gOk

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable5gOk() (*bool, bool)`

GetDeliveryEnable5gOk returns a tuple with the DeliveryEnable5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable5g

`func (o *ApAdvancedQosOpenApiVO) SetDeliveryEnable5g(v bool)`

SetDeliveryEnable5g sets DeliveryEnable5g field to given value.

### HasDeliveryEnable5g

`func (o *ApAdvancedQosOpenApiVO) HasDeliveryEnable5g() bool`

HasDeliveryEnable5g returns a boolean if a field has been set.

### GetDeliveryEnable5g2

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable5g2() bool`

GetDeliveryEnable5g2 returns the DeliveryEnable5g2 field if non-nil, zero value otherwise.

### GetDeliveryEnable5g2Ok

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable5g2Ok() (*bool, bool)`

GetDeliveryEnable5g2Ok returns a tuple with the DeliveryEnable5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable5g2

`func (o *ApAdvancedQosOpenApiVO) SetDeliveryEnable5g2(v bool)`

SetDeliveryEnable5g2 sets DeliveryEnable5g2 field to given value.

### HasDeliveryEnable5g2

`func (o *ApAdvancedQosOpenApiVO) HasDeliveryEnable5g2() bool`

HasDeliveryEnable5g2 returns a boolean if a field has been set.

### GetDeliveryEnable6g

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable6g() bool`

GetDeliveryEnable6g returns the DeliveryEnable6g field if non-nil, zero value otherwise.

### GetDeliveryEnable6gOk

`func (o *ApAdvancedQosOpenApiVO) GetDeliveryEnable6gOk() (*bool, bool)`

GetDeliveryEnable6gOk returns a tuple with the DeliveryEnable6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEnable6g

`func (o *ApAdvancedQosOpenApiVO) SetDeliveryEnable6g(v bool)`

SetDeliveryEnable6g sets DeliveryEnable6g field to given value.

### HasDeliveryEnable6g

`func (o *ApAdvancedQosOpenApiVO) HasDeliveryEnable6g() bool`

HasDeliveryEnable6g returns a boolean if a field has been set.

### GetSupport2g

`func (o *ApAdvancedQosOpenApiVO) GetSupport2g() bool`

GetSupport2g returns the Support2g field if non-nil, zero value otherwise.

### GetSupport2gOk

`func (o *ApAdvancedQosOpenApiVO) GetSupport2gOk() (*bool, bool)`

GetSupport2gOk returns a tuple with the Support2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport2g

`func (o *ApAdvancedQosOpenApiVO) SetSupport2g(v bool)`

SetSupport2g sets Support2g field to given value.

### HasSupport2g

`func (o *ApAdvancedQosOpenApiVO) HasSupport2g() bool`

HasSupport2g returns a boolean if a field has been set.

### GetSupport5g

`func (o *ApAdvancedQosOpenApiVO) GetSupport5g() bool`

GetSupport5g returns the Support5g field if non-nil, zero value otherwise.

### GetSupport5gOk

`func (o *ApAdvancedQosOpenApiVO) GetSupport5gOk() (*bool, bool)`

GetSupport5gOk returns a tuple with the Support5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g

`func (o *ApAdvancedQosOpenApiVO) SetSupport5g(v bool)`

SetSupport5g sets Support5g field to given value.

### HasSupport5g

`func (o *ApAdvancedQosOpenApiVO) HasSupport5g() bool`

HasSupport5g returns a boolean if a field has been set.

### GetSupport5g2

`func (o *ApAdvancedQosOpenApiVO) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *ApAdvancedQosOpenApiVO) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *ApAdvancedQosOpenApiVO) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *ApAdvancedQosOpenApiVO) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetSupport6g

`func (o *ApAdvancedQosOpenApiVO) GetSupport6g() bool`

GetSupport6g returns the Support6g field if non-nil, zero value otherwise.

### GetSupport6gOk

`func (o *ApAdvancedQosOpenApiVO) GetSupport6gOk() (*bool, bool)`

GetSupport6gOk returns a tuple with the Support6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport6g

`func (o *ApAdvancedQosOpenApiVO) SetSupport6g(v bool)`

SetSupport6g sets Support6g field to given value.

### HasSupport6g

`func (o *ApAdvancedQosOpenApiVO) HasSupport6g() bool`

HasSupport6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


