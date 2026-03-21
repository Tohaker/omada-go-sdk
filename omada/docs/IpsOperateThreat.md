# IpsOperateThreat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockName** | Pointer to **string** | IPS block name should not be empty when type is 0 or 1. | [optional] 
**SignatureSuppression** | Pointer to [**SignatureSuppression**](SignatureSuppression.md) |  | [optional] 
**ThreatId** | Pointer to [**[]IpsOperateThreatIdAndTime**](IpsOperateThreatIdAndTime.md) | IPS signature type should be a value as follows: 0: all traffic; 1: packet tracking respectively. | [optional] 
**Type** | **int32** | IPS Operate Threat type should be a value as follows: 0: block; 1:isolate device; 2: signature Suppression; 3: allow | 

## Methods

### NewIpsOperateThreat

`func NewIpsOperateThreat(type_ int32, ) *IpsOperateThreat`

NewIpsOperateThreat instantiates a new IpsOperateThreat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsOperateThreatWithDefaults

`func NewIpsOperateThreatWithDefaults() *IpsOperateThreat`

NewIpsOperateThreatWithDefaults instantiates a new IpsOperateThreat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockName

`func (o *IpsOperateThreat) GetBlockName() string`

GetBlockName returns the BlockName field if non-nil, zero value otherwise.

### GetBlockNameOk

`func (o *IpsOperateThreat) GetBlockNameOk() (*string, bool)`

GetBlockNameOk returns a tuple with the BlockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockName

`func (o *IpsOperateThreat) SetBlockName(v string)`

SetBlockName sets BlockName field to given value.

### HasBlockName

`func (o *IpsOperateThreat) HasBlockName() bool`

HasBlockName returns a boolean if a field has been set.

### GetSignatureSuppression

`func (o *IpsOperateThreat) GetSignatureSuppression() SignatureSuppression`

GetSignatureSuppression returns the SignatureSuppression field if non-nil, zero value otherwise.

### GetSignatureSuppressionOk

`func (o *IpsOperateThreat) GetSignatureSuppressionOk() (*SignatureSuppression, bool)`

GetSignatureSuppressionOk returns a tuple with the SignatureSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureSuppression

`func (o *IpsOperateThreat) SetSignatureSuppression(v SignatureSuppression)`

SetSignatureSuppression sets SignatureSuppression field to given value.

### HasSignatureSuppression

`func (o *IpsOperateThreat) HasSignatureSuppression() bool`

HasSignatureSuppression returns a boolean if a field has been set.

### GetThreatId

`func (o *IpsOperateThreat) GetThreatId() []IpsOperateThreatIdAndTime`

GetThreatId returns the ThreatId field if non-nil, zero value otherwise.

### GetThreatIdOk

`func (o *IpsOperateThreat) GetThreatIdOk() (*[]IpsOperateThreatIdAndTime, bool)`

GetThreatIdOk returns a tuple with the ThreatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatId

`func (o *IpsOperateThreat) SetThreatId(v []IpsOperateThreatIdAndTime)`

SetThreatId sets ThreatId field to given value.

### HasThreatId

`func (o *IpsOperateThreat) HasThreatId() bool`

HasThreatId returns a boolean if a field has been set.

### GetType

`func (o *IpsOperateThreat) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpsOperateThreat) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpsOperateThreat) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


