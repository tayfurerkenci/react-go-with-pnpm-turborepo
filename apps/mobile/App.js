import React, { useState } from 'react';
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  ScrollView,
  Alert,
} from 'react-native';
import { registerRootComponent } from 'expo';

// Import screens
import MoviesScreen from './app/movies';
import TVScreen from './app/tv';

export default function App() {
  const [activeScreen, setActiveScreen] = useState('home');

  // Debug için
  console.log('App rendered, activeScreen:', activeScreen);

  const handleMoviesPress = () => {
    console.log('🎬 Movies button pressed!');
    console.log('Current activeScreen:', activeScreen);
    setActiveScreen('movies');
  };

  const handleTVShowsPress = () => {
    console.log('📺 TV Shows button pressed!');
    console.log('Current activeScreen:', activeScreen);
    setActiveScreen('tv');
  };

  const handleBackPress = () => {
    console.log('� Back button pressed');
    setActiveScreen('home');
  };

  // Render different screens based on activeScreen
  if (activeScreen === 'movies') {
    return (
      <View style={styles.container}>
        <View style={styles.header}>
          <TouchableOpacity onPress={handleBackPress} style={styles.backButton}>
            <Text style={styles.backButtonText}>← Back</Text>
          </TouchableOpacity>
          <Text style={styles.headerTitle}>Popular Movies</Text>
        </View>
        <MoviesScreen />
      </View>
    );
  }

  if (activeScreen === 'tv') {
    return (
      <View style={styles.container}>
        <View style={styles.header}>
          <TouchableOpacity onPress={handleBackPress} style={styles.backButton}>
            <Text style={styles.backButtonText}>← Back</Text>
          </TouchableOpacity>
          <Text style={styles.headerTitle}>Popular TV Shows</Text>
        </View>
        <TVScreen />
      </View>
    );
  }
  return (
    <View style={styles.container}>
      {/* Header */}
      <View style={styles.header}>
        <Text style={styles.headerTitle}>MovieDB Mobile</Text>
      </View>

      {/* Content */}
      <ScrollView style={styles.content}>
        <Text style={styles.sectionTitle}>Quick Actions</Text>

        {/* Movies Card */}
        <TouchableOpacity style={styles.card} onPress={handleMoviesPress}>
          <View style={styles.cardContent}>
            <View style={styles.cardLeft}>
              <View
                style={[styles.iconContainer, { backgroundColor: '#2563eb' }]}
              >
                <Text style={styles.iconText}>🎬</Text>
              </View>
              <View>
                <Text style={styles.cardTitle}>Popular Movies</Text>
                <Text style={styles.cardSubtitle}>
                  Discover trending movies
                </Text>
              </View>
            </View>
            <View style={styles.cardRight}>
              <Text style={styles.cardBadge}>⭐ Popular</Text>
            </View>
          </View>
        </TouchableOpacity>

        {/* TV Shows Card */}
        <TouchableOpacity style={styles.card} onPress={handleTVShowsPress}>
          <View style={styles.cardContent}>
            <View style={styles.cardLeft}>
              <View
                style={[styles.iconContainer, { backgroundColor: '#7c3aed' }]}
              >
                <Text style={styles.iconText}>📺</Text>
              </View>
              <View>
                <Text style={styles.cardTitle}>Popular TV Shows</Text>
                <Text style={styles.cardSubtitle}>Explore trending series</Text>
              </View>
            </View>
            <View style={styles.cardRight}>
              <Text style={styles.cardBadge}>⭐ Trending</Text>
            </View>
          </View>
        </TouchableOpacity>

        {/* Info Section */}
        <View style={styles.infoSection}>
          <Text style={styles.infoTitle}>🚀 Built with Turborepo</Text>
          <Text style={styles.infoText}>
            This mobile app is part of a scalable monorepo architecture
            featuring:
          </Text>
          <View style={styles.featureList}>
            <Text style={styles.featureItem}>• Shared UI components</Text>
            <Text style={styles.featureItem}>• Cross-platform API client</Text>
            <Text style={styles.featureItem}>• Type-safe development</Text>
            <Text style={styles.featureItem}>
              • Optimized builds with caching
            </Text>
          </View>
        </View>
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#111827',
  },
  header: {
    backgroundColor: '#1f2937',
    paddingHorizontal: 16,
    paddingVertical: 12,
    paddingTop: 16,
  },
  headerTitle: {
    color: '#fff',
    fontSize: 18,
    fontWeight: 'bold',
  },
  backButton: {
    paddingVertical: 8,
    paddingHorizontal: 12,
    marginRight: 12,
  },
  backButtonText: {
    color: '#3b82f6',
    fontSize: 16,
    fontWeight: '600',
  },
  content: {
    flex: 1,
    padding: 24,
  },
  sectionTitle: {
    color: '#fff',
    fontSize: 20,
    fontWeight: '600',
    marginBottom: 16,
  },
  card: {
    backgroundColor: '#1f2937',
    borderRadius: 12,
    padding: 20,
    marginBottom: 16,
    borderWidth: 1,
    borderColor: '#374151',
  },
  cardContent: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  cardLeft: {
    flexDirection: 'row',
    alignItems: 'center',
    flex: 1,
  },
  iconContainer: {
    width: 48,
    height: 48,
    borderRadius: 24,
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: 16,
  },
  cardTitle: {
    color: '#fff',
    fontSize: 18,
    fontWeight: '600',
    marginBottom: 4,
  },
  cardSubtitle: {
    color: '#9ca3af',
    fontSize: 14,
  },
  cardRight: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  cardBadge: {
    color: '#fbbf24',
    fontSize: 12,
    marginLeft: 4,
  },
  infoSection: {
    backgroundColor: '#1f2937',
    borderRadius: 12,
    padding: 20,
    marginTop: 8,
  },
  infoTitle: {
    color: '#fff',
    fontSize: 18,
    fontWeight: '600',
    marginBottom: 12,
  },
  infoText: {
    color: '#9ca3af',
    fontSize: 14,
    lineHeight: 20,
    marginBottom: 12,
  },
  featureList: {
    marginLeft: 8,
  },
  featureItem: {
    color: '#d1d5db',
    fontSize: 14,
    lineHeight: 20,
    marginBottom: 4,
  },
});

registerRootComponent(App);
