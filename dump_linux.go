package dump

func Dump() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "golang"
	}
	fileName := appName + "_" + time.Now().Format("2006-01-02")
	logFilename := "/home/logs/" + fileName + ".log"
	os.Mkdir("/home/logs", os.ModePerm)
	logFile, _ := os.OpenFile(logFilename, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	redirectStderr(logFile)
	os.Stderr.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + "\n"))
}

func redirectStderr(f *os.File) {
	os.Stderr = f
	err := syscall.Dup2(int(f.Fd()), 2)
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
