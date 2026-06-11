package config

type Config struct {
	// --- Основні шляхи ---
	Input        string
	Output       string
	AbsInput     string
	AbsOutput    string
	Exepath      string
	Exedir       string
	Datadir      string
	FinalExename string

	// --- Тимчасові директорії ---
	TempDir string // Базова папка temp
	Temp    string // Унікальна папка для поточного білда

	// --- Залежності ---
	PathGo      string
	PathPython  string
	Interpreter string

	// --- Цільові шляхи ---
	FinalExePath string
	TargetPyDir  string

	// --- Прапорці та налаштування ---
	CLI  bool
	Log  string
	Test bool
	Err  error
}

func New() *Config {
	return &Config{}
}
