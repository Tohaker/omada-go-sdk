# OswSnmpVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | Contact | [optional] 
**Location** | Pointer to **string** | Location | [optional] 

## Methods

### NewOswSnmpVO

`func NewOswSnmpVO() *OswSnmpVO`

NewOswSnmpVO instantiates a new OswSnmpVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswSnmpVOWithDefaults

`func NewOswSnmpVOWithDefaults() *OswSnmpVO`

NewOswSnmpVOWithDefaults instantiates a new OswSnmpVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OswSnmpVO) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OswSnmpVO) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OswSnmpVO) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OswSnmpVO) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLocation

`func (o *OswSnmpVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswSnmpVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswSnmpVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswSnmpVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


