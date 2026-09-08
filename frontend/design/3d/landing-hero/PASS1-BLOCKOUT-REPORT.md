# PASS 1 — Environment Blockout

## Scene Summary

Primitive-only personal workspace, built incrementally through the connected Blender MCP in Blender 5.2.0 LTS / EEVEE. The dedicated default scene was inspected and saved before editing. Architectural, workspace, proxy and screen stages were inspected through viewport screenshots; lighting and camera decisions were inspected in rendered images.

The room is approximately 4.2 × 4.0 × 2.8 m, with a broad left window opening, sill, mullion and folded blind mass. A rear credenza separates the desk from the background. The desk is 1.60 × 0.78 m with its top at approximately 0.78 m. A low-backed chair holds an asymmetric seated proxy on the right. The monitor sits toward the desk's rear, unobstructed from the left-rear camera. An open diagram notebook, abstract résumé sheet, reference book, pen, vessel and headphones establish preparation. A modest lamp and book mass supply background domestic cues.

Eight named collections separate guides, architecture, furniture, props, candidate, screen, lights and cameras. All geometry is self-made primitives. Materials are plain value/color separators, with no textures or final shaders. No second physical interviewer exists: the interviewer symbol is contained within the replaceable screen collection.

## Camera A

**Strengths:** 50 mm, camera height 1.65 m. Best balance of monitor, attentive candidate, visible input hand and preparation objects. Candidate head is around the right third; screen is just left of center. Left wall provides a quiet light field. The lamp was moved after a render exposed a monitor overlap. Notebook diagram and résumé rules improved the preparation cue.

**Weaknesses:** Window architecture is outside this frame; daylight is implied through illumination. Primitive anatomy still looks like a mannequin, and the room needs later surface and personal-detail development. The screen alone communicates a remote interview/conversation, not specifically an AI technical interview; approved UI and adjacent landing copy must complete that meaning. Shadows and warm practical contribution are restrained to the point of being weak.

**Responsive crop assessment:** Visually verified the 1600 × 900 master, centered 1350 × 900 / 3:2 crop, and offset 720 × 900 / 4:5 crop. Portrait uses master x=32–77%, rather than a centered crop. It preserves the complete screen, candidate head, keyboard hand and notebook; it trims the résumé, outer shoulder/chair and background. Mobile has no useful left-side copy area. Head/hair expansion in PASS 2 must be checked against the portrait right margin. Crop metadata is saved in `00_GUIDES`.

## Camera B

**Strengths:** 50 mm, height 1.60 m, closer and more directly behind the candidate. Largest screen and most immediate screen/hand relationship. The notebook remains clear on desktop. Useful product-detail alternate.

**Weaknesses:** Foreground torso and chair crowd the right edge; lamp overlaps the head silhouette in projection. Less home-office context and less emotional/profile information. The composition concentrates on computer interaction more than the career-preparation setting.

**Responsive crop assessment:** 16:9 reads well; 3:2 trims the outer chair/arm. A 4:5 crop cannot comfortably retain the full notebook, monitor and head together with safe margins. Requires a dedicated portrait camera adjustment, not an object-fit assumption. Not selected for further refinement this pass.

## Camera C

**Strengths:** 45 mm, height 1.85 m. Shows desk legs, seated legs, chair base, floor depth and a sliver of the left window architecture without a wide-angle dollhouse effect. Largest calm left area and clearest full-body ergonomic reference.

**Weaknesses:** Excess blank wall makes the scene more generic and distant. Product and paper cues shrink at landing size. The room starts to compete with the subject; less psychological closeness than A.

**Responsive crop assessment:** 3:2 remains workable. A portrait crop shifted toward the interaction can retain the head, screen and hand, but discards the window and much of the environmental rationale. Only A received additional rendered crop proofs. C is a secondary composition reference, not the mobile master.

### Rendered-image self-critique

