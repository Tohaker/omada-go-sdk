# OswVrfConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Enable** | Pointer to **bool** | Whether to enable ipv4. Ipv4Enable should be true. This parameter is required when creating a VRF and is not needed when modifying a VRF. | [optional] 
**Ipv6Enable** | **bool** | Whether to enable ipv6 | 
**Vrf** | Pointer to **string** | VRF name: 1 to 15 characters consisting of numbers (0 to 9), uppercase and lowercase letters (A to Z, a to z), and symbols -_@.+; cannot be &#39;.&#39; or &#39;..&#39; only. This parameter is required when creating a VRF and is not needed when modifying a VRF. | [optional] 

## Methods

### NewOswVrfConfigOpenApiVO

`func NewOswVrfConfigOpenApiVO(ipv6Enable bool, ) *OswVrfConfigOpenApiVO`

NewOswVrfConfigOpenApiVO instantiates a new OswVrfConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrfConfigOpenApiVOWithDefaults

`func NewOswVrfConfigOpenApiVOWithDefaults() *OswVrfConfigOpenApiVO`

NewOswVrfConfigOpenApiVOWithDefaults instantiates a new OswVrfConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Enable

`func (o *OswVrfConfigOpenApiVO) GetIpv4Enable() bool`

GetIpv4Enable returns the Ipv4Enable field if non-nil, zero value otherwise.

### GetIpv4EnableOk

`func (o *OswVrfConfigOpenApiVO) GetIpv4EnableOk() (*bool, bool)`

GetIpv4EnableOk returns a tuple with the Ipv4Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enable

`func (o *OswVrfConfigOpenApiVO) SetIpv4Enable(v bool)`

SetIpv4Enable sets Ipv4Enable field to given value.

### HasIpv4Enable

`func (o *OswVrfConfigOpenApiVO) HasIpv4Enable() bool`

HasIpv4Enable returns a boolean if a field has been set.

### GetIpv6Enable

`func (o *OswVrfConfigOpenApiVO) GetIpv6Enable() bool`

GetIpv6Enable returns the Ipv6Enable field if non-nil, zero value otherwise.

### GetIpv6EnableOk

`func (o *OswVrfConfigOpenApiVO) GetIpv6EnableOk() (*bool, bool)`

GetIpv6EnableOk returns a tuple with the Ipv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enable

`func (o *OswVrfConfigOpenApiVO) SetIpv6Enable(v bool)`

SetIpv6Enable sets Ipv6Enable field to given value.


### GetVrf

`func (o *OswVrfConfigOpenApiVO) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *OswVrfConfigOpenApiVO) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *OswVrfConfigOpenApiVO) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *OswVrfConfigOpenApiVO) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


