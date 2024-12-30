class Response:
    def __init__(self, id=-1,original_text="Nothing was found", resp_category="", current_index=0):
        self.id = id
        self.original_text = original_text
        self.resp_category = resp_category
        self.current_index = current_index

    def __repr__(self):
        return (f"<CategoryResponse id={self.id}, "
                f"original_text='{self.original_text}', "
                f"resp_category='{self.resp_category}', "
                f"current_index={self.current_index}>")