| Required question | Camera A | Camera B | Camera C |
| --- | --- | --- | --- |
| 1. Immediately reads as candidate preparing/interviewing? | Preparation plus remote conversation; technical/AI specificity needs approved UI/copy. | Remote conversation strongest; preparation secondary. | Workspace preparation visible, interview less immediate. |
| 2. Candidate connected to monitor? | Yes: gaze axis and keyboard hand. | Yes, strongest connection. | Yes, but smaller. |
| 3. Product screen understood? | Two-panel interview placeholder is clear. | Clearest/largest placeholder. | Panel layout visible; detail small. |
| 4. Personal rather than corporate? | Moderately: notebook, small desk, lamp, storage. | Moderately; domestic context reduced. | Room scale helps, blank walls remain generic. |
| 5. Natural candidate silhouette? | Seated asymmetry works; primitive joints need replacement. | Attention reads; torso dominates. | Full pose reads; joint/contact refinement needed. |
| 6. Useful left negative space? | Yes, especially upper-left desktop. | Yes, but less overall breathing room. | Yes, arguably excessive. |
| 7. Avoids generic Blender-scene appearance? | Composition has purpose; clay/mannequin appearance remains intentionally unfinished. | Product focus helps; proxy dominates. | Weakest: broad empty room remains generic. |
| 8. Believable at landing size? | Yes at composition level. | Yes for a product-detail image. | Plausible but product loses weight. |
| 9. Critical content survives responsive crop? | Yes with verified offset portrait crop; secondary props/chair trimmed. | Not all in portrait; reframe required. | Core interaction can survive an offset crop; environment lost. |
| 10. Worth high-quality character investment? | Yes: proceed first with proportion/pose test. | Keep alternate, not investment target. | Keep ergonomic/environment reference. |

## Recommended Camera

**Camera A.** It keeps preparation, candidate and product in one readable relationship while reserving light negative space for editorial desktop use. Unlike B, it has a demonstrated portrait crop containing the essential story. Unlike C, the room does not overpower the product. Recommendation approves the composition for the next blocking phase, not final image quality.

## Candidate Placement Notes

Keep pelvis near (0.66, -0.43, 0.59) m, head center near (0.60, -0.275, 1.385) m, and left hand near (0.31, 0.065, 0.816) m. Preserve relaxed shoulders, slight forward attention and asymmetric elbows/feet. Maintain monitor clearance and notebook visibility. Fit the rig to chair and desk contacts; do not simply scale a finished character to the proxy bounds. Refine neck, elbow and knee transitions and verify soles, seat compression and wrist contact. The faceless proxy proves a sightline, not eye direction or confidence expression. Test future hair/head width in A's portrait crop before detailed work.

## Screen Placement Notes

The monitor bezel is approximately 0.66 × 0.40 m, centered at (0.19, 0.60, 1.20) m; the replaceable face points toward negative Y. `50_SCREEN` contains the base surface and shallow neutral panel geometry. Hide/remove the placeholder panel geometry when placing approved Figma/product artwork on the screen surface. Preserve the screen's dimensions and camera visibility. Current symbols and bars are layout placeholders, not feature contracts, transcripts, metrics or employer claims. Screen luminance is controlled by simple non-emissive materials; final integration needs reflection/luminance testing and approved content at actual delivery size.

## Lighting Direction Notes

A broad 2.4 m area source inside the left opening establishes cool-neutral daylight, with a large gentle front/interior fill and a small warm practical pool. The resulting image is bright and readable without neon or a glowing monitor. Improve directional contrast, soft contact shadows, candidate/background separation and the practical's local warmth during lookdev; current lighting is somewhat flat. Retain controlled whites. No HDRI, external textures, bloom, animation or final material work was used.

## PASS 2 Gate

**Composition GO for an MPFB proportion/pose test in a separately authorized PASS 2.** Use A as the investment target and re-check its desktop and offset portrait crops immediately after replacing the proxy. This does not approve final anatomy, likeness, wardrobe, lighting or production delivery. Follow the approved asset ledger/license gate before installation/import. MPFB has not been installed or imported, and PASS 2 has not begun.

## Reproducibility and Verification

`scripts/pass1_environment_blockout.py` was consolidated after visual review. It recreates the important primitive structure, screen, three cameras, lights and crop metadata. Run it in a fresh factory-startup Blender session; it refuses a scene containing non-default named objects and writes `scenes/pass1-reproduced.blend`, distinct from the reviewed file. Python AST syntax validation passed. It was deliberately not executed over the reviewed scene; runtime reproduction is not claimed as tested.

The reviewed `.blend`, this report and script exist (intermediate blockout renders and `.blend1` auto-save backup files were pruned during repository cleanup). No external assets were downloaded, add-ons installed, GLB exported, frontend application source changed, or files outside this landing-hero directory intentionally written. Git status before/after shows the same pre-existing changes to frontend/bun.lock, frontend/components.json and frontend/next-env.d.ts; these were preserved. No tracked diff outside frontend is present.

Frontend lint/build/test commands were not run for this Blender-only pass: they are unrelated to the artifact and can create caches/generated files outside the user's explicit landing-hero-only write boundary. Verification instead covered visual renders, saved Blender scene state, artifact dimensions and Python syntax.
