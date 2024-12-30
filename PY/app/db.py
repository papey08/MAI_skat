import time

from entities.locations import locations
from entities.place import Place

import psycopg2


class Db:
    def __init__(self, db_host: str, db_port: int, db_username: str, db_password: str, db_database_name: str):
        time.sleep(30) # ждем чтобы бд успела подняться в докере
        self.connection = psycopg2.connect(
            host=db_host,
            port=db_port,
            user=db_username,
            password=db_password,
            dbname=db_database_name
        )
        self.connection.autocommit = True

    def add_user(self, username: str):
        with self.connection.cursor() as cursor:
            cursor.execute(UserQuery.ADD_USER, (username,))

    def add_response(self, username: str, place_id: int):
        with self.connection.cursor() as cursor:
            cursor.execute(UserQuery.ADD_RESPONSE, (place_id, username))

    def execute_find_place(self, query: str, vars: tuple):
        with self.connection.cursor() as cursor:
            cursor.execute(query, vars)
            row = cursor.fetchone()
            if row:
                return Place(
                    id    =row[0],
                    name=row[1],
                    url=row[2],
                    location=row[3]
                )

    def find_place(self, category: str, location: str):
        place = self.execute_find_place(UserQuery.FIND_PLACE_BY_CATEGORY_AND_LOCATION, (category, location));
        if place:
            return place
        neighbors = locations.get(location, [])
        for neighbor in neighbors:
            place = self.execute_find_place(UserQuery.FIND_PLACE_BY_CATEGORY_AND_LOCATION, (neighbor, location));
            if place:
                return place
        place = self.execute_find_place(UserQuery.FIND_PLACE_BY_CATEGORY, (category,));
        if place:
            return place
        return Place()

class UserQuery:
    ADD_USER="""
    INSERT INTO tg_user (tg_username)
    VALUES (%s)
    ON CONFLICT (tg_username) DO NOTHING;
    """
    ADD_RESPONSE="""
    INSERT INTO response (user_id, place_id)
    SELECT id, %s FROM tg_user WHERE tg_username = %s;
    """
    FIND_PLACE_BY_CATEGORY_AND_LOCATION = """
    SELECT p.*
    FROM place p
    JOIN place_category pc ON p.id = pc.place_id
    WHERE pc.category = %s AND p.location = %s
    LIMIT 1;
    """
    FIND_PLACE_BY_CATEGORY = """
    SELECT p.*
    FROM place p
    JOIN place_category pc ON p.id = pc.place_id
    WHERE pc.category = %s
    LIMIT 1;
    """
