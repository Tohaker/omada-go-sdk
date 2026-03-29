# TelephoneNumberListUpdateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTelephoneNumbers** | Pointer to [**[]TelephoneNumberWithoutStatusOpenApiVO**](TelephoneNumberWithoutStatusOpenApiVO.md) | Telephone numbers to be added. | [optional] 
**DeleteTelephoneNumberIds** | Pointer to **[]string** | Telephone number IDs to be deleted. | [optional] 
**ModifyTelephoneNumbers** | Pointer to [**[]TelephoneNumberWithoutStatusOpenApiVO**](TelephoneNumberWithoutStatusOpenApiVO.md) | Telephone numbers to be modified. | [optional] 

## Methods

### NewTelephoneNumberListUpdateOpenApiVO

`func NewTelephoneNumberListUpdateOpenApiVO() *TelephoneNumberListUpdateOpenApiVO`

NewTelephoneNumberListUpdateOpenApiVO instantiates a new TelephoneNumberListUpdateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephoneNumberListUpdateOpenApiVOWithDefaults

`func NewTelephoneNumberListUpdateOpenApiVOWithDefaults() *TelephoneNumberListUpdateOpenApiVO`

NewTelephoneNumberListUpdateOpenApiVOWithDefaults instantiates a new TelephoneNumberListUpdateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTelephoneNumbers

`func (o *TelephoneNumberListUpdateOpenApiVO) GetAddTelephoneNumbers() []TelephoneNumberWithoutStatusOpenApiVO`

GetAddTelephoneNumbers returns the AddTelephoneNumbers field if non-nil, zero value otherwise.

### GetAddTelephoneNumbersOk

`func (o *TelephoneNumberListUpdateOpenApiVO) GetAddTelephoneNumbersOk() (*[]TelephoneNumberWithoutStatusOpenApiVO, bool)`

GetAddTelephoneNumbersOk returns a tuple with the AddTelephoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTelephoneNumbers

`func (o *TelephoneNumberListUpdateOpenApiVO) SetAddTelephoneNumbers(v []TelephoneNumberWithoutStatusOpenApiVO)`

SetAddTelephoneNumbers sets AddTelephoneNumbers field to given value.

### HasAddTelephoneNumbers

`func (o *TelephoneNumberListUpdateOpenApiVO) HasAddTelephoneNumbers() bool`

HasAddTelephoneNumbers returns a boolean if a field has been set.

### GetDeleteTelephoneNumberIds

`func (o *TelephoneNumberListUpdateOpenApiVO) GetDeleteTelephoneNumberIds() []string`

GetDeleteTelephoneNumberIds returns the DeleteTelephoneNumberIds field if non-nil, zero value otherwise.

### GetDeleteTelephoneNumberIdsOk

`func (o *TelephoneNumberListUpdateOpenApiVO) GetDeleteTelephoneNumberIdsOk() (*[]string, bool)`

GetDeleteTelephoneNumberIdsOk returns a tuple with the DeleteTelephoneNumberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteTelephoneNumberIds

`func (o *TelephoneNumberListUpdateOpenApiVO) SetDeleteTelephoneNumberIds(v []string)`

SetDeleteTelephoneNumberIds sets DeleteTelephoneNumberIds field to given value.

### HasDeleteTelephoneNumberIds

`func (o *TelephoneNumberListUpdateOpenApiVO) HasDeleteTelephoneNumberIds() bool`

HasDeleteTelephoneNumberIds returns a boolean if a field has been set.

### GetModifyTelephoneNumbers

`func (o *TelephoneNumberListUpdateOpenApiVO) GetModifyTelephoneNumbers() []TelephoneNumberWithoutStatusOpenApiVO`

GetModifyTelephoneNumbers returns the ModifyTelephoneNumbers field if non-nil, zero value otherwise.

### GetModifyTelephoneNumbersOk

`func (o *TelephoneNumberListUpdateOpenApiVO) GetModifyTelephoneNumbersOk() (*[]TelephoneNumberWithoutStatusOpenApiVO, bool)`

GetModifyTelephoneNumbersOk returns a tuple with the ModifyTelephoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTelephoneNumbers

`func (o *TelephoneNumberListUpdateOpenApiVO) SetModifyTelephoneNumbers(v []TelephoneNumberWithoutStatusOpenApiVO)`

SetModifyTelephoneNumbers sets ModifyTelephoneNumbers field to given value.

### HasModifyTelephoneNumbers

`func (o *TelephoneNumberListUpdateOpenApiVO) HasModifyTelephoneNumbers() bool`

HasModifyTelephoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


