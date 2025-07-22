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

interface TVShow {
  id: number;
  name: string;
  overview: string;
  poster_path: string;
  vote_average: number;
  first_air_date: string;
}

export default function TVScreen() {
  const [tvShows, setTVShows] = useState<TVShow[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    console.log('📺 TVScreen component mounted');
    fetchPopularTVShows();
  }, []);

  const fetchPopularTVShows = async () => {
    try {
      console.log('📡 Fetching TV shows from backend...');
      setLoading(true);
      // Backend'den TV show verilerini çek
      const response = await fetch(
        'http://192.168.1.42:8080/api/v1/tv/popular'
      );
      console.log('📡 Response status:', response.status);

      const data = await response.json();
      console.log('📡 Response data:', data);

      if (response.ok) {
        setTVShows(data.results || []);
        console.log('✅ TV Shows loaded:', data.results?.length || 0);
      } else {
        setError('Failed to fetch TV shows');
        console.error('❌ Failed to fetch TV shows');
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
        <ActivityIndicator size="large" color="#7c3aed" />
        <Text style={styles.loadingText}>Loading popular TV shows...</Text>
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
          onPress={fetchPopularTVShows}
        >
          <Text style={styles.retryButtonText}>Retry</Text>
        </TouchableOpacity>
      </View>
    );
  }

  return (
    <ScrollView style={styles.container}>
      <View style={styles.showsGrid}>
        {tvShows.map((show) => (
          <TouchableOpacity key={show.id} style={styles.showCard}>
            <Image
              source={{
                uri: show.poster_path
                  ? `https://image.tmdb.org/t/p/w500${show.poster_path}`
                  : 'https://via.placeholder.com/500x750/1f2937/ffffff?text=No+Image',
              }}
              style={styles.showPoster}
              resizeMode="cover"
            />
            <View style={styles.showInfo}>
              <Text style={styles.showTitle} numberOfLines={2}>
                {show.name}
              </Text>
              <Text style={styles.showYear}>
                {show.first_air_date
                  ? new Date(show.first_air_date).getFullYear()
                  : 'N/A'}
              </Text>
              <View style={styles.ratingContainer}>
                <Text style={styles.starIcon}>⭐</Text>
                <Text style={styles.ratingText}>
                  {show.vote_average ? show.vote_average.toFixed(1) : 'N/A'}
                </Text>
              </View>
              {show.overview && (
                <Text style={styles.showOverview} numberOfLines={3}>
                  {show.overview}
                </Text>
              )}
            </View>
          </TouchableOpacity>
        ))}
      </View>
    </ScrollView>
  );
}

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
  retryButton: {
    backgroundColor: '#7c3aed',
    paddingHorizontal: 20,
    paddingVertical: 10,
    borderRadius: 8,
  },
  retryButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  showsGrid: {
    padding: 16,
  },
  showCard: {
    backgroundColor: '#1f2937',
    borderRadius: 12,
    marginBottom: 16,
    overflow: 'hidden',
    borderWidth: 1,
    borderColor: '#374151',
  },
  showPoster: {
    width: '100%',
    height: 300,
  },
  showInfo: {
    padding: 16,
  },
  showTitle: {
    color: '#fff',
    fontSize: 18,
    fontWeight: '600',
    marginBottom: 4,
  },
  showYear: {
    color: '#9ca3af',
    fontSize: 14,
    marginBottom: 8,
  },
  ratingContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 8,
  },
  errorIcon: {
    fontSize: 48,
    marginBottom: 12,
  },
  starIcon: {
    fontSize: 16,
    marginRight: 4,
  },
  ratingText: {
    color: '#fbbf24',
    fontSize: 14,
    marginLeft: 4,
    fontWeight: '600',
  },
  showOverview: {
    color: '#d1d5db',
    fontSize: 14,
    lineHeight: 20,
  },
});
