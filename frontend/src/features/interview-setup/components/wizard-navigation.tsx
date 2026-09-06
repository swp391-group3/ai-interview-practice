"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { interviewSteps } from "@/config/navigation";
export function WizardNavigation() {
  const pathname = usePathname();
  return (
    <nav aria-label="Interview creation steps">
      <ol className="flex flex-wrap gap-4">
        {interviewSteps.map((step, index) => (
          <li key={step.href}>
            <Link
              href={step.href}
              aria-current={pathname === step.href ? "step" : undefined}
              className="rounded-sm text-sm underline-offset-4 hover:underline aria-[current=step]:font-semibold aria-[current=step]:underline"
            >
              {index + 1}. {step.label}
            </Link>
          </li>
        ))}
      </ol>
    </nav>
  );
}
