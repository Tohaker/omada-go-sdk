# InternetBasicInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | Port ID | [optional] 
**PortMode** | Pointer to **int32** | Port mode should be a value as follows: 0: WAN; 1: LAN. | [optional] 
**PortName** | Pointer to **string** | Port name | [optional] 
**PortType** | Pointer to **int32** | Port type should be a value as follows: 0: WAN; 1: WAN/LAN. | [optional] 

## Methods

### NewInternetBasicInfo

`func NewInternetBasicInfo() *InternetBasicInfo`

NewInternetBasicInfo instantiates a new InternetBasicInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetBasicInfoWithDefaults

`func NewInternetBasicInfoWithDefaults() *InternetBasicInfo`

NewInternetBasicInfoWithDefaults instantiates a new InternetBasicInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InternetBasicInfo) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InternetBasicInfo) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InternetBasicInfo) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InternetBasicInfo) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortMode

`func (o *InternetBasicInfo) GetPortMode() int32`

GetPortMode returns the PortMode field if non-nil, zero value otherwise.

### GetPortModeOk

`func (o *InternetBasicInfo) GetPortModeOk() (*int32, bool)`

GetPortModeOk returns a tuple with the PortMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMode

`func (o *InternetBasicInfo) SetPortMode(v int32)`

SetPortMode sets PortMode field to given value.

### HasPortMode

`func (o *InternetBasicInfo) HasPortMode() bool`

HasPortMode returns a boolean if a field has been set.

### GetPortName

`func (o *InternetBasicInfo) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *InternetBasicInfo) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *InternetBasicInfo) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *InternetBasicInfo) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortType

`func (o *InternetBasicInfo) GetPortType() int32`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *InternetBasicInfo) GetPortTypeOk() (*int32, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *InternetBasicInfo) SetPortType(v int32)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *InternetBasicInfo) HasPortType() bool`

HasPortType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


