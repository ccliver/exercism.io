class HighScores(object):
	def __init__(self, scores):
		self.scores = scores

	def latest(self):
		return self.scores[-1]

	def personal_best(self):
		self.sorted_scores = self.scores.copy()
		self.sorted_scores.sort()
		return self.sorted_scores[-1]

	def personal_top(self):
		self.sorted_scores = self.scores.copy()
		self.sorted_scores.sort()
		self.sorted_scores.reverse()
		return self.sorted_scores[:3]

	def report(self):
		if self.latest() == self.personal_best():
			return "Your latest score was " + str(self.latest()) + ". That's your personal best!"
		else:
			return "Your latest score was " + str(self.latest()) + ". That's " + str(self.personal_best() - self.latest()) + " short of your personal best!"
