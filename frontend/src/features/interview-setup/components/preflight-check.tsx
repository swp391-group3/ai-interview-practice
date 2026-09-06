"use client";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { getBrowserCapabilities } from "@/features/interview/audio/browser-capabilities";
export function PreflightCheck() {
  const [capabilities, setCapabilities] = useState<ReturnType<
    typeof getBrowserCapabilities
  > | null>(null);
  return (
    <div className="space-y-4">
      <Button onClick={() => setCapabilities(getBrowserCapabilities())}>
        Check browser capabilities
      </Button>
      <p className="text-sm text-muted-foreground">
        Checks API availability only. No camera or microphone is activated.
        Device permissions, network quality and actual recording still need
        validation.
      </p>
      <div aria-live="polite">
        {capabilities && (
          <dl className="grid grid-cols-2 gap-2 text-sm">
            {Object.entries(capabilities).map(([name, available]) => (
              <div key={name}>
                <dt className="font-medium">{name}</dt>
                <dd>{available ? "Available" : "Unavailable"}</dd>
              </div>
            ))}
          </dl>
        )}
      </div>
    </div>
  );
}
