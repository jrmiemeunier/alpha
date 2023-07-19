import json
import os

class PlainDataLoader():
    def __init__(self, config_path: str=DATAS_CONFIG) -> None:
        self._path = config_path
        with open(config_path, 'r', encoding='utf-8') as config:
            data = json.load(config)
            self.top_level_path:str = data["cp_path"]
            self.datasets:dict = data["datasets"]
            self.id_table = {
                v["id"]: k for (k, v) in self.datasets.items()
            }
    
    def body_extractor(self, target: str) -> list:
        if target not in self.datasets:
            print(f"{target} is not included in datas.json as a dataset")
            return None
        configs = self.datasets[target]
        tag = configs["tag"]
        body = []  # may get a bit huge... 
        full_path = os.path.join(self.top_level_path, configs["path"])
        if os.path.isfile(full_path):  # single file json
            with open(full_path, mode='r', encoding='utf-8') as file:
                data = json.load(file)
                for poem in data:
                    body += poem[tag]
            return body
        # a dir, probably with a skip list
        subpaths = os.listdir(full_path)
        for filename in subpaths:
            if filename in configs["excludes"]:
                continue
            with open(os.path.join(full_path, filename), mode='r', encoding='utf-8') as file:
                data = json.load(file)
                for poem in data:
                    body += poem[tag]
        return body
