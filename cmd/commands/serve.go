package commands

import (
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/bogatyr285/golang-boilerplate/api/v1/doer"
	"github.com/bogatyr285/golang-boilerplate/config"
	"github.com/bogatyr285/golang-boilerplate/internal/buildinfo"
	grpcServer "github.com/bogatyr285/golang-boilerplate/internal/grpc"
	"github.com/bogatyr285/golang-boilerplate/internal/grpc/controllers"
	"github.com/bogatyr285/golang-boilerplate/internal/http"
	"github.com/bogatyr285/golang-boilerplate/internal/jaeger"
	"github.com/bogatyr285/golang-boilerplate/internal/prometheus"
	"github.com/bogatyr285/golang-boilerplate/internal/services/doer"
	"github.com/bogatyr285/golang-boilerplate/internal/services/repository"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServeCmd() *cobra.Command {
	var configPath string
	c := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
			defer cancel()
			// todo configure logger
			logger := zap.NewExample()

			var conf config.Config
			err := config.ReadYaml(configPath, &conf)
			if err != nil {
				return err
			}
			logger.Info("config parsed", zap.Any("conf", conf))

			/* all this inits could be in DI*/
			buildInfo := buildinfo.New()

			promExporter, err := prometheus.NewExporter(conf.App.Name)
			if err != nil {
				return err
			}

			jaegerExporter, err := jaeger.NewExporter(conf.Tracing.Endpoint, conf.App.Name, buildInfo)
			if err != nil {
				return err
			}

			instrumentalHTTP, err := http.NewHTTPServer(conf.Metrics.HTTP, logger)
			if err != nil {
				return err
			}

			instrumentalHTTP.Handle("/metrics", promExporter)
			instrumentalHTTP.Handle("/build", buildinfo.Handler(buildInfo))
			instrumentalHTTP.AddSwagger(pb.APISwagger)

			db := repository.NewMongoDB()
			wrappedDB := repository.NewDatabaseRepositoryWithTracing(db, "mongodb")
			doerService := doer.NewDoer(wrappedDB)
			doerGRPCCtrls := controllers.NewDoerController(doerService, logger)
			doerGRPCServer := grpcServer.NewDoerGRPCServer(doerGRPCCtrls, logger)

			grpcSrv, err := grpcServer.NewServer(conf.App.Listen.GRPC, doerGRPCServer, logger)
			if err != nil {
				return err
			}

			gwMux, err := grpcServer.NewGateway(cmd.Context(), conf.App.Listen.GRPC, conf.Listen.HTTP, logger)
			if err != nil {
				return err
			}

			cancelGRPC, err := grpcSrv.Run()
			if err != nil {
				return err
			}

			cancelGw, err := gwMux.Start()
			if err != nil {
				return err
			}

			cancelHTTP, err := instrumentalHTTP.Start()
			if err != nil {
				return err
			}

			<-ctx.Done()
			// close everything
			wg := &sync.WaitGroup{}
			errCh := make(chan error, 5)
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := cancelGw(); err != nil {
					errCh <- err
				}
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := cancelGRPC(); err != nil {
					errCh <- err
				}
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := cancelHTTP(); err != nil {
					errCh <- err
				}
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := jaegerExporter.Stop(); err != nil {
					errCh <- err
				}
			}()

			wg.Wait()
			close(errCh)
			for err := range errCh {
				logger.Error("closing", zap.Error(err))
			}

			return nil
		},
	}
	c.Flags().StringVar(&configPath, "config", "", "path to config")
	return c
}
