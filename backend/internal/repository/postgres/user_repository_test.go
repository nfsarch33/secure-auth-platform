package postgres_test

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository/postgres"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pashagolub/pgxmock/v3"
)

var _ = Describe("PostgresUserRepository", func() {
	var (
		repo *postgres.PostgresUserRepository
		mock pgxmock.PgxPoolIface
		user *models.User
		ctx  context.Context
	)

	BeforeEach(func() {
		var err error
		mock, err = pgxmock.NewPool()
		Expect(err).NotTo(HaveOccurred())

		repo = postgres.NewPostgresUserRepository(mock)
		ctx = context.Background()

		user = &models.User{
			ID:           uuid.New(),
			Email:        "test@example.com",
			PasswordHash: "hashed_password",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
	})

	AfterEach(func() {
		mock.ExpectationsWereMet()
	})

	Describe("Create", func() {
		It("should create a user successfully", func() {
			mock.ExpectExec("INSERT INTO users").
				WithArgs(user.ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt).
				WillReturnResult(pgxmock.NewResult("INSERT", 1))

			err := repo.Create(ctx, user)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error if database fails", func() {
			mock.ExpectExec("INSERT INTO users").
				WithArgs(user.ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt).
				WillReturnError(errors.New("database error"))

			err := repo.Create(ctx, user)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("database error"))
		})
	})

	Describe("GetByEmail", func() {
		It("should return user if found", func() {
			rows := pgxmock.NewRows([]string{"id", "email", "password_hash", "created_at", "updated_at"}).
				AddRow(user.ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)

			mock.ExpectQuery("SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = \\$1").
				WithArgs(user.Email).
				WillReturnRows(rows)

			result, err := repo.GetByEmail(ctx, user.Email)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).NotTo(BeNil())
			Expect(result.Email).To(Equal(user.Email))
		})

		It("should return error if user not found", func() {
			mock.ExpectQuery("SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = \\$1").
				WithArgs(user.Email).
				WillReturnError(pgx.ErrNoRows)

			result, err := repo.GetByEmail(ctx, user.Email)
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
			Expect(err).To(Equal(pgx.ErrNoRows))
		})
	})
})
