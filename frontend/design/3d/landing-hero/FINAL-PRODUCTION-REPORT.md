# Landing Hero Final Production

## Final Visual

Camera A developed into a bright, quiet home-office interview still. The seated candidate occupies the right, the monitor anchors the center-left, and a visible corner-window return provides daylight architecture on the left. An open preparation notebook, résumé sheet, reference book, pen, coffee and headphones retain the career-preparation story. A plant, warm lamp and original abstract print add domestic specificity without a cluttered room.

The PASS 1 file was reopened and preserved. Work continued in `scenes/landing-hero-production.blend`; the deliverable is `scenes/landing-hero-final.blend`. Structural, pose, dressed-character, material, groom, screen and Cycles lighting decisions were checked through actual screenshots or renders during production (intermediate lookdev and blockout test renders were pruned in repository cleanup; the final master render and verified crops are preserved in `renders/final/`).

## Character

One original adult candidate generated with official MPFB2 v2.0.17, using mixed default ancestry macro values and adjusted age, gender, muscle, weight and height. No real-person reference or likeness was used. The core standard rig carries a deliberately asymmetric seated pose. The initial torso lean and wrist roll were corrected after rendered inspection; the visible hand rests over the keyboard instead of presenting its palm sideways. Head rotation turns attention toward the screen.

Core casual shirt/jeans, shoes, eyes, eyebrows and short hair were fitted to the body. Shirt color, sheen and roughness were tuned for subdued contemporary casual styling. The short-hair base received reduced specular/normal intensity, darker color and an original directional curve-strand overlay. The proxy is hidden in the retained working scene, not the final visible candidate.

The body is intended for this medium-distance still. Facial expression is primarily conveyed through head direction and relaxed pose, not a close-up smile. The supplied core wardrobe is adapted rather than bespoke garment simulation. No animation, lip-sync or motion production was started.

## Environment

The existing compact room, desk and storage layout remain. A back-left window return, reveal, mullion, sill and gathered linen mesh make the daylight source visible from Camera A. The exterior is a quiet original sky field. A custom plant with tapered leaf blades replaces oval leaf placeholders. The original abstract wall print and individual background book volumes replace coarse blockout masses.

Furniture gained softened edges, upholstery, chair arm supports and welt detail, storage pulls, a monitor cable, keyboard keys, layered notebook edges and cover, curved headphones, and a ceramic coffee vessel. The monitor grew approximately 10%, remaining a plausible large desktop display rather than a television. The desk remains 1.60 × 0.78 m.

## Product Screen

Original, temporary SVG interview-room artwork shows a neutral illustrated AI interviewer, an active-session label and a restrained response workspace. It contains no employer, metrics, scores, testimonials, real personal details or transcript claims. These labels are illustrative product artwork, not documentation of implemented features.

`assets/props/interview-screen.svg` is the editable source and `interview-screen.png` its raster derivative. A UV-mapped plane in `50_SCREEN` holds the artwork beneath a separate restrained cover-glass layer. Replace the artwork with approved Figma/product output before shipping if product accuracy is required. Screen emission is low; daylight remains the scene's primary light.

## Lighting

Final rendering uses Cycles, with a broad left daylight area, a softer window-return area, reduced interior fill, a restrained warm lamp pool and subtle late-morning sun. World strength is low enough that the room is not lit by flat ambient alone. Exposure was reduced after the first Cycles test appeared too pale. The lamp provides a local warm contrast without heavy orange/blue grading.

128 maximum samples, adaptive threshold 0.018, denoising, seven total bounces and four diffuse bounces are used for the master. AgX color management, exposure −0.55. Rendering uses the available CPU; no GPU was detected. No HDRI, bloom or post-render generative retouch was used.

## Materials

Poly Haven's CC0 Wood Table 001 supplies 2K color, roughness and GL normal maps. The roughness range was raised after the first test appeared too glossy. Longitudinal edge shading replaces visibly stretched edge grain. Plaster, pale floor grain, upholstery and linen have restrained procedural surface variation at separate scales. Powder-coated metal, paper, ceramic and screen glass use distinct material responses rather than one shared clay roughness.

Skin uses a core color map with restrained SSS and moderate roughness; hair and clothing have separate tuned responses. All image textures are packed in the final `.blend`. The work targets landing-scale material read rather than pore-level or fiber-level close-ups.

## Camera

Camera A remains the foundation: left-rear three-quarter direction, 48 mm lens, near seated height. Candidate on the right, monitor around center-left, light architectural space on the left. A focus target between candidate and monitor and f/4.5 depth of field soften distant architecture while keeping the interaction readable. Cameras B and C remain as earlier blockout references, not alternative final masters.

## Asset Sources

Official MakeHuman Community MPFB2 source release and CC0 system pack; Poly Haven Wood Table 001. Exact URLs, providers, acquisition date, selected pack paths and usage are recorded in `ASSET-LEDGER.md`. No other asset marketplace or stock room was used.

## Licensing

MakeHuman core graphics are CC0; MPFB software is GPL. Poly Haven texture maps are CC0. Neither source requires attribution for these graphics, though provenance is retained. The screen and wall artwork are original. No third-party brand, employer, celebrity likeness or unclear-license asset is used.

