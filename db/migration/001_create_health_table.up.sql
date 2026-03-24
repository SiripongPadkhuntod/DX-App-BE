-- Create the health table following the requested structure (translated to PostgreSQL)
CREATE TABLE IF NOT EXISTS health (
    health_id   VARCHAR(36) DEFAULT (md5(random()::text)) PRIMARY KEY,
    service     VARCHAR(255) NOT NULL,
    status      SMALLINT,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert initial record
INSERT INTO health (health_id, service, status) VALUES ('initial-id', 'dexter-transport', 1);
