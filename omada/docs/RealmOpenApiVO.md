# RealmOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eap** | [**[]EapMethodOpenApiVO**](EapMethodOpenApiVO.md) | EAP Method list.&lt;br /&gt;Note: Up to 4 entries are allowed for the EAP Method list. | 
**Encoding** | **int32** | Encoding format.&lt;br /&gt;Parameter encoding should be a value as follows:[0:RFC4282;1:UTF-8]. | 
**Name** | **string** | The name of the NAI realm. Usually the domain name of the service provider.&lt;br /&gt;Note: It should contain 1 to 64 UTF-8 characters. | 

## Methods

### NewRealmOpenApiVO

`func NewRealmOpenApiVO(eap []EapMethodOpenApiVO, encoding int32, name string, ) *RealmOpenApiVO`

NewRealmOpenApiVO instantiates a new RealmOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmOpenApiVOWithDefaults

`func NewRealmOpenApiVOWithDefaults() *RealmOpenApiVO`

NewRealmOpenApiVOWithDefaults instantiates a new RealmOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEap

`func (o *RealmOpenApiVO) GetEap() []EapMethodOpenApiVO`

GetEap returns the Eap field if non-nil, zero value otherwise.

### GetEapOk

`func (o *RealmOpenApiVO) GetEapOk() (*[]EapMethodOpenApiVO, bool)`

GetEapOk returns a tuple with the Eap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEap

`func (o *RealmOpenApiVO) SetEap(v []EapMethodOpenApiVO)`

SetEap sets Eap field to given value.


### GetEncoding

`func (o *RealmOpenApiVO) GetEncoding() int32`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *RealmOpenApiVO) GetEncodingOk() (*int32, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *RealmOpenApiVO) SetEncoding(v int32)`

SetEncoding sets Encoding field to given value.


### GetName

`func (o *RealmOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealmOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealmOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


