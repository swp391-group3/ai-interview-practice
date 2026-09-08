# Specification: AI-Powered Technical Interview Simulation Platform with 3D Virtual Interviewer

**Project Code**: `09_GFA26SE84`  
**Document Version**: `1.1.0` (Monetization removed)  
**Status**: `Ready for Review (Phase 1: Specify)`  
**Authors**: Capstone Group 09 (Tô Chí Bảo, Huỳnh Minh Khang, Nguyễn Huỳnh Nhật Anh, Nguyễn Tấn Trọng, Đặng Phương Nam)  
**Supervisor**: Mr. Nguyễn Thế Hoàng  

---

## 1. Executive Summary & Objective

### 1.1 Problem Statement
Technical interview preparation is high-stakes yet poorly served by traditional static question banks and generic mock interview platforms:
1. **Lack of JD Personalization:** Practice questions fail to target specific employer Job Descriptions (JDs), technical stacks, or seniority levels.
2. **Limited Availability of Interviewers:** Candidates lack access to experienced senior engineers for continuous, stress-free mock interviews.
3. **Static & Non-Adaptive Interactivity:** Static quizzes cannot drill down into candidate responses with contextual technical follow-up questions.
4. **Superficial Feedback:** Candidates receive pass/fail marks or generic scores without granular competency breakdowns, root-cause knowledge gap identification, or actionable learning roadmaps.

### 1.2 Proposed Solution
An end-to-end web platform featuring:
- **Intelligent JD Parser & Blueprint Generator:** Extracts required skills, frameworks, tools, and seniority expectations from raw text or PDF JDs to formulate a custom interview blueprint.
- **Real-Time Voice-Driven 3D AI Interviewer:** A WebGL/Three.js 3D avatar with real-time Speech-to-Text (STT), Text-to-Speech (TTS), and synchronized blend-shape viseme lip-sync that asks adaptive questions and dynamic follow-ups.
- **Automated Multi-Dimensional Evaluation Engine:** Post-interview scoring across 5 core competencies (Technical Accuracy, Depth of Understanding, Problem-Solving, Answer Relevance, and Communication Clarity) with domain radar charts and question-by-question improvement recommendations.
- **Session History & Administrative Governance:** Session monitoring, candidate management, technical domain taxonomies, avatar catalogs, and platform analytics.

### 1.3 Target Personas
1. **Candidate:** Prepares for specific job openings by uploading JDs, configuring interview sessions, talking to 3D virtual avatars, and analyzing performance reports.
2. **Admin:** Manages platform health, candidate accounts, technical domain tags, avatar/voice catalog, difficulty presets, and aggregate skill weakness analytics.

---

## 2. Tech Stack & Commands

### 2.1 Technology Matrix

| Layer | Technology | Rationale & Version |
|---|---|---|
| **Frontend Framework** | Next.js 15+ (App Router, React 19, TypeScript) | SSR/SSG, fast routing, optimized asset streaming |
| **Styling & UI** | Tailwind CSS v4, shadcn/ui, Lucide Icons | Accessible, modular design system |
| **3D Avatar & Graphics** | Three.js, `@react-three/fiber`, `@react-three/drei` | Real-time GLTF/GLB avatar rendering and blend-shape animation |
| **Speech & Audio** | Web Audio API, MediaRecorder API | Real-time candidate audio capture, VAD (Voice Activity Detection) |
| **Backend Framework** | Golang (Go 1.23+), Gin / Chi Web Framework | High concurrency, low latency, robust WebSocket handling |
| **Database Access & Migration** | PostgreSQL 16+, `pgx` / sqlc / golang-migrate | ACID compliance, relational integrity, JSONB support for blueprints |
| **Cache & Queue** | Redis 7+ | Session state, rate limiting, pub/sub for WebSocket events |
| **AI LLM Orchestration** | OpenAI / Gemini API / Claude 3.5 Sonnet | Structured output parsing (JSON schema), adaptive dialog generation |
| **Speech-to-Text (STT)** | OpenAI Whisper API / Deepgram Nova-2 | High-accuracy technical transcription with latency < 500ms |
| **Text-to-Speech (TTS)** | Azure Speech / ElevenLabs / OpenAI TTS + Rhubarb Visemes | High fidelity audio generation with phoneme/viseme timestamps for 3D mouth morphs |
| **Container & CI/CD** | Docker, Docker Compose, GitHub Actions | Reproducible environments and automated test gates |

