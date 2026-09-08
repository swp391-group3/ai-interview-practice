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

- `exploration/open-design-study/`: Reference reconstruction, visual and motion audits, and extracted tokens.
- `exploration/datn/`: Approved final landing design output and design system handoff specifications.

## Source of Truth

Entry point for UI agents:
- Consult `design/DESIGN-CONTRACT.md` before implementing or modifying UI.

Future hierarchy:

```
Product requirements
→ design/DESIGN-CONTRACT.md
→ Figma / Approved Design Artifacts
→ frontend/SKILL.md
→ implementation
```

- Realtime 3D / Blender exploration is currently DEFERRED; do not introduce 3D assets during landing implementation.
- OpenDesign, Stitch, and other generation tools are exploration tools, not canonical sources.

## Git & Artifact Guidance

- Do not create generated binaries or large cache assets in git.
- Do not commit crawler caches (`RECON/`), installer archives, or unapproved temporary prototypes.

