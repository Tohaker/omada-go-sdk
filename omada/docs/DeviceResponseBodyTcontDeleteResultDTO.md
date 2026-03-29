# DeviceResponseBodyTcontDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TcontDeleteResultDTO**](TcontDeleteResultDTO.md) |  | [optional] 
**DeviceType** | Pointer to **string** | deviceType should be a value as follows:ap,gateway,switch,pro ap,pro gateway,pro switch,festa ap,festa gateway,festa switch,ems,olt,onu. | [optional] 
**Errcode** | Pointer to **int32** | Device error code. | [optional] 
**Message** | Pointer to **string** | Device error message | [optional] 

## Methods

### NewDeviceResponseBodyTcontDeleteResultDTO

`func NewDeviceResponseBodyTcontDeleteResultDTO() *DeviceResponseBodyTcontDeleteResultDTO`

NewDeviceResponseBodyTcontDeleteResultDTO instantiates a new DeviceResponseBodyTcontDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceResponseBodyTcontDeleteResultDTOWithDefaults

`func NewDeviceResponseBodyTcontDeleteResultDTOWithDefaults() *DeviceResponseBodyTcontDeleteResultDTO`

NewDeviceResponseBodyTcontDeleteResultDTOWithDefaults instantiates a new DeviceResponseBodyTcontDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetData() TcontDeleteResultDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetDataOk() (*TcontDeleteResultDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceResponseBodyTcontDeleteResultDTO) SetData(v TcontDeleteResultDTO)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceResponseBodyTcontDeleteResultDTO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceResponseBodyTcontDeleteResultDTO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceResponseBodyTcontDeleteResultDTO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetErrcode

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *DeviceResponseBodyTcontDeleteResultDTO) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *DeviceResponseBodyTcontDeleteResultDTO) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceResponseBodyTcontDeleteResultDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceResponseBodyTcontDeleteResultDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceResponseBodyTcontDeleteResultDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


