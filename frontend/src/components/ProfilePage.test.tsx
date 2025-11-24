import { render, screen, waitFor, fireEvent, act } from '@testing-library/react';
import { vi, describe, it, expect, beforeEach, afterEach } from 'vitest';
import { ProfilePage } from './ProfilePage';
import { DefaultService } from '../api';
import { BrowserRouter } from 'react-router-dom';
import { AuthProvider } from '../contexts/AuthContext';

// Mock the API service
vi.mock('../api', async (importOriginal) => {
  const actual = await importOriginal();
  return {
    ...actual,
    DefaultService: {
      getCurrentUser: vi.fn(),
      signOut: vi.fn(),
    },
  };
});

// Mock navigate
const mockedNavigate = vi.fn();
vi.mock('react-router-dom', async (importOriginal) => {
  const actual = await importOriginal();
  return {
    ...actual,
    useNavigate: () => mockedNavigate,
  };
});

describe('ProfilePage', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    localStorage.clear();
  });

  afterEach(() => {
    vi.useRealTimers();
  });

  it('displays user profile data on success', async () => {
    localStorage.setItem('token', 'test-token');
    
    const mockUser = {
      id: '123',
      email: 'test@example.com',
      created_at: '2023-01-01T00:00:00Z',
    };
    vi.mocked(DefaultService.getCurrentUser).mockResolvedValue(mockUser);

    render(
      <AuthProvider>
        <BrowserRouter>
          <ProfilePage />
        </BrowserRouter>
      </AuthProvider>
    );

    await waitFor(() => {
      expect(screen.getByText(/Your Profile/i)).toBeInTheDocument();
      expect(screen.getByText('test@example.com')).toBeInTheDocument();
      expect(screen.getByText('123')).toBeInTheDocument();
    });
  });

  it('sign out button calls API and clears token', async () => {
    localStorage.setItem('token', 'test-token');
    const mockUser = {
        id: '123',
        email: 'test@example.com',
        created_at: '2023-01-01T00:00:00Z',
      };
    vi.mocked(DefaultService.getCurrentUser).mockResolvedValue(mockUser);
    vi.mocked(DefaultService.signOut).mockResolvedValue({});
    
    render(
      <AuthProvider>
        <BrowserRouter>
          <ProfilePage />
        </BrowserRouter>
      </AuthProvider>
      );
  
    await waitFor(() => {
        expect(screen.getByText('Sign Out')).toBeInTheDocument();
    });

    await act(async () => {
      fireEvent.click(screen.getByText('Sign Out'));
    });

    expect(DefaultService.signOut).toHaveBeenCalled();
    expect(mockedNavigate).toHaveBeenCalledWith('/signin');
  });
});
