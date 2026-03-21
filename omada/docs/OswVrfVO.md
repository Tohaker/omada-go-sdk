# OswVrfVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | VRF ID | [optional] 
**Ipv4Enable** | **bool** | Whether to enable ipv4. Ipv4Enable should be true | 
**Ipv6Enable** | **bool** | Whether to enable ipv6 | 
**Vrf** | **string** | VRF should be 1 to 15 characters consisting of numbers (0 to 9), uppercase and lowercase letters (A to Z, a to z), and symbols -_@.+ but cannot be . or .. only. | 

## Methods

### NewOswVrfVO

`func NewOswVrfVO(ipv4Enable bool, ipv6Enable bool, vrf string, ) *OswVrfVO`

NewOswVrfVO instantiates a new OswVrfVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrfVOWithDefaults

`func NewOswVrfVOWithDefaults() *OswVrfVO`

NewOswVrfVOWithDefaults instantiates a new OswVrfVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OswVrfVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswVrfVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswVrfVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswVrfVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4Enable

`func (o *OswVrfVO) GetIpv4Enable() bool`

GetIpv4Enable returns the Ipv4Enable field if non-nil, zero value otherwise.

### GetIpv4EnableOk

`func (o *OswVrfVO) GetIpv4EnableOk() (*bool, bool)`

GetIpv4EnableOk returns a tuple with the Ipv4Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enable

`func (o *OswVrfVO) SetIpv4Enable(v bool)`

SetIpv4Enable sets Ipv4Enable field to given value.


### GetIpv6Enable

`func (o *OswVrfVO) GetIpv6Enable() bool`

GetIpv6Enable returns the Ipv6Enable field if non-nil, zero value otherwise.

### GetIpv6EnableOk

`func (o *OswVrfVO) GetIpv6EnableOk() (*bool, bool)`

GetIpv6EnableOk returns a tuple with the Ipv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enable

`func (o *OswVrfVO) SetIpv6Enable(v bool)`

SetIpv6Enable sets Ipv6Enable field to given value.


### GetVrf

`func (o *OswVrfVO) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *OswVrfVO) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *OswVrfVO) SetVrf(v string)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


