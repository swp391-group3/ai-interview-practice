"use client";
import {
  QueryClient,
  QueryClientProvider,
  isServer,
} from "@tanstack/react-query";
import { lazy, Suspense } from "react";
const Devtools =
  process.env.NODE_ENV === "development"
    ? lazy(() =>
        import("@tanstack/react-query-devtools").then((module) => ({
          default: module.ReactQueryDevtools,
        })),
      )
    : null;
function createQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: { staleTime: 60_000, refetchOnWindowFocus: false, retry: 1 },
      mutations: { retry: 0 },
    },
  });
}
let browserQueryClient: QueryClient | undefined;
function getQueryClient() {
  if (isServer) return createQueryClient();
  return (browserQueryClient ??= createQueryClient());
}
export function QueryProvider({ children }: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={getQueryClient()}>
      {children}
      {Devtools && (
        <Suspense fallback={null}>
          <Devtools initialIsOpen={false} />
        </Suspense>
      )}
    </QueryClientProvider>
  );
}
