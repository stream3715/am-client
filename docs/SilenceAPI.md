# \SilenceAPI

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSilence**](SilenceAPI.md#DeleteSilence) | **Delete** /silence/{silenceID} | 
[**GetSilence**](SilenceAPI.md#GetSilence) | **Get** /silence/{silenceID} | 
[**GetSilences**](SilenceAPI.md#GetSilences) | **Get** /silences | 
[**PostSilences**](SilenceAPI.md#PostSilences) | **Post** /silences | 



## DeleteSilence

> DeleteSilence(ctx, silenceID).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/stream3715/am-client"
)

func main() {
	silenceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the silence to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SilenceAPI.DeleteSilence(context.Background(), silenceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SilenceAPI.DeleteSilence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**silenceID** | **string** | ID of the silence to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSilence

> GettableSilence GetSilence(ctx, silenceID).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/stream3715/am-client"
)

func main() {
	silenceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the silence to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SilenceAPI.GetSilence(context.Background(), silenceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SilenceAPI.GetSilence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSilence`: GettableSilence
	fmt.Fprintf(os.Stdout, "Response from `SilenceAPI.GetSilence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**silenceID** | **string** | ID of the silence to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GettableSilence**](GettableSilence.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSilences

> []GettableSilence GetSilences(ctx).Filter(filter).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/stream3715/am-client"
)

func main() {
	filter := []string{"Inner_example"} // []string | A list of matchers to filter silences by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SilenceAPI.GetSilences(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SilenceAPI.GetSilences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSilences`: []GettableSilence
	fmt.Fprintf(os.Stdout, "Response from `SilenceAPI.GetSilences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSilencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **[]string** | A list of matchers to filter silences by | 

### Return type

[**[]GettableSilence**](GettableSilence.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSilences

> PostSilences200Response PostSilences(ctx).Silence(silence).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/stream3715/am-client"
)

func main() {
	silence := *openapiclient.NewPostableSilence([]openapiclient.Matcher{*openapiclient.NewMatcher("Name_example", "Value_example", false)}, time.Now(), time.Now(), "CreatedBy_example", "Comment_example") // PostableSilence | The silence to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SilenceAPI.PostSilences(context.Background()).Silence(silence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SilenceAPI.PostSilences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSilences`: PostSilences200Response
	fmt.Fprintf(os.Stdout, "Response from `SilenceAPI.PostSilences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSilencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **silence** | [**PostableSilence**](PostableSilence.md) | The silence to create | 

### Return type

[**PostSilences200Response**](PostSilences200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

