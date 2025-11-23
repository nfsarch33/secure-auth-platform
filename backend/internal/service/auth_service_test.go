package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	mocks "github.com/nfsarch33/secure-auth-platform/backend/internal/mocks/repository"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/password"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = Describe("AuthService", func() {
	var (
		ctrl         *gomock.Controller
		mockRepo     *mocks.MockUserRepository
		authService  service.AuthService
		tokenService *jwt.TokenService
		ctx          context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockUserRepository(ctrl)
		tokenService = jwt.NewTokenService("secret", "auth-service", 24*time.Hour)
		authService = service.NewAuthService(mockRepo, tokenService)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("SignUp", func() {
		Context("when user does not exist", func() {
			It("should create a new user and return token", func() {
				email := "test@example.com"
				password := "password123"

				mockRepo.EXPECT().GetByEmail(ctx, email).Return(nil, repository.ErrUserNotFound)
				mockRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)

				user, token, err := authService.SignUp(ctx, email, password)
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(user.Email).To(Equal(email))
				Expect(token).NotTo(BeEmpty())
			})
		})

		Context("when user already exists", func() {
			It("should return error", func() {
				email := "existing@example.com"
				password := "password123"
				existingUser := &models.User{
					ID:    uuid.New(),
					Email: email,
				}

				mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

				user, token, err := authService.SignUp(ctx, email, password)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(service.ErrUserAlreadyExists))
				Expect(user).To(BeNil())
				Expect(token).To(BeEmpty())
			})
		})
	})

	Describe("SignIn", func() {
		Context("with valid credentials", func() {
			It("should return user and token", func() {
				email := "test@example.com"
				plainPassword := "password123"
				hashedPassword, _ := password.HashPassword(plainPassword)
				
				existingUser := &models.User{
					ID:           uuid.New(),
					Email:        email,
					PasswordHash: hashedPassword,
				}

				mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

				user, token, err := authService.SignIn(ctx, email, plainPassword)
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(token).NotTo(BeEmpty())
			})
		})

		Context("with invalid password", func() {
			It("should return invalid credentials error", func() {
				email := "test@example.com"
				plainPassword := "wrongpassword"
				hashedPassword, _ := password.HashPassword("correctpassword")
				
				existingUser := &models.User{
					ID:           uuid.New(),
					Email:        email,
					PasswordHash: hashedPassword,
				}

				mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

				user, token, err := authService.SignIn(ctx, email, plainPassword)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(service.ErrInvalidCredentials))
				Expect(user).To(BeNil())
				Expect(token).To(BeEmpty())
			})
		})
	})
})
