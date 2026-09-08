# RoleCue Figma product design report

## Delivery status

The RoleCue Figma file was reorganized successfully into the requested three-page Starter-plan structure:

- `Foundations` contains the original minimal Cover title block.
- `Marketing` remains the public-design workspace.
- `Product` is the former Cover page, renamed for candidate, admin, auth, runtime, and system Sections.

The first primitive collection, `RoleCue / Primitives`, was created with 15 confirmed color variables. The Starter plan rejected a second variable mode, and the Figma MCP server then rejected all additional calls because the Starter View seat reached its monthly 20-call allowance. No complete screen or component construction can be truthfully claimed until that external limit is lifted.

The accompanying [screen inventory](./SCREEN-INVENTORY.md) is the complete implementation manifest, prepared from the actual route structure and user-approved product scope so Figma work can resume without reconsidering information architecture.

## Final visual thesis

**A bright editorial interface for serious technical interview practice.**

RoleCue should feel like a composed preparation environment, with near-white canvas, near-black typographic hierarchy, generous negative space, and a single restrained cue green. Soft blue and pink environmental washes belong to public and authentication framing; they should not spill into every candidate screen. Candidate work is quiet and focused. Admin has more density, but uses rules, alignment, and typography before resorting to cards.

The Split Halo icon appears as a small contextual signal. The full RoleCue lockup is reserved for marketing, workspace introductions, and high-level account surfaces. There is no DATN, SEP490, or generic AI branding in the planned visible copy.

## Page and Section organization

The three-page plan respects the Starter plan's page cap:

1. **Foundations**: `00 Cover` through `13 Design Notes`, including the components and system-state library.
2. **Marketing**: public landing and pricing/credit-purchase frames at 1440 and 390 widths.
3. **Product**: labeled Sections A–U covering Authentication; candidate shell/dashboard; JD; setup; preflight; interview runtime; evaluation; reports; history; billing; profile; admin; and system states.

All proposed names are in [SCREEN-INVENTORY.md](./SCREEN-INVENTORY.md). Pages are not needed for Admin, Components, Notes, or System States; Sections keep the file coherent within the plan limit.

## Design-system plan

The local Figma system should use:

- **Color:** RoleCue Canvas `#FBF9F3`, Ink `#121814`, Cue green `#428762`, Quiet ink `#3B4540`, Night `#101612`, Mist `#E8ECE7`, plus quiet blue/pink environmental washes and accessible semantic colors.
- **Typography:** Inter in Figma, using Display XL, Display L, H1–H3, Body L, Body, Body S, Label, and Caption. Current frontend CSS is Arial scaffolding; the approved brand guidance takes precedence for design until production typography is intentionally updated later.
- **Spacing:** a 4 px base scale through 96 px. Canonical layouts: 1440 desktop, 1280 laptop, 390 public/auth mobile.
- **Shape and depth:** mostly 8–20 px radius, 1 px quiet borders, two subtle elevated-surface shadows, and one more pronounced marketing media shadow. Surfaces are selective; the product must not become a field of identical cards.
- **Motion:** 160–240 ms small opacity/translate reveals, 240–360 ms media entrances, a condensing public navigation, 120–160 ms control feedback, and clear reduced-motion alternatives. Interview runtime transitions describe state changes without spectacle.

The available Material 3 and Simple Design System libraries were audited but rejected for construction because their token models and visual language conflict with the approved RoleCue direction. The component inventory is deliberately local and compact: buttons, fields, selection controls, navigation, surfaces, table/filter primitives, status/feedback patterns, interviewer/avatar primitives, and report-chart primitives.

## Key UX decisions

### Candidate flow

The candidate experience leads with the next meaningful action, then role context, recent progress, and credits. It avoids metric-card overload. The JD flow follows confirmed steps: paste/upload a description, calm analysis transition, review extraction, refine skills, configure the session, select interviewer/voice, then preflight. Wizard state maps to explicit URLs; design must not imply an opaque client-only flow.

### Authentication

