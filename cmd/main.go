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
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // #nosec G102
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ip, err := getOutboundIP()
	if err != nil {
		log.Fatalf("Failed to get outbound IP: %v", err)
	}
	fmt.Printf("Outbound IP: %s\n", ip)

	store := storage.NewInMemoryStorage()
	geoClient := geocoder.NewOpenStreetMapGeocoder()
	srv := server.NewCityManagerServer(store, geoClient)

	s := grpc.NewServer()
	pb.RegisterCityManagerServer(s, srv)

	// Реализация Graceful shutdown
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

func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection: %w", err)
	}
	defer func() {
		closeErr := conn.Close()
		if closeErr != nil {
			log.Printf("Error closing connection: %v", closeErr)
		}
	}()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return nil, fmt.Errorf("failed to get local address")
	}

	return localAddr.IP, nil
}
