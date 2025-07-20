# 🎬 MovieDB - Turborepo Demo Uygulaması

Bu proje, Turborepo monorepo yapısında Go backend, RTK Query API client ve React frontend kullanılarak geliştirilmiş modern bir film ve dizi kataloğu uygulamasıdır.

## 🏗️ Proje Yapısı

```
my-monorepo/
├── apps/
│   ├── backend/          # Go + Gin + MongoDB backend
│   └── frontend/         # React + TypeScript + Vite frontend
└── packages/
    ├── api-client/       # RTK Query API client (auto-generated)
    ├── shared/           # Paylaşılan utilities
    └── oas/             # OpenAPI 3.0 schema
```

## 🚀 Teknolojiler

### Backend

- **Go 1.21** - Modern, performanslı backend dili
- **Gin Framework** - Hızlı HTTP web framework
- **MongoDB** - NoSQL veritabanı
- **Clean Architecture** - Katmanlı mimari

### Frontend

- **React 18** - Modern React hooks ve components
- **TypeScript** - Type-safe development
- **Vite** - Hızlı build tool
- **shadcn/ui** - Modern UI component library
- **Tailwind CSS** - Utility-first CSS framework
- **RTK Query** - Type-safe data fetching
- **Vitest** - Fast unit testing framework

### DevOps & Tooling

- **Turborepo** - Monorepo build system ve cache
- **pnpm** - Hızlı package manager
- **OpenAPI 3.0** - API dokumentasyonu
- **RTK Query CodeGen** - Type-safe API client generation

## 🎯 Özellikler

- ✅ **10 Popüler Film** listesi (TMDB API'den gelen gerçek veriler)
- ✅ **10 Popüler Dizi** listesi (mock data ile)
- ✅ **Responsive tasarım** (mobil-uyumlu)
- ✅ **Type-safe API calls** (RTK Query hooks)
- ✅ **Modern UI** (shadcn/ui components)
- ✅ **Unit testler** (Vitest + Testing Library)
- ✅ **Monorepo cache** (Turborepo optimizasyonu)
- ✅ **Hot reload** (geliştirme modu)

## 🛠️ Kurulum ve Çalıştırma

### Gereksinimler

- Node.js 18+
- pnpm 8+
- Go 1.21+
- MongoDB (isteğe bağlı)

### 1. Dependency Installation

```bash
# Root seviyede tüm bağımlılıkları yükle
pnpm install
```

### 2. API Client Generation

```bash
# OpenAPI schema'dan RTK Query client oluştur
pnpm generate:api
```

### 3. Development Mode

```bash
# Tüm uygulamaları geliştirme modunda başlat
pnpm dev

# Sadece frontend'i başlat
cd apps/frontend && pnpm dev

# Sadece backend'i başlat
cd apps/backend && go run main.go
```

### 4. Production Build

```bash
# Tüm projeyi build et (Turborepo cache ile)
pnpm build

# Sadece frontend'i build et
cd apps/frontend && pnpm build
```

### 5. Testing

```bash
# Tüm testleri çalıştır
pnpm test

# Frontend testlerini çalıştır
cd apps/frontend && pnpm test

# Test UI'ı aç
cd apps/frontend && pnpm test:ui
```

## 📱 Demo Ekranları

### Ana Sayfa - Film Listesi

- Modern card layout ile 10 popüler film
- Film posteri, IMDB puanı, çıkış yılı
- Hover efektleri ve animasyonlar
- Responsive grid layout

### Dizi Listesi

- 10 popüler dizi mock datası
- Sezon/bölüm sayıları
- Durum bilgisi (Devam ediyor/Bitti)
- TMDB poster görselleri

### UI/UX Özellikleri

- **Dark/Light mode** desteği
- **Loading skeletons**
- **Error handling**
- **Mobile-first** tasarım
- **Accessibility** standartları

## 🔧 Geliştirme Komutları

```bash
# Kod formatlama
pnpm format

# Lint kontrolü
pnpm lint

# Type checking
pnpm type-check

# Cache temizleme
pnpm clean

# API client yeniden oluştur
pnpm generate:api
```

## 📦 Package Scripts

### Root Level (Turborepo)

- `pnpm dev` - Tüm uygulamaları geliştirme modunda başlat
- `pnpm build` - Tüm projeyi build et
- `pnpm test` - Tüm testleri çalıştır
- `pnpm lint` - Tüm projeyi lint et

### Frontend

- `pnpm dev` - Development server (http://localhost:3000)
- `pnpm build` - Production build
- `pnpm test` - Vitest testlerini çalıştır
- `pnpm test:ui` - Test UI'ını aç

### Backend

- `go run main.go` - Server başlat (http://localhost:8080)
- `go build` - Binary oluştur
- `go test` - Go testlerini çalıştır

## 🎨 UI Component Library

Projede shadcn/ui kullanılmıştır:

- `Button` - Çeşitli varyantlarda butonlar
- `Card` - İçerik kartları
- `Badge` - Durum göstergeleri
- `Loading Skeletons` - Yükleme animasyonları

## 🔄 API Integration

### RTK Query Hooks

```typescript
// Film listesi
const { data, isLoading, error } = useGetPopularMoviesQuery({ page: 1 });

// Film detayı
const { data: movie } = useGetMovieByTmdbIdQuery({ tmdbId: 123 });

// Film arama
const { data: results } = useSearchMoviesQuery({ query: 'Batman', page: 1 });
```

### Type Safety

Tüm API calls TypeScript ile type-safe:

```typescript
interface Movie {
  id: string;
  title: string;
  overview?: string;
  voteAverage?: number;
  releaseDate?: string;
  posterPath?: string;
  // ...diğer alanlar
}
```

## 🚀 Deployment

### Frontend (Vercel/Netlify)

```bash
cd apps/frontend
pnpm build
# dist/ klasörünü deploy et
```

### Backend (Docker)

```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

## 📈 Performance

- **Turborepo Cache** - Değişmeyen packageleri cache'ler
- **Code Splitting** - Vite ile otomatik bundle splitting
- **Image Optimization** - TMDB API'den optimize poster görselleri
- **Tree Shaking** - Kullanılmayan kodları bundle'a dahil etmez

## 🎯 Gelecek Özellikler

- [ ] Film/dizi favorilere ekleme
- [ ] Kullanıcı giriş sistemi
- [ ] Arama ve filtreleme
- [ ] Detay sayfaları
- [ ] Watchlist özelliği
- [ ] MongoDB tam entegrasyonu
- [ ] Docker compose setup
- [ ] E2E testler (Playwright)

## 🤝 Katkıda Bulunma

1. Repository'yi fork edin
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add amazing feature'`)
4. Branch'i push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## 📄 Lisans

Bu proje MIT lisansı altında lisanslanmıştır.

---

**Demo Adresi:** http://localhost:3000
**API Endpoint:** http://localhost:8080/api/v1
**API Docs:** http://localhost:8080/swagger