### 2.2 Executable Commands

#### Backend (Golang)
```bash
# Development & Hot Reload
go run cmd/server/main.go
# Or with Air for hot reload:
air

# Run all unit and integration tests with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Linting & Static Analysis
golangci-lint run --timeout=5m

# Database Migrations (golang-migrate)
migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/ai_interview?sslmode=disable" up
migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/ai_interview?sslmode=disable" down 1

# Build Production Binary
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/server cmd/server/main.go
```

#### Frontend (Next.js)
```bash
# Install dependencies
pnpm install

# Start local development server
pnpm dev

# Type check & Lint
pnpm typecheck
pnpm lint --fix

# Run Tests
pnpm test
pnpm test:coverage

# Run End-to-End Tests (Playwright)
pnpm test:e2e

# Build Production Bundle
pnpm build
pnpm start
```

#### Full Stack Infrastructure (Docker Compose)
```bash
# Start Postgres, Redis, Backend, and Frontend
docker compose -f docker-compose.dev.yml up --build -d

# Stop all containers
docker compose -f docker-compose.dev.yml down -v
```

---

## 3. Project Structure

```
graduation-thesis/
├── .github/
│   └── workflows/
│       ├── backend-ci.yml           # Go linting, tests, build
│       └── frontend-ci.yml          # Next.js linting, typecheck, tests
├── docker-compose.dev.yml           # Local dev orchestration (Postgres, Redis, Mock STT/TTS)
├── docker-compose.prod.yml          # Production deployment spec
├── Makefile                         # Unified root automation tasks
├── backend/                         # Golang Backend Service
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Entrypoint
│   ├── internal/
│   │   ├── domain/                  # Core entities, value objects, interfaces
│   │   │   ├── user.go
│   │   │   ├── job_description.go
│   │   │   ├── interview_session.go
│   │   │   └── evaluation.go
│   │   ├── usecase/                 # Business logic handlers
│   │   │   ├── auth_usecase.go
│   │   │   ├── jd_usecase.go
│   │   │   ├── interview_orchestrator.go
│   │   │   └── evaluation_usecase.go
│   │   ├── repository/              # PostgreSQL data access layer
│   │   │   ├── postgres/
│   │   │   └── redis/
│   │   ├── delivery/
│   │   │   ├── http/                # REST HTTP routes & controllers
│   │   │   │   ├── router.go
│   │   │   │   ├── auth_handler.go
│   │   │   │   ├── jd_handler.go
│   │   │   │   ├── interview_handler.go
│   │   │   │   └── admin_handler.go
│   │   │   └── ws/                  # Real-time WebSocket connection manager
│   │   │       ├── hub.go
│   │   │       └── interview_socket.go
│   │   ├── platform/                # External adapters
│   │   │   ├── llm/                 # OpenAI / Gemini API client
│   │   │   ├── stt/                 # Whisper / Deepgram speech recognition
│   │   │   └── tts/                 # Speech synthesis + viseme extractors
│   │   └── middleware/              # JWT auth, RBAC, CORS, rate limiting
│   ├── migrations/                  # golang-migrate SQL migrations
│   ├── pkg/                         # Reusable utilities (logger, errors, validator)
│   ├── go.mod
│   └── go.sum
├── frontend/                        # Next.js 15 Web Application
│   ├── src/
│   │   ├── app/                     # App Router pages
│   │   │   ├── (auth)/              # Login, Register, Forgot Password
│   │   │   ├── (candidate)/         # Candidate portal
│   │   │   │   ├── dashboard/       # Overview, recent interviews, quick start
│   │   │   │   ├── jd/              # JD upload, skill extraction, review
│   │   │   │   ├── interview/
│   │   │   │   │   ├── setup/       # Difficulty, avatar selection, config
│   │   │   │   │   └── [id]/room/   # 3D Avatar interview room (WebSocket)
│   │   │   │   ├── history/         # Past sessions list
│   │   │   │   └── report/[id]/     # AI Evaluation & Radar chart report
│   │   │   ├── (admin)/             # Admin portal
│   │   │   │   ├── dashboard/       # System metrics, active sessions
│   │   │   │   ├── users/           # Candidate management & locking
│   │   │   │   ├── sessions/        # Session audit logs & status tracker
│   │   │   │   ├── domains/         # Skill & domain taxonomy manager
│   │   │   │   └── avatars/         # 3D Avatar & voice persona settings
│   │   │   ├── api/                 # Next.js edge / proxy routes (if needed)
│   │   │   └── layout.tsx
│   │   ├── components/
│   │   │   ├── 3d/                  # Three.js Canvas, AvatarLoader, LipSyncController
│   │   │   ├── audio/               # Visualizer, MicPermissionModal, AudioRecorder
│   │   │   ├── interview/           # QuestionCounter, Timer, TranscriptView
│   │   │   ├── report/              # RadarChart, CompetencyCard, FeedbackAccordion
│   │   │   └── ui/                  # shadcn/ui components (Button, Modal, etc.)
│   │   ├── hooks/                   # useWebSocket, useAudioRecorder, useAvatarLipSync
│   │   ├── lib/                     # API client (Axios/Ky), JWT token helper, utils
│   │   ├── types/                   # TypeScript interfaces (DTOs, WS packets)
│   │   └── public/
│   │       ├── models/              # 3D Avatar GLTF/GLB models
│   │       └── textures/            # Avatar textures and environment maps
│   ├── package.json
│   ├── tsconfig.json
│   └── tailwind.config.ts
└── docs/                            # Capstone Reports & Architecture Specs
```

