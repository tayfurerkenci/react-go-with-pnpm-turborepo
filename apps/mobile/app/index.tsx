import React from 'react';
import {
  Text,
  View,
  StyleSheet,
  ScrollView,
  TouchableOpacity,
} from 'react-native';
import { StatusBar } from 'expo-status-bar';
import { MaterialIcons } from '@expo/vector-icons';

interface HomeScreenProps {
  onNavigate?: (screen: string) => void;
}

export default function HomeScreen({ onNavigate }: HomeScreenProps) {
  const handleNavigation = (screen: string) => {
    if (onNavigate) {
      onNavigate(screen);
    }
  };

  return (
    <ScrollView style={styles.container}>
      {/* Content */}
      <View style={styles.content}>
        <Text style={styles.sectionTitle}>Quick Actions</Text>

        {/* Movies Card */}
        <TouchableOpacity
          style={styles.card}
          onPress={() => handleNavigation('Movies')}
        >
          <View style={styles.cardContent}>
            <View style={styles.cardLeft}>
              <View
                style={[styles.iconContainer, { backgroundColor: '#2563eb' }]}
              >
                <MaterialIcons name="movie" size={24} color="#fff" />
              </View>
              <View>
                <Text style={styles.cardTitle}>Popular Movies</Text>
                <Text style={styles.cardSubtitle}>
                  Discover trending movies
                </Text>
              </View>
            </View>
            <View style={styles.cardRight}>
              <MaterialIcons name="star" size={16} color="#fbbf24" />
              <Text style={styles.cardBadge}>Popular</Text>
            </View>
          </View>
        </TouchableOpacity>

        {/* TV Shows Card */}
        <TouchableOpacity
          style={styles.card}
          onPress={() => handleNavigation('TV')}
        >
          <View style={styles.cardContent}>
            <View style={styles.cardLeft}>
              <View
                style={[styles.iconContainer, { backgroundColor: '#7c3aed' }]}
              >
                <MaterialIcons name="tv" size={24} color="#fff" />
              </View>
              <View>
                <Text style={styles.cardTitle}>Popular TV Shows</Text>
                <Text style={styles.cardSubtitle}>Explore trending series</Text>
              </View>
            </View>
            <View style={styles.cardRight}>
              <MaterialIcons name="star" size={16} color="#fbbf24" />
              <Text style={styles.cardBadge}>Trending</Text>
            </View>
          </View>
        </TouchableOpacity>

        {/* Info Section */}
        <View style={styles.infoSection}>
          <Text style={styles.infoTitle}>Built with Turborepo</Text>
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
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#111827',
  },
  header: {
    backgroundColor: '#1f2937',
    padding: 24,
    paddingTop: 60,
  },
  headerContent: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 8,
  },
  headerTitle: {
    color: '#fff',
    fontSize: 24,
    fontWeight: 'bold',
    marginLeft: 8,
  },
  headerSubtitle: {
    color: '#9ca3af',
    fontSize: 16,
    lineHeight: 24,
  },
  content: {
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
