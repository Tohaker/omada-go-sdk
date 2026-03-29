# DeviceResponseBodyTrafficProfileDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TrafficProfileDeleteResultDTO**](TrafficProfileDeleteResultDTO.md) |  | [optional] 
**DeviceType** | Pointer to **string** | deviceType should be a value as follows:ap,gateway,switch,pro ap,pro gateway,pro switch,festa ap,festa gateway,festa switch,ems,olt,onu. | [optional] 
**Errcode** | Pointer to **int32** | Device error code. | [optional] 
**Message** | Pointer to **string** | Device error message | [optional] 

## Methods

### NewDeviceResponseBodyTrafficProfileDeleteResultDTO

`func NewDeviceResponseBodyTrafficProfileDeleteResultDTO() *DeviceResponseBodyTrafficProfileDeleteResultDTO`

NewDeviceResponseBodyTrafficProfileDeleteResultDTO instantiates a new DeviceResponseBodyTrafficProfileDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceResponseBodyTrafficProfileDeleteResultDTOWithDefaults

`func NewDeviceResponseBodyTrafficProfileDeleteResultDTOWithDefaults() *DeviceResponseBodyTrafficProfileDeleteResultDTO`

NewDeviceResponseBodyTrafficProfileDeleteResultDTOWithDefaults instantiates a new DeviceResponseBodyTrafficProfileDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetData() TrafficProfileDeleteResultDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetDataOk() (*TrafficProfileDeleteResultDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) SetData(v TrafficProfileDeleteResultDTO)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetErrcode

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceResponseBodyTrafficProfileDeleteResultDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


