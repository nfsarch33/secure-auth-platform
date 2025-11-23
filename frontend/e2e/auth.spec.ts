import { test, expect } from '@playwright/test';
import { AuthPage } from './pom/AuthPage';
import AxeBuilder from '@axe-core/playwright';

test.describe('Authentication Flow', () => {
  let authPage: AuthPage;

  test.beforeEach(async ({ page }) => {
    authPage = new AuthPage(page);
  });

  test('should allow a user to sign up and then sign in', async ({ page }) => {
    // 1. Sign Up
    await authPage.gotoSignUp();
    await expect(authPage.signUpHeading).toBeVisible();

    // Check Accessibility on Sign Up Page
    const accessibilityScanResults = await new AxeBuilder({ page }).analyze();
    expect(accessibilityScanResults.violations).toEqual([]);

    const timestamp = Date.now();
    const email = `test-${timestamp}@example.com`;
    const password = 'Password123!';

    await authPage.fillEmail(email);
    await authPage.fillPassword(password);
    await authPage.submitSignUp();

    await authPage.expectSignUpSuccess();

    // 2. Navigate to Sign In
    await authPage.gotoSignIn();
    await expect(authPage.signInHeading).toBeVisible();

    // Check Accessibility on Sign In Page
    const accessibilityScanResultsSignIn = await new AxeBuilder({ page }).analyze();
    expect(accessibilityScanResultsSignIn.violations).toEqual([]);

    await authPage.fillEmail(email);
    await authPage.fillPassword(password);
    await authPage.submitSignIn();

    await authPage.expectSignInSuccess();
  });

  test('should show validation errors', async ({ page }) => {
    await authPage.gotoSignUp();
    await authPage.submitSignUp();

    await authPage.expectValidationError('Email is required');
    await authPage.expectValidationError('Password must be at least 8 characters');
  });

  test('should fail with invalid credentials', async ({ page }) => {
    await authPage.gotoSignIn();
    await expect(authPage.signInHeading).toBeVisible();

    await authPage.fillEmail('nonexistent@example.com');
    await authPage.fillPassword('WrongPass123!');
    await authPage.submitSignIn();

    // Based on previous runs, we verify the generic failure message
    // Adjust selector if 'status' role is not used for this specific error message in the component
    // If the component uses a simple div or p for the global error, we might need to look for text directly.
    await expect(page.getByText(/Sign in failed/i)).toBeVisible();
  });
});

test.describe('Security & Regression', () => {
  test('should have secure headers and rate limit headers', async ({ request }) => {
    // Use process.env.BACKEND_URL if available, otherwise default to backend service name in docker
    const backendUrl = process.env.BACKEND_URL || 'http://backend:8080';
    // Send a POST request to trigger the handler and middleware stack securely
    const response = await request.post(`${backendUrl}/auth/signin`, {
      data: {
        email: "test@example.com",
        password: "wrongpassword"
      }
    }); 
    
    const headers = response.headers();
    // console.log('Received Headers:', headers);

    // 1. Security Headers
    expect(headers['x-frame-options']).toBe('DENY');
    expect(headers['x-content-type-options']).toBe('nosniff');
    expect(headers['strict-transport-security']).toBeDefined();
    expect(headers['content-security-policy']).toBeDefined();
    expect(headers['referrer-policy']).toBe('strict-origin-when-cross-origin');
    expect(headers['permissions-policy']).toBeDefined();

    // 2. Rate Limit Verification
    // The current middleware does not expose X-RateLimit headers.
    // To test rate limiting, we would need to exceed the burst/rate.
    // Config: 60 req/min, burst 5.
    
    // We can try to trigger the rate limiter by sending > 5 requests rapidly.
    let rateLimited = false;
    for (let i = 0; i < 10; i++) {
      const res = await request.post(`${backendUrl}/auth/signin`, {
        data: { email: "test@example.com", password: "wrong" }
      });
      if (res.status() === 429) {
        rateLimited = true;
        break;
      }
    }
    expect(rateLimited).toBeTruthy();
  });
});
