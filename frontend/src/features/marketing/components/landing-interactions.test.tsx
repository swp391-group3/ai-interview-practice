import { fireEvent, render, screen } from "@testing-library/react";
import { expect, it } from "vitest";
import { FaqSection } from "./faq-section";
import { PracticeSection } from "./practice-section";

it("switches the active practice stage with tab semantics", () => {
  render(<PracticeSection />);

  const responseTab = screen.getByRole("tab", { name: /try the response/i });
  fireEvent.click(responseTab);

  expect(responseTab).toHaveAttribute("aria-selected", "true");
  expect(screen.getByRole("tabpanel")).toHaveAttribute(
    "data-stage",
    "response",
  );
  expect(
    screen.getByRole("heading", {
      name: "Give the answer room to take shape.",
    }),
  ).toBeVisible();
});

it("opens and closes an FAQ response with accessible disclosure state", () => {
  render(<FaqSection />);

  const question = screen.getByRole("button", {
    name: /is rolecue trying to script an interview/i,
  });
  fireEvent.click(question);

  expect(question).toHaveAttribute("aria-expanded", "true");
  expect(
    screen.getByText(/centers reflection, examples, and clearer choices/i),
  ).toBeVisible();

  fireEvent.click(question);
  expect(question).toHaveAttribute("aria-expanded", "false");
});