---

## 4. Code Style & Engineering Conventions

### 4.1 Golang Code Conventions
- **Error Handling:** Use the project's `apperror` package for application and API errors. Do not expose internal errors directly to API consumers.
- **Database Access:** Use `sqlc`-generated database access code for PostgreSQL queries.
- **Migration:** Use `golang-migrate` for PostgreSQL schema migrations.
- **Implementation Accuracy:** The specification should describe the backend structure and conventions actually used by the project rather than prescribing Clean Architecture, a specific dependency-injection pattern, or domain-layer abstractions.

### 4.2 Frontend (TypeScript / React) Conventions
- **Strict TypeScript:** `noImplicitAny: true`, `strictNullChecks: true`.
- **Functional Components with Explicit Types:** Props defined as `interface Props`.
- **State Management:** URL search params for filters, React Query / SWR for server state, Zustand for real-time interview state (avatar visemes, mic status).

#### Good Frontend Code Example (3D Lip-Sync Hook)
```typescript
// frontend/src/hooks/useAvatarLipSync.ts
import { useEffect, useRef } from 'react';
import * as THREE from 'three';

export interface VisemeFrame {
  timeOffsetMs: number;
  viseme: string; // e.g., 'viseme_aa', 'viseme_O', 'viseme_sil'
  weight: number; // 0.0 to 1.0
}

export function useAvatarLipSync(skinnedMesh: THREE.SkinnedMesh | null, visemes: VisemeFrame[], isSpeaking: boolean) {
  const animationFrameRef = useRef<number | null>(null);
  const startTimeRef = useRef<number>(0);

  useEffect(() => {
    if (!skinnedMesh || !isSpeaking || visemes.length === 0) {
      if (skinnedMesh && skinnedMesh.morphTargetInfluences) {
        skinnedMesh.morphTargetInfluences.fill(0);
      }
      return;
    }

    startTimeRef.current = performance.now();

    const animate = () => {
      const elapsed = performance.now() - startTimeRef.current;
      const currentViseme = visemes.find(v => Math.abs(v.timeOffsetMs - elapsed) < 50);

      if (currentViseme && skinnedMesh.morphTargetDictionary && skinnedMesh.morphTargetInfluences) {
        const targetIndex = skinnedMesh.morphTargetDictionary[currentViseme.viseme];
        if (targetIndex !== undefined) {
          // Smooth blend-shape interpolation (LERP)
          skinnedMesh.morphTargetInfluences[targetIndex] = THREE.MathUtils.lerp(
            skinnedMesh.morphTargetInfluences[targetIndex],
            currentViseme.weight,
            0.3
          );
        }
      }

      animationFrameRef.current = requestAnimationFrame(animate);
    };

    animationFrameRef.current = requestAnimationFrame(animate);

    return () => {
      if (animationFrameRef.current) cancelAnimationFrame(animationFrameRef.current);
    };
  }, [skinnedMesh, visemes, isSpeaking]);
}
```

