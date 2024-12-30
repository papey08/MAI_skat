import yaml
import os

import app.tgbot as tgbot
import app.db as db
import app.mlmodel as mlmodel
import app.service as service


class Application:
    def __init__(self, api_key, path_to_model, db_config):
        self.db = db.Db(
            db_config['host'],
            db_config['port'],
            db_config['username'],
            db_config['password'],
            db_config['dbname']
        )
        self.mlmodel = mlmodel.MlModel(path_to_model)
        self.service = service.Service(self.mlmodel, self.db)
        self.tgbot = tgbot.Tgbot(api_key, self.service)

    def run(self):
        self.tgbot.run()

def get_config():
    with open('config.yaml', 'r') as config_file:
        config = yaml.safe_load(config_file)
        return os.environ['API_KEY'], config['path_to_model'], config['db_config']

if __name__ == '__main__':
    api_key, path_to_model, db_config = get_config()
    app = Application(api_key, path_to_model, db_config)
    print('🎉bot started🎉')
    app.run()
