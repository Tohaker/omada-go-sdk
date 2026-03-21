# ClientFilteringOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **[]string** | Device list. | [optional] 
**DeviceType** | Pointer to [**[]ClientCategoryOptionsOpenApiVO**](ClientCategoryOptionsOpenApiVO.md) | Type list of clients. | [optional] 
**Network** | Pointer to **[]string** | Network list. | [optional] 
**Ssid** | Pointer to **[]string** | Ssid list. | [optional] 
**Vendor** | Pointer to **[]string** | Vendor list of clients. | [optional] 

## Methods

### NewClientFilteringOptions

`func NewClientFilteringOptions() *ClientFilteringOptions`

NewClientFilteringOptions instantiates a new ClientFilteringOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientFilteringOptionsWithDefaults

`func NewClientFilteringOptionsWithDefaults() *ClientFilteringOptions`

NewClientFilteringOptionsWithDefaults instantiates a new ClientFilteringOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *ClientFilteringOptions) GetDevice() []string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ClientFilteringOptions) GetDeviceOk() (*[]string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ClientFilteringOptions) SetDevice(v []string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ClientFilteringOptions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceType

`func (o *ClientFilteringOptions) GetDeviceType() []ClientCategoryOptionsOpenApiVO`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClientFilteringOptions) GetDeviceTypeOk() (*[]ClientCategoryOptionsOpenApiVO, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClientFilteringOptions) SetDeviceType(v []ClientCategoryOptionsOpenApiVO)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClientFilteringOptions) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetNetwork

`func (o *ClientFilteringOptions) GetNetwork() []string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClientFilteringOptions) GetNetworkOk() (*[]string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClientFilteringOptions) SetNetwork(v []string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClientFilteringOptions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSsid

`func (o *ClientFilteringOptions) GetSsid() []string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientFilteringOptions) GetSsidOk() (*[]string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientFilteringOptions) SetSsid(v []string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientFilteringOptions) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetVendor

`func (o *ClientFilteringOptions) GetVendor() []string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ClientFilteringOptions) GetVendorOk() (*[]string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ClientFilteringOptions) SetVendor(v []string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ClientFilteringOptions) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


