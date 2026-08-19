from __future__ import annotations
import json
from .scenario import cyclone_partition_demo
if __name__ == "__main__": print(json.dumps(cyclone_partition_demo(), indent=2, sort_keys=True))
