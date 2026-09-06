export type InterviewListFilters = Readonly<{
  status?: "active" | "completed";
  page?: number;
}>;
export const interviewKeys = {
  all: ["interviews"] as const,
  lists: () => [...interviewKeys.all, "list"] as const,
  list: (filters: InterviewListFilters) =>
    [...interviewKeys.lists(), filters] as const,
  details: () => [...interviewKeys.all, "detail"] as const,
  detail: (id: string) => [...interviewKeys.details(), id] as const,
};
