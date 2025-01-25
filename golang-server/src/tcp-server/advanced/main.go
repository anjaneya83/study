// This is an advanced TCP server with more commands for reading and writing data with proper error handling and log library
package main

import (
	"net"
	"os"
	"path/filepath"

	"bufio"
	"io"

	"encoding/binary"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Define the log folder & file
	logFolder := "../../../log"
	logFile := "advanced_server.log"
	logFilePath := filepath.Join(logFolder, logFile)

	//open the log file
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	defer file.Close()

	//Create a console encoder same as zap.NewDevelopment()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	//Configure the core to write to the log file
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), //Human-readable format
		zapcore.AddSync(file),                    //Write to file specified above
		zapcore.DebugLevel,                       //Log level
	)

	//Create the logger
	logger := zap.New(core)
	defer logger.Sync()
	// Logger fully configured
	logger.Info("This is an advanced and improved TCP server Program with proper logging and read write capabilities")
	//Create TCP server
	ln, err := net.Listen("tcp", "127.0.0.1:12002")
	if err != nil {
		logger.Error("Failed to start listening on port 12002")
		return
	}
	logger.Info("Server is now listening on port 12002")
	//now we listen for incoming connection and spawn a new goroutine to handle per connection request
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error("Error when accepting new client connection on port 12002")
			return
		}
		logger.Info("Server has been approached by a client")
		go handleConnection(conn, logger)
	}
}

func handleConnection(conn net.Conn, logger *zap.Logger) {
	//Close the connection from server when done
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		//Read the Length of the incoming data ( 4 bytes, big endian)
		lengthBytes := make([]byte, 4)
		_, err := io.ReadFull(reader, lengthBytes)
		if err != nil {
			if err == io.EOF {
				logger.Warn("Client closed the connection")
			} else {
				logger.Error("Error reading the data length")
			}
			return
		}
		dataLength := binary.BigEndian.Uint32(lengthBytes)
		logger.Info("Expecting bytes of data", zap.Int("length", int(dataLength)))
		//Now read the actual data
		data := make([]byte, dataLength)   //creat a buffer to store data
		_, err = io.ReadFull(reader, data) // read data from tcp stream
		if err != nil {
			logger.Error("Error reading the data")
			return
		}
		logger.Info("Received:", zap.String("data", string(data)))
		// send a response back to client
		response := "Message received"
		err = sendMessage(conn, []byte(response))
		if err != nil {
			logger.Error("Error sending response back to client")
			return
		}
	}
}

// sendMessage sends a message with a length prefix
func sendMessage(conn net.Conn, message []byte) error {
	length := uint32(len(message)) //as protocol, first send how many bytes you are sending
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length) //convert the int into big endian format
	// send the length and the message
	_, err := conn.Write(append(lengthBytes, message...))
	return err
}
