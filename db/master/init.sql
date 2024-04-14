CREATE TABLE dialogs (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id_1 UUID NOT NULL,
    user_id_2 UUID NOT NULL
);

CREATE TABLE messages (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    dialog_id UUID NOT NULL,
    from_id UUID NOT NULL,
    to_id UUID NOT NULL,
    text_ TEXT,
    created_at DATE NOT NULL DEFAULT NOW(),
    updated_at DATE NOT NULL DEFAULT NOW(),

   FOREIGN KEY (dialog_id) REFERENCES dialogs(id)
);

SELECT create_distributed_table('dialogs', 'id');
SELECT create_distributed_table('messages', 'dialog_id', colocate_with => 'dialogs');