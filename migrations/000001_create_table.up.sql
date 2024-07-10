CREATE TABLE evaluations (
    id UUID primary key DEFAULT gen_random_uuid(),
    evaluator_id UUID,
    employee_id  UUID,
    rating FLOAT NOT NULL,
    comment TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);
