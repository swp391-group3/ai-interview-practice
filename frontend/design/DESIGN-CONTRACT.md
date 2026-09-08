# RoleCue Design Contract

This file is the entry point for agents implementing or materially changing RoleCue UI.

## Source priority

1. Explicit current product requirements
2. Approved RoleCue landing artifact
3. Approved OpenDesign reference study
4. RoleCue brand system
5. Existing implementation

When sources visibly disagree, follow the highest-priority source.

## Landing canonical references

All paths are relative to `frontend/`.

### Approved RoleCue landing

- **Final landing HTML:** `design/exploration/datn/datn-landing.html`
- **Design system & handoff notes:** `design/exploration/datn/DATN-DESIGN-SYSTEM.md`
- **Local landing assets:** `design/exploration/datn/assets/` (approved typography fonts)

### Hero media direction

- The former DATN hero renders (`datn-hero-16x9.webp`, `datn-hero-3x2.webp`, and `datn-hero-4x5.webp`) are intentionally obsolete and are not expected to exist in this repository.
- Reimplement the hero as a polished 2D editorial media composition. Match the approved OpenDesign grammar through typography, environmental treatment, construction geometry, framing, and restrained motion rather than restoring or replacing those renders.
- Realtime 3D, Blender-derived renders, Three.js, R3F, GLB, and WebGL are deferred and must not be introduced for this purpose.

### Visual grammar

- **Reference reconstruction HTML:** `design/exploration/open-design-study/reference-study.html`
- **Visual reference audit:** `design/exploration/open-design-study/REFERENCE-AUDIT.md`
- **Reference design tokens:** `design/exploration/open-design-study/TOKENS.css`
- **Exploration notes & clone reports:** `design/exploration/open-design-study/NOTES.md`, `design/exploration/open-design-study/CLONE_REPORT.md`, `design/exploration/open-design-study/CLONE_AUDIT.md`

### Motion

- **Motion contract & choreography:** `design/exploration/open-design-study/MOTION-AUDIT.md`

### Brand

- **Brand identity decisions:** `design/brand/BRAND-DECISION.md`
- **Logo system specification:** `design/brand/LOGO-SYSTEM.md`
- **Brand preview:** `design/brand/rolecue-brand-preview.html`
- **SVG vector assets:**
  - Full logo: `design/brand/rolecue-logo.svg`
  - Icon mark: `design/brand/rolecue-icon.svg`
  - Wordmark: `design/brand/rolecue-wordmark.svg`

## Implementation contract

- OpenDesign output is the visual/interaction source of truth for the landing.
- Do NOT blindly paste generated HTML/CSS/JS into production.
- Reimplement the approved design using the real Next.js / React / Tailwind architecture.
- Preserve:
  - composition
  - typography
  - spacing
  - hierarchy
  - responsive behavior
  - motion
  - interaction intent
  - environmental washes
  - borders/shadows
  - media proportions
- Do not invent a new visual direction unless explicitly requested.
- RoleCue branding must replace any old DATN/SEP490 placeholder branding.
- Treat obsolete image references inside the historical DATN landing HTML as conceptual visual guidance only; they are not production asset dependencies.

## 3D status

Realtime 3D / Blender is currently DEFERRED and is NOT part of the landing design contract.

Agents must not introduce Three.js / R3F / GLB merely because earlier exploration considered it.

If 3D is reintroduced later, it requires a new explicit design decision.

## Validation contract

Future landing implementation should be visually reviewed against the approved OpenDesign artifact at minimum at:

- 1440 desktop
- 1280 laptop
- 390 mobile

Implementation should verify:

- typography
- line breaks
- spacing
- section rhythm
- responsive layout
- navigation behavior
- motion timing
- visual density
