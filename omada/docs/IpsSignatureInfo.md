# IpsSignatureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **int32** | Attack service wide category. | 
**Classification** | **string** | Attack service concrete classification. | 
**Direction** | Pointer to **int32** | Direction should be a value as follow: 0: both direction; 1: source direction; 2: destination direction | [optional] 
**Id** | **string** | The unique identity of the signature suppresses. | 
**Service** | **string** | Attack service service, the same as parameter[signature] | 
**Sid** | **int64** | IPS signature rule SID which is from device. | 
**Signature** | **string** | Attack service signature. | 
**TrafficSource** | Pointer to **string** | IPS signature traffic Source. If parameter [trafficType] is 0, parameter [trafficSource] should be IPV4 address. If parameter [trafficType] is 1, parameter [trafficSource] should be subnet address. | [optional] 
**TrafficType** | Pointer to **int32** | TrafficType should be a value as follow: 0: ip address; 1: subnet | [optional] 
**Type** | **int32** | Type should be a value as follow: 0: all traffic; 1: packet tracking | 

## Methods

### NewIpsSignatureInfo

`func NewIpsSignatureInfo(category int32, classification string, id string, service string, sid int64, signature string, type_ int32, ) *IpsSignatureInfo`

NewIpsSignatureInfo instantiates a new IpsSignatureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsSignatureInfoWithDefaults

`func NewIpsSignatureInfoWithDefaults() *IpsSignatureInfo`

NewIpsSignatureInfoWithDefaults instantiates a new IpsSignatureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *IpsSignatureInfo) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IpsSignatureInfo) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IpsSignatureInfo) SetCategory(v int32)`

SetCategory sets Category field to given value.


### GetClassification

`func (o *IpsSignatureInfo) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *IpsSignatureInfo) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *IpsSignatureInfo) SetClassification(v string)`

SetClassification sets Classification field to given value.


### GetDirection

`func (o *IpsSignatureInfo) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IpsSignatureInfo) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IpsSignatureInfo) SetDirection(v int32)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *IpsSignatureInfo) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *IpsSignatureInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsSignatureInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsSignatureInfo) SetId(v string)`

SetId sets Id field to given value.


### GetService

`func (o *IpsSignatureInfo) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IpsSignatureInfo) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IpsSignatureInfo) SetService(v string)`

SetService sets Service field to given value.


### GetSid

`func (o *IpsSignatureInfo) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpsSignatureInfo) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpsSignatureInfo) SetSid(v int64)`

SetSid sets Sid field to given value.


### GetSignature

`func (o *IpsSignatureInfo) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *IpsSignatureInfo) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *IpsSignatureInfo) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTrafficSource

`func (o *IpsSignatureInfo) GetTrafficSource() string`

GetTrafficSource returns the TrafficSource field if non-nil, zero value otherwise.

### GetTrafficSourceOk

`func (o *IpsSignatureInfo) GetTrafficSourceOk() (*string, bool)`

GetTrafficSourceOk returns a tuple with the TrafficSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSource

`func (o *IpsSignatureInfo) SetTrafficSource(v string)`

SetTrafficSource sets TrafficSource field to given value.

### HasTrafficSource

`func (o *IpsSignatureInfo) HasTrafficSource() bool`

HasTrafficSource returns a boolean if a field has been set.

### GetTrafficType

`func (o *IpsSignatureInfo) GetTrafficType() int32`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *IpsSignatureInfo) GetTrafficTypeOk() (*int32, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *IpsSignatureInfo) SetTrafficType(v int32)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *IpsSignatureInfo) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.

### GetType

`func (o *IpsSignatureInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpsSignatureInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpsSignatureInfo) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


