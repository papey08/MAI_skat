from api.http_server import HttpServer

from common.get_config import get_config

if __name__ == "__main__":
    config = get_config()
    http_server = HttpServer(
        auth_nats_url=config['common']['nats_url'], 
        user_nats_url=config['common']['nats_url'],
        core_nats_url=config['common']['nats_url'],
        password_salt=config['common']['password_salt'],
        jwt_secret=config['common']['jwt_secret'])
    
    http_server.run(
        host=config['api-service']['host'],
        port=config['api-service']['port'])
