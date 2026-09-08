# OpenDesign landing page — motion audit

Observed through the live page, automated interaction probes, and source CSS inspection on 2026-09-08. Timings are approximate unless a source value is noted.

## Motion token summary

- Micro feedback: 160–180 ms.
- Control/background transition: 180–260 ms.
- Navigation geometry transition: 360–420 ms.
- Entrance reveal: 600–720 ms.
- Dominant media reveal: 800–900 ms.
- Stagger: about 60–80 ms.
- Reveal travel: 18–28 px.
- Standard easing: `cubic-bezier(.22,.61,.36,1)`.
- Premium reveal easing: `cubic-bezier(.16,1,.3,1)`.
- No bounce, elastic overshoot, or continuous attention-seeking CTA motion.

## Interaction matrix

| Observed interaction | Trigger | Visual effect | Approx. duration | Approx. easing | Implementation recommendation |
|---|---|---|---:|---|---|
| Campaign dismiss | Click close | Strip collapses; page chrome shifts upward | 220–300 ms | standard glide | Animate height and opacity together; remove from tab order after completion. |
| Navigation condense | Scroll beyond first region | Full-width transparent nav becomes a centered blurred pill with smaller logo and shadow | 420 ms | `cubic-bezier(.22,.61,.36,1)` | Toggle one `is-condensed` state; transition width, padding, radius, background, and shadow as a coordinated unit. |
| Dropdown open | Hover, click, or keyboard focus | Menu fades and moves down a few pixels; trigger caret rotates subtly | 180–220 ms | ease-out | Use opacity + `translateY(6px)`; preserve keyboard escape and outside-click dismissal. |
| Primary button hover | Pointer hover | Button rises 2–3 px; shadow expands; foreground contrast is unchanged | 160–180 ms | `cubic-bezier(.23,1,.32,1)` | Move the whole control, never only the label. |
| Secondary pill hover | Pointer hover | Border darkens, surface becomes fully white, control rises about 3 px | 180 ms | ease-out | Pair border/background changes with transform; keep dark text. |
| Hero media hover | Pointer hover | Poster scales to about 1.015; circular play control scales slightly | 160–400 ms | ease | Clip scaling inside the media radius. |
| Section entrance | Intersection at about 15–20% visibility | Opacity 0→1 with translateY 20–24 px→0 | 680 ms | `cubic-bezier(.16,1,.3,1)` | One observer with stagger custom properties; trigger once. |
| Large media entrance | Intersection | Opacity + slightly longer rise/scale settle | 850–900 ms | premium reveal | Delay by about 90 ms after heading; avoid perspective or spring. |
| Explanation text reveal | Scroll through sticky region | Words progress from low opacity (~.18) to full opacity | 90 ms per word update; scroll-linked | linear per word | Map viewport progress to word index; turn off sticky behavior under reduced motion. |
| Process step change | Scroll or direct selection | Active pill darkens; media annotations fade/slide to next state | 260–420 ms | standard glide | Keep one active panel in flow; crossfade details and move by <16 px. |
| Decorative background drift | Idle/scroll | Edge geometry and wash layers translate at different low speeds | 12–20 s idle or 0.03–0.08 scroll ratio | linear / scroll-linked | Move only decorative layers; cap travel around 24 px. |
| Proof tile hover | Pointer hover | Tile rises 4–6 px; image/field scales about 1.02; circular arrow shifts | 220–300 ms | standard glide | Preserve the dark stage; no color inversion. |
| FAQ accordion | Click or keyboard | Answer height and opacity reveal; plus rotates to × | 260–320 ms | standard glide | Animate a grid row or measured height; maintain `aria-expanded`. |
| Mobile menu | Click | Compact panel fades and drops below nav; body remains stable | 220–260 ms | ease-out | Use fixed placement beneath the condensed header and restore focus on close. |

## Where motion is allowed

- Navigation geometry, dropdowns, controls, media framing, scroll entrances, progressive word reveal, FAQ disclosure, and non-content decorative depth.
- Large visual blocks may reveal once. Hover motion should remain under 6 px and under 1.02 scale.

## Where motion remains static

- Body text after it has revealed, numerical proof values, form fields while typing, and all core layout dimensions during ordinary reading.
- No perpetual motion on primary CTA labels and no automatic carousel.

## Reduced-motion behavior

- Disable smooth scrolling, parallax, decorative drift, word-by-word opacity sequencing, and entrance transforms.
- Keep dropdown and accordion state changes effectively immediate while preserving visibility and focus feedback.
- Do not hide any content behind a motion-dependent state.

