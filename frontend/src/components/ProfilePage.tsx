import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

export const ProfilePage: React.FC = () => {
  const { user, signOut } = useAuth();
  const navigate = useNavigate();

  // ProtectedRoute handles authentication check and loading state
  if (!user) {
     return null;
  }

  return (
    <div className="auth-container">
      <h1 id="profile-heading">Your Profile</h1>
      <div aria-labelledby="profile-heading" className="profile-details">
        <div className="form-group">
          <label>ID</label>
          <div className="value">{user.id}</div>
        </div>
        <div className="form-group">
          <label>Email</label>
          <div className="value">{user.email}</div>
        </div>
        <div className="form-group">
          <label>Joined</label>
          <div className="value">{new Date(user.created_at).toLocaleDateString()}</div>
        </div>
      </div>
      <button onClick={async () => {
          await signOut();
          navigate('/signin');
      }} style={{ marginTop: '2rem', backgroundColor: '#dc2626' }}>
        Sign Out
      </button>
    </div>
  );
};

