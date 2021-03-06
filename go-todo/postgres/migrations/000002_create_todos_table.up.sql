CREATE TABLE todos (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    body TEXT,
    completed BOOLEAN DEFAULT FALSE,
    user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW ()
);
