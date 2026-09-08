# OpenDesign landing reference study

## Source information

- Original URL: https://open-design.ai/
- Source repository: no credible public landing-page source repository found in the initial GitHub search.
- License: the public web page and its brand assets are treated as all-rights-reserved for this study; the study must not be publicly redeployed as an OpenDesign-branded replica.
- Reference date: 2026-09-08.

## Technical profile

- Runtime signals: static/server-rendered HTML with 8 scripts, no React/Vue/Next/Astro runtime signal, 2 canvas elements, and 1 video.
- Loaded fonts: self-hosted Albert Sans normal/italic and Remix Icon.
- Motion signals: six sticky/fixed elements; no Lenis, GSAP, smooth-scroll, or scroll-snap signal.

## Pre-clone assessment

- Complexity: L4, animation-heavy brand landing page.
- Mode: visual reconstruction study.
- High-fidelity scope: macro layout, spacing, typography, control geometry, media scale, section cadence, background construction language, and restrained interaction timing.
- Approximate/replaced scope: all logo artwork, trademark assets, exact copy, photography, videos, illustrations, data claims, contributor portraits, and agent marks.
- Not cloned: real downloads, analytics, newsletter backend, external navigation, production tracking, proprietary media, or brand identity.
- Primary risks: confusing the study for a deployable branded replica; replacing branded imagery without preserving its visual weight; losing the page's unusually generous scroll rhythm.

## Run locally

```bash
cd design/exploration/open-design-study
python3 -m http.server 8123
```

Open `http://127.0.0.1:8123/reference-study.html`.

## Content boundary

- Neutral wordmark: `Reference Study`.
- Original marketing claims: replaced with neutral design-study copy.
- Proprietary imagery: replaced with CSS-built abstract compositions.
- Live-site imagery captured under `assets/` is retained only as private recon evidence and is not referenced by the final study page.

## Validation record

- Live desktop, tablet, and mobile screenshots captured under `RECON/screenshots/`.
- Live route map, interaction evidence, network record, source-map hunt, and asset manifest captured under `RECON/`.
- Local page rendered at 360, 390, 430, 600, 768, 820, 1024, 1366, 1440, and 1920 px with no console or page errors.
- No captured section exceeded the available content viewport at any tested width.
- Desktop scroll height: 13,193 px; live reference: 13,108 px (0.65% difference).
- Mobile 390 px scroll height: 10,372 px; live reference: 10,933 px (5.1% difference).
- Interaction probe: 15 of 22 safe actions changed visible state, URL, or scroll position; dropdowns, tabs, FAQ, navigation, and CTA paths are active.
- Strict clone audit: passed all font, local-asset, palette, and tracking gates. The remaining brand-name and source-URL findings occur only in the required reference documentation, not in the runnable study page.
- Visual pixel-diff rate: 23.6%. This number is expected to stay material because the task explicitly replaces all photographs, video frames, branded art, logos, and copy while preserving their geometry and visual weight.

## Study score

- Source evidence: 4/5 — full runtime recon, source CSS inspection, screenshots, route crawl, interaction probe, and asset capture; no public source repository identified.
- Structural fidelity: 4/5 — comparable region progression and desktop/mobile scroll length; reference sub-navigation routes intentionally excluded.
- Visual grammar fidelity: 4/5 — measured container, hero, media, rhythm, typography, palette, border, and shadow behavior; branded media replaced.
- Motion fidelity: 4/5 — condensed navigation, dropdowns, hover lift, reveals, word progression, depth drift, tabs, and accordion recreated without heavy runtime libraries.
- Responsive fidelity: 5/5 — ten widths captured from 360 to 1920 px with no console/page errors or horizontal section overflow.
- Functional completeness: 4/5 — all study interactions work; live download, newsletter backend, analytics, and external routes intentionally absent.
- Content neutralization: 5/5 — runnable study contains no OpenDesign wordmark, marketing copy, photographs, video, branded art, product names, contributor portraits, or agent marks.
- Deployment safety: 4/5 — self-contained local study with no tracking or remote dependencies; documentation retains the reference name and URL for provenance.
