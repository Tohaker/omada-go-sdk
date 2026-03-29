# AuthenticationConfigAddDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationConfigInfo** | Pointer to [**AuthenticationConfigInfoDTO**](AuthenticationConfigInfoDTO.md) |  | [optional] 
**OnuId** | Pointer to **int32** | ONU ID,This parameter is optional. If not specified, the system will automatically allocate the smallest available ONU number under the current port.OnuId should be within the range of 0 to 127 | [optional] 
**PortId** | **string** | Pon port ID.e.g.GPON 1/0/1 | 

## Methods

### NewAuthenticationConfigAddDTO

`func NewAuthenticationConfigAddDTO(portId string, ) *AuthenticationConfigAddDTO`

NewAuthenticationConfigAddDTO instantiates a new AuthenticationConfigAddDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationConfigAddDTOWithDefaults

`func NewAuthenticationConfigAddDTOWithDefaults() *AuthenticationConfigAddDTO`

NewAuthenticationConfigAddDTOWithDefaults instantiates a new AuthenticationConfigAddDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationConfigInfo

`func (o *AuthenticationConfigAddDTO) GetAuthenticationConfigInfo() AuthenticationConfigInfoDTO`

GetAuthenticationConfigInfo returns the AuthenticationConfigInfo field if non-nil, zero value otherwise.

### GetAuthenticationConfigInfoOk

`func (o *AuthenticationConfigAddDTO) GetAuthenticationConfigInfoOk() (*AuthenticationConfigInfoDTO, bool)`

GetAuthenticationConfigInfoOk returns a tuple with the AuthenticationConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationConfigInfo

`func (o *AuthenticationConfigAddDTO) SetAuthenticationConfigInfo(v AuthenticationConfigInfoDTO)`

SetAuthenticationConfigInfo sets AuthenticationConfigInfo field to given value.

### HasAuthenticationConfigInfo

`func (o *AuthenticationConfigAddDTO) HasAuthenticationConfigInfo() bool`

HasAuthenticationConfigInfo returns a boolean if a field has been set.

### GetOnuId

`func (o *AuthenticationConfigAddDTO) GetOnuId() int32`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *AuthenticationConfigAddDTO) GetOnuIdOk() (*int32, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *AuthenticationConfigAddDTO) SetOnuId(v int32)`

SetOnuId sets OnuId field to given value.

### HasOnuId

`func (o *AuthenticationConfigAddDTO) HasOnuId() bool`

HasOnuId returns a boolean if a field has been set.

### GetPortId

`func (o *AuthenticationConfigAddDTO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *AuthenticationConfigAddDTO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *AuthenticationConfigAddDTO) SetPortId(v string)`

SetPortId sets PortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


