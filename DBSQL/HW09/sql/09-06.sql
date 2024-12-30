explain select city, row_number() over (partition by city order by city) as city_rank from airports;
