package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		args, err := parseArgs("show collection --help")
		assert.Nil(t, err)
		assert.Equal(t, []string{"show", "collection", "--help"}, args)
	})

	t.Run("command with arguments", func(t *testing.T) {
		args, err := parseArgs("set deadline --time '2022/08/08 12:00:00'")
		assert.Nil(t, err)
		assert.Equal(t, []string{"set", "deadline", "--time", "2022/08/08 12:00:00"}, args)
	})

	t.Run("command with complext args", func(t *testing.T) {
		args, err := parseArgs("show collection --deadline='2022/08/08 12:00:00' coll_name -v")
		assert.Nil(t, err)
		assert.Equal(t, []string{"show", "collection", "--deadline", "2022/08/08 12:00:00", "coll_name", "-v"}, args)
	})

	t.Run("invalid complex args", func(t *testing.T) {
		args, err := parseArgs("show collection -d 'test\"' --cmd \"cmd-test' coll_name --verbose")
		assert.Error(t, err)
		assert.Nil(t, args)
	})
}
