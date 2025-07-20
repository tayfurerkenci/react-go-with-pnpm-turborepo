import { useGetPopularMoviesQuery } from '@api/generated/api';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Film, Star } from 'lucide-react';

interface MovieListProps {
  page?: number;
}

export function MovieList({ page = 1 }: MovieListProps) {
  const {
    data: moviesData,
    isLoading,
    error,
  } = useGetPopularMoviesQuery({ page });

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {Array.from({ length: 8 }).map((_, index) => (
          <Card key={index} className="animate-pulse">
            <CardHeader>
              <div className="h-4 bg-gray-200 rounded w-3/4"></div>
              <div className="h-3 bg-gray-200 rounded w-1/2"></div>
            </CardHeader>
            <CardContent>
              <div className="h-32 bg-gray-200 rounded mb-4"></div>
              <div className="h-3 bg-gray-200 rounded mb-2"></div>
              <div className="h-3 bg-gray-200 rounded w-2/3"></div>
            </CardContent>
          </Card>
        ))}
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col items-center justify-center p-8">
        <Film className="h-16 w-16 text-muted-foreground mb-4" />
        <h2 className="text-xl font-semibold mb-2">Filmler yüklenemiyor</h2>
        <p className="text-muted-foreground text-center">
          Bir hata oluştu. Lütfen daha sonra tekrar deneyin.
        </p>
      </div>
    );
  }

  if (!moviesData?.movies?.length) {
    return (
      <div className="flex flex-col items-center justify-center p-8">
        <Film className="h-16 w-16 text-muted-foreground mb-4" />
        <h2 className="text-xl font-semibold mb-2">Film bulunamadı</h2>
        <p className="text-muted-foreground">Henüz gösterilecek film yok.</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold tracking-tight">Popüler Filmler</h2>
        <div className="text-sm text-muted-foreground">
          Sayfa {moviesData.page} / {moviesData.totalPages}
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {moviesData.movies.slice(0, 10).map((movie) => (
          <Card
            key={movie.id}
            className="group hover:shadow-lg transition-shadow"
          >
            <CardHeader className="pb-3">
              <CardTitle className="text-lg line-clamp-1">
                {movie.title}
              </CardTitle>
              <CardDescription className="flex items-center gap-2">
                {movie.voteAverage && (
                  <span className="flex items-center gap-1">
                    <Star className="h-3 w-3 fill-yellow-400 text-yellow-400" />
                    <span className="text-xs font-medium">
                      {movie.voteAverage.toFixed(1)}
                    </span>
                  </span>
                )}
                {movie.releaseDate && (
                  <span className="text-xs">
                    ({new Date(movie.releaseDate).getFullYear()})
                  </span>
                )}
              </CardDescription>
            </CardHeader>

            <CardContent className="space-y-3">
              {movie.posterPath && (
                <div className="aspect-[2/3] bg-muted rounded-md overflow-hidden">
                  <img
                    src={`https://image.tmdb.org/t/p/w300${movie.posterPath}`}
                    alt={movie.title}
                    className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                  />
                </div>
              )}

              {movie.overview && (
                <p className="text-sm text-muted-foreground line-clamp-3">
                  {movie.overview}
                </p>
              )}

              <div className="flex items-center justify-between pt-2">
                {movie.runtime && (
                  <span className="text-xs text-muted-foreground">
                    {movie.runtime} dk
                  </span>
                )}
                <Button variant="outline" size="sm">
                  Detay
                </Button>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
