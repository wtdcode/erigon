package downloader

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	dir2 "github.com/ledgerwatch/erigon-lib/common/dir"
)

// TorrentFiles - does provide thread-safe CRUD operations on .torrent files
type TorrentFiles struct {
	lock sync.Mutex
	dir  string
}

func NewAtomicTorrentFiles(dir string) *TorrentFiles {
	return &TorrentFiles{dir: dir}
}

func (tf *TorrentFiles) Exists(name string) bool {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.exists(name)
}

func (tf *TorrentFiles) exists(name string) bool {
	fPath := filepath.Join(tf.dir, name)
	return dir2.FileExist(fPath + ".torrent")
}
func (tf *TorrentFiles) Delete(name string) error {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.delete(name)
}

func (tf *TorrentFiles) delete(name string) error {
	fPath := filepath.Join(tf.dir, name)
	return os.Remove(fPath + ".torrent")
}

func (tf *TorrentFiles) Create(torrentFilePath string, res []byte) error {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.create(torrentFilePath, res)
}
func (tf *TorrentFiles) create(torrentFilePath string, res []byte) error {
	if len(res) == 0 {
		return fmt.Errorf("try to write 0 bytes to file: %s", torrentFilePath)
	}
	f, err := os.Create(torrentFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(res); err != nil {
		return err
	}
	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

func (tf *TorrentFiles) CreateTorrentFromMetaInfo(fPath string, mi *metainfo.MetaInfo) error {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.createTorrentFromMetaInfo(fPath, mi)
}
func (tf *TorrentFiles) createTorrentFromMetaInfo(fPath string, mi *metainfo.MetaInfo) error {
	file, err := os.Create(fPath + ".tmp")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := mi.Write(file); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(fPath+".tmp", fPath); err != nil {
		return err
	}
	return nil
}

func (tf *TorrentFiles) LoadByName(fName string) (*torrent.TorrentSpec, error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	fPath := filepath.Join(tf.dir, fName+".torrent")
	return tf.load(fPath)
}

func (tf *TorrentFiles) LoadByPath(fPath string) (*torrent.TorrentSpec, error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.load(fPath)
}

func (tf *TorrentFiles) load(fPath string) (*torrent.TorrentSpec, error) {
	mi, err := metainfo.LoadFromFile(fPath)
	if err != nil {
		return nil, fmt.Errorf("LoadFromFile: %w, file=%s", err, fPath)
	}
	mi.AnnounceList = Trackers
	return torrent.TorrentSpecFromMetaInfoErr(mi)
}

const ProhibitNewDownloadsFileName = "prohibit_new_downloads.lock"

// Erigon "download once" - means restart/upgrade/downgrade will not download files (and will be fast)
// After "download once" - Erigon will produce and seed new files
// Downloader will able: seed new files (already existing on FS), download uncomplete parts of existing files (if Verify found some bad parts)
func (tf *TorrentFiles) prohibitNewDownloads() error {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	fPath := filepath.Join(tf.dir, ProhibitNewDownloadsFileName)
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}
func (tf *TorrentFiles) newDownloadsAreProhibited() bool {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return dir2.FileExist(filepath.Join(tf.dir, ProhibitNewDownloadsFileName))
}
