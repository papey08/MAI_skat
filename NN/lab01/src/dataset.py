import csv
from typing import List, Tuple

from src.training import Pair, Pairs

def to_pair(record: List[str]) -> Tuple[Pair]:
    res = int(float(record[0])) - 1
    res_encoded = [0.0] * 3
    res_encoded[res] = 1.0
    features = [float(x) for x in record[1:]]
    return Pair(input_data=features, response=res_encoded)

def load(path: str) -> Tuple[Pairs]:
    with open(path, mode='r', newline='') as f:
        reader = csv.reader(f)
        pairs = Pairs([])
        for record in reader:
            pair = to_pair(record)
            pairs.pairs.append(pair)
        return pairs
