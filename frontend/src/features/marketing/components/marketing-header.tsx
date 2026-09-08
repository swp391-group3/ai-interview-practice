"use client";

import Image from "next/image";
import Link from "next/link";
import { Menu, X } from "lucide-react";
import { useEffect, useState } from "react";
import { routes } from "@/config/routes";
import styles from "./rolecue-landing.module.css";

const navigation = [
  { href: "#practice", label: "Practice" },
  { href: "#method", label: "Method" },
  { href: "#feedback", label: "Feedback" },
] as const;

export function MarketingHeader() {
  const [isCondensed, setIsCondensed] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [isCampaignDismissed, setIsCampaignDismissed] = useState(false);

  useEffect(() => {
    const syncHeader = () => setIsCondensed(window.scrollY > 64);

    syncHeader();
    window.addEventListener("scroll", syncHeader, { passive: true });
    return () => window.removeEventListener("scroll", syncHeader);
  }, []);

  useEffect(() => {
    const closeOnEscape = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setIsMenuOpen(false);
      }
    };

    window.addEventListener("keydown", closeOnEscape);
    return () => window.removeEventListener("keydown", closeOnEscape);
  }, []);

  const closeMenu = () => setIsMenuOpen(false);

  return (
    <>
      <a className={styles.skipLink} href="#main-content">
        Skip to content
      </a>
      <div
        aria-hidden={isCampaignDismissed}
        className={`${styles.campaign} ${
          isCampaignDismissed ? styles.campaignDismissed : ""
        }`}
      >
        <span className={styles.campaignSignal} />
        <strong>RoleCue / Practice</strong>
        <span>Bring a clearer point of view to the next conversation.</span>
        <button
          aria-label="Dismiss announcement"
          className={styles.campaignClose}
          onClick={() => setIsCampaignDismissed(true)}
          tabIndex={isCampaignDismissed ? -1 : 0}
          type="button"
        >
          <X aria-hidden="true" size={15} strokeWidth={1.8} />
        </button>
      </div>
      <header className={styles.siteChrome}>
        <nav
          aria-label="Main navigation"
          className={`${styles.navShell} ${
            isCondensed ? styles.navCondensed : ""
          }`}
        >
          <div className={styles.navInner}>
            <Link
              aria-label="RoleCue home"
              className={styles.brand}
              href={routes.home}
              onClick={closeMenu}
            >
              <Image
                alt="RoleCue"
                height={48}
                priority
                src="/rolecue-logo.svg"
                width={314}
              />
            </Link>
            <div className={styles.desktopNavigation}>
              {navigation.map((item) => (
                <a className={styles.navLink} href={item.href} key={item.href}>
                  {item.label}
                </a>
              ))}
            </div>
            <div className={styles.navEnd}>
              <span className={styles.navNote}>ROLE / RESPONSE / REVIEW</span>
              <Link
                className={styles.navCta}
                href={routes.interviews.new.jobDescription}
              >
                Start practice
              </Link>
              <button
                aria-controls="rolecue-mobile-navigation"
                aria-expanded={isMenuOpen}
                aria-label={isMenuOpen ? "Close navigation" : "Open navigation"}
                className={styles.mobileToggle}
                onClick={() => setIsMenuOpen((open) => !open)}
                type="button"
              >
                {isMenuOpen ? (
                  <X aria-hidden="true" size={19} strokeWidth={1.8} />
                ) : (
                  <Menu aria-hidden="true" size={20} strokeWidth={1.8} />
                )}
              </button>
            </div>
            <div
              className={`${styles.mobilePanel} ${
                isMenuOpen ? styles.mobilePanelOpen : ""
              }`}
              id="rolecue-mobile-navigation"
            >
              {navigation.map((item) => (
                <a href={item.href} key={item.href} onClick={closeMenu}>
                  {item.label}
                  <span aria-hidden="true">↗</span>
                </a>
              ))}
              <Link
                href={routes.interviews.new.jobDescription}
                onClick={closeMenu}
              >
                Start practice
                <span aria-hidden="true">↗</span>
              </Link>
            </div>
          </div>
        </nav>
      </header>
    </>
  );
}
