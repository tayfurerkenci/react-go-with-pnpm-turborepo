# My Monorepo

Go + MongoDB backend ile React + Vite frontend kullanan modern bir monorepo yapısı.

## 🏗️ Yapı

```
my-monorepo/
├── apps/
│   ├── frontend/         # React + Vite + TypeScript
│   └── backend/          # Go + MongoDB + Gin
├── packages/
│   ├── api-client/       # OAS'den RTK Query client generator
│   ├── oas/              # OpenAPI schema tanımları
│   └── shared/           # Ortak types, utils ve constants
├── turbo.json            # Turborepo pipeline konfigürasyonu
├── tsconfig.base.json    # Base TypeScript konfigürasyonu
└── pnpm-workspace.yaml   # pnpm workspace tanımı
```

## 🚀 Teknoloji Stack'i

| Katman | Teknoloji | Açıklama |
|--------|-----------|----------|
| **Backend** | Go + MongoDB + Gin | Performanslı REST API |
| **Frontend** | React + Vite + TypeScript | Modern SPA |
| **API Schema** | OpenAPI 3.0 | Merkezi API tanımı |
| **API Client** | RTK Query | Otomatik tip güvenli client |
| **Monorepo** | pnpm + Turborepo | Hızlı workspace yönetimi |
| **Styling** | Tailwind CSS | Utility-first CSS |

## 🔥 Özellikler

- ✅ **Type-safe API**: OpenAPI schema'dan otomatik TypeScript tip üretimi
- ✅ **Otomatik RTK Query hooks**: API endpoints için hazır React hooks
- ✅ **Hot reload**: Frontend ve backend için anlık yeniden yükleme
- ✅ **Shared packages**: Ortak types, utils ve constants
- ✅ **MongoDB Atlas ready**: Bulut veritabanı desteği
- ✅ **Modern tooling**: ESLint, Prettier, Turborepo cache
- ✅ **Production ready**: Docker, CI/CD hazır yapı

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
# Tüm projeyi build et
pnpm build

# Development mode (frontend + backend)
pnpm dev

# API client'i yeniden generate et
pnpm generate:api

# Linting
pnpm lint

# Type checking
pnpm type-check

# Test'leri çalıştır
pnpm test

# Format kod
pnpm format
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

## 🔧 API Geliştirme Workflow'u

1. **OpenAPI schema'yı güncelleyin** (`packages/oas/openapi.yaml`)
2. **API client'i yeniden generate edin:**
   ```bash
   pnpm generate:api
   ```
3. **Backend'de yeni endpoint'leri implement edin**
4. **Frontend'de yeni RTK Query hooks'ları kullanın**

## 📂 Paket Detayları

### `@oas/schema`
OpenAPI 3.0 schema tanımları. Tüm API endpoint'leri buradan yönetilir.

### `@api/client`
OpenAPI schema'dan otomatik generate edilen RTK Query client:
- Type-safe API calls
- Automatic caching
- React hooks for all endpoints

### `@shared/utils`
Ortak utilities, types ve constants:
- API types
- Form validation helpers
- Date/string utilities
- Constants

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

### Backend
```bash
cd apps/backend
go build -o bin/server .
./bin/server
```

### Frontend
```bash
cd apps/frontend
pnpm build
# dist/ klasörünü static host'a deploy edin
```

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
