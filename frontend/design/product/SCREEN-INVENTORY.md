# RoleCue Figma screen inventory

## Status and source of truth

This is the complete planned Figma frame manifest for `RoleCue — Product Design` (`w1VGwSlSF38RoS7AAAaJ57`). It is grounded in the current Next route structure, the interview state machine, and the approved RoleCue brand guidance.

**Figma implementation status: blocked.** The authenticated Figma Starter View seat reached its monthly MCP tool-call quota during the initial design-system setup. Only the file reorganization and the first primitive color collection were confirmed before the cap. The frames below are the exact inventory to create after access is restored; they must not be represented as existing Figma frames until then.

## Confirmed Figma organization

| Page | Confirmed role | Last confirmed state |
| --- | --- | --- |
| Foundations | Brand system, documentation, components, states | Contains the moved original Cover title block (`RoleCue / Cover title block`) |
| Marketing | Public acquisition and credit-purchase surfaces | Empty |
| Product | Candidate, admin, auth, runtime, and system Sections | Empty after the former Cover page was renamed |

### Foundations Sections

| Section | Planned contents |
| --- | --- |
| 00 Cover | RoleCue design-workspace title block, thesis, file index |
| 01 Brand | Full lockup, icon, wordmark, mono variants, usage notes |
| 02 Color | Primitive and semantic swatches, dark contextual surfaces, semantic states |
| 03 Typography | Display XL through Caption specimens and usage notes |
| 04 Spacing & Grid | 4 px spacing scale; 1440, 1280, and 390 layout rules |
| 05 Radius & Borders | Surface radius, control radius, border hierarchy |
| 06 Shadows & Depth | Quiet elevation and media depth examples |
| 07 Icons | Split Halo usage, navigation/action icon sizing, icon-label rules |
| 08 Motion | Reveal, navigation, media, dialog, state, and reduced-motion rules |
| 09 Components | Component-set inventory below |
| 10 Application States | Empty, loading, error, no-credit, and status patterns |
| 11 Responsive Rules | Desktop/laptop/mobile structural rules |
| 12 Accessibility | AA contrast, focus, touch target, keyboard, and status guidance |
| 13 Design Notes | Handoff notes, unresolved backend contracts, screen naming rules |

### Component family inventory

| Component family | Required useful states or variants |
| --- | --- |
| Brand lockup | Full, icon only, monochrome light, monochrome dark |
| Button | Primary, Secondary, Ghost, Danger; default, hover, disabled, loading where needed |
| Icon button | Default, hover, selected, disabled; tooltip association |
| Input and password input | Default, focus, filled, error, disabled |
| Textarea, Search, Select, Combobox | Default, focus, populated, error/empty as applicable |
| Checkbox, Radio, Switch | Unselected, selected, disabled; focus ring |
| Form field | Label, control, help, error association |
| Tabs, Segmented control, Stepper | Default, selected/current, disabled |
| Badge, Interview status, Credit indicator | Neutral, positive, warning, error/contextual |
| Tooltip, Popover, Dropdown, Context menu | Open and keyboard-focus examples |
| Dialog, Confirmation dialog, Drawer/sheet | Desktop and compact-width behavior |
| Toast/notification | Neutral, success, warning, error |
| Candidate sidebar/top bar/admin navigation | Desktop and collapsed/laptop behavior |
| Breadcrumb | Default and current-page state |
| Card/surface and Editorial feature surface | Quiet default, elevated media use |
| Table, row, pagination, filters | Default, selected, empty, pagination |
| Empty state, skeleton, loading, inline error | Reusable candidate/admin application states |
| Progress indicator | Setup step progress and interview context progress |
| Avatar/interviewer thumbnail | Available, selected, unavailable/degraded context |
| Chart primitives | Competency comparison, time trend, improvement plan action list |

## Marketing frame inventory

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Marketing / Public | Marketing / Landing / Desktop / 1440 | `/` | Primary public narrative, hero media, credits, CTA, footer |
| Marketing / Public | Marketing / Landing / Mobile / 390 | `/` | Intentional mobile public layout |
| Marketing / Public | Marketing / Pricing / Desktop | `/pricing` | Credit purchase explanation; package values intentionally unapproved/placeholders |
| Marketing / Public | Marketing / Pricing / Mobile | `/pricing` | Compact public credits view |

## Product frame inventory

