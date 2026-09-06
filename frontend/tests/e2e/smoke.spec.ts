import { expect, test } from "@playwright/test";
test("public, wizard and admin routes render without a backend", async ({
  page,
}) => {
  await page.goto("/");
  await expect(
    page.getByRole("heading", { name: "Welcome to Interview Practice" }),
  ).toBeVisible();
  await page.goto("/interviews");
  await page
    .getByRole("link", { name: "Create an interview", exact: true })
    .click();
  await expect(page).toHaveURL(/\/interviews\/new\/job-description$/);
  await page.getByRole("link", { name: "2. Skills", exact: true }).click();
  await expect(
    page.getByRole("heading", { name: "Review skills" }),
  ).toBeVisible();
  await page.reload();
  await expect(
    page.getByRole("link", { name: "2. Skills", exact: true }),
  ).toHaveAttribute("aria-current", "step");
  await page.goBack();
  await expect(
    page.getByRole("heading", { name: "Job description" }),
  ).toBeVisible();
  await page.goto("/admin/users");
  await expect(
    page.getByRole("heading", { name: "Users", exact: true }),
  ).toBeVisible();
});
