# SignatureSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **int32** | Direction should be a value as follow: 0: both direction; 1: source direction; 2: destination direction | [optional] 
**Ip** | Pointer to **string** | IPS signature traffic source. If parameter [trackBy] is 0, parameter [ip] is needed.  | [optional] 
**Subnet** | Pointer to **string** | IPS signature traffic source. If parameter [trackBy] is 1, parameter [subnet] is needed. | [optional] 
**TrackBy** | Pointer to **int32** | TrackBy should be a value as follow: 0: ip address; 1: subnet | [optional] 
**Type** | **int32** | Type should be a value as follow: 0: all traffic; 1: packet tracking. | 

## Methods

### NewSignatureSuppression

`func NewSignatureSuppression(type_ int32, ) *SignatureSuppression`

NewSignatureSuppression instantiates a new SignatureSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureSuppressionWithDefaults

`func NewSignatureSuppressionWithDefaults() *SignatureSuppression`

NewSignatureSuppressionWithDefaults instantiates a new SignatureSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SignatureSuppression) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SignatureSuppression) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SignatureSuppression) SetDirection(v int32)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SignatureSuppression) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetIp

`func (o *SignatureSuppression) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SignatureSuppression) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SignatureSuppression) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SignatureSuppression) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSubnet

`func (o *SignatureSuppression) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *SignatureSuppression) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *SignatureSuppression) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *SignatureSuppression) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetTrackBy

`func (o *SignatureSuppression) GetTrackBy() int32`

GetTrackBy returns the TrackBy field if non-nil, zero value otherwise.

### GetTrackByOk

`func (o *SignatureSuppression) GetTrackByOk() (*int32, bool)`

GetTrackByOk returns a tuple with the TrackBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackBy

`func (o *SignatureSuppression) SetTrackBy(v int32)`

SetTrackBy sets TrackBy field to given value.

### HasTrackBy

`func (o *SignatureSuppression) HasTrackBy() bool`

HasTrackBy returns a boolean if a field has been set.

### GetType

`func (o *SignatureSuppression) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignatureSuppression) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignatureSuppression) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


