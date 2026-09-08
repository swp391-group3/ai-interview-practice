# Landing Hero Asset Ledger

Acquisition date: 2026-09-08. Production sources only; no brand artwork or real-person likenesses.

| Asset | Source / direct URL | Provider | License | Usage | Attribution |
| --- | --- | --- | --- | --- | --- |
| MPFB2 v2.0.17 source and bundled core data | https://github.com/makehumancommunity/mpfb2/tree/v2.0.17 ; download https://github.com/makehumancommunity/mpfb2/archive/refs/tags/v2.0.17.zip | MakeHuman Community | Code GPL; core graphics CC0 per https://static.makehumancommunity.org/about/license.html | Original candidate generation and rig; upstream provenance documented above | Core graphics: not required. Retain software license with upstream source. |

The original room, props, temporary screen artwork and custom lookdev are authored for this project. Additional acquisitions will be recorded here before use.

| Asset | Source / direct URL | Provider | License | Usage | Attribution |
| --- | --- | --- | --- | --- | --- |
| MakeHuman system assets pack | https://static.makehumancommunity.org/assets/assetpacks/makehuman_system_assets.html ; https://files2.makehumancommunity.org/asset_packs/makehuman_system_assets/makehuman_system_assets_cc0.zip | MakeHuman Community / makehuman_system | CC0, individually listed on official pack page | Selected casual clothing, eyes, skin and short hair only; unused pack assets are not scene content | Not required |
| Wood Table 001, 2K diffuse / roughness / GL normal | https://polyhaven.com/a/wood_table_001 ; https://dl.polyhaven.org/file/ph-assets/Textures/jpg/2k/wood_table_001/wood_table_001_diff_2k.jpg ; same directory `wood_table_001_rough_2k.jpg`, `wood_table_001_nor_gl_2k.jpg` | Poly Haven; Dimitrios Savva (photography), Rico Cilliers (processing) | CC0, https://polyhaven.com/license | Desk and storage wood, scale adapted from 1.5 m scan | Not required |

System pack mirror actually used: https://files.makehumancommunity.org/asset_packs/makehuman_system_assets/makehuman_system_assets_cc0.zip (primary mirror was too slow). MPFB extension installed in user_default; no unrelated preferences deliberately changed.

## Selected core assets actually used

All acquired on 2026-09-08 from the system pack above, by `makehuman_system`, CC0, no attribution required. Direct acquisition URL and official licensing page are in the pack row.

| Pack-relative asset | Intended / actual usage |
| --- | --- |
| `clothes/male_casualsuit01/` | Unbranded casual shirt and jeans; fitted to the original body and seated rig; diffuse desaturated/darkened, roughness and cloth sheen adjusted. |
| `clothes/shoes01/` | Seated footwear, largely below the hero frame. |
| `hair/short02/` | Short hair base; specular/normal response reduced, color adjusted, supplemented by original directional curve groom. |
| `skins/young_asian_male/young_lightskinned_male_diffuse3.png` | Skin color map on the original mixed-macro body; custom restrained SSS and roughness. Asset filename is provenance, not an assertion about a real person. |
| `eyes/high-poly/` | Fitted core eye geometry. |
| `eyebrows/eyebrow001/` | Fitted eyebrow geometry/material. |

Other extracted casual clothing and short-hair variants were inspected as source options, not used in the final scene. No Mixamo, BlenderKit, Sketchfab or ambientCG assets were acquired. No HDRI was needed. Poly Haven remains the sole external environment-material source.

## Original artwork

`assets/props/interview-screen.svg` and its raster derivative are authored for this scene, with generic session/response labels and an original neutral illustrated interviewer. They contain no copied product artwork, real employer, score, transcript, testimonial or personal data. SVG is the editable source. The scene's abstract wall print, leaf blades, strand groom, prop detailing, architecture, light rig and procedural material additions are original work.

## Portable scene
 
Image textures are packed in the final Blender file. The MPFB-generated meshes and armature are stored in the scene; viewing/rendering does not require downloading the source pack again. Downloadable vendor software packages and archive mirrors were removed during repository cleanup to keep git lean; upstream MPFB software source and licenses (`LICENSE.CODE.md`, `LICENSE.ASSETS.md`) are documented at https://github.com/makehumancommunity/mpfb2/tree/v2.0.17 and can be re-downloaded if needed. The GPL software license does not replace the CC0 license on the core graphics.
 
## Acquisition checksums
 
SHA-256:
 
- System pack (`system-mirror.zip`): `b542127a8e25547c7c29c19f2d1d2adb9a664c80396ecd694095dbc8028a0107`
- MPFB source release: `d08e726c798fdc4eefb02b06b6c4efe37d40b5439777e53cf96dce0e5073297d`
- Wood diffuse 2K: `63ae5cd186197b40f18bc020fe2b652bb08df4d9fc6240c6cbf8a7e3bc32c096`
 
Downloaded vendor archives and unused clothing/hair option variants were pruned during repository cleanup. Selected core assets used by the scene are retained, and all image textures are packed in the final deliverable scene.
