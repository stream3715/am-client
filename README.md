# Go API client for openapi

API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/stream3715/am-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlertAPI* | [**GetAlerts**](docs/AlertAPI.md#getalerts) | **Get** /alerts | 
*AlertAPI* | [**PostAlerts**](docs/AlertAPI.md#postalerts) | **Post** /alerts | 
*AlertgroupAPI* | [**GetAlertGroups**](docs/AlertgroupAPI.md#getalertgroups) | **Get** /alerts/groups | 
*GeneralAPI* | [**GetStatus**](docs/GeneralAPI.md#getstatus) | **Get** /status | 
*ReceiverAPI* | [**GetReceivers**](docs/ReceiverAPI.md#getreceivers) | **Get** /receivers | 
*SilenceAPI* | [**DeleteSilence**](docs/SilenceAPI.md#deletesilence) | **Delete** /silence/{silenceID} | 
*SilenceAPI* | [**GetSilence**](docs/SilenceAPI.md#getsilence) | **Get** /silence/{silenceID} | 
*SilenceAPI* | [**GetSilences**](docs/SilenceAPI.md#getsilences) | **Get** /silences | 
*SilenceAPI* | [**PostSilences**](docs/SilenceAPI.md#postsilences) | **Post** /silences | 


## Documentation For Models

 - [Alert](docs/Alert.md)
 - [AlertGroup](docs/AlertGroup.md)
 - [AlertStatus](docs/AlertStatus.md)
 - [AlertmanagerConfig](docs/AlertmanagerConfig.md)
 - [AlertmanagerStatus](docs/AlertmanagerStatus.md)
 - [ClusterStatus](docs/ClusterStatus.md)
 - [GettableAlert](docs/GettableAlert.md)
 - [GettableSilence](docs/GettableSilence.md)
 - [Matcher](docs/Matcher.md)
 - [PeerStatus](docs/PeerStatus.md)
 - [PostSilences200Response](docs/PostSilences200Response.md)
 - [PostableAlert](docs/PostableAlert.md)
 - [PostableSilence](docs/PostableSilence.md)
 - [Receiver](docs/Receiver.md)
 - [Silence](docs/Silence.md)
 - [SilenceStatus](docs/SilenceStatus.md)
 - [VersionInfo](docs/VersionInfo.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



