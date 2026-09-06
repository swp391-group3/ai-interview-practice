import { routes } from "./routes";
export const publicNavigation = [
  { label: "Home", href: routes.home },
  { label: "Pricing", href: routes.pricing },
  { label: "Log in", href: routes.login },
];
export const candidateNavigation = [
  { label: "Dashboard", href: routes.dashboard },
  { label: "Interviews", href: routes.interviews.root },
  { label: "History", href: routes.history },
  { label: "Billing", href: routes.billing },
  { label: "Profile", href: routes.profile },
  { label: "Settings", href: routes.settings },
];
export const adminNavigation = [
  { label: "Dashboard", href: routes.admin.dashboard },
  { label: "Users", href: routes.admin.users },
  { label: "Interviews", href: routes.admin.interviews },
  { label: "Domains", href: routes.admin.domains },
  { label: "Questions", href: routes.admin.questions },
  { label: "Avatars", href: routes.admin.avatars },
  { label: "Voices", href: routes.admin.voices },
  { label: "Billing", href: routes.admin.billing },
  { label: "Settings", href: routes.admin.settings },
];
export const interviewSteps = [
  { label: "Job description", href: routes.interviews.new.jobDescription },
  { label: "Skills", href: routes.interviews.new.skills },
  { label: "Setup", href: routes.interviews.new.setup },
  { label: "Interviewer", href: routes.interviews.new.interviewer },
  { label: "Preflight", href: routes.interviews.new.preflight },
];
