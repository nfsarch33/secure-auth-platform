	Describe("SignIn", func() {
		It("should return 200 OK with token on successful signin", func() {
			reqBody := api.SignInRequest{
				Email:    "test@example.com",
				Password: "Password123!",
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signin", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			token := "mock-token"
			mockService.EXPECT().SignIn(gomock.Any(), reqBody.Email, reqBody.Password).
				Return(&models.User{Email: reqBody.Email}, token, nil)

			router.POST("/auth/signin", handler.SignIn)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var resp api.AuthResponse
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.User.Email).To(Equal(reqBody.Email))
			Expect(resp.Token).To(PointTo(Equal(token)))
		})

		It("should return 401 Unauthorized on invalid credentials", func() {
			reqBody := api.SignInRequest{
				Email:    "test@example.com",
				Password: "WrongPassword",
			}
			jsonBody, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/auth/signin", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			mockService.EXPECT().SignIn(gomock.Any(), reqBody.Email, reqBody.Password).
				Return(nil, "", service.ErrInvalidCredentials)

			router.POST("/auth/signin", handler.SignIn)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})
