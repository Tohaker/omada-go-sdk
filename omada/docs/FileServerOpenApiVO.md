# FileServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | File server host name or IP. Parameter [filePath] should be 1 - 128 ASCII characters. | 
**Password** | Pointer to **string** | File server access password, if needed in FTP/SFTP protocol. Parameter [password] should be 1 - 128 ASCII characters. | [optional] 
**Port** | **int32** | File server port. The value must be between 1 and 65535. | 
**Protocol** | **string** | File server protocol, protocol should be a value as follows: FTP: FTP protocal; SFTP: SFTP protocal; TFTP: TFTP protocal; SCP: SCP protocal. | 
**Username** | Pointer to **string** | File server access username, if needed in FTP/SFTP/SCP protocol. Parameter [username] should be 1 - 128 ASCII characters. | [optional] 

## Methods

### NewFileServerOpenApiVO

`func NewFileServerOpenApiVO(hostname string, port int32, protocol string, ) *FileServerOpenApiVO`

NewFileServerOpenApiVO instantiates a new FileServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileServerOpenApiVOWithDefaults

`func NewFileServerOpenApiVOWithDefaults() *FileServerOpenApiVO`

NewFileServerOpenApiVOWithDefaults instantiates a new FileServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *FileServerOpenApiVO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FileServerOpenApiVO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FileServerOpenApiVO) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPassword

`func (o *FileServerOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FileServerOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FileServerOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FileServerOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *FileServerOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FileServerOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FileServerOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *FileServerOpenApiVO) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FileServerOpenApiVO) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FileServerOpenApiVO) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetUsername

`func (o *FileServerOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FileServerOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FileServerOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FileServerOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


