# AuthenticationParamOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | EAP authentication parameter identifier.&lt;br /&gt;Parameter id should be a value as follows: [2: None-EAP Inner Authentication Type; 3: Inner Authentication EAP Method Type; 5: Credential Type; 6: Tunneled EAP Method Credential Type] | 
**Value** | **int32** | Authentication parameters.&lt;br /&gt;When id &#x3D; 2, parameter value should be a value as follows:[1:PAP;2:CHAP;3:MSCHAP;4:MSCHAPV2].&lt;br /&gt;When id &#x3D; 3, parameter value should be a value as follows:[0:None;1:Identity;2:Notification;5:One-Time Password (OTP);6:Generic Token Card (GTC);13:EAP-TLS;18:GSM Subscriber Identity Modules (EAP-SIM);21:EAP-TTLS;23:EAP-AKA Authentication;25:PEAP;28:CRYPTOCard;29:EAP-MSCHAP-V2]&lt;br /&gt;When id &#x3D; 5, parameter value should be a value as follows:[1:SIM;2:USIM;3:NFC Secure Element;4:Hardware Token;5:Softoken;6:Certificate;7:username/password;8:None;10:Vendor Specific]&lt;br /&gt;When id &#x3D; 6, parameter value should be a value as follows:[1:SIM;2:USIM;3:NFC Secure Element;4:Hardware Token;5:Softoken;6:Certificate;7:username/password;9:Anonymous;10:Vendor Specific]. | 

## Methods

### NewAuthenticationParamOpenApiVO

`func NewAuthenticationParamOpenApiVO(id int32, value int32, ) *AuthenticationParamOpenApiVO`

NewAuthenticationParamOpenApiVO instantiates a new AuthenticationParamOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationParamOpenApiVOWithDefaults

`func NewAuthenticationParamOpenApiVOWithDefaults() *AuthenticationParamOpenApiVO`

NewAuthenticationParamOpenApiVOWithDefaults instantiates a new AuthenticationParamOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationParamOpenApiVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationParamOpenApiVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationParamOpenApiVO) SetId(v int32)`

SetId sets Id field to given value.


### GetValue

`func (o *AuthenticationParamOpenApiVO) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthenticationParamOpenApiVO) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthenticationParamOpenApiVO) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