### A — Authentication

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Product / A Authentication | Auth / Login / Desktop | `/login` | Default, focus, validation error, loading, disabled, Google sign-in |
| Product / A Authentication | Auth / Login / Mobile | `/login` | 390 px public/auth composition |
| Product / A Authentication | Auth / Register / Desktop | `/register` | Default, focus, validation error, loading, disabled, Google sign-in |
| Product / A Authentication | Auth / Register / Mobile | `/register` | 390 px public/auth composition |
| Product / A Authentication | Auth / Forgot Password | `/forgot-password` | Default, submitted/confirmation, error |
| Product / A Authentication | Auth / Reset Password | `/reset-password` | Default, password validation error, submitting |
| Product / A Authentication | Auth / Verify Email | Auth verification route/transition | Awaiting verification, resend available |
| Product / A Authentication | Auth / Email Verified | Auth verification route/transition | Confirmed, continue to dashboard |

### B–F — Candidate shell, dashboard, JD, setup, and preflight

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Product / B Candidate shell | Candidate Shell / Desktop 1440 | Candidate route group layout | Expanded navigation, title, credits, account |
| Product / B Candidate shell | Candidate Shell / Laptop 1280 | Candidate route group layout | Condensed navigation and content rail |
| Product / C Dashboard | Candidate / Dashboard | `/dashboard` | Next action, role context, resume/recommendation, credit state |
| Product / C Dashboard | Dashboard / Empty New User | `/dashboard` | No JD or session; calm first action |
| Product / C Dashboard | Dashboard / Returning User | `/dashboard` | Resume/repeat action and recent report recommendation |
| Product / D Job descriptions | JD / Empty State | `/interviews/new/job-description` | No saved description |
| Product / D Job descriptions | JD / Add Job Description | `/interviews/new/job-description` | Paste or upload document; labelled control and validation |
| Product / D Job descriptions | JD / Analyzing | Analysis transition | Calm status, no fabricated percentage |
| Product / D Job descriptions | JD / Analysis Result | Analysis transition | Role context, extracted technical skills |
| Product / D Job descriptions | JD / Review & Edit Skills | `/interviews/new/skills` | Skills, focus, priority/relevance editing |
| Product / E Interview setup | Interview Setup / General | `/interviews/new/setup` | Difficulty, duration, focus areas, sensible defaults |
| Product / E Interview setup | Interview Setup / Interviewer | `/interviews/new/interviewer` | Curated interviewer selection |
| Product / E Interview setup | Interview Setup / Voice | `/interviews/new/interviewer` | Voice choice as part of interviewer decision |
| Product / E Interview setup | Interview Setup / Credits Gate | Setup transition | Insufficient credits; transparent purchase route |
| Product / F Preflight | Preflight / Checking | `/interviews/new/preflight` | Capability checks in progress |
| Product / F Preflight | Preflight / Ready | `/interviews/new/preflight` | Required checks clear; optional status explained |
| Product / F Preflight | Preflight / Permission Required | `/interviews/new/preflight` | User-triggered device permission guidance |
| Product / F Preflight | Preflight / Failed Check | `/interviews/new/preflight` | Explain retry, browser/device alternatives, no proctoring tone |

### G — Interview runtime

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Product / G Interview room | Interview / Connecting | `/interviews/[interviewId]/room` | Establishing session |
| Product / G Interview room | Interview / Asset Loading | `/interviews/[interviewId]/room` | Preparing interviewer stage |
| Product / G Interview room | Interview / Interviewer Speaking | `/interviews/[interviewId]/room` | Avatar stage, question context, listening controls inactive |
| Product / G Interview room | Interview / Listening | `/interviews/[interviewId]/room` | Candidate speaking/listening state and clear stop control |
| Product / G Interview room | Interview / Processing Answer | `/interviews/[interviewId]/room` | Brief response processing state |
| Product / G Interview room | Interview / Transitioning | `/interviews/[interviewId]/room` | Quiet question transition |
| Product / G Interview room | Interview / Paused | `/interviews/[interviewId]/room` | Resume and end-session confirmation path |
| Product / G Interview room | Interview / Reconnecting | `/interviews/[interviewId]/room` | Session-safe language only where confirmed by contract |
| Product / G Interview room | Interview / Degraded 2D | `/interviews/[interviewId]/room` | Explicit 2D interviewer fallback acknowledgement |
| Product / G Interview room | Interview / Completed | `/interviews/[interviewId]/room` | Proceed to evaluation |
| Product / G Interview room | Interview / Terminated Early | `/interviews/[interviewId]/room` | Clear consequence and next step |
| Product / G Interview room | Interview / Fatal Error | `/interviews/[interviewId]/room` | Plain recovery guidance, no engineering jargon |

