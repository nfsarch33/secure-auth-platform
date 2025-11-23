	Describe("SignIn", func() {
		It("should authenticate user successfully", func() {
			email := "test@example.com"
			password := "Password123!"
			hashedPassword, _ := password.HashPassword(password)

			existingUser := &models.User{
				ID:           uuid.New(),
				Email:        email,
				PasswordHash: hashedPassword,
			}

			mockRepo.EXPECT().GetByEmail(ctx, email).Return(existingUser, nil)

			user, token, err := authService.SignIn(ctx, email, password)
			Expect(err).NotTo(HaveOccurred())
			Expect(user).NotTo(BeNil())
			Expect(user.Email).To(Equal(email))
			Expect(token).NotTo(BeEmpty())
		})

		It("should return error for invalid credentials (user not found)", func() {
			email := "test@example.com"
			password := "Password123!"

			mockRepo.EXPECT().GetByEmail(ctx, email).Return(nil, nil)

			user, token, err := authService.SignIn(ctx, email, password)
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
			Expect(token).To(BeEmpty())
			Expect(err).To(MatchError(service.ErrInvalidCredentials))
		})

		It("should return error for invalid credentials (wrong password)", func() {
			email := "test@example.com"
			password := "Password123!"
			wrongPassword := "WrongPass!"
			hashedPassword, _ := password.HashPassword(password)

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
