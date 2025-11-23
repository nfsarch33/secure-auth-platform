import React, { useState, useEffect } from 'react';
import { useGoogleReCaptcha } from 'react-google-recaptcha-v3';

interface CaptchaProps {
  onVerify: (token: string) => void;
  action: string;
}

export const Captcha: React.FC<CaptchaProps> = ({ onVerify, action }) => {
  const { executeRecaptcha } = useGoogleReCaptcha();
  const [token, setToken] = useState<string>('');

  useEffect(() => {
    const handleReCaptchaVerify = async () => {
      if (!executeRecaptcha) {
        return;
      }
      const token = await executeRecaptcha(action);
      setToken(token);
      onVerify(token);
    };

    handleReCaptchaVerify();
  }, [executeRecaptcha, action, onVerify]);

  return null; // Invisible reCAPTCHA
};
