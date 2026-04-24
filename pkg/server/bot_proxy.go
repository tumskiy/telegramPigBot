package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"piggifbot/env"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

var (
	proxyAddr = env.ParseEnv("PROXY_ADDR")
	proxyUser = env.ParseEnv("PROXY_USER")
	proxyPass = env.ParseEnv("PROXY_PASS")

	httpClientBot HttpClientBot
)

type HttpClientBot struct {
	Client *http.Client
}

func getAuth() *proxy.Auth {
	if proxyUser == "" || proxyPass == "" {
		return nil
	}
	return &proxy.Auth{
		User:     proxyUser,
		Password: proxyPass,
	}
}

func NewHTTPClient() *HttpClientBot {
	if proxyAddr == "" {
		logrus.Infof("Прокси не настроен, работаем напрямую")
		return &HttpClientBot{
			Client: &http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					MaxIdleConns:    100,
					IdleConnTimeout: 90 * time.Second,
				},
			},
		}
	}

	dialer, err := proxy.SOCKS5("tcp", proxyAddr, getAuth(), proxy.Direct)
	if err != nil {
		logrus.Errorf("Ошибка настройки прокси %s: %v, работаем без прокси", proxyAddr, err)
		return &HttpClientBot{
			Client: &http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					MaxIdleConns:    100,
					IdleConnTimeout: 90 * time.Second,
				},
			},
		}
	}

	logrus.Infof("SOCKS5 прокси настроен: %s", proxyAddr)
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			return dialer.Dial(network, addr)
		},
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ResponseHeaderTimeout: 60 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
	}

	return &HttpClientBot{
		Client: &http.Client{
			Transport: transport,
			Timeout:   120 * time.Second,
		},
	}
}

var defaultHTTPClient *HttpClientBot

func GetHTTPClient() *HttpClientBot {
	if defaultHTTPClient == nil {
		defaultHTTPClient = NewHTTPClient()
	}
	return defaultHTTPClient
}

func GetProxyInfo() string {
	if proxyAddr == "" {
		return "Прокси не настроен"
	}
	return fmt.Sprintf("Прокси: %s, пользователь: %s", proxyAddr, proxyUser)
}
