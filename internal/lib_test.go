package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readDBFromFile(t *testing.T) {
	t.Parallel()

	t.Run("works with non existing file", func(t *testing.T) {
		t.Parallel()

		var want []string

		got, err := readDBFromFile("non_existing_file.txt")
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("handles empty file", func(t *testing.T) {
		t.Parallel()

		tmpfile, err := os.CreateTemp("", "some_file")
		assert.NoError(t, err)
		defer func() {
			_ = os.Remove(tmpfile.Name())
		}()

		var want []string

		got, err := readDBFromFile(tmpfile.Name())
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("reads data from file", func(t *testing.T) {
		t.Parallel()

		tmpfile, err := os.CreateTemp("", "some_file")
		assert.NoError(t, err)
		defer func() {
			_ = os.Remove(tmpfile.Name())
		}()

		err = os.WriteFile(tmpfile.Name(), []byte("s01e02\ns01e01\n"), 0666)
		assert.NoError(t, err)

		want := []string{
			"s01e01",
			"s01e02",
		}

		got, err := readDBFromFile(tmpfile.Name())
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}

func Test_selectNextEpisode(t *testing.T) {
	t.Parallel()

	t.Run("returns any episode at all", func(t *testing.T) {
		t.Parallel()

		got, err := selectNextEpisode(nil)
		assert.NoError(t, err)
		assert.Contains(t, EPISODES, got)
	})

	t.Run("error if there is no unseen episodes", func(t *testing.T) {
		t.Parallel()

		_, err := selectNextEpisode(EPISODES[:])
		assert.EqualError(t, err, "no unseen episodes")
	})
}

func Test_saveDBToFile(t *testing.T) {
	t.Parallel()

	t.Run("saves empty list to file", func(t *testing.T) {
		t.Parallel()

		tmpfile, err := os.CreateTemp("", "some_file")
		assert.NoError(t, err)
		defer func() {
			_ = os.Remove(tmpfile.Name())
		}()

		err = saveDBToFile(nil, tmpfile.Name())
		assert.NoError(t, err)

		bytes, err := os.ReadFile(tmpfile.Name())
		assert.NoError(t, err)
		assert.Len(t, bytes, 0)
	})

	t.Run("saves non empty list to file in reverse order", func(t *testing.T) {
		t.Parallel()

		tmpfile, err := os.CreateTemp("", "some_file")
		assert.NoError(t, err)
		defer func() {
			_ = os.Remove(tmpfile.Name())
		}()

		seenEpisodes := []string{
			"s01e01",
			"s01e02",
		}

		err = saveDBToFile(seenEpisodes, tmpfile.Name())
		assert.NoError(t, err)

		bytes, err := os.ReadFile(tmpfile.Name())
		assert.NoError(t, err)
		assert.Equal(t, "s01e02\ns01e01\n", string(bytes))
	})
}
