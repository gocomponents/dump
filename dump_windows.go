package dump

import (
	"log"
	"os"
	"syscall"
	"time"
)

func Dump()  {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "golang"
	}
	fileName := appName + "_" +  time.Now().Format("2006-01-02")
	logFilename := "c:\\logs\\" + fileName + ".log"
	logFile, _ := os.OpenFile(logFilename, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	redirectStderr(logFile)
	os.Stderr.Write([]byte("\r\n"+time.Now().Format("2006-01-02 15:04:05") + "\r\n"))
}

func redirectStderr(f *os.File) {
	err := setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
	os.Stderr = f
}

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func setStdHandle(stdHandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdHandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}
