import React from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

export const ProtectedRoute: React.FC = () => {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <div className="auth-container">
        <p>Loading...</p>
      </div>
    );
  }

  return isAuthenticated ? <Outlet /> : <Navigate to="/signin" replace />;
};

