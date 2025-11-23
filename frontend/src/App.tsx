import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { SignUpForm } from './components/SignUpForm';
import { SignInForm } from './components/SignInForm';

const App: React.FC = () => {
  return (
    <BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
      <main>
        <nav>
          <ul>
            <li>
              <Link to="/signup" style={{ color: '#000000', textDecoration: 'underline' }}>Sign Up</Link>
            </li>
            <li>
              <Link to="/signin" style={{ color: '#000000', textDecoration: 'underline' }}>Sign In</Link>
            </li>
          </ul>
        </nav>

        <Routes>
          <Route path="/signup" element={<SignUpForm />} />
          <Route path="/signin" element={<SignInForm />} />
          <Route path="/" element={<h1>Welcome to Secure Auth Platform</h1>} />
        </Routes>
      </main>
    </BrowserRouter>
  );
};

export default App;
