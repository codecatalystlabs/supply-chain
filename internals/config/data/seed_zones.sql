-- Seed zones: one row per region with the same zone name and code (Zone 1 / Z1),
-- and each region's id and code. Run after regions are populated.
-- Skips regions that already have a zone.
INSERT INTO zones (region_id, name, code, created_at, region_code)
SELECT r.id, 'Zone 1', 'Z1', NOW(), r.code
FROM regions r
WHERE NOT EXISTS (SELECT 1 FROM zones z WHERE z.region_id = r.id);
