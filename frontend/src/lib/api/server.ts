import "server-only";
import { getApiBaseUrl } from "./config";
import { createApiTransport } from "./transport";
import type { ApiRequest, AuthHeadersProvider } from "./types";
/** Construct per server request. Never cache user credentials in a module singleton. */
export function createServerApiClient(authHeaders?: AuthHeadersProvider) {
  const transport = createApiTransport({
    baseUrl: getApiBaseUrl(),
    authHeaders,
  });
  return {
    request: <T>(path: string, options: ApiRequest<T>) =>
      transport.request(path, { ...options, cache: "no-store" }),
  };
}
