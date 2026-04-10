# IPsecFailoverInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | Pointer to [**[]IPsecInfoOpenApiVO**](IPsecInfoOpenApiVO.md) | Candidates of the ipsec fail over. | [optional] 
**FailBack** | Pointer to **bool** | FailBack of the ipsec fail over. | [optional] 
**FailBackTime** | Pointer to **int32** | FailBackTime of the ipsec fail over. | [optional] 
**Id** | Pointer to **string** | ID of the ipsec fail over. | [optional] 
**Name** | Pointer to **string** | Name of the ipsec fail over. | [optional] 
**Primary** | Pointer to [**IPsecInfoOpenApiVO**](IPsecInfoOpenApiVO.md) |  | [optional] 

## Methods

### NewIPsecFailoverInformation

`func NewIPsecFailoverInformation() *IPsecFailoverInformation`

NewIPsecFailoverInformation instantiates a new IPsecFailoverInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPsecFailoverInformationWithDefaults

`func NewIPsecFailoverInformationWithDefaults() *IPsecFailoverInformation`

NewIPsecFailoverInformationWithDefaults instantiates a new IPsecFailoverInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *IPsecFailoverInformation) GetCandidates() []IPsecInfoOpenApiVO`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *IPsecFailoverInformation) GetCandidatesOk() (*[]IPsecInfoOpenApiVO, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *IPsecFailoverInformation) SetCandidates(v []IPsecInfoOpenApiVO)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *IPsecFailoverInformation) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetFailBack

`func (o *IPsecFailoverInformation) GetFailBack() bool`

GetFailBack returns the FailBack field if non-nil, zero value otherwise.

### GetFailBackOk

`func (o *IPsecFailoverInformation) GetFailBackOk() (*bool, bool)`

GetFailBackOk returns a tuple with the FailBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailBack

`func (o *IPsecFailoverInformation) SetFailBack(v bool)`

SetFailBack sets FailBack field to given value.

### HasFailBack

`func (o *IPsecFailoverInformation) HasFailBack() bool`

HasFailBack returns a boolean if a field has been set.

### GetFailBackTime

`func (o *IPsecFailoverInformation) GetFailBackTime() int32`

GetFailBackTime returns the FailBackTime field if non-nil, zero value otherwise.

### GetFailBackTimeOk

`func (o *IPsecFailoverInformation) GetFailBackTimeOk() (*int32, bool)`

GetFailBackTimeOk returns a tuple with the FailBackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailBackTime

`func (o *IPsecFailoverInformation) SetFailBackTime(v int32)`

SetFailBackTime sets FailBackTime field to given value.

### HasFailBackTime

`func (o *IPsecFailoverInformation) HasFailBackTime() bool`

HasFailBackTime returns a boolean if a field has been set.

### GetId

`func (o *IPsecFailoverInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPsecFailoverInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPsecFailoverInformation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPsecFailoverInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IPsecFailoverInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPsecFailoverInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPsecFailoverInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPsecFailoverInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *IPsecFailoverInformation) GetPrimary() IPsecInfoOpenApiVO`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *IPsecFailoverInformation) GetPrimaryOk() (*IPsecInfoOpenApiVO, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *IPsecFailoverInformation) SetPrimary(v IPsecInfoOpenApiVO)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *IPsecFailoverInformation) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


