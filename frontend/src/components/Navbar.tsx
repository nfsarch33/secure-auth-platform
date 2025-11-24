import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

export const Navbar: React.FC = () => {
  const { isAuthenticated, signOut } = useAuth();
  const navigate = useNavigate();

  const handleSignOut = async () => {
    await signOut();
    navigate('/signin');
  };

  return (
    <nav className="navbar">
      <div className="navbar-brand">
        <Link to="/">SecureAuth</Link>
      </div>
      <ul className="navbar-menu">
        <li>
          <Link to="/">Home</Link>
        </li>
        {!isAuthenticated && (
          <>
            <li>
              <Link to="/signin">Sign In</Link>
            </li>
            <li>
              <Link to="/signup">Sign Up</Link>
            </li>
          </>
        )}
        {isAuthenticated && (
          <>
            <li>
              <Link to="/me">Profile</Link>
            </li>
            <li>
              <button onClick={handleSignOut} className="nav-button">
                Sign Out
              </button>
            </li>
          </>
        )}
      </ul>
    </nav>
  );
};

