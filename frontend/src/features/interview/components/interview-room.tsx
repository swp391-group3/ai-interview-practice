import { RoutePlaceholder } from "@/components/feedback/route-placeholder";
export function InterviewRoom({ interviewId }: { interviewId: string }) {
  return (
    <RoutePlaceholder
      title="Interview room"
      description="The real-time interview runtime will coordinate conversation, reconnect/resume, audio and the interviewer here."
    >
      <p className="text-sm text-muted-foreground">
        Session reference: {interviewId}
      </p>
      <p className="text-sm">Runtime integration is not connected.</p>
    </RoutePlaceholder>
  );
}
