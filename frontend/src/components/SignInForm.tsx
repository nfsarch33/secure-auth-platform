import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { DefaultService as AuthService } from '../api';

const signInSchema = z.object({
  email: z.string().email('Email is required'),
  password: z.string().min(1, 'Password is required'),
});

type SignInFormData = z.infer<typeof signInSchema>;

export const SignInForm: React.FC = () => {
  const [message, setMessage] = useState<string>('');
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<SignInFormData>({
    resolver: zodResolver(signInSchema),
  });

  const onSubmit = async (data: SignInFormData) => {
    try {
      const response = await AuthService.signIn({
        email: data.email,
        password: data.password,
      });
      setMessage('Sign in successful!');
      if (response.token) {
        localStorage.setItem('token', response.token);
      }
    } catch (error) {
      console.error(error);
      setMessage('Sign in failed.');
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} aria-labelledby="signin-heading">
      <h1 id="signin-heading">Sign In</h1>
      {message && <div role="status">{message}</div>}
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
        Sign In
      </button>
    </form>
  );
};
