package gifcache

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"piggifbot/env"
	"strconv"
	"strings"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var myID = env.ParseEnv("MY_ID")

type CacheGifs struct {
	mu        sync.RWMutex
	bot       *tgbotapi.BotAPI
	cacheData *CacheData
	cacheDir  string
}

type Cache struct {
	FileName string `json:"file_name"`
	TgFileID string `json:"tg_file_id"`
	Command  string `json:"command"`
	Caption  string `json:"caption"`
}

type CacheData struct {
	Gifs map[string]Cache
}

func NewCache(bot *tgbotapi.BotAPI, cacheDir string) *CacheGifs {
	return &CacheGifs{
		bot:      bot,
		cacheDir: cacheDir,
		cacheData: &CacheData{
			Gifs: make(map[string]Cache),
		},
	}
}

func (c *CacheGifs) LoadAllGifsWithCaptions(dirPath, commandName string, captionsMap map[string]string) error {
	pattern := filepath.Join(dirPath, "*.gif")

	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("ошибка поиска гифок: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("в директории %s нет гифок", dirPath)
	}

	loaded := 0
	for _, filePath := range files {
		fileName := filepath.Base(filePath)

		c.mu.RLock()
		_, exists := c.cacheData.Gifs[fileName]
		c.mu.RUnlock()

		if exists {
			logrus.Debugf("Gif %s already in cache, skipping", fileName)
			continue
		}

		fileID, err := c.uploadGifToTelegram(filePath)
		if err != nil {
			logrus.Errorf("Ошибка загрузки %s: %v", fileName, err)
			continue
		}

		caption := ""
		if captionsMap != nil {
			if val, ok := captionsMap[fileName]; ok {
				caption = val
			}
		}

		c.mu.Lock()
		c.cacheData.Gifs[fileName] = Cache{
			FileName: fileName,
			TgFileID: fileID,
			Command:  commandName,
			Caption:  caption,
		}
		c.mu.Unlock()

		loaded++
		logrus.Infof("Загружена гифка %s, file_id: %s...", fileName, fileID[:20])
		time.Sleep(500 * time.Millisecond)
	}

	if loaded > 0 {
		if err := c.SaveCache(); err != nil {
			logrus.Errorf("Ошибка сохранения кэша: %v", err)
		}
	}

	return nil
}

func (c *CacheGifs) uploadGifToTelegram(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileReader := tgbotapi.FileReader{
		Name:   filepath.Base(filePath),
		Reader: file,
	}

	myid, _ := strconv.Atoi(myID)
	msg := tgbotapi.NewDocument(int64(myid), fileReader)

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		resp, err := c.bot.Send(msg)
		if err == nil {
			return resp.Document.FileID, nil
		}
		if strings.Contains(err.Error(), "429") {
			waitTime := time.Duration(i+1) * 30 * time.Second
			logrus.Warnf("Rate limit, waiting %v before retry %d/%d", waitTime, i+1, maxRetries)
			time.Sleep(waitTime)
			continue
		}
		return "", fmt.Errorf("ошибка отправки в Telegram: %w", err)
	}

	return "", fmt.Errorf("failed after %d retries", maxRetries)
}

func (c *CacheGifs) SaveCache() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cacheFile := filepath.Join(c.cacheDir, "gif_cache.json")

	if err := os.MkdirAll(c.cacheDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории кэша: %w", err)
	}

	file, err := os.Create(cacheFile)
	if err != nil {
		return fmt.Errorf("ошибка создания файла кэша: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c.cacheData); err != nil {
		return fmt.Errorf("ошибка сохранения кэша: %w", err)
	}

	logrus.Infof("Кэш сохранён: %d гифок", len(c.cacheData.Gifs))
	return nil
}

func (c *CacheGifs) GetRandomGif(commandName string) (Cache, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var commandGifs []Cache
	for _, gif := range c.cacheData.Gifs {
		if gif.Command == commandName {
			commandGifs = append(commandGifs, gif)
		}
	}

	if len(commandGifs) == 0 {
		return Cache{}, fmt.Errorf("no gifs found for command: %s", commandName)
	}

	random := rand.Intn(len(commandGifs))
	return commandGifs[random], nil
}

func (c *CacheGifs) GetGifByName(fileName string) (Cache, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	gif, exists := c.cacheData.Gifs[fileName]
	if !exists {
		return Cache{}, fmt.Errorf("gif not found: %s", fileName)
	}
	return gif, nil
}

func (c *CacheGifs) PreloadCache() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheFile := filepath.Join(c.cacheDir, "gif_cache.json")

	file, err := os.Open(cacheFile)
	if err != nil {
		if os.IsNotExist(err) {
			logrus.Info("Cache file not found, will create new")
			return nil
		}
		return fmt.Errorf("failed to open cache file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	tempData := &CacheData{
		Gifs: make(map[string]Cache),
	}

	if err := decoder.Decode(tempData); err != nil {
		return fmt.Errorf("failed to decode cache: %w", err)
	}

	c.cacheData = tempData
	logrus.Infof("Loaded %d gifs from cache", len(c.cacheData.Gifs))
	return nil
}
