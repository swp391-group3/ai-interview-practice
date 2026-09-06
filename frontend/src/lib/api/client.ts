"use client";
import { getApiBaseUrl } from "./config";
import { createApiTransport } from "./transport";
import type { AuthHeadersProvider } from "./types";
/** Create at the application auth integration boundary, never inside individual components. */
export function createBrowserApiClient(authHeaders?: AuthHeadersProvider) {
  const transport = createApiTransport({
    baseUrl: getApiBaseUrl(),
    authHeaders,
  });
  return {
    request: <T>(path: string, options: import("./types").ApiRequest<T>) =>
      transport.request(path, { credentials: "include", ...options }),
  };
}
