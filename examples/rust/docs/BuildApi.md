# \BuildApi

All URIs are relative to *https://api.server.test/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_build**](BuildApi.md#add_build) | **POST** /build | Add a new build to run
[**get_build_by_id**](BuildApi.md#get_build_by_id) | **GET** /build/{buildId} | Return build info
[**get_build_output_manifest_by_id**](BuildApi.md#get_build_output_manifest_by_id) | **GET** /build/{buildId}/outputs | return manifest of build outputs
[**get_slsaby_id**](BuildApi.md#get_slsaby_id) | **GET** /build/{buildId}/slsa | return SLSA attestation



## add_build

> crate::models::CdBuildstarted add_build(build)
Add a new build to run

Add a new build

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**build** | [**Build**](Build.md) | Create a new build | [required] |

### Return type

[**crate::models::CdBuildstarted**](cd_buildstarted.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/x-www-form-urlencoded
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_build_by_id

> crate::models::BuildStatus get_build_by_id(build_id)
Return build info

Return build info by buildId

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**build_id** | **i32** | Id of build | [required] |

### Return type

[**crate::models::BuildStatus**](BuildStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_build_output_manifest_by_id

> crate::models::BuildOutputManifest get_build_output_manifest_by_id(build_id)
return manifest of build outputs

return manifest build outputs by buildId

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**build_id** | **i32** | Id of build | [required] |

### Return type

[**crate::models::BuildOutputManifest**](BuildOutputManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_slsaby_id

> crate::models::SlsaProvenanceV01 get_slsaby_id(build_id)
return SLSA attestation

return SLSA attestation by buildId

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**build_id** | **i32** | Id of build | [required] |

### Return type

[**crate::models::SlsaProvenanceV01**](slsa_provenance_v01.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

