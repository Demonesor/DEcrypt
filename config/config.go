package config

// PathsConfig відповідає за всі шляхи на машині користувача
type PathsConfig struct {
	Input     string // Прямий шлях до вхідного файлу (напр. test.py)
	Output    string // Директорія для вихідного .exe
	AbsInput  string // Глобальний абсолютний шлях до вхідного .py
	AbsOutput string // Глобальний абсолютний шлях до вихідної директорії
	ExePath   string // Шлях до виконуваного бінарника пакера
	ExeDir    string // Шлях до папки, де лежить пакер
	DataDir   string // Шлях до папки pdata (кеш завантажень)

}

// TempConfig відповідає виключно за ізольоване тимчасове середовище
type TempConfig struct {
	BaseDir      string // Базова системна папка temp
	WorkDir      string // Унікальна робоча папка для поточного білда (колишній Temp)
	Mainfiletemp string // посиланя на головний файл скрипта в WorkDir
}

// DepsConfig тримає шляхи до розпакованих залежностей
type DepsConfig struct {
	GoZip       string // Шлях до архіву Go в pdata
	GoTempDir   string // Шлях до розпакованого Go в Temp
	PyZip       string // Шлях до архіву Python в pdata
	PyTempDir   string // Шлях до розпакованого Python в Temp
	Interpreter string // Прямий шлях до python.exe всередині PyTempDir
}

// BuildConfig містить параметри для фінальної генерації та збірки
type BuildConfig struct {
	TargetPyDir  string // Шлях до папки зі скриптом, яку будемо пакувати
	FinalExeName string // Остаточна назва готового бінарника (напр. app_packed.exe)
	FinalExePath string // Повний шлях для збереження готового бінарника
	LdFlags      string // Команди збірки (напр. "-s -w" для зменшення розміру)
	Hash         string // хеш ісходного файла
}

// FlagsConfig тримає налаштування поведінки утиліти
type FlagsConfig struct {
	CLI     bool // Чи використовується інтерфейс CLI
	Verbose bool // Чи виводити логування в консоль (замість Islog)
	RunTest bool // Чи запускати тест перед фінальною збіркою
}

// Config — головна структура, яка об'єднує всі модулі (Single Source of Truth)
type Config struct {
	Paths PathsConfig
	Temp  TempConfig
	Deps  DepsConfig
	Build BuildConfig
	Flags FlagsConfig
}

// New ініціалізує конфіг і задає безпечні значення за замовчуванням
func New() *Config {
	return &Config{
		Build: BuildConfig{
			// Одразу прописуємо дефолтні прапорці компілятора, щоб не хардкодити їх у логіці
			LdFlags: "-s -w",
		},
		Flags: FlagsConfig{
			Verbose: true, // Вмикаємо логи за замовчуванням
			RunTest: true, // робимо тести  зо замовчиванням
		},
	}
}
