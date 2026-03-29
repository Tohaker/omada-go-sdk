# IpsBlockListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | The destination IP in defined blocked or isolated policy. | [optional] 
**Id** | Pointer to **string** | Block list entry ID. | [optional] 
**Name** | Pointer to **string** | The name of the defined blocked or isolated policy. | [optional] 
**SourceIp** | Pointer to **string** | Source IP in defined blocked or isolated policy. | [optional] 
**SourceLocation** | Pointer to **string** | The region code of the source IP. | [optional] 

## Methods

### NewIpsBlockListEntry

`func NewIpsBlockListEntry() *IpsBlockListEntry`

NewIpsBlockListEntry instantiates a new IpsBlockListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsBlockListEntryWithDefaults

`func NewIpsBlockListEntryWithDefaults() *IpsBlockListEntry`

NewIpsBlockListEntryWithDefaults instantiates a new IpsBlockListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *IpsBlockListEntry) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *IpsBlockListEntry) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *IpsBlockListEntry) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *IpsBlockListEntry) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetId

`func (o *IpsBlockListEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsBlockListEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsBlockListEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpsBlockListEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpsBlockListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpsBlockListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpsBlockListEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpsBlockListEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceIp

`func (o *IpsBlockListEntry) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *IpsBlockListEntry) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *IpsBlockListEntry) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *IpsBlockListEntry) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetSourceLocation

`func (o *IpsBlockListEntry) GetSourceLocation() string`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *IpsBlockListEntry) GetSourceLocationOk() (*string, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *IpsBlockListEntry) SetSourceLocation(v string)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *IpsBlockListEntry) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


