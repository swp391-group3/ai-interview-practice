import { SiteShell } from "@/components/layout/site-shell";
import { adminNavigation } from "@/config/navigation";
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <SiteShell navigation={adminNavigation} label="Admin">
      {children}
    </SiteShell>
  );
}
