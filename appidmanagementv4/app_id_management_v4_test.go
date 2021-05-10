/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package appidmanagementv4_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/appid-go-sdk/appidmanagementv4"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`AppIdManagementV4`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListApplications(listApplicationsOptions *ListApplicationsOptions) - Operation response error`, func() {
		tenantID := "testString"
		listApplicationsPath := "/management/v4/testString/applications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApplicationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListApplications with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(appidmanagementv4.ListApplicationsOptions)
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListApplications(listApplicationsOptions *ListApplicationsOptions)`, func() {
		tenantID := "testString"
		listApplicationsPath := "/management/v4/testString/applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}]}`)
				}))
			})
			It(`Invoke ListApplications successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(appidmanagementv4.ListApplicationsOptions)
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.ListApplicationsWithContext(ctx, listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.ListApplicationsWithContext(ctx, listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}]}`)
				}))
			})
			It(`Invoke ListApplications successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.ListApplications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(appidmanagementv4.ListApplicationsOptions)
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListApplications with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(appidmanagementv4.ListApplicationsOptions)
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RegisterApplication(registerApplicationOptions *RegisterApplicationOptions) - Operation response error`, func() {
		tenantID := "testString"
		registerApplicationPath := "/management/v4/testString/applications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerApplicationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RegisterApplication with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the RegisterApplicationOptions model
				registerApplicationOptionsModel := new(appidmanagementv4.RegisterApplicationOptions)
				registerApplicationOptionsModel.Name = core.StringPtr("testString")
				registerApplicationOptionsModel.Type = core.StringPtr("testString")
				registerApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.RegisterApplication(registerApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.RegisterApplication(registerApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RegisterApplication(registerApplicationOptions *RegisterApplicationOptions)`, func() {
		tenantID := "testString"
		registerApplicationPath := "/management/v4/testString/applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerApplicationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}`)
				}))
			})
			It(`Invoke RegisterApplication successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the RegisterApplicationOptions model
				registerApplicationOptionsModel := new(appidmanagementv4.RegisterApplicationOptions)
				registerApplicationOptionsModel.Name = core.StringPtr("testString")
				registerApplicationOptionsModel.Type = core.StringPtr("testString")
				registerApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.RegisterApplicationWithContext(ctx, registerApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.RegisterApplication(registerApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.RegisterApplicationWithContext(ctx, registerApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerApplicationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}`)
				}))
			})
			It(`Invoke RegisterApplication successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.RegisterApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RegisterApplicationOptions model
				registerApplicationOptionsModel := new(appidmanagementv4.RegisterApplicationOptions)
				registerApplicationOptionsModel.Name = core.StringPtr("testString")
				registerApplicationOptionsModel.Type = core.StringPtr("testString")
				registerApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.RegisterApplication(registerApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RegisterApplication with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the RegisterApplicationOptions model
				registerApplicationOptionsModel := new(appidmanagementv4.RegisterApplicationOptions)
				registerApplicationOptionsModel.Name = core.StringPtr("testString")
				registerApplicationOptionsModel.Type = core.StringPtr("testString")
				registerApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.RegisterApplication(registerApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RegisterApplicationOptions model with no property values
				registerApplicationOptionsModelNew := new(appidmanagementv4.RegisterApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.RegisterApplication(registerApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplication(getApplicationOptions *GetApplicationOptions) - Operation response error`, func() {
		tenantID := "testString"
		getApplicationPath := "/management/v4/testString/applications/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplication with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(appidmanagementv4.GetApplicationOptions)
				getApplicationOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetApplication(getApplicationOptions *GetApplicationOptions)`, func() {
		tenantID := "testString"
		getApplicationPath := "/management/v4/testString/applications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}`)
				}))
			})
			It(`Invoke GetApplication successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(appidmanagementv4.GetApplicationOptions)
				getApplicationOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetApplicationWithContext(ctx, getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetApplicationWithContext(ctx, getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}`)
				}))
			})
			It(`Invoke GetApplication successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(appidmanagementv4.GetApplicationOptions)
				getApplicationOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplication with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(appidmanagementv4.GetApplicationOptions)
				getApplicationOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationOptions model with no property values
				getApplicationOptionsModelNew := new(appidmanagementv4.GetApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetApplication(getApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateApplication(updateApplicationOptions *UpdateApplicationOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateApplicationPath := "/management/v4/testString/applications/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApplicationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateApplication with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateApplicationOptions model
				updateApplicationOptionsModel := new(appidmanagementv4.UpdateApplicationOptions)
				updateApplicationOptionsModel.ClientID = core.StringPtr("testString")
				updateApplicationOptionsModel.Name = core.StringPtr("testString")
				updateApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateApplication(updateApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateApplication(updateApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateApplication(updateApplicationOptions *UpdateApplicationOptions)`, func() {
		tenantID := "testString"
		updateApplicationPath := "/management/v4/testString/applications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApplicationPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateApplication successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateApplicationOptions model
				updateApplicationOptionsModel := new(appidmanagementv4.UpdateApplicationOptions)
				updateApplicationOptionsModel.ClientID = core.StringPtr("testString")
				updateApplicationOptionsModel.Name = core.StringPtr("testString")
				updateApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateApplicationWithContext(ctx, updateApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateApplication(updateApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateApplicationWithContext(ctx, updateApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApplicationPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"clientId": "ClientID", "tenantId": "TenantID", "secret": "Secret", "name": "Name", "oAuthServerUrl": "OAuthServerURL", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateApplication successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateApplicationOptions model
				updateApplicationOptionsModel := new(appidmanagementv4.UpdateApplicationOptions)
				updateApplicationOptionsModel.ClientID = core.StringPtr("testString")
				updateApplicationOptionsModel.Name = core.StringPtr("testString")
				updateApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateApplication(updateApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateApplication with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateApplicationOptions model
				updateApplicationOptionsModel := new(appidmanagementv4.UpdateApplicationOptions)
				updateApplicationOptionsModel.ClientID = core.StringPtr("testString")
				updateApplicationOptionsModel.Name = core.StringPtr("testString")
				updateApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateApplication(updateApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateApplicationOptions model with no property values
				updateApplicationOptionsModelNew := new(appidmanagementv4.UpdateApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateApplication(updateApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteApplication(deleteApplicationOptions *DeleteApplicationOptions)`, func() {
		tenantID := "testString"
		deleteApplicationPath := "/management/v4/testString/applications/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteApplicationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteApplication successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.DeleteApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteApplicationOptions model
				deleteApplicationOptionsModel := new(appidmanagementv4.DeleteApplicationOptions)
				deleteApplicationOptionsModel.ClientID = core.StringPtr("testString")
				deleteApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.DeleteApplication(deleteApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteApplication with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteApplicationOptions model
				deleteApplicationOptionsModel := new(appidmanagementv4.DeleteApplicationOptions)
				deleteApplicationOptionsModel.ClientID = core.StringPtr("testString")
				deleteApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.DeleteApplication(deleteApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteApplicationOptions model with no property values
				deleteApplicationOptionsModelNew := new(appidmanagementv4.DeleteApplicationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.DeleteApplication(deleteApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplicationScopes(getApplicationScopesOptions *GetApplicationScopesOptions) - Operation response error`, func() {
		tenantID := "testString"
		getApplicationScopesPath := "/management/v4/testString/applications/testString/scopes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationScopesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplicationScopes with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetApplicationScopesOptions model
				getApplicationScopesOptionsModel := new(appidmanagementv4.GetApplicationScopesOptions)
				getApplicationScopesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetApplicationScopes(getApplicationScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetApplicationScopes(getApplicationScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetApplicationScopes(getApplicationScopesOptions *GetApplicationScopesOptions)`, func() {
		tenantID := "testString"
		getApplicationScopesPath := "/management/v4/testString/applications/testString/scopes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationScopesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": ["Scopes"]}`)
				}))
			})
			It(`Invoke GetApplicationScopes successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationScopesOptions model
				getApplicationScopesOptionsModel := new(appidmanagementv4.GetApplicationScopesOptions)
				getApplicationScopesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetApplicationScopesWithContext(ctx, getApplicationScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetApplicationScopes(getApplicationScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetApplicationScopesWithContext(ctx, getApplicationScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationScopesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": ["Scopes"]}`)
				}))
			})
			It(`Invoke GetApplicationScopes successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetApplicationScopes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationScopesOptions model
				getApplicationScopesOptionsModel := new(appidmanagementv4.GetApplicationScopesOptions)
				getApplicationScopesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetApplicationScopes(getApplicationScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplicationScopes with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetApplicationScopesOptions model
				getApplicationScopesOptionsModel := new(appidmanagementv4.GetApplicationScopesOptions)
				getApplicationScopesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetApplicationScopes(getApplicationScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationScopesOptions model with no property values
				getApplicationScopesOptionsModelNew := new(appidmanagementv4.GetApplicationScopesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetApplicationScopes(getApplicationScopesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutApplicationsScopes(putApplicationsScopesOptions *PutApplicationsScopesOptions) - Operation response error`, func() {
		tenantID := "testString"
		putApplicationsScopesPath := "/management/v4/testString/applications/testString/scopes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putApplicationsScopesPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutApplicationsScopes with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PutApplicationsScopesOptions model
				putApplicationsScopesOptionsModel := new(appidmanagementv4.PutApplicationsScopesOptions)
				putApplicationsScopesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsScopesOptionsModel.Scopes = []string{"cartoons", "horror", "animated"}
				putApplicationsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.PutApplicationsScopes(putApplicationsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.PutApplicationsScopes(putApplicationsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutApplicationsScopes(putApplicationsScopesOptions *PutApplicationsScopesOptions)`, func() {
		tenantID := "testString"
		putApplicationsScopesPath := "/management/v4/testString/applications/testString/scopes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putApplicationsScopesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": ["Scopes"]}`)
				}))
			})
			It(`Invoke PutApplicationsScopes successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the PutApplicationsScopesOptions model
				putApplicationsScopesOptionsModel := new(appidmanagementv4.PutApplicationsScopesOptions)
				putApplicationsScopesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsScopesOptionsModel.Scopes = []string{"cartoons", "horror", "animated"}
				putApplicationsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.PutApplicationsScopesWithContext(ctx, putApplicationsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.PutApplicationsScopes(putApplicationsScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.PutApplicationsScopesWithContext(ctx, putApplicationsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putApplicationsScopesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": ["Scopes"]}`)
				}))
			})
			It(`Invoke PutApplicationsScopes successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.PutApplicationsScopes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PutApplicationsScopesOptions model
				putApplicationsScopesOptionsModel := new(appidmanagementv4.PutApplicationsScopesOptions)
				putApplicationsScopesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsScopesOptionsModel.Scopes = []string{"cartoons", "horror", "animated"}
				putApplicationsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.PutApplicationsScopes(putApplicationsScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PutApplicationsScopes with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PutApplicationsScopesOptions model
				putApplicationsScopesOptionsModel := new(appidmanagementv4.PutApplicationsScopesOptions)
				putApplicationsScopesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsScopesOptionsModel.Scopes = []string{"cartoons", "horror", "animated"}
				putApplicationsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.PutApplicationsScopes(putApplicationsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutApplicationsScopesOptions model with no property values
				putApplicationsScopesOptionsModelNew := new(appidmanagementv4.PutApplicationsScopesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.PutApplicationsScopes(putApplicationsScopesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplicationRoles(getApplicationRolesOptions *GetApplicationRolesOptions) - Operation response error`, func() {
		tenantID := "testString"
		getApplicationRolesPath := "/management/v4/testString/applications/testString/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationRolesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplicationRoles with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetApplicationRolesOptions model
				getApplicationRolesOptionsModel := new(appidmanagementv4.GetApplicationRolesOptions)
				getApplicationRolesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetApplicationRoles(getApplicationRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetApplicationRoles(getApplicationRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetApplicationRoles(getApplicationRolesOptions *GetApplicationRolesOptions)`, func() {
		tenantID := "testString"
		getApplicationRolesPath := "/management/v4/testString/applications/testString/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationRolesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "adult"}]}`)
				}))
			})
			It(`Invoke GetApplicationRoles successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationRolesOptions model
				getApplicationRolesOptionsModel := new(appidmanagementv4.GetApplicationRolesOptions)
				getApplicationRolesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetApplicationRolesWithContext(ctx, getApplicationRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetApplicationRoles(getApplicationRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetApplicationRolesWithContext(ctx, getApplicationRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationRolesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "adult"}]}`)
				}))
			})
			It(`Invoke GetApplicationRoles successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetApplicationRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationRolesOptions model
				getApplicationRolesOptionsModel := new(appidmanagementv4.GetApplicationRolesOptions)
				getApplicationRolesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetApplicationRoles(getApplicationRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplicationRoles with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetApplicationRolesOptions model
				getApplicationRolesOptionsModel := new(appidmanagementv4.GetApplicationRolesOptions)
				getApplicationRolesOptionsModel.ClientID = core.StringPtr("testString")
				getApplicationRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetApplicationRoles(getApplicationRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationRolesOptions model with no property values
				getApplicationRolesOptionsModelNew := new(appidmanagementv4.GetApplicationRolesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetApplicationRoles(getApplicationRolesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutApplicationsRoles(putApplicationsRolesOptions *PutApplicationsRolesOptions) - Operation response error`, func() {
		tenantID := "testString"
		putApplicationsRolesPath := "/management/v4/testString/applications/testString/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putApplicationsRolesPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutApplicationsRoles with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the PutApplicationsRolesOptions model
				putApplicationsRolesOptionsModel := new(appidmanagementv4.PutApplicationsRolesOptions)
				putApplicationsRolesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				putApplicationsRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.PutApplicationsRoles(putApplicationsRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.PutApplicationsRoles(putApplicationsRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutApplicationsRoles(putApplicationsRolesOptions *PutApplicationsRolesOptions)`, func() {
		tenantID := "testString"
		putApplicationsRolesPath := "/management/v4/testString/applications/testString/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putApplicationsRolesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "child"}]}`)
				}))
			})
			It(`Invoke PutApplicationsRoles successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the PutApplicationsRolesOptions model
				putApplicationsRolesOptionsModel := new(appidmanagementv4.PutApplicationsRolesOptions)
				putApplicationsRolesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				putApplicationsRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.PutApplicationsRolesWithContext(ctx, putApplicationsRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.PutApplicationsRoles(putApplicationsRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.PutApplicationsRolesWithContext(ctx, putApplicationsRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putApplicationsRolesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "child"}]}`)
				}))
			})
			It(`Invoke PutApplicationsRoles successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.PutApplicationsRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the PutApplicationsRolesOptions model
				putApplicationsRolesOptionsModel := new(appidmanagementv4.PutApplicationsRolesOptions)
				putApplicationsRolesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				putApplicationsRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.PutApplicationsRoles(putApplicationsRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PutApplicationsRoles with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the PutApplicationsRolesOptions model
				putApplicationsRolesOptionsModel := new(appidmanagementv4.PutApplicationsRolesOptions)
				putApplicationsRolesOptionsModel.ClientID = core.StringPtr("testString")
				putApplicationsRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				putApplicationsRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.PutApplicationsRoles(putApplicationsRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutApplicationsRolesOptions model with no property values
				putApplicationsRolesOptionsModelNew := new(appidmanagementv4.PutApplicationsRolesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.PutApplicationsRoles(putApplicationsRolesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListCloudDirectoryUsers(listCloudDirectoryUsersOptions *ListCloudDirectoryUsersOptions) - Operation response error`, func() {
		tenantID := "testString"
		listCloudDirectoryUsersPath := "/management/v4/testString/cloud_directory/Users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCloudDirectoryUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCloudDirectoryUsers with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListCloudDirectoryUsersOptions model
				listCloudDirectoryUsersOptionsModel := new(appidmanagementv4.ListCloudDirectoryUsersOptions)
				listCloudDirectoryUsersOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				listCloudDirectoryUsersOptionsModel.Count = core.Int64Ptr(int64(0))
				listCloudDirectoryUsersOptionsModel.Query = core.StringPtr("testString")
				listCloudDirectoryUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.ListCloudDirectoryUsers(listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.ListCloudDirectoryUsers(listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCloudDirectoryUsers(listCloudDirectoryUsersOptions *ListCloudDirectoryUsersOptions)`, func() {
		tenantID := "testString"
		listCloudDirectoryUsersPath := "/management/v4/testString/cloud_directory/Users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCloudDirectoryUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"totalResults": 12, "itemsPerPage": 12, "Resources": [{"anyKey": "anyValue"}]}`)
				}))
			})
			It(`Invoke ListCloudDirectoryUsers successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListCloudDirectoryUsersOptions model
				listCloudDirectoryUsersOptionsModel := new(appidmanagementv4.ListCloudDirectoryUsersOptions)
				listCloudDirectoryUsersOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				listCloudDirectoryUsersOptionsModel.Count = core.Int64Ptr(int64(0))
				listCloudDirectoryUsersOptionsModel.Query = core.StringPtr("testString")
				listCloudDirectoryUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.ListCloudDirectoryUsersWithContext(ctx, listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.ListCloudDirectoryUsers(listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.ListCloudDirectoryUsersWithContext(ctx, listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCloudDirectoryUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"totalResults": 12, "itemsPerPage": 12, "Resources": [{"anyKey": "anyValue"}]}`)
				}))
			})
			It(`Invoke ListCloudDirectoryUsers successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.ListCloudDirectoryUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCloudDirectoryUsersOptions model
				listCloudDirectoryUsersOptionsModel := new(appidmanagementv4.ListCloudDirectoryUsersOptions)
				listCloudDirectoryUsersOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				listCloudDirectoryUsersOptionsModel.Count = core.Int64Ptr(int64(0))
				listCloudDirectoryUsersOptionsModel.Query = core.StringPtr("testString")
				listCloudDirectoryUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.ListCloudDirectoryUsers(listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCloudDirectoryUsers with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListCloudDirectoryUsersOptions model
				listCloudDirectoryUsersOptionsModel := new(appidmanagementv4.ListCloudDirectoryUsersOptions)
				listCloudDirectoryUsersOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				listCloudDirectoryUsersOptionsModel.Count = core.Int64Ptr(int64(0))
				listCloudDirectoryUsersOptionsModel.Query = core.StringPtr("testString")
				listCloudDirectoryUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.ListCloudDirectoryUsers(listCloudDirectoryUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCloudDirectoryUser(createCloudDirectoryUserOptions *CreateCloudDirectoryUserOptions)`, func() {
		tenantID := "testString"
		createCloudDirectoryUserPath := "/management/v4/testString/cloud_directory/Users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCloudDirectoryUserPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCloudDirectoryUser successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.CreateCloudDirectoryUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)

				// Construct an instance of the CreateCloudDirectoryUserOptions model
				createCloudDirectoryUserOptionsModel := new(appidmanagementv4.CreateCloudDirectoryUserOptions)
				createCloudDirectoryUserOptionsModel.Emails = []appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}
				createCloudDirectoryUserOptionsModel.Password = core.StringPtr("userPassword")
				createCloudDirectoryUserOptionsModel.Active = core.BoolPtr(true)
				createCloudDirectoryUserOptionsModel.UserName = core.StringPtr("myUserName")
				createCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.CreateCloudDirectoryUser(createCloudDirectoryUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateCloudDirectoryUser with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)

				// Construct an instance of the CreateCloudDirectoryUserOptions model
				createCloudDirectoryUserOptionsModel := new(appidmanagementv4.CreateCloudDirectoryUserOptions)
				createCloudDirectoryUserOptionsModel.Emails = []appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}
				createCloudDirectoryUserOptionsModel.Password = core.StringPtr("userPassword")
				createCloudDirectoryUserOptionsModel.Active = core.BoolPtr(true)
				createCloudDirectoryUserOptionsModel.UserName = core.StringPtr("myUserName")
				createCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.CreateCloudDirectoryUser(createCloudDirectoryUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateCloudDirectoryUserOptions model with no property values
				createCloudDirectoryUserOptionsModelNew := new(appidmanagementv4.CreateCloudDirectoryUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.CreateCloudDirectoryUser(createCloudDirectoryUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectoryUser(getCloudDirectoryUserOptions *GetCloudDirectoryUserOptions)`, func() {
		tenantID := "testString"
		getCloudDirectoryUserPath := "/management/v4/testString/cloud_directory/Users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryUserPath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCloudDirectoryUser successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.GetCloudDirectoryUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetCloudDirectoryUserOptions model
				getCloudDirectoryUserOptionsModel := new(appidmanagementv4.GetCloudDirectoryUserOptions)
				getCloudDirectoryUserOptionsModel.UserID = core.StringPtr("testString")
				getCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.GetCloudDirectoryUser(getCloudDirectoryUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetCloudDirectoryUser with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryUserOptions model
				getCloudDirectoryUserOptionsModel := new(appidmanagementv4.GetCloudDirectoryUserOptions)
				getCloudDirectoryUserOptionsModel.UserID = core.StringPtr("testString")
				getCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.GetCloudDirectoryUser(getCloudDirectoryUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetCloudDirectoryUserOptions model with no property values
				getCloudDirectoryUserOptionsModelNew := new(appidmanagementv4.GetCloudDirectoryUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.GetCloudDirectoryUser(getCloudDirectoryUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCloudDirectoryUser(updateCloudDirectoryUserOptions *UpdateCloudDirectoryUserOptions)`, func() {
		tenantID := "testString"
		updateCloudDirectoryUserPath := "/management/v4/testString/cloud_directory/Users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCloudDirectoryUserPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCloudDirectoryUser successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UpdateCloudDirectoryUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)

				// Construct an instance of the UpdateCloudDirectoryUserOptions model
				updateCloudDirectoryUserOptionsModel := new(appidmanagementv4.UpdateCloudDirectoryUserOptions)
				updateCloudDirectoryUserOptionsModel.UserID = core.StringPtr("testString")
				updateCloudDirectoryUserOptionsModel.Emails = []appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}
				updateCloudDirectoryUserOptionsModel.Active = core.BoolPtr(true)
				updateCloudDirectoryUserOptionsModel.UserName = core.StringPtr("myUserName")
				updateCloudDirectoryUserOptionsModel.Password = core.StringPtr("userPassword")
				updateCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UpdateCloudDirectoryUser(updateCloudDirectoryUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateCloudDirectoryUser with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)

				// Construct an instance of the UpdateCloudDirectoryUserOptions model
				updateCloudDirectoryUserOptionsModel := new(appidmanagementv4.UpdateCloudDirectoryUserOptions)
				updateCloudDirectoryUserOptionsModel.UserID = core.StringPtr("testString")
				updateCloudDirectoryUserOptionsModel.Emails = []appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}
				updateCloudDirectoryUserOptionsModel.Active = core.BoolPtr(true)
				updateCloudDirectoryUserOptionsModel.UserName = core.StringPtr("myUserName")
				updateCloudDirectoryUserOptionsModel.Password = core.StringPtr("userPassword")
				updateCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UpdateCloudDirectoryUser(updateCloudDirectoryUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateCloudDirectoryUserOptions model with no property values
				updateCloudDirectoryUserOptionsModelNew := new(appidmanagementv4.UpdateCloudDirectoryUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UpdateCloudDirectoryUser(updateCloudDirectoryUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCloudDirectoryUser(deleteCloudDirectoryUserOptions *DeleteCloudDirectoryUserOptions)`, func() {
		tenantID := "testString"
		deleteCloudDirectoryUserPath := "/management/v4/testString/cloud_directory/Users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCloudDirectoryUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCloudDirectoryUser successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.DeleteCloudDirectoryUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCloudDirectoryUserOptions model
				deleteCloudDirectoryUserOptionsModel := new(appidmanagementv4.DeleteCloudDirectoryUserOptions)
				deleteCloudDirectoryUserOptionsModel.UserID = core.StringPtr("testString")
				deleteCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.DeleteCloudDirectoryUser(deleteCloudDirectoryUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCloudDirectoryUser with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteCloudDirectoryUserOptions model
				deleteCloudDirectoryUserOptionsModel := new(appidmanagementv4.DeleteCloudDirectoryUserOptions)
				deleteCloudDirectoryUserOptionsModel.UserID = core.StringPtr("testString")
				deleteCloudDirectoryUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.DeleteCloudDirectoryUser(deleteCloudDirectoryUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCloudDirectoryUserOptions model with no property values
				deleteCloudDirectoryUserOptionsModelNew := new(appidmanagementv4.DeleteCloudDirectoryUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.DeleteCloudDirectoryUser(deleteCloudDirectoryUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SsoLogoutFromAllApps(ssoLogoutFromAllAppsOptions *SsoLogoutFromAllAppsOptions)`, func() {
		tenantID := "testString"
		ssoLogoutFromAllAppsPath := "/management/v4/testString/cloud_directory/Users/testString/sso/logout"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(ssoLogoutFromAllAppsPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke SsoLogoutFromAllApps successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.SsoLogoutFromAllApps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SsoLogoutFromAllAppsOptions model
				ssoLogoutFromAllAppsOptionsModel := new(appidmanagementv4.SsoLogoutFromAllAppsOptions)
				ssoLogoutFromAllAppsOptionsModel.UserID = core.StringPtr("testString")
				ssoLogoutFromAllAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.SsoLogoutFromAllApps(ssoLogoutFromAllAppsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SsoLogoutFromAllApps with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SsoLogoutFromAllAppsOptions model
				ssoLogoutFromAllAppsOptionsModel := new(appidmanagementv4.SsoLogoutFromAllAppsOptions)
				ssoLogoutFromAllAppsOptionsModel.UserID = core.StringPtr("testString")
				ssoLogoutFromAllAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.SsoLogoutFromAllApps(ssoLogoutFromAllAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SsoLogoutFromAllAppsOptions model with no property values
				ssoLogoutFromAllAppsOptionsModelNew := new(appidmanagementv4.SsoLogoutFromAllAppsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.SsoLogoutFromAllApps(ssoLogoutFromAllAppsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CloudDirectoryExport(cloudDirectoryExportOptions *CloudDirectoryExportOptions) - Operation response error`, func() {
		tenantID := "testString"
		cloudDirectoryExportPath := "/management/v4/testString/cloud_directory/export"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryExportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["encryption_secret"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CloudDirectoryExport with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryExportOptions model
				cloudDirectoryExportOptionsModel := new(appidmanagementv4.CloudDirectoryExportOptions)
				cloudDirectoryExportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				cloudDirectoryExportOptionsModel.Count = core.Int64Ptr(int64(0))
				cloudDirectoryExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.CloudDirectoryExport(cloudDirectoryExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.CloudDirectoryExport(cloudDirectoryExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CloudDirectoryExport(cloudDirectoryExportOptions *CloudDirectoryExportOptions)`, func() {
		tenantID := "testString"
		cloudDirectoryExportPath := "/management/v4/testString/cloud_directory/export"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryExportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["encryption_secret"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"users": [{"scimUser": {"anyKey": "anyValue"}, "passwordHash": "PasswordHash", "passwordHashAlg": "PasswordHashAlg", "profile": {"attributes": {"anyKey": "anyValue"}}, "roles": ["Roles"]}]}`)
				}))
			})
			It(`Invoke CloudDirectoryExport successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the CloudDirectoryExportOptions model
				cloudDirectoryExportOptionsModel := new(appidmanagementv4.CloudDirectoryExportOptions)
				cloudDirectoryExportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				cloudDirectoryExportOptionsModel.Count = core.Int64Ptr(int64(0))
				cloudDirectoryExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.CloudDirectoryExportWithContext(ctx, cloudDirectoryExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.CloudDirectoryExport(cloudDirectoryExportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.CloudDirectoryExportWithContext(ctx, cloudDirectoryExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryExportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["encryption_secret"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"users": [{"scimUser": {"anyKey": "anyValue"}, "passwordHash": "PasswordHash", "passwordHashAlg": "PasswordHashAlg", "profile": {"attributes": {"anyKey": "anyValue"}}, "roles": ["Roles"]}]}`)
				}))
			})
			It(`Invoke CloudDirectoryExport successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.CloudDirectoryExport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CloudDirectoryExportOptions model
				cloudDirectoryExportOptionsModel := new(appidmanagementv4.CloudDirectoryExportOptions)
				cloudDirectoryExportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				cloudDirectoryExportOptionsModel.Count = core.Int64Ptr(int64(0))
				cloudDirectoryExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.CloudDirectoryExport(cloudDirectoryExportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CloudDirectoryExport with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryExportOptions model
				cloudDirectoryExportOptionsModel := new(appidmanagementv4.CloudDirectoryExportOptions)
				cloudDirectoryExportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				cloudDirectoryExportOptionsModel.Count = core.Int64Ptr(int64(0))
				cloudDirectoryExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.CloudDirectoryExport(cloudDirectoryExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CloudDirectoryExportOptions model with no property values
				cloudDirectoryExportOptionsModelNew := new(appidmanagementv4.CloudDirectoryExportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.CloudDirectoryExport(cloudDirectoryExportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CloudDirectoryImport(cloudDirectoryImportOptions *CloudDirectoryImportOptions) - Operation response error`, func() {
		tenantID := "testString"
		cloudDirectoryImportPath := "/management/v4/testString/cloud_directory/import"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryImportPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["encryption_secret"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CloudDirectoryImport with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ExportUserUsersItemProfile model
				exportUserUsersItemProfileModel := new(appidmanagementv4.ExportUserUsersItemProfile)
				exportUserUsersItemProfileModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ExportUserUsersItem model
				exportUserUsersItemModel := new(appidmanagementv4.ExportUserUsersItem)
				exportUserUsersItemModel.ScimUser = map[string]interface{}{"anyKey": "anyValue"}
				exportUserUsersItemModel.PasswordHash = core.StringPtr("testString")
				exportUserUsersItemModel.PasswordHashAlg = core.StringPtr("testString")
				exportUserUsersItemModel.Profile = exportUserUsersItemProfileModel
				exportUserUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the CloudDirectoryImportOptions model
				cloudDirectoryImportOptionsModel := new(appidmanagementv4.CloudDirectoryImportOptions)
				cloudDirectoryImportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryImportOptionsModel.Users = []appidmanagementv4.ExportUserUsersItem{*exportUserUsersItemModel}
				cloudDirectoryImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.CloudDirectoryImport(cloudDirectoryImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.CloudDirectoryImport(cloudDirectoryImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CloudDirectoryImport(cloudDirectoryImportOptions *CloudDirectoryImportOptions)`, func() {
		tenantID := "testString"
		cloudDirectoryImportPath := "/management/v4/testString/cloud_directory/import"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryImportPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["encryption_secret"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": 5, "failed": 6, "failReasons": [{"originalId": "OriginalID", "id": "ID", "email": "Email", "userName": "UserName", "error": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke CloudDirectoryImport successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ExportUserUsersItemProfile model
				exportUserUsersItemProfileModel := new(appidmanagementv4.ExportUserUsersItemProfile)
				exportUserUsersItemProfileModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ExportUserUsersItem model
				exportUserUsersItemModel := new(appidmanagementv4.ExportUserUsersItem)
				exportUserUsersItemModel.ScimUser = map[string]interface{}{"anyKey": "anyValue"}
				exportUserUsersItemModel.PasswordHash = core.StringPtr("testString")
				exportUserUsersItemModel.PasswordHashAlg = core.StringPtr("testString")
				exportUserUsersItemModel.Profile = exportUserUsersItemProfileModel
				exportUserUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the CloudDirectoryImportOptions model
				cloudDirectoryImportOptionsModel := new(appidmanagementv4.CloudDirectoryImportOptions)
				cloudDirectoryImportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryImportOptionsModel.Users = []appidmanagementv4.ExportUserUsersItem{*exportUserUsersItemModel}
				cloudDirectoryImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.CloudDirectoryImportWithContext(ctx, cloudDirectoryImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.CloudDirectoryImport(cloudDirectoryImportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.CloudDirectoryImportWithContext(ctx, cloudDirectoryImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryImportPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["encryption_secret"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": 5, "failed": 6, "failReasons": [{"originalId": "OriginalID", "id": "ID", "email": "Email", "userName": "UserName", "error": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke CloudDirectoryImport successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.CloudDirectoryImport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExportUserUsersItemProfile model
				exportUserUsersItemProfileModel := new(appidmanagementv4.ExportUserUsersItemProfile)
				exportUserUsersItemProfileModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ExportUserUsersItem model
				exportUserUsersItemModel := new(appidmanagementv4.ExportUserUsersItem)
				exportUserUsersItemModel.ScimUser = map[string]interface{}{"anyKey": "anyValue"}
				exportUserUsersItemModel.PasswordHash = core.StringPtr("testString")
				exportUserUsersItemModel.PasswordHashAlg = core.StringPtr("testString")
				exportUserUsersItemModel.Profile = exportUserUsersItemProfileModel
				exportUserUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the CloudDirectoryImportOptions model
				cloudDirectoryImportOptionsModel := new(appidmanagementv4.CloudDirectoryImportOptions)
				cloudDirectoryImportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryImportOptionsModel.Users = []appidmanagementv4.ExportUserUsersItem{*exportUserUsersItemModel}
				cloudDirectoryImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.CloudDirectoryImport(cloudDirectoryImportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CloudDirectoryImport with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ExportUserUsersItemProfile model
				exportUserUsersItemProfileModel := new(appidmanagementv4.ExportUserUsersItemProfile)
				exportUserUsersItemProfileModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ExportUserUsersItem model
				exportUserUsersItemModel := new(appidmanagementv4.ExportUserUsersItem)
				exportUserUsersItemModel.ScimUser = map[string]interface{}{"anyKey": "anyValue"}
				exportUserUsersItemModel.PasswordHash = core.StringPtr("testString")
				exportUserUsersItemModel.PasswordHashAlg = core.StringPtr("testString")
				exportUserUsersItemModel.Profile = exportUserUsersItemProfileModel
				exportUserUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the CloudDirectoryImportOptions model
				cloudDirectoryImportOptionsModel := new(appidmanagementv4.CloudDirectoryImportOptions)
				cloudDirectoryImportOptionsModel.EncryptionSecret = core.StringPtr("testString")
				cloudDirectoryImportOptionsModel.Users = []appidmanagementv4.ExportUserUsersItem{*exportUserUsersItemModel}
				cloudDirectoryImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.CloudDirectoryImport(cloudDirectoryImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CloudDirectoryImportOptions model with no property values
				cloudDirectoryImportOptionsModelNew := new(appidmanagementv4.CloudDirectoryImportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.CloudDirectoryImport(cloudDirectoryImportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptions *CloudDirectoryGetUserinfoOptions) - Operation response error`, func() {
		tenantID := "testString"
		cloudDirectoryGetUserinfoPath := "/management/v4/testString/cloud_directory/testString/userinfo"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryGetUserinfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CloudDirectoryGetUserinfo with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryGetUserinfoOptions model
				cloudDirectoryGetUserinfoOptionsModel := new(appidmanagementv4.CloudDirectoryGetUserinfoOptions)
				cloudDirectoryGetUserinfoOptionsModel.UserID = core.StringPtr("testString")
				cloudDirectoryGetUserinfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptions *CloudDirectoryGetUserinfoOptions)`, func() {
		tenantID := "testString"
		cloudDirectoryGetUserinfoPath := "/management/v4/testString/cloud_directory/testString/userinfo"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryGetUserinfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sub": "Sub", "identities": [{"provider": "Provider", "id": "ID", "idpUserInfo": {"anyKey": "anyValue"}}], "attributes": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CloudDirectoryGetUserinfo successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the CloudDirectoryGetUserinfoOptions model
				cloudDirectoryGetUserinfoOptionsModel := new(appidmanagementv4.CloudDirectoryGetUserinfoOptions)
				cloudDirectoryGetUserinfoOptionsModel.UserID = core.StringPtr("testString")
				cloudDirectoryGetUserinfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.CloudDirectoryGetUserinfoWithContext(ctx, cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.CloudDirectoryGetUserinfoWithContext(ctx, cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryGetUserinfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sub": "Sub", "identities": [{"provider": "Provider", "id": "ID", "idpUserInfo": {"anyKey": "anyValue"}}], "attributes": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CloudDirectoryGetUserinfo successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.CloudDirectoryGetUserinfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CloudDirectoryGetUserinfoOptions model
				cloudDirectoryGetUserinfoOptionsModel := new(appidmanagementv4.CloudDirectoryGetUserinfoOptions)
				cloudDirectoryGetUserinfoOptionsModel.UserID = core.StringPtr("testString")
				cloudDirectoryGetUserinfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CloudDirectoryGetUserinfo with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryGetUserinfoOptions model
				cloudDirectoryGetUserinfoOptionsModel := new(appidmanagementv4.CloudDirectoryGetUserinfoOptions)
				cloudDirectoryGetUserinfoOptionsModel.UserID = core.StringPtr("testString")
				cloudDirectoryGetUserinfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CloudDirectoryGetUserinfoOptions model with no property values
				cloudDirectoryGetUserinfoOptionsModelNew := new(appidmanagementv4.CloudDirectoryGetUserinfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`StartSignUp(startSignUpOptions *StartSignUpOptions)`, func() {
		tenantID := "testString"
		startSignUpPath := "/management/v4/testString/cloud_directory/sign_up"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startSignUpPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for shouldCreateProfile query parameter
					Expect(req.URL.Query()["language"]).To(Equal([]string{"testString"}))
					res.WriteHeader(201)
				}))
			})
			It(`Invoke StartSignUp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.StartSignUp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)

				// Construct an instance of the StartSignUpOptions model
				startSignUpOptionsModel := new(appidmanagementv4.StartSignUpOptions)
				startSignUpOptionsModel.ShouldCreateProfile = core.BoolPtr(true)
				startSignUpOptionsModel.Emails = []appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}
				startSignUpOptionsModel.Password = core.StringPtr("userPassword")
				startSignUpOptionsModel.Active = core.BoolPtr(true)
				startSignUpOptionsModel.UserName = core.StringPtr("myUserName")
				startSignUpOptionsModel.Language = core.StringPtr("testString")
				startSignUpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.StartSignUp(startSignUpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke StartSignUp with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)

				// Construct an instance of the StartSignUpOptions model
				startSignUpOptionsModel := new(appidmanagementv4.StartSignUpOptions)
				startSignUpOptionsModel.ShouldCreateProfile = core.BoolPtr(true)
				startSignUpOptionsModel.Emails = []appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}
				startSignUpOptionsModel.Password = core.StringPtr("userPassword")
				startSignUpOptionsModel.Active = core.BoolPtr(true)
				startSignUpOptionsModel.UserName = core.StringPtr("myUserName")
				startSignUpOptionsModel.Language = core.StringPtr("testString")
				startSignUpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.StartSignUp(startSignUpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the StartSignUpOptions model with no property values
				startSignUpOptionsModelNew := new(appidmanagementv4.StartSignUpOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.StartSignUp(startSignUpOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UserVerificationResult(userVerificationResultOptions *UserVerificationResultOptions) - Operation response error`, func() {
		tenantID := "testString"
		userVerificationResultPath := "/management/v4/testString/cloud_directory/sign_up/confirmation_result"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userVerificationResultPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UserVerificationResult with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UserVerificationResultOptions model
				userVerificationResultOptionsModel := new(appidmanagementv4.UserVerificationResultOptions)
				userVerificationResultOptionsModel.Context = core.StringPtr("testString")
				userVerificationResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UserVerificationResult(userVerificationResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UserVerificationResult(userVerificationResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UserVerificationResult(userVerificationResultOptions *UserVerificationResultOptions)`, func() {
		tenantID := "testString"
		userVerificationResultPath := "/management/v4/testString/cloud_directory/sign_up/confirmation_result"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userVerificationResultPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false, "uuid": "UUID"}`)
				}))
			})
			It(`Invoke UserVerificationResult successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UserVerificationResultOptions model
				userVerificationResultOptionsModel := new(appidmanagementv4.UserVerificationResultOptions)
				userVerificationResultOptionsModel.Context = core.StringPtr("testString")
				userVerificationResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UserVerificationResultWithContext(ctx, userVerificationResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UserVerificationResult(userVerificationResultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UserVerificationResultWithContext(ctx, userVerificationResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userVerificationResultPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false, "uuid": "UUID"}`)
				}))
			})
			It(`Invoke UserVerificationResult successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UserVerificationResult(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserVerificationResultOptions model
				userVerificationResultOptionsModel := new(appidmanagementv4.UserVerificationResultOptions)
				userVerificationResultOptionsModel.Context = core.StringPtr("testString")
				userVerificationResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UserVerificationResult(userVerificationResultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UserVerificationResult with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UserVerificationResultOptions model
				userVerificationResultOptionsModel := new(appidmanagementv4.UserVerificationResultOptions)
				userVerificationResultOptionsModel.Context = core.StringPtr("testString")
				userVerificationResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UserVerificationResult(userVerificationResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UserVerificationResultOptions model with no property values
				userVerificationResultOptionsModelNew := new(appidmanagementv4.UserVerificationResultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UserVerificationResult(userVerificationResultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`StartForgotPassword(startForgotPasswordOptions *StartForgotPasswordOptions)`, func() {
		tenantID := "testString"
		startForgotPasswordPath := "/management/v4/testString/cloud_directory/forgot_password"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startForgotPasswordPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["language"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke StartForgotPassword successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.StartForgotPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the StartForgotPasswordOptions model
				startForgotPasswordOptionsModel := new(appidmanagementv4.StartForgotPasswordOptions)
				startForgotPasswordOptionsModel.User = core.StringPtr("testString")
				startForgotPasswordOptionsModel.Language = core.StringPtr("testString")
				startForgotPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.StartForgotPassword(startForgotPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke StartForgotPassword with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the StartForgotPasswordOptions model
				startForgotPasswordOptionsModel := new(appidmanagementv4.StartForgotPasswordOptions)
				startForgotPasswordOptionsModel.User = core.StringPtr("testString")
				startForgotPasswordOptionsModel.Language = core.StringPtr("testString")
				startForgotPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.StartForgotPassword(startForgotPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the StartForgotPasswordOptions model with no property values
				startForgotPasswordOptionsModelNew := new(appidmanagementv4.StartForgotPasswordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.StartForgotPassword(startForgotPasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ForgotPasswordResult(forgotPasswordResultOptions *ForgotPasswordResultOptions) - Operation response error`, func() {
		tenantID := "testString"
		forgotPasswordResultPath := "/management/v4/testString/cloud_directory/forgot_password/confirmation_result"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forgotPasswordResultPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ForgotPasswordResult with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ForgotPasswordResultOptions model
				forgotPasswordResultOptionsModel := new(appidmanagementv4.ForgotPasswordResultOptions)
				forgotPasswordResultOptionsModel.Context = core.StringPtr("testString")
				forgotPasswordResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.ForgotPasswordResult(forgotPasswordResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.ForgotPasswordResult(forgotPasswordResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ForgotPasswordResult(forgotPasswordResultOptions *ForgotPasswordResultOptions)`, func() {
		tenantID := "testString"
		forgotPasswordResultPath := "/management/v4/testString/cloud_directory/forgot_password/confirmation_result"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forgotPasswordResultPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false, "uuid": "UUID"}`)
				}))
			})
			It(`Invoke ForgotPasswordResult successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ForgotPasswordResultOptions model
				forgotPasswordResultOptionsModel := new(appidmanagementv4.ForgotPasswordResultOptions)
				forgotPasswordResultOptionsModel.Context = core.StringPtr("testString")
				forgotPasswordResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.ForgotPasswordResultWithContext(ctx, forgotPasswordResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.ForgotPasswordResult(forgotPasswordResultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.ForgotPasswordResultWithContext(ctx, forgotPasswordResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forgotPasswordResultPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false, "uuid": "UUID"}`)
				}))
			})
			It(`Invoke ForgotPasswordResult successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.ForgotPasswordResult(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ForgotPasswordResultOptions model
				forgotPasswordResultOptionsModel := new(appidmanagementv4.ForgotPasswordResultOptions)
				forgotPasswordResultOptionsModel.Context = core.StringPtr("testString")
				forgotPasswordResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.ForgotPasswordResult(forgotPasswordResultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ForgotPasswordResult with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ForgotPasswordResultOptions model
				forgotPasswordResultOptionsModel := new(appidmanagementv4.ForgotPasswordResultOptions)
				forgotPasswordResultOptionsModel.Context = core.StringPtr("testString")
				forgotPasswordResultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.ForgotPasswordResult(forgotPasswordResultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ForgotPasswordResultOptions model with no property values
				forgotPasswordResultOptionsModelNew := new(appidmanagementv4.ForgotPasswordResultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.ForgotPasswordResult(forgotPasswordResultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangePassword(changePasswordOptions *ChangePasswordOptions)`, func() {
		tenantID := "testString"
		changePasswordPath := "/management/v4/testString/cloud_directory/change_password"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(changePasswordPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["language"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ChangePassword successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.ChangePassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ChangePasswordOptions model
				changePasswordOptionsModel := new(appidmanagementv4.ChangePasswordOptions)
				changePasswordOptionsModel.NewPassword = core.StringPtr("testString")
				changePasswordOptionsModel.UUID = core.StringPtr("testString")
				changePasswordOptionsModel.ChangedIpAddress = core.StringPtr("testString")
				changePasswordOptionsModel.Language = core.StringPtr("testString")
				changePasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.ChangePassword(changePasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ChangePassword with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ChangePasswordOptions model
				changePasswordOptionsModel := new(appidmanagementv4.ChangePasswordOptions)
				changePasswordOptionsModel.NewPassword = core.StringPtr("testString")
				changePasswordOptionsModel.UUID = core.StringPtr("testString")
				changePasswordOptionsModel.ChangedIpAddress = core.StringPtr("testString")
				changePasswordOptionsModel.Language = core.StringPtr("testString")
				changePasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.ChangePassword(changePasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ChangePasswordOptions model with no property values
				changePasswordOptionsModelNew := new(appidmanagementv4.ChangePasswordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.ChangePassword(changePasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResendNotification(resendNotificationOptions *ResendNotificationOptions) - Operation response error`, func() {
		tenantID := "testString"
		resendNotificationPath := "/management/v4/testString/cloud_directory/resend/USER_VERIFICATION"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resendNotificationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["language"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ResendNotification with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ResendNotificationOptions model
				resendNotificationOptionsModel := new(appidmanagementv4.ResendNotificationOptions)
				resendNotificationOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				resendNotificationOptionsModel.UUID = core.StringPtr("testString")
				resendNotificationOptionsModel.Language = core.StringPtr("testString")
				resendNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.ResendNotification(resendNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.ResendNotification(resendNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ResendNotification(resendNotificationOptions *ResendNotificationOptions)`, func() {
		tenantID := "testString"
		resendNotificationPath := "/management/v4/testString/cloud_directory/resend/USER_VERIFICATION"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resendNotificationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["language"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				}))
			})
			It(`Invoke ResendNotification successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ResendNotificationOptions model
				resendNotificationOptionsModel := new(appidmanagementv4.ResendNotificationOptions)
				resendNotificationOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				resendNotificationOptionsModel.UUID = core.StringPtr("testString")
				resendNotificationOptionsModel.Language = core.StringPtr("testString")
				resendNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.ResendNotificationWithContext(ctx, resendNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.ResendNotification(resendNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.ResendNotificationWithContext(ctx, resendNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resendNotificationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["language"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				}))
			})
			It(`Invoke ResendNotification successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.ResendNotification(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResendNotificationOptions model
				resendNotificationOptionsModel := new(appidmanagementv4.ResendNotificationOptions)
				resendNotificationOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				resendNotificationOptionsModel.UUID = core.StringPtr("testString")
				resendNotificationOptionsModel.Language = core.StringPtr("testString")
				resendNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.ResendNotification(resendNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ResendNotification with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ResendNotificationOptions model
				resendNotificationOptionsModel := new(appidmanagementv4.ResendNotificationOptions)
				resendNotificationOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				resendNotificationOptionsModel.UUID = core.StringPtr("testString")
				resendNotificationOptionsModel.Language = core.StringPtr("testString")
				resendNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.ResendNotification(resendNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ResendNotificationOptions model with no property values
				resendNotificationOptionsModelNew := new(appidmanagementv4.ResendNotificationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.ResendNotification(resendNotificationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CloudDirectoryRemove(cloudDirectoryRemoveOptions *CloudDirectoryRemoveOptions)`, func() {
		tenantID := "testString"
		cloudDirectoryRemovePath := "/management/v4/testString/cloud_directory/remove/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cloudDirectoryRemovePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CloudDirectoryRemove successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.CloudDirectoryRemove(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CloudDirectoryRemoveOptions model
				cloudDirectoryRemoveOptionsModel := new(appidmanagementv4.CloudDirectoryRemoveOptions)
				cloudDirectoryRemoveOptionsModel.UserID = core.StringPtr("testString")
				cloudDirectoryRemoveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.CloudDirectoryRemove(cloudDirectoryRemoveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CloudDirectoryRemove with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryRemoveOptions model
				cloudDirectoryRemoveOptionsModel := new(appidmanagementv4.CloudDirectoryRemoveOptions)
				cloudDirectoryRemoveOptionsModel.UserID = core.StringPtr("testString")
				cloudDirectoryRemoveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.CloudDirectoryRemove(cloudDirectoryRemoveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CloudDirectoryRemoveOptions model with no property values
				cloudDirectoryRemoveOptionsModelNew := new(appidmanagementv4.CloudDirectoryRemoveOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.CloudDirectoryRemove(cloudDirectoryRemoveOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetTokensConfig(getTokensConfigOptions *GetTokensConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		getTokensConfigPath := "/management/v4/testString/config/tokens"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTokensConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTokensConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetTokensConfigOptions model
				getTokensConfigOptionsModel := new(appidmanagementv4.GetTokensConfigOptions)
				getTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetTokensConfig(getTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetTokensConfig(getTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTokensConfig(getTokensConfigOptions *GetTokensConfigOptions)`, func() {
		tenantID := "testString"
		getTokensConfigPath := "/management/v4/testString/config/tokens"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTokensConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"idTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "accessTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "access": [{"expires_in": 9}], "refresh": [{"expires_in": 9, "enabled": false}], "anonymousAccess": [{"expires_in": 9}]}`)
				}))
			})
			It(`Invoke GetTokensConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetTokensConfigOptions model
				getTokensConfigOptionsModel := new(appidmanagementv4.GetTokensConfigOptions)
				getTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetTokensConfigWithContext(ctx, getTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetTokensConfig(getTokensConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetTokensConfigWithContext(ctx, getTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTokensConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"idTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "accessTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "access": [{"expires_in": 9}], "refresh": [{"expires_in": 9, "enabled": false}], "anonymousAccess": [{"expires_in": 9}]}`)
				}))
			})
			It(`Invoke GetTokensConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetTokensConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTokensConfigOptions model
				getTokensConfigOptionsModel := new(appidmanagementv4.GetTokensConfigOptions)
				getTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetTokensConfig(getTokensConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTokensConfig with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetTokensConfigOptions model
				getTokensConfigOptionsModel := new(appidmanagementv4.GetTokensConfigOptions)
				getTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetTokensConfig(getTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutTokensConfig(putTokensConfigOptions *PutTokensConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		putTokensConfigPath := "/management/v4/testString/config/tokens"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putTokensConfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutTokensConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the TokenClaimMapping model
				tokenClaimMappingModel := new(appidmanagementv4.TokenClaimMapping)
				tokenClaimMappingModel.Source = core.StringPtr("saml")
				tokenClaimMappingModel.SourceClaim = core.StringPtr("testString")
				tokenClaimMappingModel.DestinationClaim = core.StringPtr("testString")

				// Construct an instance of the TokenConfigParams model
				tokenConfigParamsModel := new(appidmanagementv4.TokenConfigParams)
				tokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RefreshTokenConfigParams model
				refreshTokenConfigParamsModel := new(appidmanagementv4.RefreshTokenConfigParams)
				refreshTokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))
				refreshTokenConfigParamsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the PutTokensConfigOptions model
				putTokensConfigOptionsModel := new(appidmanagementv4.PutTokensConfigOptions)
				putTokensConfigOptionsModel.IdTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.AccessTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.Access = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Refresh = []appidmanagementv4.RefreshTokenConfigParams{*refreshTokenConfigParamsModel}
				putTokensConfigOptionsModel.AnonymousAccess = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.PutTokensConfig(putTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.PutTokensConfig(putTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutTokensConfig(putTokensConfigOptions *PutTokensConfigOptions)`, func() {
		tenantID := "testString"
		putTokensConfigPath := "/management/v4/testString/config/tokens"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putTokensConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"idTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "accessTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "access": [{"expires_in": 9}], "refresh": [{"expires_in": 9, "enabled": false}], "anonymousAccess": [{"expires_in": 9}]}`)
				}))
			})
			It(`Invoke PutTokensConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the TokenClaimMapping model
				tokenClaimMappingModel := new(appidmanagementv4.TokenClaimMapping)
				tokenClaimMappingModel.Source = core.StringPtr("saml")
				tokenClaimMappingModel.SourceClaim = core.StringPtr("testString")
				tokenClaimMappingModel.DestinationClaim = core.StringPtr("testString")

				// Construct an instance of the TokenConfigParams model
				tokenConfigParamsModel := new(appidmanagementv4.TokenConfigParams)
				tokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RefreshTokenConfigParams model
				refreshTokenConfigParamsModel := new(appidmanagementv4.RefreshTokenConfigParams)
				refreshTokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))
				refreshTokenConfigParamsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the PutTokensConfigOptions model
				putTokensConfigOptionsModel := new(appidmanagementv4.PutTokensConfigOptions)
				putTokensConfigOptionsModel.IdTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.AccessTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.Access = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Refresh = []appidmanagementv4.RefreshTokenConfigParams{*refreshTokenConfigParamsModel}
				putTokensConfigOptionsModel.AnonymousAccess = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.PutTokensConfigWithContext(ctx, putTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.PutTokensConfig(putTokensConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.PutTokensConfigWithContext(ctx, putTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putTokensConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"idTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "accessTokenClaims": [{"source": "saml", "sourceClaim": "SourceClaim", "destinationClaim": "DestinationClaim"}], "access": [{"expires_in": 9}], "refresh": [{"expires_in": 9, "enabled": false}], "anonymousAccess": [{"expires_in": 9}]}`)
				}))
			})
			It(`Invoke PutTokensConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.PutTokensConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TokenClaimMapping model
				tokenClaimMappingModel := new(appidmanagementv4.TokenClaimMapping)
				tokenClaimMappingModel.Source = core.StringPtr("saml")
				tokenClaimMappingModel.SourceClaim = core.StringPtr("testString")
				tokenClaimMappingModel.DestinationClaim = core.StringPtr("testString")

				// Construct an instance of the TokenConfigParams model
				tokenConfigParamsModel := new(appidmanagementv4.TokenConfigParams)
				tokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RefreshTokenConfigParams model
				refreshTokenConfigParamsModel := new(appidmanagementv4.RefreshTokenConfigParams)
				refreshTokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))
				refreshTokenConfigParamsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the PutTokensConfigOptions model
				putTokensConfigOptionsModel := new(appidmanagementv4.PutTokensConfigOptions)
				putTokensConfigOptionsModel.IdTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.AccessTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.Access = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Refresh = []appidmanagementv4.RefreshTokenConfigParams{*refreshTokenConfigParamsModel}
				putTokensConfigOptionsModel.AnonymousAccess = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.PutTokensConfig(putTokensConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PutTokensConfig with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the TokenClaimMapping model
				tokenClaimMappingModel := new(appidmanagementv4.TokenClaimMapping)
				tokenClaimMappingModel.Source = core.StringPtr("saml")
				tokenClaimMappingModel.SourceClaim = core.StringPtr("testString")
				tokenClaimMappingModel.DestinationClaim = core.StringPtr("testString")

				// Construct an instance of the TokenConfigParams model
				tokenConfigParamsModel := new(appidmanagementv4.TokenConfigParams)
				tokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RefreshTokenConfigParams model
				refreshTokenConfigParamsModel := new(appidmanagementv4.RefreshTokenConfigParams)
				refreshTokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))
				refreshTokenConfigParamsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the PutTokensConfigOptions model
				putTokensConfigOptionsModel := new(appidmanagementv4.PutTokensConfigOptions)
				putTokensConfigOptionsModel.IdTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.AccessTokenClaims = []appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}
				putTokensConfigOptionsModel.Access = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Refresh = []appidmanagementv4.RefreshTokenConfigParams{*refreshTokenConfigParamsModel}
				putTokensConfigOptionsModel.AnonymousAccess = []appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}
				putTokensConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.PutTokensConfig(putTokensConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRedirectUris(getRedirectUrisOptions *GetRedirectUrisOptions) - Operation response error`, func() {
		tenantID := "testString"
		getRedirectUrisPath := "/management/v4/testString/config/redirect_uris"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRedirectUrisPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRedirectUris with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetRedirectUrisOptions model
				getRedirectUrisOptionsModel := new(appidmanagementv4.GetRedirectUrisOptions)
				getRedirectUrisOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetRedirectUris(getRedirectUrisOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetRedirectUris(getRedirectUrisOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRedirectUris(getRedirectUrisOptions *GetRedirectUrisOptions)`, func() {
		tenantID := "testString"
		getRedirectUrisPath := "/management/v4/testString/config/redirect_uris"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRedirectUrisPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"redirectUris": ["RedirectUris"]}`)
				}))
			})
			It(`Invoke GetRedirectUris successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetRedirectUrisOptions model
				getRedirectUrisOptionsModel := new(appidmanagementv4.GetRedirectUrisOptions)
				getRedirectUrisOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetRedirectUrisWithContext(ctx, getRedirectUrisOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetRedirectUris(getRedirectUrisOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetRedirectUrisWithContext(ctx, getRedirectUrisOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRedirectUrisPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"redirectUris": ["RedirectUris"]}`)
				}))
			})
			It(`Invoke GetRedirectUris successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetRedirectUris(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRedirectUrisOptions model
				getRedirectUrisOptionsModel := new(appidmanagementv4.GetRedirectUrisOptions)
				getRedirectUrisOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetRedirectUris(getRedirectUrisOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRedirectUris with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetRedirectUrisOptions model
				getRedirectUrisOptionsModel := new(appidmanagementv4.GetRedirectUrisOptions)
				getRedirectUrisOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetRedirectUris(getRedirectUrisOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateRedirectUris(updateRedirectUrisOptions *UpdateRedirectUrisOptions)`, func() {
		tenantID := "testString"
		updateRedirectUrisPath := "/management/v4/testString/config/redirect_uris"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRedirectUrisPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateRedirectUris successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UpdateRedirectUris(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RedirectUriConfig model
				redirectUriConfigModel := new(appidmanagementv4.RedirectUriConfig)
				redirectUriConfigModel.RedirectUris = []string{"http://localhost:3000/oauth-callback"}
				redirectUriConfigModel.TrustCloudIAMRedirectUris = core.BoolPtr(true)
				redirectUriConfigModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the UpdateRedirectUrisOptions model
				updateRedirectUrisOptionsModel := new(appidmanagementv4.UpdateRedirectUrisOptions)
				updateRedirectUrisOptionsModel.RedirectUrisArray = redirectUriConfigModel
				updateRedirectUrisOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UpdateRedirectUris(updateRedirectUrisOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateRedirectUris with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the RedirectUriConfig model
				redirectUriConfigModel := new(appidmanagementv4.RedirectUriConfig)
				redirectUriConfigModel.RedirectUris = []string{"http://localhost:3000/oauth-callback"}
				redirectUriConfigModel.TrustCloudIAMRedirectUris = core.BoolPtr(true)
				redirectUriConfigModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the UpdateRedirectUrisOptions model
				updateRedirectUrisOptionsModel := new(appidmanagementv4.UpdateRedirectUrisOptions)
				updateRedirectUrisOptionsModel.RedirectUrisArray = redirectUriConfigModel
				updateRedirectUrisOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UpdateRedirectUris(updateRedirectUrisOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateRedirectUrisOptions model with no property values
				updateRedirectUrisOptionsModelNew := new(appidmanagementv4.UpdateRedirectUrisOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UpdateRedirectUris(updateRedirectUrisOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserProfilesConfig(getUserProfilesConfigOptions *GetUserProfilesConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		getUserProfilesConfigPath := "/management/v4/testString/config/users_profile"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserProfilesConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserProfilesConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfilesConfigOptions model
				getUserProfilesConfigOptionsModel := new(appidmanagementv4.GetUserProfilesConfigOptions)
				getUserProfilesConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetUserProfilesConfig(getUserProfilesConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetUserProfilesConfig(getUserProfilesConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUserProfilesConfig(getUserProfilesConfigOptions *GetUserProfilesConfigOptions)`, func() {
		tenantID := "testString"
		getUserProfilesConfigPath := "/management/v4/testString/config/users_profile"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserProfilesConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true}`)
				}))
			})
			It(`Invoke GetUserProfilesConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetUserProfilesConfigOptions model
				getUserProfilesConfigOptionsModel := new(appidmanagementv4.GetUserProfilesConfigOptions)
				getUserProfilesConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetUserProfilesConfigWithContext(ctx, getUserProfilesConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetUserProfilesConfig(getUserProfilesConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetUserProfilesConfigWithContext(ctx, getUserProfilesConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserProfilesConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true}`)
				}))
			})
			It(`Invoke GetUserProfilesConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetUserProfilesConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserProfilesConfigOptions model
				getUserProfilesConfigOptionsModel := new(appidmanagementv4.GetUserProfilesConfigOptions)
				getUserProfilesConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetUserProfilesConfig(getUserProfilesConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUserProfilesConfig with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfilesConfigOptions model
				getUserProfilesConfigOptionsModel := new(appidmanagementv4.GetUserProfilesConfigOptions)
				getUserProfilesConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetUserProfilesConfig(getUserProfilesConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateUserProfilesConfig(updateUserProfilesConfigOptions *UpdateUserProfilesConfigOptions)`, func() {
		tenantID := "testString"
		updateUserProfilesConfigPath := "/management/v4/testString/config/users_profile"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserProfilesConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateUserProfilesConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UpdateUserProfilesConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateUserProfilesConfigOptions model
				updateUserProfilesConfigOptionsModel := new(appidmanagementv4.UpdateUserProfilesConfigOptions)
				updateUserProfilesConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateUserProfilesConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UpdateUserProfilesConfig(updateUserProfilesConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateUserProfilesConfig with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserProfilesConfigOptions model
				updateUserProfilesConfigOptionsModel := new(appidmanagementv4.UpdateUserProfilesConfigOptions)
				updateUserProfilesConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateUserProfilesConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UpdateUserProfilesConfig(updateUserProfilesConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateUserProfilesConfigOptions model with no property values
				updateUserProfilesConfigOptionsModelNew := new(appidmanagementv4.UpdateUserProfilesConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UpdateUserProfilesConfig(updateUserProfilesConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetThemeText(getThemeTextOptions *GetThemeTextOptions) - Operation response error`, func() {
		tenantID := "testString"
		getThemeTextPath := "/management/v4/testString/config/ui/theme_text"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getThemeTextPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetThemeText with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetThemeTextOptions model
				getThemeTextOptionsModel := new(appidmanagementv4.GetThemeTextOptions)
				getThemeTextOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetThemeText(getThemeTextOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetThemeText(getThemeTextOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetThemeText(getThemeTextOptions *GetThemeTextOptions)`, func() {
		tenantID := "testString"
		getThemeTextPath := "/management/v4/testString/config/ui/theme_text"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getThemeTextPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"footnote": "Footnote", "tabTitle": "TabTitle"}`)
				}))
			})
			It(`Invoke GetThemeText successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetThemeTextOptions model
				getThemeTextOptionsModel := new(appidmanagementv4.GetThemeTextOptions)
				getThemeTextOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetThemeTextWithContext(ctx, getThemeTextOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetThemeText(getThemeTextOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetThemeTextWithContext(ctx, getThemeTextOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getThemeTextPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"footnote": "Footnote", "tabTitle": "TabTitle"}`)
				}))
			})
			It(`Invoke GetThemeText successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetThemeText(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetThemeTextOptions model
				getThemeTextOptionsModel := new(appidmanagementv4.GetThemeTextOptions)
				getThemeTextOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetThemeText(getThemeTextOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetThemeText with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetThemeTextOptions model
				getThemeTextOptionsModel := new(appidmanagementv4.GetThemeTextOptions)
				getThemeTextOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetThemeText(getThemeTextOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostThemeText(postThemeTextOptions *PostThemeTextOptions)`, func() {
		tenantID := "testString"
		postThemeTextPath := "/management/v4/testString/config/ui/theme_text"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postThemeTextPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke PostThemeText successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.PostThemeText(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PostThemeTextOptions model
				postThemeTextOptionsModel := new(appidmanagementv4.PostThemeTextOptions)
				postThemeTextOptionsModel.TabTitle = core.StringPtr("Login")
				postThemeTextOptionsModel.Footnote = core.StringPtr("Powered by App ID")
				postThemeTextOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.PostThemeText(postThemeTextOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PostThemeText with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostThemeTextOptions model
				postThemeTextOptionsModel := new(appidmanagementv4.PostThemeTextOptions)
				postThemeTextOptionsModel.TabTitle = core.StringPtr("Login")
				postThemeTextOptionsModel.Footnote = core.StringPtr("Powered by App ID")
				postThemeTextOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.PostThemeText(postThemeTextOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetThemeColor(getThemeColorOptions *GetThemeColorOptions) - Operation response error`, func() {
		tenantID := "testString"
		getThemeColorPath := "/management/v4/testString/config/ui/theme_color"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getThemeColorPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetThemeColor with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetThemeColorOptions model
				getThemeColorOptionsModel := new(appidmanagementv4.GetThemeColorOptions)
				getThemeColorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetThemeColor(getThemeColorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetThemeColor(getThemeColorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetThemeColor(getThemeColorOptions *GetThemeColorOptions)`, func() {
		tenantID := "testString"
		getThemeColorPath := "/management/v4/testString/config/ui/theme_color"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getThemeColorPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"headerColor": "HeaderColor"}`)
				}))
			})
			It(`Invoke GetThemeColor successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetThemeColorOptions model
				getThemeColorOptionsModel := new(appidmanagementv4.GetThemeColorOptions)
				getThemeColorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetThemeColorWithContext(ctx, getThemeColorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetThemeColor(getThemeColorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetThemeColorWithContext(ctx, getThemeColorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getThemeColorPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"headerColor": "HeaderColor"}`)
				}))
			})
			It(`Invoke GetThemeColor successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetThemeColor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetThemeColorOptions model
				getThemeColorOptionsModel := new(appidmanagementv4.GetThemeColorOptions)
				getThemeColorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetThemeColor(getThemeColorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetThemeColor with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetThemeColorOptions model
				getThemeColorOptionsModel := new(appidmanagementv4.GetThemeColorOptions)
				getThemeColorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetThemeColor(getThemeColorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostThemeColor(postThemeColorOptions *PostThemeColorOptions)`, func() {
		tenantID := "testString"
		postThemeColorPath := "/management/v4/testString/config/ui/theme_color"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postThemeColorPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke PostThemeColor successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.PostThemeColor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PostThemeColorOptions model
				postThemeColorOptionsModel := new(appidmanagementv4.PostThemeColorOptions)
				postThemeColorOptionsModel.HeaderColor = core.StringPtr("#EEF2F5")
				postThemeColorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.PostThemeColor(postThemeColorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PostThemeColor with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostThemeColorOptions model
				postThemeColorOptionsModel := new(appidmanagementv4.PostThemeColorOptions)
				postThemeColorOptionsModel.HeaderColor = core.StringPtr("#EEF2F5")
				postThemeColorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.PostThemeColor(postThemeColorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMedia(getMediaOptions *GetMediaOptions) - Operation response error`, func() {
		tenantID := "testString"
		getMediaPath := "/management/v4/testString/config/ui/media"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMediaPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMedia with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetMediaOptions model
				getMediaOptionsModel := new(appidmanagementv4.GetMediaOptions)
				getMediaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetMedia(getMediaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetMedia(getMediaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMedia(getMediaOptions *GetMediaOptions)`, func() {
		tenantID := "testString"
		getMediaPath := "/management/v4/testString/config/ui/media"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMediaPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"image": "Image"}`)
				}))
			})
			It(`Invoke GetMedia successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetMediaOptions model
				getMediaOptionsModel := new(appidmanagementv4.GetMediaOptions)
				getMediaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetMediaWithContext(ctx, getMediaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetMedia(getMediaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetMediaWithContext(ctx, getMediaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMediaPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"image": "Image"}`)
				}))
			})
			It(`Invoke GetMedia successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetMedia(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMediaOptions model
				getMediaOptionsModel := new(appidmanagementv4.GetMediaOptions)
				getMediaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetMedia(getMediaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMedia with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetMediaOptions model
				getMediaOptionsModel := new(appidmanagementv4.GetMediaOptions)
				getMediaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetMedia(getMediaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostMedia(postMediaOptions *PostMediaOptions)`, func() {
		tenantID := "testString"
		postMediaPath := "/management/v4/testString/config/ui/media"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postMediaPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["mediaType"]).To(Equal([]string{"logo"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke PostMedia successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.PostMedia(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PostMediaOptions model
				postMediaOptionsModel := new(appidmanagementv4.PostMediaOptions)
				postMediaOptionsModel.MediaType = core.StringPtr("logo")
				postMediaOptionsModel.File = CreateMockReader("This is a mock file.")
				postMediaOptionsModel.FileContentType = core.StringPtr("testString")
				postMediaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.PostMedia(postMediaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PostMedia with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostMediaOptions model
				postMediaOptionsModel := new(appidmanagementv4.PostMediaOptions)
				postMediaOptionsModel.MediaType = core.StringPtr("logo")
				postMediaOptionsModel.File = CreateMockReader("This is a mock file.")
				postMediaOptionsModel.FileContentType = core.StringPtr("testString")
				postMediaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.PostMedia(postMediaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PostMediaOptions model with no property values
				postMediaOptionsModelNew := new(appidmanagementv4.PostMediaOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.PostMedia(postMediaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSamlMetadata(getSamlMetadataOptions *GetSamlMetadataOptions)`, func() {
		tenantID := "testString"
		getSamlMetadataPath := "/management/v4/testString/config/saml_metadata"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSamlMetadataPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"<SPSSODescriptor WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</NameIDFormat><AssertionConsumerService index="1" isDefault="true" Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://us-south.appid.cloud.ibm.com/saml2/v1/login-acs"/></SPSSODescriptor>"`)
				}))
			})
			It(`Invoke GetSamlMetadata successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetSamlMetadataOptions model
				getSamlMetadataOptionsModel := new(appidmanagementv4.GetSamlMetadataOptions)
				getSamlMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetSamlMetadataWithContext(ctx, getSamlMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetSamlMetadata(getSamlMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetSamlMetadataWithContext(ctx, getSamlMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSamlMetadataPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"<SPSSODescriptor WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</NameIDFormat><AssertionConsumerService index="1" isDefault="true" Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://us-south.appid.cloud.ibm.com/saml2/v1/login-acs"/></SPSSODescriptor>"`)
				}))
			})
			It(`Invoke GetSamlMetadata successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetSamlMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSamlMetadataOptions model
				getSamlMetadataOptionsModel := new(appidmanagementv4.GetSamlMetadataOptions)
				getSamlMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetSamlMetadata(getSamlMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSamlMetadata with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetSamlMetadataOptions model
				getSamlMetadataOptionsModel := new(appidmanagementv4.GetSamlMetadataOptions)
				getSamlMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetSamlMetadata(getSamlMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTemplate(getTemplateOptions *GetTemplateOptions) - Operation response error`, func() {
		tenantID := "testString"
		getTemplatePath := "/management/v4/testString/config/cloud_directory/templates/USER_VERIFICATION/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTemplate with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(appidmanagementv4.GetTemplateOptions)
				getTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				getTemplateOptionsModel.Language = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTemplate(getTemplateOptions *GetTemplateOptions)`, func() {
		tenantID := "testString"
		getTemplatePath := "/management/v4/testString/config/cloud_directory/templates/USER_VERIFICATION/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"subject": "Subject", "html_body": "HTMLBody", "base64_encoded_html_body": "Base64EncodedHTMLBody", "plain_text_body": "PlainTextBody"}`)
				}))
			})
			It(`Invoke GetTemplate successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(appidmanagementv4.GetTemplateOptions)
				getTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				getTemplateOptionsModel.Language = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetTemplateWithContext(ctx, getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetTemplateWithContext(ctx, getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"subject": "Subject", "html_body": "HTMLBody", "base64_encoded_html_body": "Base64EncodedHTMLBody", "plain_text_body": "PlainTextBody"}`)
				}))
			})
			It(`Invoke GetTemplate successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(appidmanagementv4.GetTemplateOptions)
				getTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				getTemplateOptionsModel.Language = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTemplate with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(appidmanagementv4.GetTemplateOptions)
				getTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				getTemplateOptionsModel.Language = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTemplateOptions model with no property values
				getTemplateOptionsModelNew := new(appidmanagementv4.GetTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetTemplate(getTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTemplate(updateTemplateOptions *UpdateTemplateOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateTemplatePath := "/management/v4/testString/config/cloud_directory/templates/USER_VERIFICATION/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTemplate with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(appidmanagementv4.UpdateTemplateOptions)
				updateTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				updateTemplateOptionsModel.Language = core.StringPtr("testString")
				updateTemplateOptionsModel.Subject = core.StringPtr("testString")
				updateTemplateOptionsModel.HTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Base64EncodedHTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.PlainTextBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTemplate(updateTemplateOptions *UpdateTemplateOptions)`, func() {
		tenantID := "testString"
		updateTemplatePath := "/management/v4/testString/config/cloud_directory/templates/USER_VERIFICATION/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplatePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"subject": "Subject", "html_body": "HTMLBody", "base64_encoded_html_body": "Base64EncodedHTMLBody", "plain_text_body": "PlainTextBody"}`)
				}))
			})
			It(`Invoke UpdateTemplate successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(appidmanagementv4.UpdateTemplateOptions)
				updateTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				updateTemplateOptionsModel.Language = core.StringPtr("testString")
				updateTemplateOptionsModel.Subject = core.StringPtr("testString")
				updateTemplateOptionsModel.HTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Base64EncodedHTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.PlainTextBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateTemplateWithContext(ctx, updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateTemplateWithContext(ctx, updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplatePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"subject": "Subject", "html_body": "HTMLBody", "base64_encoded_html_body": "Base64EncodedHTMLBody", "plain_text_body": "PlainTextBody"}`)
				}))
			})
			It(`Invoke UpdateTemplate successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(appidmanagementv4.UpdateTemplateOptions)
				updateTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				updateTemplateOptionsModel.Language = core.StringPtr("testString")
				updateTemplateOptionsModel.Subject = core.StringPtr("testString")
				updateTemplateOptionsModel.HTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Base64EncodedHTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.PlainTextBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTemplate with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(appidmanagementv4.UpdateTemplateOptions)
				updateTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				updateTemplateOptionsModel.Language = core.StringPtr("testString")
				updateTemplateOptionsModel.Subject = core.StringPtr("testString")
				updateTemplateOptionsModel.HTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Base64EncodedHTMLBody = core.StringPtr("testString")
				updateTemplateOptionsModel.PlainTextBody = core.StringPtr("testString")
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTemplateOptions model with no property values
				updateTemplateOptionsModelNew := new(appidmanagementv4.UpdateTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateTemplate(updateTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions)`, func() {
		tenantID := "testString"
		deleteTemplatePath := "/management/v4/testString/config/cloud_directory/templates/USER_VERIFICATION/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTemplate successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.DeleteTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTemplateOptions model
				deleteTemplateOptionsModel := new(appidmanagementv4.DeleteTemplateOptions)
				deleteTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				deleteTemplateOptionsModel.Language = core.StringPtr("testString")
				deleteTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.DeleteTemplate(deleteTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTemplate with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteTemplateOptions model
				deleteTemplateOptionsModel := new(appidmanagementv4.DeleteTemplateOptions)
				deleteTemplateOptionsModel.TemplateName = core.StringPtr("USER_VERIFICATION")
				deleteTemplateOptionsModel.Language = core.StringPtr("testString")
				deleteTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.DeleteTemplate(deleteTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTemplateOptions model with no property values
				deleteTemplateOptionsModelNew := new(appidmanagementv4.DeleteTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.DeleteTemplate(deleteTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLocalization(getLocalizationOptions *GetLocalizationOptions) - Operation response error`, func() {
		tenantID := "testString"
		getLocalizationPath := "/management/v4/testString/config/ui/languages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocalizationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLocalization with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetLocalizationOptions model
				getLocalizationOptionsModel := new(appidmanagementv4.GetLocalizationOptions)
				getLocalizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetLocalization(getLocalizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetLocalization(getLocalizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLocalization(getLocalizationOptions *GetLocalizationOptions)`, func() {
		tenantID := "testString"
		getLocalizationPath := "/management/v4/testString/config/ui/languages"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocalizationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"languages": ["Languages"]}`)
				}))
			})
			It(`Invoke GetLocalization successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetLocalizationOptions model
				getLocalizationOptionsModel := new(appidmanagementv4.GetLocalizationOptions)
				getLocalizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetLocalizationWithContext(ctx, getLocalizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetLocalization(getLocalizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetLocalizationWithContext(ctx, getLocalizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocalizationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"languages": ["Languages"]}`)
				}))
			})
			It(`Invoke GetLocalization successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetLocalization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLocalizationOptions model
				getLocalizationOptionsModel := new(appidmanagementv4.GetLocalizationOptions)
				getLocalizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetLocalization(getLocalizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLocalization with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetLocalizationOptions model
				getLocalizationOptionsModel := new(appidmanagementv4.GetLocalizationOptions)
				getLocalizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetLocalization(getLocalizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateLocalization(updateLocalizationOptions *UpdateLocalizationOptions)`, func() {
		tenantID := "testString"
		updateLocalizationPath := "/management/v4/testString/config/ui/languages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLocalizationPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateLocalization successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UpdateLocalization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateLocalizationOptions model
				updateLocalizationOptionsModel := new(appidmanagementv4.UpdateLocalizationOptions)
				updateLocalizationOptionsModel.Languages = []string{"testString"}
				updateLocalizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UpdateLocalization(updateLocalizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateLocalization with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateLocalizationOptions model
				updateLocalizationOptionsModel := new(appidmanagementv4.UpdateLocalizationOptions)
				updateLocalizationOptionsModel.Languages = []string{"testString"}
				updateLocalizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UpdateLocalization(updateLocalizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptions *GetCloudDirectorySenderDetailsOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCloudDirectorySenderDetailsPath := "/management/v4/testString/config/cloud_directory/sender_details"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectorySenderDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCloudDirectorySenderDetails with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectorySenderDetailsOptions model
				getCloudDirectorySenderDetailsOptionsModel := new(appidmanagementv4.GetCloudDirectorySenderDetailsOptions)
				getCloudDirectorySenderDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptions *GetCloudDirectorySenderDetailsOptions)`, func() {
		tenantID := "testString"
		getCloudDirectorySenderDetailsPath := "/management/v4/testString/config/cloud_directory/sender_details"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectorySenderDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"senderDetails": {"from": {"name": "Name", "email": "Email"}, "reply_to": {"name": "Name", "email": "Email"}, "linkExpirationSec": 900}}`)
				}))
			})
			It(`Invoke GetCloudDirectorySenderDetails successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCloudDirectorySenderDetailsOptions model
				getCloudDirectorySenderDetailsOptionsModel := new(appidmanagementv4.GetCloudDirectorySenderDetailsOptions)
				getCloudDirectorySenderDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCloudDirectorySenderDetailsWithContext(ctx, getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCloudDirectorySenderDetailsWithContext(ctx, getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectorySenderDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"senderDetails": {"from": {"name": "Name", "email": "Email"}, "reply_to": {"name": "Name", "email": "Email"}, "linkExpirationSec": 900}}`)
				}))
			})
			It(`Invoke GetCloudDirectorySenderDetails successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCloudDirectorySenderDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCloudDirectorySenderDetailsOptions model
				getCloudDirectorySenderDetailsOptionsModel := new(appidmanagementv4.GetCloudDirectorySenderDetailsOptions)
				getCloudDirectorySenderDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCloudDirectorySenderDetails with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectorySenderDetailsOptions model
				getCloudDirectorySenderDetailsOptionsModel := new(appidmanagementv4.GetCloudDirectorySenderDetailsOptions)
				getCloudDirectorySenderDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCloudDirectorySenderDetails(setCloudDirectorySenderDetailsOptions *SetCloudDirectorySenderDetailsOptions)`, func() {
		tenantID := "testString"
		setCloudDirectorySenderDetailsPath := "/management/v4/testString/config/cloud_directory/sender_details"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectorySenderDetailsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke SetCloudDirectorySenderDetails successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.SetCloudDirectorySenderDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetailsFrom model
				cloudDirectorySenderDetailsSenderDetailsFromModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsFrom)
				cloudDirectorySenderDetailsSenderDetailsFromModel.Name = core.StringPtr("testString")
				cloudDirectorySenderDetailsSenderDetailsFromModel.Email = core.StringPtr("testString")

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetailsReplyTo model
				cloudDirectorySenderDetailsSenderDetailsReplyToModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsReplyTo)
				cloudDirectorySenderDetailsSenderDetailsReplyToModel.Name = core.StringPtr("testString")
				cloudDirectorySenderDetailsSenderDetailsReplyToModel.Email = core.StringPtr("testString")

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetails model
				cloudDirectorySenderDetailsSenderDetailsModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetails)
				cloudDirectorySenderDetailsSenderDetailsModel.From = cloudDirectorySenderDetailsSenderDetailsFromModel
				cloudDirectorySenderDetailsSenderDetailsModel.ReplyTo = cloudDirectorySenderDetailsSenderDetailsReplyToModel
				cloudDirectorySenderDetailsSenderDetailsModel.LinkExpirationSec = core.Float64Ptr(float64(900))

				// Construct an instance of the SetCloudDirectorySenderDetailsOptions model
				setCloudDirectorySenderDetailsOptionsModel := new(appidmanagementv4.SetCloudDirectorySenderDetailsOptions)
				setCloudDirectorySenderDetailsOptionsModel.SenderDetails = cloudDirectorySenderDetailsSenderDetailsModel
				setCloudDirectorySenderDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.SetCloudDirectorySenderDetails(setCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SetCloudDirectorySenderDetails with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetailsFrom model
				cloudDirectorySenderDetailsSenderDetailsFromModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsFrom)
				cloudDirectorySenderDetailsSenderDetailsFromModel.Name = core.StringPtr("testString")
				cloudDirectorySenderDetailsSenderDetailsFromModel.Email = core.StringPtr("testString")

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetailsReplyTo model
				cloudDirectorySenderDetailsSenderDetailsReplyToModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsReplyTo)
				cloudDirectorySenderDetailsSenderDetailsReplyToModel.Name = core.StringPtr("testString")
				cloudDirectorySenderDetailsSenderDetailsReplyToModel.Email = core.StringPtr("testString")

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetails model
				cloudDirectorySenderDetailsSenderDetailsModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetails)
				cloudDirectorySenderDetailsSenderDetailsModel.From = cloudDirectorySenderDetailsSenderDetailsFromModel
				cloudDirectorySenderDetailsSenderDetailsModel.ReplyTo = cloudDirectorySenderDetailsSenderDetailsReplyToModel
				cloudDirectorySenderDetailsSenderDetailsModel.LinkExpirationSec = core.Float64Ptr(float64(900))

				// Construct an instance of the SetCloudDirectorySenderDetailsOptions model
				setCloudDirectorySenderDetailsOptionsModel := new(appidmanagementv4.SetCloudDirectorySenderDetailsOptions)
				setCloudDirectorySenderDetailsOptionsModel.SenderDetails = cloudDirectorySenderDetailsSenderDetailsModel
				setCloudDirectorySenderDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.SetCloudDirectorySenderDetails(setCloudDirectorySenderDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SetCloudDirectorySenderDetailsOptions model with no property values
				setCloudDirectorySenderDetailsOptionsModelNew := new(appidmanagementv4.SetCloudDirectorySenderDetailsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.SetCloudDirectorySenderDetails(setCloudDirectorySenderDetailsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptions *GetCloudDirectoryActionUrlOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCloudDirectoryActionURLPath := "/management/v4/testString/config/cloud_directory/action_url/on_user_verified"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryActionURLPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCloudDirectoryActionURL with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryActionUrlOptions model
				getCloudDirectoryActionUrlOptionsModel := new(appidmanagementv4.GetCloudDirectoryActionUrlOptions)
				getCloudDirectoryActionUrlOptionsModel.Action = core.StringPtr("on_user_verified")
				getCloudDirectoryActionUrlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptions *GetCloudDirectoryActionUrlOptions)`, func() {
		tenantID := "testString"
		getCloudDirectoryActionURLPath := "/management/v4/testString/config/cloud_directory/action_url/on_user_verified"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryActionURLPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"actionUrl": "ActionURL"}`)
				}))
			})
			It(`Invoke GetCloudDirectoryActionURL successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCloudDirectoryActionUrlOptions model
				getCloudDirectoryActionUrlOptionsModel := new(appidmanagementv4.GetCloudDirectoryActionUrlOptions)
				getCloudDirectoryActionUrlOptionsModel.Action = core.StringPtr("on_user_verified")
				getCloudDirectoryActionUrlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCloudDirectoryActionURLWithContext(ctx, getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCloudDirectoryActionURLWithContext(ctx, getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryActionURLPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"actionUrl": "ActionURL"}`)
				}))
			})
			It(`Invoke GetCloudDirectoryActionURL successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCloudDirectoryActionURL(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCloudDirectoryActionUrlOptions model
				getCloudDirectoryActionUrlOptionsModel := new(appidmanagementv4.GetCloudDirectoryActionUrlOptions)
				getCloudDirectoryActionUrlOptionsModel.Action = core.StringPtr("on_user_verified")
				getCloudDirectoryActionUrlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCloudDirectoryActionURL with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryActionUrlOptions model
				getCloudDirectoryActionUrlOptionsModel := new(appidmanagementv4.GetCloudDirectoryActionUrlOptions)
				getCloudDirectoryActionUrlOptionsModel.Action = core.StringPtr("on_user_verified")
				getCloudDirectoryActionUrlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCloudDirectoryActionUrlOptions model with no property values
				getCloudDirectoryActionUrlOptionsModelNew := new(appidmanagementv4.GetCloudDirectoryActionUrlOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCloudDirectoryAction(setCloudDirectoryActionOptions *SetCloudDirectoryActionOptions) - Operation response error`, func() {
		tenantID := "testString"
		setCloudDirectoryActionPath := "/management/v4/testString/config/cloud_directory/action_url/on_user_verified"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryActionPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCloudDirectoryAction with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SetCloudDirectoryActionOptions model
				setCloudDirectoryActionOptionsModel := new(appidmanagementv4.SetCloudDirectoryActionOptions)
				setCloudDirectoryActionOptionsModel.Action = core.StringPtr("on_user_verified")
				setCloudDirectoryActionOptionsModel.ActionURL = core.StringPtr("testString")
				setCloudDirectoryActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAction(setCloudDirectoryActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryAction(setCloudDirectoryActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCloudDirectoryAction(setCloudDirectoryActionOptions *SetCloudDirectoryActionOptions)`, func() {
		tenantID := "testString"
		setCloudDirectoryActionPath := "/management/v4/testString/config/cloud_directory/action_url/on_user_verified"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryActionPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"actionUrl": "ActionURL"}`)
				}))
			})
			It(`Invoke SetCloudDirectoryAction successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the SetCloudDirectoryActionOptions model
				setCloudDirectoryActionOptionsModel := new(appidmanagementv4.SetCloudDirectoryActionOptions)
				setCloudDirectoryActionOptionsModel.Action = core.StringPtr("on_user_verified")
				setCloudDirectoryActionOptionsModel.ActionURL = core.StringPtr("testString")
				setCloudDirectoryActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetCloudDirectoryActionWithContext(ctx, setCloudDirectoryActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAction(setCloudDirectoryActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetCloudDirectoryActionWithContext(ctx, setCloudDirectoryActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryActionPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"actionUrl": "ActionURL"}`)
				}))
			})
			It(`Invoke SetCloudDirectoryAction successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetCloudDirectoryActionOptions model
				setCloudDirectoryActionOptionsModel := new(appidmanagementv4.SetCloudDirectoryActionOptions)
				setCloudDirectoryActionOptionsModel.Action = core.StringPtr("on_user_verified")
				setCloudDirectoryActionOptionsModel.ActionURL = core.StringPtr("testString")
				setCloudDirectoryActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryAction(setCloudDirectoryActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetCloudDirectoryAction with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SetCloudDirectoryActionOptions model
				setCloudDirectoryActionOptionsModel := new(appidmanagementv4.SetCloudDirectoryActionOptions)
				setCloudDirectoryActionOptionsModel.Action = core.StringPtr("on_user_verified")
				setCloudDirectoryActionOptionsModel.ActionURL = core.StringPtr("testString")
				setCloudDirectoryActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAction(setCloudDirectoryActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetCloudDirectoryActionOptions model with no property values
				setCloudDirectoryActionOptionsModelNew := new(appidmanagementv4.SetCloudDirectoryActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryAction(setCloudDirectoryActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteActionURL(deleteActionUrlOptions *DeleteActionUrlOptions)`, func() {
		tenantID := "testString"
		deleteActionURLPath := "/management/v4/testString/config/cloud_directory/action_url/on_user_verified"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteActionURLPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteActionURL successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.DeleteActionURL(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteActionUrlOptions model
				deleteActionUrlOptionsModel := new(appidmanagementv4.DeleteActionUrlOptions)
				deleteActionUrlOptionsModel.Action = core.StringPtr("on_user_verified")
				deleteActionUrlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.DeleteActionURL(deleteActionUrlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteActionURL with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteActionUrlOptions model
				deleteActionUrlOptionsModel := new(appidmanagementv4.DeleteActionUrlOptions)
				deleteActionUrlOptionsModel.Action = core.StringPtr("on_user_verified")
				deleteActionUrlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.DeleteActionURL(deleteActionUrlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteActionUrlOptions model with no property values
				deleteActionUrlOptionsModelNew := new(appidmanagementv4.DeleteActionUrlOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.DeleteActionURL(deleteActionUrlOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptions *GetCloudDirectoryPasswordRegexOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCloudDirectoryPasswordRegexPath := "/management/v4/testString/config/cloud_directory/password_regex"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryPasswordRegexPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCloudDirectoryPasswordRegex with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryPasswordRegexOptions model
				getCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.GetCloudDirectoryPasswordRegexOptions)
				getCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptions *GetCloudDirectoryPasswordRegexOptions)`, func() {
		tenantID := "testString"
		getCloudDirectoryPasswordRegexPath := "/management/v4/testString/config/cloud_directory/password_regex"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryPasswordRegexPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regex": "Regex", "base64_encoded_regex": "Base64EncodedRegex", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke GetCloudDirectoryPasswordRegex successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCloudDirectoryPasswordRegexOptions model
				getCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.GetCloudDirectoryPasswordRegexOptions)
				getCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCloudDirectoryPasswordRegexWithContext(ctx, getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCloudDirectoryPasswordRegexWithContext(ctx, getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryPasswordRegexPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regex": "Regex", "base64_encoded_regex": "Base64EncodedRegex", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke GetCloudDirectoryPasswordRegex successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCloudDirectoryPasswordRegex(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCloudDirectoryPasswordRegexOptions model
				getCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.GetCloudDirectoryPasswordRegexOptions)
				getCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCloudDirectoryPasswordRegex with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryPasswordRegexOptions model
				getCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.GetCloudDirectoryPasswordRegexOptions)
				getCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptions *SetCloudDirectoryPasswordRegexOptions) - Operation response error`, func() {
		tenantID := "testString"
		setCloudDirectoryPasswordRegexPath := "/management/v4/testString/config/cloud_directory/password_regex"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryPasswordRegexPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCloudDirectoryPasswordRegex with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SetCloudDirectoryPasswordRegexOptions model
				setCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.SetCloudDirectoryPasswordRegexOptions)
				setCloudDirectoryPasswordRegexOptionsModel.Regex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Base64EncodedRegex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.ErrorMessage = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptions *SetCloudDirectoryPasswordRegexOptions)`, func() {
		tenantID := "testString"
		setCloudDirectoryPasswordRegexPath := "/management/v4/testString/config/cloud_directory/password_regex"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryPasswordRegexPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regex": "Regex", "base64_encoded_regex": "Base64EncodedRegex", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke SetCloudDirectoryPasswordRegex successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the SetCloudDirectoryPasswordRegexOptions model
				setCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.SetCloudDirectoryPasswordRegexOptions)
				setCloudDirectoryPasswordRegexOptionsModel.Regex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Base64EncodedRegex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.ErrorMessage = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetCloudDirectoryPasswordRegexWithContext(ctx, setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetCloudDirectoryPasswordRegexWithContext(ctx, setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryPasswordRegexPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regex": "Regex", "base64_encoded_regex": "Base64EncodedRegex", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke SetCloudDirectoryPasswordRegex successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetCloudDirectoryPasswordRegex(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetCloudDirectoryPasswordRegexOptions model
				setCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.SetCloudDirectoryPasswordRegexOptions)
				setCloudDirectoryPasswordRegexOptionsModel.Regex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Base64EncodedRegex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.ErrorMessage = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetCloudDirectoryPasswordRegex with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SetCloudDirectoryPasswordRegexOptions model
				setCloudDirectoryPasswordRegexOptionsModel := new(appidmanagementv4.SetCloudDirectoryPasswordRegexOptions)
				setCloudDirectoryPasswordRegexOptionsModel.Regex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Base64EncodedRegex = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.ErrorMessage = core.StringPtr("testString")
				setCloudDirectoryPasswordRegexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptions *GetCloudDirectoryEmailDispatcherOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCloudDirectoryEmailDispatcherPath := "/management/v4/testString/config/cloud_directory/email_dispatcher"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryEmailDispatcherPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCloudDirectoryEmailDispatcher with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryEmailDispatcherOptions model
				getCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.GetCloudDirectoryEmailDispatcherOptions)
				getCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptions *GetCloudDirectoryEmailDispatcherOptions)`, func() {
		tenantID := "testString"
		getCloudDirectoryEmailDispatcherPath := "/management/v4/testString/config/cloud_directory/email_dispatcher"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryEmailDispatcherPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider": "sendgrid", "sendgrid": {"apiKey": "ApiKey"}, "custom": {"url": "URL", "authorization": {"type": "value", "value": "Value", "username": "Username", "password": "Password"}}}`)
				}))
			})
			It(`Invoke GetCloudDirectoryEmailDispatcher successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCloudDirectoryEmailDispatcherOptions model
				getCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.GetCloudDirectoryEmailDispatcherOptions)
				getCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCloudDirectoryEmailDispatcherWithContext(ctx, getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCloudDirectoryEmailDispatcherWithContext(ctx, getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryEmailDispatcherPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider": "sendgrid", "sendgrid": {"apiKey": "ApiKey"}, "custom": {"url": "URL", "authorization": {"type": "value", "value": "Value", "username": "Username", "password": "Password"}}}`)
				}))
			})
			It(`Invoke GetCloudDirectoryEmailDispatcher successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCloudDirectoryEmailDispatcher(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCloudDirectoryEmailDispatcherOptions model
				getCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.GetCloudDirectoryEmailDispatcherOptions)
				getCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCloudDirectoryEmailDispatcher with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryEmailDispatcherOptions model
				getCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.GetCloudDirectoryEmailDispatcherOptions)
				getCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptions *SetCloudDirectoryEmailDispatcherOptions) - Operation response error`, func() {
		tenantID := "testString"
		setCloudDirectoryEmailDispatcherPath := "/management/v4/testString/config/cloud_directory/email_dispatcher"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryEmailDispatcherPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCloudDirectoryEmailDispatcher with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the EmailDispatcherParamsSendgrid model
				emailDispatcherParamsSendgridModel := new(appidmanagementv4.EmailDispatcherParamsSendgrid)
				emailDispatcherParamsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustomAuthorization model
				emailDispatcherParamsCustomAuthorizationModel := new(appidmanagementv4.EmailDispatcherParamsCustomAuthorization)
				emailDispatcherParamsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailDispatcherParamsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustom model
				emailDispatcherParamsCustomModel := new(appidmanagementv4.EmailDispatcherParamsCustom)
				emailDispatcherParamsCustomModel.URL = core.StringPtr("testString")
				emailDispatcherParamsCustomModel.Authorization = emailDispatcherParamsCustomAuthorizationModel

				// Construct an instance of the SetCloudDirectoryEmailDispatcherOptions model
				setCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.SetCloudDirectoryEmailDispatcherOptions)
				setCloudDirectoryEmailDispatcherOptionsModel.Provider = core.StringPtr("sendgrid")
				setCloudDirectoryEmailDispatcherOptionsModel.Sendgrid = emailDispatcherParamsSendgridModel
				setCloudDirectoryEmailDispatcherOptionsModel.Custom = emailDispatcherParamsCustomModel
				setCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptions *SetCloudDirectoryEmailDispatcherOptions)`, func() {
		tenantID := "testString"
		setCloudDirectoryEmailDispatcherPath := "/management/v4/testString/config/cloud_directory/email_dispatcher"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryEmailDispatcherPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider": "sendgrid", "sendgrid": {"apiKey": "ApiKey"}, "custom": {"url": "URL", "authorization": {"type": "value", "value": "Value", "username": "Username", "password": "Password"}}}`)
				}))
			})
			It(`Invoke SetCloudDirectoryEmailDispatcher successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the EmailDispatcherParamsSendgrid model
				emailDispatcherParamsSendgridModel := new(appidmanagementv4.EmailDispatcherParamsSendgrid)
				emailDispatcherParamsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustomAuthorization model
				emailDispatcherParamsCustomAuthorizationModel := new(appidmanagementv4.EmailDispatcherParamsCustomAuthorization)
				emailDispatcherParamsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailDispatcherParamsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustom model
				emailDispatcherParamsCustomModel := new(appidmanagementv4.EmailDispatcherParamsCustom)
				emailDispatcherParamsCustomModel.URL = core.StringPtr("testString")
				emailDispatcherParamsCustomModel.Authorization = emailDispatcherParamsCustomAuthorizationModel

				// Construct an instance of the SetCloudDirectoryEmailDispatcherOptions model
				setCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.SetCloudDirectoryEmailDispatcherOptions)
				setCloudDirectoryEmailDispatcherOptionsModel.Provider = core.StringPtr("sendgrid")
				setCloudDirectoryEmailDispatcherOptionsModel.Sendgrid = emailDispatcherParamsSendgridModel
				setCloudDirectoryEmailDispatcherOptionsModel.Custom = emailDispatcherParamsCustomModel
				setCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetCloudDirectoryEmailDispatcherWithContext(ctx, setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetCloudDirectoryEmailDispatcherWithContext(ctx, setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryEmailDispatcherPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider": "sendgrid", "sendgrid": {"apiKey": "ApiKey"}, "custom": {"url": "URL", "authorization": {"type": "value", "value": "Value", "username": "Username", "password": "Password"}}}`)
				}))
			})
			It(`Invoke SetCloudDirectoryEmailDispatcher successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetCloudDirectoryEmailDispatcher(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EmailDispatcherParamsSendgrid model
				emailDispatcherParamsSendgridModel := new(appidmanagementv4.EmailDispatcherParamsSendgrid)
				emailDispatcherParamsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustomAuthorization model
				emailDispatcherParamsCustomAuthorizationModel := new(appidmanagementv4.EmailDispatcherParamsCustomAuthorization)
				emailDispatcherParamsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailDispatcherParamsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustom model
				emailDispatcherParamsCustomModel := new(appidmanagementv4.EmailDispatcherParamsCustom)
				emailDispatcherParamsCustomModel.URL = core.StringPtr("testString")
				emailDispatcherParamsCustomModel.Authorization = emailDispatcherParamsCustomAuthorizationModel

				// Construct an instance of the SetCloudDirectoryEmailDispatcherOptions model
				setCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.SetCloudDirectoryEmailDispatcherOptions)
				setCloudDirectoryEmailDispatcherOptionsModel.Provider = core.StringPtr("sendgrid")
				setCloudDirectoryEmailDispatcherOptionsModel.Sendgrid = emailDispatcherParamsSendgridModel
				setCloudDirectoryEmailDispatcherOptionsModel.Custom = emailDispatcherParamsCustomModel
				setCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetCloudDirectoryEmailDispatcher with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the EmailDispatcherParamsSendgrid model
				emailDispatcherParamsSendgridModel := new(appidmanagementv4.EmailDispatcherParamsSendgrid)
				emailDispatcherParamsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustomAuthorization model
				emailDispatcherParamsCustomAuthorizationModel := new(appidmanagementv4.EmailDispatcherParamsCustomAuthorization)
				emailDispatcherParamsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailDispatcherParamsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailDispatcherParamsCustom model
				emailDispatcherParamsCustomModel := new(appidmanagementv4.EmailDispatcherParamsCustom)
				emailDispatcherParamsCustomModel.URL = core.StringPtr("testString")
				emailDispatcherParamsCustomModel.Authorization = emailDispatcherParamsCustomAuthorizationModel

				// Construct an instance of the SetCloudDirectoryEmailDispatcherOptions model
				setCloudDirectoryEmailDispatcherOptionsModel := new(appidmanagementv4.SetCloudDirectoryEmailDispatcherOptions)
				setCloudDirectoryEmailDispatcherOptionsModel.Provider = core.StringPtr("sendgrid")
				setCloudDirectoryEmailDispatcherOptionsModel.Sendgrid = emailDispatcherParamsSendgridModel
				setCloudDirectoryEmailDispatcherOptionsModel.Custom = emailDispatcherParamsCustomModel
				setCloudDirectoryEmailDispatcherOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetCloudDirectoryEmailDispatcherOptions model with no property values
				setCloudDirectoryEmailDispatcherOptionsModelNew := new(appidmanagementv4.SetCloudDirectoryEmailDispatcherOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EmailSettingTest(emailSettingTestOptions *EmailSettingTestOptions) - Operation response error`, func() {
		tenantID := "testString"
		emailSettingTestPath := "/management/v4/testString/config/cloud_directory/email_settings/test"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(emailSettingTestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EmailSettingTest with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsSendgrid model
				emailSettingsTestParamsEmailSettingsSendgridModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsSendgrid)
				emailSettingsTestParamsEmailSettingsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustomAuthorization model
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustomAuthorization)
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustom model
				emailSettingsTestParamsEmailSettingsCustomModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustom)
				emailSettingsTestParamsEmailSettingsCustomModel.URL = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomModel.Authorization = emailSettingsTestParamsEmailSettingsCustomAuthorizationModel

				// Construct an instance of the EmailSettingsTestParamsEmailSettings model
				emailSettingsTestParamsEmailSettingsModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettings)
				emailSettingsTestParamsEmailSettingsModel.Provider = core.StringPtr("sendgrid")
				emailSettingsTestParamsEmailSettingsModel.Sendgrid = emailSettingsTestParamsEmailSettingsSendgridModel
				emailSettingsTestParamsEmailSettingsModel.Custom = emailSettingsTestParamsEmailSettingsCustomModel

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsFrom model
				emailSettingsTestParamsSenderDetailsFromModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsFrom)
				emailSettingsTestParamsSenderDetailsFromModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsFromModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsReplyTo model
				emailSettingsTestParamsSenderDetailsReplyToModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsReplyTo)
				emailSettingsTestParamsSenderDetailsReplyToModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsReplyToModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetails model
				emailSettingsTestParamsSenderDetailsModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetails)
				emailSettingsTestParamsSenderDetailsModel.From = emailSettingsTestParamsSenderDetailsFromModel
				emailSettingsTestParamsSenderDetailsModel.ReplyTo = emailSettingsTestParamsSenderDetailsReplyToModel

				// Construct an instance of the EmailSettingTestOptions model
				emailSettingTestOptionsModel := new(appidmanagementv4.EmailSettingTestOptions)
				emailSettingTestOptionsModel.EmailTo = core.StringPtr("testString")
				emailSettingTestOptionsModel.EmailSettings = emailSettingsTestParamsEmailSettingsModel
				emailSettingTestOptionsModel.SenderDetails = emailSettingsTestParamsSenderDetailsModel
				emailSettingTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.EmailSettingTest(emailSettingTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.EmailSettingTest(emailSettingTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EmailSettingTest(emailSettingTestOptions *EmailSettingTestOptions)`, func() {
		tenantID := "testString"
		emailSettingTestPath := "/management/v4/testString/config/cloud_directory/email_settings/test"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(emailSettingTestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false, "dispatcherStatusCode": 20, "dispatcherResponse": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke EmailSettingTest successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsSendgrid model
				emailSettingsTestParamsEmailSettingsSendgridModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsSendgrid)
				emailSettingsTestParamsEmailSettingsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustomAuthorization model
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustomAuthorization)
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustom model
				emailSettingsTestParamsEmailSettingsCustomModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustom)
				emailSettingsTestParamsEmailSettingsCustomModel.URL = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomModel.Authorization = emailSettingsTestParamsEmailSettingsCustomAuthorizationModel

				// Construct an instance of the EmailSettingsTestParamsEmailSettings model
				emailSettingsTestParamsEmailSettingsModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettings)
				emailSettingsTestParamsEmailSettingsModel.Provider = core.StringPtr("sendgrid")
				emailSettingsTestParamsEmailSettingsModel.Sendgrid = emailSettingsTestParamsEmailSettingsSendgridModel
				emailSettingsTestParamsEmailSettingsModel.Custom = emailSettingsTestParamsEmailSettingsCustomModel

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsFrom model
				emailSettingsTestParamsSenderDetailsFromModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsFrom)
				emailSettingsTestParamsSenderDetailsFromModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsFromModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsReplyTo model
				emailSettingsTestParamsSenderDetailsReplyToModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsReplyTo)
				emailSettingsTestParamsSenderDetailsReplyToModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsReplyToModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetails model
				emailSettingsTestParamsSenderDetailsModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetails)
				emailSettingsTestParamsSenderDetailsModel.From = emailSettingsTestParamsSenderDetailsFromModel
				emailSettingsTestParamsSenderDetailsModel.ReplyTo = emailSettingsTestParamsSenderDetailsReplyToModel

				// Construct an instance of the EmailSettingTestOptions model
				emailSettingTestOptionsModel := new(appidmanagementv4.EmailSettingTestOptions)
				emailSettingTestOptionsModel.EmailTo = core.StringPtr("testString")
				emailSettingTestOptionsModel.EmailSettings = emailSettingsTestParamsEmailSettingsModel
				emailSettingTestOptionsModel.SenderDetails = emailSettingsTestParamsSenderDetailsModel
				emailSettingTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.EmailSettingTestWithContext(ctx, emailSettingTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.EmailSettingTest(emailSettingTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.EmailSettingTestWithContext(ctx, emailSettingTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(emailSettingTestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false, "dispatcherStatusCode": 20, "dispatcherResponse": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke EmailSettingTest successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.EmailSettingTest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsSendgrid model
				emailSettingsTestParamsEmailSettingsSendgridModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsSendgrid)
				emailSettingsTestParamsEmailSettingsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustomAuthorization model
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustomAuthorization)
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustom model
				emailSettingsTestParamsEmailSettingsCustomModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustom)
				emailSettingsTestParamsEmailSettingsCustomModel.URL = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomModel.Authorization = emailSettingsTestParamsEmailSettingsCustomAuthorizationModel

				// Construct an instance of the EmailSettingsTestParamsEmailSettings model
				emailSettingsTestParamsEmailSettingsModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettings)
				emailSettingsTestParamsEmailSettingsModel.Provider = core.StringPtr("sendgrid")
				emailSettingsTestParamsEmailSettingsModel.Sendgrid = emailSettingsTestParamsEmailSettingsSendgridModel
				emailSettingsTestParamsEmailSettingsModel.Custom = emailSettingsTestParamsEmailSettingsCustomModel

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsFrom model
				emailSettingsTestParamsSenderDetailsFromModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsFrom)
				emailSettingsTestParamsSenderDetailsFromModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsFromModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsReplyTo model
				emailSettingsTestParamsSenderDetailsReplyToModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsReplyTo)
				emailSettingsTestParamsSenderDetailsReplyToModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsReplyToModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetails model
				emailSettingsTestParamsSenderDetailsModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetails)
				emailSettingsTestParamsSenderDetailsModel.From = emailSettingsTestParamsSenderDetailsFromModel
				emailSettingsTestParamsSenderDetailsModel.ReplyTo = emailSettingsTestParamsSenderDetailsReplyToModel

				// Construct an instance of the EmailSettingTestOptions model
				emailSettingTestOptionsModel := new(appidmanagementv4.EmailSettingTestOptions)
				emailSettingTestOptionsModel.EmailTo = core.StringPtr("testString")
				emailSettingTestOptionsModel.EmailSettings = emailSettingsTestParamsEmailSettingsModel
				emailSettingTestOptionsModel.SenderDetails = emailSettingsTestParamsSenderDetailsModel
				emailSettingTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.EmailSettingTest(emailSettingTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke EmailSettingTest with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsSendgrid model
				emailSettingsTestParamsEmailSettingsSendgridModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsSendgrid)
				emailSettingsTestParamsEmailSettingsSendgridModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustomAuthorization model
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustomAuthorization)
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Password = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustom model
				emailSettingsTestParamsEmailSettingsCustomModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustom)
				emailSettingsTestParamsEmailSettingsCustomModel.URL = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomModel.Authorization = emailSettingsTestParamsEmailSettingsCustomAuthorizationModel

				// Construct an instance of the EmailSettingsTestParamsEmailSettings model
				emailSettingsTestParamsEmailSettingsModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettings)
				emailSettingsTestParamsEmailSettingsModel.Provider = core.StringPtr("sendgrid")
				emailSettingsTestParamsEmailSettingsModel.Sendgrid = emailSettingsTestParamsEmailSettingsSendgridModel
				emailSettingsTestParamsEmailSettingsModel.Custom = emailSettingsTestParamsEmailSettingsCustomModel

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsFrom model
				emailSettingsTestParamsSenderDetailsFromModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsFrom)
				emailSettingsTestParamsSenderDetailsFromModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsFromModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsReplyTo model
				emailSettingsTestParamsSenderDetailsReplyToModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsReplyTo)
				emailSettingsTestParamsSenderDetailsReplyToModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsReplyToModel.Name = core.StringPtr("testString")

				// Construct an instance of the EmailSettingsTestParamsSenderDetails model
				emailSettingsTestParamsSenderDetailsModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetails)
				emailSettingsTestParamsSenderDetailsModel.From = emailSettingsTestParamsSenderDetailsFromModel
				emailSettingsTestParamsSenderDetailsModel.ReplyTo = emailSettingsTestParamsSenderDetailsReplyToModel

				// Construct an instance of the EmailSettingTestOptions model
				emailSettingTestOptionsModel := new(appidmanagementv4.EmailSettingTestOptions)
				emailSettingTestOptionsModel.EmailTo = core.StringPtr("testString")
				emailSettingTestOptionsModel.EmailSettings = emailSettingsTestParamsEmailSettingsModel
				emailSettingTestOptionsModel.SenderDetails = emailSettingsTestParamsSenderDetailsModel
				emailSettingTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.EmailSettingTest(emailSettingTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EmailSettingTestOptions model with no property values
				emailSettingTestOptionsModelNew := new(appidmanagementv4.EmailSettingTestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.EmailSettingTest(emailSettingTestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostEmailDispatcherTest(postEmailDispatcherTestOptions *PostEmailDispatcherTestOptions) - Operation response error`, func() {
		tenantID := "testString"
		postEmailDispatcherTestPath := "/management/v4/testString/config/cloud_directory/email_dispatcher/test"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEmailDispatcherTestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostEmailDispatcherTest with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostEmailDispatcherTestOptions model
				postEmailDispatcherTestOptionsModel := new(appidmanagementv4.PostEmailDispatcherTestOptions)
				postEmailDispatcherTestOptionsModel.Email = core.StringPtr("testString")
				postEmailDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.PostEmailDispatcherTest(postEmailDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.PostEmailDispatcherTest(postEmailDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostEmailDispatcherTest(postEmailDispatcherTestOptions *PostEmailDispatcherTestOptions)`, func() {
		tenantID := "testString"
		postEmailDispatcherTestPath := "/management/v4/testString/config/cloud_directory/email_dispatcher/test"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEmailDispatcherTestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"statusCode": 10, "headers": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke PostEmailDispatcherTest successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the PostEmailDispatcherTestOptions model
				postEmailDispatcherTestOptionsModel := new(appidmanagementv4.PostEmailDispatcherTestOptions)
				postEmailDispatcherTestOptionsModel.Email = core.StringPtr("testString")
				postEmailDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.PostEmailDispatcherTestWithContext(ctx, postEmailDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.PostEmailDispatcherTest(postEmailDispatcherTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.PostEmailDispatcherTestWithContext(ctx, postEmailDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEmailDispatcherTestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"statusCode": 10, "headers": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke PostEmailDispatcherTest successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.PostEmailDispatcherTest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostEmailDispatcherTestOptions model
				postEmailDispatcherTestOptionsModel := new(appidmanagementv4.PostEmailDispatcherTestOptions)
				postEmailDispatcherTestOptionsModel.Email = core.StringPtr("testString")
				postEmailDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.PostEmailDispatcherTest(postEmailDispatcherTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostEmailDispatcherTest with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostEmailDispatcherTestOptions model
				postEmailDispatcherTestOptionsModel := new(appidmanagementv4.PostEmailDispatcherTestOptions)
				postEmailDispatcherTestOptionsModel.Email = core.StringPtr("testString")
				postEmailDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.PostEmailDispatcherTest(postEmailDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostEmailDispatcherTestOptions model with no property values
				postEmailDispatcherTestOptionsModelNew := new(appidmanagementv4.PostEmailDispatcherTestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.PostEmailDispatcherTest(postEmailDispatcherTestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostSmsDispatcherTest(postSmsDispatcherTestOptions *PostSmsDispatcherTestOptions) - Operation response error`, func() {
		tenantID := "testString"
		postSmsDispatcherTestPath := "/management/v4/testString/config/cloud_directory/sms_dispatcher/test"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSmsDispatcherTestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostSmsDispatcherTest with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostSmsDispatcherTestOptions model
				postSmsDispatcherTestOptionsModel := new(appidmanagementv4.PostSmsDispatcherTestOptions)
				postSmsDispatcherTestOptionsModel.PhoneNumber = core.StringPtr("+1-999-999-9999")
				postSmsDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.PostSmsDispatcherTest(postSmsDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.PostSmsDispatcherTest(postSmsDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostSmsDispatcherTest(postSmsDispatcherTestOptions *PostSmsDispatcherTestOptions)`, func() {
		tenantID := "testString"
		postSmsDispatcherTestPath := "/management/v4/testString/config/cloud_directory/sms_dispatcher/test"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSmsDispatcherTestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"confirmationCode": 16, "phoneNumber": "PhoneNumber"}`)
				}))
			})
			It(`Invoke PostSmsDispatcherTest successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the PostSmsDispatcherTestOptions model
				postSmsDispatcherTestOptionsModel := new(appidmanagementv4.PostSmsDispatcherTestOptions)
				postSmsDispatcherTestOptionsModel.PhoneNumber = core.StringPtr("+1-999-999-9999")
				postSmsDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.PostSmsDispatcherTestWithContext(ctx, postSmsDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.PostSmsDispatcherTest(postSmsDispatcherTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.PostSmsDispatcherTestWithContext(ctx, postSmsDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSmsDispatcherTestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"confirmationCode": 16, "phoneNumber": "PhoneNumber"}`)
				}))
			})
			It(`Invoke PostSmsDispatcherTest successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.PostSmsDispatcherTest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostSmsDispatcherTestOptions model
				postSmsDispatcherTestOptionsModel := new(appidmanagementv4.PostSmsDispatcherTestOptions)
				postSmsDispatcherTestOptionsModel.PhoneNumber = core.StringPtr("+1-999-999-9999")
				postSmsDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.PostSmsDispatcherTest(postSmsDispatcherTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostSmsDispatcherTest with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostSmsDispatcherTestOptions model
				postSmsDispatcherTestOptionsModel := new(appidmanagementv4.PostSmsDispatcherTestOptions)
				postSmsDispatcherTestOptionsModel.PhoneNumber = core.StringPtr("+1-999-999-9999")
				postSmsDispatcherTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.PostSmsDispatcherTest(postSmsDispatcherTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostSmsDispatcherTestOptions model with no property values
				postSmsDispatcherTestOptionsModelNew := new(appidmanagementv4.PostSmsDispatcherTestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.PostSmsDispatcherTest(postSmsDispatcherTestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptions *GetCloudDirectoryAdvancedPasswordManagementOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCloudDirectoryAdvancedPasswordManagementPath := "/management/v4/testString/config/cloud_directory/advanced_password_management"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryAdvancedPasswordManagementPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCloudDirectoryAdvancedPasswordManagement with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryAdvancedPasswordManagementOptions model
				getCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.GetCloudDirectoryAdvancedPasswordManagementOptions)
				getCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptions *GetCloudDirectoryAdvancedPasswordManagementOptions)`, func() {
		tenantID := "testString"
		getCloudDirectoryAdvancedPasswordManagementPath := "/management/v4/testString/config/cloud_directory/advanced_password_management"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryAdvancedPasswordManagementPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"advancedPasswordManagement": {"enabled": false, "passwordReuse": {"enabled": false, "config": {"maxPasswordReuse": 1}}, "preventPasswordWithUsername": {"enabled": false}, "passwordExpiration": {"enabled": false, "config": {"daysToExpire": 1}}, "lockOutPolicy": {"enabled": false, "config": {"lockOutTimeSec": 60, "numOfAttempts": 1}}, "minPasswordChangeInterval": {"enabled": false, "config": {"minHoursToChangePassword": 0}}}}`)
				}))
			})
			It(`Invoke GetCloudDirectoryAdvancedPasswordManagement successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCloudDirectoryAdvancedPasswordManagementOptions model
				getCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.GetCloudDirectoryAdvancedPasswordManagementOptions)
				getCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCloudDirectoryAdvancedPasswordManagementWithContext(ctx, getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCloudDirectoryAdvancedPasswordManagementWithContext(ctx, getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryAdvancedPasswordManagementPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"advancedPasswordManagement": {"enabled": false, "passwordReuse": {"enabled": false, "config": {"maxPasswordReuse": 1}}, "preventPasswordWithUsername": {"enabled": false}, "passwordExpiration": {"enabled": false, "config": {"daysToExpire": 1}}, "lockOutPolicy": {"enabled": false, "config": {"lockOutTimeSec": 60, "numOfAttempts": 1}}, "minPasswordChangeInterval": {"enabled": false, "config": {"minHoursToChangePassword": 0}}}}`)
				}))
			})
			It(`Invoke GetCloudDirectoryAdvancedPasswordManagement successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCloudDirectoryAdvancedPasswordManagement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCloudDirectoryAdvancedPasswordManagementOptions model
				getCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.GetCloudDirectoryAdvancedPasswordManagementOptions)
				getCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCloudDirectoryAdvancedPasswordManagement with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryAdvancedPasswordManagementOptions model
				getCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.GetCloudDirectoryAdvancedPasswordManagementOptions)
				getCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptions *SetCloudDirectoryAdvancedPasswordManagementOptions) - Operation response error`, func() {
		tenantID := "testString"
		setCloudDirectoryAdvancedPasswordManagementPath := "/management/v4/testString/config/cloud_directory/advanced_password_management"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryAdvancedPasswordManagementPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCloudDirectoryAdvancedPasswordManagement with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuseConfig model
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel.MaxPasswordReuse = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuse model
				apmSchemaAdvancedPasswordManagementPasswordReuseModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuse)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Config = apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername model
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig model
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel.DaysToExpire = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpiration model
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpiration)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Config = apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig model
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.LockOutTimeSec = core.Float64Ptr(float64(60))
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.NumOfAttempts = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicy model
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicy)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Config = apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel.MinHoursToChangePassword = core.Float64Ptr(float64(0))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Config = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagement model
				apmSchemaAdvancedPasswordManagementModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagement)
				apmSchemaAdvancedPasswordManagementModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementModel.PasswordReuse = apmSchemaAdvancedPasswordManagementPasswordReuseModel
				apmSchemaAdvancedPasswordManagementModel.PreventPasswordWithUsername = apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel
				apmSchemaAdvancedPasswordManagementModel.PasswordExpiration = apmSchemaAdvancedPasswordManagementPasswordExpirationModel
				apmSchemaAdvancedPasswordManagementModel.LockOutPolicy = apmSchemaAdvancedPasswordManagementLockOutPolicyModel
				apmSchemaAdvancedPasswordManagementModel.MinPasswordChangeInterval = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel

				// Construct an instance of the SetCloudDirectoryAdvancedPasswordManagementOptions model
				setCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.SetCloudDirectoryAdvancedPasswordManagementOptions)
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.AdvancedPasswordManagement = apmSchemaAdvancedPasswordManagementModel
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptions *SetCloudDirectoryAdvancedPasswordManagementOptions)`, func() {
		tenantID := "testString"
		setCloudDirectoryAdvancedPasswordManagementPath := "/management/v4/testString/config/cloud_directory/advanced_password_management"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryAdvancedPasswordManagementPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"advancedPasswordManagement": {"enabled": false, "passwordReuse": {"enabled": false, "config": {"maxPasswordReuse": 1}}, "preventPasswordWithUsername": {"enabled": false}, "passwordExpiration": {"enabled": false, "config": {"daysToExpire": 1}}, "lockOutPolicy": {"enabled": false, "config": {"lockOutTimeSec": 60, "numOfAttempts": 1}}, "minPasswordChangeInterval": {"enabled": false, "config": {"minHoursToChangePassword": 0}}}}`)
				}))
			})
			It(`Invoke SetCloudDirectoryAdvancedPasswordManagement successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuseConfig model
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel.MaxPasswordReuse = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuse model
				apmSchemaAdvancedPasswordManagementPasswordReuseModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuse)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Config = apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername model
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig model
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel.DaysToExpire = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpiration model
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpiration)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Config = apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig model
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.LockOutTimeSec = core.Float64Ptr(float64(60))
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.NumOfAttempts = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicy model
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicy)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Config = apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel.MinHoursToChangePassword = core.Float64Ptr(float64(0))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Config = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagement model
				apmSchemaAdvancedPasswordManagementModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagement)
				apmSchemaAdvancedPasswordManagementModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementModel.PasswordReuse = apmSchemaAdvancedPasswordManagementPasswordReuseModel
				apmSchemaAdvancedPasswordManagementModel.PreventPasswordWithUsername = apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel
				apmSchemaAdvancedPasswordManagementModel.PasswordExpiration = apmSchemaAdvancedPasswordManagementPasswordExpirationModel
				apmSchemaAdvancedPasswordManagementModel.LockOutPolicy = apmSchemaAdvancedPasswordManagementLockOutPolicyModel
				apmSchemaAdvancedPasswordManagementModel.MinPasswordChangeInterval = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel

				// Construct an instance of the SetCloudDirectoryAdvancedPasswordManagementOptions model
				setCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.SetCloudDirectoryAdvancedPasswordManagementOptions)
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.AdvancedPasswordManagement = apmSchemaAdvancedPasswordManagementModel
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetCloudDirectoryAdvancedPasswordManagementWithContext(ctx, setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetCloudDirectoryAdvancedPasswordManagementWithContext(ctx, setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryAdvancedPasswordManagementPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"advancedPasswordManagement": {"enabled": false, "passwordReuse": {"enabled": false, "config": {"maxPasswordReuse": 1}}, "preventPasswordWithUsername": {"enabled": false}, "passwordExpiration": {"enabled": false, "config": {"daysToExpire": 1}}, "lockOutPolicy": {"enabled": false, "config": {"lockOutTimeSec": 60, "numOfAttempts": 1}}, "minPasswordChangeInterval": {"enabled": false, "config": {"minHoursToChangePassword": 0}}}}`)
				}))
			})
			It(`Invoke SetCloudDirectoryAdvancedPasswordManagement successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuseConfig model
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel.MaxPasswordReuse = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuse model
				apmSchemaAdvancedPasswordManagementPasswordReuseModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuse)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Config = apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername model
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig model
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel.DaysToExpire = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpiration model
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpiration)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Config = apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig model
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.LockOutTimeSec = core.Float64Ptr(float64(60))
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.NumOfAttempts = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicy model
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicy)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Config = apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel.MinHoursToChangePassword = core.Float64Ptr(float64(0))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Config = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagement model
				apmSchemaAdvancedPasswordManagementModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagement)
				apmSchemaAdvancedPasswordManagementModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementModel.PasswordReuse = apmSchemaAdvancedPasswordManagementPasswordReuseModel
				apmSchemaAdvancedPasswordManagementModel.PreventPasswordWithUsername = apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel
				apmSchemaAdvancedPasswordManagementModel.PasswordExpiration = apmSchemaAdvancedPasswordManagementPasswordExpirationModel
				apmSchemaAdvancedPasswordManagementModel.LockOutPolicy = apmSchemaAdvancedPasswordManagementLockOutPolicyModel
				apmSchemaAdvancedPasswordManagementModel.MinPasswordChangeInterval = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel

				// Construct an instance of the SetCloudDirectoryAdvancedPasswordManagementOptions model
				setCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.SetCloudDirectoryAdvancedPasswordManagementOptions)
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.AdvancedPasswordManagement = apmSchemaAdvancedPasswordManagementModel
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetCloudDirectoryAdvancedPasswordManagement with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuseConfig model
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel.MaxPasswordReuse = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuse model
				apmSchemaAdvancedPasswordManagementPasswordReuseModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuse)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Config = apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername model
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig model
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel.DaysToExpire = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpiration model
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpiration)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Config = apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig model
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.LockOutTimeSec = core.Float64Ptr(float64(60))
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.NumOfAttempts = core.Float64Ptr(float64(1))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicy model
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicy)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Config = apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel.MinHoursToChangePassword = core.Float64Ptr(float64(0))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Config = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel

				// Construct an instance of the ApmSchemaAdvancedPasswordManagement model
				apmSchemaAdvancedPasswordManagementModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagement)
				apmSchemaAdvancedPasswordManagementModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementModel.PasswordReuse = apmSchemaAdvancedPasswordManagementPasswordReuseModel
				apmSchemaAdvancedPasswordManagementModel.PreventPasswordWithUsername = apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel
				apmSchemaAdvancedPasswordManagementModel.PasswordExpiration = apmSchemaAdvancedPasswordManagementPasswordExpirationModel
				apmSchemaAdvancedPasswordManagementModel.LockOutPolicy = apmSchemaAdvancedPasswordManagementLockOutPolicyModel
				apmSchemaAdvancedPasswordManagementModel.MinPasswordChangeInterval = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel

				// Construct an instance of the SetCloudDirectoryAdvancedPasswordManagementOptions model
				setCloudDirectoryAdvancedPasswordManagementOptionsModel := new(appidmanagementv4.SetCloudDirectoryAdvancedPasswordManagementOptions)
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.AdvancedPasswordManagement = apmSchemaAdvancedPasswordManagementModel
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetCloudDirectoryAdvancedPasswordManagementOptions model with no property values
				setCloudDirectoryAdvancedPasswordManagementOptionsModelNew := new(appidmanagementv4.SetCloudDirectoryAdvancedPasswordManagementOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAuditStatus(getAuditStatusOptions *GetAuditStatusOptions)`, func() {
		tenantID := "testString"
		getAuditStatusPath := "/management/v4/testString/config/capture_runtime_activity"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAuditStatusPath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAuditStatus successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.GetAuditStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetAuditStatusOptions model
				getAuditStatusOptionsModel := new(appidmanagementv4.GetAuditStatusOptions)
				getAuditStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.GetAuditStatus(getAuditStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetAuditStatus with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetAuditStatusOptions model
				getAuditStatusOptionsModel := new(appidmanagementv4.GetAuditStatusOptions)
				getAuditStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.GetAuditStatus(getAuditStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetAuditStatus(setAuditStatusOptions *SetAuditStatusOptions)`, func() {
		tenantID := "testString"
		setAuditStatusPath := "/management/v4/testString/config/capture_runtime_activity"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAuditStatusPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke SetAuditStatus successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.SetAuditStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SetAuditStatusOptions model
				setAuditStatusOptionsModel := new(appidmanagementv4.SetAuditStatusOptions)
				setAuditStatusOptionsModel.IsActive = core.BoolPtr(true)
				setAuditStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.SetAuditStatus(setAuditStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SetAuditStatus with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SetAuditStatusOptions model
				setAuditStatusOptionsModel := new(appidmanagementv4.SetAuditStatusOptions)
				setAuditStatusOptionsModel.IsActive = core.BoolPtr(true)
				setAuditStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.SetAuditStatus(setAuditStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SetAuditStatusOptions model with no property values
				setAuditStatusOptionsModelNew := new(appidmanagementv4.SetAuditStatusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.SetAuditStatus(setAuditStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListChannels(listChannelsOptions *ListChannelsOptions) - Operation response error`, func() {
		tenantID := "testString"
		listChannelsPath := "/management/v4/testString/config/cloud_directory/mfa/channels"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listChannelsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListChannels with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListChannelsOptions model
				listChannelsOptionsModel := new(appidmanagementv4.ListChannelsOptions)
				listChannelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.ListChannels(listChannelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.ListChannels(listChannelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListChannels(listChannelsOptions *ListChannelsOptions)`, func() {
		tenantID := "testString"
		listChannelsPath := "/management/v4/testString/config/cloud_directory/mfa/channels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listChannelsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"channels": [{"type": "Type", "isActive": true, "config": {"key": "Key", "secret": "Secret", "from": "From", "provider": "Provider"}}]}`)
				}))
			})
			It(`Invoke ListChannels successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListChannelsOptions model
				listChannelsOptionsModel := new(appidmanagementv4.ListChannelsOptions)
				listChannelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.ListChannelsWithContext(ctx, listChannelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.ListChannels(listChannelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.ListChannelsWithContext(ctx, listChannelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listChannelsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"channels": [{"type": "Type", "isActive": true, "config": {"key": "Key", "secret": "Secret", "from": "From", "provider": "Provider"}}]}`)
				}))
			})
			It(`Invoke ListChannels successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.ListChannels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListChannelsOptions model
				listChannelsOptionsModel := new(appidmanagementv4.ListChannelsOptions)
				listChannelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.ListChannels(listChannelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListChannels with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListChannelsOptions model
				listChannelsOptionsModel := new(appidmanagementv4.ListChannelsOptions)
				listChannelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.ListChannels(listChannelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChannel(getChannelOptions *GetChannelOptions) - Operation response error`, func() {
		tenantID := "testString"
		getChannelPath := "/management/v4/testString/config/cloud_directory/mfa/channels/email"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChannelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChannel with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetChannelOptions model
				getChannelOptionsModel := new(appidmanagementv4.GetChannelOptions)
				getChannelOptionsModel.Channel = core.StringPtr("email")
				getChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetChannel(getChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetChannel(getChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetChannel(getChannelOptions *GetChannelOptions)`, func() {
		tenantID := "testString"
		getChannelPath := "/management/v4/testString/config/cloud_directory/mfa/channels/email"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChannelPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "type": "Type", "config": {"key": "Key", "secret": "Secret", "from": "From", "provider": "Provider"}}`)
				}))
			})
			It(`Invoke GetChannel successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetChannelOptions model
				getChannelOptionsModel := new(appidmanagementv4.GetChannelOptions)
				getChannelOptionsModel.Channel = core.StringPtr("email")
				getChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetChannelWithContext(ctx, getChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetChannel(getChannelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetChannelWithContext(ctx, getChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChannelPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "type": "Type", "config": {"key": "Key", "secret": "Secret", "from": "From", "provider": "Provider"}}`)
				}))
			})
			It(`Invoke GetChannel successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetChannel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChannelOptions model
				getChannelOptionsModel := new(appidmanagementv4.GetChannelOptions)
				getChannelOptionsModel.Channel = core.StringPtr("email")
				getChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetChannel(getChannelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetChannel with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetChannelOptions model
				getChannelOptionsModel := new(appidmanagementv4.GetChannelOptions)
				getChannelOptionsModel.Channel = core.StringPtr("email")
				getChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetChannel(getChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetChannelOptions model with no property values
				getChannelOptionsModelNew := new(appidmanagementv4.GetChannelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetChannel(getChannelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateChannel(updateChannelOptions *UpdateChannelOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateChannelPath := "/management/v4/testString/config/cloud_directory/mfa/channels/email"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateChannelPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateChannel with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateChannelOptions model
				updateChannelOptionsModel := new(appidmanagementv4.UpdateChannelOptions)
				updateChannelOptionsModel.Channel = core.StringPtr("email")
				updateChannelOptionsModel.IsActive = core.BoolPtr(true)
				updateChannelOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateChannel(updateChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateChannel(updateChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateChannel(updateChannelOptions *UpdateChannelOptions)`, func() {
		tenantID := "testString"
		updateChannelPath := "/management/v4/testString/config/cloud_directory/mfa/channels/email"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateChannelPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "type": "Type", "config": {"key": "Key", "secret": "Secret", "from": "From", "provider": "Provider"}}`)
				}))
			})
			It(`Invoke UpdateChannel successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateChannelOptions model
				updateChannelOptionsModel := new(appidmanagementv4.UpdateChannelOptions)
				updateChannelOptionsModel.Channel = core.StringPtr("email")
				updateChannelOptionsModel.IsActive = core.BoolPtr(true)
				updateChannelOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateChannelWithContext(ctx, updateChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateChannel(updateChannelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateChannelWithContext(ctx, updateChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateChannelPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "type": "Type", "config": {"key": "Key", "secret": "Secret", "from": "From", "provider": "Provider"}}`)
				}))
			})
			It(`Invoke UpdateChannel successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateChannel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateChannelOptions model
				updateChannelOptionsModel := new(appidmanagementv4.UpdateChannelOptions)
				updateChannelOptionsModel.Channel = core.StringPtr("email")
				updateChannelOptionsModel.IsActive = core.BoolPtr(true)
				updateChannelOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateChannel(updateChannelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateChannel with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateChannelOptions model
				updateChannelOptionsModel := new(appidmanagementv4.UpdateChannelOptions)
				updateChannelOptionsModel.Channel = core.StringPtr("email")
				updateChannelOptionsModel.IsActive = core.BoolPtr(true)
				updateChannelOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateChannelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateChannel(updateChannelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateChannelOptions model with no property values
				updateChannelOptionsModelNew := new(appidmanagementv4.UpdateChannelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateChannel(updateChannelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetExtensionConfig(getExtensionConfigOptions *GetExtensionConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		getExtensionConfigPath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getExtensionConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetExtensionConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetExtensionConfigOptions model
				getExtensionConfigOptionsModel := new(appidmanagementv4.GetExtensionConfigOptions)
				getExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				getExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetExtensionConfig(getExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetExtensionConfig(getExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetExtensionConfig(getExtensionConfigOptions *GetExtensionConfigOptions)`, func() {
		tenantID := "testString"
		getExtensionConfigPath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getExtensionConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"url": "URL", "headers": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetExtensionConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetExtensionConfigOptions model
				getExtensionConfigOptionsModel := new(appidmanagementv4.GetExtensionConfigOptions)
				getExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				getExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetExtensionConfigWithContext(ctx, getExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetExtensionConfig(getExtensionConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetExtensionConfigWithContext(ctx, getExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getExtensionConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"url": "URL", "headers": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetExtensionConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetExtensionConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetExtensionConfigOptions model
				getExtensionConfigOptionsModel := new(appidmanagementv4.GetExtensionConfigOptions)
				getExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				getExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetExtensionConfig(getExtensionConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetExtensionConfig with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetExtensionConfigOptions model
				getExtensionConfigOptionsModel := new(appidmanagementv4.GetExtensionConfigOptions)
				getExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				getExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetExtensionConfig(getExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetExtensionConfigOptions model with no property values
				getExtensionConfigOptionsModelNew := new(appidmanagementv4.GetExtensionConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetExtensionConfig(getExtensionConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateExtensionConfig(updateExtensionConfigOptions *UpdateExtensionConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateExtensionConfigPath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExtensionConfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateExtensionConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateExtensionConfigConfig model
				updateExtensionConfigConfigModel := new(appidmanagementv4.UpdateExtensionConfigConfig)
				updateExtensionConfigConfigModel.URL = core.StringPtr("testString")
				updateExtensionConfigConfigModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateExtensionConfigOptions model
				updateExtensionConfigOptionsModel := new(appidmanagementv4.UpdateExtensionConfigOptions)
				updateExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionConfigOptionsModel.Config = updateExtensionConfigConfigModel
				updateExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateExtensionConfig(updateExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateExtensionConfig(updateExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateExtensionConfig(updateExtensionConfigOptions *UpdateExtensionConfigOptions)`, func() {
		tenantID := "testString"
		updateExtensionConfigPath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExtensionConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"url": "URL", "headers": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke UpdateExtensionConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateExtensionConfigConfig model
				updateExtensionConfigConfigModel := new(appidmanagementv4.UpdateExtensionConfigConfig)
				updateExtensionConfigConfigModel.URL = core.StringPtr("testString")
				updateExtensionConfigConfigModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateExtensionConfigOptions model
				updateExtensionConfigOptionsModel := new(appidmanagementv4.UpdateExtensionConfigOptions)
				updateExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionConfigOptionsModel.Config = updateExtensionConfigConfigModel
				updateExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateExtensionConfigWithContext(ctx, updateExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateExtensionConfig(updateExtensionConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateExtensionConfigWithContext(ctx, updateExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExtensionConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"url": "URL", "headers": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke UpdateExtensionConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateExtensionConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateExtensionConfigConfig model
				updateExtensionConfigConfigModel := new(appidmanagementv4.UpdateExtensionConfigConfig)
				updateExtensionConfigConfigModel.URL = core.StringPtr("testString")
				updateExtensionConfigConfigModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateExtensionConfigOptions model
				updateExtensionConfigOptionsModel := new(appidmanagementv4.UpdateExtensionConfigOptions)
				updateExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionConfigOptionsModel.Config = updateExtensionConfigConfigModel
				updateExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateExtensionConfig(updateExtensionConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateExtensionConfig with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateExtensionConfigConfig model
				updateExtensionConfigConfigModel := new(appidmanagementv4.UpdateExtensionConfigConfig)
				updateExtensionConfigConfigModel.URL = core.StringPtr("testString")
				updateExtensionConfigConfigModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateExtensionConfigOptions model
				updateExtensionConfigOptionsModel := new(appidmanagementv4.UpdateExtensionConfigOptions)
				updateExtensionConfigOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionConfigOptionsModel.Config = updateExtensionConfigConfigModel
				updateExtensionConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateExtensionConfig(updateExtensionConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateExtensionConfigOptions model with no property values
				updateExtensionConfigOptionsModelNew := new(appidmanagementv4.UpdateExtensionConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateExtensionConfig(updateExtensionConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateExtensionActive(updateExtensionActiveOptions *UpdateExtensionActiveOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateExtensionActivePath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa/active"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExtensionActivePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateExtensionActive with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateExtensionActiveOptions model
				updateExtensionActiveOptionsModel := new(appidmanagementv4.UpdateExtensionActiveOptions)
				updateExtensionActiveOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionActiveOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionActiveOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateExtensionActiveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateExtensionActive(updateExtensionActiveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateExtensionActive(updateExtensionActiveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateExtensionActive(updateExtensionActiveOptions *UpdateExtensionActiveOptions)`, func() {
		tenantID := "testString"
		updateExtensionActivePath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa/active"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExtensionActivePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateExtensionActive successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateExtensionActiveOptions model
				updateExtensionActiveOptionsModel := new(appidmanagementv4.UpdateExtensionActiveOptions)
				updateExtensionActiveOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionActiveOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionActiveOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateExtensionActiveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateExtensionActiveWithContext(ctx, updateExtensionActiveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateExtensionActive(updateExtensionActiveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateExtensionActiveWithContext(ctx, updateExtensionActiveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExtensionActivePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateExtensionActive successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateExtensionActive(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateExtensionActiveOptions model
				updateExtensionActiveOptionsModel := new(appidmanagementv4.UpdateExtensionActiveOptions)
				updateExtensionActiveOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionActiveOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionActiveOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateExtensionActiveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateExtensionActive(updateExtensionActiveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateExtensionActive with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateExtensionActiveOptions model
				updateExtensionActiveOptionsModel := new(appidmanagementv4.UpdateExtensionActiveOptions)
				updateExtensionActiveOptionsModel.Name = core.StringPtr("premfa")
				updateExtensionActiveOptionsModel.IsActive = core.BoolPtr(true)
				updateExtensionActiveOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateExtensionActiveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateExtensionActive(updateExtensionActiveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateExtensionActiveOptions model with no property values
				updateExtensionActiveOptionsModelNew := new(appidmanagementv4.UpdateExtensionActiveOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateExtensionActive(updateExtensionActiveOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostExtensionsTest(postExtensionsTestOptions *PostExtensionsTestOptions) - Operation response error`, func() {
		tenantID := "testString"
		postExtensionsTestPath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa/test"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postExtensionsTestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostExtensionsTest with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostExtensionsTestOptions model
				postExtensionsTestOptionsModel := new(appidmanagementv4.PostExtensionsTestOptions)
				postExtensionsTestOptionsModel.Name = core.StringPtr("premfa")
				postExtensionsTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.PostExtensionsTest(postExtensionsTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.PostExtensionsTest(postExtensionsTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostExtensionsTest(postExtensionsTestOptions *PostExtensionsTestOptions)`, func() {
		tenantID := "testString"
		postExtensionsTestPath := "/management/v4/testString/config/cloud_directory/mfa/extensions/premfa/test"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postExtensionsTestPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"statusCode": 10, "headers": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke PostExtensionsTest successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the PostExtensionsTestOptions model
				postExtensionsTestOptionsModel := new(appidmanagementv4.PostExtensionsTestOptions)
				postExtensionsTestOptionsModel.Name = core.StringPtr("premfa")
				postExtensionsTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.PostExtensionsTestWithContext(ctx, postExtensionsTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.PostExtensionsTest(postExtensionsTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.PostExtensionsTestWithContext(ctx, postExtensionsTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postExtensionsTestPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"statusCode": 10, "headers": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke PostExtensionsTest successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.PostExtensionsTest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostExtensionsTestOptions model
				postExtensionsTestOptionsModel := new(appidmanagementv4.PostExtensionsTestOptions)
				postExtensionsTestOptionsModel.Name = core.StringPtr("premfa")
				postExtensionsTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.PostExtensionsTest(postExtensionsTestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostExtensionsTest with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the PostExtensionsTestOptions model
				postExtensionsTestOptionsModel := new(appidmanagementv4.PostExtensionsTestOptions)
				postExtensionsTestOptionsModel.Name = core.StringPtr("premfa")
				postExtensionsTestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.PostExtensionsTest(postExtensionsTestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostExtensionsTestOptions model with no property values
				postExtensionsTestOptionsModelNew := new(appidmanagementv4.PostExtensionsTestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.PostExtensionsTest(postExtensionsTestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMFAConfig(getMFAConfigOptions *GetMFAConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		getMfaConfigPath := "/management/v4/testString/config/cloud_directory/mfa"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMFAConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetMFAConfigOptions model
				getMfaConfigOptionsModel := new(appidmanagementv4.GetMFAConfigOptions)
				getMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetMFAConfig(getMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetMFAConfig(getMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMFAConfig(getMFAConfigOptions *GetMFAConfigOptions)`, func() {
		tenantID := "testString"
		getMfaConfigPath := "/management/v4/testString/config/cloud_directory/mfa"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true}`)
				}))
			})
			It(`Invoke GetMFAConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetMFAConfigOptions model
				getMfaConfigOptionsModel := new(appidmanagementv4.GetMFAConfigOptions)
				getMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetMFAConfigWithContext(ctx, getMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetMFAConfig(getMfaConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetMFAConfigWithContext(ctx, getMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true}`)
				}))
			})
			It(`Invoke GetMFAConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetMFAConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMFAConfigOptions model
				getMfaConfigOptionsModel := new(appidmanagementv4.GetMFAConfigOptions)
				getMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetMFAConfig(getMfaConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMFAConfig with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetMFAConfigOptions model
				getMfaConfigOptionsModel := new(appidmanagementv4.GetMFAConfigOptions)
				getMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetMFAConfig(getMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMFAConfig(updateMFAConfigOptions *UpdateMFAConfigOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateMfaConfigPath := "/management/v4/testString/config/cloud_directory/mfa"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMfaConfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMFAConfig with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateMFAConfigOptions model
				updateMfaConfigOptionsModel := new(appidmanagementv4.UpdateMFAConfigOptions)
				updateMfaConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateMfaConfigOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateMFAConfig(updateMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateMFAConfig(updateMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMFAConfig(updateMFAConfigOptions *UpdateMFAConfigOptions)`, func() {
		tenantID := "testString"
		updateMfaConfigPath := "/management/v4/testString/config/cloud_directory/mfa"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMfaConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true}`)
				}))
			})
			It(`Invoke UpdateMFAConfig successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateMFAConfigOptions model
				updateMfaConfigOptionsModel := new(appidmanagementv4.UpdateMFAConfigOptions)
				updateMfaConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateMfaConfigOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateMFAConfigWithContext(ctx, updateMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateMFAConfig(updateMfaConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateMFAConfigWithContext(ctx, updateMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMfaConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true}`)
				}))
			})
			It(`Invoke UpdateMFAConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateMFAConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateMFAConfigOptions model
				updateMfaConfigOptionsModel := new(appidmanagementv4.UpdateMFAConfigOptions)
				updateMfaConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateMfaConfigOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateMFAConfig(updateMfaConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateMFAConfig with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateMFAConfigOptions model
				updateMfaConfigOptionsModel := new(appidmanagementv4.UpdateMFAConfigOptions)
				updateMfaConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateMfaConfigOptionsModel.Config = map[string]interface{}{"anyKey": "anyValue"}
				updateMfaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateMFAConfig(updateMfaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateMFAConfigOptions model with no property values
				updateMfaConfigOptionsModelNew := new(appidmanagementv4.UpdateMFAConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateMFAConfig(updateMfaConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSSOConfig(getSSOConfigOptions *GetSSOConfigOptions)`, func() {
		tenantID := "testString"
		getSsoConfigPath := "/management/v4/testString/config/cloud_directory/sso"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSsoConfigPath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSSOConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.GetSSOConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetSSOConfigOptions model
				getSsoConfigOptionsModel := new(appidmanagementv4.GetSSOConfigOptions)
				getSsoConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.GetSSOConfig(getSsoConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetSSOConfig with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetSSOConfigOptions model
				getSsoConfigOptionsModel := new(appidmanagementv4.GetSSOConfigOptions)
				getSsoConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.GetSSOConfig(getSsoConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateSSOConfig(updateSSOConfigOptions *UpdateSSOConfigOptions)`, func() {
		tenantID := "testString"
		updateSsoConfigPath := "/management/v4/testString/config/cloud_directory/sso"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSsoConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateSSOConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UpdateSSOConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateSSOConfigOptions model
				updateSsoConfigOptionsModel := new(appidmanagementv4.UpdateSSOConfigOptions)
				updateSsoConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateSsoConfigOptionsModel.InactivityTimeoutSeconds = core.Float64Ptr(float64(86400))
				updateSsoConfigOptionsModel.LogoutRedirectUris = []string{"http://localhost:3000/logout-callback"}
				updateSsoConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UpdateSSOConfig(updateSsoConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateSSOConfig with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateSSOConfigOptions model
				updateSsoConfigOptionsModel := new(appidmanagementv4.UpdateSSOConfigOptions)
				updateSsoConfigOptionsModel.IsActive = core.BoolPtr(true)
				updateSsoConfigOptionsModel.InactivityTimeoutSeconds = core.Float64Ptr(float64(86400))
				updateSsoConfigOptionsModel.LogoutRedirectUris = []string{"http://localhost:3000/logout-callback"}
				updateSsoConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UpdateSSOConfig(updateSsoConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateSSOConfigOptions model with no property values
				updateSsoConfigOptionsModelNew := new(appidmanagementv4.UpdateSSOConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UpdateSSOConfig(updateSsoConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRateLimitConfig(getRateLimitConfigOptions *GetRateLimitConfigOptions)`, func() {
		tenantID := "testString"
		getRateLimitConfigPath := "/management/v4/testString/config/cloud_directory/rate_limit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRateLimitConfigPath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetRateLimitConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.GetRateLimitConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetRateLimitConfigOptions model
				getRateLimitConfigOptionsModel := new(appidmanagementv4.GetRateLimitConfigOptions)
				getRateLimitConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.GetRateLimitConfig(getRateLimitConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetRateLimitConfig with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetRateLimitConfigOptions model
				getRateLimitConfigOptionsModel := new(appidmanagementv4.GetRateLimitConfigOptions)
				getRateLimitConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.GetRateLimitConfig(getRateLimitConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateRateLimitConfig(updateRateLimitConfigOptions *UpdateRateLimitConfigOptions)`, func() {
		tenantID := "testString"
		updateRateLimitConfigPath := "/management/v4/testString/config/cloud_directory/rate_limit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRateLimitConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateRateLimitConfig successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UpdateRateLimitConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateRateLimitConfigOptions model
				updateRateLimitConfigOptionsModel := new(appidmanagementv4.UpdateRateLimitConfigOptions)
				updateRateLimitConfigOptionsModel.SignUpLimitPerMinute = core.Int64Ptr(int64(50))
				updateRateLimitConfigOptionsModel.SignInLimitPerMinute = core.Int64Ptr(int64(60))
				updateRateLimitConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UpdateRateLimitConfig(updateRateLimitConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateRateLimitConfig with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateRateLimitConfigOptions model
				updateRateLimitConfigOptionsModel := new(appidmanagementv4.UpdateRateLimitConfigOptions)
				updateRateLimitConfigOptionsModel.SignUpLimitPerMinute = core.Int64Ptr(int64(50))
				updateRateLimitConfigOptionsModel.SignInLimitPerMinute = core.Int64Ptr(int64(60))
				updateRateLimitConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UpdateRateLimitConfig(updateRateLimitConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateRateLimitConfigOptions model with no property values
				updateRateLimitConfigOptionsModelNew := new(appidmanagementv4.UpdateRateLimitConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UpdateRateLimitConfig(updateRateLimitConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetFacebookIdp(getFacebookIdpOptions *GetFacebookIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		getFacebookIdpPath := "/management/v4/testString/config/idps/facebook"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFacebookIdpPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFacebookIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetFacebookIdpOptions model
				getFacebookIdpOptionsModel := new(appidmanagementv4.GetFacebookIdpOptions)
				getFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetFacebookIdp(getFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetFacebookIdp(getFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFacebookIdp(getFacebookIdpOptions *GetFacebookIdpOptions)`, func() {
		tenantID := "testString"
		getFacebookIdpPath := "/management/v4/testString/config/idps/facebook"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFacebookIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke GetFacebookIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetFacebookIdpOptions model
				getFacebookIdpOptionsModel := new(appidmanagementv4.GetFacebookIdpOptions)
				getFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetFacebookIdpWithContext(ctx, getFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetFacebookIdp(getFacebookIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetFacebookIdpWithContext(ctx, getFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFacebookIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke GetFacebookIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetFacebookIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFacebookIdpOptions model
				getFacebookIdpOptionsModel := new(appidmanagementv4.GetFacebookIdpOptions)
				getFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetFacebookIdp(getFacebookIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetFacebookIdp with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetFacebookIdpOptions model
				getFacebookIdpOptionsModel := new(appidmanagementv4.GetFacebookIdpOptions)
				getFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetFacebookIdp(getFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetFacebookIdp(setFacebookIdpOptions *SetFacebookIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		setFacebookIdpPath := "/management/v4/testString/config/idps/facebook"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setFacebookIdpPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetFacebookIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetFacebookIdpOptions model
				setFacebookIdpOptionsModel := new(appidmanagementv4.SetFacebookIdpOptions)
				setFacebookIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetFacebookIdp(setFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetFacebookIdp(setFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetFacebookIdp(setFacebookIdpOptions *SetFacebookIdpOptions)`, func() {
		tenantID := "testString"
		setFacebookIdpPath := "/management/v4/testString/config/idps/facebook"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setFacebookIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke SetFacebookIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetFacebookIdpOptions model
				setFacebookIdpOptionsModel := new(appidmanagementv4.SetFacebookIdpOptions)
				setFacebookIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetFacebookIdpWithContext(ctx, setFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetFacebookIdp(setFacebookIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetFacebookIdpWithContext(ctx, setFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setFacebookIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke SetFacebookIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetFacebookIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetFacebookIdpOptions model
				setFacebookIdpOptionsModel := new(appidmanagementv4.SetFacebookIdpOptions)
				setFacebookIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetFacebookIdp(setFacebookIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetFacebookIdp with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetFacebookIdpOptions model
				setFacebookIdpOptionsModel := new(appidmanagementv4.SetFacebookIdpOptions)
				setFacebookIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setFacebookIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetFacebookIdp(setFacebookIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetFacebookIdpOptions model with no property values
				setFacebookIdpOptionsModelNew := new(appidmanagementv4.SetFacebookIdpOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetFacebookIdp(setFacebookIdpOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGoogleIdp(getGoogleIdpOptions *GetGoogleIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		getGoogleIdpPath := "/management/v4/testString/config/idps/google"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGoogleIdpPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGoogleIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetGoogleIdpOptions model
				getGoogleIdpOptionsModel := new(appidmanagementv4.GetGoogleIdpOptions)
				getGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetGoogleIdp(getGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetGoogleIdp(getGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGoogleIdp(getGoogleIdpOptions *GetGoogleIdpOptions)`, func() {
		tenantID := "testString"
		getGoogleIdpPath := "/management/v4/testString/config/idps/google"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGoogleIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke GetGoogleIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetGoogleIdpOptions model
				getGoogleIdpOptionsModel := new(appidmanagementv4.GetGoogleIdpOptions)
				getGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetGoogleIdpWithContext(ctx, getGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetGoogleIdp(getGoogleIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetGoogleIdpWithContext(ctx, getGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGoogleIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke GetGoogleIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetGoogleIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGoogleIdpOptions model
				getGoogleIdpOptionsModel := new(appidmanagementv4.GetGoogleIdpOptions)
				getGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetGoogleIdp(getGoogleIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGoogleIdp with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetGoogleIdpOptions model
				getGoogleIdpOptionsModel := new(appidmanagementv4.GetGoogleIdpOptions)
				getGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetGoogleIdp(getGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetGoogleIdp(setGoogleIdpOptions *SetGoogleIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		setGoogleIdpPath := "/management/v4/testString/config/idps/google"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setGoogleIdpPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetGoogleIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetGoogleIdpOptions model
				setGoogleIdpOptionsModel := new(appidmanagementv4.SetGoogleIdpOptions)
				setGoogleIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetGoogleIdp(setGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetGoogleIdp(setGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetGoogleIdp(setGoogleIdpOptions *SetGoogleIdpOptions)`, func() {
		tenantID := "testString"
		setGoogleIdpPath := "/management/v4/testString/config/idps/google"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setGoogleIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke SetGoogleIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetGoogleIdpOptions model
				setGoogleIdpOptionsModel := new(appidmanagementv4.SetGoogleIdpOptions)
				setGoogleIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetGoogleIdpWithContext(ctx, setGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetGoogleIdp(setGoogleIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetGoogleIdpWithContext(ctx, setGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setGoogleIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"idpId": "IdpID", "secret": "Secret"}}`)
				}))
			})
			It(`Invoke SetGoogleIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetGoogleIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetGoogleIdpOptions model
				setGoogleIdpOptionsModel := new(appidmanagementv4.SetGoogleIdpOptions)
				setGoogleIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetGoogleIdp(setGoogleIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetGoogleIdp with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetGoogleIdpOptions model
				setGoogleIdpOptionsModel := new(appidmanagementv4.SetGoogleIdpOptions)
				setGoogleIdpOptionsModel.Idp = facebookGoogleConfigParamsModel
				setGoogleIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetGoogleIdp(setGoogleIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetGoogleIdpOptions model with no property values
				setGoogleIdpOptionsModelNew := new(appidmanagementv4.SetGoogleIdpOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetGoogleIdp(setGoogleIdpOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomIdp(getCustomIdpOptions *GetCustomIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCustomIdpPath := "/management/v4/testString/config/idps/custom"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomIdpPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCustomIdpOptions model
				getCustomIdpOptionsModel := new(appidmanagementv4.GetCustomIdpOptions)
				getCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCustomIdp(getCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCustomIdp(getCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCustomIdp(getCustomIdpOptions *GetCustomIdpOptions)`, func() {
		tenantID := "testString"
		getCustomIdpPath := "/management/v4/testString/config/idps/custom"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"publicKey": "PublicKey"}}`)
				}))
			})
			It(`Invoke GetCustomIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCustomIdpOptions model
				getCustomIdpOptionsModel := new(appidmanagementv4.GetCustomIdpOptions)
				getCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCustomIdpWithContext(ctx, getCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCustomIdp(getCustomIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCustomIdpWithContext(ctx, getCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"publicKey": "PublicKey"}}`)
				}))
			})
			It(`Invoke GetCustomIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCustomIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomIdpOptions model
				getCustomIdpOptionsModel := new(appidmanagementv4.GetCustomIdpOptions)
				getCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCustomIdp(getCustomIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCustomIdp with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCustomIdpOptions model
				getCustomIdpOptionsModel := new(appidmanagementv4.GetCustomIdpOptions)
				getCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCustomIdp(getCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCustomIdp(setCustomIdpOptions *SetCustomIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		setCustomIdpPath := "/management/v4/testString/config/idps/custom"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCustomIdpPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCustomIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CustomIdPConfigParamsConfig model
				customIdPConfigParamsConfigModel := new(appidmanagementv4.CustomIdPConfigParamsConfig)
				customIdPConfigParamsConfigModel.PublicKey = core.StringPtr("testString")

				// Construct an instance of the SetCustomIdpOptions model
				setCustomIdpOptionsModel := new(appidmanagementv4.SetCustomIdpOptions)
				setCustomIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCustomIdpOptionsModel.Config = customIdPConfigParamsConfigModel
				setCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetCustomIdp(setCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetCustomIdp(setCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCustomIdp(setCustomIdpOptions *SetCustomIdpOptions)`, func() {
		tenantID := "testString"
		setCustomIdpPath := "/management/v4/testString/config/idps/custom"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCustomIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"publicKey": "PublicKey"}}`)
				}))
			})
			It(`Invoke SetCustomIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the CustomIdPConfigParamsConfig model
				customIdPConfigParamsConfigModel := new(appidmanagementv4.CustomIdPConfigParamsConfig)
				customIdPConfigParamsConfigModel.PublicKey = core.StringPtr("testString")

				// Construct an instance of the SetCustomIdpOptions model
				setCustomIdpOptionsModel := new(appidmanagementv4.SetCustomIdpOptions)
				setCustomIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCustomIdpOptionsModel.Config = customIdPConfigParamsConfigModel
				setCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetCustomIdpWithContext(ctx, setCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetCustomIdp(setCustomIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetCustomIdpWithContext(ctx, setCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCustomIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"publicKey": "PublicKey"}}`)
				}))
			})
			It(`Invoke SetCustomIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetCustomIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CustomIdPConfigParamsConfig model
				customIdPConfigParamsConfigModel := new(appidmanagementv4.CustomIdPConfigParamsConfig)
				customIdPConfigParamsConfigModel.PublicKey = core.StringPtr("testString")

				// Construct an instance of the SetCustomIdpOptions model
				setCustomIdpOptionsModel := new(appidmanagementv4.SetCustomIdpOptions)
				setCustomIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCustomIdpOptionsModel.Config = customIdPConfigParamsConfigModel
				setCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetCustomIdp(setCustomIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetCustomIdp with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CustomIdPConfigParamsConfig model
				customIdPConfigParamsConfigModel := new(appidmanagementv4.CustomIdPConfigParamsConfig)
				customIdPConfigParamsConfigModel.PublicKey = core.StringPtr("testString")

				// Construct an instance of the SetCustomIdpOptions model
				setCustomIdpOptionsModel := new(appidmanagementv4.SetCustomIdpOptions)
				setCustomIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCustomIdpOptionsModel.Config = customIdPConfigParamsConfigModel
				setCustomIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetCustomIdp(setCustomIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetCustomIdpOptions model with no property values
				setCustomIdpOptionsModelNew := new(appidmanagementv4.SetCustomIdpOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetCustomIdp(setCustomIdpOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCloudDirectoryIdp(getCloudDirectoryIdpOptions *GetCloudDirectoryIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		getCloudDirectoryIdpPath := "/management/v4/testString/config/idps/cloud_directory"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryIdpPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCloudDirectoryIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryIdpOptions model
				getCloudDirectoryIdpOptionsModel := new(appidmanagementv4.GetCloudDirectoryIdpOptions)
				getCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetCloudDirectoryIdp(getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryIdp(getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCloudDirectoryIdp(getCloudDirectoryIdpOptions *GetCloudDirectoryIdpOptions)`, func() {
		tenantID := "testString"
		getCloudDirectoryIdpPath := "/management/v4/testString/config/idps/cloud_directory"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"selfServiceEnabled": true, "signupEnabled": true, "interactions": {"identityConfirmation": {"accessMode": "FULL", "methods": ["email"]}, "welcomeEnabled": false, "resetPasswordEnabled": false, "resetPasswordNotificationEnable": true}, "identityField": "email"}}`)
				}))
			})
			It(`Invoke GetCloudDirectoryIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCloudDirectoryIdpOptions model
				getCloudDirectoryIdpOptionsModel := new(appidmanagementv4.GetCloudDirectoryIdpOptions)
				getCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetCloudDirectoryIdpWithContext(ctx, getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetCloudDirectoryIdp(getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetCloudDirectoryIdpWithContext(ctx, getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCloudDirectoryIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"selfServiceEnabled": true, "signupEnabled": true, "interactions": {"identityConfirmation": {"accessMode": "FULL", "methods": ["email"]}, "welcomeEnabled": false, "resetPasswordEnabled": false, "resetPasswordNotificationEnable": true}, "identityField": "email"}}`)
				}))
			})
			It(`Invoke GetCloudDirectoryIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetCloudDirectoryIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCloudDirectoryIdpOptions model
				getCloudDirectoryIdpOptionsModel := new(appidmanagementv4.GetCloudDirectoryIdpOptions)
				getCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetCloudDirectoryIdp(getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCloudDirectoryIdp with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetCloudDirectoryIdpOptions model
				getCloudDirectoryIdpOptionsModel := new(appidmanagementv4.GetCloudDirectoryIdpOptions)
				getCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetCloudDirectoryIdp(getCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCloudDirectoryIdp(setCloudDirectoryIdpOptions *SetCloudDirectoryIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		setCloudDirectoryIdpPath := "/management/v4/testString/config/idps/cloud_directory"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryIdpPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCloudDirectoryIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryConfigParamsInteractionsIdentityConfirmation model
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractionsIdentityConfirmation)
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.AccessMode = core.StringPtr("FULL")
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.Methods = []string{"email"}

				// Construct an instance of the CloudDirectoryConfigParamsInteractions model
				cloudDirectoryConfigParamsInteractionsModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractions)
				cloudDirectoryConfigParamsInteractionsModel.IdentityConfirmation = cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel
				cloudDirectoryConfigParamsInteractionsModel.WelcomeEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordNotificationEnable = core.BoolPtr(true)

				// Construct an instance of the CloudDirectoryConfigParams model
				cloudDirectoryConfigParamsModel := new(appidmanagementv4.CloudDirectoryConfigParams)
				cloudDirectoryConfigParamsModel.SelfServiceEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.SignupEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.Interactions = cloudDirectoryConfigParamsInteractionsModel
				cloudDirectoryConfigParamsModel.IdentityField = core.StringPtr("email")

				// Construct an instance of the SetCloudDirectoryIdpOptions model
				setCloudDirectoryIdpOptionsModel := new(appidmanagementv4.SetCloudDirectoryIdpOptions)
				setCloudDirectoryIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCloudDirectoryIdpOptionsModel.Config = cloudDirectoryConfigParamsModel
				setCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetCloudDirectoryIdp(setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryIdp(setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCloudDirectoryIdp(setCloudDirectoryIdpOptions *SetCloudDirectoryIdpOptions)`, func() {
		tenantID := "testString"
		setCloudDirectoryIdpPath := "/management/v4/testString/config/idps/cloud_directory"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"selfServiceEnabled": true, "signupEnabled": true, "interactions": {"identityConfirmation": {"accessMode": "FULL", "methods": ["email"]}, "welcomeEnabled": false, "resetPasswordEnabled": false, "resetPasswordNotificationEnable": true}, "identityField": "email"}}`)
				}))
			})
			It(`Invoke SetCloudDirectoryIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the CloudDirectoryConfigParamsInteractionsIdentityConfirmation model
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractionsIdentityConfirmation)
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.AccessMode = core.StringPtr("FULL")
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.Methods = []string{"email"}

				// Construct an instance of the CloudDirectoryConfigParamsInteractions model
				cloudDirectoryConfigParamsInteractionsModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractions)
				cloudDirectoryConfigParamsInteractionsModel.IdentityConfirmation = cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel
				cloudDirectoryConfigParamsInteractionsModel.WelcomeEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordNotificationEnable = core.BoolPtr(true)

				// Construct an instance of the CloudDirectoryConfigParams model
				cloudDirectoryConfigParamsModel := new(appidmanagementv4.CloudDirectoryConfigParams)
				cloudDirectoryConfigParamsModel.SelfServiceEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.SignupEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.Interactions = cloudDirectoryConfigParamsInteractionsModel
				cloudDirectoryConfigParamsModel.IdentityField = core.StringPtr("email")

				// Construct an instance of the SetCloudDirectoryIdpOptions model
				setCloudDirectoryIdpOptionsModel := new(appidmanagementv4.SetCloudDirectoryIdpOptions)
				setCloudDirectoryIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCloudDirectoryIdpOptionsModel.Config = cloudDirectoryConfigParamsModel
				setCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetCloudDirectoryIdpWithContext(ctx, setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetCloudDirectoryIdp(setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetCloudDirectoryIdpWithContext(ctx, setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setCloudDirectoryIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"selfServiceEnabled": true, "signupEnabled": true, "interactions": {"identityConfirmation": {"accessMode": "FULL", "methods": ["email"]}, "welcomeEnabled": false, "resetPasswordEnabled": false, "resetPasswordNotificationEnable": true}, "identityField": "email"}}`)
				}))
			})
			It(`Invoke SetCloudDirectoryIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetCloudDirectoryIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CloudDirectoryConfigParamsInteractionsIdentityConfirmation model
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractionsIdentityConfirmation)
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.AccessMode = core.StringPtr("FULL")
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.Methods = []string{"email"}

				// Construct an instance of the CloudDirectoryConfigParamsInteractions model
				cloudDirectoryConfigParamsInteractionsModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractions)
				cloudDirectoryConfigParamsInteractionsModel.IdentityConfirmation = cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel
				cloudDirectoryConfigParamsInteractionsModel.WelcomeEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordNotificationEnable = core.BoolPtr(true)

				// Construct an instance of the CloudDirectoryConfigParams model
				cloudDirectoryConfigParamsModel := new(appidmanagementv4.CloudDirectoryConfigParams)
				cloudDirectoryConfigParamsModel.SelfServiceEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.SignupEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.Interactions = cloudDirectoryConfigParamsInteractionsModel
				cloudDirectoryConfigParamsModel.IdentityField = core.StringPtr("email")

				// Construct an instance of the SetCloudDirectoryIdpOptions model
				setCloudDirectoryIdpOptionsModel := new(appidmanagementv4.SetCloudDirectoryIdpOptions)
				setCloudDirectoryIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCloudDirectoryIdpOptionsModel.Config = cloudDirectoryConfigParamsModel
				setCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryIdp(setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetCloudDirectoryIdp with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CloudDirectoryConfigParamsInteractionsIdentityConfirmation model
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractionsIdentityConfirmation)
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.AccessMode = core.StringPtr("FULL")
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.Methods = []string{"email"}

				// Construct an instance of the CloudDirectoryConfigParamsInteractions model
				cloudDirectoryConfigParamsInteractionsModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractions)
				cloudDirectoryConfigParamsInteractionsModel.IdentityConfirmation = cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel
				cloudDirectoryConfigParamsInteractionsModel.WelcomeEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordNotificationEnable = core.BoolPtr(true)

				// Construct an instance of the CloudDirectoryConfigParams model
				cloudDirectoryConfigParamsModel := new(appidmanagementv4.CloudDirectoryConfigParams)
				cloudDirectoryConfigParamsModel.SelfServiceEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.SignupEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.Interactions = cloudDirectoryConfigParamsInteractionsModel
				cloudDirectoryConfigParamsModel.IdentityField = core.StringPtr("email")

				// Construct an instance of the SetCloudDirectoryIdpOptions model
				setCloudDirectoryIdpOptionsModel := new(appidmanagementv4.SetCloudDirectoryIdpOptions)
				setCloudDirectoryIdpOptionsModel.IsActive = core.BoolPtr(true)
				setCloudDirectoryIdpOptionsModel.Config = cloudDirectoryConfigParamsModel
				setCloudDirectoryIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetCloudDirectoryIdp(setCloudDirectoryIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetCloudDirectoryIdpOptions model with no property values
				setCloudDirectoryIdpOptionsModelNew := new(appidmanagementv4.SetCloudDirectoryIdpOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetCloudDirectoryIdp(setCloudDirectoryIdpOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSamlIdp(getSamlIdpOptions *GetSamlIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		getSamlIdpPath := "/management/v4/testString/config/idps/saml"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSamlIdpPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSamlIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetSamlIdpOptions model
				getSamlIdpOptionsModel := new(appidmanagementv4.GetSamlIdpOptions)
				getSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetSamlIdp(getSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetSamlIdp(getSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSamlIdp(getSamlIdpOptions *GetSamlIdpOptions)`, func() {
		tenantID := "testString"
		getSamlIdpPath := "/management/v4/testString/config/idps/saml"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSamlIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"entityID": "EntityID", "signInUrl": "SignInURL", "certificates": ["Certificates"], "displayName": "DisplayName", "authnContext": {"class": ["urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"], "comparison": "exact"}, "signRequest": false, "encryptResponse": false, "includeScoping": false}}`)
				}))
			})
			It(`Invoke GetSamlIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetSamlIdpOptions model
				getSamlIdpOptionsModel := new(appidmanagementv4.GetSamlIdpOptions)
				getSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetSamlIdpWithContext(ctx, getSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetSamlIdp(getSamlIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetSamlIdpWithContext(ctx, getSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSamlIdpPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"entityID": "EntityID", "signInUrl": "SignInURL", "certificates": ["Certificates"], "displayName": "DisplayName", "authnContext": {"class": ["urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"], "comparison": "exact"}, "signRequest": false, "encryptResponse": false, "includeScoping": false}}`)
				}))
			})
			It(`Invoke GetSamlIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetSamlIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSamlIdpOptions model
				getSamlIdpOptionsModel := new(appidmanagementv4.GetSamlIdpOptions)
				getSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetSamlIdp(getSamlIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSamlIdp with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetSamlIdpOptions model
				getSamlIdpOptionsModel := new(appidmanagementv4.GetSamlIdpOptions)
				getSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetSamlIdp(getSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetSamlIdp(setSamlIdpOptions *SetSamlIdpOptions) - Operation response error`, func() {
		tenantID := "testString"
		setSamlIdpPath := "/management/v4/testString/config/idps/saml"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSamlIdpPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetSamlIdp with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SamlConfigParamsAuthnContext model
				samlConfigParamsAuthnContextModel := new(appidmanagementv4.SamlConfigParamsAuthnContext)
				samlConfigParamsAuthnContextModel.Class = []string{"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"}
				samlConfigParamsAuthnContextModel.Comparison = core.StringPtr("exact")

				// Construct an instance of the SamlConfigParams model
				samlConfigParamsModel := new(appidmanagementv4.SamlConfigParams)
				samlConfigParamsModel.EntityID = core.StringPtr("testString")
				samlConfigParamsModel.SignInURL = core.StringPtr("testString")
				samlConfigParamsModel.Certificates = []string{"testString"}
				samlConfigParamsModel.DisplayName = core.StringPtr("testString")
				samlConfigParamsModel.AuthnContext = samlConfigParamsAuthnContextModel
				samlConfigParamsModel.SignRequest = core.BoolPtr(false)
				samlConfigParamsModel.EncryptResponse = core.BoolPtr(false)
				samlConfigParamsModel.IncludeScoping = core.BoolPtr(false)
				samlConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetSamlIdpOptions model
				setSamlIdpOptionsModel := new(appidmanagementv4.SetSamlIdpOptions)
				setSamlIdpOptionsModel.IsActive = core.BoolPtr(true)
				setSamlIdpOptionsModel.Config = samlConfigParamsModel
				setSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.SetSamlIdp(setSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.SetSamlIdp(setSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetSamlIdp(setSamlIdpOptions *SetSamlIdpOptions)`, func() {
		tenantID := "testString"
		setSamlIdpPath := "/management/v4/testString/config/idps/saml"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSamlIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"entityID": "EntityID", "signInUrl": "SignInURL", "certificates": ["Certificates"], "displayName": "DisplayName", "authnContext": {"class": ["urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"], "comparison": "exact"}, "signRequest": false, "encryptResponse": false, "includeScoping": false}, "validation_data": {"certificates": [{"certificate_index": 16, "expiration_timestamp": 19, "warning": "Warning"}]}}`)
				}))
			})
			It(`Invoke SetSamlIdp successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the SamlConfigParamsAuthnContext model
				samlConfigParamsAuthnContextModel := new(appidmanagementv4.SamlConfigParamsAuthnContext)
				samlConfigParamsAuthnContextModel.Class = []string{"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"}
				samlConfigParamsAuthnContextModel.Comparison = core.StringPtr("exact")

				// Construct an instance of the SamlConfigParams model
				samlConfigParamsModel := new(appidmanagementv4.SamlConfigParams)
				samlConfigParamsModel.EntityID = core.StringPtr("testString")
				samlConfigParamsModel.SignInURL = core.StringPtr("testString")
				samlConfigParamsModel.Certificates = []string{"testString"}
				samlConfigParamsModel.DisplayName = core.StringPtr("testString")
				samlConfigParamsModel.AuthnContext = samlConfigParamsAuthnContextModel
				samlConfigParamsModel.SignRequest = core.BoolPtr(false)
				samlConfigParamsModel.EncryptResponse = core.BoolPtr(false)
				samlConfigParamsModel.IncludeScoping = core.BoolPtr(false)
				samlConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetSamlIdpOptions model
				setSamlIdpOptionsModel := new(appidmanagementv4.SetSamlIdpOptions)
				setSamlIdpOptionsModel.IsActive = core.BoolPtr(true)
				setSamlIdpOptionsModel.Config = samlConfigParamsModel
				setSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.SetSamlIdpWithContext(ctx, setSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.SetSamlIdp(setSamlIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.SetSamlIdpWithContext(ctx, setSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSamlIdpPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"isActive": true, "config": {"entityID": "EntityID", "signInUrl": "SignInURL", "certificates": ["Certificates"], "displayName": "DisplayName", "authnContext": {"class": ["urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"], "comparison": "exact"}, "signRequest": false, "encryptResponse": false, "includeScoping": false}, "validation_data": {"certificates": [{"certificate_index": 16, "expiration_timestamp": 19, "warning": "Warning"}]}}`)
				}))
			})
			It(`Invoke SetSamlIdp successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.SetSamlIdp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SamlConfigParamsAuthnContext model
				samlConfigParamsAuthnContextModel := new(appidmanagementv4.SamlConfigParamsAuthnContext)
				samlConfigParamsAuthnContextModel.Class = []string{"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"}
				samlConfigParamsAuthnContextModel.Comparison = core.StringPtr("exact")

				// Construct an instance of the SamlConfigParams model
				samlConfigParamsModel := new(appidmanagementv4.SamlConfigParams)
				samlConfigParamsModel.EntityID = core.StringPtr("testString")
				samlConfigParamsModel.SignInURL = core.StringPtr("testString")
				samlConfigParamsModel.Certificates = []string{"testString"}
				samlConfigParamsModel.DisplayName = core.StringPtr("testString")
				samlConfigParamsModel.AuthnContext = samlConfigParamsAuthnContextModel
				samlConfigParamsModel.SignRequest = core.BoolPtr(false)
				samlConfigParamsModel.EncryptResponse = core.BoolPtr(false)
				samlConfigParamsModel.IncludeScoping = core.BoolPtr(false)
				samlConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetSamlIdpOptions model
				setSamlIdpOptionsModel := new(appidmanagementv4.SetSamlIdpOptions)
				setSamlIdpOptionsModel.IsActive = core.BoolPtr(true)
				setSamlIdpOptionsModel.Config = samlConfigParamsModel
				setSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.SetSamlIdp(setSamlIdpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetSamlIdp with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the SamlConfigParamsAuthnContext model
				samlConfigParamsAuthnContextModel := new(appidmanagementv4.SamlConfigParamsAuthnContext)
				samlConfigParamsAuthnContextModel.Class = []string{"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"}
				samlConfigParamsAuthnContextModel.Comparison = core.StringPtr("exact")

				// Construct an instance of the SamlConfigParams model
				samlConfigParamsModel := new(appidmanagementv4.SamlConfigParams)
				samlConfigParamsModel.EntityID = core.StringPtr("testString")
				samlConfigParamsModel.SignInURL = core.StringPtr("testString")
				samlConfigParamsModel.Certificates = []string{"testString"}
				samlConfigParamsModel.DisplayName = core.StringPtr("testString")
				samlConfigParamsModel.AuthnContext = samlConfigParamsAuthnContextModel
				samlConfigParamsModel.SignRequest = core.BoolPtr(false)
				samlConfigParamsModel.EncryptResponse = core.BoolPtr(false)
				samlConfigParamsModel.IncludeScoping = core.BoolPtr(false)
				samlConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SetSamlIdpOptions model
				setSamlIdpOptionsModel := new(appidmanagementv4.SetSamlIdpOptions)
				setSamlIdpOptionsModel.IsActive = core.BoolPtr(true)
				setSamlIdpOptionsModel.Config = samlConfigParamsModel
				setSamlIdpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.SetSamlIdp(setSamlIdpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetSamlIdpOptions model with no property values
				setSamlIdpOptionsModelNew := new(appidmanagementv4.SetSamlIdpOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.SetSamlIdp(setSamlIdpOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListRoles(listRolesOptions *ListRolesOptions) - Operation response error`, func() {
		tenantID := "testString"
		listRolesPath := "/management/v4/testString/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRolesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRoles with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(appidmanagementv4.ListRolesOptions)
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListRoles(listRolesOptions *ListRolesOptions)`, func() {
		tenantID := "testString"
		listRolesPath := "/management/v4/testString/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRolesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "12345678-1234-1234-1234-123456789012", "name": "adult", "description": "No movie retrictions in place.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}]}`)
				}))
			})
			It(`Invoke ListRoles successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(appidmanagementv4.ListRolesOptions)
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.ListRolesWithContext(ctx, listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.ListRolesWithContext(ctx, listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRolesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "12345678-1234-1234-1234-123456789012", "name": "adult", "description": "No movie retrictions in place.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}]}`)
				}))
			})
			It(`Invoke ListRoles successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.ListRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(appidmanagementv4.ListRolesOptions)
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRoles with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(appidmanagementv4.ListRolesOptions)
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRole(createRoleOptions *CreateRoleOptions) - Operation response error`, func() {
		tenantID := "testString"
		createRolePath := "/management/v4/testString/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRolePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRole with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CreateRoleParamsAccessItem model
				createRoleParamsAccessItemModel := new(appidmanagementv4.CreateRoleParamsAccessItem)
				createRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				createRoleParamsAccessItemModel.Scopes = []string{"cartoons"}

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(appidmanagementv4.CreateRoleOptions)
				createRoleOptionsModel.Name = core.StringPtr("child")
				createRoleOptionsModel.Access = []appidmanagementv4.CreateRoleParamsAccessItem{*createRoleParamsAccessItemModel}
				createRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateRole(createRoleOptions *CreateRoleOptions)`, func() {
		tenantID := "testString"
		createRolePath := "/management/v4/testString/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRolePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789013", "name": "child", "description": "Limits the available movie options to those that might be more appropriate for younger viewers.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}`)
				}))
			})
			It(`Invoke CreateRole successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateRoleParamsAccessItem model
				createRoleParamsAccessItemModel := new(appidmanagementv4.CreateRoleParamsAccessItem)
				createRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				createRoleParamsAccessItemModel.Scopes = []string{"cartoons"}

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(appidmanagementv4.CreateRoleOptions)
				createRoleOptionsModel.Name = core.StringPtr("child")
				createRoleOptionsModel.Access = []appidmanagementv4.CreateRoleParamsAccessItem{*createRoleParamsAccessItemModel}
				createRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.CreateRoleWithContext(ctx, createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.CreateRoleWithContext(ctx, createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRolePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789013", "name": "child", "description": "Limits the available movie options to those that might be more appropriate for younger viewers.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}`)
				}))
			})
			It(`Invoke CreateRole successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.CreateRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateRoleParamsAccessItem model
				createRoleParamsAccessItemModel := new(appidmanagementv4.CreateRoleParamsAccessItem)
				createRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				createRoleParamsAccessItemModel.Scopes = []string{"cartoons"}

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(appidmanagementv4.CreateRoleOptions)
				createRoleOptionsModel.Name = core.StringPtr("child")
				createRoleOptionsModel.Access = []appidmanagementv4.CreateRoleParamsAccessItem{*createRoleParamsAccessItemModel}
				createRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRole with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the CreateRoleParamsAccessItem model
				createRoleParamsAccessItemModel := new(appidmanagementv4.CreateRoleParamsAccessItem)
				createRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				createRoleParamsAccessItemModel.Scopes = []string{"cartoons"}

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(appidmanagementv4.CreateRoleOptions)
				createRoleOptionsModel.Name = core.StringPtr("child")
				createRoleOptionsModel.Access = []appidmanagementv4.CreateRoleParamsAccessItem{*createRoleParamsAccessItemModel}
				createRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRoleOptions model with no property values
				createRoleOptionsModelNew := new(appidmanagementv4.CreateRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.CreateRole(createRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRole(getRoleOptions *GetRoleOptions) - Operation response error`, func() {
		tenantID := "testString"
		getRolePath := "/management/v4/testString/roles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRole with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(appidmanagementv4.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRole(getRoleOptions *GetRoleOptions)`, func() {
		tenantID := "testString"
		getRolePath := "/management/v4/testString/roles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789012", "name": "adult", "description": "No movie retrictions in place.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}`)
				}))
			})
			It(`Invoke GetRole successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(appidmanagementv4.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetRoleWithContext(ctx, getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetRoleWithContext(ctx, getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789012", "name": "adult", "description": "No movie retrictions in place.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}`)
				}))
			})
			It(`Invoke GetRole successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(appidmanagementv4.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRole with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(appidmanagementv4.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRoleOptions model with no property values
				getRoleOptionsModelNew := new(appidmanagementv4.GetRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetRole(getRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRole(updateRoleOptions *UpdateRoleOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateRolePath := "/management/v4/testString/roles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRolePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRole with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleParamsAccessItem model
				updateRoleParamsAccessItemModel := new(appidmanagementv4.UpdateRoleParamsAccessItem)
				updateRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				updateRoleParamsAccessItemModel.Scopes = []string{"cartoons", "animated"}

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(appidmanagementv4.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.Name = core.StringPtr("child")
				updateRoleOptionsModel.Access = []appidmanagementv4.UpdateRoleParamsAccessItem{*updateRoleParamsAccessItemModel}
				updateRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateRole(updateRoleOptions *UpdateRoleOptions)`, func() {
		tenantID := "testString"
		updateRolePath := "/management/v4/testString/roles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRolePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789013", "name": "child", "description": "Limits the available movie options to those that might be more appropriate for younger viewers.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}`)
				}))
			})
			It(`Invoke UpdateRole successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateRoleParamsAccessItem model
				updateRoleParamsAccessItemModel := new(appidmanagementv4.UpdateRoleParamsAccessItem)
				updateRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				updateRoleParamsAccessItemModel.Scopes = []string{"cartoons", "animated"}

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(appidmanagementv4.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.Name = core.StringPtr("child")
				updateRoleOptionsModel.Access = []appidmanagementv4.UpdateRoleParamsAccessItem{*updateRoleParamsAccessItemModel}
				updateRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateRoleWithContext(ctx, updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateRoleWithContext(ctx, updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRolePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789013", "name": "child", "description": "Limits the available movie options to those that might be more appropriate for younger viewers.", "access": [{"application_id": "de33d272-f8a7-4406-8fe8-ab28fd457be5", "scopes": ["Scopes"]}]}`)
				}))
			})
			It(`Invoke UpdateRole successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateRoleParamsAccessItem model
				updateRoleParamsAccessItemModel := new(appidmanagementv4.UpdateRoleParamsAccessItem)
				updateRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				updateRoleParamsAccessItemModel.Scopes = []string{"cartoons", "animated"}

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(appidmanagementv4.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.Name = core.StringPtr("child")
				updateRoleOptionsModel.Access = []appidmanagementv4.UpdateRoleParamsAccessItem{*updateRoleParamsAccessItemModel}
				updateRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRole with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleParamsAccessItem model
				updateRoleParamsAccessItemModel := new(appidmanagementv4.UpdateRoleParamsAccessItem)
				updateRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				updateRoleParamsAccessItemModel.Scopes = []string{"cartoons", "animated"}

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(appidmanagementv4.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.Name = core.StringPtr("child")
				updateRoleOptionsModel.Access = []appidmanagementv4.UpdateRoleParamsAccessItem{*updateRoleParamsAccessItemModel}
				updateRoleOptionsModel.Description = core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRoleOptions model with no property values
				updateRoleOptionsModelNew := new(appidmanagementv4.UpdateRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateRole(updateRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteRole(deleteRoleOptions *DeleteRoleOptions)`, func() {
		tenantID := "testString"
		deleteRolePath := "/management/v4/testString/roles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRolePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteRole successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.DeleteRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRoleOptions model
				deleteRoleOptionsModel := new(appidmanagementv4.DeleteRoleOptions)
				deleteRoleOptionsModel.RoleID = core.StringPtr("testString")
				deleteRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRole with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteRoleOptions model
				deleteRoleOptionsModel := new(appidmanagementv4.DeleteRoleOptions)
				deleteRoleOptionsModel.RoleID = core.StringPtr("testString")
				deleteRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRoleOptions model with no property values
				deleteRoleOptionsModelNew := new(appidmanagementv4.DeleteRoleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.DeleteRole(deleteRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		tenantID := "testString"
		It(`Instantiate service client`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL: "https://appidmanagementv4/api",
				TenantID: core.StringPtr(tenantID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{})
			Expect(appIdManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		tenantID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					URL: "https://testService/api",
					TenantID: core.StringPtr(tenantID),
				})
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
					TenantID: core.StringPtr(tenantID),
				})
				err := appIdManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appIdManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appIdManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appIdManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appIdManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_URL": "https://appidmanagementv4/api",
				"APP_ID_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_ID_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4UsingExternalConfig(&appidmanagementv4.AppIdManagementV4Options{
				URL: "{BAD_URL_STRING",
				TenantID: core.StringPtr(tenantID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(appIdManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appidmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`UsersSearchUserProfile(usersSearchUserProfileOptions *UsersSearchUserProfileOptions) - Operation response error`, func() {
		tenantID := "testString"
		usersSearchUserProfilePath := "/management/v4/testString/users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersSearchUserProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["dataScope"]).To(Equal([]string{"index"}))
					Expect(req.URL.Query()["email"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UsersSearchUserProfile with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersSearchUserProfileOptions model
				usersSearchUserProfileOptionsModel := new(appidmanagementv4.UsersSearchUserProfileOptions)
				usersSearchUserProfileOptionsModel.DataScope = core.StringPtr("index")
				usersSearchUserProfileOptionsModel.Email = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				usersSearchUserProfileOptionsModel.Count = core.Int64Ptr(int64(0))
				usersSearchUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UsersSearchUserProfile(usersSearchUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UsersSearchUserProfile(usersSearchUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UsersSearchUserProfile(usersSearchUserProfileOptions *UsersSearchUserProfileOptions)`, func() {
		tenantID := "testString"
		usersSearchUserProfilePath := "/management/v4/testString/users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersSearchUserProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["dataScope"]).To(Equal([]string{"index"}))
					Expect(req.URL.Query()["email"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"totalResults": 12, "itemsPerPage": 12, "requestOptions": {"startIndex": 10, "count": 5}, "users": [{"id": "ID", "idp": "Idp", "email": "Email"}]}`)
				}))
			})
			It(`Invoke UsersSearchUserProfile successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UsersSearchUserProfileOptions model
				usersSearchUserProfileOptionsModel := new(appidmanagementv4.UsersSearchUserProfileOptions)
				usersSearchUserProfileOptionsModel.DataScope = core.StringPtr("index")
				usersSearchUserProfileOptionsModel.Email = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				usersSearchUserProfileOptionsModel.Count = core.Int64Ptr(int64(0))
				usersSearchUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UsersSearchUserProfileWithContext(ctx, usersSearchUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UsersSearchUserProfile(usersSearchUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UsersSearchUserProfileWithContext(ctx, usersSearchUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersSearchUserProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["dataScope"]).To(Equal([]string{"index"}))
					Expect(req.URL.Query()["email"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"totalResults": 12, "itemsPerPage": 12, "requestOptions": {"startIndex": 10, "count": 5}, "users": [{"id": "ID", "idp": "Idp", "email": "Email"}]}`)
				}))
			})
			It(`Invoke UsersSearchUserProfile successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UsersSearchUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UsersSearchUserProfileOptions model
				usersSearchUserProfileOptionsModel := new(appidmanagementv4.UsersSearchUserProfileOptions)
				usersSearchUserProfileOptionsModel.DataScope = core.StringPtr("index")
				usersSearchUserProfileOptionsModel.Email = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				usersSearchUserProfileOptionsModel.Count = core.Int64Ptr(int64(0))
				usersSearchUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UsersSearchUserProfile(usersSearchUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UsersSearchUserProfile with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersSearchUserProfileOptions model
				usersSearchUserProfileOptionsModel := new(appidmanagementv4.UsersSearchUserProfileOptions)
				usersSearchUserProfileOptionsModel.DataScope = core.StringPtr("index")
				usersSearchUserProfileOptionsModel.Email = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersSearchUserProfileOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				usersSearchUserProfileOptionsModel.Count = core.Int64Ptr(int64(0))
				usersSearchUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UsersSearchUserProfile(usersSearchUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UsersSearchUserProfileOptions model with no property values
				usersSearchUserProfileOptionsModelNew := new(appidmanagementv4.UsersSearchUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UsersSearchUserProfile(usersSearchUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UsersNominateUser(usersNominateUserOptions *UsersNominateUserOptions)`, func() {
		tenantID := "testString"
		usersNominateUserPath := "/management/v4/testString/users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersNominateUserPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(201)
				}))
			})
			It(`Invoke UsersNominateUser successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UsersNominateUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UsersNominateUserParamsProfile model
				usersNominateUserParamsProfileModel := new(appidmanagementv4.UsersNominateUserParamsProfile)
				usersNominateUserParamsProfileModel.Attributes = make(map[string]interface{})

				// Construct an instance of the UsersNominateUserOptions model
				usersNominateUserOptionsModel := new(appidmanagementv4.UsersNominateUserOptions)
				usersNominateUserOptionsModel.Idp = core.StringPtr("saml")
				usersNominateUserOptionsModel.IdpIdentity = core.StringPtr("appid@ibm.com")
				usersNominateUserOptionsModel.Profile = usersNominateUserParamsProfileModel
				usersNominateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UsersNominateUser(usersNominateUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UsersNominateUser with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersNominateUserParamsProfile model
				usersNominateUserParamsProfileModel := new(appidmanagementv4.UsersNominateUserParamsProfile)
				usersNominateUserParamsProfileModel.Attributes = make(map[string]interface{})

				// Construct an instance of the UsersNominateUserOptions model
				usersNominateUserOptionsModel := new(appidmanagementv4.UsersNominateUserOptions)
				usersNominateUserOptionsModel.Idp = core.StringPtr("saml")
				usersNominateUserOptionsModel.IdpIdentity = core.StringPtr("appid@ibm.com")
				usersNominateUserOptionsModel.Profile = usersNominateUserParamsProfileModel
				usersNominateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UsersNominateUser(usersNominateUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UsersNominateUserOptions model with no property values
				usersNominateUserOptionsModelNew := new(appidmanagementv4.UsersNominateUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UsersNominateUser(usersNominateUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UserProfilesExport(userProfilesExportOptions *UserProfilesExportOptions) - Operation response error`, func() {
		tenantID := "testString"
		userProfilesExportPath := "/management/v4/testString/users/export"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userProfilesExportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UserProfilesExport with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UserProfilesExportOptions model
				userProfilesExportOptionsModel := new(appidmanagementv4.UserProfilesExportOptions)
				userProfilesExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				userProfilesExportOptionsModel.Count = core.Int64Ptr(int64(0))
				userProfilesExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UserProfilesExport(userProfilesExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UserProfilesExport(userProfilesExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UserProfilesExport(userProfilesExportOptions *UserProfilesExportOptions)`, func() {
		tenantID := "testString"
		userProfilesExportPath := "/management/v4/testString/users/export"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userProfilesExportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"users": [{"id": "ID", "identities": [{"provider": "Provider", "id": "ID", "idpUserInfo": {"anyKey": "anyValue"}}], "attributes": {"anyKey": "anyValue"}, "name": "Name", "email": "Email", "picture": "Picture", "gender": "Gender", "locale": "Locale", "preferred_username": "PreferredUsername", "idp": "Idp", "hashedIdpId": "HashedIdpID", "hashedEmail": "HashedEmail", "roles": ["Roles"]}]}`)
				}))
			})
			It(`Invoke UserProfilesExport successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UserProfilesExportOptions model
				userProfilesExportOptionsModel := new(appidmanagementv4.UserProfilesExportOptions)
				userProfilesExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				userProfilesExportOptionsModel.Count = core.Int64Ptr(int64(0))
				userProfilesExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UserProfilesExportWithContext(ctx, userProfilesExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UserProfilesExport(userProfilesExportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UserProfilesExportWithContext(ctx, userProfilesExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userProfilesExportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["startIndex"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"users": [{"id": "ID", "identities": [{"provider": "Provider", "id": "ID", "idpUserInfo": {"anyKey": "anyValue"}}], "attributes": {"anyKey": "anyValue"}, "name": "Name", "email": "Email", "picture": "Picture", "gender": "Gender", "locale": "Locale", "preferred_username": "PreferredUsername", "idp": "Idp", "hashedIdpId": "HashedIdpID", "hashedEmail": "HashedEmail", "roles": ["Roles"]}]}`)
				}))
			})
			It(`Invoke UserProfilesExport successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UserProfilesExport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserProfilesExportOptions model
				userProfilesExportOptionsModel := new(appidmanagementv4.UserProfilesExportOptions)
				userProfilesExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				userProfilesExportOptionsModel.Count = core.Int64Ptr(int64(0))
				userProfilesExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UserProfilesExport(userProfilesExportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UserProfilesExport with error: Operation request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UserProfilesExportOptions model
				userProfilesExportOptionsModel := new(appidmanagementv4.UserProfilesExportOptions)
				userProfilesExportOptionsModel.StartIndex = core.Int64Ptr(int64(38))
				userProfilesExportOptionsModel.Count = core.Int64Ptr(int64(0))
				userProfilesExportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UserProfilesExport(userProfilesExportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UserProfilesImport(userProfilesImportOptions *UserProfilesImportOptions) - Operation response error`, func() {
		tenantID := "testString"
		userProfilesImportPath := "/management/v4/testString/users/import"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userProfilesImportPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UserProfilesImport with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ExportUserProfileUsersItemIdentitiesItem model
				exportUserProfileUsersItemIdentitiesItemModel := new(appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem)
				exportUserProfileUsersItemIdentitiesItemModel.Provider = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.IdpUserInfo = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemIdentitiesItemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ExportUserProfileUsersItem model
				exportUserProfileUsersItemModel := new(appidmanagementv4.ExportUserProfileUsersItem)
				exportUserProfileUsersItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Identities = []appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{*exportUserProfileUsersItemIdentitiesItemModel}
				exportUserProfileUsersItemModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemModel.Name = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Email = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Picture = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Gender = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Locale = core.StringPtr("testString")
				exportUserProfileUsersItemModel.PreferredUsername = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Idp = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedIdpID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedEmail = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the UserProfilesImportOptions model
				userProfilesImportOptionsModel := new(appidmanagementv4.UserProfilesImportOptions)
				userProfilesImportOptionsModel.Users = []appidmanagementv4.ExportUserProfileUsersItem{*exportUserProfileUsersItemModel}
				userProfilesImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UserProfilesImport(userProfilesImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UserProfilesImport(userProfilesImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UserProfilesImport(userProfilesImportOptions *UserProfilesImportOptions)`, func() {
		tenantID := "testString"
		userProfilesImportPath := "/management/v4/testString/users/import"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userProfilesImportPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": 5, "failed": 6, "failReasons": [{"originalId": "OriginalID", "idp": "Idp", "error": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke UserProfilesImport successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the ExportUserProfileUsersItemIdentitiesItem model
				exportUserProfileUsersItemIdentitiesItemModel := new(appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem)
				exportUserProfileUsersItemIdentitiesItemModel.Provider = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.IdpUserInfo = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemIdentitiesItemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ExportUserProfileUsersItem model
				exportUserProfileUsersItemModel := new(appidmanagementv4.ExportUserProfileUsersItem)
				exportUserProfileUsersItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Identities = []appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{*exportUserProfileUsersItemIdentitiesItemModel}
				exportUserProfileUsersItemModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemModel.Name = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Email = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Picture = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Gender = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Locale = core.StringPtr("testString")
				exportUserProfileUsersItemModel.PreferredUsername = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Idp = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedIdpID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedEmail = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the UserProfilesImportOptions model
				userProfilesImportOptionsModel := new(appidmanagementv4.UserProfilesImportOptions)
				userProfilesImportOptionsModel.Users = []appidmanagementv4.ExportUserProfileUsersItem{*exportUserProfileUsersItemModel}
				userProfilesImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UserProfilesImportWithContext(ctx, userProfilesImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UserProfilesImport(userProfilesImportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UserProfilesImportWithContext(ctx, userProfilesImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(userProfilesImportPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": 5, "failed": 6, "failReasons": [{"originalId": "OriginalID", "idp": "Idp", "error": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke UserProfilesImport successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UserProfilesImport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExportUserProfileUsersItemIdentitiesItem model
				exportUserProfileUsersItemIdentitiesItemModel := new(appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem)
				exportUserProfileUsersItemIdentitiesItemModel.Provider = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.IdpUserInfo = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemIdentitiesItemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ExportUserProfileUsersItem model
				exportUserProfileUsersItemModel := new(appidmanagementv4.ExportUserProfileUsersItem)
				exportUserProfileUsersItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Identities = []appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{*exportUserProfileUsersItemIdentitiesItemModel}
				exportUserProfileUsersItemModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemModel.Name = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Email = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Picture = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Gender = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Locale = core.StringPtr("testString")
				exportUserProfileUsersItemModel.PreferredUsername = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Idp = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedIdpID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedEmail = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the UserProfilesImportOptions model
				userProfilesImportOptionsModel := new(appidmanagementv4.UserProfilesImportOptions)
				userProfilesImportOptionsModel.Users = []appidmanagementv4.ExportUserProfileUsersItem{*exportUserProfileUsersItemModel}
				userProfilesImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UserProfilesImport(userProfilesImportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UserProfilesImport with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the ExportUserProfileUsersItemIdentitiesItem model
				exportUserProfileUsersItemIdentitiesItemModel := new(appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem)
				exportUserProfileUsersItemIdentitiesItemModel.Provider = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.IdpUserInfo = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemIdentitiesItemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ExportUserProfileUsersItem model
				exportUserProfileUsersItemModel := new(appidmanagementv4.ExportUserProfileUsersItem)
				exportUserProfileUsersItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Identities = []appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{*exportUserProfileUsersItemIdentitiesItemModel}
				exportUserProfileUsersItemModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemModel.Name = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Email = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Picture = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Gender = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Locale = core.StringPtr("testString")
				exportUserProfileUsersItemModel.PreferredUsername = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Idp = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedIdpID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedEmail = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Roles = []string{"testString"}

				// Construct an instance of the UserProfilesImportOptions model
				userProfilesImportOptionsModel := new(appidmanagementv4.UserProfilesImportOptions)
				userProfilesImportOptionsModel.Users = []appidmanagementv4.ExportUserProfileUsersItem{*exportUserProfileUsersItemModel}
				userProfilesImportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UserProfilesImport(userProfilesImportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UserProfilesImportOptions model with no property values
				userProfilesImportOptionsModelNew := new(appidmanagementv4.UserProfilesImportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UserProfilesImport(userProfilesImportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UsersDeleteUserProfile(usersDeleteUserProfileOptions *UsersDeleteUserProfileOptions)`, func() {
		tenantID := "testString"
		usersDeleteUserProfilePath := "/management/v4/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersDeleteUserProfilePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UsersDeleteUserProfile successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UsersDeleteUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UsersDeleteUserProfileOptions model
				usersDeleteUserProfileOptionsModel := new(appidmanagementv4.UsersDeleteUserProfileOptions)
				usersDeleteUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersDeleteUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UsersDeleteUserProfile(usersDeleteUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UsersDeleteUserProfile with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersDeleteUserProfileOptions model
				usersDeleteUserProfileOptionsModel := new(appidmanagementv4.UsersDeleteUserProfileOptions)
				usersDeleteUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersDeleteUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UsersDeleteUserProfile(usersDeleteUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UsersDeleteUserProfileOptions model with no property values
				usersDeleteUserProfileOptionsModelNew := new(appidmanagementv4.UsersDeleteUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UsersDeleteUserProfile(usersDeleteUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UsersRevokeRefreshToken(usersRevokeRefreshTokenOptions *UsersRevokeRefreshTokenOptions)`, func() {
		tenantID := "testString"
		usersRevokeRefreshTokenPath := "/management/v4/testString/users/testString/revoke_refresh_token"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersRevokeRefreshTokenPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UsersRevokeRefreshToken successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UsersRevokeRefreshToken(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UsersRevokeRefreshTokenOptions model
				usersRevokeRefreshTokenOptionsModel := new(appidmanagementv4.UsersRevokeRefreshTokenOptions)
				usersRevokeRefreshTokenOptionsModel.ID = core.StringPtr("testString")
				usersRevokeRefreshTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UsersRevokeRefreshToken(usersRevokeRefreshTokenOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UsersRevokeRefreshToken with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersRevokeRefreshTokenOptions model
				usersRevokeRefreshTokenOptionsModel := new(appidmanagementv4.UsersRevokeRefreshTokenOptions)
				usersRevokeRefreshTokenOptionsModel.ID = core.StringPtr("testString")
				usersRevokeRefreshTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UsersRevokeRefreshToken(usersRevokeRefreshTokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UsersRevokeRefreshTokenOptions model with no property values
				usersRevokeRefreshTokenOptionsModelNew := new(appidmanagementv4.UsersRevokeRefreshTokenOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UsersRevokeRefreshToken(usersRevokeRefreshTokenOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UsersGetUserProfile(usersGetUserProfileOptions *UsersGetUserProfileOptions)`, func() {
		tenantID := "testString"
		usersGetUserProfilePath := "/management/v4/testString/users/testString/profile"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersGetUserProfilePath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UsersGetUserProfile successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UsersGetUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UsersGetUserProfileOptions model
				usersGetUserProfileOptionsModel := new(appidmanagementv4.UsersGetUserProfileOptions)
				usersGetUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersGetUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UsersGetUserProfile(usersGetUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UsersGetUserProfile with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersGetUserProfileOptions model
				usersGetUserProfileOptionsModel := new(appidmanagementv4.UsersGetUserProfileOptions)
				usersGetUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersGetUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UsersGetUserProfile(usersGetUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UsersGetUserProfileOptions model with no property values
				usersGetUserProfileOptionsModelNew := new(appidmanagementv4.UsersGetUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UsersGetUserProfile(usersGetUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UsersSetUserProfile(usersSetUserProfileOptions *UsersSetUserProfileOptions)`, func() {
		tenantID := "testString"
		usersSetUserProfilePath := "/management/v4/testString/users/testString/profile"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(usersSetUserProfilePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UsersSetUserProfile successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appIdManagementService.UsersSetUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UsersSetUserProfileOptions model
				usersSetUserProfileOptionsModel := new(appidmanagementv4.UsersSetUserProfileOptions)
				usersSetUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersSetUserProfileOptionsModel.Attributes = make(map[string]interface{})
				usersSetUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appIdManagementService.UsersSetUserProfile(usersSetUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UsersSetUserProfile with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UsersSetUserProfileOptions model
				usersSetUserProfileOptionsModel := new(appidmanagementv4.UsersSetUserProfileOptions)
				usersSetUserProfileOptionsModel.ID = core.StringPtr("testString")
				usersSetUserProfileOptionsModel.Attributes = make(map[string]interface{})
				usersSetUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appIdManagementService.UsersSetUserProfile(usersSetUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UsersSetUserProfileOptions model with no property values
				usersSetUserProfileOptionsModelNew := new(appidmanagementv4.UsersSetUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appIdManagementService.UsersSetUserProfile(usersSetUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserRoles(getUserRolesOptions *GetUserRolesOptions) - Operation response error`, func() {
		tenantID := "testString"
		getUserRolesPath := "/management/v4/testString/users/testString/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserRolesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserRoles with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserRolesOptions model
				getUserRolesOptionsModel := new(appidmanagementv4.GetUserRolesOptions)
				getUserRolesOptionsModel.ID = core.StringPtr("testString")
				getUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.GetUserRoles(getUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.GetUserRoles(getUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUserRoles(getUserRolesOptions *GetUserRolesOptions)`, func() {
		tenantID := "testString"
		getUserRolesPath := "/management/v4/testString/users/testString/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserRolesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "adult"}]}`)
				}))
			})
			It(`Invoke GetUserRoles successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetUserRolesOptions model
				getUserRolesOptionsModel := new(appidmanagementv4.GetUserRolesOptions)
				getUserRolesOptionsModel.ID = core.StringPtr("testString")
				getUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.GetUserRolesWithContext(ctx, getUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.GetUserRoles(getUserRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.GetUserRolesWithContext(ctx, getUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserRolesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "adult"}]}`)
				}))
			})
			It(`Invoke GetUserRoles successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.GetUserRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserRolesOptions model
				getUserRolesOptionsModel := new(appidmanagementv4.GetUserRolesOptions)
				getUserRolesOptionsModel.ID = core.StringPtr("testString")
				getUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.GetUserRoles(getUserRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUserRoles with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserRolesOptions model
				getUserRolesOptionsModel := new(appidmanagementv4.GetUserRolesOptions)
				getUserRolesOptionsModel.ID = core.StringPtr("testString")
				getUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.GetUserRoles(getUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserRolesOptions model with no property values
				getUserRolesOptionsModelNew := new(appidmanagementv4.GetUserRolesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.GetUserRoles(getUserRolesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUserRoles(updateUserRolesOptions *UpdateUserRolesOptions) - Operation response error`, func() {
		tenantID := "testString"
		updateUserRolesPath := "/management/v4/testString/users/testString/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserRolesPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateUserRoles with error: Operation response processing error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the UpdateUserRolesOptions model
				updateUserRolesOptionsModel := new(appidmanagementv4.UpdateUserRolesOptions)
				updateUserRolesOptionsModel.ID = core.StringPtr("testString")
				updateUserRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				updateUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appIdManagementService.UpdateUserRoles(updateUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appIdManagementService.EnableRetries(0, 0)
				result, response, operationErr = appIdManagementService.UpdateUserRoles(updateUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateUserRoles(updateUserRolesOptions *UpdateUserRolesOptions)`, func() {
		tenantID := "testString"
		updateUserRolesPath := "/management/v4/testString/users/testString/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserRolesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "child"}]}`)
				}))
			})
			It(`Invoke UpdateUserRoles successfully with retries`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())
				appIdManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the UpdateUserRolesOptions model
				updateUserRolesOptionsModel := new(appidmanagementv4.UpdateUserRolesOptions)
				updateUserRolesOptionsModel.ID = core.StringPtr("testString")
				updateUserRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				updateUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appIdManagementService.UpdateUserRolesWithContext(ctx, updateUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appIdManagementService.DisableRetries()
				result, response, operationErr := appIdManagementService.UpdateUserRoles(updateUserRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appIdManagementService.UpdateUserRolesWithContext(ctx, updateUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserRolesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"roles": [{"id": "111c22c3-38ea-4de8-b5d4-338744d83b0f", "name": "child"}]}`)
				}))
			})
			It(`Invoke UpdateUserRoles successfully`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appIdManagementService.UpdateUserRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the UpdateUserRolesOptions model
				updateUserRolesOptionsModel := new(appidmanagementv4.UpdateUserRolesOptions)
				updateUserRolesOptionsModel.ID = core.StringPtr("testString")
				updateUserRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				updateUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appIdManagementService.UpdateUserRoles(updateUserRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateUserRoles with error: Operation validation and request error`, func() {
				appIdManagementService, serviceErr := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					TenantID: core.StringPtr(tenantID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(appIdManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}

				// Construct an instance of the UpdateUserRolesOptions model
				updateUserRolesOptionsModel := new(appidmanagementv4.UpdateUserRolesOptions)
				updateUserRolesOptionsModel.ID = core.StringPtr("testString")
				updateUserRolesOptionsModel.Roles = updateUserRolesParamsRolesModel
				updateUserRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appIdManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appIdManagementService.UpdateUserRoles(updateUserRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateUserRolesOptions model with no property values
				updateUserRolesOptionsModelNew := new(appidmanagementv4.UpdateUserRolesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appIdManagementService.UpdateUserRoles(updateUserRolesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			tenantID := "testString"
			appIdManagementService, _ := appidmanagementv4.NewAppIdManagementV4(&appidmanagementv4.AppIdManagementV4Options{
				URL:           "http://appidmanagementv4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				TenantID: core.StringPtr(tenantID),
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagement successfully`, func() {
				enabled := true
				var passwordReuse *appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuse = nil
				var preventPasswordWithUsername *appidmanagementv4.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername = nil
				var passwordExpiration *appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpiration = nil
				var lockOutPolicy *appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicy = nil
				_, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagement(enabled, passwordReuse, preventPasswordWithUsername, passwordExpiration, lockOutPolicy)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementLockOutPolicy successfully`, func() {
				enabled := true
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementLockOutPolicy(enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementLockOutPolicyConfig successfully`, func() {
				lockOutTimeSec := float64(60)
				numOfAttempts := float64(1)
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementLockOutPolicyConfig(lockOutTimeSec, numOfAttempts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval successfully`, func() {
				enabled := true
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval(enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig successfully`, func() {
				minHoursToChangePassword := float64(0)
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig(minHoursToChangePassword)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementPasswordExpiration successfully`, func() {
				enabled := true
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementPasswordExpiration(enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementPasswordExpirationConfig successfully`, func() {
				daysToExpire := float64(1)
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementPasswordExpirationConfig(daysToExpire)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementPasswordReuse successfully`, func() {
				enabled := true
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementPasswordReuse(enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementPasswordReuseConfig successfully`, func() {
				maxPasswordReuse := float64(1)
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementPasswordReuseConfig(maxPasswordReuse)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername successfully`, func() {
				enabled := true
				model, err := appIdManagementService.NewApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername(enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewChangePasswordOptions successfully`, func() {
				// Construct an instance of the ChangePasswordOptions model
				changePasswordOptionsNewPassword := "testString"
				changePasswordOptionsUUID := "testString"
				changePasswordOptionsModel := appIdManagementService.NewChangePasswordOptions(changePasswordOptionsNewPassword, changePasswordOptionsUUID)
				changePasswordOptionsModel.SetNewPassword("testString")
				changePasswordOptionsModel.SetUUID("testString")
				changePasswordOptionsModel.SetChangedIpAddress("testString")
				changePasswordOptionsModel.SetLanguage("testString")
				changePasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changePasswordOptionsModel).ToNot(BeNil())
				Expect(changePasswordOptionsModel.NewPassword).To(Equal(core.StringPtr("testString")))
				Expect(changePasswordOptionsModel.UUID).To(Equal(core.StringPtr("testString")))
				Expect(changePasswordOptionsModel.ChangedIpAddress).To(Equal(core.StringPtr("testString")))
				Expect(changePasswordOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(changePasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCloudDirectoryConfigParamsInteractions successfully`, func() {
				var identityConfirmation *appidmanagementv4.CloudDirectoryConfigParamsInteractionsIdentityConfirmation = nil
				welcomeEnabled := false
				resetPasswordEnabled := false
				resetPasswordNotificationEnable := true
				_, err := appIdManagementService.NewCloudDirectoryConfigParamsInteractions(identityConfirmation, welcomeEnabled, resetPasswordEnabled, resetPasswordNotificationEnable)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCloudDirectoryConfigParamsInteractionsIdentityConfirmation successfully`, func() {
				accessMode := "FULL"
				model, err := appIdManagementService.NewCloudDirectoryConfigParamsInteractionsIdentityConfirmation(accessMode)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCloudDirectoryExportOptions successfully`, func() {
				// Construct an instance of the CloudDirectoryExportOptions model
				encryptionSecret := "testString"
				cloudDirectoryExportOptionsModel := appIdManagementService.NewCloudDirectoryExportOptions(encryptionSecret)
				cloudDirectoryExportOptionsModel.SetEncryptionSecret("testString")
				cloudDirectoryExportOptionsModel.SetStartIndex(int64(38))
				cloudDirectoryExportOptionsModel.SetCount(int64(0))
				cloudDirectoryExportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cloudDirectoryExportOptionsModel).ToNot(BeNil())
				Expect(cloudDirectoryExportOptionsModel.EncryptionSecret).To(Equal(core.StringPtr("testString")))
				Expect(cloudDirectoryExportOptionsModel.StartIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(cloudDirectoryExportOptionsModel.Count).To(Equal(core.Int64Ptr(int64(0))))
				Expect(cloudDirectoryExportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCloudDirectoryGetUserinfoOptions successfully`, func() {
				// Construct an instance of the CloudDirectoryGetUserinfoOptions model
				userID := "testString"
				cloudDirectoryGetUserinfoOptionsModel := appIdManagementService.NewCloudDirectoryGetUserinfoOptions(userID)
				cloudDirectoryGetUserinfoOptionsModel.SetUserID("testString")
				cloudDirectoryGetUserinfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cloudDirectoryGetUserinfoOptionsModel).ToNot(BeNil())
				Expect(cloudDirectoryGetUserinfoOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(cloudDirectoryGetUserinfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCloudDirectoryImportOptions successfully`, func() {
				// Construct an instance of the ExportUserUsersItemProfile model
				exportUserUsersItemProfileModel := new(appidmanagementv4.ExportUserUsersItemProfile)
				Expect(exportUserUsersItemProfileModel).ToNot(BeNil())
				exportUserUsersItemProfileModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				Expect(exportUserUsersItemProfileModel.Attributes).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the ExportUserUsersItem model
				exportUserUsersItemModel := new(appidmanagementv4.ExportUserUsersItem)
				Expect(exportUserUsersItemModel).ToNot(BeNil())
				exportUserUsersItemModel.ScimUser = map[string]interface{}{"anyKey": "anyValue"}
				exportUserUsersItemModel.PasswordHash = core.StringPtr("testString")
				exportUserUsersItemModel.PasswordHashAlg = core.StringPtr("testString")
				exportUserUsersItemModel.Profile = exportUserUsersItemProfileModel
				exportUserUsersItemModel.Roles = []string{"testString"}
				Expect(exportUserUsersItemModel.ScimUser).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(exportUserUsersItemModel.PasswordHash).To(Equal(core.StringPtr("testString")))
				Expect(exportUserUsersItemModel.PasswordHashAlg).To(Equal(core.StringPtr("testString")))
				Expect(exportUserUsersItemModel.Profile).To(Equal(exportUserUsersItemProfileModel))
				Expect(exportUserUsersItemModel.Roles).To(Equal([]string{"testString"}))

				// Construct an instance of the CloudDirectoryImportOptions model
				encryptionSecret := "testString"
				cloudDirectoryImportOptionsUsers := []appidmanagementv4.ExportUserUsersItem{}
				cloudDirectoryImportOptionsModel := appIdManagementService.NewCloudDirectoryImportOptions(encryptionSecret, cloudDirectoryImportOptionsUsers)
				cloudDirectoryImportOptionsModel.SetEncryptionSecret("testString")
				cloudDirectoryImportOptionsModel.SetUsers([]appidmanagementv4.ExportUserUsersItem{*exportUserUsersItemModel})
				cloudDirectoryImportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cloudDirectoryImportOptionsModel).ToNot(BeNil())
				Expect(cloudDirectoryImportOptionsModel.EncryptionSecret).To(Equal(core.StringPtr("testString")))
				Expect(cloudDirectoryImportOptionsModel.Users).To(Equal([]appidmanagementv4.ExportUserUsersItem{*exportUserUsersItemModel}))
				Expect(cloudDirectoryImportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCloudDirectoryRemoveOptions successfully`, func() {
				// Construct an instance of the CloudDirectoryRemoveOptions model
				userID := "testString"
				cloudDirectoryRemoveOptionsModel := appIdManagementService.NewCloudDirectoryRemoveOptions(userID)
				cloudDirectoryRemoveOptionsModel.SetUserID("testString")
				cloudDirectoryRemoveOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cloudDirectoryRemoveOptionsModel).ToNot(BeNil())
				Expect(cloudDirectoryRemoveOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(cloudDirectoryRemoveOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCloudDirectorySenderDetailsSenderDetails successfully`, func() {
				var from *appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsFrom = nil
				_, err := appIdManagementService.NewCloudDirectorySenderDetailsSenderDetails(from)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCloudDirectorySenderDetailsSenderDetailsFrom successfully`, func() {
				email := "testString"
				model, err := appIdManagementService.NewCloudDirectorySenderDetailsSenderDetailsFrom(email)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCloudDirectoryUserOptions successfully`, func() {
				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				Expect(createNewUserEmailsItemModel).ToNot(BeNil())
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)
				Expect(createNewUserEmailsItemModel.Value).To(Equal(core.StringPtr("user@mail.com")))
				Expect(createNewUserEmailsItemModel.Primary).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateCloudDirectoryUserOptions model
				createCloudDirectoryUserOptionsEmails := []appidmanagementv4.CreateNewUserEmailsItem{}
				createCloudDirectoryUserOptionsPassword := "userPassword"
				createCloudDirectoryUserOptionsModel := appIdManagementService.NewCreateCloudDirectoryUserOptions(createCloudDirectoryUserOptionsEmails, createCloudDirectoryUserOptionsPassword)
				createCloudDirectoryUserOptionsModel.SetEmails([]appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel})
				createCloudDirectoryUserOptionsModel.SetPassword("userPassword")
				createCloudDirectoryUserOptionsModel.SetActive(true)
				createCloudDirectoryUserOptionsModel.SetUserName("myUserName")
				createCloudDirectoryUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCloudDirectoryUserOptionsModel).ToNot(BeNil())
				Expect(createCloudDirectoryUserOptionsModel.Emails).To(Equal([]appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}))
				Expect(createCloudDirectoryUserOptionsModel.Password).To(Equal(core.StringPtr("userPassword")))
				Expect(createCloudDirectoryUserOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(createCloudDirectoryUserOptionsModel.UserName).To(Equal(core.StringPtr("myUserName")))
				Expect(createCloudDirectoryUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateNewUserEmailsItem successfully`, func() {
				value := "user@mail.com"
				model, err := appIdManagementService.NewCreateNewUserEmailsItem(value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateRoleOptions successfully`, func() {
				// Construct an instance of the CreateRoleParamsAccessItem model
				createRoleParamsAccessItemModel := new(appidmanagementv4.CreateRoleParamsAccessItem)
				Expect(createRoleParamsAccessItemModel).ToNot(BeNil())
				createRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				createRoleParamsAccessItemModel.Scopes = []string{"cartoons"}
				Expect(createRoleParamsAccessItemModel.ApplicationID).To(Equal(core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")))
				Expect(createRoleParamsAccessItemModel.Scopes).To(Equal([]string{"cartoons"}))

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsName := "child"
				createRoleOptionsAccess := []appidmanagementv4.CreateRoleParamsAccessItem{}
				createRoleOptionsModel := appIdManagementService.NewCreateRoleOptions(createRoleOptionsName, createRoleOptionsAccess)
				createRoleOptionsModel.SetName("child")
				createRoleOptionsModel.SetAccess([]appidmanagementv4.CreateRoleParamsAccessItem{*createRoleParamsAccessItemModel})
				createRoleOptionsModel.SetDescription("Limits the available movie options to those that might be more appropriate for younger viewers.")
				createRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRoleOptionsModel).ToNot(BeNil())
				Expect(createRoleOptionsModel.Name).To(Equal(core.StringPtr("child")))
				Expect(createRoleOptionsModel.Access).To(Equal([]appidmanagementv4.CreateRoleParamsAccessItem{*createRoleParamsAccessItemModel}))
				Expect(createRoleOptionsModel.Description).To(Equal(core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")))
				Expect(createRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRoleParamsAccessItem successfully`, func() {
				applicationID := "de33d272-f8a7-4406-8fe8-ab28fd457be5"
				scopes := []string{"cartoons"}
				model, err := appIdManagementService.NewCreateRoleParamsAccessItem(applicationID, scopes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteActionUrlOptions successfully`, func() {
				// Construct an instance of the DeleteActionUrlOptions model
				action := "on_user_verified"
				deleteActionUrlOptionsModel := appIdManagementService.NewDeleteActionUrlOptions(action)
				deleteActionUrlOptionsModel.SetAction("on_user_verified")
				deleteActionUrlOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteActionUrlOptionsModel).ToNot(BeNil())
				Expect(deleteActionUrlOptionsModel.Action).To(Equal(core.StringPtr("on_user_verified")))
				Expect(deleteActionUrlOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteApplicationOptions successfully`, func() {
				// Construct an instance of the DeleteApplicationOptions model
				clientID := "testString"
				deleteApplicationOptionsModel := appIdManagementService.NewDeleteApplicationOptions(clientID)
				deleteApplicationOptionsModel.SetClientID("testString")
				deleteApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteApplicationOptionsModel).ToNot(BeNil())
				Expect(deleteApplicationOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(deleteApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCloudDirectoryUserOptions successfully`, func() {
				// Construct an instance of the DeleteCloudDirectoryUserOptions model
				userID := "testString"
				deleteCloudDirectoryUserOptionsModel := appIdManagementService.NewDeleteCloudDirectoryUserOptions(userID)
				deleteCloudDirectoryUserOptionsModel.SetUserID("testString")
				deleteCloudDirectoryUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCloudDirectoryUserOptionsModel).ToNot(BeNil())
				Expect(deleteCloudDirectoryUserOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCloudDirectoryUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRoleOptions successfully`, func() {
				// Construct an instance of the DeleteRoleOptions model
				roleID := "testString"
				deleteRoleOptionsModel := appIdManagementService.NewDeleteRoleOptions(roleID)
				deleteRoleOptionsModel.SetRoleID("testString")
				deleteRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRoleOptionsModel).ToNot(BeNil())
				Expect(deleteRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteTemplateOptions model
				templateName := "USER_VERIFICATION"
				language := "testString"
				deleteTemplateOptionsModel := appIdManagementService.NewDeleteTemplateOptions(templateName, language)
				deleteTemplateOptionsModel.SetTemplateName("USER_VERIFICATION")
				deleteTemplateOptionsModel.SetLanguage("testString")
				deleteTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteTemplateOptionsModel.TemplateName).To(Equal(core.StringPtr("USER_VERIFICATION")))
				Expect(deleteTemplateOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEmailDispatcherParamsCustom successfully`, func() {
				url := "testString"
				var authorization *appidmanagementv4.EmailDispatcherParamsCustomAuthorization = nil
				_, err := appIdManagementService.NewEmailDispatcherParamsCustom(url, authorization)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewEmailDispatcherParamsCustomAuthorization successfully`, func() {
				typeVar := "value"
				model, err := appIdManagementService.NewEmailDispatcherParamsCustomAuthorization(typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailDispatcherParamsSendgrid successfully`, func() {
				apiKey := "testString"
				model, err := appIdManagementService.NewEmailDispatcherParamsSendgrid(apiKey)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailSettingTestOptions successfully`, func() {
				// Construct an instance of the EmailSettingsTestParamsEmailSettingsSendgrid model
				emailSettingsTestParamsEmailSettingsSendgridModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsSendgrid)
				Expect(emailSettingsTestParamsEmailSettingsSendgridModel).ToNot(BeNil())
				emailSettingsTestParamsEmailSettingsSendgridModel.ApiKey = core.StringPtr("testString")
				Expect(emailSettingsTestParamsEmailSettingsSendgridModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustomAuthorization model
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustomAuthorization)
				Expect(emailSettingsTestParamsEmailSettingsCustomAuthorizationModel).ToNot(BeNil())
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Password = core.StringPtr("testString")
				Expect(emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Type).To(Equal(core.StringPtr("value")))
				Expect(emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(emailSettingsTestParamsEmailSettingsCustomAuthorizationModel.Password).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EmailSettingsTestParamsEmailSettingsCustom model
				emailSettingsTestParamsEmailSettingsCustomModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustom)
				Expect(emailSettingsTestParamsEmailSettingsCustomModel).ToNot(BeNil())
				emailSettingsTestParamsEmailSettingsCustomModel.URL = core.StringPtr("testString")
				emailSettingsTestParamsEmailSettingsCustomModel.Authorization = emailSettingsTestParamsEmailSettingsCustomAuthorizationModel
				Expect(emailSettingsTestParamsEmailSettingsCustomModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(emailSettingsTestParamsEmailSettingsCustomModel.Authorization).To(Equal(emailSettingsTestParamsEmailSettingsCustomAuthorizationModel))

				// Construct an instance of the EmailSettingsTestParamsEmailSettings model
				emailSettingsTestParamsEmailSettingsModel := new(appidmanagementv4.EmailSettingsTestParamsEmailSettings)
				Expect(emailSettingsTestParamsEmailSettingsModel).ToNot(BeNil())
				emailSettingsTestParamsEmailSettingsModel.Provider = core.StringPtr("sendgrid")
				emailSettingsTestParamsEmailSettingsModel.Sendgrid = emailSettingsTestParamsEmailSettingsSendgridModel
				emailSettingsTestParamsEmailSettingsModel.Custom = emailSettingsTestParamsEmailSettingsCustomModel
				Expect(emailSettingsTestParamsEmailSettingsModel.Provider).To(Equal(core.StringPtr("sendgrid")))
				Expect(emailSettingsTestParamsEmailSettingsModel.Sendgrid).To(Equal(emailSettingsTestParamsEmailSettingsSendgridModel))
				Expect(emailSettingsTestParamsEmailSettingsModel.Custom).To(Equal(emailSettingsTestParamsEmailSettingsCustomModel))

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsFrom model
				emailSettingsTestParamsSenderDetailsFromModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsFrom)
				Expect(emailSettingsTestParamsSenderDetailsFromModel).ToNot(BeNil())
				emailSettingsTestParamsSenderDetailsFromModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsFromModel.Name = core.StringPtr("testString")
				Expect(emailSettingsTestParamsSenderDetailsFromModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(emailSettingsTestParamsSenderDetailsFromModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EmailSettingsTestParamsSenderDetailsReplyTo model
				emailSettingsTestParamsSenderDetailsReplyToModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetailsReplyTo)
				Expect(emailSettingsTestParamsSenderDetailsReplyToModel).ToNot(BeNil())
				emailSettingsTestParamsSenderDetailsReplyToModel.Email = core.StringPtr("testString")
				emailSettingsTestParamsSenderDetailsReplyToModel.Name = core.StringPtr("testString")
				Expect(emailSettingsTestParamsSenderDetailsReplyToModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(emailSettingsTestParamsSenderDetailsReplyToModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EmailSettingsTestParamsSenderDetails model
				emailSettingsTestParamsSenderDetailsModel := new(appidmanagementv4.EmailSettingsTestParamsSenderDetails)
				Expect(emailSettingsTestParamsSenderDetailsModel).ToNot(BeNil())
				emailSettingsTestParamsSenderDetailsModel.From = emailSettingsTestParamsSenderDetailsFromModel
				emailSettingsTestParamsSenderDetailsModel.ReplyTo = emailSettingsTestParamsSenderDetailsReplyToModel
				Expect(emailSettingsTestParamsSenderDetailsModel.From).To(Equal(emailSettingsTestParamsSenderDetailsFromModel))
				Expect(emailSettingsTestParamsSenderDetailsModel.ReplyTo).To(Equal(emailSettingsTestParamsSenderDetailsReplyToModel))

				// Construct an instance of the EmailSettingTestOptions model
				emailSettingTestOptionsEmailTo := "testString"
				var emailSettingTestOptionsEmailSettings *appidmanagementv4.EmailSettingsTestParamsEmailSettings = nil
				var emailSettingTestOptionsSenderDetails *appidmanagementv4.EmailSettingsTestParamsSenderDetails = nil
				emailSettingTestOptionsModel := appIdManagementService.NewEmailSettingTestOptions(emailSettingTestOptionsEmailTo, emailSettingTestOptionsEmailSettings, emailSettingTestOptionsSenderDetails)
				emailSettingTestOptionsModel.SetEmailTo("testString")
				emailSettingTestOptionsModel.SetEmailSettings(emailSettingsTestParamsEmailSettingsModel)
				emailSettingTestOptionsModel.SetSenderDetails(emailSettingsTestParamsSenderDetailsModel)
				emailSettingTestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(emailSettingTestOptionsModel).ToNot(BeNil())
				Expect(emailSettingTestOptionsModel.EmailTo).To(Equal(core.StringPtr("testString")))
				Expect(emailSettingTestOptionsModel.EmailSettings).To(Equal(emailSettingsTestParamsEmailSettingsModel))
				Expect(emailSettingTestOptionsModel.SenderDetails).To(Equal(emailSettingsTestParamsSenderDetailsModel))
				Expect(emailSettingTestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEmailSettingsTestParamsEmailSettings successfully`, func() {
				provider := "sendgrid"
				model, err := appIdManagementService.NewEmailSettingsTestParamsEmailSettings(provider)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailSettingsTestParamsEmailSettingsCustom successfully`, func() {
				url := "testString"
				var authorization *appidmanagementv4.EmailSettingsTestParamsEmailSettingsCustomAuthorization = nil
				_, err := appIdManagementService.NewEmailSettingsTestParamsEmailSettingsCustom(url, authorization)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewEmailSettingsTestParamsEmailSettingsCustomAuthorization successfully`, func() {
				typeVar := "value"
				model, err := appIdManagementService.NewEmailSettingsTestParamsEmailSettingsCustomAuthorization(typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailSettingsTestParamsEmailSettingsSendgrid successfully`, func() {
				apiKey := "testString"
				model, err := appIdManagementService.NewEmailSettingsTestParamsEmailSettingsSendgrid(apiKey)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailSettingsTestParamsSenderDetails successfully`, func() {
				var from *appidmanagementv4.EmailSettingsTestParamsSenderDetailsFrom = nil
				_, err := appIdManagementService.NewEmailSettingsTestParamsSenderDetails(from)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewEmailSettingsTestParamsSenderDetailsFrom successfully`, func() {
				email := "testString"
				model, err := appIdManagementService.NewEmailSettingsTestParamsSenderDetailsFrom(email)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailSettingsTestParamsSenderDetailsReplyTo successfully`, func() {
				email := "testString"
				model, err := appIdManagementService.NewEmailSettingsTestParamsSenderDetailsReplyTo(email)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExportUserProfileUsersItem successfully`, func() {
				id := "testString"
				identities := []appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{}
				attributes := map[string]interface{}{"anyKey": "anyValue"}
				model, err := appIdManagementService.NewExportUserProfileUsersItem(id, identities, attributes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExportUserUsersItem successfully`, func() {
				scimUser := map[string]interface{}{"anyKey": "anyValue"}
				passwordHash := "testString"
				passwordHashAlg := "testString"
				var profile *appidmanagementv4.ExportUserUsersItemProfile = nil
				roles := []string{"testString"}
				_, err := appIdManagementService.NewExportUserUsersItem(scimUser, passwordHash, passwordHashAlg, profile, roles)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewExportUserUsersItemProfile successfully`, func() {
				attributes := map[string]interface{}{"anyKey": "anyValue"}
				model, err := appIdManagementService.NewExportUserUsersItemProfile(attributes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFacebookGoogleConfigParamsConfig successfully`, func() {
				idpID := "appID"
				secret := "appsecret"
				model, err := appIdManagementService.NewFacebookGoogleConfigParamsConfig(idpID, secret)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewForgotPasswordResultOptions successfully`, func() {
				// Construct an instance of the ForgotPasswordResultOptions model
				forgotPasswordResultOptionsContext := "testString"
				forgotPasswordResultOptionsModel := appIdManagementService.NewForgotPasswordResultOptions(forgotPasswordResultOptionsContext)
				forgotPasswordResultOptionsModel.SetContext("testString")
				forgotPasswordResultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(forgotPasswordResultOptionsModel).ToNot(BeNil())
				Expect(forgotPasswordResultOptionsModel.Context).To(Equal(core.StringPtr("testString")))
				Expect(forgotPasswordResultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationOptions successfully`, func() {
				// Construct an instance of the GetApplicationOptions model
				clientID := "testString"
				getApplicationOptionsModel := appIdManagementService.NewGetApplicationOptions(clientID)
				getApplicationOptionsModel.SetClientID("testString")
				getApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationOptionsModel).ToNot(BeNil())
				Expect(getApplicationOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationRolesOptions successfully`, func() {
				// Construct an instance of the GetApplicationRolesOptions model
				clientID := "testString"
				getApplicationRolesOptionsModel := appIdManagementService.NewGetApplicationRolesOptions(clientID)
				getApplicationRolesOptionsModel.SetClientID("testString")
				getApplicationRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationRolesOptionsModel).ToNot(BeNil())
				Expect(getApplicationRolesOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationScopesOptions successfully`, func() {
				// Construct an instance of the GetApplicationScopesOptions model
				clientID := "testString"
				getApplicationScopesOptionsModel := appIdManagementService.NewGetApplicationScopesOptions(clientID)
				getApplicationScopesOptionsModel.SetClientID("testString")
				getApplicationScopesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationScopesOptionsModel).ToNot(BeNil())
				Expect(getApplicationScopesOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationScopesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAuditStatusOptions successfully`, func() {
				// Construct an instance of the GetAuditStatusOptions model
				getAuditStatusOptionsModel := appIdManagementService.NewGetAuditStatusOptions()
				getAuditStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAuditStatusOptionsModel).ToNot(BeNil())
				Expect(getAuditStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChannelOptions successfully`, func() {
				// Construct an instance of the GetChannelOptions model
				channel := "email"
				getChannelOptionsModel := appIdManagementService.NewGetChannelOptions(channel)
				getChannelOptionsModel.SetChannel("email")
				getChannelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChannelOptionsModel).ToNot(BeNil())
				Expect(getChannelOptionsModel.Channel).To(Equal(core.StringPtr("email")))
				Expect(getChannelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectoryActionUrlOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectoryActionUrlOptions model
				action := "on_user_verified"
				getCloudDirectoryActionUrlOptionsModel := appIdManagementService.NewGetCloudDirectoryActionUrlOptions(action)
				getCloudDirectoryActionUrlOptionsModel.SetAction("on_user_verified")
				getCloudDirectoryActionUrlOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectoryActionUrlOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectoryActionUrlOptionsModel.Action).To(Equal(core.StringPtr("on_user_verified")))
				Expect(getCloudDirectoryActionUrlOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectoryAdvancedPasswordManagementOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectoryAdvancedPasswordManagementOptions model
				getCloudDirectoryAdvancedPasswordManagementOptionsModel := appIdManagementService.NewGetCloudDirectoryAdvancedPasswordManagementOptions()
				getCloudDirectoryAdvancedPasswordManagementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectoryAdvancedPasswordManagementOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectoryEmailDispatcherOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectoryEmailDispatcherOptions model
				getCloudDirectoryEmailDispatcherOptionsModel := appIdManagementService.NewGetCloudDirectoryEmailDispatcherOptions()
				getCloudDirectoryEmailDispatcherOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectoryEmailDispatcherOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectoryEmailDispatcherOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectoryIdpOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectoryIdpOptions model
				getCloudDirectoryIdpOptionsModel := appIdManagementService.NewGetCloudDirectoryIdpOptions()
				getCloudDirectoryIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectoryIdpOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectoryIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectoryPasswordRegexOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectoryPasswordRegexOptions model
				getCloudDirectoryPasswordRegexOptionsModel := appIdManagementService.NewGetCloudDirectoryPasswordRegexOptions()
				getCloudDirectoryPasswordRegexOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectoryPasswordRegexOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectoryPasswordRegexOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectorySenderDetailsOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectorySenderDetailsOptions model
				getCloudDirectorySenderDetailsOptionsModel := appIdManagementService.NewGetCloudDirectorySenderDetailsOptions()
				getCloudDirectorySenderDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectorySenderDetailsOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectorySenderDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCloudDirectoryUserOptions successfully`, func() {
				// Construct an instance of the GetCloudDirectoryUserOptions model
				userID := "testString"
				getCloudDirectoryUserOptionsModel := appIdManagementService.NewGetCloudDirectoryUserOptions(userID)
				getCloudDirectoryUserOptionsModel.SetUserID("testString")
				getCloudDirectoryUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCloudDirectoryUserOptionsModel).ToNot(BeNil())
				Expect(getCloudDirectoryUserOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getCloudDirectoryUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomIdpOptions successfully`, func() {
				// Construct an instance of the GetCustomIdpOptions model
				getCustomIdpOptionsModel := appIdManagementService.NewGetCustomIdpOptions()
				getCustomIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomIdpOptionsModel).ToNot(BeNil())
				Expect(getCustomIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetExtensionConfigOptions successfully`, func() {
				// Construct an instance of the GetExtensionConfigOptions model
				name := "premfa"
				getExtensionConfigOptionsModel := appIdManagementService.NewGetExtensionConfigOptions(name)
				getExtensionConfigOptionsModel.SetName("premfa")
				getExtensionConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getExtensionConfigOptionsModel).ToNot(BeNil())
				Expect(getExtensionConfigOptionsModel.Name).To(Equal(core.StringPtr("premfa")))
				Expect(getExtensionConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFacebookIdpOptions successfully`, func() {
				// Construct an instance of the GetFacebookIdpOptions model
				getFacebookIdpOptionsModel := appIdManagementService.NewGetFacebookIdpOptions()
				getFacebookIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFacebookIdpOptionsModel).ToNot(BeNil())
				Expect(getFacebookIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGoogleIdpOptions successfully`, func() {
				// Construct an instance of the GetGoogleIdpOptions model
				getGoogleIdpOptionsModel := appIdManagementService.NewGetGoogleIdpOptions()
				getGoogleIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGoogleIdpOptionsModel).ToNot(BeNil())
				Expect(getGoogleIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLocalizationOptions successfully`, func() {
				// Construct an instance of the GetLocalizationOptions model
				getLocalizationOptionsModel := appIdManagementService.NewGetLocalizationOptions()
				getLocalizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLocalizationOptionsModel).ToNot(BeNil())
				Expect(getLocalizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMFAConfigOptions successfully`, func() {
				// Construct an instance of the GetMFAConfigOptions model
				getMfaConfigOptionsModel := appIdManagementService.NewGetMFAConfigOptions()
				getMfaConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMfaConfigOptionsModel).ToNot(BeNil())
				Expect(getMfaConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMediaOptions successfully`, func() {
				// Construct an instance of the GetMediaOptions model
				getMediaOptionsModel := appIdManagementService.NewGetMediaOptions()
				getMediaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMediaOptionsModel).ToNot(BeNil())
				Expect(getMediaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRateLimitConfigOptions successfully`, func() {
				// Construct an instance of the GetRateLimitConfigOptions model
				getRateLimitConfigOptionsModel := appIdManagementService.NewGetRateLimitConfigOptions()
				getRateLimitConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRateLimitConfigOptionsModel).ToNot(BeNil())
				Expect(getRateLimitConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRedirectUrisOptions successfully`, func() {
				// Construct an instance of the GetRedirectUrisOptions model
				getRedirectUrisOptionsModel := appIdManagementService.NewGetRedirectUrisOptions()
				getRedirectUrisOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRedirectUrisOptionsModel).ToNot(BeNil())
				Expect(getRedirectUrisOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRoleOptions successfully`, func() {
				// Construct an instance of the GetRoleOptions model
				roleID := "testString"
				getRoleOptionsModel := appIdManagementService.NewGetRoleOptions(roleID)
				getRoleOptionsModel.SetRoleID("testString")
				getRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRoleOptionsModel).ToNot(BeNil())
				Expect(getRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(getRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSSOConfigOptions successfully`, func() {
				// Construct an instance of the GetSSOConfigOptions model
				getSsoConfigOptionsModel := appIdManagementService.NewGetSSOConfigOptions()
				getSsoConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSsoConfigOptionsModel).ToNot(BeNil())
				Expect(getSsoConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSamlIdpOptions successfully`, func() {
				// Construct an instance of the GetSamlIdpOptions model
				getSamlIdpOptionsModel := appIdManagementService.NewGetSamlIdpOptions()
				getSamlIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSamlIdpOptionsModel).ToNot(BeNil())
				Expect(getSamlIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSamlMetadataOptions successfully`, func() {
				// Construct an instance of the GetSamlMetadataOptions model
				getSamlMetadataOptionsModel := appIdManagementService.NewGetSamlMetadataOptions()
				getSamlMetadataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSamlMetadataOptionsModel).ToNot(BeNil())
				Expect(getSamlMetadataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTemplateOptions successfully`, func() {
				// Construct an instance of the GetTemplateOptions model
				templateName := "USER_VERIFICATION"
				language := "testString"
				getTemplateOptionsModel := appIdManagementService.NewGetTemplateOptions(templateName, language)
				getTemplateOptionsModel.SetTemplateName("USER_VERIFICATION")
				getTemplateOptionsModel.SetLanguage("testString")
				getTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTemplateOptionsModel).ToNot(BeNil())
				Expect(getTemplateOptionsModel.TemplateName).To(Equal(core.StringPtr("USER_VERIFICATION")))
				Expect(getTemplateOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetThemeColorOptions successfully`, func() {
				// Construct an instance of the GetThemeColorOptions model
				getThemeColorOptionsModel := appIdManagementService.NewGetThemeColorOptions()
				getThemeColorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getThemeColorOptionsModel).ToNot(BeNil())
				Expect(getThemeColorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetThemeTextOptions successfully`, func() {
				// Construct an instance of the GetThemeTextOptions model
				getThemeTextOptionsModel := appIdManagementService.NewGetThemeTextOptions()
				getThemeTextOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getThemeTextOptionsModel).ToNot(BeNil())
				Expect(getThemeTextOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTokensConfigOptions successfully`, func() {
				// Construct an instance of the GetTokensConfigOptions model
				getTokensConfigOptionsModel := appIdManagementService.NewGetTokensConfigOptions()
				getTokensConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTokensConfigOptionsModel).ToNot(BeNil())
				Expect(getTokensConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserProfilesConfigOptions successfully`, func() {
				// Construct an instance of the GetUserProfilesConfigOptions model
				getUserProfilesConfigOptionsModel := appIdManagementService.NewGetUserProfilesConfigOptions()
				getUserProfilesConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserProfilesConfigOptionsModel).ToNot(BeNil())
				Expect(getUserProfilesConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserRolesOptions successfully`, func() {
				// Construct an instance of the GetUserRolesOptions model
				id := "testString"
				getUserRolesOptionsModel := appIdManagementService.NewGetUserRolesOptions(id)
				getUserRolesOptionsModel.SetID("testString")
				getUserRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserRolesOptionsModel).ToNot(BeNil())
				Expect(getUserRolesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getUserRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListApplicationsOptions successfully`, func() {
				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := appIdManagementService.NewListApplicationsOptions()
				listApplicationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listApplicationsOptionsModel).ToNot(BeNil())
				Expect(listApplicationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListChannelsOptions successfully`, func() {
				// Construct an instance of the ListChannelsOptions model
				listChannelsOptionsModel := appIdManagementService.NewListChannelsOptions()
				listChannelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listChannelsOptionsModel).ToNot(BeNil())
				Expect(listChannelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCloudDirectoryUsersOptions successfully`, func() {
				// Construct an instance of the ListCloudDirectoryUsersOptions model
				listCloudDirectoryUsersOptionsModel := appIdManagementService.NewListCloudDirectoryUsersOptions()
				listCloudDirectoryUsersOptionsModel.SetStartIndex(int64(38))
				listCloudDirectoryUsersOptionsModel.SetCount(int64(0))
				listCloudDirectoryUsersOptionsModel.SetQuery("testString")
				listCloudDirectoryUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCloudDirectoryUsersOptionsModel).ToNot(BeNil())
				Expect(listCloudDirectoryUsersOptionsModel.StartIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listCloudDirectoryUsersOptionsModel.Count).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listCloudDirectoryUsersOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(listCloudDirectoryUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRolesOptions successfully`, func() {
				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := appIdManagementService.NewListRolesOptions()
				listRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRolesOptionsModel).ToNot(BeNil())
				Expect(listRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostEmailDispatcherTestOptions successfully`, func() {
				// Construct an instance of the PostEmailDispatcherTestOptions model
				postEmailDispatcherTestOptionsEmail := "testString"
				postEmailDispatcherTestOptionsModel := appIdManagementService.NewPostEmailDispatcherTestOptions(postEmailDispatcherTestOptionsEmail)
				postEmailDispatcherTestOptionsModel.SetEmail("testString")
				postEmailDispatcherTestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postEmailDispatcherTestOptionsModel).ToNot(BeNil())
				Expect(postEmailDispatcherTestOptionsModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(postEmailDispatcherTestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostExtensionsTestOptions successfully`, func() {
				// Construct an instance of the PostExtensionsTestOptions model
				name := "premfa"
				postExtensionsTestOptionsModel := appIdManagementService.NewPostExtensionsTestOptions(name)
				postExtensionsTestOptionsModel.SetName("premfa")
				postExtensionsTestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postExtensionsTestOptionsModel).ToNot(BeNil())
				Expect(postExtensionsTestOptionsModel.Name).To(Equal(core.StringPtr("premfa")))
				Expect(postExtensionsTestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostMediaOptions successfully`, func() {
				// Construct an instance of the PostMediaOptions model
				mediaType := "logo"
				file := CreateMockReader("This is a mock file.")
				postMediaOptionsModel := appIdManagementService.NewPostMediaOptions(mediaType, file)
				postMediaOptionsModel.SetMediaType("logo")
				postMediaOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				postMediaOptionsModel.SetFileContentType("testString")
				postMediaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postMediaOptionsModel).ToNot(BeNil())
				Expect(postMediaOptionsModel.MediaType).To(Equal(core.StringPtr("logo")))
				Expect(postMediaOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(postMediaOptionsModel.FileContentType).To(Equal(core.StringPtr("testString")))
				Expect(postMediaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostSmsDispatcherTestOptions successfully`, func() {
				// Construct an instance of the PostSmsDispatcherTestOptions model
				postSmsDispatcherTestOptionsPhoneNumber := "+1-999-999-9999"
				postSmsDispatcherTestOptionsModel := appIdManagementService.NewPostSmsDispatcherTestOptions(postSmsDispatcherTestOptionsPhoneNumber)
				postSmsDispatcherTestOptionsModel.SetPhoneNumber("+1-999-999-9999")
				postSmsDispatcherTestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postSmsDispatcherTestOptionsModel).ToNot(BeNil())
				Expect(postSmsDispatcherTestOptionsModel.PhoneNumber).To(Equal(core.StringPtr("+1-999-999-9999")))
				Expect(postSmsDispatcherTestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostThemeColorOptions successfully`, func() {
				// Construct an instance of the PostThemeColorOptions model
				postThemeColorOptionsModel := appIdManagementService.NewPostThemeColorOptions()
				postThemeColorOptionsModel.SetHeaderColor("#EEF2F5")
				postThemeColorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postThemeColorOptionsModel).ToNot(BeNil())
				Expect(postThemeColorOptionsModel.HeaderColor).To(Equal(core.StringPtr("#EEF2F5")))
				Expect(postThemeColorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostThemeTextOptions successfully`, func() {
				// Construct an instance of the PostThemeTextOptions model
				postThemeTextOptionsModel := appIdManagementService.NewPostThemeTextOptions()
				postThemeTextOptionsModel.SetTabTitle("Login")
				postThemeTextOptionsModel.SetFootnote("Powered by App ID")
				postThemeTextOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postThemeTextOptionsModel).ToNot(BeNil())
				Expect(postThemeTextOptionsModel.TabTitle).To(Equal(core.StringPtr("Login")))
				Expect(postThemeTextOptionsModel.Footnote).To(Equal(core.StringPtr("Powered by App ID")))
				Expect(postThemeTextOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutApplicationsRolesOptions successfully`, func() {
				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				Expect(updateUserRolesParamsRolesModel).ToNot(BeNil())
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}
				Expect(updateUserRolesParamsRolesModel.Ids).To(Equal([]string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}))

				// Construct an instance of the PutApplicationsRolesOptions model
				clientID := "testString"
				var putApplicationsRolesOptionsRoles *appidmanagementv4.UpdateUserRolesParamsRoles = nil
				putApplicationsRolesOptionsModel := appIdManagementService.NewPutApplicationsRolesOptions(clientID, putApplicationsRolesOptionsRoles)
				putApplicationsRolesOptionsModel.SetClientID("testString")
				putApplicationsRolesOptionsModel.SetRoles(updateUserRolesParamsRolesModel)
				putApplicationsRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putApplicationsRolesOptionsModel).ToNot(BeNil())
				Expect(putApplicationsRolesOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(putApplicationsRolesOptionsModel.Roles).To(Equal(updateUserRolesParamsRolesModel))
				Expect(putApplicationsRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutApplicationsScopesOptions successfully`, func() {
				// Construct an instance of the PutApplicationsScopesOptions model
				clientID := "testString"
				putApplicationsScopesOptionsModel := appIdManagementService.NewPutApplicationsScopesOptions(clientID)
				putApplicationsScopesOptionsModel.SetClientID("testString")
				putApplicationsScopesOptionsModel.SetScopes([]string{"cartoons", "horror", "animated"})
				putApplicationsScopesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putApplicationsScopesOptionsModel).ToNot(BeNil())
				Expect(putApplicationsScopesOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(putApplicationsScopesOptionsModel.Scopes).To(Equal([]string{"cartoons", "horror", "animated"}))
				Expect(putApplicationsScopesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutTokensConfigOptions successfully`, func() {
				// Construct an instance of the TokenClaimMapping model
				tokenClaimMappingModel := new(appidmanagementv4.TokenClaimMapping)
				Expect(tokenClaimMappingModel).ToNot(BeNil())
				tokenClaimMappingModel.Source = core.StringPtr("saml")
				tokenClaimMappingModel.SourceClaim = core.StringPtr("testString")
				tokenClaimMappingModel.DestinationClaim = core.StringPtr("testString")
				Expect(tokenClaimMappingModel.Source).To(Equal(core.StringPtr("saml")))
				Expect(tokenClaimMappingModel.SourceClaim).To(Equal(core.StringPtr("testString")))
				Expect(tokenClaimMappingModel.DestinationClaim).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TokenConfigParams model
				tokenConfigParamsModel := new(appidmanagementv4.TokenConfigParams)
				Expect(tokenConfigParamsModel).ToNot(BeNil())
				tokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))
				Expect(tokenConfigParamsModel.ExpiresIn).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the RefreshTokenConfigParams model
				refreshTokenConfigParamsModel := new(appidmanagementv4.RefreshTokenConfigParams)
				Expect(refreshTokenConfigParamsModel).ToNot(BeNil())
				refreshTokenConfigParamsModel.ExpiresIn = core.Float64Ptr(float64(72.5))
				refreshTokenConfigParamsModel.Enabled = core.BoolPtr(true)
				Expect(refreshTokenConfigParamsModel.ExpiresIn).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(refreshTokenConfigParamsModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PutTokensConfigOptions model
				putTokensConfigOptionsModel := appIdManagementService.NewPutTokensConfigOptions()
				putTokensConfigOptionsModel.SetIdTokenClaims([]appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel})
				putTokensConfigOptionsModel.SetAccessTokenClaims([]appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel})
				putTokensConfigOptionsModel.SetAccess([]appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel})
				putTokensConfigOptionsModel.SetRefresh([]appidmanagementv4.RefreshTokenConfigParams{*refreshTokenConfigParamsModel})
				putTokensConfigOptionsModel.SetAnonymousAccess([]appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel})
				putTokensConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putTokensConfigOptionsModel).ToNot(BeNil())
				Expect(putTokensConfigOptionsModel.IdTokenClaims).To(Equal([]appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}))
				Expect(putTokensConfigOptionsModel.AccessTokenClaims).To(Equal([]appidmanagementv4.TokenClaimMapping{*tokenClaimMappingModel}))
				Expect(putTokensConfigOptionsModel.Access).To(Equal([]appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}))
				Expect(putTokensConfigOptionsModel.Refresh).To(Equal([]appidmanagementv4.RefreshTokenConfigParams{*refreshTokenConfigParamsModel}))
				Expect(putTokensConfigOptionsModel.AnonymousAccess).To(Equal([]appidmanagementv4.TokenConfigParams{*tokenConfigParamsModel}))
				Expect(putTokensConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRegisterApplicationOptions successfully`, func() {
				// Construct an instance of the RegisterApplicationOptions model
				registerApplicationOptionsName := "testString"
				registerApplicationOptionsModel := appIdManagementService.NewRegisterApplicationOptions(registerApplicationOptionsName)
				registerApplicationOptionsModel.SetName("testString")
				registerApplicationOptionsModel.SetType("testString")
				registerApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(registerApplicationOptionsModel).ToNot(BeNil())
				Expect(registerApplicationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(registerApplicationOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(registerApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResendNotificationOptions successfully`, func() {
				// Construct an instance of the ResendNotificationOptions model
				templateName := "USER_VERIFICATION"
				resendNotificationOptionsUUID := "testString"
				resendNotificationOptionsModel := appIdManagementService.NewResendNotificationOptions(templateName, resendNotificationOptionsUUID)
				resendNotificationOptionsModel.SetTemplateName("USER_VERIFICATION")
				resendNotificationOptionsModel.SetUUID("testString")
				resendNotificationOptionsModel.SetLanguage("testString")
				resendNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resendNotificationOptionsModel).ToNot(BeNil())
				Expect(resendNotificationOptionsModel.TemplateName).To(Equal(core.StringPtr("USER_VERIFICATION")))
				Expect(resendNotificationOptionsModel.UUID).To(Equal(core.StringPtr("testString")))
				Expect(resendNotificationOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(resendNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetAuditStatusOptions successfully`, func() {
				// Construct an instance of the SetAuditStatusOptions model
				setAuditStatusOptionsIsActive := true
				setAuditStatusOptionsModel := appIdManagementService.NewSetAuditStatusOptions(setAuditStatusOptionsIsActive)
				setAuditStatusOptionsModel.SetIsActive(true)
				setAuditStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setAuditStatusOptionsModel).ToNot(BeNil())
				Expect(setAuditStatusOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(setAuditStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCloudDirectoryActionOptions successfully`, func() {
				// Construct an instance of the SetCloudDirectoryActionOptions model
				action := "on_user_verified"
				setCloudDirectoryActionOptionsActionURL := "testString"
				setCloudDirectoryActionOptionsModel := appIdManagementService.NewSetCloudDirectoryActionOptions(action, setCloudDirectoryActionOptionsActionURL)
				setCloudDirectoryActionOptionsModel.SetAction("on_user_verified")
				setCloudDirectoryActionOptionsModel.SetActionURL("testString")
				setCloudDirectoryActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCloudDirectoryActionOptionsModel).ToNot(BeNil())
				Expect(setCloudDirectoryActionOptionsModel.Action).To(Equal(core.StringPtr("on_user_verified")))
				Expect(setCloudDirectoryActionOptionsModel.ActionURL).To(Equal(core.StringPtr("testString")))
				Expect(setCloudDirectoryActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCloudDirectoryAdvancedPasswordManagementOptions successfully`, func() {
				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuseConfig model
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
				Expect(apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel.MaxPasswordReuse = core.Float64Ptr(float64(1))
				Expect(apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel.MaxPasswordReuse).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordReuse model
				apmSchemaAdvancedPasswordManagementPasswordReuseModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordReuse)
				Expect(apmSchemaAdvancedPasswordManagementPasswordReuseModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordReuseModel.Config = apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel
				Expect(apmSchemaAdvancedPasswordManagementPasswordReuseModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(apmSchemaAdvancedPasswordManagementPasswordReuseModel.Config).To(Equal(apmSchemaAdvancedPasswordManagementPasswordReuseConfigModel))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername model
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
				Expect(apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel.Enabled = core.BoolPtr(true)
				Expect(apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig model
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
				Expect(apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel.DaysToExpire = core.Float64Ptr(float64(1))
				Expect(apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel.DaysToExpire).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementPasswordExpiration model
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementPasswordExpiration)
				Expect(apmSchemaAdvancedPasswordManagementPasswordExpirationModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Config = apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel
				Expect(apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(apmSchemaAdvancedPasswordManagementPasswordExpirationModel.Config).To(Equal(apmSchemaAdvancedPasswordManagementPasswordExpirationConfigModel))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig model
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
				Expect(apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.LockOutTimeSec = core.Float64Ptr(float64(60))
				apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.NumOfAttempts = core.Float64Ptr(float64(1))
				Expect(apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.LockOutTimeSec).To(Equal(core.Float64Ptr(float64(60))))
				Expect(apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel.NumOfAttempts).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementLockOutPolicy model
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementLockOutPolicy)
				Expect(apmSchemaAdvancedPasswordManagementLockOutPolicyModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Config = apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel
				Expect(apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(apmSchemaAdvancedPasswordManagementLockOutPolicyModel.Config).To(Equal(apmSchemaAdvancedPasswordManagementLockOutPolicyConfigModel))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
				Expect(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel.MinHoursToChangePassword = core.Float64Ptr(float64(0))
				Expect(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel.MinHoursToChangePassword).To(Equal(core.Float64Ptr(float64(0))))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval model
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
				Expect(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Config = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel
				Expect(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel.Config).To(Equal(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfigModel))

				// Construct an instance of the ApmSchemaAdvancedPasswordManagement model
				apmSchemaAdvancedPasswordManagementModel := new(appidmanagementv4.ApmSchemaAdvancedPasswordManagement)
				Expect(apmSchemaAdvancedPasswordManagementModel).ToNot(BeNil())
				apmSchemaAdvancedPasswordManagementModel.Enabled = core.BoolPtr(true)
				apmSchemaAdvancedPasswordManagementModel.PasswordReuse = apmSchemaAdvancedPasswordManagementPasswordReuseModel
				apmSchemaAdvancedPasswordManagementModel.PreventPasswordWithUsername = apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel
				apmSchemaAdvancedPasswordManagementModel.PasswordExpiration = apmSchemaAdvancedPasswordManagementPasswordExpirationModel
				apmSchemaAdvancedPasswordManagementModel.LockOutPolicy = apmSchemaAdvancedPasswordManagementLockOutPolicyModel
				apmSchemaAdvancedPasswordManagementModel.MinPasswordChangeInterval = apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel
				Expect(apmSchemaAdvancedPasswordManagementModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(apmSchemaAdvancedPasswordManagementModel.PasswordReuse).To(Equal(apmSchemaAdvancedPasswordManagementPasswordReuseModel))
				Expect(apmSchemaAdvancedPasswordManagementModel.PreventPasswordWithUsername).To(Equal(apmSchemaAdvancedPasswordManagementPreventPasswordWithUsernameModel))
				Expect(apmSchemaAdvancedPasswordManagementModel.PasswordExpiration).To(Equal(apmSchemaAdvancedPasswordManagementPasswordExpirationModel))
				Expect(apmSchemaAdvancedPasswordManagementModel.LockOutPolicy).To(Equal(apmSchemaAdvancedPasswordManagementLockOutPolicyModel))
				Expect(apmSchemaAdvancedPasswordManagementModel.MinPasswordChangeInterval).To(Equal(apmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalModel))

				// Construct an instance of the SetCloudDirectoryAdvancedPasswordManagementOptions model
				var setCloudDirectoryAdvancedPasswordManagementOptionsAdvancedPasswordManagement *appidmanagementv4.ApmSchemaAdvancedPasswordManagement = nil
				setCloudDirectoryAdvancedPasswordManagementOptionsModel := appIdManagementService.NewSetCloudDirectoryAdvancedPasswordManagementOptions(setCloudDirectoryAdvancedPasswordManagementOptionsAdvancedPasswordManagement)
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.SetAdvancedPasswordManagement(apmSchemaAdvancedPasswordManagementModel)
				setCloudDirectoryAdvancedPasswordManagementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCloudDirectoryAdvancedPasswordManagementOptionsModel).ToNot(BeNil())
				Expect(setCloudDirectoryAdvancedPasswordManagementOptionsModel.AdvancedPasswordManagement).To(Equal(apmSchemaAdvancedPasswordManagementModel))
				Expect(setCloudDirectoryAdvancedPasswordManagementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCloudDirectoryEmailDispatcherOptions successfully`, func() {
				// Construct an instance of the EmailDispatcherParamsSendgrid model
				emailDispatcherParamsSendgridModel := new(appidmanagementv4.EmailDispatcherParamsSendgrid)
				Expect(emailDispatcherParamsSendgridModel).ToNot(BeNil())
				emailDispatcherParamsSendgridModel.ApiKey = core.StringPtr("testString")
				Expect(emailDispatcherParamsSendgridModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EmailDispatcherParamsCustomAuthorization model
				emailDispatcherParamsCustomAuthorizationModel := new(appidmanagementv4.EmailDispatcherParamsCustomAuthorization)
				Expect(emailDispatcherParamsCustomAuthorizationModel).ToNot(BeNil())
				emailDispatcherParamsCustomAuthorizationModel.Type = core.StringPtr("value")
				emailDispatcherParamsCustomAuthorizationModel.Value = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Username = core.StringPtr("testString")
				emailDispatcherParamsCustomAuthorizationModel.Password = core.StringPtr("testString")
				Expect(emailDispatcherParamsCustomAuthorizationModel.Type).To(Equal(core.StringPtr("value")))
				Expect(emailDispatcherParamsCustomAuthorizationModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(emailDispatcherParamsCustomAuthorizationModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(emailDispatcherParamsCustomAuthorizationModel.Password).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EmailDispatcherParamsCustom model
				emailDispatcherParamsCustomModel := new(appidmanagementv4.EmailDispatcherParamsCustom)
				Expect(emailDispatcherParamsCustomModel).ToNot(BeNil())
				emailDispatcherParamsCustomModel.URL = core.StringPtr("testString")
				emailDispatcherParamsCustomModel.Authorization = emailDispatcherParamsCustomAuthorizationModel
				Expect(emailDispatcherParamsCustomModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(emailDispatcherParamsCustomModel.Authorization).To(Equal(emailDispatcherParamsCustomAuthorizationModel))

				// Construct an instance of the SetCloudDirectoryEmailDispatcherOptions model
				setCloudDirectoryEmailDispatcherOptionsProvider := "sendgrid"
				setCloudDirectoryEmailDispatcherOptionsModel := appIdManagementService.NewSetCloudDirectoryEmailDispatcherOptions(setCloudDirectoryEmailDispatcherOptionsProvider)
				setCloudDirectoryEmailDispatcherOptionsModel.SetProvider("sendgrid")
				setCloudDirectoryEmailDispatcherOptionsModel.SetSendgrid(emailDispatcherParamsSendgridModel)
				setCloudDirectoryEmailDispatcherOptionsModel.SetCustom(emailDispatcherParamsCustomModel)
				setCloudDirectoryEmailDispatcherOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCloudDirectoryEmailDispatcherOptionsModel).ToNot(BeNil())
				Expect(setCloudDirectoryEmailDispatcherOptionsModel.Provider).To(Equal(core.StringPtr("sendgrid")))
				Expect(setCloudDirectoryEmailDispatcherOptionsModel.Sendgrid).To(Equal(emailDispatcherParamsSendgridModel))
				Expect(setCloudDirectoryEmailDispatcherOptionsModel.Custom).To(Equal(emailDispatcherParamsCustomModel))
				Expect(setCloudDirectoryEmailDispatcherOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCloudDirectoryIdpOptions successfully`, func() {
				// Construct an instance of the CloudDirectoryConfigParamsInteractionsIdentityConfirmation model
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractionsIdentityConfirmation)
				Expect(cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel).ToNot(BeNil())
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.AccessMode = core.StringPtr("FULL")
				cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.Methods = []string{"email"}
				Expect(cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.AccessMode).To(Equal(core.StringPtr("FULL")))
				Expect(cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel.Methods).To(Equal([]string{"email"}))

				// Construct an instance of the CloudDirectoryConfigParamsInteractions model
				cloudDirectoryConfigParamsInteractionsModel := new(appidmanagementv4.CloudDirectoryConfigParamsInteractions)
				Expect(cloudDirectoryConfigParamsInteractionsModel).ToNot(BeNil())
				cloudDirectoryConfigParamsInteractionsModel.IdentityConfirmation = cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel
				cloudDirectoryConfigParamsInteractionsModel.WelcomeEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordEnabled = core.BoolPtr(false)
				cloudDirectoryConfigParamsInteractionsModel.ResetPasswordNotificationEnable = core.BoolPtr(true)
				Expect(cloudDirectoryConfigParamsInteractionsModel.IdentityConfirmation).To(Equal(cloudDirectoryConfigParamsInteractionsIdentityConfirmationModel))
				Expect(cloudDirectoryConfigParamsInteractionsModel.WelcomeEnabled).To(Equal(core.BoolPtr(false)))
				Expect(cloudDirectoryConfigParamsInteractionsModel.ResetPasswordEnabled).To(Equal(core.BoolPtr(false)))
				Expect(cloudDirectoryConfigParamsInteractionsModel.ResetPasswordNotificationEnable).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CloudDirectoryConfigParams model
				cloudDirectoryConfigParamsModel := new(appidmanagementv4.CloudDirectoryConfigParams)
				Expect(cloudDirectoryConfigParamsModel).ToNot(BeNil())
				cloudDirectoryConfigParamsModel.SelfServiceEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.SignupEnabled = core.BoolPtr(true)
				cloudDirectoryConfigParamsModel.Interactions = cloudDirectoryConfigParamsInteractionsModel
				cloudDirectoryConfigParamsModel.IdentityField = core.StringPtr("email")
				Expect(cloudDirectoryConfigParamsModel.SelfServiceEnabled).To(Equal(core.BoolPtr(true)))
				Expect(cloudDirectoryConfigParamsModel.SignupEnabled).To(Equal(core.BoolPtr(true)))
				Expect(cloudDirectoryConfigParamsModel.Interactions).To(Equal(cloudDirectoryConfigParamsInteractionsModel))
				Expect(cloudDirectoryConfigParamsModel.IdentityField).To(Equal(core.StringPtr("email")))

				// Construct an instance of the SetCloudDirectoryIdpOptions model
				setCloudDirectoryIdpOptionsIsActive := true
				var setCloudDirectoryIdpOptionsConfig *appidmanagementv4.CloudDirectoryConfigParams = nil
				setCloudDirectoryIdpOptionsModel := appIdManagementService.NewSetCloudDirectoryIdpOptions(setCloudDirectoryIdpOptionsIsActive, setCloudDirectoryIdpOptionsConfig)
				setCloudDirectoryIdpOptionsModel.SetIsActive(true)
				setCloudDirectoryIdpOptionsModel.SetConfig(cloudDirectoryConfigParamsModel)
				setCloudDirectoryIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCloudDirectoryIdpOptionsModel).ToNot(BeNil())
				Expect(setCloudDirectoryIdpOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(setCloudDirectoryIdpOptionsModel.Config).To(Equal(cloudDirectoryConfigParamsModel))
				Expect(setCloudDirectoryIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCloudDirectoryPasswordRegexOptions successfully`, func() {
				// Construct an instance of the SetCloudDirectoryPasswordRegexOptions model
				setCloudDirectoryPasswordRegexOptionsModel := appIdManagementService.NewSetCloudDirectoryPasswordRegexOptions()
				setCloudDirectoryPasswordRegexOptionsModel.SetRegex("testString")
				setCloudDirectoryPasswordRegexOptionsModel.SetBase64EncodedRegex("testString")
				setCloudDirectoryPasswordRegexOptionsModel.SetErrorMessage("testString")
				setCloudDirectoryPasswordRegexOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCloudDirectoryPasswordRegexOptionsModel).ToNot(BeNil())
				Expect(setCloudDirectoryPasswordRegexOptionsModel.Regex).To(Equal(core.StringPtr("testString")))
				Expect(setCloudDirectoryPasswordRegexOptionsModel.Base64EncodedRegex).To(Equal(core.StringPtr("testString")))
				Expect(setCloudDirectoryPasswordRegexOptionsModel.ErrorMessage).To(Equal(core.StringPtr("testString")))
				Expect(setCloudDirectoryPasswordRegexOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCloudDirectorySenderDetailsOptions successfully`, func() {
				// Construct an instance of the CloudDirectorySenderDetailsSenderDetailsFrom model
				cloudDirectorySenderDetailsSenderDetailsFromModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsFrom)
				Expect(cloudDirectorySenderDetailsSenderDetailsFromModel).ToNot(BeNil())
				cloudDirectorySenderDetailsSenderDetailsFromModel.Name = core.StringPtr("testString")
				cloudDirectorySenderDetailsSenderDetailsFromModel.Email = core.StringPtr("testString")
				Expect(cloudDirectorySenderDetailsSenderDetailsFromModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(cloudDirectorySenderDetailsSenderDetailsFromModel.Email).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetailsReplyTo model
				cloudDirectorySenderDetailsSenderDetailsReplyToModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetailsReplyTo)
				Expect(cloudDirectorySenderDetailsSenderDetailsReplyToModel).ToNot(BeNil())
				cloudDirectorySenderDetailsSenderDetailsReplyToModel.Name = core.StringPtr("testString")
				cloudDirectorySenderDetailsSenderDetailsReplyToModel.Email = core.StringPtr("testString")
				Expect(cloudDirectorySenderDetailsSenderDetailsReplyToModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(cloudDirectorySenderDetailsSenderDetailsReplyToModel.Email).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CloudDirectorySenderDetailsSenderDetails model
				cloudDirectorySenderDetailsSenderDetailsModel := new(appidmanagementv4.CloudDirectorySenderDetailsSenderDetails)
				Expect(cloudDirectorySenderDetailsSenderDetailsModel).ToNot(BeNil())
				cloudDirectorySenderDetailsSenderDetailsModel.From = cloudDirectorySenderDetailsSenderDetailsFromModel
				cloudDirectorySenderDetailsSenderDetailsModel.ReplyTo = cloudDirectorySenderDetailsSenderDetailsReplyToModel
				cloudDirectorySenderDetailsSenderDetailsModel.LinkExpirationSec = core.Float64Ptr(float64(900))
				Expect(cloudDirectorySenderDetailsSenderDetailsModel.From).To(Equal(cloudDirectorySenderDetailsSenderDetailsFromModel))
				Expect(cloudDirectorySenderDetailsSenderDetailsModel.ReplyTo).To(Equal(cloudDirectorySenderDetailsSenderDetailsReplyToModel))
				Expect(cloudDirectorySenderDetailsSenderDetailsModel.LinkExpirationSec).To(Equal(core.Float64Ptr(float64(900))))

				// Construct an instance of the SetCloudDirectorySenderDetailsOptions model
				var setCloudDirectorySenderDetailsOptionsSenderDetails *appidmanagementv4.CloudDirectorySenderDetailsSenderDetails = nil
				setCloudDirectorySenderDetailsOptionsModel := appIdManagementService.NewSetCloudDirectorySenderDetailsOptions(setCloudDirectorySenderDetailsOptionsSenderDetails)
				setCloudDirectorySenderDetailsOptionsModel.SetSenderDetails(cloudDirectorySenderDetailsSenderDetailsModel)
				setCloudDirectorySenderDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCloudDirectorySenderDetailsOptionsModel).ToNot(BeNil())
				Expect(setCloudDirectorySenderDetailsOptionsModel.SenderDetails).To(Equal(cloudDirectorySenderDetailsSenderDetailsModel))
				Expect(setCloudDirectorySenderDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCustomIdpOptions successfully`, func() {
				// Construct an instance of the CustomIdPConfigParamsConfig model
				customIdPConfigParamsConfigModel := new(appidmanagementv4.CustomIdPConfigParamsConfig)
				Expect(customIdPConfigParamsConfigModel).ToNot(BeNil())
				customIdPConfigParamsConfigModel.PublicKey = core.StringPtr("testString")
				Expect(customIdPConfigParamsConfigModel.PublicKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SetCustomIdpOptions model
				setCustomIdpOptionsIsActive := true
				setCustomIdpOptionsModel := appIdManagementService.NewSetCustomIdpOptions(setCustomIdpOptionsIsActive)
				setCustomIdpOptionsModel.SetIsActive(true)
				setCustomIdpOptionsModel.SetConfig(customIdPConfigParamsConfigModel)
				setCustomIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCustomIdpOptionsModel).ToNot(BeNil())
				Expect(setCustomIdpOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(setCustomIdpOptionsModel.Config).To(Equal(customIdPConfigParamsConfigModel))
				Expect(setCustomIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetFacebookIdpOptions successfully`, func() {
				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				Expect(facebookGoogleConfigParamsConfigModel).ToNot(BeNil())
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")
				Expect(facebookGoogleConfigParamsConfigModel.IdpID).To(Equal(core.StringPtr("appID")))
				Expect(facebookGoogleConfigParamsConfigModel.Secret).To(Equal(core.StringPtr("appsecret")))

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				Expect(facebookGoogleConfigParamsModel).ToNot(BeNil())
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(facebookGoogleConfigParamsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(facebookGoogleConfigParamsModel.Config).To(Equal(facebookGoogleConfigParamsConfigModel))
				Expect(facebookGoogleConfigParamsModel.GetProperties()).ToNot(BeEmpty())
				Expect(facebookGoogleConfigParamsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SetFacebookIdpOptions model
				var idp *appidmanagementv4.FacebookGoogleConfigParams = nil
				setFacebookIdpOptionsModel := appIdManagementService.NewSetFacebookIdpOptions(idp)
				setFacebookIdpOptionsModel.SetIdp(facebookGoogleConfigParamsModel)
				setFacebookIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setFacebookIdpOptionsModel).ToNot(BeNil())
				Expect(setFacebookIdpOptionsModel.Idp).To(Equal(facebookGoogleConfigParamsModel))
				Expect(setFacebookIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetGoogleIdpOptions successfully`, func() {
				// Construct an instance of the FacebookGoogleConfigParamsConfig model
				facebookGoogleConfigParamsConfigModel := new(appidmanagementv4.FacebookGoogleConfigParamsConfig)
				Expect(facebookGoogleConfigParamsConfigModel).ToNot(BeNil())
				facebookGoogleConfigParamsConfigModel.IdpID = core.StringPtr("appID")
				facebookGoogleConfigParamsConfigModel.Secret = core.StringPtr("appsecret")
				Expect(facebookGoogleConfigParamsConfigModel.IdpID).To(Equal(core.StringPtr("appID")))
				Expect(facebookGoogleConfigParamsConfigModel.Secret).To(Equal(core.StringPtr("appsecret")))

				// Construct an instance of the FacebookGoogleConfigParams model
				facebookGoogleConfigParamsModel := new(appidmanagementv4.FacebookGoogleConfigParams)
				Expect(facebookGoogleConfigParamsModel).ToNot(BeNil())
				facebookGoogleConfigParamsModel.IsActive = core.BoolPtr(true)
				facebookGoogleConfigParamsModel.Config = facebookGoogleConfigParamsConfigModel
				facebookGoogleConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(facebookGoogleConfigParamsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(facebookGoogleConfigParamsModel.Config).To(Equal(facebookGoogleConfigParamsConfigModel))
				Expect(facebookGoogleConfigParamsModel.GetProperties()).ToNot(BeEmpty())
				Expect(facebookGoogleConfigParamsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SetGoogleIdpOptions model
				var idp *appidmanagementv4.FacebookGoogleConfigParams = nil
				setGoogleIdpOptionsModel := appIdManagementService.NewSetGoogleIdpOptions(idp)
				setGoogleIdpOptionsModel.SetIdp(facebookGoogleConfigParamsModel)
				setGoogleIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setGoogleIdpOptionsModel).ToNot(BeNil())
				Expect(setGoogleIdpOptionsModel.Idp).To(Equal(facebookGoogleConfigParamsModel))
				Expect(setGoogleIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetSamlIdpOptions successfully`, func() {
				// Construct an instance of the SamlConfigParamsAuthnContext model
				samlConfigParamsAuthnContextModel := new(appidmanagementv4.SamlConfigParamsAuthnContext)
				Expect(samlConfigParamsAuthnContextModel).ToNot(BeNil())
				samlConfigParamsAuthnContextModel.Class = []string{"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"}
				samlConfigParamsAuthnContextModel.Comparison = core.StringPtr("exact")
				Expect(samlConfigParamsAuthnContextModel.Class).To(Equal([]string{"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"}))
				Expect(samlConfigParamsAuthnContextModel.Comparison).To(Equal(core.StringPtr("exact")))

				// Construct an instance of the SamlConfigParams model
				samlConfigParamsModel := new(appidmanagementv4.SamlConfigParams)
				Expect(samlConfigParamsModel).ToNot(BeNil())
				samlConfigParamsModel.EntityID = core.StringPtr("testString")
				samlConfigParamsModel.SignInURL = core.StringPtr("testString")
				samlConfigParamsModel.Certificates = []string{"testString"}
				samlConfigParamsModel.DisplayName = core.StringPtr("testString")
				samlConfigParamsModel.AuthnContext = samlConfigParamsAuthnContextModel
				samlConfigParamsModel.SignRequest = core.BoolPtr(false)
				samlConfigParamsModel.EncryptResponse = core.BoolPtr(false)
				samlConfigParamsModel.IncludeScoping = core.BoolPtr(false)
				samlConfigParamsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(samlConfigParamsModel.EntityID).To(Equal(core.StringPtr("testString")))
				Expect(samlConfigParamsModel.SignInURL).To(Equal(core.StringPtr("testString")))
				Expect(samlConfigParamsModel.Certificates).To(Equal([]string{"testString"}))
				Expect(samlConfigParamsModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(samlConfigParamsModel.AuthnContext).To(Equal(samlConfigParamsAuthnContextModel))
				Expect(samlConfigParamsModel.SignRequest).To(Equal(core.BoolPtr(false)))
				Expect(samlConfigParamsModel.EncryptResponse).To(Equal(core.BoolPtr(false)))
				Expect(samlConfigParamsModel.IncludeScoping).To(Equal(core.BoolPtr(false)))
				Expect(samlConfigParamsModel.GetProperties()).ToNot(BeEmpty())
				Expect(samlConfigParamsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SetSamlIdpOptions model
				setSamlIdpOptionsIsActive := true
				setSamlIdpOptionsModel := appIdManagementService.NewSetSamlIdpOptions(setSamlIdpOptionsIsActive)
				setSamlIdpOptionsModel.SetIsActive(true)
				setSamlIdpOptionsModel.SetConfig(samlConfigParamsModel)
				setSamlIdpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setSamlIdpOptionsModel).ToNot(BeNil())
				Expect(setSamlIdpOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(setSamlIdpOptionsModel.Config).To(Equal(samlConfigParamsModel))
				Expect(setSamlIdpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSsoLogoutFromAllAppsOptions successfully`, func() {
				// Construct an instance of the SsoLogoutFromAllAppsOptions model
				userID := "testString"
				ssoLogoutFromAllAppsOptionsModel := appIdManagementService.NewSsoLogoutFromAllAppsOptions(userID)
				ssoLogoutFromAllAppsOptionsModel.SetUserID("testString")
				ssoLogoutFromAllAppsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(ssoLogoutFromAllAppsOptionsModel).ToNot(BeNil())
				Expect(ssoLogoutFromAllAppsOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(ssoLogoutFromAllAppsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewStartForgotPasswordOptions successfully`, func() {
				// Construct an instance of the StartForgotPasswordOptions model
				startForgotPasswordOptionsUser := "testString"
				startForgotPasswordOptionsModel := appIdManagementService.NewStartForgotPasswordOptions(startForgotPasswordOptionsUser)
				startForgotPasswordOptionsModel.SetUser("testString")
				startForgotPasswordOptionsModel.SetLanguage("testString")
				startForgotPasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(startForgotPasswordOptionsModel).ToNot(BeNil())
				Expect(startForgotPasswordOptionsModel.User).To(Equal(core.StringPtr("testString")))
				Expect(startForgotPasswordOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(startForgotPasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewStartSignUpOptions successfully`, func() {
				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				Expect(createNewUserEmailsItemModel).ToNot(BeNil())
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)
				Expect(createNewUserEmailsItemModel.Value).To(Equal(core.StringPtr("user@mail.com")))
				Expect(createNewUserEmailsItemModel.Primary).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the StartSignUpOptions model
				shouldCreateProfile := true
				startSignUpOptionsEmails := []appidmanagementv4.CreateNewUserEmailsItem{}
				startSignUpOptionsPassword := "userPassword"
				startSignUpOptionsModel := appIdManagementService.NewStartSignUpOptions(shouldCreateProfile, startSignUpOptionsEmails, startSignUpOptionsPassword)
				startSignUpOptionsModel.SetShouldCreateProfile(true)
				startSignUpOptionsModel.SetEmails([]appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel})
				startSignUpOptionsModel.SetPassword("userPassword")
				startSignUpOptionsModel.SetActive(true)
				startSignUpOptionsModel.SetUserName("myUserName")
				startSignUpOptionsModel.SetLanguage("testString")
				startSignUpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(startSignUpOptionsModel).ToNot(BeNil())
				Expect(startSignUpOptionsModel.ShouldCreateProfile).To(Equal(core.BoolPtr(true)))
				Expect(startSignUpOptionsModel.Emails).To(Equal([]appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}))
				Expect(startSignUpOptionsModel.Password).To(Equal(core.StringPtr("userPassword")))
				Expect(startSignUpOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(startSignUpOptionsModel.UserName).To(Equal(core.StringPtr("myUserName")))
				Expect(startSignUpOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(startSignUpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateApplicationOptions successfully`, func() {
				// Construct an instance of the UpdateApplicationOptions model
				clientID := "testString"
				updateApplicationOptionsName := "testString"
				updateApplicationOptionsModel := appIdManagementService.NewUpdateApplicationOptions(clientID, updateApplicationOptionsName)
				updateApplicationOptionsModel.SetClientID("testString")
				updateApplicationOptionsModel.SetName("testString")
				updateApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateApplicationOptionsModel).ToNot(BeNil())
				Expect(updateApplicationOptionsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(updateApplicationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateChannelOptions successfully`, func() {
				// Construct an instance of the UpdateChannelOptions model
				channel := "email"
				updateChannelOptionsIsActive := true
				updateChannelOptionsModel := appIdManagementService.NewUpdateChannelOptions(channel, updateChannelOptionsIsActive)
				updateChannelOptionsModel.SetChannel("email")
				updateChannelOptionsModel.SetIsActive(true)
				updateChannelOptionsModel.SetConfig(map[string]interface{}{"anyKey": "anyValue"})
				updateChannelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateChannelOptionsModel).ToNot(BeNil())
				Expect(updateChannelOptionsModel.Channel).To(Equal(core.StringPtr("email")))
				Expect(updateChannelOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(updateChannelOptionsModel.Config).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateChannelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCloudDirectoryUserOptions successfully`, func() {
				// Construct an instance of the CreateNewUserEmailsItem model
				createNewUserEmailsItemModel := new(appidmanagementv4.CreateNewUserEmailsItem)
				Expect(createNewUserEmailsItemModel).ToNot(BeNil())
				createNewUserEmailsItemModel.Value = core.StringPtr("user@mail.com")
				createNewUserEmailsItemModel.Primary = core.BoolPtr(true)
				Expect(createNewUserEmailsItemModel.Value).To(Equal(core.StringPtr("user@mail.com")))
				Expect(createNewUserEmailsItemModel.Primary).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateCloudDirectoryUserOptions model
				userID := "testString"
				updateCloudDirectoryUserOptionsEmails := []appidmanagementv4.CreateNewUserEmailsItem{}
				updateCloudDirectoryUserOptionsModel := appIdManagementService.NewUpdateCloudDirectoryUserOptions(userID, updateCloudDirectoryUserOptionsEmails)
				updateCloudDirectoryUserOptionsModel.SetUserID("testString")
				updateCloudDirectoryUserOptionsModel.SetEmails([]appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel})
				updateCloudDirectoryUserOptionsModel.SetActive(true)
				updateCloudDirectoryUserOptionsModel.SetUserName("myUserName")
				updateCloudDirectoryUserOptionsModel.SetPassword("userPassword")
				updateCloudDirectoryUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCloudDirectoryUserOptionsModel).ToNot(BeNil())
				Expect(updateCloudDirectoryUserOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(updateCloudDirectoryUserOptionsModel.Emails).To(Equal([]appidmanagementv4.CreateNewUserEmailsItem{*createNewUserEmailsItemModel}))
				Expect(updateCloudDirectoryUserOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(updateCloudDirectoryUserOptionsModel.UserName).To(Equal(core.StringPtr("myUserName")))
				Expect(updateCloudDirectoryUserOptionsModel.Password).To(Equal(core.StringPtr("userPassword")))
				Expect(updateCloudDirectoryUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateExtensionActiveOptions successfully`, func() {
				// Construct an instance of the UpdateExtensionActiveOptions model
				name := "premfa"
				updateExtensionActiveOptionsIsActive := true
				updateExtensionActiveOptionsModel := appIdManagementService.NewUpdateExtensionActiveOptions(name, updateExtensionActiveOptionsIsActive)
				updateExtensionActiveOptionsModel.SetName("premfa")
				updateExtensionActiveOptionsModel.SetIsActive(true)
				updateExtensionActiveOptionsModel.SetConfig(map[string]interface{}{"anyKey": "anyValue"})
				updateExtensionActiveOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateExtensionActiveOptionsModel).ToNot(BeNil())
				Expect(updateExtensionActiveOptionsModel.Name).To(Equal(core.StringPtr("premfa")))
				Expect(updateExtensionActiveOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(updateExtensionActiveOptionsModel.Config).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateExtensionActiveOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateExtensionConfigOptions successfully`, func() {
				// Construct an instance of the UpdateExtensionConfigConfig model
				updateExtensionConfigConfigModel := new(appidmanagementv4.UpdateExtensionConfigConfig)
				Expect(updateExtensionConfigConfigModel).ToNot(BeNil())
				updateExtensionConfigConfigModel.URL = core.StringPtr("testString")
				updateExtensionConfigConfigModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				Expect(updateExtensionConfigConfigModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(updateExtensionConfigConfigModel.HeadersVar).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the UpdateExtensionConfigOptions model
				name := "premfa"
				updateExtensionConfigOptionsIsActive := true
				updateExtensionConfigOptionsModel := appIdManagementService.NewUpdateExtensionConfigOptions(name, updateExtensionConfigOptionsIsActive)
				updateExtensionConfigOptionsModel.SetName("premfa")
				updateExtensionConfigOptionsModel.SetIsActive(true)
				updateExtensionConfigOptionsModel.SetConfig(updateExtensionConfigConfigModel)
				updateExtensionConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateExtensionConfigOptionsModel).ToNot(BeNil())
				Expect(updateExtensionConfigOptionsModel.Name).To(Equal(core.StringPtr("premfa")))
				Expect(updateExtensionConfigOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(updateExtensionConfigOptionsModel.Config).To(Equal(updateExtensionConfigConfigModel))
				Expect(updateExtensionConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateLocalizationOptions successfully`, func() {
				// Construct an instance of the UpdateLocalizationOptions model
				updateLocalizationOptionsModel := appIdManagementService.NewUpdateLocalizationOptions()
				updateLocalizationOptionsModel.SetLanguages([]string{"testString"})
				updateLocalizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLocalizationOptionsModel).ToNot(BeNil())
				Expect(updateLocalizationOptionsModel.Languages).To(Equal([]string{"testString"}))
				Expect(updateLocalizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMFAConfigOptions successfully`, func() {
				// Construct an instance of the UpdateMFAConfigOptions model
				updateMfaConfigOptionsIsActive := true
				updateMfaConfigOptionsModel := appIdManagementService.NewUpdateMFAConfigOptions(updateMfaConfigOptionsIsActive)
				updateMfaConfigOptionsModel.SetIsActive(true)
				updateMfaConfigOptionsModel.SetConfig(map[string]interface{}{"anyKey": "anyValue"})
				updateMfaConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMfaConfigOptionsModel).ToNot(BeNil())
				Expect(updateMfaConfigOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(updateMfaConfigOptionsModel.Config).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateMfaConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRateLimitConfigOptions successfully`, func() {
				// Construct an instance of the UpdateRateLimitConfigOptions model
				updateRateLimitConfigOptionsSignUpLimitPerMinute := int64(50)
				updateRateLimitConfigOptionsSignInLimitPerMinute := int64(60)
				updateRateLimitConfigOptionsModel := appIdManagementService.NewUpdateRateLimitConfigOptions(updateRateLimitConfigOptionsSignUpLimitPerMinute, updateRateLimitConfigOptionsSignInLimitPerMinute)
				updateRateLimitConfigOptionsModel.SetSignUpLimitPerMinute(int64(50))
				updateRateLimitConfigOptionsModel.SetSignInLimitPerMinute(int64(60))
				updateRateLimitConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRateLimitConfigOptionsModel).ToNot(BeNil())
				Expect(updateRateLimitConfigOptionsModel.SignUpLimitPerMinute).To(Equal(core.Int64Ptr(int64(50))))
				Expect(updateRateLimitConfigOptionsModel.SignInLimitPerMinute).To(Equal(core.Int64Ptr(int64(60))))
				Expect(updateRateLimitConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRedirectUrisOptions successfully`, func() {
				// Construct an instance of the RedirectUriConfig model
				redirectUriConfigModel := new(appidmanagementv4.RedirectUriConfig)
				Expect(redirectUriConfigModel).ToNot(BeNil())
				redirectUriConfigModel.RedirectUris = []string{"http://localhost:3000/oauth-callback"}
				redirectUriConfigModel.TrustCloudIAMRedirectUris = core.BoolPtr(true)
				redirectUriConfigModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(redirectUriConfigModel.RedirectUris).To(Equal([]string{"http://localhost:3000/oauth-callback"}))
				Expect(redirectUriConfigModel.TrustCloudIAMRedirectUris).To(Equal(core.BoolPtr(true)))
				Expect(redirectUriConfigModel.GetProperties()).ToNot(BeEmpty())
				Expect(redirectUriConfigModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateRedirectUrisOptions model
				var redirectUrisArray *appidmanagementv4.RedirectUriConfig = nil
				updateRedirectUrisOptionsModel := appIdManagementService.NewUpdateRedirectUrisOptions(redirectUrisArray)
				updateRedirectUrisOptionsModel.SetRedirectUrisArray(redirectUriConfigModel)
				updateRedirectUrisOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRedirectUrisOptionsModel).ToNot(BeNil())
				Expect(updateRedirectUrisOptionsModel.RedirectUrisArray).To(Equal(redirectUriConfigModel))
				Expect(updateRedirectUrisOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRoleOptions successfully`, func() {
				// Construct an instance of the UpdateRoleParamsAccessItem model
				updateRoleParamsAccessItemModel := new(appidmanagementv4.UpdateRoleParamsAccessItem)
				Expect(updateRoleParamsAccessItemModel).ToNot(BeNil())
				updateRoleParamsAccessItemModel.ApplicationID = core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")
				updateRoleParamsAccessItemModel.Scopes = []string{"cartoons", "animated"}
				Expect(updateRoleParamsAccessItemModel.ApplicationID).To(Equal(core.StringPtr("de33d272-f8a7-4406-8fe8-ab28fd457be5")))
				Expect(updateRoleParamsAccessItemModel.Scopes).To(Equal([]string{"cartoons", "animated"}))

				// Construct an instance of the UpdateRoleOptions model
				roleID := "testString"
				updateRoleOptionsName := "child"
				updateRoleOptionsAccess := []appidmanagementv4.UpdateRoleParamsAccessItem{}
				updateRoleOptionsModel := appIdManagementService.NewUpdateRoleOptions(roleID, updateRoleOptionsName, updateRoleOptionsAccess)
				updateRoleOptionsModel.SetRoleID("testString")
				updateRoleOptionsModel.SetName("child")
				updateRoleOptionsModel.SetAccess([]appidmanagementv4.UpdateRoleParamsAccessItem{*updateRoleParamsAccessItemModel})
				updateRoleOptionsModel.SetDescription("Limits the available movie options to those that might be more appropriate for younger viewers.")
				updateRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRoleOptionsModel).ToNot(BeNil())
				Expect(updateRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(updateRoleOptionsModel.Name).To(Equal(core.StringPtr("child")))
				Expect(updateRoleOptionsModel.Access).To(Equal([]appidmanagementv4.UpdateRoleParamsAccessItem{*updateRoleParamsAccessItemModel}))
				Expect(updateRoleOptionsModel.Description).To(Equal(core.StringPtr("Limits the available movie options to those that might be more appropriate for younger viewers.")))
				Expect(updateRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRoleParamsAccessItem successfully`, func() {
				applicationID := "de33d272-f8a7-4406-8fe8-ab28fd457be5"
				scopes := []string{"cartoons", "animated"}
				model, err := appIdManagementService.NewUpdateRoleParamsAccessItem(applicationID, scopes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateSSOConfigOptions successfully`, func() {
				// Construct an instance of the UpdateSSOConfigOptions model
				updateSsoConfigOptionsIsActive := true
				updateSsoConfigOptionsInactivityTimeoutSeconds := float64(86400)
				updateSsoConfigOptionsLogoutRedirectUris := []string{"http://localhost:3000/logout-callback"}
				updateSsoConfigOptionsModel := appIdManagementService.NewUpdateSSOConfigOptions(updateSsoConfigOptionsIsActive, updateSsoConfigOptionsInactivityTimeoutSeconds, updateSsoConfigOptionsLogoutRedirectUris)
				updateSsoConfigOptionsModel.SetIsActive(true)
				updateSsoConfigOptionsModel.SetInactivityTimeoutSeconds(float64(86400))
				updateSsoConfigOptionsModel.SetLogoutRedirectUris([]string{"http://localhost:3000/logout-callback"})
				updateSsoConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSsoConfigOptionsModel).ToNot(BeNil())
				Expect(updateSsoConfigOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(updateSsoConfigOptionsModel.InactivityTimeoutSeconds).To(Equal(core.Float64Ptr(float64(86400))))
				Expect(updateSsoConfigOptionsModel.LogoutRedirectUris).To(Equal([]string{"http://localhost:3000/logout-callback"}))
				Expect(updateSsoConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTemplateOptions successfully`, func() {
				// Construct an instance of the UpdateTemplateOptions model
				templateName := "USER_VERIFICATION"
				language := "testString"
				updateTemplateOptionsSubject := "testString"
				updateTemplateOptionsModel := appIdManagementService.NewUpdateTemplateOptions(templateName, language, updateTemplateOptionsSubject)
				updateTemplateOptionsModel.SetTemplateName("USER_VERIFICATION")
				updateTemplateOptionsModel.SetLanguage("testString")
				updateTemplateOptionsModel.SetSubject("testString")
				updateTemplateOptionsModel.SetHTMLBody("testString")
				updateTemplateOptionsModel.SetBase64EncodedHTMLBody("testString")
				updateTemplateOptionsModel.SetPlainTextBody("testString")
				updateTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTemplateOptionsModel).ToNot(BeNil())
				Expect(updateTemplateOptionsModel.TemplateName).To(Equal(core.StringPtr("USER_VERIFICATION")))
				Expect(updateTemplateOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.Subject).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.HTMLBody).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.Base64EncodedHTMLBody).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.PlainTextBody).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserProfilesConfigOptions successfully`, func() {
				// Construct an instance of the UpdateUserProfilesConfigOptions model
				updateUserProfilesConfigOptionsIsActive := true
				updateUserProfilesConfigOptionsModel := appIdManagementService.NewUpdateUserProfilesConfigOptions(updateUserProfilesConfigOptionsIsActive)
				updateUserProfilesConfigOptionsModel.SetIsActive(true)
				updateUserProfilesConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserProfilesConfigOptionsModel).ToNot(BeNil())
				Expect(updateUserProfilesConfigOptionsModel.IsActive).To(Equal(core.BoolPtr(true)))
				Expect(updateUserProfilesConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserRolesOptions successfully`, func() {
				// Construct an instance of the UpdateUserRolesParamsRoles model
				updateUserRolesParamsRolesModel := new(appidmanagementv4.UpdateUserRolesParamsRoles)
				Expect(updateUserRolesParamsRolesModel).ToNot(BeNil())
				updateUserRolesParamsRolesModel.Ids = []string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}
				Expect(updateUserRolesParamsRolesModel.Ids).To(Equal([]string{"111c22c3-38ea-4de8-b5d4-338744d83b0f"}))

				// Construct an instance of the UpdateUserRolesOptions model
				id := "testString"
				var updateUserRolesOptionsRoles *appidmanagementv4.UpdateUserRolesParamsRoles = nil
				updateUserRolesOptionsModel := appIdManagementService.NewUpdateUserRolesOptions(id, updateUserRolesOptionsRoles)
				updateUserRolesOptionsModel.SetID("testString")
				updateUserRolesOptionsModel.SetRoles(updateUserRolesParamsRolesModel)
				updateUserRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserRolesOptionsModel).ToNot(BeNil())
				Expect(updateUserRolesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserRolesOptionsModel.Roles).To(Equal(updateUserRolesParamsRolesModel))
				Expect(updateUserRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUserProfilesExportOptions successfully`, func() {
				// Construct an instance of the UserProfilesExportOptions model
				userProfilesExportOptionsModel := appIdManagementService.NewUserProfilesExportOptions()
				userProfilesExportOptionsModel.SetStartIndex(int64(38))
				userProfilesExportOptionsModel.SetCount(int64(0))
				userProfilesExportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(userProfilesExportOptionsModel).ToNot(BeNil())
				Expect(userProfilesExportOptionsModel.StartIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(userProfilesExportOptionsModel.Count).To(Equal(core.Int64Ptr(int64(0))))
				Expect(userProfilesExportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUserProfilesImportOptions successfully`, func() {
				// Construct an instance of the ExportUserProfileUsersItemIdentitiesItem model
				exportUserProfileUsersItemIdentitiesItemModel := new(appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem)
				Expect(exportUserProfileUsersItemIdentitiesItemModel).ToNot(BeNil())
				exportUserProfileUsersItemIdentitiesItemModel.Provider = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemIdentitiesItemModel.IdpUserInfo = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemIdentitiesItemModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(exportUserProfileUsersItemIdentitiesItemModel.Provider).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemIdentitiesItemModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemIdentitiesItemModel.IdpUserInfo).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(exportUserProfileUsersItemIdentitiesItemModel.GetProperties()).ToNot(BeEmpty())
				Expect(exportUserProfileUsersItemIdentitiesItemModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ExportUserProfileUsersItem model
				exportUserProfileUsersItemModel := new(appidmanagementv4.ExportUserProfileUsersItem)
				Expect(exportUserProfileUsersItemModel).ToNot(BeNil())
				exportUserProfileUsersItemModel.ID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Identities = []appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{*exportUserProfileUsersItemIdentitiesItemModel}
				exportUserProfileUsersItemModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				exportUserProfileUsersItemModel.Name = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Email = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Picture = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Gender = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Locale = core.StringPtr("testString")
				exportUserProfileUsersItemModel.PreferredUsername = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Idp = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedIdpID = core.StringPtr("testString")
				exportUserProfileUsersItemModel.HashedEmail = core.StringPtr("testString")
				exportUserProfileUsersItemModel.Roles = []string{"testString"}
				Expect(exportUserProfileUsersItemModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Identities).To(Equal([]appidmanagementv4.ExportUserProfileUsersItemIdentitiesItem{*exportUserProfileUsersItemIdentitiesItemModel}))
				Expect(exportUserProfileUsersItemModel.Attributes).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(exportUserProfileUsersItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Picture).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Gender).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Locale).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.PreferredUsername).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Idp).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.HashedIdpID).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.HashedEmail).To(Equal(core.StringPtr("testString")))
				Expect(exportUserProfileUsersItemModel.Roles).To(Equal([]string{"testString"}))

				// Construct an instance of the UserProfilesImportOptions model
				userProfilesImportOptionsUsers := []appidmanagementv4.ExportUserProfileUsersItem{}
				userProfilesImportOptionsModel := appIdManagementService.NewUserProfilesImportOptions(userProfilesImportOptionsUsers)
				userProfilesImportOptionsModel.SetUsers([]appidmanagementv4.ExportUserProfileUsersItem{*exportUserProfileUsersItemModel})
				userProfilesImportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(userProfilesImportOptionsModel).ToNot(BeNil())
				Expect(userProfilesImportOptionsModel.Users).To(Equal([]appidmanagementv4.ExportUserProfileUsersItem{*exportUserProfileUsersItemModel}))
				Expect(userProfilesImportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUserVerificationResultOptions successfully`, func() {
				// Construct an instance of the UserVerificationResultOptions model
				userVerificationResultOptionsContext := "testString"
				userVerificationResultOptionsModel := appIdManagementService.NewUserVerificationResultOptions(userVerificationResultOptionsContext)
				userVerificationResultOptionsModel.SetContext("testString")
				userVerificationResultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(userVerificationResultOptionsModel).ToNot(BeNil())
				Expect(userVerificationResultOptionsModel.Context).To(Equal(core.StringPtr("testString")))
				Expect(userVerificationResultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUsersDeleteUserProfileOptions successfully`, func() {
				// Construct an instance of the UsersDeleteUserProfileOptions model
				id := "testString"
				usersDeleteUserProfileOptionsModel := appIdManagementService.NewUsersDeleteUserProfileOptions(id)
				usersDeleteUserProfileOptionsModel.SetID("testString")
				usersDeleteUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(usersDeleteUserProfileOptionsModel).ToNot(BeNil())
				Expect(usersDeleteUserProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(usersDeleteUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUsersGetUserProfileOptions successfully`, func() {
				// Construct an instance of the UsersGetUserProfileOptions model
				id := "testString"
				usersGetUserProfileOptionsModel := appIdManagementService.NewUsersGetUserProfileOptions(id)
				usersGetUserProfileOptionsModel.SetID("testString")
				usersGetUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(usersGetUserProfileOptionsModel).ToNot(BeNil())
				Expect(usersGetUserProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(usersGetUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUsersNominateUserOptions successfully`, func() {
				// Construct an instance of the UsersNominateUserParamsProfile model
				usersNominateUserParamsProfileModel := new(appidmanagementv4.UsersNominateUserParamsProfile)
				Expect(usersNominateUserParamsProfileModel).ToNot(BeNil())
				usersNominateUserParamsProfileModel.Attributes = make(map[string]interface{})
				Expect(usersNominateUserParamsProfileModel.Attributes).To(Equal(make(map[string]interface{})))

				// Construct an instance of the UsersNominateUserOptions model
				usersNominateUserOptionsIdp := "saml"
				usersNominateUserOptionsIdpIdentity := "appid@ibm.com"
				usersNominateUserOptionsModel := appIdManagementService.NewUsersNominateUserOptions(usersNominateUserOptionsIdp, usersNominateUserOptionsIdpIdentity)
				usersNominateUserOptionsModel.SetIdp("saml")
				usersNominateUserOptionsModel.SetIdpIdentity("appid@ibm.com")
				usersNominateUserOptionsModel.SetProfile(usersNominateUserParamsProfileModel)
				usersNominateUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(usersNominateUserOptionsModel).ToNot(BeNil())
				Expect(usersNominateUserOptionsModel.Idp).To(Equal(core.StringPtr("saml")))
				Expect(usersNominateUserOptionsModel.IdpIdentity).To(Equal(core.StringPtr("appid@ibm.com")))
				Expect(usersNominateUserOptionsModel.Profile).To(Equal(usersNominateUserParamsProfileModel))
				Expect(usersNominateUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUsersRevokeRefreshTokenOptions successfully`, func() {
				// Construct an instance of the UsersRevokeRefreshTokenOptions model
				id := "testString"
				usersRevokeRefreshTokenOptionsModel := appIdManagementService.NewUsersRevokeRefreshTokenOptions(id)
				usersRevokeRefreshTokenOptionsModel.SetID("testString")
				usersRevokeRefreshTokenOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(usersRevokeRefreshTokenOptionsModel).ToNot(BeNil())
				Expect(usersRevokeRefreshTokenOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(usersRevokeRefreshTokenOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUsersSearchUserProfileOptions successfully`, func() {
				// Construct an instance of the UsersSearchUserProfileOptions model
				dataScope := "index"
				usersSearchUserProfileOptionsModel := appIdManagementService.NewUsersSearchUserProfileOptions(dataScope)
				usersSearchUserProfileOptionsModel.SetDataScope("index")
				usersSearchUserProfileOptionsModel.SetEmail("testString")
				usersSearchUserProfileOptionsModel.SetID("testString")
				usersSearchUserProfileOptionsModel.SetStartIndex(int64(38))
				usersSearchUserProfileOptionsModel.SetCount(int64(0))
				usersSearchUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(usersSearchUserProfileOptionsModel).ToNot(BeNil())
				Expect(usersSearchUserProfileOptionsModel.DataScope).To(Equal(core.StringPtr("index")))
				Expect(usersSearchUserProfileOptionsModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(usersSearchUserProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(usersSearchUserProfileOptionsModel.StartIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(usersSearchUserProfileOptionsModel.Count).To(Equal(core.Int64Ptr(int64(0))))
				Expect(usersSearchUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUsersSetUserProfileOptions successfully`, func() {
				// Construct an instance of the UsersSetUserProfileOptions model
				id := "testString"
				usersSetUserProfileOptionsAttributes := make(map[string]interface{})
				usersSetUserProfileOptionsModel := appIdManagementService.NewUsersSetUserProfileOptions(id, usersSetUserProfileOptionsAttributes)
				usersSetUserProfileOptionsModel.SetID("testString")
				usersSetUserProfileOptionsModel.SetAttributes(make(map[string]interface{}))
				usersSetUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(usersSetUserProfileOptionsModel).ToNot(BeNil())
				Expect(usersSetUserProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(usersSetUserProfileOptionsModel.Attributes).To(Equal(make(map[string]interface{})))
				Expect(usersSetUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewApmSchema successfully`, func() {
				var advancedPasswordManagement *appidmanagementv4.ApmSchemaAdvancedPasswordManagement = nil
				_, err := appIdManagementService.NewApmSchema(advancedPasswordManagement)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCloudDirectoryConfigParams successfully`, func() {
				selfServiceEnabled := true
				var interactions *appidmanagementv4.CloudDirectoryConfigParamsInteractions = nil
				_, err := appIdManagementService.NewCloudDirectoryConfigParams(selfServiceEnabled, interactions)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCloudDirectorySenderDetails successfully`, func() {
				var senderDetails *appidmanagementv4.CloudDirectorySenderDetailsSenderDetails = nil
				_, err := appIdManagementService.NewCloudDirectorySenderDetails(senderDetails)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCustomIdPConfigParams successfully`, func() {
				isActive := true
				model, err := appIdManagementService.NewCustomIdPConfigParams(isActive)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEmailDispatcherParams successfully`, func() {
				provider := "sendgrid"
				model, err := appIdManagementService.NewEmailDispatcherParams(provider)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExportUser successfully`, func() {
				users := []appidmanagementv4.ExportUserUsersItem{}
				model, err := appIdManagementService.NewExportUser(users)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExportUserProfile successfully`, func() {
				users := []appidmanagementv4.ExportUserProfileUsersItem{}
				model, err := appIdManagementService.NewExportUserProfile(users)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExtensionActive successfully`, func() {
				isActive := true
				model, err := appIdManagementService.NewExtensionActive(isActive)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFacebookGoogleConfigParams successfully`, func() {
				isActive := true
				model, err := appIdManagementService.NewFacebookGoogleConfigParams(isActive)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetLanguages successfully`, func() {
				languages := []string{"testString"}
				model, err := appIdManagementService.NewGetLanguages(languages)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRefreshTokenConfigParams successfully`, func() {
				expiresIn := float64(72.5)
				enabled := true
				model, err := appIdManagementService.NewRefreshTokenConfigParams(expiresIn, enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSamlConfigParams successfully`, func() {
				entityID := "testString"
				signInURL := "testString"
				certificates := []string{"testString"}
				model, err := appIdManagementService.NewSamlConfigParams(entityID, signInURL, certificates)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTokenClaimMapping successfully`, func() {
				source := "saml"
				model, err := appIdManagementService.NewTokenClaimMapping(source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTokenConfigParams successfully`, func() {
				expiresIn := float64(72.5)
				model, err := appIdManagementService.NewTokenConfigParams(expiresIn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateExtensionConfig successfully`, func() {
				isActive := true
				model, err := appIdManagementService.NewUpdateExtensionConfig(isActive)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
