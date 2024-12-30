SET datestyle TO 'MDY';

SELECT '18-05-2016'::timestamp;

SELECT '05-18-2016'::timestamp;

SET datestyle TO DEFAULT;

SHOW datestyle;

SET datestyle TO 'ISO, DMY';

SHOW datestyle;
