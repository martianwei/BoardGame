
-- build a mock data for users table
INSERT INTO users (id, name, password,score, created_at, updated_at)
VALUES
  ('f3a31764-bb25-4c12-89d0-3d7d112b51c1', 'martianwei', 'password',0, NOW(), NOW()),
  ('3f6e7e75-d153-4f0e-b726-62f2eac4e157', 'yoya', 'password',0, NOW(), NOW()),
  ('adefdb45-2ae1-44d1-9120-50b0c0a8db8c', 'pig', 'password',0, NOW(), NOW()),
  ('6e8760f3-9ff0-4f2e-bc94-1f1f0e7b3051', 'dog', 'password',0, NOW(), NOW());

-- build a mock data for formations table
INSERT INTO formations (id, user_id, name, created_at, updated_at)
VALUES
  ('d84a5259-8b18-4dab-8c3f-02c5f08ffbe5', 'f3a31764-bb25-4c12-89d0-3d7d112b51c1', 'Formation 1', NOW(), NOW()),
  ('7c95b73e-868a-4d45-8949-7959447f7c78', '3f6e7e75-d153-4f0e-b726-62f2eac4e157', 'Formation 2', NOW(), NOW()),
  ('d1a708f6-8d71-480e-80a5-924a7f9c855d', 'adefdb45-2ae1-44d1-9120-50b0c0a8db8c', 'Formation 3', NOW(), NOW()),
  ('f2e8cdaa-8d0f-4c6e-9065-6d1ab6efabe8', '6e8760f3-9ff0-4f2e-bc94-1f1f0e7b3051', 'Formation 4', NOW(), NOW());

-- build a mock data for formation_elements table
INSERT INTO formation_elements (formation_id, commission, position)
VALUES
  ('d84a5259-8b18-4dab-8c3f-02c5f08ffbe5', 1, 1),
  ('d84a5259-8b18-4dab-8c3f-02c5f08ffbe5', 2, 2),
  ('d84a5259-8b18-4dab-8c3f-02c5f08ffbe5', 3, 3),
  ('7c95b73e-868a-4d45-8949-7959447f7c78', 4, 1),
  ('7c95b73e-868a-4d45-8949-7959447f7c78', 5, 2),
  ('7c95b73e-868a-4d45-8949-7959447f7c78', 6, 3),
  ('d1a708f6-8d71-480e-80a5-924a7f9c855d', 7, 1),
  ('d1a708f6-8d71-480e-80a5-924a7f9c855d', 8, 2),
  ('d1a708f6-8d71-480e-80a5-924a7f9c855d', 9, 3),
  ('f2e8cdaa-8d0f-4c6e-9065-6d1ab6efabe8', 10, 1),
  ('f2e8cdaa-8d0f-4c6e-9065-6d1ab6efabe8', 11, 2),
  ('f2e8cdaa-8d0f-4c6e-9065-6d1ab6efabe8', 12, 3);

-- build a mock data for game table
INSERT INTO games (id,created_at)
VALUES
  ('9c253f05-f8d2-47be-bd30-17f02515c682', NOW());
