import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { DefaultService as AuthService } from '../api';
import { useGoogleReCaptcha } from 'react-google-recaptcha-v3';
import { logger } from '../utils/logger';

const signUpSchema = z.object({
  email: z.string().email('Email is required'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
});

type SignUpFormData = z.infer<typeof signUpSchema>;

export const SignUpForm: React.FC = () => {
  const [message, setMessage] = useState<string>('');
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<SignUpFormData>({
    resolver: zodResolver(signUpSchema),
  });

  const { executeRecaptcha } = useGoogleReCaptcha();

  const onSubmit = async (data: SignUpFormData) => {
    try {
      let captchaToken = '';
      if (executeRecaptcha) {
        captchaToken = await executeRecaptcha('signup');
      }

      await AuthService.signUp({
        email: data.email,
        password: data.password,
        captchaToken: captchaToken,
      });
      setMessage('Sign up successful! Please sign in.');
        } catch (error: unknown) {
          logger.error('SignUp Error:', error);
          // Display detailed error for debugging
          let errorMsg = 'Unknown error';
          if (error instanceof Error) {
             errorMsg = error.message;
          } else if (typeof error === 'object' && error !== null && 'body' in error) {
             errorMsg = JSON.stringify((error as { body: unknown }).body);
          }
          setMessage(`Sign up failed: ${errorMsg}`);
        }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} aria-labelledby="signup-heading">
      <h1 id="signup-heading">Sign Up</h1>
      {message && <div role="alert" aria-live="assertive">{message}</div>}
      <div>
        <label htmlFor="email">Email</label>
        <input id="email" type="email" {...register('email')} />
        {errors.email && <span role="alert">{errors.email.message}</span>}
      </div>
      <div>
        <label htmlFor="password">Password</label>
        <input id="password" type="password" {...register('password')} />
        {errors.password && <span role="alert">{errors.password.message}</span>}
      </div>
      <button type="submit" disabled={isSubmitting}>
        Sign Up
      </button>
    </form>
  );
};
