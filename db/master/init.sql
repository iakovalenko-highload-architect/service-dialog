CREATE TABLE messages (
   id SERIAL NOT NULL PRIMARY KEY,
   from_id UUID NOT NULL,
   to_id UUID NOT NULL,
   text_ TEXT,
   created_at DATE NOT NULL DEFAULT NOW(),
   updated_at DATE NOT NULL DEFAULT NOW()
);
