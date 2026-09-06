import { describe, expect, it, vi } from "vitest";
import { z } from "zod";
import { createApiTransport } from "./transport";
import { ApiConfigurationError, ApiError } from "./errors";
describe("API transport", () => {
  it("does not fetch without configuration", async () => {
    const fetcher = vi.fn();
    await expect(
      createApiTransport({ fetcher }).request("/resource", {
        decode: z.unknown().parse,
      }),
    ).rejects.toBeInstanceOf(ApiConfigurationError);
    expect(fetcher).not.toHaveBeenCalled();
  });
  it("decodes responses and injects authority headers", async () => {
    const fetcher = vi.fn<typeof fetch>().mockResolvedValue(
      new Response('{"value":1}', {
        headers: { "content-type": "application/json" },
      }),
    );
    const client = createApiTransport({
      baseUrl: "https://example.test/api",
      fetcher,
      authHeaders: () => ({ "X-Test-Authority": "test-only" }),
    });
    expect(
      await client.request("/resource", {
        decode: z.object({ value: z.number() }).parse,
      }),
    ).toEqual({ value: 1 });
    expect(fetcher.mock.calls[0]?.[0].toString()).toBe(
      "https://example.test/api/resource",
    );
    expect(
      new Headers(fetcher.mock.calls[0]?.[1]?.headers).get("X-Test-Authority"),
    ).toBe("test-only");
  });
  it("normalizes non-JSON HTTP failures", async () => {
    const client = createApiTransport({
      baseUrl: "https://example.test",
      fetcher: vi
        .fn<typeof fetch>()
        .mockResolvedValue(new Response("Unavailable", { status: 503 })),
    });
    await expect(
      client.request("/resource", { decode: z.unknown().parse }),
    ).rejects.toMatchObject({
      name: "ApiError",
      status: 503,
      body: "Unavailable",
    });
  });
  it("normalizes malformed JSON and supports empty responses", async () => {
    const fetcher = vi
      .fn<typeof fetch>()
      .mockResolvedValueOnce(
        new Response("{", { headers: { "content-type": "application/json" } }),
      )
      .mockResolvedValueOnce(new Response(null, { status: 204 }));
    const client = createApiTransport({
      baseUrl: "https://example.test",
      fetcher,
    });
    await expect(
      client.request("/resource", { decode: z.unknown().parse }),
    ).rejects.toBeInstanceOf(ApiError);
    await expect(
      client.request("/resource", { decode: z.undefined().parse }),
    ).resolves.toBeUndefined();
  });
  it("rejects externally supplied origins", async () => {
    const fetcher = vi.fn();
    const client = createApiTransport({
      baseUrl: "https://example.test",
      fetcher,
    });
    await expect(
      client.request("//other.test", { decode: z.unknown().parse }),
    ).rejects.toBeInstanceOf(TypeError);
    expect(fetcher).not.toHaveBeenCalled();
  });
});
