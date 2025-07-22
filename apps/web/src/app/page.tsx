import React from 'react';

export default function HomePage(): React.JSX.Element {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 to-slate-800">
      {/* Hero Section */}
      <div className="container mx-auto px-6 py-24">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-white mb-6">
            Welcome to MovieDB
          </h1>
          <p className="text-xl text-gray-300 mb-12 max-w-2xl mx-auto">
            Discover movies and TV shows powered by our modern monorepo
            architecture with Turborepo, Next.js, and Go backend.
          </p>

          {/* Feature Cards */}
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8 mt-16">
            {/* Movies Card */}
            <div className="bg-slate-800 rounded-lg p-6 border border-slate-700 hover:border-blue-500 transition-colors">
              <div className="text-blue-400 text-4xl mb-4">🎬</div>
              <h3 className="text-xl font-semibold text-white mb-2">
                Popular Movies
              </h3>
              <p className="text-gray-400">
                Browse the latest and most popular movies from around the world.
              </p>
            </div>

            {/* TV Shows Card */}
            <div className="bg-slate-800 rounded-lg p-6 border border-slate-700 hover:border-purple-500 transition-colors">
              <div className="text-purple-400 text-4xl mb-4">📺</div>
              <h3 className="text-xl font-semibold text-white mb-2">
                TV Shows
              </h3>
              <p className="text-gray-400">
                Explore trending TV series and discover your next binge-watch.
              </p>
            </div>

            {/* API Card */}
            <div className="bg-slate-800 rounded-lg p-6 border border-slate-700 hover:border-green-500 transition-colors">
              <div className="text-green-400 text-4xl mb-4">⚡</div>
              <h3 className="text-xl font-semibold text-white mb-2">
                Fast API
              </h3>
              <p className="text-gray-400">
                Powered by Go backend with MongoDB for lightning-fast responses.
              </p>
            </div>
          </div>

          {/* Tech Stack */}
          <div className="mt-16">
            <h2 className="text-2xl font-bold text-white mb-8">
              Built with Modern Tech Stack
            </h2>
            <div className="flex flex-wrap justify-center gap-4">
              {[
                'Next.js',
                'React Native',
                'Go',
                'MongoDB',
                'Turborepo',
                'TypeScript',
                'Tailwind CSS',
              ].map((tech) => (
                <span
                  key={tech}
                  className="px-4 py-2 bg-slate-700 text-gray-300 rounded-full text-sm"
                >
                  {tech}
                </span>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
