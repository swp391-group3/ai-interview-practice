import { SiteShell } from "@/components/layout/site-shell";
import { publicNavigation } from "@/config/navigation";
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <SiteShell navigation={publicNavigation} label="Public">
      {children}
    </SiteShell>
  );
}
