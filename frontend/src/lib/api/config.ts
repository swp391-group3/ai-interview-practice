import { getEnvironment } from "@/config/env";
export function getApiBaseUrl() {
  return getEnvironment().apiBaseUrl;
}
