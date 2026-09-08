import Link from "next/link";
import { ArrowUpRight } from "lucide-react";
import { routes } from "@/config/routes";
import { FaqSection } from "./faq-section";
import { LandingMotion } from "./landing-motion";
import { MarketingHeader } from "./marketing-header";
import { PracticeSection } from "./practice-section";
import styles from "./rolecue-landing.module.css";

const processSteps = [
  {
    number: "01",
    title: "Name the role",
    detail: "Give the conversation its real context before you rehearse it.",
  },
  {
    number: "02",
    title: "Find the example",
    detail: "Make a specific moment do the work of a broad claim.",
  },
  {
    number: "03",
    title: "Keep the cue",
    detail: "Carry the useful adjustment into the next attempt.",
  },
] as const;

const workingModes = [
  {
    number: "01 / Context",
    title: "Start with the work that matters.",
    shape: "square",
  },
  {
    number: "02 / Rehearsal",
    title: "Make room for the answer to change.",
    shape: "circle",
  },
  {
    number: "03 / Review",
    title: "Leave with a better next response.",
    shape: "line",
  },
] as const;

function Arrow() {
  return <ArrowUpRight aria-hidden="true" size={16} strokeWidth={1.8} />;
}

function HeroMedia() {
  return (
    <figure
      aria-label="A RoleCue editorial practice workspace"
      className={styles.heroMedia}
      data-rolecue-reveal
    >
      <div aria-hidden="true" className={styles.heroMediaWash} />
      <div aria-hidden="true" className={styles.heroMediaGrid} />
      <div className={styles.heroMediaWindow}>
        <aside aria-hidden="true" className={styles.heroMediaRail}>
          <div className={styles.windowDots}>
            <i />
            <i />
            <i />
          </div>
          <span className={styles.railActive} />
          <span />
          <span />
          <span />
          <span />
        </aside>
        <div className={styles.heroMediaCanvas}>
          <div className={styles.heroMediaTopline}>
            <span>Practice field / 01</span>
            <span>Role focus</span>
          </div>
          <div className={styles.heroMediaPrompt}>
            <p>Tell us about a decision that changed the work.</p>
            <i aria-hidden="true" />
          </div>
          <div className={styles.heroMediaNote}>
            <span>Keep the detail</span>
            <strong>Start with the constraint, then name the choice.</strong>
          </div>
          <div aria-hidden="true" className={styles.heroMediaHalo} />
          <span aria-hidden="true" className={styles.heroMediaTick} />
        </div>
      </div>
      <figcaption className={styles.mediaCaption}>
        ROLECUE / PRACTICE FIELD
        <br />A quiet screen for a considered answer
      </figcaption>
    </figure>
  );
}

function ProcessArtifact() {
  return (
    <div className={styles.processArtifact} data-rolecue-reveal>
      <div aria-hidden="true" className={styles.processArtifactGrid} />
      <div className={styles.processArtifactPage}>
        <p>RoleCue note / 03</p>
        <strong>
          One clear example is more memorable than a perfect script.
        </strong>
        <span>Save the cue, not the performance.</span>
      </div>
      <span aria-hidden="true" className={styles.processArtifactRing} />
      <span aria-hidden="true" className={styles.processArtifactTick} />
    </div>
  );
}

function ClosingMedia() {
  return (
    <div className={styles.closingMedia} data-rolecue-reveal>
      <div aria-hidden="true" className={styles.closingMediaField} />
      <div className={styles.closingWindow}>
        <div className={styles.closingWindowTopline}>
          <span>RoleCue</span>
          <span>Practice note</span>
        </div>
        <p>Make the next response feel like yours.</p>
        <div className={styles.closingCursor}>
          <span>One thing to carry forward…</span>
          <i aria-hidden="true" />
        </div>
      </div>
    </div>
  );
}

