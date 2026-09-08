# OpenDesign landing page — reference audit

Observed from the live public home page at `https://open-design.ai/` on 2026-09-08. This is a design-study record, not a publishing specification. Values are computed where the recon probe exposed them and visually estimated where noted.

## Layout

- Desktop capture: 1440 × 13,108 px; the content viewport measured 1,425 px because of the browser scrollbar.
- The root content container is `max-width: 1360px`, with `64px` horizontal padding on desktop. Several signature sections deliberately narrow to roughly 960–1080 px or use 238 px page insets.
- The page contains 13 major structural regions: campaign band, navigation, hero, explanation, process, showcase, agents, community, proof, closing statement, newsletter, FAQ, and footer.
- Alignment alternates between centered editorial statements and asymmetric two-column moments. The page avoids a repeated feature-card grid as its default syntax.
- Vertical pacing is unusually spacious. Major sections range from approximately 766 to 2,548 px; the explanation section is about 2,244 px and the process section about 2,548 px.

## Typography

- Loaded family: Albert Sans variable, normal and italic, self-hosted. The page does not use a developer-terminal typographic voice.
- Hero headline: heavy sans, approximately 54–64 px for the lead line at desktop, with a 34 px secondary statement. The composition is centered and tightly tracked.
- Large section statements range visually from about 42–66 px, while scroll-reveal body statements sit near 26 px with 1.35 line-height.
- Body copy is 15–18 px with relaxed line-height. Labels and pills are 11–13 px, semibold, and compact.
- Hierarchy comes from scale and isolation rather than multiple font families, although italic type is used sparingly as emphasis.

## Palette

- Computed canvas: `rgb(250,250,250)` / `#fafafa`.
- Computed text: `rgb(38,38,38)` / `#262626`; secondary inks include `#434343`, `#595959`, and `#8c8c8c`.
- Computed dividers: `#d9d9d9`, `#f0f0f0`, and `#f5f5f5`.
- Primary accent: highly luminous green (`#63fe13` observed in root variables), used for campaign emphasis, selection states, highlight strokes, and data figures.
- The mostly neutral canvas is periodically energized by colorful media. Around the hero, faint cool-blue and pink environmental washes create depth without tinting every surface.

## Spacing

- Base spacing is a fine-grained 2–68 px scale, but layout rhythm is governed by much larger editorial intervals.
- Desktop navigation uses 22 px vertical padding; hero copy begins after a responsive top offset and is capped near 1,080 px.
- The hero title box uses roughly 16–28 px vertical and 18–48 px horizontal padding. Pills are separated by 8 px; the CTA row sits about 42–44 px lower.
- Dominant media follows the CTA row with approximately 48 px separation.
- Quiet whitespace is structural: several 700–1,200 px breathing intervals delay the next statement until the previous visual has settled.

## Surfaces

- Most sections sit directly on the shared near-white canvas. Depth is introduced selectively through white media windows, dark proof containers, and translucent navigation glass.
- Media frames use white or dark stage surfaces rather than generic cards.
- The post-scroll navigation becomes a floating capsule with translucent white fill, backdrop blur, subtle inset highlights, and a light layered shadow.

## Borders

- Borders are consistently 1 px and low contrast.
- Hero media uses `#d9d9d9`; secondary controls use dark ink around 18–20% opacity.
- Selection and construction-line moments use the green accent at full strength.
- FAQ rows use fine full-width dividers; circular toggles repeat the same restrained border treatment.

## Shadows

- Hero and major media: broad, diffused shadow around `0 30px 80px -40px rgba(38,38,38,.28)`.
- Buttons gain a compact lifted shadow on hover rather than a strong glow.
- Navigation glass uses multiple inset highlights plus low-opacity ambient shadows. Shadows remain neutral, never tinted green.

## Background Geometry

- The hero contains faint grids, axis-like lines, and oversized construction diagrams. Their role is spatial: they make the white field feel designed without competing with the headline.
- Thin green selection guides and small square handles frame the hero copy.
- Sparse dimensional objects and soft color washes sit near viewport edges, partially cropped.
- Later sections generally reduce geometry, allowing the featured media to carry visual intensity.

## Media Treatment

- Hero media is the dominant object: about 1,080 px wide at desktop, 16:9, 16 px radius, 1 px border, and a soft deep shadow.
- Supporting product media uses similarly rounded frames and natural aspect ratios, often with annotation panels inside rather than overlapping outside the image.
- Showcase and proof sections shift to more saturated visual stages, followed by quiet white space.
- The study intentionally replaces proprietary photography, logos, videos, and branded artwork with original CSS-built abstract compositions of comparable scale and weight.

## Navigation

- Campaign strip: 44 px high. Main navigation region: about 85 px high in the initial state.
- Desktop structure is a three-part grid: wordmark left, primary links centered, social proof and primary action right.
- The live logo area is about 108 × 40 px. The reconstruction uses a neutral wordmark with the same footprint.
- Initial nav is transparent. After scroll it condenses into a centered pill approximately `min(1280px, 100% - 32px)` with 9 px vertical padding.
- Dropdowns open below their triggers and use opacity/translation rather than theatrical motion.
- Mobile replaces the center and right clusters with one compact menu control while retaining the primary CTA.

## Section Rhythm

1. Campaign band and transparent navigation.
2. Centered, bordered hero statement with pill metadata and two-level CTA hierarchy.
3. Dominant 16:9 media stage.
4. Quiet gap followed by a long explanation and tabbed media proof.
5. Very large pause leading into a process statement and asymmetric step/media arrangement.
6. Centered artifact statement with a large visual.
7. Agents and community sections become more graphic and asymmetric.
8. Proof condenses into a dark stage of six saturated tiles.
9. Closing product statement returns to one dominant media window.
10. Newsletter, split FAQ, quiet multi-column footer, oversized closing wordmark.

## Responsive Behavior

- At 390 px, the hero remains centered; it does not become a left-aligned dashboard hero.
- Display type scales down to roughly 28–36 px, tags wrap into multiple rows, and the CTA group stacks/wraps.
- Desktop two-column structures collapse to one column. Content order becomes statement → controls → media.
- Proof tiles collapse from 3 × 2 to a single vertical list.
- Footer columns wrap into a denser two-column field.
- Horizontal padding falls to roughly 20 px. Media stays full-width with preserved aspect ratio and no horizontal overflow.

