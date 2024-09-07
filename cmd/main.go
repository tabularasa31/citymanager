package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/tabularasa31/citymanager/api/gen"
	"github.com/tabularasa31/citymanager/internal/geocoder"
	"github.com/tabularasa31/citymanager/internal/server"
	"github.com/tabularasa31/citymanager/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ip := getOutboundIP()

	store := storage.NewInMemoryStorage()
	geocoder := geocoder.NewOpenStreetMapGeocoder()
	srv := server.NewCityManagerServer(store, geocoder)

	s := grpc.NewServer()
	pb.RegisterCityManagerServer(s, srv)

	// Реализация Grasefull shutdown
	// Канал для получения сигналов от операционной системы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("gRPC server is running and listening on %s:50051...", ip)
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Failed to serve: %v\n", err)
		}
	}()

	// Ожидаем сигналы на завершение
	<-sigChan
	fmt.Println("Graceful shutdown initiated...")

	// Плавное завершение
	s.GracefulStop()
	fmt.Println("gRPC server stopped gracefully")
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
