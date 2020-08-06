// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vearne/tinyurl/config"
	zlog "github.com/vearne/tinyurl/log"
	"github.com/vearne/tinyurl/router"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web service",
	Long:  `web service`,
	Run:   runWeb,
}

func init() {
	RootCmd.AddCommand(webCmd)
}

func runWeb(cmd *cobra.Command, args []string) {
	go func() {
		//golang /debug/pprof
		http.ListenAndServe("0.0.0.0:18080", nil)
	}()

	// ------------- run ---------------------
	mode := viper.GetString("RUN_MODE")
	if mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	ginServer := router.NewServer()
	server := &http.Server{
		Addr:           config.GetOpts().Web.ListenAddress,
		Handler:        ginServer,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()
	// 设置优雅退出
	gracefulExitWeb(server)
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	zlog.Info("got a signal", zap.Any("signal", sig))
	now := time.Now()
	// 最长等待3秒
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(cxt)
	if err != nil {
		zlog.Error("shutdown error", zap.Error(err))
	}
	// 看看实际退出所耗费的时间
	zlog.Info("------exited--------",
		zap.Duration("exit_cost", time.Since(now)))
}
