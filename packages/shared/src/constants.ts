// API configuration constants
export const API_CONFIG = {
  BASE_URL: process.env.VITE_API_URL || 'http://localhost:8080/api/v1',
  TIMEOUT: 10000,
  RETRY_ATTEMPTS: 3,
} as const;

// Pagination defaults
export const PAGINATION = {
  DEFAULT_LIMIT: 10,
  MAX_LIMIT: 100,
  DEFAULT_OFFSET: 0,
} as const;

// Validation rules
export const VALIDATION = {
  USER: {
    EMAIL_MAX_LENGTH: 100,
    FIRST_NAME_MAX_LENGTH: 50,
    LAST_NAME_MAX_LENGTH: 50,
    MIN_AGE: 0,
    MAX_AGE: 150,
  },
} as const;

// HTTP status codes
export const HTTP_STATUS = {
  OK: 200,
  CREATED: 201,
  NO_CONTENT: 204,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  NOT_FOUND: 404,
  INTERNAL_SERVER_ERROR: 500,
} as const;

// App constants
export const APP = {
  NAME: 'My Monorepo App',
  VERSION: '1.0.0',
  DESCRIPTION: 'Go + MongoDB backend with React frontend',
} as const;
