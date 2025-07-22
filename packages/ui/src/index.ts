// Re-export shadcn/ui components for cross-platform use
export { Button, buttonVariants, type ButtonProps } from './components/button';
export {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from './components/card';
export { cn } from './lib/utils';

// Platform-specific exports
export type { ComponentProps, ReactNode } from 'react';
