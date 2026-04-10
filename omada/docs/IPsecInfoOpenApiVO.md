# IPsecInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the IPSec info. | [optional] 
**Name** | Pointer to **string** | Name of the IPSec info. | [optional] 
**Status** | Pointer to **int32** | Status of the IPSec info. 0: Active; 1: Standby; 2: Disconnected. | [optional] 

## Methods

### NewIPsecInfoOpenApiVO

`func NewIPsecInfoOpenApiVO() *IPsecInfoOpenApiVO`

NewIPsecInfoOpenApiVO instantiates a new IPsecInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPsecInfoOpenApiVOWithDefaults

`func NewIPsecInfoOpenApiVOWithDefaults() *IPsecInfoOpenApiVO`

NewIPsecInfoOpenApiVOWithDefaults instantiates a new IPsecInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPsecInfoOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPsecInfoOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPsecInfoOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPsecInfoOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IPsecInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPsecInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPsecInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPsecInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IPsecInfoOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPsecInfoOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPsecInfoOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IPsecInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


