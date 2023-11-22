package vz_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"

	"github.com/Code-Hex/vz/v3"
)

func TestVirtioSocketConnectFailure(t *testing.T) {
	var conn net.Conn
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	container := newVirtualizationMachine(t)
	defer container.Close()

	vm := container.VirtualMachine

	socketDevice := vm.SocketDevices()[0] // already tested in newVirtualizationMachine

	port := 43218
	done := make(chan struct{})

	go func() {
		var err error
		defer close(done)

		conn, err = socketDevice.Connect(uint32(port))
		if err != nil {
			t.Errorf("failed to accept connection: %v", err)
			return
		}
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatalf("timeout when attempting to connect to vsock")
	}
}

func TestVirtioSocketListener(t *testing.T) {
	container := newVirtualizationMachine(t)
	t.Cleanup(func() {
		if err := container.Shutdown(); err != nil {
			log.Println(err)
		}
	})

	vm := container.VirtualMachine

	socketDevice := vm.SocketDevices()[0] // already tested in newVirtualizationMachine

	port := 43218
	wantData := "hello"
	done := make(chan struct{})

	listener, err := socketDevice.Listen(uint32(port))
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	go func() {
		defer close(done)

		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("failed to accept connection: %v", err)
			return
		}
		defer conn.Close()

		destPort := conn.(*vz.VirtioSocketConnection).DestinationPort()
		if port != int(destPort) {
			t.Errorf("want destination port %d but got %d", destPort, port)
			return
		}

		buf := make([]byte, len(wantData))
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			t.Errorf("failed to read data: %v", err)
			return
		}
		got := string(buf[:n])

		if wantData != got {
			t.Errorf("want %q but got %q", wantData, got)
		}
	}()

	session := container.NewSession(t)
	var buf bytes.Buffer
	session.Stderr = &buf
	cmd := fmt.Sprintf("echo %s | socat - VSOCK-CONNECT:2:%d", wantData, port)
	if err := session.Run(cmd); err != nil {
		t.Fatalf("failed to write data to vsock: %v\nstderr: %q", err, buf)
	}
	session.Close()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatalf("timeout connection handling after accepted")
	}
}
