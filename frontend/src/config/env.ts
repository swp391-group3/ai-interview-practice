import { z } from "zod";
const environmentSchema = z.object({ apiBaseUrl: z.url().optional() });
export function getEnvironment() {
  return environmentSchema.parse({
    apiBaseUrl: process.env.NEXT_PUBLIC_API_BASE_URL || undefined,
  });
}