---

## 5. Testing Strategy

```
               ┌───────────────────────┐
               │    E2E Tests (10%)    │  Playwright (Complete Interview Flow)
               ├───────────────────────┤
               │ Integration Tests(30%)│  API Routes, DB Queries, Redis, WebSocket Hub
               ├───────────────────────┤
               │   Unit Tests (60%)    │  Business Logic, LLM Prompt Parsing, Scoring
               └───────────────────────┘
```

| Level | Scope | Framework / Tool | Success Target |
|---|---|---|---|
| **Unit (Backend)** | Usecases, token parsers, scoring calculations, DTO validations | `testing`, `testify/assert`, `testify/mock` | > 80% line coverage |
| **Unit (Frontend)** | Hooks, UI components, state machines, audio buffer decoders | Vitest, React Testing Library | > 75% branch coverage |
| **Integration** | DB migrations, repository queries, Redis caching, auth workflows | `testcontainers-go` (real PostgreSQL & Redis) | 100% pass on all API contracts |
| **Real-time / WS** | WebSocket lifecycle (Connect, Offer Question, Stream Audio, Visemes, Terminate) | Custom Go WS test harness | Robust reconnection handling |
| **E2E** | User Register -> Upload JD -> Configure Room -> Voice Exchange -> Report View | Playwright | Full happy paths green |

---

## 6. Engineering Boundaries

### Always Do
- **Always** validate all incoming HTTP payloads and WebSocket frames with strict schema validation.
- **Always** enforce role-based access control (Admin vs Candidate) and ensure candidates can only access their own sessions, transcripts, and reports.
- **Always** sanitize and validate raw JD inputs before passing them into LLM prompt templates (guard against prompt injection).
- **Always** handle audio stream disconnects gracefully by saving state and allowing session resumption within 5 minutes.
- **Always** run unit tests and linter before submitting commits or pull requests.

### Ask First
- **Ask First** before making breaking changes to PostgreSQL database schemas or altering existing migration scripts.
- **Ask First** before adding external third-party dependencies or introducing new cloud SaaS vendor APIs.
- **Ask First** before modifying scoring criteria weights or interview prompt templates.

### Never Do
- **Never** commit API keys, LLM tokens, or database credentials to Git (use `.env` and secret managers).
- **Never** perform direct external AI API calls inside database transactions.
- **Never** store unhashed passwords in the database (always use Bcrypt with work factor >= 12).
- **Never** skip failing unit/integration tests or comment out test assertions to pass CI.

---

## 7. Success Criteria & Quality Gates

