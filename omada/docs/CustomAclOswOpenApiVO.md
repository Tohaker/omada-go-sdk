# CustomAclOswOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | MAC | 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**Vrf** | **string** | VRF should be 1 to 15 characters consisting of numbers (0 to 9), uppercase and lowercase letters (A to Z, a to z), and symbols -_@.+ but cannot be . or .. only. | 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewCustomAclOswOpenApiVO

`func NewCustomAclOswOpenApiVO(mac string, vrf string, ) *CustomAclOswOpenApiVO`

NewCustomAclOswOpenApiVO instantiates a new CustomAclOswOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAclOswOpenApiVOWithDefaults

`func NewCustomAclOswOpenApiVOWithDefaults() *CustomAclOswOpenApiVO`

NewCustomAclOswOpenApiVOWithDefaults instantiates a new CustomAclOswOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *CustomAclOswOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CustomAclOswOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CustomAclOswOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetStackId

`func (o *CustomAclOswOpenApiVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CustomAclOswOpenApiVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CustomAclOswOpenApiVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CustomAclOswOpenApiVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetVrf

`func (o *CustomAclOswOpenApiVO) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *CustomAclOswOpenApiVO) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *CustomAclOswOpenApiVO) SetVrf(v string)`

SetVrf sets Vrf field to given value.


### GetVrfId

`func (o *CustomAclOswOpenApiVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *CustomAclOswOpenApiVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *CustomAclOswOpenApiVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *CustomAclOswOpenApiVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


