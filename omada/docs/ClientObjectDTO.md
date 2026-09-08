# ClientObjectDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | Client type: iPhone, iPod, Android, PC, printer, TV... | [optional] 
**Ip** | Pointer to **string** | Client IP address. | [optional] 
**Manager** | Pointer to **bool** | Whether it is the client currently being managed. | [optional] 
**Model** | Pointer to **string** | Client model. | [optional] 
**Name** | Pointer to **string** | Client Name, alias. | [optional] 
**Port** | Pointer to **int32** | (Wired) Port ID. | [optional] 
**RadioId** | Pointer to **int32** | (Wireless) Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz-1; 2:5GHz-2; 3: 6GHz. | [optional] 
**Ssid** | Pointer to **string** | (Wireless)  SSID name. | [optional] 

## Methods

### NewClientObjectDTO

`func NewClientObjectDTO() *ClientObjectDTO`

NewClientObjectDTO instantiates a new ClientObjectDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientObjectDTOWithDefaults

`func NewClientObjectDTOWithDefaults() *ClientObjectDTO`

NewClientObjectDTOWithDefaults instantiates a new ClientObjectDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *ClientObjectDTO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClientObjectDTO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClientObjectDTO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClientObjectDTO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIp

`func (o *ClientObjectDTO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientObjectDTO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientObjectDTO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientObjectDTO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetManager

`func (o *ClientObjectDTO) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClientObjectDTO) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClientObjectDTO) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ClientObjectDTO) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetModel

`func (o *ClientObjectDTO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientObjectDTO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientObjectDTO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClientObjectDTO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ClientObjectDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientObjectDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientObjectDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientObjectDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ClientObjectDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClientObjectDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClientObjectDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ClientObjectDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRadioId

`func (o *ClientObjectDTO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ClientObjectDTO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ClientObjectDTO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ClientObjectDTO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetSsid

`func (o *ClientObjectDTO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientObjectDTO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientObjectDTO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientObjectDTO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


