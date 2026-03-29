# IPsecFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | **[]string** | Candidates of the IPSec failover. IPsec VPN can be created using &#39;Create site-to-site VPN&#39; interface, and ID can be obtained from &#39;Get site-to-site VPN list&#39; interface. | 
**Failback** | Pointer to **bool** | Failback of the IPSec failover. | [optional] 
**FailbackTime** | Pointer to **int32** | Failback time should be within the range of 10–3600s. | [optional] 
**Id** | Pointer to **string** | ID of the IPSec failover. | [optional] 
**Name** | **string** | Name should contain 1 to 64 characters. | 
**Primary** | **string** | Primary of the IPSec failover. IPsec VPN can be created using &#39;Create site-to-site VPN&#39; interface, and ID can be obtained from &#39;Get site-to-site VPN list&#39; interface. | 

## Methods

### NewIPsecFailover

`func NewIPsecFailover(candidates []string, name string, primary string, ) *IPsecFailover`

NewIPsecFailover instantiates a new IPsecFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPsecFailoverWithDefaults

`func NewIPsecFailoverWithDefaults() *IPsecFailover`

NewIPsecFailoverWithDefaults instantiates a new IPsecFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *IPsecFailover) GetCandidates() []string`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *IPsecFailover) GetCandidatesOk() (*[]string, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *IPsecFailover) SetCandidates(v []string)`

SetCandidates sets Candidates field to given value.


### GetFailback

`func (o *IPsecFailover) GetFailback() bool`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *IPsecFailover) GetFailbackOk() (*bool, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *IPsecFailover) SetFailback(v bool)`

SetFailback sets Failback field to given value.

### HasFailback

`func (o *IPsecFailover) HasFailback() bool`

HasFailback returns a boolean if a field has been set.

### GetFailbackTime

`func (o *IPsecFailover) GetFailbackTime() int32`

GetFailbackTime returns the FailbackTime field if non-nil, zero value otherwise.

### GetFailbackTimeOk

`func (o *IPsecFailover) GetFailbackTimeOk() (*int32, bool)`

GetFailbackTimeOk returns a tuple with the FailbackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailbackTime

`func (o *IPsecFailover) SetFailbackTime(v int32)`

SetFailbackTime sets FailbackTime field to given value.

### HasFailbackTime

`func (o *IPsecFailover) HasFailbackTime() bool`

HasFailbackTime returns a boolean if a field has been set.

### GetId

`func (o *IPsecFailover) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPsecFailover) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPsecFailover) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPsecFailover) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IPsecFailover) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPsecFailover) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPsecFailover) SetName(v string)`

SetName sets Name field to given value.


### GetPrimary

`func (o *IPsecFailover) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *IPsecFailover) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *IPsecFailover) SetPrimary(v string)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


