package ImageFileMicroservices

func RunService() {
	workToDo := make(chan string, 1024)
	finishedWorkMap := make(map[string]bool)
	go startProcessor(workToDo, &finishedWorkMap)
	setupWebInterface(workToDo, &finishedWorkMap)
}
