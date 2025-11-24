package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/handlers"
	recaptchaMocks "github.com/nfsarch33/secure-auth-platform/backend/internal/mocks/recaptcha"
	serviceMocks "github.com/nfsarch33/secure-auth-platform/backend/internal/mocks/service"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

var _ = Describe("AuthHandler", func() {
	var (
		ctrl         *gomock.Controller
		mockService  *serviceMocks.MockAuthService
		mockVerifier *recaptchaMocks.MockVerifier
		handler      *handlers.AuthHandler
		router       *gin.Engine
		w            *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockService = serviceMocks.NewMockAuthService(ctrl)
		mockVerifier = recaptchaMocks.NewMockVerifier(ctrl)
		handler = handlers.NewAuthHandler(mockService, mockVerifier)

		gin.SetMode(gin.TestMode)
		router = gin.New()
		w = httptest.NewRecorder()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("SignUp", func() {
		It("should return 201 Created on successful signup", func() {
			reqBody := api.SignUpRequest{
				Email:    openapi_types.Email("test@example.com"),
				Password: "Password123!",
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			token := "mock-token"
			mockService.EXPECT().SignUp(gomock.Any(), string(reqBody.Email), reqBody.Password).
				Return(&models.User{Email: string(reqBody.Email), ID: uuid.New(), CreatedAt: time.Now()}, token, nil)

			router.POST("/auth/signup", handler.SignUp)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))

			var resp api.AuthResponse
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.User.Email).To(Equal(reqBody.Email))
			Expect(resp.Token).To(Equal(token))
		})

		It("should return 400 if CAPTCHA is invalid", func() {
			token := "invalid-token"
			reqBody := api.SignUpRequest{
				Email:        openapi_types.Email("test@example.com"),
				Password:     "Password123!",
				CaptchaToken: &token,
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			mockVerifier.EXPECT().Verify(gomock.Any(), token).Return(false, nil)

			router.POST("/auth/signup", handler.SignUp)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			var resp api.ErrorResponse
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Error).To(Equal("Invalid CAPTCHA"))
		})

		//nolint:dupl
		It("should return 409 Conflict if user already exists", func() {
			reqBody := api.SignUpRequest{
				Email:    openapi_types.Email("existing@example.com"),
				Password: "Password123!",
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			mockService.EXPECT().SignUp(gomock.Any(), string(reqBody.Email), reqBody.Password).
				Return(nil, "", service.ErrUserAlreadyExists)

			router.POST("/auth/signup", handler.SignUp)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusConflict))
		})
	})

	Describe("SignIn", func() {
		It("should return 200 OK with token on successful signin", func() {
			reqBody := api.SignInRequest{
				Email:    openapi_types.Email("test@example.com"),
				Password: "Password123!",
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signin", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			token := "mock-token"
			mockService.EXPECT().SignIn(gomock.Any(), string(reqBody.Email), reqBody.Password).
				Return(&models.User{Email: string(reqBody.Email), ID: uuid.New(), CreatedAt: time.Now()}, token, nil)

			router.POST("/auth/signin", handler.SignIn)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var resp api.AuthResponse
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Token).To(Equal(token))
		})

		//nolint:dupl
		It("should return 401 Unauthorized on invalid credentials", func() {
			reqBody := api.SignInRequest{
				Email:    openapi_types.Email("test@example.com"),
				Password: "WrongPassword",
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signin", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			mockService.EXPECT().SignIn(gomock.Any(), string(reqBody.Email), reqBody.Password).
				Return(nil, "", service.ErrInvalidCredentials)

			router.POST("/auth/signin", handler.SignIn)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})
})
