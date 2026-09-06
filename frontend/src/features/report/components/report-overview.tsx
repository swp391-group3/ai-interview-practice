import { RoutePlaceholder } from "@/components/feedback/route-placeholder";
export function ReportOverview({ reportId }: { reportId: string }) {
  return (
    <RoutePlaceholder
      title="Performance report"
      description="Evaluation processing status, technical feedback and performance insights will appear here."
    >
      <p className="text-sm text-muted-foreground">
        Report reference: {reportId}
      </p>
    </RoutePlaceholder>
  );
}
