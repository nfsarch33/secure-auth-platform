package repository_test

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pashagolub/pgxmock/v3"

	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository"
)

var _ = Describe("PostgresUserRepository", func() {
	var (
		mock pgxmock.PgxPoolIface
		repo repository.UserRepository
		ctx  context.Context
	)

	BeforeEach(func() {
		var err error
		mock, err = pgxmock.NewPool()
		Expect(err).NotTo(HaveOccurred())
		repo = repository.NewPostgresUserRepository(mock)
		ctx = context.Background()
	})

	AfterEach(func() {
		mock.Close()
	})

	Describe("Create", func() {
		It("should create a user successfully", func() {
			user := &models.User{
				ID:           uuid.New(),
				Email:        "test@example.com",
				PasswordHash: "hashedpassword",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			mock.ExpectExec("INSERT INTO users").
				WithArgs(user.ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt).
				WillReturnResult(pgxmock.NewResult("INSERT", 1))

			err := repo.Create(ctx, user)
			Expect(err).NotTo(HaveOccurred())
			Expect(mock.ExpectationsWereMet()).To(Succeed())
		})

		It("should return error if execution fails", func() {
			user := &models.User{
				ID:           uuid.New(),
				Email:        "test@example.com",
				PasswordHash: "hashedpassword",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			mock.ExpectExec("INSERT INTO users").
				WithArgs(user.ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt).
				WillReturnError(errors.New("db error"))

			err := repo.Create(ctx, user)
			Expect(err).To(HaveOccurred())
			Expect(mock.ExpectationsWereMet()).To(Succeed())
		})
	})
})