### 7.1 Functional Success Criteria
- [ ] **JD Analysis:** System accepts raw text or PDF/DOCX JDs, parses technical skills, libraries, databases, and domain topics in < 5 seconds with an editable review interface.
- [ ] **Interview Blueprinting:** System generates a structured blueprint (4 to 8 questions categorized into Core Concepts, System Design / Architecture, Practical Problem Solving, and Scenario Questions) matching candidate level (Junior, Middle, Senior).
- [ ] **3D Voice Interview:** Real-time bi-directional voice interview with WebGL 3D avatar rendering 60 FPS on standard desktop GPUs, featuring audio synchronized with mouth blend-shapes (visemes).
- [ ] **Adaptive Follow-up Questioning:** AI asks dynamic follow-up questions when a candidate gives an incomplete, vague, or particularly insightful answer.
- [ ] **Comprehensive Evaluation Report:** Generates an evaluation report within 15 seconds after interview completion, featuring overall score (0-100), domain breakdown radar chart, 5 core competency ratings, question transcripts, and targeted learning recommendations.
- [ ] **Session History:** Candidates can view complete records of past interview attempts and performance reports.
- [ ] **Admin Dashboard:** Admin can lock/unlock accounts, monitor all live and past sessions, update domain catalogs, and view platform aggregate analytics.

### 7.2 Non-Functional Metrics (SLAs & Performance)
- **API Response Time:** Common CRUD requests (login, profile, list sessions, view reports) resolve in **< 500ms** (95th percentile < 1.5s).
- **STT/TTS Turnaround Latency:** Time from candidate speech stop to interviewer voice start < **1.8s**.
- **Concurrent Load Capacity:** Minimum **20 concurrent active interview voice sessions** running simultaneously without packet drops or backend crashes.
- **Security & Integrity:** Zero SQL injection, zero XSS, zero unauthorized cross-tenant data access; OWASP Top 10 compliance.
- **Browser Compatibility:** Chromium (Chrome, Edge, Brave), Firefox, and Safari on desktop and modern mobile devices.

---

## 8. Detailed System Architecture & Data Flow

### 8.1 High-Level Architecture Diagram

```mermaid
flowchart TB
    subgraph Client["Frontend Client (Next.js 15 + Three.js)"]
        UI["Candidate / Admin Web UI"]
        AvatarEngine["3D Avatar Engine (Three.js + LipSync)"]
        AudioEngine["Web Audio / Mic Streamer"]
    end

    subgraph Backend["Golang Backend Application"]
        HTTPHandler["REST API Handlers (Auth, JD, Interview, Admin)"]
        WSHandler["WebSocket Hub (Session Orchestrator)"]
        EvalEngine["AI Evaluation & Scoring Pipeline"]
    end

    subgraph ExternalServices["External AI & Third-Party APIs"]
        LLM["LLM (OpenAI / Gemini / Claude)"]
        STT["STT Service (Whisper / Deepgram)"]
        TTS["TTS + Viseme Engine (Azure / ElevenLabs)"]
    end

    subgraph Storage["Data Tier"]
        Postgres[(PostgreSQL 16 Database)]
        Redis[(Redis 7 Cache & Session Store)]
    end

    UI -->|HTTPS / REST| HTTPHandler
    AvatarEngine <-->|WebSocket Stream| WSHandler
    AudioEngine <-->|Audio Chunks| WSHandler

    HTTPHandler --> Postgres
    HTTPHandler --> Redis
    HTTPHandler --> LLM

    WSHandler --> STT
    WSHandler --> LLM
    WSHandler --> TTS
    WSHandler --> Redis
    WSHandler --> EvalEngine

    EvalEngine --> LLM
    EvalEngine --> Postgres
```

### 8.2 Real-Time Interview Session Sequence Flow

