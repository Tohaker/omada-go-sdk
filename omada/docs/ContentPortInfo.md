# ContentPortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Port Uuid. | 
**Ip** | **bool** | Whether to query the port IP. | 
**Name** | Pointer to **string** | Port name. | [optional] 
**Usage** | **bool** | Whether to query the port data usage. | 

## Methods

### NewContentPortInfo

`func NewContentPortInfo(id string, ip bool, usage bool, ) *ContentPortInfo`

NewContentPortInfo instantiates a new ContentPortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPortInfoWithDefaults

`func NewContentPortInfoWithDefaults() *ContentPortInfo`

NewContentPortInfoWithDefaults instantiates a new ContentPortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentPortInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentPortInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentPortInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIp

`func (o *ContentPortInfo) GetIp() bool`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ContentPortInfo) GetIpOk() (*bool, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ContentPortInfo) SetIp(v bool)`

SetIp sets Ip field to given value.


### GetName

`func (o *ContentPortInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentPortInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentPortInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentPortInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsage

`func (o *ContentPortInfo) GetUsage() bool`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ContentPortInfo) GetUsageOk() (*bool, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ContentPortInfo) SetUsage(v bool)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


