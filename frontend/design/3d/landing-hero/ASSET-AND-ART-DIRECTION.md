# Landing Hero — Asset & Art Direction

**Research status:** planning only; no assets were downloaded, no accounts were used, and no Blender or frontend source was changed. Source facts below were checked on 2026-09-08. “Verified” means stated by the linked provider; all other design choices are recommendations/inferences.

## 1. Executive Recommendation

Use strategy **C: a hybrid MPFB2 base character, custom Blender art direction, and optional Mixamo body-motion reference/animation**.
Build the hero room, desk silhouette, screen, and all storytelling composition ourselves in Blender.
Use [Poly Haven](https://polyhaven.com/) as the primary CC0 source for HDRIs and PBR surfaces; use its models only where their style fits.
Use [ambientCG](https://ambientcg.com/) only to fill a specific CC0 material gap, not as a second indiscriminate library.
Treat [Blendkit/BlenderKit](https://www.blendkit.com/) as a later, fast iteration source after an asset-level license and quality review.
Treat [Sketchfab](https://sketchfab.com/) as exceptional-case sourcing only; review the selected model’s license, author, provenance, and distribution implications individually.
Frame a three-quarter, near-over-the-shoulder view: candidate on the right, interview screen readable near centre, calm daylight negative space on the left.
Put the AI interviewer inside the product screen, not as a second person in the physical room.
Deliver a carefully art-directed **still image first** (desktop AVIF/WebP with intentional crops); defer loop/video and R3F until a still proves the composition.

## 2. Visual Narrative

The scene is a personal, intentional home workspace in late-morning daylight. An adult software engineer sits at a real desk rather than a generic “future office” set. They have just moved from handwritten preparation into a live practice session: the notebook is open to a small system-design sketch, a clean résumé sheet is partly visible, and the monitor shows an AI interviewer in the web application.

The candidate is attentive rather than stressed—upright, one hand near the keyboard or desk, shoulders released, gaze meeting the screen. This is the short emotional beat after preparation has paid off: they are participating capably and can imagine succeeding. The room’s warmth and personal specificity make the landing visitor think “this could be my next interview,” not “this is a distant fictional avatar.”

**ART-DIRECTION RECOMMENDATION / INFERENCE:** Story progression is carried by the relationship between preparation objects, posture, and screen rather than a literal multi-scene timeline. That keeps one still readable at landing-page scale.

## 3. Character Strategy

### Mixamo

**VERIFIED SOURCE INFORMATION.** Adobe’s [Mixamo FAQ](https://helpx.adobe.com/creative-cloud/faq/mixamo-faq.html) says the service is free with an Adobe ID, its auto-rigger and animation library are for bipedal humanoids, and its custom-character requirements include a clean, centred, neutral humanoid mesh. Adobe’s [rigging/export workflow](https://helpx.adobe.com/creative-cloud/help/mixamo-rigging-animation.html) documents uploading FBX, OBJ, or ZIP assets for auto-rigging, placing body markers, and downloading rigged characters/animations. Adobe’s [licensing FAQ on its community](https://community.adobe.com/questions-696/mixamo-faq-licensing-royalties-ownership-eula-and-tos-589400?lang=en) states that the then-current preview permitted commercial and non-commercial creative use but prohibited distributing raw character/animation files; it also states that Mixamo content must not be used to train machine-learning models.

**ART-DIRECTION RECOMMENDATION / INFERENCE.** Mixamo is useful for a rigged base and, more importantly, as a source of bipedal seated/typing/listening motion to retarget or use as pose reference. Do not expect a stock Mixamo character, hair, clothing, facial expression, or skin lookdev to meet this hero’s quality bar unchanged. Its generic anatomy, game-ready topology, limited face system, and common recognisability make it a poor final character with minimal modification. Exact seated clips must be manually confirmed in the current library before acquisition; no account was used for this research.

### MakeHuman / MPFB

**VERIFIED SOURCE INFORMATION.** [MPFB](https://static.makehumancommunity.org/mpfb.html) is a free, open-source human generator for Blender. Its [getting-started documentation](https://static.makehumancommunity.org/mpfb/docs/getting_started.html) says MPFB2 requires Blender 4.2+ (therefore the installed Blender 5.2 LTS exceeds the documented minimum), can create and adjust a human from scratch in Blender, load skin/eye/hair/clothing assets, add a standard rig, and add IK helpers. The [MakeHuman–MPFB comparison](https://static.makehumancommunity.org/mpfb/faq/differences_between_mpfb2_and_makehuman.html) identifies MakeHuman as the standalone application and MPFB2 as the Blender add-on; they share core system assets, while MPFB exports through Blender. The [asset download guide](https://static.makehumancommunity.org/assets/downloadassets.html) recommends the MakeHuman system asset pack for MPFB and identifies other assets as user-contributed repositories. The [license page](https://static.makehumancommunity.org/about/license.html) says core graphics assets are CC0, while the MPFB code is GPL and MakeHuman code is AGPL. The [MPFB 2.0.16 release notes](https://static.makehumancommunity.org/mpfb/releases/release_2016.html) document ARKit-compatible facial shape assets and a modern Rigify workflow.

**ART-DIRECTION RECOMMENDATION / INFERENCE.** MPFB is the better character base for this task because authoring, proportional edits, shading, Rigify, pose work, and eventual export remain inside Blender. It is not a one-click AAA person: the base face, supplied hair/clothes, eyelashes, teeth, and skin must be evaluated as starting material, then refined or replaced selectively. Use it to create an original, non-celebrity character, not an aggregate of arbitrary library presets.

### Other serious candidate if any

No third character platform is adopted for this project. A paid, proprietary digital-human ecosystem could raise the off-the-shelf face ceiling, but it would add licensing, account, conversion, and web-optimization uncertainty while weakening the Blender-first workflow. That does not beat a deliberately art-directed MPFB hybrid for a single, medium-distance landing still.

| Option | Time to usable base | Quality ceiling for this hero | Rig / pose / animation | Face, hair, clothing | Blender workflow | Licensing confidence | Verdict |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Full human from scratch | Slowest; large sculpt, retopo, UV, groom, rig, and lookdev effort | Highest in principle, but high delivery risk | Must build or obtain a rig; expensive to iterate | Can be excellent, but every department is bespoke | Native after substantial setup | Original work, with any source textures separately cleared | Do not choose for the initial hero. |
| Mixamo character with minimal modification | Fastest first pose | Low-to-medium; recognisably stock without extensive rework | Strong biped animation/auto-rig utility | Weakest differentiation and face control | Import/cleanup/retarget required | Verify Adobe terms at acquisition; no raw-file redistribution | Do not choose as final strategy. |
| **Hybrid: MPFB2 base + custom Blender art direction + Mixamo motion as needed** | Moderate, with early visual control | High enough for a medium-distance cinematic still | MPFB standard/Rigify + optional Mixamo body motion | Custom head shaping, shader, groom and wardrobe can carry identity | Strongest: character remains editable in Blender | MPFB core graphics CC0; inspect every non-core asset; Mixamo terms re-check required | **Primary choice.** |

### Decision

**Primary strategy:** **Hybrid MPFB2 base character + custom Blender art direction.** Create one original adult candidate in MPFB2, choose/construct a Blender-friendly rig, then customise silhouette, head proportions, simple haircut/groom, cloth, material response, hands, and seated pose. Mixamo is permitted only as an optional body-animation source or pose reference after terms are checked again at download time.

**Fallback:** **Mixamo character/rig + substantial Blender replacement lookdev.** Use only if MPFB’s current Blender-5.2 workflow or its base mesh prevents timely posing. Replace or heavily modify the visible hair, clothes, face proportions, materials, and accessories; do not present a stock character unchanged.

## 4. Environment Strategy

Model the hero-specific architecture ourselves: room envelope, window wall, built-in shelf or wall plane, desk, monitor/laptop assembly, screen bezel, keyboard, notebook, résumé pages, cable routing, and the hero chair silhouette. These elements control scale, composition, the candidate’s ergonomics, and the visual language; sourcing them as unrelated assets risks a kitbashed room.

Source only genuinely generic, secondary ingredients after review: one chair or lamp as a topology starting point, a plant, a few books, coffee vessel, paper/fabric/wood/metal surfaces, and neutral HDRI lookdev support. Adapt dimensions, shaders, wear, and colour to the scene. Avoid importing a whole prebuilt office, scene, branded computer, or a clutter pack.

**VERIFIED SOURCE INFORMATION.** [Poly Haven’s models catalogue](https://polyhaven.com/models) includes Furniture, Lighting, Office & Stationery, Electronics & Appliances, and other relevant categories; an example [metal office desk](https://polyhaven.com/a/metal_office_desk) provides PBR map types, 7K triangles, scale information, and a CC0 designation. Its [Blender add-on documentation](https://docs.polyhaven.com/en/guides/blender-addon) says the add-on can browse HDRIs, textures, and models in Blender, defaults to 1K imports, and that assets are generally available at at least 8K or higher. This is useful later, not an authorization to download now.

## 5. Asset Source Matrix

| Source | Asset Type | Why Useful | License | Risk | Recommended Use |
| --- | --- | --- | --- | --- | --- |
| [Adobe Mixamo](https://www.mixamo.com/) | Biped characters, automatic rigging, body-animation library | Quick seated/listening/typing body-motion source; can auto-rig an eligible custom humanoid. [Adobe documentation](https://helpx.adobe.com/creative-cloud/help/mixamo-rigging-animation.html) verifies uploads and rig workflow. | Adobe’s [published FAQ](https://community.adobe.com/questions-696/mixamo-faq-licensing-royalties-ownership-eula-and-tos-589400?lang=en) describes royalty-free creative use during preview and disallows raw-file distribution; **re-verify terms at acquisition.** | Stock look, generic face, current preview/terms status, raw asset redistribution restriction. | Motion/pose aid and fallback base only; never a minimally changed hero. |
| [MakeHuman / MPFB](https://www.makehumancommunity.org/) | Parametric base human, core graphics, Blender add-on, rig/face tools | Blender-native editing and proportional control; MPFB docs confirm rigs, IK helpers, hair/clothes asset slots, and shape adjustments. | [Core graphics assets are CC0](https://static.makehumancommunity.org/about/license.html); code has GPL/AGPL licenses. User-contributed asset packs require per-asset verification. | Base realism and contributed asset quality vary; non-core asset provenance must be checked. | **Primary character base**; use only verified core assets or separately cleared clothing/hair. |
| [Poly Haven](https://polyhaven.com/) | HDRIs, PBR textures, selected models | Curated source for believable surface response and lighting support; [models](https://polyhaven.com/models) and [textures](https://polyhaven.com/textures) support props/lookdev. | [All provided HDRIs, textures, and models are CC0](https://polyhaven.com/license). | Catalogue may not supply a coherent complete office; high-resolution maps may be unnecessarily heavy. | **Primary environment/material source**; select only assets that fit the scene, then set resolution per final use. |
| [ambientCG](https://ambientcg.com/) | PBR surfaces, HDRIs, 3D assets | Verified catalogue of PBR materials, HDRIs, and models; a useful specific-gap supplement to Poly Haven. | [All downloadable assets are CC0](https://docs.ambientcg.com/license/). | Mixing scans/procedural scale conventions can create inconsistent materials; avoid duplicate library browsing. | Secondary material source only when Poly Haven lacks a needed surface. |
| [Blendkit / BlenderKit](https://www.blendkit.com/) | Blender-integrated models, materials, HDRIs, scenes, add-ons | Its catalogue is directly searchable in Blender and covers office furniture, computers, lights, decor, materials, and HDRIs; this can accelerate local iteration. | [Blendkit licenses](https://www.blendkit.com/docs/licenses/) are CC0 or Royalty Free; both allow commercial higher-level derivatives, while Royalty Free forbids re-sale of the 3D asset itself. | Requires account/add-on access and an asset-by-asset license/quality/provenance check; library styling varies; subscription availability affects iteration. | Optional, later-stage secondary props only. Prefer CC0 when practical; record the exact asset and license. |
| [Sketchfab](https://sketchfab.com/) | Downloadable models: furniture, rooms, props, characters; glTF/original file formats | Very broad selection; [Sketchfab confirms glTF is available for downloadable models](https://sketchfab.com/features/gltf). | Licensing is selected per item: official [license guidance](https://sketchfab.com/blogs/community/refine-downloadable-model-searches-with-new-license-filters/) distinguishes CC0, CC BY, NC, ND, SA, Standard, and Editorial. **LICENSE REQUIRES MANUAL REVIEW for every candidate.** | Source provenance, attribution, NC/ND/SA restrictions, editorial-only content, heavy or poorly constructed meshes, and style mismatch. | Last-resort gap filler only; accept CC0 or a fully documented commercial-compatible license after review. Never use Editorial/NC/ND assets. |

## 6. Asset Licensing Rules

1. Before any future download, create an asset record with provider, asset title/ID, direct URL, author, license text/link, acquisition date, intended scene use, and a local copy of any required attribution.
2. “Free” never means “cleared.” Accept only a license that explicitly permits the intended commercial/promotional landing use and necessary derivative work.
3. **Poly Haven and ambientCG:** their sites state CC0 for downloadable assets; retain source records even though attribution is not required.
4. **MPFB/MakeHuman:** treat only core graphics as verified CC0. Treat every user-contributed hair, garment, body part, or pack as **LICENSE REQUIRES MANUAL REVIEW** unless its item record clearly states compatible rights.
5. **Mixamo:** use under the current Adobe terms shown at acquisition; preserve the terms URL/date. Do not redistribute its raw character or animation files in a repository, asset pack, public download, or web-delivered GLB that exposes them as an extractable source asset without confirmed permission.
6. **Blendkit:** record whether the exact asset is CC0 or Royalty Free. Royalty Free is suitable for a rendered still/video and potentially an inseparably packaged runtime asset, but do not re-sell or separately distribute the model.
7. **Sketchfab:** manually inspect the model page, author, license, and download terms. Reject Editorial, CC BY-NC, CC BY-ND, and any asset with unclear provenance. CC BY requires correct attribution; CC BY-SA may impose downstream sharing obligations and should be reviewed by the project owner before use.
8. Do not use recognizable brand marks, copied UI, real résumés, real company names, celebrity likenesses, or unauthorised scans. Create generic product UI and fictional preparation materials.
9. Do not commit raw third-party source files to the frontend repository unless their license and web-delivery implications are expressly approved. Keep source assets in the approved design asset workflow with their ledger.

## 7. Room Art Direction

### Architecture

A compact, well-composed home office: warm off-white plaster walls, a broad left-side window with translucent linen blind or curtain, medium-oak floor, a built-in shelf or low credenza in the background, and enough depth behind the candidate for soft separation. It should feel attainable and personal, not luxury real estate or an empty showroom.

### Palette

Daylight neutrals: chalk, warm grey, tobacco/oak, charcoal, muted navy or forest green clothing, and a single restrained rust/terracotta accent in a book spine, ceramic, or soft furnishing. The screen provides a quiet cool-neutral counterpoint. No neon, saturated cyan-magenta scheme, RGB strips, or black sci-fi void.

### Materials

Use timber with directional grain and small hue/roughness variation; chalky plaster with low-frequency tonal variation; woven chair/curtain textiles; anodized or powder-coated metal; matte paper; restrained ceramic; and a physically plausible screen/glass layer. Slight edge wear belongs only where use justifies it.

### Props and storytelling

Keep roughly five meaningful clusters, each with a job:

- open interview notebook with a simplified boxes/arrows system-design sketch;
- one unstated, generic résumé/CV sheet, not a readable real person’s data;
- a technical paperback or API/reference book with no copied brand cover;
- coffee or water, a pen, and a small stack of preparation cards;
- headphones set aside, implying focused prep rather than a gamer setup;
- one personal cue such as a modest plant, camera, or framed abstract print—no recognisable logo.

**ART-DIRECTION RECOMMENDATION / INFERENCE:** Prop density should be selective: about 70% clean surfaces and 30% purposeful lived-in detail. Objects should support the candidate’s decision-making, not compete with the face, hands, or product screen.

## 8. Character Art Direction

- **Age range:** 26–35. Present an adult early/mid-career software candidate without suggesting a particular real person or demographic stereotype.
- **Professional styling:** capable and contemporary: a textured knit or overshirt over a quiet T-shirt, or an unbranded shirt with soft tailoring. Avoid suit-and-tie stiffness, hoodie cliché, gaming headset, or tech-bro costume.
- **Silhouette:** recognisable at laptop scale through shoulder line, hair outline, and a strong seated profile; avoid tiny head/oversized hands or exaggerated idealized proportions.
- **Clothing:** two or three real textile layers, with seams, hems, weave, compression at elbows/waist, and restrained folds driven by the seated pose.
- **Pose:** neutral asymmetric seated stance—slight forward attention, stable pelvis/feet, released shoulders, one hand resting and the other near input—rather than a frozen typing pose or theatrical gesture.
- **Expression:** calm concentration with a faint easing around brow/jaw; if the face is visible, a small, genuine confidence cue rather than a broad advertising smile.
- **Realism level:** high-end stylised realism. Facial landmarks, skin response, hands, and clothing obey reality, while camera distance, soft light, and restrained imperfections avoid photoreal close-up scrutiny and uncanny valley.

## 9. Camera Strategy

| Candidate shot | Strength | Limitation | Use |
| --- | --- | --- | --- |
| Over-the-shoulder, candidate toward screen | Most immediate product legibility; viewer inhabits the candidate role; hides face-risk | Can reduce environment and character emotional read; screen may dominate | Strong alternate crop or product-detail shot. |
| Cinematic side/three-quarter | Shows seated human, desk ritual, screen, and room depth in one frame; supports a readable profile and premium lighting | Needs disciplined blocking so no silhouette overlaps the screen | **Main landing hero.** |
| Wider environmental storytelling shot | Best sense of home, daylight, and aspirations; accommodates typography | Candidate and product can become too small on laptop/mobile crops | Secondary editorial still or section image. |

### Main hero composition

Use a **three-quarter, near-over-the-shoulder camera from the candidate’s left-rear side**, at approximately seated eye/shoulder height. The camera sees the candidate’s profile/cheek enough to read focus, their right-side silhouette on the right third, the monitor in a three-quarter view near the centre, and the sunlit room receding left. Start with a 45–55 mm full-frame-equivalent lens: natural enough for a personal room, compressed enough to feel cinematic, with no wide-angle “CG dollhouse” distortion.

Preserve a low-detail, light-value safe area on the left third for optional landing typography. Keep the candidate’s head, visible hand, screen face, notebook, and key chair contour inside the centre 60% so they survive a 4:5 crop. Test desktop 16:9, 3:2 laptop, and 4:5 mobile crops from the blockout; produce deliberately re-framed crops where one master does not work.

## 10. Lighting Strategy

Bright cinematic does not mean flat. Make daylight through the left window the soft, broad key at roughly 5200–6000K; let the exterior be several stops brighter but controlled by sheer fabric or an off-camera diffusion plane. Add a very soft interior fill so the candidate’s face and shirt retain shadow shape. Add a warm, low-intensity practical lamp at roughly 2700–3200K behind or beside the candidate; it should tint the midground and add emotional warmth, not act as a visible orange spotlight.

Use soft bounced indirect light from walls/floor, a subtle edge/rim from window direction to separate hair and shoulder from the background, and a low-contrast area light only if screen-side facial values need recovery. The monitor should give a faint cool lift to the closest cheek/hand, never become the primary blue light source. Preserve contact shadows under objects, stable soft penumbrae, specular highlights on wood/ceramic/eyes, and a gentle colour-temperature dialogue of cool daylight versus warm interior practical.

**ART-DIRECTION RECOMMENDATION / INFERENCE:** Grade around warm highlights, neutral skin, soft green/blue shadow information, and protected whites. Avoid crushed blacks, bloom-heavy neon, teal/orange cliché, and uniform ambient illumination.

## 11. Material / Lookdev Strategy

Premium output comes more from calibrated light response and hierarchy than polygon count. Apply material scale in real units and avoid one universal roughness value.

- **Skin:** plausible melanin/albedo breakup at low amplitude, per-region roughness, soft subsurface scattering, normal detail only at macro/micro scale appropriate to camera distance, wetline/lip subtlety, and an eye model with cornea/specular separation. Never use pore detail as noise wallpaper.
- **Hair:** clean silhouette first; use clumped directional strands/cards/curves appropriate to delivery mode, varied roughness, subtle flyaways, and controlled shadow density. A single plastic helmet shape is unacceptable.
- **Cloth:** weave normal/bump at correct scale, roughness variation, seam topology, fold compression at contact points, and colour variation constrained to the fabric, not random dirt.
- **Wood/plaster/metal:** map orientation follows fabrication; wood has grain/edge treatment, plaster has low-frequency variation, powder coat has soft broad highlights, and metal receives only physically credible wear.
- **Paper/ceramic/glass/monitor:** use real thinness/bevels, contact shadows, off-white paper, ceramic micro-roughness, correct IOR/reflection behaviour, and monitor screen/glass separation. Make screen content legible through contrast—not emissive glare.
- **Geometry:** bevel hard edges, smooth deliberately, prevent perfect intersections, add only visible manufacturing joins. Resolve shade normals before adding expensive texture detail.

The anti-plastic checklist: no perfectly uniform roughness, no razor-sharp asset edges, no repeated texture tiles in the hero field, no pure black plastics, no default white walls, no full-strength normal maps, no unmotivated fingerprints/scratches, and no all-surfaces-equal reflection.

## 12. Emotional Storytelling

| Emotional beat | Visible evidence | What it tells the visitor |
| --- | --- | --- |
| Preparation | Note pages, system sketch, tidy reference book, headphones intentionally set aside | The candidate took this opportunity seriously. |
| Anticipation | Open posture with a small forward lean, coffee still warm, focused gaze | This moment matters but is manageable. |
| Engagement | Candidate and clearly framed interviewer UI share a sightline; hands are ready rather than frantic | The product is the place where practice becomes conversation. |
| Confidence / relief | Relaxed shoulders, quiet facial softening, warm practical in an ordered room | The candidate leaves the encounter more prepared and optimistic. |

Avoid literal status badges, confetti, exaggerated smile, floating code, holograms, or “AI magic” effects. Those make the emotion generic and date the image faster than human behaviour does.

## 13. Screen / Product Integration

Place the AI interviewer **inside the monitor/laptop UI**. This has the clearest causal story: the physical candidate is preparing in their own room; the product is the interface that hosts the interviewer. A second physical 3D person would turn the story into a staged meeting, confuse whether the platform is remote or in-room, and use precious composition area without explaining the web application.

The screen should show a restrained future-compatible interview-room composition: interviewer avatar/video panel, interview context, response/controls region, and no invented backend state, fake transcript, fake employer, or fake metrics. Use generic placeholder labels in early Blender passes. Before final delivery, replace them only with approved product UI from Figma/the actual frontend design. The app is currently a placeholder and its interview runtime is not connected, so this research does not treat any particular UI as a shipping contract.

For a rendered still, keep the screen’s luminance below a clipped white rectangle and use a screen-safe UI plane/material that can be swapped in compositing. The physical screen must remain an attractive object even when responsive implementation uses a separate accessible text/copy treatment outside the image.

## 14. Landing Integration

Design the master as 16:9 cinematic media for a bright editorial Next.js page, normally below or beside landing copy rather than a full-screen takeover. The asset must work as an art-directed media block:

- export a clean 16:9 desktop master, then evaluate intended 3:2 and 4:5 crops rather than relying on `object-fit: cover` alone;
- protect candidate head/profile, monitor UI, hand, notebook, and desk edge from crop boundaries;
- retain a calm left-side negative-space zone for a possible overlay, but do not make text-overlay dependence mandatory;
- use foreground/midground/background depth and high-contrast silhouettes that stay readable at laptop width;
- avoid details that need zooming to understand the product, and avoid embedded copy that would be unreadable, stale, or inaccessible;
- offer meaningful adjacent HTML text and image alt text; the image is supportive hero media, not the only explanation of the platform;
- respect reduced-motion if a later video is approved, and provide a still fallback.

**Future implementation constraint verified from the repository:** Next.js 16, React 19, Three.js/R3F/Drei are present. The current `AvatarStage` is a client-only R3F canvas intended to mount only when real assets exist, and the landing route is currently a placeholder. That supports a light initial image delivery and avoids binding this artwork prematurely to the interview runtime.

## 15. Delivery Strategy

| Mode | Visual ceiling | Weight / GPU / mobile | Responsive behaviour | Implementation and iteration cost | Decision |
| --- | --- | --- | --- | --- | --- |
| High-quality still render | Highest reliable quality per byte; full Cycles lookdev, sampling, denoise, and grade are practical | Optimised AVIF/WebP is small; negligible runtime GPU; strongest mobile safety | Requires intentional desktop/mobile crops but is predictable | Lowest page complexity; easy to replace after art review | **Recommended first release.** |
| Short optimised cinematic loop/video | Can add breathing, subtle gaze, curtain/light motion; quality can remain high if pre-rendered | Heavier media decode/bandwidth; mobile battery/data and LCP/autoplay policy require care | Needs poster, crops, no-autoplay/reduced-motion logic | Animation, render, encoding, and QA cost rises substantially | Defer until still composition proves value; use only as optional desktop enhancement. |
| Realtime Three.js / React Three Fiber | Interactive angle/parallax possible, but browser quality ceiling and asset budget constrain hair/skin/lighting | Highest GPU/memory risk; asset decoding and texture budget threaten Core Web Vitals/mobile | Responsive camera can adapt, but every device/performance tier needs testing | Highest engineering/art optimization cost; would compete with the real interviewer runtime | Do not use for the landing hero now. |

### Recommended path

Render a high-quality, graded still first, then deliver responsive static derivatives. Keep the Blender scene and screen layer reusable so a 4–6 second pre-rendered desktop loop can be explored later if it demonstrably improves the landing page. Do not introduce a landing R3F scene merely because R3F exists in the product: its runtime performance burden is not justified by this non-interactive visual.

## 16. Blender Production Passes

Each pass uses the established loop: **inspect → smallest isolated change → viewport/camera screenshot → art-direction critique → iterate**. Do not generate the entire scene with one monolithic Python script.

| Pass | Scope and proof of completion |
| --- | --- |
| **PASS 0 Reference** | Compile a rights-safe, visual-only reference board for daylight home-office architecture, seated posture, material response, wardrobe, framing, and editorial crops. Write a short shot brief, scale targets, palette, and negative-space map. No production asset download. |
| **PASS 1 Environment blockout** | Create reproducible room shell, window, desk, chair volume, monitor volume, shelf/credenza massing, prop placeholders, three cameras, and crop guides. Use simple geometry only. Validate silhouette/readability in screenshots. Structural collections, transforms, cameras, and light placeholders should be reproducible through small Blender Python scripts saved later under `scripts/`. |
| **PASS 2 Candidate** | After license gate and acquisition, create the MPFB candidate, apply/select a pose rig, establish ergonomic chair/desk contact, and test silhouette/head/hand visibility in the approved camera. Do not detail the face before the body reads at landing size. |
| **PASS 3 Product screen** | Create the screen surface, bezel, viewing angle, and generic UI placeholder. Verify screen luminance/reflections and its swap/compositing path. Use approved future Figma/product content only when supplied. |
| **PASS 4 Lookdev** | Develop hero materials in camera order: skin/eyes/hair, clothing, desk/wood, walls/floor, chair textile, and hand-held/story props. Check real-world texture scale, roughness breakup, and asset style cohesion. |
| **PASS 5 Lighting** | Establish daylight key, bounce/fill, warm practical, minimal screen lift, exposure, filmic grade, and contact-shadow integrity. Compare clay, albedo, roughness, normal, and final-light screenshots; correct causes rather than grading over defects. |
| **PASS 6 Final composition** | Lock the main 16:9 frame and independently solve 3:2/4:5 crops. Refine pose, prop clusters, depth of field, facial read, screen clarity, and final render settings. Conduct art-direction review at actual landing dimensions. |
| **PASS 7 Landing delivery optimization** | Produce the approved still master and responsive encoded derivatives; make an optional video proof only if approved. Measure dimensions, byte size, decode behaviour, alt/copy integration, and page performance before any frontend implementation. Export GLB only for a separately approved real-time experiment. |

## 17. First Asset Shopping List

Do not download these yet. PASS 1 deliberately requires **no external asset**: it should prove the shot with self-made primitives. The single planned acquisition below is the smallest requirement for PASS 2 after the asset ledger and license checks—not a general shopping list.

| Need | Search terms / selection criteria | Preferred source | Why it is required now |
| --- | --- | --- | --- |
| Editable candidate base and rig | `MPFB2`, `MakeHuman system assets`, `standard rig`, `Rigify rig`, `faceunits01` only if facial shape work is in scope | [MPFB downloads / docs](https://static.makehumancommunity.org/mpfb/docs/getting_started.html) | PASS 2 needs an editable original human and controllable seated pose. Verify pack licensing; core system assets are the intended initial set. |

Explicitly defer until their relevant pass: daylight HDRI (PASS 5), wood/plaster/textile PBR materials (PASS 4), a Mixamo seated-motion aid (only if PASS 2 pose work needs it), finished office/room models, hair/clothing packs, branded hardware, random decor, and all Sketchfab/Blendkit models. PASS 1 can prove the shot with self-made primitives; those assets would add choice noise before the composition is known.

## 18. Risks

| Risk | Prevention / decision gate |
| --- | --- |
| Uncanny valley | Keep camera at medium distance; prioritise silhouette, pose, eyes, skin response, and calm expression. Reject a face that needs close-up realism to work. |
| Licensing | Use the ledger and rules above. Manual-review all user-contributed MPFB, Mixamo at acquisition, Blendkit asset level, and every Sketchfab item. |
| Polycount / texture bloat | Block out first; apply camera-distance budgets; use instancing and low/mid detail beyond focus; choose only needed texture resolution. Do not assume a Cycles source scene can become a web GLB. |
| Web performance | Still first, responsive encodes/crops, no runtime canvas or unoptimised source assets. Later video/R3F requires separate performance budget and device testing. |
| Asset mismatch | Give sourced objects one scale, palette, bevel, material, and wear language; prefer modelling hero objects rather than assembling a showroom. |
| Character quality | Make a PASS 2 camera test early. If MPFB cannot achieve a convincing medium-shot candidate after proportion/shader/groom tests, trigger the Mixamo fallback rather than overinvesting. |
| Render time | Establish samples/denoise/lighting in low-resolution camera tests; reserve high-resolution finals for a locked composition. Avoid hair/volumes/detail that do not survive landing scale. |
| Style inconsistency | Use a written palette/material brief and review every new asset in the hero camera under the actual light. No cyberpunk, neon, plastic furniture, stock corporate set dressing, or mobile-game proportions. |
| Screen/UI obsolescence | Keep screen content as a swappable plane/layer until approved Figma or real product UI exists; never let imagined product details become a claimed implementation. |

## 19. Go / No-Go Decision

**GO — enough information exists to begin Blender PASS 1.** The visual narrative, main camera, environment scope, material/lighting direction, character primary/fallback strategy, asset-source hierarchy, and licensing gates are defined. PASS 1 needs no downloaded asset and should use a self-made primitive blockout only.

**Gate before PASS 2:** create the asset ledger, re-check the current source terms for any selected non-core asset, and choose MPFB’s verified core system assets (or invoke the documented Mixamo fallback). Do not proceed to production downloads or character import before that gate is met.
