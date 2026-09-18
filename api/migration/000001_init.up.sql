CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE account_role AS ENUM (
    'participant',
    'jury',
    'admin'
);

CREATE TYPE seniority_level AS ENUM (
    'intern',
    'junior',
    'mid',
    'senior',
    'lead'
);

CREATE TYPE jd_status AS ENUM (
    'uploaded',
    'parsing',
    'parsed',
    'customized',
    'locked',
    'failed'
);

CREATE TYPE interview_status AS ENUM (
    'created',
    'in_progress',
    'completed',
    'interrupted',
    'terminated_early',
    'failed'
);

CREATE TYPE interview_difficulty AS ENUM (
    'easy',
    'medium',
    'hard'
);

CREATE TABLE IF NOT EXISTS accounts (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

    email text UNIQUE NOT NULL,
    full_name text NOT NULL,
    password_hash varchar(128) NOT NULL,

    role account_role NOT NULL DEFAULT 'participant'::account_role,
    is_locked bool NOT NULL DEFAULT false,

    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS technical_domains (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text UNIQUE NOT NULL,
    description text,
    is_active bool NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS skills (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    domain_id uuid NOT NULL REFERENCES technical_domains(id) ON DELETE CASCADE,
    name text NOT NULL,
    category varchar(128),
    is_active bool NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (domain_id, name)
);

CREATE TABLE IF NOT EXISTS job_descriptions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    title text NOT NULL,
    seniority_level seniority_level NOT NULL,
    raw_text text NOT NULL,
    parsed_data jsonb,
    blueprint jsonb,
    status jd_status NOT NULL DEFAULT 'uploaded'::jd_status,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS job_description_skills (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_description_id uuid NOT NULL REFERENCES job_descriptions(id) ON DELETE CASCADE,
    skill_id uuid NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    is_required boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (job_description_id, skill_id)
);

CREATE TABLE IF NOT EXISTS avatar_profiles (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    model_url text NOT NULL,
    voice_id text NOT NULL,
    default_camera jsonb,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS interview_sessions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    jd_id uuid REFERENCES job_descriptions(id) ON DELETE SET NULL,
    avatar_id uuid REFERENCES avatar_profiles(id) ON DELETE SET NULL,
    status interview_status NOT NULL DEFAULT 'created'::interview_status,
    difficulty interview_difficulty NOT NULL,
    total_questions integer NOT NULL DEFAULT 0,
    duration_minutes integer NOT NULL DEFAULT 0,
    current_question_index integer NOT NULL DEFAULT 0,
    blueprint_snapshot jsonb NOT NULL,
    candidate_audio_url text,
    started_at timestamptz,
    ended_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS session_turns (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id uuid NOT NULL REFERENCES interview_sessions(id) ON DELETE CASCADE,
    turn_index integer NOT NULL,
    topic text,
    question_text text NOT NULL,
    candidate_transcript text,
    is_follow_up boolean NOT NULL DEFAULT false,
    started_at timestamptz,
    completed_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (session_id, turn_index)
);

CREATE TABLE IF NOT EXISTS performance_reports (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id uuid UNIQUE NOT NULL REFERENCES interview_sessions(id) ON DELETE CASCADE,
    overall_score decimal(5, 2) NOT NULL,
    competency_scores jsonb NOT NULL,
    domain_scores jsonb NOT NULL,
    strengths jsonb NOT NULL,
    weaknesses jsonb NOT NULL,
    question_feedback jsonb NOT NULL,
    roadmap jsonb NOT NULL,
    is_partial boolean NOT NULL DEFAULT false,
    evaluation_notice text,
    generated_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_skills_domain_id ON skills(domain_id);
CREATE INDEX IF NOT EXISTS idx_job_descriptions_user_id ON job_descriptions(user_id);
CREATE INDEX IF NOT EXISTS idx_job_description_skills_jd_id ON job_description_skills(job_description_id);
CREATE INDEX IF NOT EXISTS idx_job_description_skills_skill_id ON job_description_skills(skill_id);
CREATE INDEX IF NOT EXISTS idx_interview_sessions_user_id ON interview_sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_interview_sessions_jd_id ON interview_sessions(jd_id);
CREATE INDEX IF NOT EXISTS idx_interview_sessions_avatar_id ON interview_sessions(avatar_id);
CREATE INDEX IF NOT EXISTS idx_session_turns_session_id ON session_turns(session_id);
