import { Stack } from 'expo-router';
import { StatusBar } from 'expo-status-bar';

export default function RootLayout() {
  return (
    <>
      <Stack
        screenOptions={{
          headerStyle: {
            backgroundColor: '#1f2937',
          },
          headerTintColor: '#fff',
          headerTitleStyle: {
            fontWeight: 'bold',
          },
        }}
      >
        <Stack.Screen
          name="index"
          options={{
            title: 'MovieDB Mobile',
            headerShown: true,
          }}
        />
        <Stack.Screen
          name="movies"
          options={{
            title: 'Popular Movies',
            presentation: 'modal',
          }}
        />
        <Stack.Screen
          name="tv"
          options={{
            title: 'Popular TV Shows',
            presentation: 'modal',
          }}
        />
      </Stack>
      <StatusBar style="light" />
    </>
  );
}
