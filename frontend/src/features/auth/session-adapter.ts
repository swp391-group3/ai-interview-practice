import type { ProductRole } from "@/config/permissions";
import type { AuthHeadersProvider } from "@/lib/api/types";
/** Internal app contract; adapt the eventual backend response at the authority boundary. */
export type Session = { userId: string; role: ProductRole } | null;
export interface SessionAdapter {
  getSession(): Promise<Session>;
  getAuthHeaders: AuthHeadersProvider;
}
/** No implementation is supplied. Route layouts currently provide no access control. */
