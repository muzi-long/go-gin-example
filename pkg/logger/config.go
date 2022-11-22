package logger

type Config struct {
	Path       string // 日志文件路径，默认 os.TempDir()
	FileName   string // 日志文件名称，默认 os.TempDir()
	MaxSize    int    // 每个日志文件保存10M，默认 100M
	MaxBackups int    // 保留30个备份，默认不限
	MaxAge     int    // 保留7天，默认不限
	Compress   bool   // 是否压缩，默认不压缩
	Level      string // 日志级别 debug->info->warn->error
}
