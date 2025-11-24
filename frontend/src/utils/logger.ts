import log from 'loglevel';

// Set default level based on environment
// In development, we want to see debug messages
// In production, we typically only want warnings and errors
const level = import.meta.env.DEV ? 'debug' : 'warn';

log.setLevel(level);

/**
 * Structured logger wrapper around 'loglevel'
 * Provides a consistent logging interface for the frontend application.
 */
export const logger = {
  trace: (...msg: any[]) => log.trace(...msg),
  debug: (...msg: any[]) => log.debug(...msg),
  info: (...msg: any[]) => log.info(...msg),
  warn: (...msg: any[]) => log.warn(...msg),
  error: (...msg: any[]) => log.error(...msg),
};

