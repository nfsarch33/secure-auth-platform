import { test, expect } from '@playwright/test';

test.describe('Authentication Flow', () => {
  test('should allow a user to sign up and then sign in', async ({ page }) => {
    // 1. Sign Up
    await page.goto('/signup');
    await expect(page.getByRole('heading', { name: 'Sign Up' })).toBeVisible();

    const timestamp = Date.now();
    const email = `test-${timestamp}@example.com`;
    const password = 'Password123!';

    await page.getByLabel('Email').fill(email);
    await page.getByLabel('Password').fill(password);
    await page.getByRole('button', { name: 'Sign Up' }).click();

    // Expect success message
    await expect(page.getByText('Sign up successful! Please sign in.')).toBeVisible();

    // 2. Navigate to Sign In
    // In this app, we stay on the same page or user navigates manually. 
    // Let's assume user clicks "Sign In" link if it existed, or we go directly.
    await page.goto('/signin');
    await expect(page.getByRole('heading', { name: 'Sign In' })).toBeVisible();

    await page.getByLabel('Email').fill(email);
    await page.getByLabel('Password').fill(password);
    await page.getByRole('button', { name: 'Sign In' }).click();

    await expect(page.getByText('Sign in successful!')).toBeVisible();
  });

  test('should show validation errors', async ({ page }) => {
    await page.goto('/signup');
    await page.getByRole('button', { name: 'Sign Up' }).click();

    await expect(page.getByText('Email is required')).toBeVisible();
    await expect(page.getByText('Password must be at least 8 characters')).toBeVisible();
  });

  test('should fail with invalid credentials', async ({ page }) => {
    await page.goto('/signin');
    await expect(page.getByRole('heading', { name: 'Sign In' })).toBeVisible();

    await page.getByLabel('Email').fill('nonexistent@example.com');
    await page.getByLabel('Password').fill('WrongPass123!');
    await page.getByRole('button', { name: 'Sign In' }).click();

    // Expect error message (text depends on frontend implementation, usually 'Invalid credentials')
    // Based on unit tests it might be "Sign in failed." or specific error.
    // Let's check for "Sign in failed" as seen in unit test mocks or "Invalid credentials" if backend returns 401.
    // The unit test expects: expect(screen.getByRole('status')).toHaveTextContent('Sign in failed.');
    await expect(page.getByRole('status')).toContainText(/Sign in failed/i);
  });
});

test.describe('Security Regression', () => {
  test('should have secure headers', async ({ request }) => {
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

    // Check if headers exist (keys are lower-cased)
    expect(headers['x-frame-options']).toBe('DENY');
    expect(headers['x-content-type-options']).toBe('nosniff');
    // HSTS is usually set by middleware
    expect(headers['strict-transport-security']).toBeDefined();
    expect(headers['content-security-policy']).toBeDefined();
  });
});
