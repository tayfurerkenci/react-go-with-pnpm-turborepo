import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

// API configuration
const API_CONFIG = {
  BASE_URL: 'http://localhost:8080/api/v1',
  TIMEOUT: 10000,
  RETRY_ATTEMPTS: 3,
} as const;

// Base API configuration
export const baseApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl: API_CONFIG.BASE_URL,
    prepareHeaders: (headers) => {
      // Add auth headers if needed
      const token = localStorage.getItem('authToken');
      if (token) {
        headers.set('authorization', `Bearer ${token}`);
      }
      headers.set('content-type', 'application/json');
      return headers;
    },
  }),
  tagTypes: ['Movie', 'Genre', 'Health'],
  endpoints: () => ({}),
});
