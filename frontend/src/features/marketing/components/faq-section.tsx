"use client";

import { useId, useState } from "react";
import styles from "./rolecue-landing.module.css";

const questions = [
  {
    question: "What does a practice start with?",
    answer:
      "A role and the context you want to prepare for. The landing points to the existing interview setup flow; it does not claim a backend workflow that is not connected yet.",
  },
  {
    question: "Is RoleCue trying to script an interview?",
    answer:
      "No. The visual and content direction centers reflection, examples, and clearer choices—not a synthetic answer generator or a promise of an outcome.",
  },
  {
    question: "What should feel different after a session?",
    answer:
      "You should leave with a more legible version of your own experience: what matters, which example proves it, and how to bring it into the next conversation.",
  },
  {
    question: "Can I return to a specific part of the page?",
    answer:
      "Yes. The navigation and disclosure controls work with keyboard focus, named anchors, and reduced-motion preferences as well as pointer input.",
  },
] as const;

export function FaqSection() {
  const [openIndex, setOpenIndex] = useState<number | null>(null);
  const listId = useId();

  return (
    <section className={styles.faqSection} id="feedback">
      <div className={styles.container}>
        <div className={styles.faqGrid}>
          <div className={styles.faqIntro} data-rolecue-reveal>
            <p className={styles.eyebrow}>A few useful details</p>
            <h2 className={styles.sectionTitle}>
              Preparation should make the next question feel less strange.
            </h2>
            <p className={styles.sectionIntro}>
              A quiet, practical layer for the things people usually need to
              know before they begin.
            </p>
          </div>
          <div className={styles.faqList} data-rolecue-reveal>
            {questions.map((item, index) => {
              const isOpen = openIndex === index;
              const answerId = `${listId}-${index}`;
              return (
                <article className={styles.faqItem} key={item.question}>
                  <button
                    aria-controls={answerId}
                    aria-expanded={isOpen}
                    className={styles.faqQuestion}
                    onClick={() => setOpenIndex(isOpen ? null : index)}
                    type="button"
                  >
                    <span>{String(index + 1).padStart(2, "0")}</span>
                    <strong>{item.question}</strong>
                    <i aria-hidden="true">+</i>
                  </button>
                  <div
                    className={`${styles.faqAnswerWrap} ${
                      isOpen ? styles.faqAnswerOpen : ""
                    }`}
                    id={answerId}
                  >
                    <div>
                      <p>{item.answer}</p>
                    </div>
                  </div>
                </article>
              );
            })}
          </div>
        </div>
      </div>
    </section>
  );
}
