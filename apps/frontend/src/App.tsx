import { MovieList } from '@/components/MovieList';
import { TvShowList } from '@/components/TvShowList';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Film, Tv, Play, Star } from 'lucide-react';
import { useState } from 'react';

function AppContent() {
  const [activeTab, setActiveTab] = useState<'movies' | 'tv'>('movies');

  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <header className="border-b">
        <div className="container mx-auto px-4 py-6">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <Play className="h-8 w-8 text-primary" />
              <h1 className="text-3xl font-bold">MovieDB</h1>
            </div>
            <div className="flex items-center space-x-2">
              <Button
                variant={activeTab === 'movies' ? 'default' : 'outline'}
                onClick={() => setActiveTab('movies')}
                className="flex items-center space-x-2"
              >
                <Film className="h-4 w-4" />
                <span>Filmler</span>
              </Button>
              <Button
                variant={activeTab === 'tv' ? 'default' : 'outline'}
                onClick={() => setActiveTab('tv')}
                className="flex items-center space-x-2"
              >
                <Tv className="h-4 w-4" />
                <span>Diziler</span>
              </Button>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="container mx-auto px-4 py-8">
        {/* Welcome Card */}
        <Card className="mb-8">
          <CardHeader>
            <CardTitle className="flex items-center space-x-2">
              <Star className="h-5 w-5 text-yellow-500" />
              <span>Turborepo + Go Backend + React Frontend Demo</span>
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-muted-foreground">
              Bu uygulama Turborepo monorepo yapısında Go backend, RTK Query API
              client ve React frontend kullanılarak geliştirilmiştir. shadcn/ui
              ve Tailwind CSS ile modern bir tasarım sunmaktadır.
            </p>
          </CardContent>
        </Card>

        {/* Content based on active tab */}
        {activeTab === 'movies' ? <MovieList /> : <TvShowList />}
      </main>

      {/* Footer */}
      <footer className="border-t mt-16">
        <div className="container mx-auto px-4 py-6">
          <div className="flex items-center justify-between text-sm text-muted-foreground">
            <p>© 2025 MovieDB. Turborepo + Go + React Demo.</p>
            <div className="flex items-center space-x-4">
              <span>shadcn/ui</span>
              <span>•</span>
              <span>Tailwind CSS</span>
              <span>•</span>
              <span>RTK Query</span>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

function App() {
  return <AppContent />;
}

export default App;
