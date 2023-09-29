import logging
import datetime
import time
import math

n = 3090869112548711415389914349925751666928911216642414835736649468121

logging.basicConfig(filename='log.txt', level=logging.INFO, format='%(asctime)s %(message)s')

for p in range(2, n):
    if n % p == 0:
        q = n / p
        logging.info('ANSWER: p = {}   q = {}'.format(p, q))
        break

    if datetime.datetime.now().minute == 0 and datetime.datetime.now().second == 0:
        progress = p // math.sqrt(n)
        logging.info('Last calculated p = {}   progress = {}'.format(p, progress))
        time.sleep(1)
