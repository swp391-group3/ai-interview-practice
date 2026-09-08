# Landing Hero 3D

## Goal

Create a premium game-cinematic 3D scene for the landing page of an AI Virtual Technical Interview Platform.

The scene should communicate:

> "A job candidate is preparing for and participating in a realistic AI technical interview from their own workspace."

## Target Feeling

- AAA-inspired visual quality
- cinematic
- human
- warm
- aspirational
- professional
- realistic but not uncanny
- bright enough to integrate with the editorial landing design

## Current Mental Model

Candidate in a carefully designed home office / personal workspace.

The scene tells a small emotional story:

preparing for an interview  
→ opening the interview web application  
→ speaking with the AI interviewer  
→ confidence / relief / optimism  

The AI interviewer may appear through the application displayed on the candidate's monitor/laptop rather than physically occupying the room.

## Directory Structure & Purposes

```text
design/3d/landing-hero/
├── README.md
├── references/
├── assets/
│   ├── characters/
│   ├── environment/
│   ├── props/
│   ├── materials/
│   └── hdri/
├── scripts/
├── scenes/
├── renders/
│   ├── blockout/
│   ├── lookdev/
│   └── final/
└── exports/
```

- **`references/`**: Visual references, composition studies, lighting references, screenshots, mood references, and external research. Inspiration/research only; do not treat third-party reference material as production assets.
- **`assets/characters/`**: Character source files used during exploration (e.g., Mixamo, MakeHuman / MPFB, properly licensed downloadable character assets). Do not download anything yet.
- **`assets/environment/`**: Reusable room/environment geometry.
- **`assets/props/`**: Furniture and small scene props such as desk, chair, monitor, laptop, lamp, books, plant.
- **`assets/materials/`**: Locally stored PBR material assets when required.
- **`assets/hdri/`**: HDRIs used for look development and lighting.
- **`scripts/`**: Deterministic Blender Python scripts to recreate structural scene setups programmatically without relying purely on manual Blender operations.
- **`scenes/`**: `.blend` working files. Expected future naming pattern: `interview-room-blockout.blend`, `interview-room-lookdev.blend`, `interview-room-final.blend`. Do not create those files yet.
- **`renders/blockout/`**: Composition and camera tests.
- **`renders/lookdev/`**: Material and lighting development renders.
- **`renders/final/`**: Approved final hero renders.
- **`exports/`**: Future optimized delivery assets (GLB, WebP/AVIF, video). Do not export anything yet.

## Production Strategy

Expected approach:

**Hybrid asset pipeline**

- custom Blender environment/composition
- reusable licensed/base character
- curated props/assets
- professional lighting/lookdev
- Blender/Codex-assisted iteration

Do not commit to a specific character source yet. That decision belongs to the asset-strategy research phase.

## Landing Integration

The scene is being created specifically for a Next.js landing page.

Design for:

- large 16:9 hero media treatment
- responsive cropping
- desktop/laptop priority
- possible still-image delivery
- possible video loop
- possible optimized realtime R3F/GLB delivery later

Do not assume realtime 3D is required. Performance and visual quality will determine the final delivery mode.

## Source of Truth

- **Figma** will eventually become the canonical visual composition.
- **Blender** contains the canonical 3D source.
- `DESIGN.md` will later document final production rules.

## Asset Licensing

Every externally sourced production asset must have:

- source URL
- author/provider where relevant
- license
- usage restrictions
- date acquired

Do not use an asset in production if its license cannot be verified.