## Responsive Crops

The 16:9 master is 2560 × 1440. A centered 3:2 derivative is 2160 × 1440. The 4:5 derivative is an intentional offset 1152 × 1440 crop preserving the screen, head, visible keyboard hand and preparation notebook. The portrait trims the room/window, outer chair and part of the résumé; it is an interaction-focused composition and has no left typography-safe field. Use the actual portrait asset rather than relying on centered `object-fit`.

## Known Limitations

The screen is temporary artwork awaiting real product/Figma content. Facial emotion is subtle and mostly inferred from posture in a rear three-quarter view. The supplied garment and short-hair bases have been adapted for this view, not authored as close-up digital-human assets. Feet, the far hand and some secondary props are partly or wholly occluded by the production camera. The exterior is a soft sky field, not a modeled neighborhood. Desktop negative space is useful; mobile deliberately sacrifices it.

## Future Landing Integration

Use the lossless master for future encoding and the WebP as a review/integration preview. Select the intentional aspect-ratio asset at breakpoints. Keep marketing copy and essential meaning in accessible HTML. Suggested alt text: “A candidate reviews interview notes while practicing a technical interview with an AI interviewer at a home-office desk.” No frontend source, R3F implementation, production GLB, video or animation was created.

## Final Quality Verdict

**Completed: a finished Cycles still and verified intentional derivatives, substantially beyond the PASS 1 blockout.** The actual 2560 × 1440 master, both crops, and separate head/hand detail crops were inspected. No obvious visible proxy geometry, detached hands, broken rig, missing texture or screen occlusion remains. This is a restrained stylized production still for landing-scale use; “AAA-inspired” describes the art-direction ambition, not an independently certified quality rating.

| Final review question | Assessment of actual output |
| --- | --- |
| Beyond PASS 1? | Yes: real rigged human, textured clothing/skin/hair, material separation, custom detail, actual screen artwork and Cycles lighting. |
| Believable human? | Yes at the intended medium shot; no close-up facial realism claim. |
| Personal/lived-in room? | Preparation objects, plant, print, coffee and small domestic storage establish it. The room remains deliberately tidy. |
| Interview preparation? | Notebook diagram, paper, reference book and interview UI work together; confidence is quiet posture rather than an overt expression. |
| Screen obvious but controlled? | Yes; enlarged monitor and dark bezel preserve hierarchy, with no clipped emissive rectangle. |
| Cinematic lighting? | Soft window-led direction, indirect contact, warm practical pool and modest depth separation. It remains bright and restrained rather than dramatically contrasty. |
| Materials avoid universal plastic? | Yes: wood, garment, skin, metal, paper, glass and upholstery use distinct responses; surface detail is restrained. |
| Skin/hair/clothing at landing scale? | Acceptable in the inspected medium shot. The hair/wardrobe retain core-asset foundations and are not close-up hero assets. |
| Intersections/rigging errors? | No obvious visible defects found in final master and detail crops. Hidden surfaces were not certified for alternate camera use. |
| Hands? | Visible hand/wrist sits naturally over input; fingers read at delivery size. Far hand is occluded. |
| Chair/body contact? | Seated pelvis/thigh and chair supports read plausibly; no floating visible body. |
| Scale? | Adult proportions and ordinary desk/monitor/chair dimensions remain believable. |
| Useful left space? | Yes on desktop, including a quiet wall field beside visible window architecture. |
| 3:2 / 4:5 survival? | Visually verified. Portrait uses x=870 through 2021 of the master and retains the full monitor, head, keyboard hand and notebook. |
| Premium landing use? | Suitable as the completed still artwork, subject to replacing temporary product content when the real UI is approved. |

### Deliverable audit

- `scenes/landing-hero-final.blend`: saved with Camera A active, Cycles, 128 samples, 2560 × 1440 render settings; 13 packed images, no unpacked external image dependencies and no visible candidate proxy.
- `renders/final/landing-hero-final-16x9.png`: 2560 × 1440, 16-bit RGB PNG, approximately 18.06 MB.
- `renders/final/landing-hero-final-3x2.png`: 2160 × 1440, 16-bit PNG, approximately 11.08 MB.
- `renders/final/landing-hero-final-4x5.png`: 1152 × 1440, 16-bit PNG, approximately 6.31 MB.
- `renders/final/landing-hero-final-preview.webp`: 1600 × 900, approximately 46 KB; review/integration preview, not the lossless master.
- `scripts/export_final_still.sh`: reproducible crop/preview export, shell syntax checked and executed successfully.

The crop PNGs are exact pixel crops of the checked master, not independently inconsistent re-renders. No alternate was produced because it did not improve the selected result.

All project writes remained inside `landing-hero`. The only authorized outside-project changes were MPFB's user-level extension installation and its own configuration/cache directories. The pre-existing frontend `bun.lock`, `components.json` and `next-env.d.ts` changes were preserved; the tracked diff path list is unchanged from the start. No application/backend/root source, GLB, R3F, video or animation changes were made. Frontend build/test commands were not run for this artwork-only task because their generated caches/files would cross the explicit project write boundary.
