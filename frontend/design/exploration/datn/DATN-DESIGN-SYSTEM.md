# DATN landing page — design system and handoff

## Narrative

DATN is positioned as a working network for people who direct, design, research, and build digital culture. The story moves from a simple assertion—direction is the differentiator—to the practices, programs, network structure, and membership invitation that make that assertion credible.

## Visual system

The page uses an editorial, near-white canvas with dark ink and a single acid-green accent. Albert Sans is used as the shared display and reading face: the display uses a heavy weight, tight tracking, and compact leading; utility labels use the system mono stack. The hero is centered and intentionally airy. Subsequent chapters alternate a quiet centered statement with asymmetric working material.

| Token | Value | Use |
| --- | --- | --- |
| `--bg` | `oklch(98.4% .006 98)` | Page canvas |
| `--surface` | `oklch(100% 0 0)` | Elevated surfaces |
| `--ink` | `oklch(19% .012 265)` | Main type and dark stage |
| `--muted` | `oklch(47% .014 255)` | Supporting copy |
| `--line` | `oklch(86% .011 105)` | Borders and dividers |
| `--acid` | `oklch(84% .23 143)` | Selection points and emphasis |
| `--blue` | `oklch(68% .14 235)` | Internal visual media only |
| `--pink` | `oklch(74% .12 8)` | Internal visual media only |

## Component rules

- The primary CTA uses dark ink with white text. Secondary actions are transparent with an ink border.
- Use the acid accent only for selection geometry, a badge, and contained visual details. Do not add a second page-level accent.
- The condensed navigation becomes a white translucent pill after 64 px of scroll.
- Cards are reserved for the dark program chapter. Elsewhere, borders and spacing define groups.
- Media remains inside a measured 16:9, 3:2, or 4:5 frame. The delivered production render has each responsive crop available in `assets/`.

## Figma handoff

1. Create a 1360 px desktop content frame with 64 px side padding; use 20 px side padding at mobile.
2. Set display text to Albert Sans 820, -7.3% tracking, 0.91 line-height. Large section titles use 790 with -6.4% tracking.
3. Use a 16 px hero radius and 22 px dark-stage radius. Keep all structural borders at 1 px.
4. Preserve the hero crop contract: 16:9 at desktop, 3:2 between 621–1024 px, 4:5 at 620 px and below.
5. Prototype button lift at 3 px over 180 ms and section reveals at 700 ms with a 22 px upward travel. Respect reduced-motion preferences by rendering visible states without transitions.
6. Maintain one primary action per action group. The mobile header reduces to mark plus menu control; the document body stays single-column below 760 px.

## Asset inventory

- `assets/fonts/fonts.css` — local Albert Sans font-face declarations
- `assets/datn-hero-16x9.webp` — 1920 × 1080 desktop hero
- `assets/datn-hero-3x2.webp` — 1440 × 960 tablet hero
- `assets/datn-hero-4x5.webp` — 960 × 1200 mobile hero

The supplied renders depict an interview-practice product environment. They remain credited as the user-provided production asset and should be replaced only when DATN’s final product imagery is ready.
