import { GoogleReCaptchaProvider } from 'react-google-recaptcha-v3';

// Use a dummy key for development/test if not provided
const RECAPTCHA_KEY = import.meta.env.VITE_RECAPTCHA_SITE_KEY || '6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI';

export const CaptchaProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <GoogleReCaptchaProvider reCaptchaKey={RECAPTCHA_KEY}>
      {children}
    </GoogleReCaptchaProvider>
  );
};

