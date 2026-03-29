# EthUnit1PortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**DuplexLink** | Pointer to **string** |  | [optional] 
**FlowControl** | Pointer to **string** |  | [optional] 
**Lag** | Pointer to **string** |  | [optional] 
**LinkStatus** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Port** | **string** |  | 
**Speed** | Pointer to **int32** |  | [optional] 
**SpeedLink** | Pointer to **int32** |  | [optional] 
**SpeedMax** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEthUnit1PortDTO

`func NewEthUnit1PortDTO(port string, ) *EthUnit1PortDTO`

NewEthUnit1PortDTO instantiates a new EthUnit1PortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthUnit1PortDTOWithDefaults

`func NewEthUnit1PortDTOWithDefaults() *EthUnit1PortDTO`

NewEthUnit1PortDTOWithDefaults instantiates a new EthUnit1PortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EthUnit1PortDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EthUnit1PortDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EthUnit1PortDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EthUnit1PortDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *EthUnit1PortDTO) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *EthUnit1PortDTO) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *EthUnit1PortDTO) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *EthUnit1PortDTO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDuplexLink

`func (o *EthUnit1PortDTO) GetDuplexLink() string`

GetDuplexLink returns the DuplexLink field if non-nil, zero value otherwise.

### GetDuplexLinkOk

`func (o *EthUnit1PortDTO) GetDuplexLinkOk() (*string, bool)`

GetDuplexLinkOk returns a tuple with the DuplexLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexLink

`func (o *EthUnit1PortDTO) SetDuplexLink(v string)`

SetDuplexLink sets DuplexLink field to given value.

### HasDuplexLink

`func (o *EthUnit1PortDTO) HasDuplexLink() bool`

HasDuplexLink returns a boolean if a field has been set.

### GetFlowControl

`func (o *EthUnit1PortDTO) GetFlowControl() string`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *EthUnit1PortDTO) GetFlowControlOk() (*string, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *EthUnit1PortDTO) SetFlowControl(v string)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *EthUnit1PortDTO) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetLag

`func (o *EthUnit1PortDTO) GetLag() string`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *EthUnit1PortDTO) GetLagOk() (*string, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *EthUnit1PortDTO) SetLag(v string)`

SetLag sets Lag field to given value.

### HasLag

`func (o *EthUnit1PortDTO) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetLinkStatus

`func (o *EthUnit1PortDTO) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *EthUnit1PortDTO) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *EthUnit1PortDTO) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *EthUnit1PortDTO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetMediaType

`func (o *EthUnit1PortDTO) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *EthUnit1PortDTO) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *EthUnit1PortDTO) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *EthUnit1PortDTO) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetPort

`func (o *EthUnit1PortDTO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EthUnit1PortDTO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EthUnit1PortDTO) SetPort(v string)`

SetPort sets Port field to given value.


### GetSpeed

`func (o *EthUnit1PortDTO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EthUnit1PortDTO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EthUnit1PortDTO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EthUnit1PortDTO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSpeedLink

`func (o *EthUnit1PortDTO) GetSpeedLink() int32`

GetSpeedLink returns the SpeedLink field if non-nil, zero value otherwise.

### GetSpeedLinkOk

`func (o *EthUnit1PortDTO) GetSpeedLinkOk() (*int32, bool)`

GetSpeedLinkOk returns a tuple with the SpeedLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedLink

`func (o *EthUnit1PortDTO) SetSpeedLink(v int32)`

SetSpeedLink sets SpeedLink field to given value.

### HasSpeedLink

`func (o *EthUnit1PortDTO) HasSpeedLink() bool`

HasSpeedLink returns a boolean if a field has been set.

### GetSpeedMax

`func (o *EthUnit1PortDTO) GetSpeedMax() int32`

GetSpeedMax returns the SpeedMax field if non-nil, zero value otherwise.

### GetSpeedMaxOk

`func (o *EthUnit1PortDTO) GetSpeedMaxOk() (*int32, bool)`

GetSpeedMaxOk returns a tuple with the SpeedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMax

`func (o *EthUnit1PortDTO) SetSpeedMax(v int32)`

SetSpeedMax sets SpeedMax field to given value.

### HasSpeedMax

`func (o *EthUnit1PortDTO) HasSpeedMax() bool`

HasSpeedMax returns a boolean if a field has been set.

### GetStatus

`func (o *EthUnit1PortDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EthUnit1PortDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EthUnit1PortDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EthUnit1PortDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EthUnit1PortDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EthUnit1PortDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EthUnit1PortDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EthUnit1PortDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


