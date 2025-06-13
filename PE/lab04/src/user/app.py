import asyncio

from nats_server import NatsServer
from common.get_config import get_config

async def main():
    config = get_config()
    await NatsServer(
        config['common']['nats_url'],
        config['user-service']['postgres_url'],
    ).run()

if __name__ == "__main__":
    asyncio.run(main())
