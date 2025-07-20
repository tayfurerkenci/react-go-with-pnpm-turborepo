import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Tv, Star, Calendar } from 'lucide-react';

// Mock TV show data since we don't have a TV shows endpoint yet
const mockTvShows = [
  {
    id: '1',
    name: 'Breaking Bad',
    overview:
      "A high school chemistry teacher diagnosed with inoperable lung cancer turns to manufacturing and selling methamphetamine in order to secure his family's future.",
    posterPath: '/3xnWaLQjelJDDF7LT1WBo6f4BRe.jpg',
    firstAirDate: '2008-01-20',
    voteAverage: 9.5,
    numberOfSeasons: 5,
    numberOfEpisodes: 62,
    status: 'Ended',
  },
  {
    id: '2',
    name: 'Stranger Things',
    overview:
      'When a young boy vanishes, a small town uncovers a mystery involving secret experiments, terrifying supernatural forces, and one strange little girl.',
    posterPath: '/49WJfeN0moxb9IPfGn8AIqMGskD.jpg',
    firstAirDate: '2016-07-15',
    voteAverage: 8.7,
    numberOfSeasons: 4,
    numberOfEpisodes: 42,
    status: 'Returning Series',
  },
  {
    id: '3',
    name: 'The Crown',
    overview:
      "Follows the political rivalries and romance of Queen Elizabeth II's reign and the events that shaped the second half of the twentieth century.",
    posterPath: '/1M876KPjulVwppEpldhdc8V4o68.jpg',
    firstAirDate: '2016-11-04',
    voteAverage: 8.2,
    numberOfSeasons: 6,
    numberOfEpisodes: 60,
    status: 'Ended',
  },
  {
    id: '4',
    name: 'Game of Thrones',
    overview:
      'Seven noble families fight for control of the mythical land of Westeros. Friction between the houses leads to full-scale war.',
    posterPath: '/u3bZgnGQ9T01sWNhyveQz0wH0Hl.jpg',
    firstAirDate: '2011-04-17',
    voteAverage: 9.3,
    numberOfSeasons: 8,
    numberOfEpisodes: 73,
    status: 'Ended',
  },
  {
    id: '5',
    name: 'The Office',
    overview:
      'A mockumentary on a group of typical office workers, where the workday consists of ego clashes, inappropriate behavior, and tedium.',
    posterPath: '/7DJKHzAi83BmQrWLrYYOqcoKfhR.jpg',
    firstAirDate: '2005-03-24',
    voteAverage: 8.7,
    numberOfSeasons: 9,
    numberOfEpisodes: 201,
    status: 'Ended',
  },
  {
    id: '6',
    name: 'Friends',
    overview:
      'Follows the personal and professional lives of six twenty to thirty-something-year-old friends living in Manhattan.',
    posterPath: '/f496cm9enuEsZkSPzCwnTESEK5s.jpg',
    firstAirDate: '1994-09-22',
    voteAverage: 8.9,
    numberOfSeasons: 10,
    numberOfEpisodes: 236,
    status: 'Ended',
  },
  {
    id: '7',
    name: 'The Mandalorian',
    overview:
      'The travels of a lone bounty hunter in the outer reaches of the galaxy, far from the authority of the New Republic.',
    posterPath: '/sWgBv7LV2PRoQgkxwlibdGXKz1S.jpg',
    firstAirDate: '2019-11-12',
    voteAverage: 8.7,
    numberOfSeasons: 3,
    numberOfEpisodes: 24,
    status: 'Returning Series',
  },
  {
    id: '8',
    name: 'The Witcher',
    overview:
      'Geralt of Rivia, a mutated monster-hunter for hire, journeys toward his destiny in a turbulent world where people often prove more wicked than beasts.',
    posterPath: '/7vjaCdMw15FEbXyLQTVa04URsPm.jpg',
    firstAirDate: '2019-12-20',
    voteAverage: 8.2,
    numberOfSeasons: 3,
    numberOfEpisodes: 24,
    status: 'Returning Series',
  },
  {
    id: '9',
    name: 'House of the Dragon',
    overview:
      'The Targaryen dynasty is at the absolute apex of its power, with more than 15 dragons under their yoke.',
    posterPath: '/1X4h40fcB4WWUmIBK0auT4lVmo8.jpg',
    firstAirDate: '2022-08-21',
    voteAverage: 8.4,
    numberOfSeasons: 2,
    numberOfEpisodes: 18,
    status: 'Returning Series',
  },
  {
    id: '10',
    name: 'Wednesday',
    overview:
      "A sleuthing, supernaturally infused mystery charting Wednesday Addams' years as a student at Nevermore Academy.",
    posterPath: '/9PFonBhy4cQy7Jz20NpMygczOkv.jpg',
    firstAirDate: '2022-11-23',
    voteAverage: 8.6,
    numberOfSeasons: 1,
    numberOfEpisodes: 8,
    status: 'Returning Series',
  },
];

export function TvShowList() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold tracking-tight">Popüler Diziler</h2>
        <div className="text-sm text-muted-foreground">
          {mockTvShows.length} dizi
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {mockTvShows.map((show) => (
          <Card
            key={show.id}
            className="group hover:shadow-lg transition-shadow"
          >
            <CardHeader className="pb-3">
              <CardTitle className="text-lg line-clamp-1">
                {show.name}
              </CardTitle>
              <CardDescription className="flex items-center gap-2">
                <span className="flex items-center gap-1">
                  <Star className="h-3 w-3 fill-yellow-400 text-yellow-400" />
                  <span className="text-xs font-medium">
                    {show.voteAverage.toFixed(1)}
                  </span>
                </span>
                <span className="flex items-center gap-1">
                  <Calendar className="h-3 w-3" />
                  <span className="text-xs">
                    {new Date(show.firstAirDate).getFullYear()}
                  </span>
                </span>
              </CardDescription>
            </CardHeader>

            <CardContent className="space-y-3">
              <div className="aspect-[2/3] bg-muted rounded-md overflow-hidden">
                <img
                  src={`https://image.tmdb.org/t/p/w300${show.posterPath}`}
                  alt={show.name}
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                />
              </div>

              <p className="text-sm text-muted-foreground line-clamp-3">
                {show.overview}
              </p>

              <div className="space-y-2">
                <div className="flex items-center justify-between text-xs text-muted-foreground">
                  <span>{show.numberOfSeasons} Sezon</span>
                  <span>{show.numberOfEpisodes} Bölüm</span>
                </div>

                <div className="flex items-center justify-between">
                  <span
                    className={`text-xs px-2 py-1 rounded-full ${
                      show.status === 'Ended'
                        ? 'bg-red-100 text-red-800'
                        : 'bg-green-100 text-green-800'
                    }`}
                  >
                    {show.status === 'Ended' ? 'Bitti' : 'Devam Ediyor'}
                  </span>
                  <Button variant="outline" size="sm">
                    <Tv className="h-3 w-3 mr-1" />
                    Detay
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
