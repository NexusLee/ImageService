package ImageFileMicroservices

func RunService() {
	workToDo := make(chan string)
	finishedWorkMap := make(map[string]bool)
	go startProcessor(workToDo, &finishedWorkMap)
	setupWebInterface(workToDo, &finishedWorkMap)
}
