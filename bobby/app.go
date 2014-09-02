// bobby is the sample application. 
package bobby

// Application specific constants. For libraries to be able to either be used
// by an interface harness or by other applications, making these into
// variables that are settable might be better for the later use case.
const (
	// appName is the application name.
	appName = "quine"

	// appCode is the code for the application, This is used to both prefix
	// things, like environment variables, and to suffix things, like 
	// application specific file extension. The shortcode should match
	// those appplications
	appCode = "qn"
)


func AppName() string {
	return appName
}

func AppCode() string {
	return appCode
}
