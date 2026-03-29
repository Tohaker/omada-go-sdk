# OsgSnmpVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 

## Methods

### NewOsgSnmpVO

`func NewOsgSnmpVO() *OsgSnmpVO`

NewOsgSnmpVO instantiates a new OsgSnmpVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgSnmpVOWithDefaults

`func NewOsgSnmpVOWithDefaults() *OsgSnmpVO`

NewOsgSnmpVOWithDefaults instantiates a new OsgSnmpVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OsgSnmpVO) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OsgSnmpVO) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OsgSnmpVO) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OsgSnmpVO) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLocation

`func (o *OsgSnmpVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OsgSnmpVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OsgSnmpVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OsgSnmpVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


