# IPsecFailoverConfiguration

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

### NewIPsecFailoverConfiguration

`func NewIPsecFailoverConfiguration(candidates []string, name string, primary string, ) *IPsecFailoverConfiguration`

NewIPsecFailoverConfiguration instantiates a new IPsecFailoverConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPsecFailoverConfigurationWithDefaults

`func NewIPsecFailoverConfigurationWithDefaults() *IPsecFailoverConfiguration`

NewIPsecFailoverConfigurationWithDefaults instantiates a new IPsecFailoverConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *IPsecFailoverConfiguration) GetCandidates() []string`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *IPsecFailoverConfiguration) GetCandidatesOk() (*[]string, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *IPsecFailoverConfiguration) SetCandidates(v []string)`

SetCandidates sets Candidates field to given value.


### GetFailback

`func (o *IPsecFailoverConfiguration) GetFailback() bool`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *IPsecFailoverConfiguration) GetFailbackOk() (*bool, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *IPsecFailoverConfiguration) SetFailback(v bool)`

SetFailback sets Failback field to given value.

### HasFailback

`func (o *IPsecFailoverConfiguration) HasFailback() bool`

HasFailback returns a boolean if a field has been set.

### GetFailbackTime

`func (o *IPsecFailoverConfiguration) GetFailbackTime() int32`

GetFailbackTime returns the FailbackTime field if non-nil, zero value otherwise.

### GetFailbackTimeOk

`func (o *IPsecFailoverConfiguration) GetFailbackTimeOk() (*int32, bool)`

GetFailbackTimeOk returns a tuple with the FailbackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailbackTime

`func (o *IPsecFailoverConfiguration) SetFailbackTime(v int32)`

SetFailbackTime sets FailbackTime field to given value.

### HasFailbackTime

`func (o *IPsecFailoverConfiguration) HasFailbackTime() bool`

HasFailbackTime returns a boolean if a field has been set.

### GetId

`func (o *IPsecFailoverConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPsecFailoverConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPsecFailoverConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPsecFailoverConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IPsecFailoverConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPsecFailoverConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPsecFailoverConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetPrimary

`func (o *IPsecFailoverConfiguration) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *IPsecFailoverConfiguration) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *IPsecFailoverConfiguration) SetPrimary(v string)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


