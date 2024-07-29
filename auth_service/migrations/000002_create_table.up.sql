-- -- Insert mock data into users table
-- INSERT INTO users (id, username, password, email) VALUES
-- ('00000000-0000-0000-0000-000000000001', 'johndoe', 'securepassword', 'johndoe@example.com'),
-- ('00000000-0000-0000-0000-000000000002', 'janedoe', 'anotherpassword', 'janedoe@example.com'),
-- ('11111111-1111-1111-1111-111111111111', 'user1', 'password1', 'user1@example.com'),
-- ('22222222-2222-2222-2222-222222222222', 'user2', 'password2', 'user2@example.com');

-- Ensure the pgcrypto extension is available
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Insert mock data into users table with bcrypt encrypted passwords
INSERT INTO users (id, username, password, email) VALUES
('00000000-0000-0000-0000-000000000001', 'johndoe', crypt('securepassword', gen_salt('bf')), 'johndoe@example.com'),
('00000000-0000-0000-0000-000000000002', 'janedoe', crypt('anotherpassword', gen_salt('bf')), 'janedoe@example.com'),
('11111111-1111-1111-1111-111111111111', 'user1', crypt('password1', gen_salt('bf')), 'user1@example.com'),
('22222222-2222-2222-2222-222222222222', 'user2', crypt('password2', gen_salt('bf')), 'user2@example.com');
