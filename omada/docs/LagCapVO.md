# LagCapVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LacpModSupport** | Pointer to **bool** | Lacp Mod Support | [optional] 
**LagHashAlgSupport** | Pointer to **bool** | Lag Hash Algorithm Support | [optional] 
**LagHashAlgs** | Pointer to **[]int32** | Support Lag Hash Algorithms, 0: SRC_MAC; 1: DST_MAC; 2: SRC_MAC_DST_MAC; 3: SRC_IP; 4: DST_IP; 5: SRC_IP_DST_IP | [optional] 

## Methods

### NewLagCapVO

`func NewLagCapVO() *LagCapVO`

NewLagCapVO instantiates a new LagCapVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagCapVOWithDefaults

`func NewLagCapVOWithDefaults() *LagCapVO`

NewLagCapVOWithDefaults instantiates a new LagCapVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLacpModSupport

`func (o *LagCapVO) GetLacpModSupport() bool`

GetLacpModSupport returns the LacpModSupport field if non-nil, zero value otherwise.

### GetLacpModSupportOk

`func (o *LagCapVO) GetLacpModSupportOk() (*bool, bool)`

GetLacpModSupportOk returns a tuple with the LacpModSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpModSupport

`func (o *LagCapVO) SetLacpModSupport(v bool)`

SetLacpModSupport sets LacpModSupport field to given value.

### HasLacpModSupport

`func (o *LagCapVO) HasLacpModSupport() bool`

HasLacpModSupport returns a boolean if a field has been set.

### GetLagHashAlgSupport

`func (o *LagCapVO) GetLagHashAlgSupport() bool`

GetLagHashAlgSupport returns the LagHashAlgSupport field if non-nil, zero value otherwise.

### GetLagHashAlgSupportOk

`func (o *LagCapVO) GetLagHashAlgSupportOk() (*bool, bool)`

GetLagHashAlgSupportOk returns a tuple with the LagHashAlgSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagHashAlgSupport

`func (o *LagCapVO) SetLagHashAlgSupport(v bool)`

SetLagHashAlgSupport sets LagHashAlgSupport field to given value.

### HasLagHashAlgSupport

`func (o *LagCapVO) HasLagHashAlgSupport() bool`

HasLagHashAlgSupport returns a boolean if a field has been set.

### GetLagHashAlgs

`func (o *LagCapVO) GetLagHashAlgs() []int32`

GetLagHashAlgs returns the LagHashAlgs field if non-nil, zero value otherwise.

### GetLagHashAlgsOk

`func (o *LagCapVO) GetLagHashAlgsOk() (*[]int32, bool)`

GetLagHashAlgsOk returns a tuple with the LagHashAlgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagHashAlgs

`func (o *LagCapVO) SetLagHashAlgs(v []int32)`

SetLagHashAlgs sets LagHashAlgs field to given value.

### HasLagHashAlgs

`func (o *LagCapVO) HasLagHashAlgs() bool`

HasLagHashAlgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


