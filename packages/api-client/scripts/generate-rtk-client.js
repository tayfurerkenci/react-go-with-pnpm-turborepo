const fs = require('fs');
const path = require('path');

// RTK Query client generator script
// Bu script OpenAPI schema'dan RTK Query endpoints üretir

function generateRTKClient() {
  const template = `import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import type { paths } from './schema';
import { API_CONFIG } from '@shared/utils';

// Base API slice
export const api = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl: API_CONFIG.BASE_URL,
    prepareHeaders: (headers) => {
      // Add auth headers if needed
      // const token = localStorage.getItem('authToken');
      // if (token) {
      //   headers.set('authorization', \`Bearer \${token}\`);
      // }
      headers.set('content-type', 'application/json');
      return headers;
    },
  }),
  tagTypes: ['User', 'Health'],
  endpoints: (builder) => ({
    // Health endpoints
    getHealth: builder.query<
      paths['/health']['get']['responses']['200']['content']['application/json'],
      void
    >({
      query: () => '/health',
      providesTags: ['Health'],
    }),

    // User endpoints
    getUsers: builder.query<
      paths['/users']['get']['responses']['200']['content']['application/json'],
      paths['/users']['get']['parameters']['query']
    >({
      query: (params) => ({
        url: '/users',
        params,
      }),
      providesTags: (result) =>
        result
          ? [
              ...result.users.map(({ id }) => ({ type: 'User' as const, id })),
              { type: 'User', id: 'LIST' },
            ]
          : [{ type: 'User', id: 'LIST' }],
    }),

    getUserById: builder.query<
      paths['/users/{id}']['get']['responses']['200']['content']['application/json'],
      string
    >({
      query: (id) => \`/users/\${id}\`,
      providesTags: (result, error, id) => [{ type: 'User', id }],
    }),

    createUser: builder.mutation<
      paths['/users']['post']['responses']['201']['content']['application/json'],
      paths['/users']['post']['requestBody']['content']['application/json']
    >({
      query: (body) => ({
        url: '/users',
        method: 'POST',
        body,
      }),
      invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    updateUser: builder.mutation<
      paths['/users/{id}']['put']['responses']['200']['content']['application/json'],
      { id: string; body: paths['/users/{id}']['put']['requestBody']['content']['application/json'] }
    >({
      query: ({ id, body }) => ({
        url: \`/users/\${id}\`,
        method: 'PUT',
        body,
      }),
      invalidatesTags: (result, error, { id }) => [
        { type: 'User', id },
        { type: 'User', id: 'LIST' },
      ],
    }),

    deleteUser: builder.mutation<void, string>({
      query: (id) => ({
        url: \`/users/\${id}\`,
        method: 'DELETE',
      }),
      invalidatesTags: (result, error, id) => [
        { type: 'User', id },
        { type: 'User', id: 'LIST' },
      ],
    }),
  }),
});

// Export hooks for usage in functional components
export const {
  useGetHealthQuery,
  useGetUsersQuery,
  useGetUserByIdQuery,
  useCreateUserMutation,
  useUpdateUserMutation,
  useDeleteUserMutation,
} = api;

// Export endpoints for direct usage
export const {
  getHealth,
  getUsers,
  getUserById,
  createUser,
  updateUser,
  deleteUser,
} = api.endpoints;
`;

  const outputPath = path.join(__dirname, '../src/generated/api.ts');
  const outputDir = path.dirname(outputPath);

  // Create directory if it doesn't exist
  if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
  }

  fs.writeFileSync(outputPath, template);
  console.log('✅ RTK Query client generated successfully!');
}

generateRTKClient();
