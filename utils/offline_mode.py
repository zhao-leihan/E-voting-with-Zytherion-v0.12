import json

class OfflineMode:
    def __init__(self, file_path="offline_votes.json"):
        self.file_path = file_path

    def save_vote(self, vote_data):
        # Save vote data to a local file
        with open(self.file_path, 'a') as f:
            json.dump(vote_data, f)
            f.write('\n')

    def load_votes(self):
        # Load votes from the local file
        votes = []
        with open(self.file_path, 'r') as f:
            for line in f:
                votes.append(json.loads(line))
        return votes