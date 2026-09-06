import { expect, it } from "vitest";
import { routes } from "./routes";
it("encodes dynamic route segments and retains correct prefixes", () => {
  expect(routes.interviews.room("a/b ?")).toBe("/interviews/a%2Fb%20%3F/room");
  expect(routes.reports.detail("report#1")).toBe("/reports/report%231");
  expect(routes.dashboard).toBe("/dashboard");
  expect(routes.admin.users).toBe("/admin/users");
});
