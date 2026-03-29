# ApSnmpVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | Contact, contact should contain 0 to 128 ASCII characters, spaces are not allowed. | [optional] 
**Location** | Pointer to **string** | Location, location should contain 0 to 128 ASCII characters, spaces are allowed, and leading and trailing spaces are not allowed. | [optional] 

## Methods

### NewApSnmpVO

`func NewApSnmpVO() *ApSnmpVO`

NewApSnmpVO instantiates a new ApSnmpVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApSnmpVOWithDefaults

`func NewApSnmpVOWithDefaults() *ApSnmpVO`

NewApSnmpVOWithDefaults instantiates a new ApSnmpVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ApSnmpVO) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ApSnmpVO) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ApSnmpVO) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ApSnmpVO) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLocation

`func (o *ApSnmpVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApSnmpVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApSnmpVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApSnmpVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


