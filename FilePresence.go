package main

type FilePresence struct {
	IsFilePresent   bool
	FilePathIfExist string
}

func FileAbsent() FilePresence {
	return FilePresence{IsFilePresent: false, FilePathIfExist: ""}
}

func FilePresent(filePath string) FilePresence {
	return FilePresence{IsFilePresent: true, FilePathIfExist: filePath}
}
