# OswVrfOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | VRF ID | [optional] 
**Ipv4Enable** | **bool** | Indicates whether ipv4 is enabled | 
**Ipv6Enable** | **bool** | Indicates whether ipv6 is enabled | 
**IsDefault** | Pointer to **bool** | Indicates whether vrf is default vrf | [optional] 
**Resource** | Pointer to **int32** | Resource | [optional] 
**Vrf** | **string** | VRF Name | 

## Methods

### NewOswVrfOpenApiVO

`func NewOswVrfOpenApiVO(ipv4Enable bool, ipv6Enable bool, vrf string, ) *OswVrfOpenApiVO`

NewOswVrfOpenApiVO instantiates a new OswVrfOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrfOpenApiVOWithDefaults

`func NewOswVrfOpenApiVOWithDefaults() *OswVrfOpenApiVO`

NewOswVrfOpenApiVOWithDefaults instantiates a new OswVrfOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OswVrfOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswVrfOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswVrfOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswVrfOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4Enable

`func (o *OswVrfOpenApiVO) GetIpv4Enable() bool`

GetIpv4Enable returns the Ipv4Enable field if non-nil, zero value otherwise.

### GetIpv4EnableOk

`func (o *OswVrfOpenApiVO) GetIpv4EnableOk() (*bool, bool)`

GetIpv4EnableOk returns a tuple with the Ipv4Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enable

`func (o *OswVrfOpenApiVO) SetIpv4Enable(v bool)`

SetIpv4Enable sets Ipv4Enable field to given value.


### GetIpv6Enable

`func (o *OswVrfOpenApiVO) GetIpv6Enable() bool`

GetIpv6Enable returns the Ipv6Enable field if non-nil, zero value otherwise.

### GetIpv6EnableOk

`func (o *OswVrfOpenApiVO) GetIpv6EnableOk() (*bool, bool)`

GetIpv6EnableOk returns a tuple with the Ipv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enable

`func (o *OswVrfOpenApiVO) SetIpv6Enable(v bool)`

SetIpv6Enable sets Ipv6Enable field to given value.


### GetIsDefault

`func (o *OswVrfOpenApiVO) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OswVrfOpenApiVO) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OswVrfOpenApiVO) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OswVrfOpenApiVO) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetResource

`func (o *OswVrfOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswVrfOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswVrfOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswVrfOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetVrf

`func (o *OswVrfOpenApiVO) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *OswVrfOpenApiVO) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *OswVrfOpenApiVO) SetVrf(v string)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