### H–L — Evaluation, reports, history, billing, and profile

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Product / H Evaluation | Evaluation / Processing | Post-interview transition | Meaningful status, no fake percentage |
| Product / I Reports | Report / Overview | `/reports/[reportId]` | Strengths → gaps → next action hierarchy |
| Product / I Reports | Report / Competency Detail | `/reports/[reportId]` detail state | Technical and reasoning dimensions with restrained charting |
| Product / I Reports | Report / Improvement Plan | `/reports/[reportId]` action state | Concrete practice plan and repeat action |
| Product / J History | History / Interview List | `/history`, `/interviews` | Practical filters and report links |
| Product / J History | History / Empty | `/history` | First-practice path |
| Product / J History | History / Session Summary | `/history` selected session | View report, practice again, reuse JD context |
| Product / K Billing | Billing / Credits Overview | `/billing` | Current balance, transparent use history |
| Product / K Billing | Billing / Purchase Credits | `/billing` purchase state | Unapproved package values marked design placeholders |
| Product / K Billing | Billing / Payment Processing | Billing transition | No gateway branding assumed |
| Product / K Billing | Billing / Payment Success | Billing transition | Confirmation and updated balance context |
| Product / K Billing | Billing / Payment Failed | Billing transition | Explain retry and payment-method route without blame |
| Product / K Billing | Billing / Transaction History | `/billing` | Creditable, readable ledger |
| Product / L Profile | Profile / Personal Information | `/profile` | Candidate account details |
| Product / L Profile | Profile / Preferences | `/settings` | Interview/product preferences only |
| Product / L Profile | Profile / Security | `/profile`, `/settings` | Password and account security route |

### M–T — Admin

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Product / M Admin shell | Admin Shell / Desktop | `/admin` layout | Dense but calm operational navigation |
| Product / N Admin dashboard | Admin / Dashboard | `/admin/dashboard` | Operational sessions, service health context, pending review |
| Product / O Admin users | Admin / Users / List | `/admin/users` | Search, filter, account status, credits, session context |
| Product / O Admin users | Admin / Users / Detail | `/admin/users` selected user | Relevant account/status/credit/interview context |
| Product / P Admin sessions | Admin / Sessions / List | `/admin/interviews` | Status, candidate, role, duration, report availability, technical-error flag |
| Product / P Admin sessions | Admin / Session / Detail | `/admin/interviews` selected session | Session metadata; no fabricated transcript |
| Product / Q Admin technical content | Admin / Technical Domains | `/admin/domains` | Domains and skill configuration |
| Product / Q Admin technical content | Admin / Question / Content List | `/admin/questions` | Question/configuration content list |
| Product / Q Admin technical content | Admin / Content Edit | `/admin/questions` editing state | Structured question/configuration form |
| Product / R Admin avatars and voices | Admin / Avatars | `/admin/avatars` | Interviewer asset catalog |
| Product / R Admin avatars and voices | Admin / Avatar Detail/Edit | `/admin/avatars` detail state | Asset metadata and availability |
| Product / R Admin avatars and voices | Admin / Voices | `/admin/voices` | Voice catalog |
| Product / R Admin avatars and voices | Admin / Voice Detail/Edit | `/admin/voices` detail state | Voice metadata and availability |
| Product / S Admin billing | Admin / Billing Overview | `/admin/billing` | Credit/payment operations overview |
| Product / S Admin billing | Admin / Transactions | `/admin/billing` transaction state | Filtered operational ledger |
| Product / T Admin settings | Admin / Settings | `/admin/settings` | Meaningful operational settings only |

### U — System and edge states

| Figma section | Frame name | Route/product mapping | Major state |
| --- | --- | --- | --- |
| Product / U System states | System / 403 Permission Denied | Route/error boundary | Ask user to return to an allowed area |
| Product / U System states | System / 404 Not Found | `not-found.tsx` | Clear route recovery |
| Product / U System states | System / 500 Unexpected Error | `global-error.tsx`, candidate error | Plain retry/recovery guidance |
| Product / U System states | System / Offline Connection Issue | Network-dependent routes | Show saved-context/retry language where possible |
| Product / U System states | System / Generic Empty | Reusable app state | Explain absence and next action |
| Product / U System states | System / Generic Skeleton | Reusable app state | Content-shape loading treatment |
| Product / U System states | System / Generic Inline Error | Reusable app state | Local, recoverable form/data error |
| Product / U System states | System / No Credits | Billing/setup gate | Clear purchase or return route |
| Product / U System states | System / No Job Descriptions | JD/history entry state | Start JD flow |
| Product / U System states | System / No Interview History | History state | Start practice flow |
| Product / U System states | System / No Report Available | Reports/history state | Explain evaluation availability and recovery |

## Counts

- Foundations documentation Sections: 14
- Component families: 20
- Marketing frames: 4
- Product screens/states: 80
- Total planned product and marketing screen frames: 84

The 84-screen total excludes component variants and documentation specimens.
