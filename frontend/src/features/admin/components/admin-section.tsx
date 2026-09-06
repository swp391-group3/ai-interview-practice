import { RoutePlaceholder } from "@/components/feedback/route-placeholder";
export function AdminSection({
  title,
  description,
}: {
  title: string;
  description: string;
}) {
  return (
    <RoutePlaceholder title={title} description={description}>
      <p className="text-sm text-muted-foreground">
        Administration is not connected. This route is not access-controlled
        yet.
      </p>
    </RoutePlaceholder>
  );
}
