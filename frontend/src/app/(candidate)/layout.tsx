import { SiteShell } from "@/components/layout/site-shell";
import { candidateNavigation } from "@/config/navigation";
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <SiteShell navigation={candidateNavigation} label="Candidate">
      {children}
    </SiteShell>
  );
}
