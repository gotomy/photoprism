package photoprism

import (
	"testing"

	"github.com/photoprism/photoprism/internal/config"
)

func TestIndexer_IndexAll(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	conf.InitializeTestData(t)

	tensorFlow := NewTensorFlow(conf)

	indexer := NewIndexer(conf, tensorFlow)

	converter := NewConverter(conf)

	importer := NewImporter(conf, indexer, converter)

	importer.ImportPhotosFromDirectory(conf.ImportPath())

	options := IndexerOptionsAll()

	indexer.IndexOriginals(options)
}
