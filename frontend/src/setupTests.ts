import '@testing-library/jest-dom';
import { vi } from 'vitest';

// Mock react-google-recaptcha-v3
vi.mock('react-google-recaptcha-v3', () => ({
  GoogleReCaptchaProvider: ({ children }: { children: React.ReactNode }) => children,
  useGoogleReCaptcha: () => ({
    executeRecaptcha: async () => 'mock-captcha-token',
  }),
}));
