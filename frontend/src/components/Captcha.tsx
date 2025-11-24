import React, { useEffect } from 'react';
import { useGoogleReCaptcha } from 'react-google-recaptcha-v3';

interface CaptchaProps {
  onVerify: (token: string) => void;
  action: string;
}

export const Captcha: React.FC<CaptchaProps> = ({ onVerify, action }) => {
  const { executeRecaptcha } = useGoogleReCaptcha();

  useEffect(() => {
    const handleReCaptchaVerify = async () => {
      if (!executeRecaptcha) {
        return;
      }
      const token = await executeRecaptcha(action);
      onVerify(token);
    };

    handleReCaptchaVerify();
  }, [executeRecaptcha, action, onVerify]);

  return null; // Invisible reCAPTCHA
};
