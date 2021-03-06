package shell

import "testing"

func TestZsh_SetupWrapper(t *testing.T) {
	z := zsh{}

	result := z.SetupWrapper("shell_logger")
	expected := "shell_logger --mode=wrapper"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestZsh_SetupHooks(t *testing.T) {
	z := zsh{}

	result := z.SetupHooks("shell_logger", "test.db")
	var expected = `
preexec () {
	__SHELL_LOGGER_START_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
}
precmd () {
	export __SHELL_LOGGER_RETURN_CODE=$?
	export __SHELL_LOGGER_COMMAND=$(fc -ln -1)
	export __SHELL_LOGGER_FAILED_COMMAND=$(fc -ln -3 | head -n 1)
	export __SHELL_LOGGER_FUCK_CMD=$(fc -ln -2 | head -n 1)
	export __SHELL_LOGGER_DB_PATH=test.db
	[ "$__SHELL_LOGGER_FUCK_CMD" = "fuck" ] && [ "$__SHELL_LOGGER_RETURN_CODE" = "0" ] && shell_logger -mode submit
}
`
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