Auth uses an editorial two-zone composition: a quiet brand/copy rail and an expansive form field, with environmental wash contained to the outer canvas. It avoids a small generic white card centered over a gradient. Google sign-in is included where requested, while its backend behavior remains unclaimed. All forms show focus, validation, loading, and disabled states with label/help/error association.

### Interview room

The interview room is desktop/laptop first and quieter than dashboard. It has a defined interviewer media stage for the future 3D avatar, a controlled 2D visual fallback, current question/context, session progress, and only essential controls. The interview state machine directly defines the visual states: connecting, asset loading, interviewer speaking, listening, processing answer, transitioning, paused, reconnecting, degraded 2D, completed, terminated early, and fatal error.

The design must never present temporary avatar art as a production 3D deliverable. Reconnecting language must only promise session safety where the final product contract supports that statement. Mobile should not compress the room into a faux video-conference grid; it should clearly explain the desktop/laptop requirement or show an intentionally degraded guidance experience.

### Reports

Performance reports begin with strengths, then improvement areas, then specific next actions. Score-first or punitive framing is rejected. Charts are reserved for competency comparisons or change over time and must carry a visible interpretation. Labels such as “failed”, “poor”, and “bad” are excluded.

### Billing and profile

Credit purchase is a transparent path from balance to purchase to payment result to transaction history. No payment gateway, persistent package price, or entitlement behavior is assumed where the backend contract is unconfirmed. Unapproved values remain explicit design placeholders. Profile contains personal information, practical preferences, and security controls only.

### Admin

Admin retains RoleCue typography and quiet surfaces while increasing information density. It supports operational concepts present in the routes: users, sessions, domains, question/configuration content, avatars, voices, billing, and settings. Session views show verified operational context and avoid invented interview transcripts or sensitive candidate data.

## Responsive rules

- **1440 desktop:** candidate shell with generous rail and content width; marketing has centered framing and large hero media.
- **1280 laptop:** navigation condenses but remains discoverable; interview runtime remains fully supported.
- **390 mobile:** marketing and auth are redesigned for narrow reading order; candidate shell uses intentional compact navigation. Live interview is not forced into a tiny desktop layout.
- **Touch/accessibility:** controls retain at least 44 px targets where touch is primary; focus is high contrast and visible; status uses labels and icon/text, not color alone.

## Accessibility requirements

All planned foundations include AA-aware color pairings, visible focus treatment, keyboard navigation, descriptive action labels, field help/error association, status text in addition to color, and reduced-motion rules. Permission and preflight design distinguish required, optional, and warning conditions without a surveillance/proctoring tone.

## Known limitations

1. **Figma MCP quota blocker:** the authenticated Starter View seat reached the server’s monthly tool-call limit. It must be upgraded to a Professional-or-higher plan with a Full or Dev seat, or access must be restored with a higher-capacity plan, before Figma design construction can continue.
2. **Starter variable-mode limit:** the selected plan allows one variable mode, so a local Light/Dark semantic collection cannot be created. Dark interview surfaces should stay as documented contextual primitives unless the plan is upgraded.
3. **OpenDesign source missing:** `design/exploration/landing/`, `design/exploration/open-design-study/`, and the requested reference artifacts were empty/unavailable. The plan uses the requested visual grammar and RoleCue brand documentation, but cannot claim direct transfer of an absent landing artifact.
4. **No production implementation changed:** `src/`, package/configuration, tests, and backend contracts remain untouched.

## Implementation handoff recommendations

When Figma access is restored, resume in this order without creating another file or any additional pages:

1. Inspect and clean up any empty `RoleCue / Semantic color` collection that may have been created before the mode-limit error.
2. Complete Foundations Sections 00–13, text/effect styles, and the local component families.
3. Build Marketing desktop/mobile frames using the approved RoleCue lockup and the approved landing media-container reference.
4. Build Product Sections A–U from the exact 80-frame inventory; validate each section with metadata and screenshots.
5. Run a final consistency and accessibility review across Landing, Auth, Dashboard, JD, Setup, Preflight, Interview, Report, Billing, and Admin.
6. Only after human approval, translate the verified screens into Next.js feature work using the existing route and API boundaries.
