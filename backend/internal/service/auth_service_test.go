package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/mocks/repository"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/password"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
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
		tokenService = jwt.NewTokenService("secret", "auth-service", time.Hour)
		authService = service.NewAuthService(mockRepo, tokenService)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("SignUp", func() {
		It("should create a new user successfully", func() {
			email := "test@example.com"
			password := "Password123!"

			// Expect GetByEmail to return nil (user does not exist)
			mockRepo.EXPECT().GetByEmail(ctx, email).Return(nil, nil)

			// Expect Create to be called
			mockRepo.EXPECT().Create(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, u *models.User) error {
				Expect(u.Email).To(Equal(email))
				Expect(u.PasswordHash).NotTo(Equal(password)) // Should be hashed
				Expect(u.PasswordHash).NotTo(BeEmpty())
				return nil
			})

			user, err := authService.SignUp(ctx, email, password)
			Expect(err).NotTo(HaveOccurred())
			Expect(user).NotTo(BeNil())
			Expect(user.Email).To(Equal(email))
		})

		It("should return error if user already exists", func() {
			email := "existing@example.com"
			password := "Password123!"

			existingUser := &models.User{ID: uuid.New(), Email: email}
			mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

			user, err := authService.SignUp(ctx, email, password)
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
			Expect(err).To(MatchError(service.ErrUserAlreadyExists))
		})
	})

	Describe("SignIn", func() {
		It("should authenticate user successfully", func() {
			email := "test@example.com"
			pass := "Password123!"
			hashedPassword, _ := password.HashPassword(pass)

			existingUser := &models.User{
				ID:           uuid.New(),
				Email:        email,
				PasswordHash: hashedPassword,
			}

			mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

			user, token, err := authService.SignIn(ctx, email, pass)
			Expect(err).NotTo(HaveOccurred())
			Expect(user).NotTo(BeNil())
			Expect(user.Email).To(Equal(email))
			Expect(token).NotTo(BeEmpty())
		})

		It("should return error for invalid credentials (user not found)", func() {
			email := "test@example.com"
			pass := "Password123!"

			mockRepo.EXPECT().GetByEmail(ctx, email).Return(nil, nil)

			user, token, err := authService.SignIn(ctx, email, pass)
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
			Expect(token).To(BeEmpty())
			Expect(err).To(MatchError(service.ErrInvalidCredentials))
		})

		It("should return error for invalid credentials (wrong password)", func() {
			email := "test@example.com"
			pass := "Password123!"
			wrongPassword := "WrongPass!"
			hashedPassword, _ := password.HashPassword(pass)

			existingUser := &models.User{
				ID:           uuid.New(),
				Email:        email,
				PasswordHash: hashedPassword,
			}

			mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

			user, token, err := authService.SignIn(ctx, email, wrongPassword)
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
			Expect(token).To(BeEmpty())
			Expect(err).To(MatchError(service.ErrInvalidCredentials))
		})
	})
})
