# Frontend agent instructions

- Work only inside frontend/ unless explicitly authorized otherwise. Preserve existing work; never initialize frontend/frontend.
- Read SKILL.md for the detailed frontend engineering playbook.
- Keep app pages focused on composition; Server Components are the default.
- Own business code inside features; shared UI cannot import features.
- Server state: TanStack Query. Forms: React Hook Form. Navigation: URL. Local UI: React. Cross-route client state: narrowly scoped Zustand only when needed.
- Use the canonical lib/api transport and domain key factories. No arbitrary token reads or competing HTTP clients.
- Keep the interview machine pure and browser runtimes behind client boundaries.
- Never invent auth, backend, AI, billing, socket or lip-sync contracts.
- Run lint, typecheck, bun test and build; discover Playwright tests and run smoke if a browser is installed.
- Do not add DESIGN.md, heavy media, empty directories or broad lint/type suppressions.
- Confirm git status/diff contains no changes outside frontend/ before handing off.
