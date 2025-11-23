import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { SignUpForm } from './SignUpForm';
import { BrowserRouter } from 'react-router-dom';

// Mock the auth service
const mockSignup = vi.fn();
    vi.mock('../api', () => ({
      DefaultService: {
        signUp: (...args: unknown[]) => mockSignup(...args),
      },
    }));

// Mock useNavigate
const mockNavigate = vi.fn();
vi.mock('react-router-dom', async (importOriginal) => {
  const actual = await importOriginal<typeof import('react-router-dom')>();
  return {
    ...actual,
    useNavigate: () => mockNavigate,
  };
});

// Mock react-google-recaptcha-v3
const mockExecuteRecaptcha = vi.fn();
vi.mock('react-google-recaptcha-v3', () => ({
  useGoogleReCaptcha: () => ({
    executeRecaptcha: mockExecuteRecaptcha,
  }),
}));

describe('SignUpForm', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    mockExecuteRecaptcha.mockResolvedValue('mock-captcha-token');
  });

  it('renders sign up form elements', () => {
    render(
      <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
        <SignUpForm />
      </BrowserRouter>
    );
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /sign up/i })).toBeInTheDocument();
  });

  it('validates required fields', async () => {
    render(
      <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
        <SignUpForm />
      </BrowserRouter>
    );
    const submitButton = screen.getByRole('button', { name: /sign up/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/email is required/i)).toBeInTheDocument();
      expect(screen.getByText(/password must be at least 8 characters/i)).toBeInTheDocument();
    });
  });

  it('submits form with valid data', async () => {
    const mockUser = { id: '123', email: 'test@example.com', createdAt: new Date().toISOString() };
    const mockToken = 'fake-jwt-token';
    mockSignup.mockResolvedValue({ user: mockUser, token: mockToken });

    render(
      <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
        <SignUpForm />
      </BrowserRouter>
    );
    const user = userEvent.setup();

    await user.type(screen.getByLabelText(/email/i), 'test@example.com');
    await user.type(screen.getByLabelText(/password/i), 'Password123!');
    await user.click(screen.getByRole('button', { name: /sign up/i }));

    await waitFor(() => {
      expect(mockSignup).toHaveBeenCalledWith({
        email: 'test@example.com',
        password: 'Password123!',
        captchaToken: 'mock-captcha-token',
      });
      // The component currently sets a message, it doesn't automatically navigate or login in the current implementation shown in read_file
      // But let's check if success message appears
      expect(screen.getByText(/Sign up successful! Please sign in./i)).toBeInTheDocument();
    });
  });

  it('displays API error message on sign up failure', async () => {
    const errorMessage = 'User already exists';
    mockSignup.mockRejectedValue(new Error(errorMessage));

    render(
      <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
        <SignUpForm />
      </BrowserRouter>
    );
    const user = userEvent.setup();

    await user.type(screen.getByLabelText(/email/i), 'existing@example.com');
    await user.type(screen.getByLabelText(/password/i), 'Password123!');
    await user.click(screen.getByRole('button', { name: /sign up/i }));

    await waitFor(() => {
      // The component catches error and sets generic "Sign up failed." message or logs it.
      // Looking at the component code: setMessage('Sign up failed.');
      expect(screen.getByRole('alert')).toHaveTextContent(/Sign up failed/i);
    });
  });
});
