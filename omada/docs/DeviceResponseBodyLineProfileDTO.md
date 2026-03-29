# DeviceResponseBodyLineProfileDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LineProfileDTO**](LineProfileDTO.md) |  | [optional] 
**DeviceType** | Pointer to **string** | deviceType should be a value as follows:ap,gateway,switch,pro ap,pro gateway,pro switch,festa ap,festa gateway,festa switch,ems,olt,onu. | [optional] 
**Errcode** | Pointer to **int32** | Device error code. | [optional] 
**Message** | Pointer to **string** | Device error message | [optional] 

## Methods

### NewDeviceResponseBodyLineProfileDTO

`func NewDeviceResponseBodyLineProfileDTO() *DeviceResponseBodyLineProfileDTO`

NewDeviceResponseBodyLineProfileDTO instantiates a new DeviceResponseBodyLineProfileDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceResponseBodyLineProfileDTOWithDefaults

`func NewDeviceResponseBodyLineProfileDTOWithDefaults() *DeviceResponseBodyLineProfileDTO`

NewDeviceResponseBodyLineProfileDTOWithDefaults instantiates a new DeviceResponseBodyLineProfileDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceResponseBodyLineProfileDTO) GetData() LineProfileDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceResponseBodyLineProfileDTO) GetDataOk() (*LineProfileDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceResponseBodyLineProfileDTO) SetData(v LineProfileDTO)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceResponseBodyLineProfileDTO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceResponseBodyLineProfileDTO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceResponseBodyLineProfileDTO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceResponseBodyLineProfileDTO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceResponseBodyLineProfileDTO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetErrcode

`func (o *DeviceResponseBodyLineProfileDTO) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *DeviceResponseBodyLineProfileDTO) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *DeviceResponseBodyLineProfileDTO) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *DeviceResponseBodyLineProfileDTO) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceResponseBodyLineProfileDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceResponseBodyLineProfileDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceResponseBodyLineProfileDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceResponseBodyLineProfileDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


