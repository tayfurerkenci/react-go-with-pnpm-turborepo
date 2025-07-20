# 🎬 Movies & TV Shows Platform

Go + MongoDB + Gin framework ile modern bir **film ve dizi platformu**. **Clean Architecture + TMDB API + pnpm + Turborepo** monorepo yapısı.

## 🏗️ Monorepo Yapısı

```
movies-platform/
├── apps/
│   ├── frontend/         # React + Vite + TypeScript
│   └── backend/          # Go + MongoDB + Gin (Clean Architecture)
│       ├── internal/
│       │   ├── domain/           # Entities & Interfaces
│       │   ├── usecase/          # Business Logic
│       │   ├── infrastructure/   # External Services & Database
│       │   ├── delivery/         # HTTP Handlers
│       │   └── middleware/       # Gin Middlewares
├── packages/
│   ├── api-client/       # OAS'den RTK Query client generator
│   ├── oas/              # OpenAPI schema tanımları
│   └── shared/           # Ortak types, utils ve constants
├── turbo.json            # Turborepo pipeline konfigürasyonu
├── tsconfig.base.json    # Base TypeScript konfigürasyonu
└── pnpm-workspace.yaml   # pnpm workspace tanımı
```

## 🚀 Teknoloji Stack'i

| Katman           | Teknoloji                 | Açıklama                     |
| ---------------- | ------------------------- | ---------------------------- |
| **Backend**      | Go + MongoDB + Gin        | Clean Architecture, TMDB API |
| **Frontend**     | React + Vite + TypeScript | Modern SPA                   |
| **API Schema**   | OpenAPI 3.0               | Merkezi API tanımı           |
| **API Client**   | RTK Query                 | Otomatik tip güvenli client  |
| **External API** | TMDB API                  | Film & dizi verileri         |
| **Architecture** | Clean Architecture        | Domain-driven design         |
| **Monorepo**     | pnpm + Turborepo          | Hızlı workspace yönetimi     |
| **Styling**      | Tailwind CSS              | Utility-first CSS            |

## 🎯 Platform Özellikleri

### 🎬 Movies API

- ✅ **TMDB entegrasyonu**: Gerçek film verileri
- ✅ **Search**: Film arama ve filtreleme
- ✅ **Popular**: Popüler filmler
- ✅ **Genres**: Türe göre filtreleme
- ✅ **Cache**: MongoDB'de local cache

### 📺 TV Shows API (Coming Soon)

- 🔄 TV show arama ve listeme
- 🔄 Season & episode detayları
- 🔄 Trending TV shows

### 👤 User Features (Planned)

- 🔄 Watchlist (İzleme listesi)
- 🔄 Ratings & Reviews
- 🔄 Favorite genres
- 🔄 User profiles

## 🔥 Turborepo + Clean Architecture Benefits

### ⚡ Turborepo Avantajları

- **Intelligent Caching**: Build output'ları cache'lenir
- **Parallel Execution**: Task'lar paralel çalışır
- **Dependency Graph**: Akıllı task orchestration
- **Remote Caching**: Team genelinde hızlı build'ler
- **Pipeline Configuration**: `turbo.json` ile task definitions

### 📦 pnpm Workspace Benefits

- **Symlink Strategy**: Disk alanından tasarruf
- **Fast Installation**: npm'den 2x daha hızlı
- **Strict Dependencies**: Phantom dependencies önlenir
- **Workspace Protocols**: Inter-package dependencies

### 🏛️ Clean Architecture Layers

#### Domain Layer (Core Business)

- **Entities**: Movie, TVShow, User, Watchlist
- **Repositories**: Data access interfaces
- **Services**: External API interfaces

#### Use Case Layer (Business Logic)

- **Business Logic**: Film arama, cache management
- **Error Handling**: Consistent error responses
- **External API**: TMDB service calls

#### Infrastructure Layer (External Concerns)

- **TMDB Service**: External API implementation
- **MongoDB Repos**: Database operations
- **HTTP Client**: API calls with timeouts

#### Delivery Layer (HTTP Interface)

- **Gin Handlers**: HTTP endpoints
- **Middlewares**: CORS, Error handling
- **Swagger Docs**: Auto-generated API docs

## 🚀 Turborepo Pipeline

`turbo.json` konfigürasyonu ile optimize edilmiş task pipeline:

```json
{
  "pipeline": {
    "build": {
      "dependsOn": ["^build"],
      "outputs": ["dist/**", "bin/**"]
    },
    "dev": {
      "cache": false,
      "persistent": true
    },
    "lint": {
      "outputs": []
    },
    "type-check": {
      "dependsOn": ["^build"]
    }
  }
}
```

### 🏃‍♂️ Pipeline Commands

```bash
# Tüm workspace'leri paralel build
turbo build

# Sadece değişen paketleri rebuild
turbo build --filter=...HEAD

# Cache'i temizle
turbo clean

# Pipeline görselleştir
turbo graph
```

## 🚀 Hızlı Başlangıç

### Gereksinimler

- Node.js >= 18
- pnpm >= 8
- Go >= 1.21
- MongoDB (local veya Atlas)

### Kurulum

1. **Bağımlılıkları yükleyin:**

```bash
pnpm install
```

2. **Backend environment'ı ayarlayın:**

```bash
cd apps/backend
cp .env.example .env
# .env dosyasını MongoDB URI ile güncelleyin
```

3. **API Client'i generate edin:**

```bash
pnpm generate:api
```

4. **Tüm projeyi build edin:**

```bash
pnpm build
```

5. **Development modunda çalıştırın:**

