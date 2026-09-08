"use client";

import { useEffect } from "react";

const visibleClass = "is-visible";

export function LandingMotion() {
  useEffect(() => {
    const root = document.documentElement;
    const targets = Array.from(
      document.querySelectorAll<HTMLElement>("[data-rolecue-reveal]"),
    );
    const reducedMotion = window.matchMedia(
      "(prefers-reduced-motion: reduce)",
    ).matches;

    const reveal = (element: HTMLElement) => {
      element.classList.add(visibleClass);
    };

    if (reducedMotion || !("IntersectionObserver" in window)) {
      targets.forEach(reveal);
      return;
    }

    const pendingTargets = targets.filter((element) => {
      const rect = element.getBoundingClientRect();
      const isInitiallyVisible = rect.top < window.innerHeight * 0.9;

      if (isInitiallyVisible) {
        reveal(element);
      }

      return !isInitiallyVisible;
    });

    root.dataset.rolecueMotion = "ready";

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            reveal(entry.target as HTMLElement);
            observer.unobserve(entry.target);
          }
        });
      },
      { rootMargin: "0px 0px -8% 0px", threshold: 0.14 },
    );

    pendingTargets.forEach((element) => observer.observe(element));

    return () => {
      observer.disconnect();
      delete root.dataset.rolecueMotion;
    };
  }, []);

  return null;
}
