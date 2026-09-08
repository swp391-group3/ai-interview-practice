# RoleCue logo system

## Concept

RoleCue uses the **Split Halo**: an open focus field plus a compact cue tick. It is a quiet sign of attention, timing, and readiness. The mark should remain abstract; do not add eyes, speech bubbles, microphones, robot parts, radar grids, or AI effects.

## Geometry

- The primary icon is drawn on a 48 × 48 artboard.
- The near-black halo uses a 5.5-unit rounded stroke and a 16-unit radial field around the center.
- The green cue tick uses the same stroke weight, is angled at 45°, and sits in the halo opening.
- The lockup uses a 12-unit gap from the icon artboard to the wordmark. Do not tighten or widen this relationship in normal use.
- The custom wordmark is a 254 × 48 geometric stroke drawing. Use the supplied SVG rather than recreating it with a substitute font.

## Icon logic

The open halo gives the icon a directional moment without becoming a target, camera, or chat symbol. The green tick is the only accent: it should read as a cue entering the field, a turn beginning, or a person becoming ready. The open center preserves calm and makes the icon legible at small sizes.

## Typography guidance

The wordmark is custom artwork and should not be typeset. For future product and Figma work, use an unobtrusive UI sans such as Inter for interface text, with regular, medium, and semibold weights. Keep headings compact and sentence case; let spacing and type scale establish the editorial feel instead of using decorative display fonts.

## Color guidance

| Token | Hex | Use |
| --- | --- | --- |
| Canvas | `#FBF9F3` | Primary near-white field |
| Ink | `#121814` | Primary text and halo |
| Cue green | `#428762` | Cue tick and restrained emphasis |
| Quiet ink | `#3B4540` | Supporting text and secondary marks |
| Night | `#101612` | Dark surfaces |
| Mist | `#E8ECE7` | Low-contrast dividers and fields |

Use green to identify a meaningful cue, action, or status detail. Keep it sparse. Do not use gradients, neon lime, or a large field of saturated green as a default surface.

## Clear space and minimum size

- Keep clear space equal to the halo stroke width (5.5 units on the source artboard) around the icon and around the full lockup.
- Use the full logo at 120 px wide or larger when space allows; do not use the wordmark below that width.
- Use the icon alone below 120 px wide. The preferred minimum is 16 px; use 20 px or larger for interface navigation.
- Do not place the mark inside a tight border, crop the cue tick, stretch it, rotate it, or separate the tick from the halo.

## Light and dark usage

- **Light:** use the supplied near-black halo and green cue tick on Canvas or white.
- **Dark:** use an all-white monochrome mark on Night or another dark neutral. Keep the green version only when the contrast and context remain clear.
- **Single-color:** set both halo and cue tick to one solid ink or one solid white. Do not use a gray cue tick in a monochrome rendition.
- Avoid photo backgrounds and busy patterned surfaces. If necessary, use a quiet solid plate with ample clear space.

## Favicon and app icon

- The supplied `rolecue-icon.svg` is the source for favicon and app-icon exports.
- Export 16, 32, 48, 180, and 512 px raster variants from the master SVG when needed; do not include the wordmark in those exports.
- For a rounded-square app tile, preserve the mark's clear space and use Canvas or Night as the tile color. Do not redraw the halo as an enclosing app border.
- At 16 px, check that the green tick retains contrast. Use the monochrome source treatment when a platform renders color too softly.

## Future product usage

Use the full lockup for the marketing header, Cover page, and high-level settings. Use the icon for compact product navigation, loading states, empty states, and section signatures. Reserve the cue green for a small number of intentional product moments so the brand continues to feel composed during technical interview work.
