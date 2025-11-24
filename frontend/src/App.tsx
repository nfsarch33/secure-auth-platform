import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { SignUpForm } from './components/SignUpForm';
import { SignInForm } from './components/SignInForm';
import { ProfilePage } from './components/ProfilePage';
import { CaptchaProvider } from './components/CaptchaProvider';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import { Navbar } from './components/Navbar';
import { ProtectedRoute } from './components/ProtectedRoute';
import './components/Navbar.css';

const HomeContent: React.FC = () => {
  const { isAuthenticated } = useAuth();

  return (
    <div className="auth-container" style={{ textAlign: 'center' }}>
      <h1>Welcome to Secure Auth Platform</h1>
      <p style={{ marginBottom: '1.5rem', color: 'var(--color-text)' }}>
        A production-ready authentication system.
      </p>
      <div style={{ display: 'flex', gap: '1rem', justifyContent: 'center' }}>
        {!isAuthenticated ? (
          <Link to="/signup" style={{ textDecoration: 'none' }}>
            <button>Get Started</button>
          </Link>
        ) : (
          <Link to="/me" style={{ textDecoration: 'none' }}>
            <button>Go to Profile</button>
          </Link>
        )}
      </div>
    </div>
  );
};

const App: React.FC = () => {
  return (
    <CaptchaProvider>
      <AuthProvider>
        <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
          <Navbar />
          <main>
            <Routes>
              <Route path="/signup" element={<SignUpForm />} />
              <Route path="/signin" element={<SignInForm />} />
              <Route element={<ProtectedRoute />}>
                <Route path="/me" element={<ProfilePage />} />
              </Route>
              <Route path="/" element={<HomeContent />} />
            </Routes>
          </main>
        </BrowserRouter>
      </AuthProvider>
    </CaptchaProvider>
  );
};

export default App;
