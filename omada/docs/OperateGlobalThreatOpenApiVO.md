# OperateGlobalThreatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockName** | Pointer to **string** | Block Name. | [optional] 
**SignatureSuppression** | Pointer to [**SignatureSuppression**](SignatureSuppression.md) |  | [optional] 
**ThreatId** | Pointer to [**[]SiteTimeIdOpenApiVO**](SiteTimeIdOpenApiVO.md) | The global view needs to pass in a siteId. | [optional] 
**Type** | **int32** | type,0 - block, 1 - isolate device, 2 - allow from all ips, 3 - allow this ip. | 

## Methods

### NewOperateGlobalThreatOpenApiVO

`func NewOperateGlobalThreatOpenApiVO(type_ int32, ) *OperateGlobalThreatOpenApiVO`

NewOperateGlobalThreatOpenApiVO instantiates a new OperateGlobalThreatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperateGlobalThreatOpenApiVOWithDefaults

`func NewOperateGlobalThreatOpenApiVOWithDefaults() *OperateGlobalThreatOpenApiVO`

NewOperateGlobalThreatOpenApiVOWithDefaults instantiates a new OperateGlobalThreatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockName

`func (o *OperateGlobalThreatOpenApiVO) GetBlockName() string`

GetBlockName returns the BlockName field if non-nil, zero value otherwise.

### GetBlockNameOk

`func (o *OperateGlobalThreatOpenApiVO) GetBlockNameOk() (*string, bool)`

GetBlockNameOk returns a tuple with the BlockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockName

`func (o *OperateGlobalThreatOpenApiVO) SetBlockName(v string)`

SetBlockName sets BlockName field to given value.

### HasBlockName

`func (o *OperateGlobalThreatOpenApiVO) HasBlockName() bool`

HasBlockName returns a boolean if a field has been set.

### GetSignatureSuppression

`func (o *OperateGlobalThreatOpenApiVO) GetSignatureSuppression() SignatureSuppression`

GetSignatureSuppression returns the SignatureSuppression field if non-nil, zero value otherwise.

### GetSignatureSuppressionOk

`func (o *OperateGlobalThreatOpenApiVO) GetSignatureSuppressionOk() (*SignatureSuppression, bool)`

GetSignatureSuppressionOk returns a tuple with the SignatureSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureSuppression

`func (o *OperateGlobalThreatOpenApiVO) SetSignatureSuppression(v SignatureSuppression)`

SetSignatureSuppression sets SignatureSuppression field to given value.

### HasSignatureSuppression

`func (o *OperateGlobalThreatOpenApiVO) HasSignatureSuppression() bool`

HasSignatureSuppression returns a boolean if a field has been set.

### GetThreatId

`func (o *OperateGlobalThreatOpenApiVO) GetThreatId() []SiteTimeIdOpenApiVO`

GetThreatId returns the ThreatId field if non-nil, zero value otherwise.

### GetThreatIdOk

`func (o *OperateGlobalThreatOpenApiVO) GetThreatIdOk() (*[]SiteTimeIdOpenApiVO, bool)`

GetThreatIdOk returns a tuple with the ThreatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatId

`func (o *OperateGlobalThreatOpenApiVO) SetThreatId(v []SiteTimeIdOpenApiVO)`

SetThreatId sets ThreatId field to given value.

### HasThreatId

`func (o *OperateGlobalThreatOpenApiVO) HasThreatId() bool`

HasThreatId returns a boolean if a field has been set.

### GetType

`func (o *OperateGlobalThreatOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperateGlobalThreatOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperateGlobalThreatOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