```mermaid
sequenceDiagram
    autonumber
    actor Candidate as Candidate (Browser)
    participant WS as Backend WebSocket Hub
    participant STT as Speech-to-Text Engine
    participant LLM as LLM Orchestrator
    participant TTS as TTS + Viseme Generator
    participant DB as PostgreSQL / Redis

    Candidate->>WS: Connect WebSocket (Session Token)
    WS->>DB: Validate Session & Active Status
    WS-->>Candidate: Connection Ready (Avatar Config, First Question Audio & Visemes)
    Candidate->>Candidate: 3D Avatar Speaks Question with Lip-Sync

    loop For each interview question & follow-up
        Candidate->>WS: Stream User Audio Chunks (Candidate Speaking)
        Candidate->>WS: User Audio Stream Finished (Silence detected)
        WS->>STT: Transcribe Audio Stream
        STT-->>WS: Transcript: "In Go, channels are used for..."
        WS->>LLM: Evaluate Answer & Determine Next Action (Follow-up or Next Question)
        LLM-->>WS: Response Text: "That is correct. How do you handle buffered vs unbuffered channels?"
        WS->>TTS: Synthesize Speech + Viseme Phoneme Timestamps
        TTS-->>WS: Audio Buffer + Viseme Array [{time: 0, viseme: 'aa', weight: 0.8}, ...]
        WS-->>Candidate: WS Packet: { type: "INTERVIEWER_RESPONSE", audio: base64, visemes: [...] }
        Candidate->>Candidate: Avatar Plays Audio & Syncs Morph Targets
    end

    Candidate->>WS: Conclude Interview / End of Questions
    WS->>LLM: Run Full Evaluation Pipeline against JD Blueprint
    LLM-->>WS: JSON Evaluation Matrix (Scores, Breakdown, Feedback)
    WS->>DB: Save Session State (Completed) + Full Evaluation Report
    WS-->>Candidate: WS Packet: { type: "INTERVIEW_COMPLETED", reportId: "..." }
```

---

## 9. Database Schema Design (PostgreSQL)

