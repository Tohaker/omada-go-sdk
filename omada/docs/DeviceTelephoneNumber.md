# DeviceTelephoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindNumberList** | Pointer to [**[]TelephoneNumberWithStatusOpenApiVO**](TelephoneNumberWithStatusOpenApiVO.md) | The list of telephone number. | [optional] 
**Id** | Pointer to **string** | The ID of voip device. | [optional] 
**Mac** | Pointer to **string** | The mac of voip device. | [optional] 

## Methods

### NewDeviceTelephoneNumber

`func NewDeviceTelephoneNumber() *DeviceTelephoneNumber`

NewDeviceTelephoneNumber instantiates a new DeviceTelephoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTelephoneNumberWithDefaults

`func NewDeviceTelephoneNumberWithDefaults() *DeviceTelephoneNumber`

NewDeviceTelephoneNumberWithDefaults instantiates a new DeviceTelephoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindNumberList

`func (o *DeviceTelephoneNumber) GetBindNumberList() []TelephoneNumberWithStatusOpenApiVO`

GetBindNumberList returns the BindNumberList field if non-nil, zero value otherwise.

### GetBindNumberListOk

`func (o *DeviceTelephoneNumber) GetBindNumberListOk() (*[]TelephoneNumberWithStatusOpenApiVO, bool)`

GetBindNumberListOk returns a tuple with the BindNumberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindNumberList

`func (o *DeviceTelephoneNumber) SetBindNumberList(v []TelephoneNumberWithStatusOpenApiVO)`

SetBindNumberList sets BindNumberList field to given value.

### HasBindNumberList

`func (o *DeviceTelephoneNumber) HasBindNumberList() bool`

HasBindNumberList returns a boolean if a field has been set.

### GetId

`func (o *DeviceTelephoneNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceTelephoneNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceTelephoneNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceTelephoneNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *DeviceTelephoneNumber) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceTelephoneNumber) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceTelephoneNumber) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceTelephoneNumber) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


