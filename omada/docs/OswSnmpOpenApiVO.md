# OswSnmpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | Contact, contact should contain 0 to 128 ASCII characters, spaces are not allowed. | [optional] 
**Location** | Pointer to **string** | Location, location should contain 0 to 128 ASCII characters, spaces are allowed, and leading and trailing spaces are not allowed. | [optional] 
**Type** | Pointer to **int32** | SNMP config type, 0: Use Site Settings; 1: Custom. | [optional] 

## Methods

### NewOswSnmpOpenApiVO

`func NewOswSnmpOpenApiVO() *OswSnmpOpenApiVO`

NewOswSnmpOpenApiVO instantiates a new OswSnmpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswSnmpOpenApiVOWithDefaults

`func NewOswSnmpOpenApiVOWithDefaults() *OswSnmpOpenApiVO`

NewOswSnmpOpenApiVOWithDefaults instantiates a new OswSnmpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OswSnmpOpenApiVO) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OswSnmpOpenApiVO) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OswSnmpOpenApiVO) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OswSnmpOpenApiVO) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLocation

`func (o *OswSnmpOpenApiVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswSnmpOpenApiVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswSnmpOpenApiVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswSnmpOpenApiVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *OswSnmpOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswSnmpOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswSnmpOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswSnmpOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


