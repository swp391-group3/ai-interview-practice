#!/usr/bin/env bash
set -euo pipefail
cd -- "$(dirname -- "$0")/.."
master="renders/final/landing-hero-final-16x9.png"
test "$(magick identify -format '%wx%h' "$master")" = "2560x1440"
# Deliberate crop windows reviewed for this camera, not centered mobile cover.
magick "$master" -crop 2160x1440+200+0 +repage renders/final/landing-hero-final-3x2.png
magick "$master" -crop 1152x1440+870+0 +repage renders/final/landing-hero-final-4x5.png
magick "$master" -resize 1600x900 -quality 86 -define webp:method=6 renders/final/landing-hero-final-preview.webp
