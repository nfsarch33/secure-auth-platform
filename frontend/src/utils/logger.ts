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
  trace: (...msg: unknown[]) => log.trace(...msg),
  debug: (...msg: unknown[]) => log.debug(...msg),
  info: (...msg: unknown[]) => log.info(...msg),
  warn: (...msg: unknown[]) => log.warn(...msg),
  error: (...msg: unknown[]) => log.error(...msg),
};

