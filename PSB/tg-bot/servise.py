from repository import Repository


class Servise:
    def __init__(self, db: Repository):
        self.db = db

    def get_next_response_in_category(self, category: str, current_id: int):
        return self.db.get_next_response_in_category(category, current_id + 1)

    def get_previous_response_in_category(self, category: str, current_id: int):
        return self.db.get_next_response_in_category(category, current_id - 1)

    def get_current_response_in_category(self, category: str, current_id: int):
        return self.db.get_next_response_in_category(category, current_id)
