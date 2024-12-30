SELECT '{
  "sports": "хоккей"
}'::jsonb || '{
  "trips": 5
}'::jsonb;

SELECT '{
  "sports": "хоккей"
}'::jsonb || '{
  "trips": 5
}'::jsonb || '{
  "pilots": "ИВАН"
}'::jsonb;