export function RoleCueLanding() {
  return (
    <div className={styles.landing}>
      <LandingMotion />
      <MarketingHeader />
      <main id="main-content">
        <section className={styles.hero} id="top">
          <div aria-hidden="true" className={styles.heroAtmosphere}>
            <span className={styles.heroOrbit} />
            <span className={styles.heroAxis} />
          </div>
          <div className={`${styles.container} ${styles.heroInner}`}>
            <div className={styles.heroFrame} data-rolecue-reveal>
              <i aria-hidden="true" className={styles.frameHandle} />
              <i aria-hidden="true" className={styles.frameHandle} />
              <i aria-hidden="true" className={styles.frameHandle} />
              <i aria-hidden="true" className={styles.frameHandle} />
              <p className={styles.eyebrow}>Technical interview practice</p>
              <h1>
                Practice the role <em>before</em> the room.
              </h1>
              <p className={styles.heroLead}>
                A calmer way to shape the examples, choices, and confidence you
                want to bring to the next technical conversation.
              </p>
            </div>
            <ul
              aria-label="RoleCue practice qualities"
              className={styles.heroTags}
            >
              <li>Role context</li>
              <li>Clear examples</li>
              <li>Useful cues</li>
            </ul>
            <div className={styles.heroActions}>
              <Link
                className={styles.buttonPrimary}
                href={routes.interviews.new.jobDescription}
              >
                Start a practice <Arrow />
              </Link>
              <a className={styles.buttonSecondary} href="#practice">
                See the method
              </a>
            </div>
          </div>
          <HeroMedia />
        </section>

        <section className={styles.editorialSection}>
          <div className={styles.editorialStatement} data-rolecue-reveal>
            <p className={styles.eyebrow}>Less performance, more preparation</p>
            <h2 className={styles.statementTitle}>
              Bring the thinking behind your work into the conversation.
            </h2>
            <p>
              The strongest interview answer does not sound rehearsed. It makes
              a decision legible: what you saw, what you chose, and why it
              mattered.
            </p>
          </div>
        </section>

        <PracticeSection />

        <section className={styles.methodSection} id="method">
          <div className={styles.container}>
            <header className={styles.methodHeader} data-rolecue-reveal>
              <p className={styles.eyebrow}>A usable rhythm</p>
              <h2 className={styles.sectionTitle}>
                From first read to a response you can stand behind.
              </h2>
              <p className={styles.sectionIntro}>
                A small sequence creates enough structure to prepare without
                turning the work into a script.
              </p>
            </header>
            <ol className={styles.processList}>
              {processSteps.map((step) => (
                <li data-rolecue-reveal key={step.number}>
                  <span>{step.number}</span>
                  <h3>{step.title}</h3>
                  <p>{step.detail}</p>
                </li>
              ))}
            </ol>
            <ProcessArtifact />
          </div>
        </section>

        <section className={styles.workingSection}>
          <div className={styles.container}>
            <header className={styles.workingHeader} data-rolecue-reveal>
              <p className={styles.eyebrow}>One private working field</p>
              <h2 className={styles.sectionTitle}>
                Give each part of the conversation its own clear place.
              </h2>
              <p className={styles.sectionIntro}>
                RoleCue keeps the interface quiet so the useful detail can stay
                in view.
              </p>
            </header>
            <div className={styles.workingModes}>
              {workingModes.map((mode) => (
                <article data-rolecue-reveal key={mode.number}>
                  <span
                    aria-hidden="true"
                    className={`${styles.modeShape} ${styles[`mode${mode.shape}`]}`}
                  />
                  <div>
                    <p>{mode.number}</p>
                    <h3>{mode.title}</h3>
                  </div>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section className={styles.reflectionSection}>
          <div className={`${styles.container} ${styles.reflectionGrid}`}>
            <div className={styles.reflectionCopy} data-rolecue-reveal>
              <p className={styles.eyebrow}>The cue stays with you</p>
              <h2 className={styles.sectionTitle}>
                Notice the pattern. Then take it into the next room.
              </h2>
              <p className={styles.sectionIntro}>
                A good review does not turn your experience into a score. It
                gives you a sharper way to recognize what is worth saying.
              </p>
              <Link
                className={styles.buttonSecondary}
                href={routes.interviews.new.jobDescription}
              >
                Prepare a role <Arrow />
              </Link>
            </div>
            <div className={styles.reflectionGraphic} data-rolecue-reveal>
              <span className={styles.reflectionOrbit} />
              <span className={styles.reflectionOrbitInner} />
              <span className={styles.reflectionNode}>ROLE</span>
              <span className={styles.reflectionNode}>CHOICE</span>
              <span className={styles.reflectionNode}>DETAIL</span>
              <span className={styles.reflectionNode}>CUE</span>
              <span className={styles.reflectionCenter}>YOU</span>
            </div>
          </div>
        </section>

        <FaqSection />

        <section className={styles.closingSection}>
          <div className={styles.container}>
            <div className={styles.closingStatement} data-rolecue-reveal>
              <p className={styles.eyebrow}>The next response</p>
              <h2 className={styles.statementTitle}>
                Walk in with more of your own work in reach.
              </h2>
              <p>
                Start with the role. Make space for the answer. Keep the cue
                that makes the next attempt better.
              </p>
              <Link
                className={styles.buttonPrimary}
                href={routes.interviews.new.jobDescription}
              >
                Start a practice <Arrow />
              </Link>
            </div>
            <ClosingMedia />
          </div>
        </section>
      </main>
      <footer className={styles.footer}>
        <div className={styles.container}>
          <div className={styles.footerTop}>
            <div>
              <p className={styles.footerKicker}>Role-focused practice</p>
              <p className={styles.footerWordmark}>
                Role<span>Cue.</span>
              </p>
            </div>
            <nav aria-label="Footer navigation" className={styles.footerLinks}>
              <a href="#practice">Practice</a>
              <a href="#method">Method</a>
              <a href="#feedback">Feedback</a>
              <Link href={routes.interviews.new.jobDescription}>Start</Link>
            </nav>
          </div>
          <div className={styles.footerMeta}>
            <span>© 2026 RoleCue</span>
            <span>Made for clearer technical conversations</span>
          </div>
        </div>
      </footer>
    </div>
  );
}
