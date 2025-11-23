import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { OpenAPI } from './api/core/OpenAPI'

// Configure API Base URL to use relative path /api which Nginx proxies to backend
OpenAPI.BASE = '/api';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