```sql
CREATE TYPE seniority_level AS ENUM ('JUNIOR', 'MIDDLE', 'SENIOR', 'LEAD');
CREATE TYPE status AS ENUM ('CREATED', 'IN_PROGRESS', 'COMPLETED', 'FAILED', 'INTERRUPTED');

-- 1. Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    phone VARCHAR(30),
    role VARCHAR(50) NOT NULL DEFAULT 'CANDIDATE', -- 'ADMIN', 'CANDIDATE'
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. Technical Domains & Skills Taxonomy
CREATE TABLE technical_domains (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL, -- e.g. 'Backend Development', 'DevOps', 'Frontend'
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE skills (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    domain_id UUID REFERENCES technical_domains(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL, -- e.g. 'Golang', 'PostgreSQL', 'Docker'
    category VARCHAR(50) NOT NULL, -- 'Language', 'Database', 'Framework', 'Tool'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. Job Descriptions (JD)
CREATE TABLE job_descriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    candidate_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    raw_content TEXT NOT NULL,
    extracted_data JSONB NOT NULL, -- { domains: [], skills: [], level: "Senior", requirements: [] }
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 4. 3D Avatars & Voice Personas
CREATE TABLE avatar_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    model_url VARCHAR(500) NOT NULL, -- GLTF/GLB asset path
    preview_image_url VARCHAR(500),
    gender VARCHAR(20) NOT NULL,
    voice_id VARCHAR(100) NOT NULL, -- Azure / ElevenLabs Voice identifier
    speaking_speed NUMERIC(3, 2) DEFAULT 1.0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. Interview Sessions
CREATE TABLE interview_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    candidate_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    jd_id UUID NOT NULL REFERENCES job_descriptions(id) ON DELETE RESTRICT,
    avatar_id UUID NOT NULL REFERENCES avatar_profiles(id) ON DELETE RESTRICT,
    difficulty seniority_level NOT NULL,
    target_duration_minutes INT NOT NULL DEFAULT 30,
    status status NOT NULL DEFAULT 'CREATED',
    total_questions INT NOT NULL DEFAULT 5,
    blueprint JSONB NOT NULL, -- Session-level snapshot of the interview blueprint used for this interview
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. Interview Question & Turn Transcripts
CREATE TABLE session_turns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL REFERENCES interview_sessions(id) ON DELETE CASCADE,
    turn_index INT NOT NULL,
    question_topic VARCHAR(100) NOT NULL,
    interviewer_question TEXT NOT NULL,
    candidate_audio_url VARCHAR(500),
    candidate_transcript TEXT,
    is_follow_up BOOLEAN DEFAULT FALSE,
    turn_duration_seconds INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 7. AI Evaluations & Performance Reports
CREATE TABLE performance_reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID UNIQUE NOT NULL REFERENCES interview_sessions(id) ON DELETE CASCADE,
    candidate_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    overall_score NUMERIC(5, 2) NOT NULL, -- 0.00 to 100.00
    technical_accuracy_score NUMERIC(5, 2) NOT NULL,
    depth_score NUMERIC(5, 2) NOT NULL,
    problem_solving_score NUMERIC(5, 2) NOT NULL,
    answer_relevance_score NUMERIC(5, 2) NOT NULL,
    communication_clarity_score NUMERIC(5, 2) NOT NULL,
    domain_scores JSONB NOT NULL, -- { "Golang": 85, "Database Design": 70, "Concurrency": 90 }
    strengths JSONB NOT NULL, -- ["Solid understanding of Go channels", "Clean architectural thinking"]
    weaknesses JSONB NOT NULL, -- ["Missed database isolation levels nuances"]
    actionable_recommendations JSONB NOT NULL,
    question_evaluations JSONB NOT NULL, -- Detailed feedback per question
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## 10. API Endpoints & WebSocket Contracts

### 10.1 REST API Routes

#### Authentication (`/api/v1/auth`)
- `POST /register`: Candidate registration (`email`, `password`, `fullName`).
- `POST /login`: Authentication; returns JWT Access Token (15 min) and Refresh Token (7 days).
- `POST /refresh`: Refresh expired Access Token.
- `GET /me`: Returns logged-in user profile and role.

#### Job Description Management (`/api/v1/jd`)
- `POST /upload`: Upload PDF/Docx or raw text. Returns extracted skills & blueprint preview.
- `GET /`: List all parsed JDs for the candidate.
- `GET /:id`: Retrieve specific JD details and extracted entities.
- `PUT /:id`: Update/override extracted skills and targets before interview creation.

#### Interview Management (`/api/v1/interviews`)
- `POST /create`: Create new interview session (`jdId`, `avatarId`, `difficulty`, `targetDuration`).
- `GET /`: List candidate's interview history (with statuses and scores).
- `GET /:id`: Get session metadata and status.
- `GET /:id/report`: Get completed AI evaluation report and domain radar metrics.

#### Admin Management (`/api/v1/admin`)
- `GET /users`: Paginated list of users with filter and lock status.
- `PUT /users/:id/lock`: Lock / Unlock candidate account.
- `GET /sessions`: All platform sessions with filters (status, candidate, date range).
- `GET /analytics/overview`: Aggregate stats (Total interviews, completion rate, avg score, common skill weaknesses).
- `CRUD /domains`, `CRUD /avatars`: Manage domains, skills, and 3D avatars.

### 10.2 Real-Time WebSocket Protocol (`/api/v1/ws/interview/:session_id`)

#### Client to Server Packets
```json
// 1. Join & Authenticate
{
  "type": "JOIN_ROOM",
  "token": "JWT_ACCESS_TOKEN"
}

// 2. Stream Audio Chunk (Base64 WebM / PCM)
{
  "type": "AUDIO_CHUNK",
  "data": "GkXfo59ChoEBQveBAULygQ8U...",
  "isFinal": false
}

// 3. User Stopped Speaking
{
  "type": "SPEECH_FINISHED"
}

// 4. Request Early Session Termination
{
  "type": "END_SESSION"
}
```

#### Server to Client Packets
```json
// 1. Session Initialized
{
  "type": "SESSION_READY",
  "avatarModelUrl": "/models/interviewer_sarah.glb",
  "totalQuestions": 6,
  "currentQuestionIndex": 1
}

// 2. Interviewer Speaking (Audio + 3D Blendshape Visemes)
{
  "type": "INTERVIEWER_SPEAK",
  "questionText": "Can you explain how goroutine scheduling works in the Go runtime?",
  "audioBase64": "UklGRi4AAABXQVZFZm10IBAAAA...",
  "visemes": [
    { "timeOffsetMs": 0, "viseme": "viseme_sil", "weight": 1.0 },
    { "timeOffsetMs": 120, "viseme": "viseme_kk", "weight": 0.9 },
    { "timeOffsetMs": 240, "viseme": "viseme_aa", "weight": 0.85 }
  ],
  "isFollowUp": false
}

