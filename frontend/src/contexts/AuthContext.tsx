import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { DefaultService as AuthService, User } from '../api';

interface AuthContextType {
  user: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  signIn: (token: string) => Promise<void>;
  signOut: () => Promise<void>;
  refreshProfile: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  // Initialize state from local storage
  useEffect(() => {
    const initAuth = async () => {
      const token = localStorage.getItem('token');
      if (token) {
        try {
            // Token configuration is already set in main.tsx via OpenAPI.TOKEN resolver
            // But we trigger a profile fetch to validate it and get user details
            const userData = await AuthService.getCurrentUser();
            setUser(userData);
        } catch (error) {
            console.error('Failed to restore auth state:', error);
            localStorage.removeItem('token');
            setUser(null);
        }
      }
      setIsLoading(false);
    };

    initAuth();
  }, []);

  const signIn = async (token: string) => {
    localStorage.setItem('token', token);
    try {
        const userData = await AuthService.getCurrentUser();
        setUser(userData);
    } catch (error) {
        console.error('Failed to fetch user after sign in:', error);
        // Even if profile fetch fails, we have the token, so we count as authenticated?
        // Better to be safe: if we can't get profile, something is wrong.
        // But for now let's assume valid token means mostly authenticated.
        // Actually, if we can't get profile, we shouldn't set user.
        // Let's rely on the caller to handle the initial signin API call success.
    }
  };

  const signOut = async () => {
    try {
      await AuthService.signOut();
    } catch (error) {
      console.error('Sign out API failed', error);
    } finally {
      localStorage.removeItem('token');
      setUser(null);
    }
  };

  const refreshProfile = async () => {
      if (localStorage.getItem('token')) {
          try {
              const userData = await AuthService.getCurrentUser();
              setUser(userData);
          } catch (error) {
              console.error('Failed to refresh profile', error);
              localStorage.removeItem('token');
              setUser(null);
          }
      }
  };

  return (
    <AuthContext.Provider value={{
      user,
      isAuthenticated: !!user,
      isLoading,
      signIn,
      signOut,
      refreshProfile
    }}>
      {children}
    </AuthContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
