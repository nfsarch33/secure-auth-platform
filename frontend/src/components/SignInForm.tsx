import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { DefaultService as AuthService } from '../api';
import { useGoogleReCaptcha } from 'react-google-recaptcha-v3';
import { logger } from '../utils/logger';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

const signInSchema = z.object({
  email: z.string().email('Email is required'),
  password: z.string().min(1, 'Password is required'),
});

type SignInFormData = z.infer<typeof signInSchema>;

export const SignInForm: React.FC = () => {
  const [message, setMessage] = useState<string>('');
  const navigate = useNavigate();
  const { signIn } = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<SignInFormData>({
    resolver: zodResolver(signInSchema),
  });

  const { executeRecaptcha } = useGoogleReCaptcha();

  const onSubmit = async (data: SignInFormData) => {
    try {
      let captchaToken = '';
      if (executeRecaptcha) {
        captchaToken = await executeRecaptcha('signin');
      }

      const response = await AuthService.signIn({
        email: data.email,
        password: data.password,
        captchaToken: captchaToken,
      });
      setMessage('Sign in successful!');
      if (response.token) {
        await signIn(response.token);
        // Navigate to profile after a short delay to show success message
        setTimeout(() => {
          navigate('/me');
        }, 1000);
      }
    } catch (error) {
      logger.error('SignIn Error:', error);
      setMessage('Sign in failed.');
    }
  };

  return (
    <div className="auth-container">
      <form onSubmit={handleSubmit(onSubmit)} aria-labelledby="signin-heading">
        <h1 id="signin-heading">Sign In</h1>
        {message && (
          <div role="status" className={`status-message ${message.includes('successful') ? 'success' : 'error'}`}>
            {message}
          </div>
        )}
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input id="email" type="email" placeholder="you@example.com" {...register('email')} />
          {errors.email && <span role="alert">{errors.email.message}</span>}
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input id="password" type="password" placeholder="••••••••" {...register('password')} />
          {errors.password && <span role="alert">{errors.password.message}</span>}
        </div>
        <button type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Signing in...' : 'Sign In'}
        </button>
      </form>
    </div>
  );
};
