# IptvDslOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpPhoneVci** | Pointer to **int32** | IP Phone VCI should be a number between 0 and 35535. | [optional] 
**IpPhoneVpi** | Pointer to **int32** | IP Phone VPI should be a number between 0 and 255. | [optional] 
**IptvVci** | Pointer to **int32** | IPTV VCI should be a number between 0 and 35535. | [optional] 
**IptvVpi** | Pointer to **int32** | IPTV VPI should be a number between 0 and 255. | [optional] 
**ModulationType** | **int32** | Modulation Type should be a value as follows: 0: ADSL, 1: VDSL. | 

## Methods

### NewIptvDslOpenApiVO

`func NewIptvDslOpenApiVO(modulationType int32, ) *IptvDslOpenApiVO`

NewIptvDslOpenApiVO instantiates a new IptvDslOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvDslOpenApiVOWithDefaults

`func NewIptvDslOpenApiVOWithDefaults() *IptvDslOpenApiVO`

NewIptvDslOpenApiVOWithDefaults instantiates a new IptvDslOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpPhoneVci

`func (o *IptvDslOpenApiVO) GetIpPhoneVci() int32`

GetIpPhoneVci returns the IpPhoneVci field if non-nil, zero value otherwise.

### GetIpPhoneVciOk

`func (o *IptvDslOpenApiVO) GetIpPhoneVciOk() (*int32, bool)`

GetIpPhoneVciOk returns a tuple with the IpPhoneVci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPhoneVci

`func (o *IptvDslOpenApiVO) SetIpPhoneVci(v int32)`

SetIpPhoneVci sets IpPhoneVci field to given value.

### HasIpPhoneVci

`func (o *IptvDslOpenApiVO) HasIpPhoneVci() bool`

HasIpPhoneVci returns a boolean if a field has been set.

### GetIpPhoneVpi

`func (o *IptvDslOpenApiVO) GetIpPhoneVpi() int32`

GetIpPhoneVpi returns the IpPhoneVpi field if non-nil, zero value otherwise.

### GetIpPhoneVpiOk

`func (o *IptvDslOpenApiVO) GetIpPhoneVpiOk() (*int32, bool)`

GetIpPhoneVpiOk returns a tuple with the IpPhoneVpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPhoneVpi

`func (o *IptvDslOpenApiVO) SetIpPhoneVpi(v int32)`

SetIpPhoneVpi sets IpPhoneVpi field to given value.

### HasIpPhoneVpi

`func (o *IptvDslOpenApiVO) HasIpPhoneVpi() bool`

HasIpPhoneVpi returns a boolean if a field has been set.

### GetIptvVci

`func (o *IptvDslOpenApiVO) GetIptvVci() int32`

GetIptvVci returns the IptvVci field if non-nil, zero value otherwise.

### GetIptvVciOk

`func (o *IptvDslOpenApiVO) GetIptvVciOk() (*int32, bool)`

GetIptvVciOk returns a tuple with the IptvVci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvVci

`func (o *IptvDslOpenApiVO) SetIptvVci(v int32)`

SetIptvVci sets IptvVci field to given value.

### HasIptvVci

`func (o *IptvDslOpenApiVO) HasIptvVci() bool`

HasIptvVci returns a boolean if a field has been set.

### GetIptvVpi

`func (o *IptvDslOpenApiVO) GetIptvVpi() int32`

GetIptvVpi returns the IptvVpi field if non-nil, zero value otherwise.

### GetIptvVpiOk

`func (o *IptvDslOpenApiVO) GetIptvVpiOk() (*int32, bool)`

GetIptvVpiOk returns a tuple with the IptvVpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvVpi

`func (o *IptvDslOpenApiVO) SetIptvVpi(v int32)`

SetIptvVpi sets IptvVpi field to given value.

### HasIptvVpi

`func (o *IptvDslOpenApiVO) HasIptvVpi() bool`

HasIptvVpi returns a boolean if a field has been set.

### GetModulationType

`func (o *IptvDslOpenApiVO) GetModulationType() int32`

GetModulationType returns the ModulationType field if non-nil, zero value otherwise.

### GetModulationTypeOk

`func (o *IptvDslOpenApiVO) GetModulationTypeOk() (*int32, bool)`

GetModulationTypeOk returns a tuple with the ModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulationType

`func (o *IptvDslOpenApiVO) SetModulationType(v int32)`

SetModulationType sets ModulationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


