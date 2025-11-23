import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { SignUpForm } from './components/SignUpForm';
import { SignInForm } from './components/SignInForm';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/signup">Sign Up</Link>
            </li>
            <li>
              <Link to="/signin">Sign In</Link>
            </li>
          </ul>
        </nav>

        <Routes>
          <Route path="/signup" element={<SignUpForm />} />
          <Route path="/signin" element={<SignInForm />} />
          <Route path="/" element={<h1>Welcome to Secure Auth Platform</h1>} />
        </Routes>
      </div>
    </BrowserRouter>
  );
};

export default App;
