# Frontend Design Workspace

This directory contains design-phase artifacts that are not yet production application assets.

## Directory Responsibilities

### references/

External visual references, screenshots, mood references, and research material.

These are used for visual study only.

Reference material must not be treated as permission to copy branding or proprietary assets.

- `references/landing/`: Visual references, research material, and mood studies for the landing page.

### exploration/

Temporary generated design artifacts such as:

- OpenDesign HTML prototypes
- layout explorations
- visual direction studies
- design review notes

Artifacts here are exploratory and are NOT canonical product design.

- `exploration/landing/`: Exploratory prototypes, layout variants, and review notes for the landing page.

### 3d/

Blender scenes, scripts, preview renders, and 3D art-direction experiments.

Experimental 3D assets should stay here until approved and optimized.

Do not place experimental assets directly in `public/`.

- `3d/landing-hero/`: 3D scenes, art-direction experiments, and preview renders for the landing hero.

## Source of Truth

Future hierarchy:

```
Product requirements
→ Figma
→ frontend/DESIGN.md
→ frontend/SKILL.md
→ implementation
```

- **Figma** will become the canonical visual source after review.
- `frontend/DESIGN.md` does not exist yet and must NOT be created during this task.
- OpenDesign, Blender, Stitch, and other generation tools are exploration tools, not canonical sources.

## Git & Artifact Guidance

- Do not create generated binaries or large assets yet.
- Do not add Blender files, renders, screenshots, or HTML prototypes during this task.
- Only create the directory structure and README.
