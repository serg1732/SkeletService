package skeletservice

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/serg1732/SkeletService/pkg/constants"
	"github.com/serg1732/SkeletService/pkg/consulwrapper"
	"github.com/serg1732/SkeletService/pkg/loggers"

	"github.com/gin-gonic/gin"
)

type Service struct {
	router       *gin.Engine
	consulClient *consulwrapper.ConsulClient
	exitDone     *sync.WaitGroup
	httpServer   *http.Server
	logger       *loggers.ILogger
}

func NewService(handlerEngine *gin.Engine, log loggers.ILogger) *Service {
	return &Service{router: handlerEngine,
		consulClient: consulwrapper.NewConsulClient(),
		exitDone:     &sync.WaitGroup{},
		logger:       &log,
	}
}

func (s *Service) Start() error {

	var err error
	if s.router == nil {
		textError := "Error! Not initialize handler Service!"
		return errors.New(textError)
	}
	s.consulClient.RegisterService()

	s.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(constants.ServicePort),
		Handler: s.router,
	}

	go func() {
		defer s.exitDone.Done()

		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	defer s.consulClient.DeRegisterService()
	s.exitDone.Add(1)
	s.exitDone.Wait()
	return err
}
func (s *Service) AddHandler(rtype constants.RequestType, path string, handler gin.HandlerFunc) error {
	var err error
	if s.router == nil {
		s.router = gin.Default()
	}
	switch rtype {
	case constants.GET:
		s.router.GET(path, handler)
	case constants.POST:
		s.router.POST(path, handler)
	case constants.PUT:
		s.router.PUT(path, handler)
	default:
		err = errors.New("Unknown method!")
	}
	return err
}
