import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import { Provider } from 'react-redux';
import { configureStore } from '@reduxjs/toolkit';
import { moviesApi } from '@api/generated/api';
import { MovieList } from '../MovieList';

// Mock store for testing
const createMockStore = () => {
  return configureStore({
    reducer: {
      [moviesApi.reducerPath]: moviesApi.reducer,
    },
    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware().concat(moviesApi.middleware),
  });
};

// Mock the API hook
vi.mock('@api/generated/api', async () => {
  const actual = await vi.importActual('@api/generated/api');
  return {
    ...actual,
    useGetPopularMoviesQuery: vi.fn(() => ({
      data: {
        movies: [
          {
            id: '1',
            title: 'Test Movie',
            overview: 'Test overview',
            voteAverage: 8.5,
            releaseDate: '2023-01-01',
            posterPath: '/test.jpg',
            runtime: 120,
          },
        ],
        page: 1,
        totalPages: 1,
      },
      isLoading: false,
      error: null,
    })),
  };
});

describe('MovieList Component', () => {
  it('renders movie list correctly', () => {
    const store = createMockStore();

    render(
      <Provider store={store}>
        <MovieList />
      </Provider>
    );

    expect(screen.getByText('Popüler Filmler')).toBeInTheDocument();
    expect(screen.getByText('Test Movie')).toBeInTheDocument();
    expect(screen.getByText('Test overview')).toBeInTheDocument();
  });

  it('shows loading state', () => {
    // Mock loading state
    const mockUseQuery = vi.fn(() => ({
      data: null,
      isLoading: true,
      error: null,
    }));

    vi.doMock('@api/generated/api', () => ({
      useGetPopularMoviesQuery: mockUseQuery,
    }));

    const store = createMockStore();

    render(
      <Provider store={store}>
        <MovieList />
      </Provider>
    );

    // Should show loading skeletons
    const skeletons = screen.getAllByRole('generic');
    expect(skeletons.length).toBeGreaterThan(0);
  });
});
