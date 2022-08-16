package utils

//golang os.OpenFile几种常用模式
//os.O_WRONLY | os.O_CREATE | O_EXCL           【如果已经存在，则失败】
//
//os.O_WRONLY | os.O_CREATE                         【如果已经存在，会覆盖写，不会清空原来的文件，而是从头直接覆盖写】
//
//os.O_WRONLY | os.O_CREATE | os.O_APPEND  【如果已经存在，则在尾部添加写】

// Write file , create if not exist
//func Write(filePath string) error {
//	os.OpenFile(filePath, O_APPEND)
//}
