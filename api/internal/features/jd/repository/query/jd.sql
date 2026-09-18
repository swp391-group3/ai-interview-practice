-- name: CreateCustomizedJD :one
INSERT INTO job_descriptions (user_id, title, seniority_level, raw_text, parsed_data, status)
VALUES (@user_id, @title, @seniority_level, @raw_text, @parsed_data, 'customized')
RETURNING *;

-- name: ListJDs :many
SELECT id, title, seniority_level, status, created_at, updated_at FROM job_descriptions
WHERE user_id = @user_id
ORDER BY created_at DESC, id DESC
LIMIT @page_limit OFFSET @page_offset;

-- name: CountJDs :one
SELECT count(*) FROM job_descriptions
WHERE user_id = @user_id;

-- name: GetJD :one
SELECT * FROM job_descriptions
WHERE id = @id AND user_id = @user_id;

-- name: UpdateJD :one
UPDATE job_descriptions
SET title = @title, seniority_level = @seniority_level,
    parsed_data = @parsed_data, updated_at = now()
WHERE id = @id AND user_id = @user_id
RETURNING *;

-- name: DeleteJD :one
DELETE FROM job_descriptions
WHERE id = @id AND user_id = @user_id
RETURNING id;
