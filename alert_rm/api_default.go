/*
 * Alert Runtime Manager
 *
 * Description of the Alert in Runtime Manager API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alert_rm

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiApiAddnewalertRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId *string
	alert *Alert
	authorization *string
	envId *string
}

func (r DefaultApiApiAddnewalertRequest) OrgId(orgId string) DefaultApiApiAddnewalertRequest {
	r.orgId = &orgId
	return r
}
func (r DefaultApiApiAddnewalertRequest) Alert(alert Alert) DefaultApiApiAddnewalertRequest {
	r.alert = &alert
	return r
}
func (r DefaultApiApiAddnewalertRequest) Authorization(authorization string) DefaultApiApiAddnewalertRequest {
	r.authorization = &authorization
	return r
}
func (r DefaultApiApiAddnewalertRequest) EnvId(envId string) DefaultApiApiAddnewalertRequest {
	r.envId = &envId
	return r
}

func (r DefaultApiApiAddnewalertRequest) Execute() (AlertWithId, *_nethttp.Response, error) {
	return r.ApiService.AddnewalertExecute(r)
}

/*
 * Addnewalert Method for Addnewalert
 * Add new alert
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DefaultApiApiAddnewalertRequest
 */
func (a *DefaultApiService) Addnewalert(ctx _context.Context) DefaultApiApiAddnewalertRequest {
	return DefaultApiApiAddnewalertRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AlertWithId
 */
func (a *DefaultApiService) AddnewalertExecute(r DefaultApiApiAddnewalertRequest) (AlertWithId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlertWithId
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.Addnewalert")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alerts/cloudhub"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}
	if r.alert == nil {
		return localVarReturnValue, nil, reportError("alert is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	localVarHeaderParams["orgId"] = parameterToString(*r.orgId, "")
	if r.envId != nil {
		localVarHeaderParams["envId"] = parameterToString(*r.envId, "")
	}
	// body params
	localVarPostBody = r.alert
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiDeleteonealertRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId *string
	envId *string
	cloudhubRmAlert string
	authorization *string
}

func (r DefaultApiApiDeleteonealertRequest) OrgId(orgId string) DefaultApiApiDeleteonealertRequest {
	r.orgId = &orgId
	return r
}
func (r DefaultApiApiDeleteonealertRequest) EnvId(envId string) DefaultApiApiDeleteonealertRequest {
	r.envId = &envId
	return r
}
func (r DefaultApiApiDeleteonealertRequest) Authorization(authorization string) DefaultApiApiDeleteonealertRequest {
	r.authorization = &authorization
	return r
}

func (r DefaultApiApiDeleteonealertRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteonealertExecute(r)
}

/*
 * Deleteonealert Deleteonealert
 * Delete one alert
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cloudhubRmAlert the name of the cloud hub alert
 * @return DefaultApiApiDeleteonealertRequest
 */
func (a *DefaultApiService) Deleteonealert(ctx _context.Context, cloudhubRmAlert string) DefaultApiApiDeleteonealertRequest {
	return DefaultApiApiDeleteonealertRequest{
		ApiService: a,
		ctx: ctx,
		cloudhubRmAlert: cloudhubRmAlert,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) DeleteonealertExecute(r DefaultApiApiDeleteonealertRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.Deleteonealert")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alerts/cloudhub/{cloudhub-rm-alert}"
	localVarPath = strings.Replace(localVarPath, "{"+"cloudhub-rm-alert"+"}", _neturl.PathEscape(parameterToString(r.cloudhubRmAlert, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return nil, reportError("orgId is required and must be specified")
	}
	if r.envId == nil {
		return nil, reportError("envId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["orgId"] = parameterToString(*r.orgId, "")
	localVarHeaderParams["envId"] = parameterToString(*r.envId, "")
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiApiModifyOneAlertRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId *string
	envId *string
	cloudhubRmAlert string
	authorization *string
	alertWithId *AlertWithId
}

func (r DefaultApiApiModifyOneAlertRequest) OrgId(orgId string) DefaultApiApiModifyOneAlertRequest {
	r.orgId = &orgId
	return r
}
func (r DefaultApiApiModifyOneAlertRequest) EnvId(envId string) DefaultApiApiModifyOneAlertRequest {
	r.envId = &envId
	return r
}
func (r DefaultApiApiModifyOneAlertRequest) Authorization(authorization string) DefaultApiApiModifyOneAlertRequest {
	r.authorization = &authorization
	return r
}
func (r DefaultApiApiModifyOneAlertRequest) AlertWithId(alertWithId AlertWithId) DefaultApiApiModifyOneAlertRequest {
	r.alertWithId = &alertWithId
	return r
}

func (r DefaultApiApiModifyOneAlertRequest) Execute() (AlertWithId, *_nethttp.Response, error) {
	return r.ApiService.ModifyOneAlertExecute(r)
}

/*
 * ModifyOneAlert Method for ModifyOneAlert
 * Modify one alert
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cloudhubRmAlert the name of the cloud hub alert
 * @return DefaultApiApiModifyOneAlertRequest
 */
func (a *DefaultApiService) ModifyOneAlert(ctx _context.Context, cloudhubRmAlert string) DefaultApiApiModifyOneAlertRequest {
	return DefaultApiApiModifyOneAlertRequest{
		ApiService: a,
		ctx: ctx,
		cloudhubRmAlert: cloudhubRmAlert,
	}
}

/*
 * Execute executes the request
 * @return AlertWithId
 */
func (a *DefaultApiService) ModifyOneAlertExecute(r DefaultApiApiModifyOneAlertRequest) (AlertWithId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlertWithId
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ModifyOneAlert")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alerts/cloudhub/{cloudhub-rm-alert}"
	localVarPath = strings.Replace(localVarPath, "{"+"cloudhub-rm-alert"+"}", _neturl.PathEscape(parameterToString(r.cloudhubRmAlert, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}
	if r.envId == nil {
		return localVarReturnValue, nil, reportError("envId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["orgId"] = parameterToString(*r.orgId, "")
	localVarHeaderParams["envId"] = parameterToString(*r.envId, "")
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	// body params
	localVarPostBody = r.alertWithId
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId *string
	envId *string
	cloudhubRmAlert string
	authorization *string
}

func (r DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest) OrgId(orgId string) DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest {
	r.orgId = &orgId
	return r
}
func (r DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest) EnvId(envId string) DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest {
	r.envId = &envId
	return r
}
func (r DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest) Authorization(authorization string) DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest {
	r.authorization = &authorization
	return r
}

func (r DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest) Execute() ([]AlertWithId, *_nethttp.Response, error) {
	return r.ApiService.V1AlertsCloudhubCloudhubRmAlertGetExecute(r)
}

/*
 * V1AlertsCloudhubCloudhubRmAlertGet Get one specific alert
 * Get one specific alert
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cloudhubRmAlert the name of the cloud hub alert
 * @return DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest
 */
func (a *DefaultApiService) V1AlertsCloudhubCloudhubRmAlertGet(ctx _context.Context, cloudhubRmAlert string) DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest {
	return DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest{
		ApiService: a,
		ctx: ctx,
		cloudhubRmAlert: cloudhubRmAlert,
	}
}

/*
 * Execute executes the request
 * @return []AlertWithId
 */
func (a *DefaultApiService) V1AlertsCloudhubCloudhubRmAlertGetExecute(r DefaultApiApiV1AlertsCloudhubCloudhubRmAlertGetRequest) ([]AlertWithId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AlertWithId
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.V1AlertsCloudhubCloudhubRmAlertGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alerts/cloudhub/{cloudhub-rm-alert}"
	localVarPath = strings.Replace(localVarPath, "{"+"cloudhub-rm-alert"+"}", _neturl.PathEscape(parameterToString(r.cloudhubRmAlert, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}
	if r.envId == nil {
		return localVarReturnValue, nil, reportError("envId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["orgId"] = parameterToString(*r.orgId, "")
	localVarHeaderParams["envId"] = parameterToString(*r.envId, "")
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiV1AlertsGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId *string
	envId *string
}

func (r DefaultApiApiV1AlertsGetRequest) OrgId(orgId string) DefaultApiApiV1AlertsGetRequest {
	r.orgId = &orgId
	return r
}
func (r DefaultApiApiV1AlertsGetRequest) EnvId(envId string) DefaultApiApiV1AlertsGetRequest {
	r.envId = &envId
	return r
}

func (r DefaultApiApiV1AlertsGetRequest) Execute() ([]AlertWithId, *_nethttp.Response, error) {
	return r.ApiService.V1AlertsGetExecute(r)
}

/*
 * V1AlertsGet Get all the alert instances from the run time manager
 * Get alerts
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DefaultApiApiV1AlertsGetRequest
 */
func (a *DefaultApiService) V1AlertsGet(ctx _context.Context) DefaultApiApiV1AlertsGetRequest {
	return DefaultApiApiV1AlertsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []AlertWithId
 */
func (a *DefaultApiService) V1AlertsGetExecute(r DefaultApiApiV1AlertsGetRequest) ([]AlertWithId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AlertWithId
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.V1AlertsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alerts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}
	if r.envId == nil {
		return localVarReturnValue, nil, reportError("envId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["orgId"] = parameterToString(*r.orgId, "")
	localVarHeaderParams["envId"] = parameterToString(*r.envId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
