INSERT INTO public.tasklist_task (id, tenantid, key, partitionid)
VALUES 
  ('4', 'NewTenantIdFor4', 2, 5),
  ('7', 'NewTenantIdFor7', 3, 6),
  ('10', 'NewTenantIdFor10', 4, 7)
ON CONFLICT (id) DO UPDATE
SET
  tenantid = EXCLUDED.tenantid,
  key = EXCLUDED.key,
  partitionid = EXCLUDED.partitionid;