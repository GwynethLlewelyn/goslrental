package main

import (
	"fmt"
	"database/sql"
	"log"
	"github.com/cznic/ql"
	"github.com/op/go-logging" // more complete package to log to different outputs; we start with file, syslog, and stderr; 
	"github.com/spf13/viper" // to read config files
	"runtime"
	"path/filepath"
	"time"
)

var (
	// Default configurations, hopefully exported to other files and packages
	// we probably should have a struct for this (or even several)
	Host, GoSLRentalDSN, URLPathPrefix, PDO_Prefix, PathToStaticFiles,
	ServerPort, FrontEnd, LSLSignaturePIN string
	logFileName string = "goslrental.log"
	logMaxSize, logMaxBackups, logMaxAge int // configuration for the go-logging logger
	logSeverityStderr, logSeverityFile, logSeveritySyslog logging.Level // more configuration for the go-logging logger
	Log = logging.MustGetLogger("goslrental")	// configuration for the go-logging logger, must be available everywhere
	logFormat logging.Formatter	// must be initialised or all hell breaks loose
)

const NullUUID = "00000000-0000-0000-0000-000000000000" // always useful when we deal with SL/OpenSimulator...

//type templateParameters map[string]string
type templateParameters map[string]interface{}

// setUp tries to create a table on the QL database for testing purposes.
func setUp(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS note (
	  id INT 
	  ,title STRING
	  ,body STRING
	  ,created_at STRING
	  ,updated_at STRING
	);
	`)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// loadConfiguration loads all the configuration from the config.toml file.
// It's a separate function because we want to be able to do a killall -HUP goslrental to force the configuration to be read again.
// Also, if the configuration file changes, this ought to read it back in again without the need of a HUP signal (20170811).
func loadConfiguration() {
	fmt.Print("Reading goslrental configuration:")	// note that we might not have go-logging active as yet, so we use fmt
	// Open our config file and extract relevant data from there
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return	// we might still get away with this!
	}
	// Without these set, we cannot do anything
	viper.SetDefault("goslrental.Host", "localhost") // to prevent bombing out with panics
	Host = viper.GetString("goslrental.Host"); fmt.Print(".")
	viper.SetDefault("goslrental.URLPathPrefix", "") // empty by default, but you might add a 'main' website for information later
	URLPathPrefix = viper.GetString("goslrental.URLPathPrefix"); fmt.Print(".")
	GoSLRentalDSN = viper.GetString("goslrental.GoSLRentalDSN"); fmt.Print(".")
	viper.SetDefault("PDO_Prefix", "ql") // for now, nothing else will work anyway...
	PDO_Prefix = viper.GetString("goslrental.PDO_Prefix"); fmt.Print(".")
	viper.SetDefault("goslrental.PathToStaticFiles", "~/go/src/goslrental")
	path, err := expandPath(viper.GetString("goslrental.PathToStaticFiles")); fmt.Print(".")
	if err != nil {
		fmt.Println("Error expanding path:", err)
		path = ""	// we might get away with this as well
	}
	PathToStaticFiles = path
	viper.SetDefault("goslrental.ServerPort", ":3333")
	ServerPort = viper.GetString("goslrental.ServerPort"); fmt.Print(".")
	FrontEnd = viper.GetString("goslrental.FrontEnd"); fmt.Print(".")
	viper.SetDefault("goslrental.LSLSignaturePIN", "9876") // better than no signature at all
	LSLSignaturePIN = viper.GetString("opensim.LSLSignaturePIN"); fmt.Print(".")
	// logging options
	viper.SetDefault("log.FileName", "log/goslrental.log")
	logFileName = viper.GetString("log.FileName"); fmt.Print(".")
	viper.SetDefault("log.Format", `%{color}%{time:2006/01/02 15:04:05.0} %{shortfile} - %{shortfunc} ▶ %{level:.4s}%{color:reset} %{message}`)
	logFormat = logging.MustStringFormatter(viper.GetString("log.Format")); fmt.Print(".")
	viper.SetDefault("log.MaxSize", 500)
	logMaxSize = viper.GetInt("log.MaxSize"); fmt.Print(".")
	viper.SetDefault("log.MaxBackups", 3)
	logMaxBackups = viper.GetInt("log.MaxBackups"); fmt.Print(".")
	viper.SetDefault("log.MaxAge", 28)
	logMaxAge = viper.GetInt("log.MaxAge"); fmt.Print(".")
	viper.SetDefault("log.SeverityStderr", logging.DEBUG)
	switch viper.GetString("log.SeverityStderr") {
		case "CRITICAL":
			logSeverityStderr = logging.CRITICAL
    	case "ERROR":
			logSeverityStderr = logging.ERROR
    	case "WARNING":
			logSeverityStderr = logging.WARNING
    	case "NOTICE":
			logSeverityStderr = logging.NOTICE
    	case "INFO":
			logSeverityStderr = logging.INFO
    	case "DEBUG":
			logSeverityStderr = logging.DEBUG
		// default case is handled directly by viper
	}
	fmt.Print(".")
	viper.SetDefault("log.SeverityFile", logging.DEBUG)
	switch viper.GetString("log.SeverityFile") {
		case "CRITICAL":
			logSeverityFile = logging.CRITICAL
    	case "ERROR":
			logSeverityFile = logging.ERROR
    	case "WARNING":
			logSeverityFile = logging.WARNING
    	case "NOTICE":
			logSeverityFile = logging.NOTICE
    	case "INFO":
			logSeverityFile = logging.INFO
    	case "DEBUG":
			logSeverityFile = logging.DEBUG
	}
	fmt.Print(".")
	viper.SetDefault("log.SeveritySyslog", logging.CRITICAL) // we don't want to swamp syslog with debugging messages!!
	switch viper.GetString("log.SeveritySyslog") {
		case "CRITICAL":
			logSeveritySyslog = logging.CRITICAL
    	case "ERROR":
			logSeveritySyslog = logging.ERROR
    	case "WARNING":
			logSeveritySyslog = logging.WARNING
    	case "NOTICE":
			logSeveritySyslog = logging.NOTICE
    	case "INFO":
			logSeveritySyslog = logging.INFO
    	case "DEBUG":
			logSeveritySyslog = logging.DEBUG
	}
	fmt.Print(".")
	fmt.Println("read!")	// note that we might not have go-logging active as yet, so we use fmt
	
	// Setup the lumberjack rotating logger. This is because we need it for the go-logging logger when writing to files. (20170813)
	rotatingLogger := &lumberjack.Logger{
	    Filename:   logFileName,	// this is an option set on the config.yaml file, eventually the others will be so, too.
	    MaxSize:    logMaxSize, // megabytes
	    MaxBackups: logMaxBackups,
	    MaxAge:     logMaxAge, //days
	}
	// Setup the go-logging Logger. (20170812) We have three loggers: one to stderr, one to a logfile, one to syslog for critical stuff. (20170813
	backendStderr	:= logging.NewLogBackend(os.Stderr, "", 0)
	backendFile		:= logging.NewLogBackend(rotatingLogger, "", 0)
	backendSyslog,_	:= logging.NewSyslogBackend("")

	// Set formatting for stderr and file (basically the same). I'm assuming syslog has its own format, but I'll have to see what happens (20170813).
	backendStderrFormatter	:= logging.NewBackendFormatter(backendStderr, logFormat)
	backendFileFormatter	:= logging.NewBackendFormatter(backendFile, logFormat)

	// Check if we're overriding the default severity for each backend. This is user-configurable. By default: DEBUG, DEBUG, CRITICAL.
	// TODO(gwyneth): What about a WebSocket backend using https://github.com/cryptix/exp/wslog ? (20170813)
	backendStderrLeveled := logging.AddModuleLevel(backendStderrFormatter)
	backendStderrLeveled.SetLevel(logSeverityStderr, "goslrental")
	backendFileLeveled := logging.AddModuleLevel(backendFileFormatter)
	backendFileLeveled.SetLevel(logSeverityFile, "goslrental")
	backendSyslogLeveled := logging.AddModuleLevel(backendSyslog)
	backendSyslogLeveled.SetLevel(logSeveritySyslog, "goslrental")

	// Set the backends to be used. Logging should commence now.
	logging.SetBackend(backendStderrLeveled, backendFileLeveled, backendSyslogLeveled)
	fmt.Println("Logging set up.")
}

// main starts here.
func main() {
		// to change the flags on the default logger
	// see https://stackoverflow.com/a/24809859/1035977
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Config viper, which reads in the configuration file every time it's needed.
	// Note that we need some hard-coded variables for the path and config file name.
	viper.SetConfigName("config")
	viper.SetConfigType("toml") // just to make sure; it's the same format as OpenSimulator (or MySQL) config files
	viper.AddConfigPath("$HOME/go/src/goslrental/") // that's how I have it
	viper.AddConfigPath("$HOME/go/src/github.com/GwynethLlewelyn/goslrental/") // that's how you'll have it
	viper.AddConfigPath(".")               // optionally look for config in the working directory

	loadConfiguration() // this gets loaded always, on the first time it runs
	viper.WatchConfig() // if the config file is changed, this is supposed to reload it (20170811)
	viper.OnConfigChange(func(e fsnotify.Event) {
		if (Log == nil) {
			fmt.Println("Config file changed:", e.Name) // if we couldn't configure the logging subsystem, it's better to print it to the console
		} else {
			Log.Info("Config file changed:", e.Name)
		}
		loadConfiguration() // I think that this needs to be here, or else, how does Viper know what to call?
	})
	
	// prepares a special channel to look for termination signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGCONT)

	// goroutine which listens to signals and calls the loadConfiguration() function if someone sends us a HUP
	go func() {
		for {
	        sig := <-sigs
	        Log.Notice("Got signal", sig)
	        switch sig {
		        case syscall.SIGUSR1:
		        case syscall.SIGUSR2:
		        case syscall.SIGHUP:
		        case syscall.SIGCONT:
				default:
		        	Log.Warning("Unknown UNIX signal", sig, "caught!! Ignoring...")
	        }
        }
    }()
	
	ql.RegisterDriver()	// this should allow us to use the 'normal' SQL Go bindings to use QL.

	// do some database tests. If it fails, it means the database is broken or corrupted and it's worthless
	//  to run this application anyway!
	Log.Info("Testing opening database connection at ", GoBotDSN, "\nPath to static files is:", PathToStaticFiles)

	db, err := sql.Open(PDO_Prefix, GoSLRentalDSN)
	defer db.Close()
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}
	
	if err = setUp(db); err != nil {
		log.Fatalf("failed to create table: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %s", err)
	}
	
	// now insert some stuff and read from the database.
	tx, err := db.Begin()
	checkErr(err)
	
	stmt, err := tx.Prepare("INSERT INTO note (id, title, body, created_at, updated_at) VALUES (?1,?2,?3,?4,?5)");
	checkErrPanic(err)

	defer stmt.Close()
	
	curTime := fmt.Sprintf("%s", time.Now())
	
	_, err = stmt.Exec(1, "this is my note", "blah blah", curTime, curTime)
	checkErr(err)

	err = tx.Commit()	
	checkErr(err)
	
		// Now prepare the web interface

	// Load all templates
	err = GoSLRentalTemplates.init(PathToStaticFiles + "/templates/*.tpl")
	checkErr(err) // abort if templates are not found

	http.HandleFunc(URLPathPrefix + "/admin/user-management/",			backofficeUserManagement)
	
	http.HandleFunc(URLPathPrefix + "/admin/lsl-register-object/",		backofficeLSLRegisterObject)

	// fallthrough for admin
	http.HandleFunc(URLPathPrefix + "/admin/",							backofficeMain)
	http.HandleFunc(URLPathPrefix + "/",								backofficeLogin) // if not auth, then get auth

    err = http.ListenAndServe(ServerPort, nil) // set listen port
    checkErr(err) // if it can't listen to all the above, then it has to abort anyway

}

// checkErrPanic logs a fatal error and panics.
func checkErrPanic(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		Log.Panic(filepath.Base(file), ":", line, ":", pc, ok, " - panic:", err)
	}
}

// checkErr checks if there is an error, and if yes, it logs it out and continues.
//  this is for 'normal' situations when we want to get a log if something goes wrong but do not need to panic
func checkErr(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		Log.Error(filepath.Base(file), ":", line, ":", pc, ok, " - error:", err)
	}
}