import React, { useEffect, useState } from 'react';
import {
  Text,
  View,
  StyleSheet,
  ScrollView,
  Image,
  ActivityIndicator,
  TouchableOpacity,
} from 'react-native';
import { StatusBar } from 'expo-status-bar';

interface Movie {
  id: number;
  title: string;
  overview: string;
  poster_path: string;
  vote_average: number;
  release_date: string;
}

const MoviesScreen: React.FC = () => {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    console.log('🎬 MoviesScreen component mounted');
    fetchPopularMovies();
  }, []);

  const fetchPopularMovies = async () => {
    try {
      console.log('📡 Fetching movies from backend...');
      setLoading(true);
      // Backend'den film verilerini çek
      const response = await fetch(
        'http://192.168.1.42:8080/api/v1/movies/popular'
      );
      console.log('📡 Response status:', response.status);

      const data = await response.json();
      console.log('📡 Response data:', data);

      if (response.ok) {
        setMovies(data.results || []);
        console.log('✅ Movies loaded:', data.results?.length || 0);
      } else {
        setError('Failed to fetch movies');
        console.error('❌ Failed to fetch movies');
      }
    } catch (err) {
      setError('Network error occurred');
      console.error('❌ Network error:', err);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <View style={styles.centerContainer}>
        <ActivityIndicator size="large" color="#3b82f6" />
        <Text style={styles.loadingText}>Loading popular movies...</Text>
      </View>
    );
  }

  if (error) {
    return (
      <View style={styles.centerContainer}>
        <Text style={styles.errorIcon}>❌</Text>
        <Text style={styles.errorText}>{error}</Text>
        <TouchableOpacity
          style={styles.retryButton}
          onPress={fetchPopularMovies}
        >
          <Text style={styles.retryButtonText}>Retry</Text>
        </TouchableOpacity>
      </View>
    );
  }

  return (
    <ScrollView style={styles.container}>
      <View style={styles.moviesGrid}>
        {movies.map((movie) => (
          <TouchableOpacity key={movie.id} style={styles.movieCard}>
            <Image
              source={{
                uri: movie.poster_path
                  ? `https://image.tmdb.org/t/p/w500${movie.poster_path}`
                  : 'https://via.placeholder.com/500x750/1f2937/ffffff?text=No+Image',
              }}
              style={styles.moviePoster}
              resizeMode="cover"
            />
            <View style={styles.movieInfo}>
              <Text style={styles.movieTitle} numberOfLines={2}>
                {movie.title}
              </Text>
              <Text style={styles.movieYear}>
                {movie.release_date
                  ? new Date(movie.release_date).getFullYear()
                  : 'N/A'}
              </Text>
              <View style={styles.ratingContainer}>
                <Text style={styles.starIcon}>⭐</Text>
                <Text style={styles.ratingText}>
                  {movie.vote_average ? movie.vote_average.toFixed(1) : 'N/A'}
                </Text>
              </View>
              {movie.overview && (
                <Text style={styles.movieOverview} numberOfLines={3}>
                  {movie.overview}
                </Text>
              )}
            </View>
          </TouchableOpacity>
        ))}
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#111827',
  },
  centerContainer: {
    flex: 1,
    backgroundColor: '#111827',
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  header: {
    backgroundColor: '#1f2937',
    padding: 20,
    paddingTop: 60,
  },
  headerTitle: {
    color: '#fff',
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 4,
  },
  headerSubtitle: {
    color: '#9ca3af',
    fontSize: 16,
  },
  loadingText: {
    color: '#9ca3af',
    fontSize: 16,
    marginTop: 12,
  },
  errorText: {
    color: '#ef4444',
    fontSize: 16,
    textAlign: 'center',
    marginTop: 12,
    marginBottom: 20,
  },
  errorIcon: {
    fontSize: 48,
    marginBottom: 12,
  },
  starIcon: {
    fontSize: 16,
    marginRight: 4,
  },
  retryButton: {
    backgroundColor: '#3b82f6',
    paddingHorizontal: 20,
    paddingVertical: 10,
    borderRadius: 8,
  },
  retryButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  moviesGrid: {
    padding: 16,
  },
  movieCard: {
    backgroundColor: '#1f2937',
    borderRadius: 12,
    marginBottom: 16,
    overflow: 'hidden',
    borderWidth: 1,
    borderColor: '#374151',
  },
  moviePoster: {
    width: '100%',
    height: 300,
  },
  movieInfo: {
    padding: 16,
  },
  movieTitle: {
    color: '#fff',
    fontSize: 18,
    fontWeight: '600',
    marginBottom: 4,
  },
  movieYear: {
    color: '#9ca3af',
    fontSize: 14,
    marginBottom: 8,
  },
  ratingContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 8,
  },
  ratingText: {
    color: '#fbbf24',
    fontSize: 14,
    marginLeft: 4,
    fontWeight: '600',
  },
  movieOverview: {
    color: '#d1d5db',
    fontSize: 14,
    lineHeight: 20,
  },
});

export default MoviesScreen;
