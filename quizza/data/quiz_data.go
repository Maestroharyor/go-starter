package data

type Question struct {
	Question string
	Answer   string
	Points   uint
}

var Questions = []Question{
	{
		Question: "What is 15 * 12?",
		Answer:   "180",
		Points:   10,
	},
	{
		Question: "Who is the first president of the United States?",
		Answer:   "George Washington",
		Points:   10,
	},
	{
		Question: "What is the capital of France?",
		Answer:   "Paris",
		Points:   10,
	},
	{
		Question: "What is the largest planet in our solar system?",
		Answer:   "Jupiter",
		Points:   10,
	},
	{
		Question: "What is the chemical symbol for gold?",
		Answer:   "Au",
		Points:   10,
	},
	{
		Question: "How many continents are there on Earth?",
		Answer:   "7",
		Points:   10,
	},
	{
		Question: "Who wrote the play 'Romeo and Juliet'?",
		Answer:   "William Shakespeare",
		Points:   10,
	},
	{
		Question: "What is the square root of 144?",
		Answer:   "12",
		Points:   10,
	},
	{
		Question: "What gas do plants absorb from the atmosphere?",
		Answer:   "Carbon Dioxide",
		Points:   10,
	},
	{
		Question: "In what year did World War II end?",
		Answer:   "1945",
		Points:   10,
	},
}
