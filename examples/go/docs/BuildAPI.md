# \BuildAPI

All URIs are relative to *https://api.server.test/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBuild**](BuildAPI.md#AddBuild) | **Post** /build | Add a new build to run
[**GetBuildById**](BuildAPI.md#GetBuildById) | **Get** /build/{buildId} | Return build info
[**GetBuildOutputManifestById**](BuildAPI.md#GetBuildOutputManifestById) | **Get** /build/{buildId}/outputs | return manifest of build outputs
[**GetSLSAById**](BuildAPI.md#GetSLSAById) | **Get** /build/{buildId}/slsa | return SLSA attestation



## AddBuild

> Build AddBuild(ctx).Build(build).Execute()

Add a new build to run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    build := *openapiclient.NewBuild() // Build | Create a new build

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildAPI.AddBuild(context.Background()).Build(build).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.AddBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBuild`: Build
    fmt.Fprintf(os.Stdout, "Response from `BuildAPI.AddBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **build** | [**Build**](Build.md) | Create a new build | 

### Return type

[**Build**](Build.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/x-www-form-urlencoded
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildById

> BuildStatus GetBuildById(ctx, buildId).Execute()

Return build info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    buildId := int32(56) // int32 | Id of build

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildAPI.GetBuildById(context.Background(), buildId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.GetBuildById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuildById`: BuildStatus
    fmt.Fprintf(os.Stdout, "Response from `BuildAPI.GetBuildById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **int32** | Id of build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildStatus**](BuildStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildOutputManifestById

> BuildOutputManifest GetBuildOutputManifestById(ctx, buildId).Execute()

return manifest of build outputs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    buildId := int32(56) // int32 | Id of build

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildAPI.GetBuildOutputManifestById(context.Background(), buildId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.GetBuildOutputManifestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuildOutputManifestById`: BuildOutputManifest
    fmt.Fprintf(os.Stdout, "Response from `BuildAPI.GetBuildOutputManifestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **int32** | Id of build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildOutputManifestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildOutputManifest**](BuildOutputManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSLSAById

> SlsaProvenanceV01 GetSLSAById(ctx, buildId).Execute()

return SLSA attestation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    buildId := int32(56) // int32 | Id of build

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildAPI.GetSLSAById(context.Background(), buildId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.GetSLSAById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLSAById`: SlsaProvenanceV01
    fmt.Fprintf(os.Stdout, "Response from `BuildAPI.GetSLSAById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **int32** | Id of build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSLSAByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlsaProvenanceV01**](SlsaProvenanceV01.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