// 3. Transcription Real-time Feedback
{
  "type": "CANDIDATE_TRANSCRIPT_UPDATE",
  "transcript": "Goroutines are managed by the Go runtime scheduler using an M:N model..."
}

// 4. Interview Concluded
{
  "type": "INTERVIEW_FINISHED",
  "reportId": "8f39bc21-7294-4d89-9c57-79a0ce6295ea"
}

// 5. Error
{
  "type": "ERROR",
  "code": "MIC_SILENCE_TIMEOUT",
  "message": "We could not hear your answer. Please check your microphone."
}
```

---

## 11. AI Orchestration, Prompt Design & Evaluation

### 11.1 Job Description Analysis Prompt Schema
```json
{
  "system_prompt": "You are a Senior Technical Recruiter and Staff Software Architect. Analyze the provided Job Description (JD) and extract technical requirements into a strict JSON schema.",
  "response_format": {
    "type": "json_object",
    "schema": {
      "title": "string",
      "seniority_level": "configured seniority_level value",
      "core_domains": ["string"],
      "must_have_skills": ["string"],
      "nice_to_have_skills": ["string"],
      "interview_blueprint": [
        {
          "topic": "string",
          "depth_level": "Fundamental | In-Depth | Architectural",
          "sample_question": "string",
          "evaluation_criteria": ["string"]
        }
      ]
    }
  }
}
```

### 11.2 Adaptive Follow-Up Questioning Strategy
1. **Direct Answer:** If candidate gives an accurate and complete answer, praise briefly and move to the next planned blueprint question.
2. **Surface Answer:** If candidate mentions keywords without depth, ask an adaptive follow-up: *"You mentioned using Redis for caching; how do you handle cache invalidation and the thundering herd problem?"*
3. **Incorrect / Stuck:** If candidate struggles, provide a gentle hint or pivot to avoid dead air, noting the knowledge gap in the final report.

### 11.3 Multi-Dimensional Evaluation Matrix
Scores are normalized from 0 to 100 based on rubric criteria:
- **Technical Accuracy (30%):** Correctness of syntax, architectural trade-offs, and design patterns.
- **Depth of Understanding (25%):** Explaining *why* a solution works beyond surface syntax (memory management, runtime mechanics, performance).
- **Problem-Solving & Adaptability (20%):** Handling edge cases, follow-up constraints, and trade-offs.
- **Answer Relevance (15%):** Staying on topic and directly addressing the question asked.
- **Communication Clarity (10%):** Structured explanations, concise articulation, and professional tone.

---

## 12. Open Questions & Assumptions for Human Review

### Assumptions Made:
1. **Avatar Rendering:** 3D avatars will be rendered client-side in the browser via Three.js with standard GLTF/GLB models using Ready Player Me / RPM compatible blend-shapes (`viseme_aa`, `viseme_E`, `viseme_I`, `viseme_O`, `viseme_U`, etc.) for zero GPU server cost.
2. **Audio Streaming:** Frontend records candidate speech using standard MediaRecorder (WebM/Opus) and sends chunks over the existing WebSocket connection upon speech pauses (Voice Activity Detection), avoiding complex WebRTC SFU infrastructure.
3. **Scope & Access:** Every registered candidate can create and practice unlimited interview sessions for evaluation and training purposes.

### Review Decisions Needed from Supervisor / Team:
1. **Speech Service Choice:** Should we default to Azure Speech Services (which provides native phoneme/viseme timestamps out of the box) or OpenAI Whisper + ElevenLabs / Rhubarb Lip Sync?
3. **Code Editor / Coding Sandbox:** Is live coding (Monaco editor execution) required during the interview, or is the platform purely conceptual / architectural technical dialog?
