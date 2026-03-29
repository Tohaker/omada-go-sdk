# EapMethodOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **int32** | EAP authentication method supported by the NAI realm.&lt;br /&gt;Parameter method should be a value as follows:[0:None;1:Identity;2:Notification;5:One-Time Password (OTP);6:Generic Token Card (GTC);13:EAP-TLS;18:GSM Subscriber Identity Modules (EAP-SIM);21:EAP-TTLS;23:EAP-AKA Authentication;25:PEAP;28:CRYPTOCard;29:EAP-MSCHAP-V2] | 
**Param** | [**[]AuthenticationParamOpenApiVO**](AuthenticationParamOpenApiVO.md) | Authentication Param list, configure the EAP authentication parameter identifier and authentication parameters.&lt;br /&gt;Note: Up to 4 entries are allowed for the Authentication Param list. | 

## Methods

### NewEapMethodOpenApiVO

`func NewEapMethodOpenApiVO(method int32, param []AuthenticationParamOpenApiVO, ) *EapMethodOpenApiVO`

NewEapMethodOpenApiVO instantiates a new EapMethodOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapMethodOpenApiVOWithDefaults

`func NewEapMethodOpenApiVOWithDefaults() *EapMethodOpenApiVO`

NewEapMethodOpenApiVOWithDefaults instantiates a new EapMethodOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *EapMethodOpenApiVO) GetMethod() int32`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EapMethodOpenApiVO) GetMethodOk() (*int32, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EapMethodOpenApiVO) SetMethod(v int32)`

SetMethod sets Method field to given value.


### GetParam

`func (o *EapMethodOpenApiVO) GetParam() []AuthenticationParamOpenApiVO`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *EapMethodOpenApiVO) GetParamOk() (*[]AuthenticationParamOpenApiVO, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *EapMethodOpenApiVO) SetParam(v []AuthenticationParamOpenApiVO)`

SetParam sets Param field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


