"use client";
import { Canvas } from "@react-three/fiber";
/** Mount via a client dynamic import (ssr: false) when real assets are available. */
export function AvatarStage({ children }: { children: React.ReactNode }) {
  return (
    <div role="img" aria-label="3D interviewer" className="h-80">
      <Canvas
        frameloop="demand"
        fallback={
          <p>3D rendering is unavailable. Use the 2D interview view.</p>
        }
      >
        {children}
      </Canvas>
    </div>
  );
}
