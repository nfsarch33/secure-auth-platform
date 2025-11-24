import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { SignInForm } from './SignInForm';
import { BrowserRouter } from 'react-router-dom';
import { AuthProvider } from '../contexts/AuthContext';

// Mock the auth service
const mockSignin = vi.fn();
const mockGetCurrentUser = vi.fn();
vi.mock('../api', () => ({
  DefaultService: {
    signIn: (...args: unknown[]) => mockSignin(...args),
    getCurrentUser: (...args: unknown[]) => mockGetCurrentUser(...args),
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

describe('SignInForm', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    mockExecuteRecaptcha.mockResolvedValue('mock-captcha-token');
    // Default mock for getCurrentUser to avoid failures in AuthProvider
    mockGetCurrentUser.mockResolvedValue({ id: '123', email: 'test@example.com' });
  });

  it('renders sign in form elements', async () => {
    render(
      <AuthProvider>
        <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
          <SignInForm />
        </BrowserRouter>
      </AuthProvider>
    );
    await waitFor(() => {
        expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    });
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument();
  });

  it('validates required fields', async () => {
    render(
      <AuthProvider>
        <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
          <SignInForm />
        </BrowserRouter>
      </AuthProvider>
    );
    // Wait for initial render
    await waitFor(() => expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument());

    const submitButton = screen.getByRole('button', { name: /sign in/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/email is required/i)).toBeInTheDocument();
      expect(screen.getByText(/password is required/i)).toBeInTheDocument();
    });
  });

  it('submits form with valid data', async () => {
    const mockUser = { id: '123', email: 'test@example.com', createdAt: new Date().toISOString() };
    const mockToken = 'fake-jwt-token';
    mockSignin.mockResolvedValue({ user: mockUser, token: mockToken });
    mockGetCurrentUser.mockResolvedValue(mockUser);

    render(
      <AuthProvider>
        <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
          <SignInForm />
        </BrowserRouter>
      </AuthProvider>
    );
    const user = userEvent.setup();

    await user.type(screen.getByLabelText(/email/i), 'test@example.com');
    await user.type(screen.getByLabelText(/password/i), 'Password123!');
    await user.click(screen.getByRole('button', { name: /sign in/i }));

    await waitFor(() => {
      expect(mockSignin).toHaveBeenCalledWith({
        email: 'test@example.com',
        password: 'Password123!',
        captchaToken: 'mock-captcha-token',
      });
      expect(screen.getByText(/Sign in successful!/i)).toBeInTheDocument();
    });
  });

  it('displays API error message on sign in failure', async () => {
    const errorMessage = 'Invalid credentials';
    mockSignin.mockRejectedValue(new Error(errorMessage));

    render(
      <AuthProvider>
        <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
          <SignInForm />
        </BrowserRouter>
      </AuthProvider>
    );
    const user = userEvent.setup();

    await user.type(screen.getByLabelText(/email/i), 'wrong@example.com');
    await user.type(screen.getByLabelText(/password/i), 'WrongPass');
    await user.click(screen.getByRole('button', { name: /sign in/i }));

    await waitFor(() => {
        expect(screen.getByRole('status')).toHaveTextContent('Sign in failed.');
    });
  });
});

