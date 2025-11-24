import { type Page, type Locator, expect } from '@playwright/test';

export class AuthPage {
  readonly page: Page;
  readonly signUpHeading: Locator;
  readonly signInHeading: Locator;
  readonly emailInput: Locator;
  readonly passwordInput: Locator;
  readonly signUpButton: Locator;
  readonly signInButton: Locator;
  readonly statusMessage: Locator;

  constructor(page: Page) {
    this.page = page;
    this.signUpHeading = page.getByRole('heading', { name: 'Create Account' });
    this.signInHeading = page.getByRole('heading', { name: 'Sign In' });
    this.emailInput = page.getByLabel('Email');
    this.passwordInput = page.getByLabel('Password');
    this.signUpButton = page.getByRole('button', { name: 'Sign Up' });
    this.signInButton = page.getByRole('button', { name: 'Sign In' });
    this.statusMessage = page.getByRole('status');
  }

  async gotoSignUp() {
    await this.page.goto('/signup');
  }

  async gotoSignIn() {
    await this.page.goto('/signin');
  }

  async fillEmail(email: string) {
    await this.emailInput.fill(email);
  }

  async fillPassword(password: string) {
    await this.passwordInput.fill(password);
  }

  async submitSignUp() {
    await this.signUpButton.click();
  }

  async submitSignIn() {
    await this.signInButton.click();
  }

  async expectSignUpSuccess() {
    await expect(this.page.getByText('Sign up successful! Please sign in.')).toBeVisible();
  }

  async expectSignInSuccess() {
    await expect(this.page.getByText('Sign in successful!')).toBeVisible();
  }

  async expectValidationError(text: string | RegExp) {
    await expect(this.page.getByText(text)).toBeVisible();
  }
}

