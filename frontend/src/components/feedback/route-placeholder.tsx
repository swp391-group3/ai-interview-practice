import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
export function RoutePlaceholder({
  title,
  description,
  children,
}: {
  title: string;
  description: string;
  children?: React.ReactNode;
}) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>
          <h1 className="text-2xl font-semibold tracking-tight">{title}</h1>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-6">
        <p className="max-w-2xl text-muted-foreground">{description}</p>
        {children}
      </CardContent>
    </Card>
  );
}