```bash
pnpm dev
```

## 📝 Komutlar

### Root level komutlar

```bash
# Tüm projeyi Turborepo ile build et
pnpm build
# veya
turbo build

# Development mode (frontend + backend paralel)
pnpm dev
# veya
turbo dev

# API client'i yeniden generate et
pnpm generate:api

# Linting (tüm workspace'ler)
pnpm lint
# veya
turbo lint

# Type checking (dependency-aware)
pnpm type-check
# veya
turbo type-check

# Test'leri çalıştır
pnpm test
# veya
turbo test

# Format kod (tüm workspace'ler)
pnpm format

# Sadece değişen paketleri rebuild
turbo build --filter=...HEAD

# Belirli workspace'i build et
turbo build --filter=backend
turbo build --filter=frontend

# Cache durumunu göster
turbo info

# Pipeline dependency graph'ini göster
turbo graph
```

### Backend komutları

```bash
cd apps/backend

# Development server
pnpm dev

# Build
pnpm build

# Go dependencies
pnpm tidy
```

### Frontend komutları

```bash
cd apps/frontend

# Development server
pnpm dev

# Build
pnpm build

# Preview
pnpm preview
```

## 🔧 Monorepo Geliştirme Workflow'u

### 1. OpenAPI Schema Update

```bash
# packages/oas/openapi.yaml dosyasını güncelleyin
cd packages/oas
# Schema değişikliklerini yapın
```

### 2. API Client Generation

```bash
# Root'tan çalıştırın - Turborepo dependency'leri handle eder
pnpm generate:api
# veya
turbo generate:api
```

### 3. Backend Implementation

```bash
# Backend'de yeni endpoint'leri implement edin
cd apps/backend
pnpm dev  # Hot reload ile development
```

### 4. Frontend Development

```bash
# Frontend'de yeni RTK Query hooks'ları kullanın
cd apps/frontend
pnpm dev  # Vite hot reload
```

### 5. Cross-workspace Testing

```bash
# Root'tan tüm test'leri çalıştırın
turbo test

# Sadece değişen paketleri test edin
turbo test --filter=...HEAD
```

## 📂 Monorepo Paket Detayları

### `@oas/schema`

OpenAPI 3.0 schema tanımları. Tüm API endpoint'leri buradan yönetilir:

- **Centralized API Design**: Tek source of truth
- **Cross-platform Compatibility**: Backend + Frontend sync
- **Auto-documentation**: Swagger UI generation

### `@api/client`

OpenAPI schema'dan otomatik generate edilen RTK Query client:

- **Type-safe API calls**: TypeScript interfaces
- **Automatic caching**: React Query optimizations
- **React hooks for all endpoints**: useGetMoviesQuery vs.
- **Turborepo Integration**: Build pipeline'da auto-generate

### `@shared/utils`

Ortak utilities, types ve constants:

- **Cross-app Sharing**: Frontend + Backend ortak kod
- **Type Definitions**: API interfaces
- **Form validation helpers**: Reusable validation logic
- **Date/string utilities**: Common formatters
- **Constants**: API endpoints, error codes

### `apps/backend`

Go Clean Architecture backend:

- **Independent Deployment**: Container-ready
- **Clean Dependencies**: Domain-driven design
- **MongoDB Integration**: Local caching layer
- **TMDB Service**: External API integration

### `apps/frontend`

React + Vite frontend:

- **Modern Tooling**: Lightning-fast HMR
- **Type Safety**: Full TypeScript integration
- **Styled Components**: Tailwind CSS utilities
- **API Integration**: RTK Query hooks

## 🔄 Turborepo Cache Strategy

### Build Caching

```bash
# İlk build - cache miss
turbo build

# İkinci build - cache hit (instant)
turbo build

# Sadece değişen dosyalar rebuild
turbo build --filter=...HEAD
```

### Remote Caching (Team Benefits)

```bash
# Team remote cache setup
turbo login
turbo link

# Artık team üyeleri cache'i paylaşır
turbo build  # Cache hits from teammates
```

## 🗄️ MongoDB Schema

### Users Collection

```typescript
{
  _id: ObjectId,
  email: string,
  firstName: string,
  lastName: string,
  age?: number,
  isActive: boolean,
  createdAt: Date,
  updatedAt: Date
}
```

## 🚀 Production Deployment

### Turborepo Build Pipeline

```bash
# Production build with caching
turbo build --filter=backend --filter=frontend

# Docker multi-stage builds
docker build -f apps/backend/Dockerfile .
docker build -f apps/frontend/Dockerfile .
```

### Backend Deployment

```bash
cd apps/backend
go build -o bin/server .
./bin/server
```

### Frontend Deployment

```bash
cd apps/frontend
turbo build --filter=frontend
# dist/ klasörünü static host'a deploy edin
```

### Monorepo CI/CD Benefits

- **Intelligent Builds**: Sadece değişen app'ler build olur
- **Parallel Deployments**: Backend + Frontend paralel deploy
- **Shared Dependencies**: Package version consistency
- **Cache Optimization**: Build times 10x daha hızlı

## 📖 API Documentation

API dokümantasyonu OpenAPI schema'dan otomatik generate edilir. Development server çalışırken:

- API Base URL: `http://localhost:8080/api/v1`
- Health Check: `GET /api/v1/health`

## 🤝 Contributing

1. Feature branch oluşturun
2. OpenAPI schema değişikliklerini kaydedin
3. `pnpm generate:api` çalıştırın
4. Backend ve frontend'i güncelleyin
5. Test'leri çalıştırın
6. PR oluşturun

## 📝 License

MIT
