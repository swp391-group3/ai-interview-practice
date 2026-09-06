# Interview Practice frontend

Frontend architecture for the AI Virtual Technical Interview Platform. This directory is the Next.js application root inside the Go monorepo.

## Quick start

Use Bun 1.4.0 and Node 22.12+ (Node 26.7.0 was used for verification). From this directory:

```sh
bun install --frozen-lockfile
bun run dev
```

No environment file or running backend is required. The optional NEXT_PUBLIC_API_BASE_URL in .env.example is read only when an API client is constructed. Public variables are build-time values, never secrets.

## Stack

Next.js App Router, React 19, strict TypeScript, Tailwind v4 and shadcn/ui (Radix Nova, neutral, Lucide). TanStack Query owns server state. React Hook Form, Zod and the resolver are installed for upcoming forms. Zustand is available for a justified cross-route draft; no store has been invented. Three.js, React Three Fiber and Drei are ready for the future avatar integration.

See package.json for exact pinned versions and bun.lock for resolved dependencies.

## Commands

| Command                               | Purpose                                               |
| ------------------------------------- | ----------------------------------------------------- |
| bun run dev                           | Development server                                    |
| bun run build                         | Production compilation and prerendering               |
| bun run start                         | Serve the production build                            |
| bun run lint                          | Next.js core web vitals and TypeScript ESLint rules   |
| bun run typecheck                     | Generate route types, then strict TypeScript checking |
| bun test                              | Native Bun bridge that runs the Vitest suite          |
| bun run test                          | Run Vitest directly                                   |
| bun run test:watch                    | Vitest watch mode                                     |
| bun run test:e2e -- --list            | Discover Playwright tests                             |
| bun run test:e2e                      | Smoke test against an existing production build       |
| bun run format / bun run format:check | Format / verify frontend files                        |

Playwright uses port 3100 and an already installed Chromium browser. Run build first. Set PLAYWRIGHT_CHROMIUM_EXECUTABLE_PATH when reusing an existing compatible Chromium executable instead of Playwright's expected revision. Browser binaries and operating-system dependencies are not installed automatically. The Bun test bridge exists because bun test invokes Bun's own runner rather than the package.json test script; Vitest owns all colocated unit/component tests.

## Source structure

- src/app: Server Component pages, route groups, layouts, redirects and error/loading boundaries.
- src/features: auth, dashboard, job-description, interview-setup, interview, report, history, billing, profile and admin.
- src/components: shared layout, feedback and domain-neutral shadcn primitives.
- src/lib: canonical fetch transport, environment-specific API factories and styling utility.
- src/config: routes, navigation, environment, site metadata and role vocabulary.
- src/providers: per-browser-lifecycle QueryClient and application provider composition.
- tests: Vitest setup, native Bun bridge and Playwright smoke test.

No global types folder or empty public asset directories are needed yet. Feature-specific contracts stay beside their owner.

## Architecture and workflow

Read AGENTS.md and SKILL.md before changing this frontend. Prefer Server Components; introduce client boundaries for interaction or browser runtime ownership only. URL routes determine wizard steps and allow refresh/back/deep links. Step guards and draft persistence are intentionally pending actual draft semantics.

The API factories share one transport and error contract. Every response requires a decoder over unknown; Zod schemas can validate real response contracts when agreed. Auth supplies headers through an adapter. Server clients are created per request and disable caching. No endpoint calls, token storage or response envelope is assumed.

Feature integration follows request function → Query wrapper → orchestration → UI. interviewKeys demonstrates hierarchical key factories without a fabricated request. Do not put navigation or notifications inside transport/query functions.

The interview machine is deterministic TypeScript independent of React and Zustand. Invalid transitions are no-ops, terminal states require RESET, and resume enters synchronization through RECONNECTING. The 2D state is an acknowledgement gate; persistent renderer mode and authoritative session resumption must be designed with the runtime contract.

The avatar stage is a small unmounted client boundary; future callers must dynamically import it with SSR disabled. There are no models or fake lip-sync. Browser capability checks are user-triggered and only test availability; they do not verify permissions, actual devices or connection quality.

## Current status and integration limits

All planned public/auth/candidate/admin route placeholders render without backend services. /interviews/new redirects to its first explicit step; /admin redirects to /admin/dashboard. Dynamic room/report routes display references without inventing session records.

**Candidate and admin layouts are not authorization barriers.** There is no login success behavior, session authority implementation or protected route claim. Implement server/backend authorization before exposing sensitive data.

Auth, AI analysis, interview wire protocol, voice capture/playback, resume reconciliation, evaluation/report data, avatar assets and payment gateway integration remain unimplemented. Billing belongs behind backend/provider boundaries, including credits, packages, checkout, transactions and optional invoices/subscriptions.

There are no product forms until real submission behavior exists; RHF/Zod/resolvers are installed and their integration rules are in SKILL.md. No theme provider, toast provider or tooltip provider is needed by the current UI. Semantic theme variables remain ready for a later visual system. DESIGN.md is intentionally absent.

## References

- [Next.js installation](https://nextjs.org/docs/app/getting-started/installation)
- [shadcn Tailwind v4](https://ui.shadcn.com/docs/tailwind-v4)
