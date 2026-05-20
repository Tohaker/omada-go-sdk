# EPDGOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain. If parameter [type] is 0, it should not be null. | [optional] 
**Ip** | Pointer to **string** | Ip. If parameter [type] is 1, it should not be null. | [optional] 
**QosPriority** | Pointer to **int32** | QOS Priority | [optional] 
**Type** | Pointer to **int32** | Type. Such as: 0: domain, 1:ip. | [optional] 

## Methods

### NewEPDGOpenApiVO

`func NewEPDGOpenApiVO() *EPDGOpenApiVO`

NewEPDGOpenApiVO instantiates a new EPDGOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEPDGOpenApiVOWithDefaults

`func NewEPDGOpenApiVOWithDefaults() *EPDGOpenApiVO`

NewEPDGOpenApiVOWithDefaults instantiates a new EPDGOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *EPDGOpenApiVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EPDGOpenApiVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EPDGOpenApiVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EPDGOpenApiVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIp

`func (o *EPDGOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EPDGOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EPDGOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EPDGOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetQosPriority

`func (o *EPDGOpenApiVO) GetQosPriority() int32`

GetQosPriority returns the QosPriority field if non-nil, zero value otherwise.

### GetQosPriorityOk

`func (o *EPDGOpenApiVO) GetQosPriorityOk() (*int32, bool)`

GetQosPriorityOk returns a tuple with the QosPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPriority

`func (o *EPDGOpenApiVO) SetQosPriority(v int32)`

SetQosPriority sets QosPriority field to given value.

### HasQosPriority

`func (o *EPDGOpenApiVO) HasQosPriority() bool`

HasQosPriority returns a boolean if a field has been set.

### GetType

`func (o *EPDGOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EPDGOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EPDGOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *EPDGOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


