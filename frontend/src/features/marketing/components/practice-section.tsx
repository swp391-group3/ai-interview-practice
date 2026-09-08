"use client";

import { useId, useState } from "react";
import styles from "./rolecue-landing.module.css";

const practiceStages = [
  {
    key: "role",
    number: "01",
    label: "Read the role",
    title: "Find the signal beneath the job description.",
    detail:
      "Start with the decisions, technical context, and point of view the conversation is really trying to surface.",
    visualLabel: "Role signal",
  },
  {
    key: "response",
    number: "02",
    label: "Try the response",
    title: "Give the answer room to take shape.",
    detail:
      "Work through one focused prompt at a time, with enough quiet to notice what is true, vague, or missing.",
    visualLabel: "Response loop",
  },
  {
    key: "review",
    number: "03",
    label: "Keep the cue",
    title: "Turn each attempt into a more useful next one.",
    detail:
      "Save the patterns worth carrying forward: clearer examples, sharper structure, and the question behind the question.",
    visualLabel: "Review note",
  },
] as const;

export function PracticeSection() {
  const [activeIndex, setActiveIndex] = useState(0);
  const panelId = useId();
  const activeStage = practiceStages[activeIndex];

  return (
    <section className={styles.practiceSection} id="practice">
      <div className={styles.container}>
        <div className={styles.practiceIntro} data-rolecue-reveal>
          <p className={styles.eyebrow}>A calmer preparation loop</p>
          <h2 className={styles.sectionTitle}>
            The best answer starts before you start speaking.
          </h2>
          <p className={styles.sectionIntro}>
            RoleCue makes a private space for the kind of preparation that feels
            more like thinking than performing.
          </p>
        </div>
        <div className={styles.practiceGrid}>
          <div className={styles.practiceCopy} data-rolecue-reveal>
            <p className={styles.eyebrow}>The practice</p>
            <div
              aria-label="Practice stages"
              className={styles.practiceTabs}
              role="tablist"
            >
              {practiceStages.map((stage, index) => {
                const tabId = `${panelId}-${stage.key}-tab`;
                return (
                  <button
                    aria-controls={panelId}
                    aria-selected={activeIndex === index}
                    className={styles.practiceTab}
                    id={tabId}
                    key={stage.key}
                    onClick={() => setActiveIndex(index)}
                    role="tab"
                    type="button"
                  >
                    <span>{stage.number}</span>
                    <strong>{stage.label}</strong>
                    <i aria-hidden="true">↗</i>
                  </button>
                );
              })}
            </div>
            <div className={styles.practiceDetail}>
              <h3>{activeStage.title}</h3>
              <p>{activeStage.detail}</p>
            </div>
          </div>
          <div
            aria-labelledby={`${panelId}-${activeStage.key}-tab`}
            className={styles.practiceStage}
            data-rolecue-reveal
            data-stage={activeStage.key}
            id={panelId}
            role="tabpanel"
          >
            <div aria-hidden="true" className={styles.practiceStageGrid} />
            <div aria-hidden="true" className={styles.practiceHalo} />
            <div aria-hidden="true" className={styles.practiceTick} />
            <div className={styles.practiceWindow}>
              <div className={styles.windowChrome}>
                <span />
                <span />
                <span />
                <p>{activeStage.visualLabel}</p>
              </div>
              <p className={styles.stageNumber}>{activeStage.number}</p>
              <p className={styles.stagePrompt}>
                What do you want the interviewer to remember?
              </p>
              <div className={styles.stageUnderline} />
              <p className={styles.stageFootnote}>
                A cue for the next response
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
