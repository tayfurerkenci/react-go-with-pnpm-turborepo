/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['@ui/components', '@api/client', '@shared/utils'],
  images: {
    domains: ['image.tmdb.org'],
  },
  experimental: {
    externalDir: true,
  },
};

module.exports = nextConfig;
