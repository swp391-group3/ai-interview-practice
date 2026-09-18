ALTER TABLE job_descriptions DROP COLUMN blueprint;

CREATE TABLE interview_blueprints (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_description_id uuid NOT NULL REFERENCES job_descriptions(id) ON DELETE RESTRICT,
    difficulty interview_difficulty NOT NULL,
    duration_minutes integer NOT NULL CHECK (duration_minutes > 0),
    question_count integer NOT NULL CHECK (question_count > 0),
    blueprint_data jsonb NOT NULL,
    contract_version integer NOT NULL CHECK (contract_version > 0),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_interview_blueprints_job_description_id ON interview_blueprints(job_description_id);
