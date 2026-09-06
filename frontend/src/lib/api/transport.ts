import { ApiConfigurationError, ApiError } from "./errors";
import type { ApiClient, AuthHeadersProvider, Fetcher } from "./types";
export function createApiTransport(config: {
  baseUrl?: string;
  authHeaders?: AuthHeadersProvider;
  fetcher?: Fetcher;
}): ApiClient {
  return {
    async request(path, { decode, headers, ...options }) {
      if (!config.baseUrl) throw new ApiConfigurationError();
      // Relative paths only: credentials must never be forwarded to a supplied origin.
      if (!path.startsWith("/") || path.startsWith("//") || path.includes("\\"))
        throw new TypeError("API path must start with a single slash.");
      const base = new URL(config.baseUrl);
      const url = new URL(base.pathname.replace(/\/$/, "") + path, base.origin);
      if (url.origin !== base.origin)
        throw new TypeError("API path must stay on the configured origin.");
      const requestHeaders = new Headers(await config.authHeaders?.());
      new Headers(headers).forEach((value, key) =>
        requestHeaders.set(key, value),
      );
      const response = await (config.fetcher ?? fetch)(url, {
        ...options,
        headers: requestHeaders,
      });
      const text = response.status === 204 ? "" : await response.text();
      let body: unknown = text || undefined;
      if (text && response.headers.get("content-type")?.includes("json")) {
        try {
          body = JSON.parse(text);
        } catch {
          if (response.ok)
            throw new ApiError(
              "Invalid JSON response.",
              response.status,
              undefined,
            );
        }
      }
      if (!response.ok)
        throw new ApiError(
          `Request failed (${response.status}).`,
          response.status,
          body,
        );
      return decode(body);
    },
  };
}
