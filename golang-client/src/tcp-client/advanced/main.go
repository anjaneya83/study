// This is an advanced TCP client program with proper logging and error handling
package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"net"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//Define the log folder & file
	logFolder := "../../../log"
	logFile := "advanced_client.log"
	logFilePath := filepath.Join(logFolder, logFile)

	//open th log file
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file:" + err.Error())
	}
	defer file.Close()
	//create a console encoder same as zap.NewDevelopment
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	//configure the core to write to the log file
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), //Human-readable format
		zapcore.AddSync(file),                    //write to file specified above
		zapcore.DebugLevel,
	)

	//create the logger
	logger := zap.New(core)
	defer logger.Sync()
	//Logger fully configured
	logger.Info("This is an advanced TCP client program with proper error handling and logging")
	conn, err := net.Dial("tcp", "127.0.0.1:12002")
	if err != nil {
		logger.Error("Error connecting to server on port 12002")
		return
	}
	defer conn.Close()
	logger.Info("Connection established with the server")
	message := "Hello my friend,the server, from the client"
	err = sendMessage(conn, []byte(message))
	if err != nil {
		logger.Info("Error sending message to server")
		return
	}

	// Read response from server
	reader := bufio.NewReader(conn) //link the tcp socket with the io stream
	response, err := readMessage(reader)
	if err != nil {
		logger.Error("Error reading response from server")
		return
	}
	logger.Info("Server response:", zap.String("response", string(response)))
}

// sendMessage sends a message with a length prefix
func sendMessage(conn net.Conn, message []byte) error {
	length := uint32(len(message))
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)
	// write the length and the message in separate writes
	_, err := conn.Write(lengthBytes)
	if err != nil {
		return err
	}
	_, err = conn.Write(message)
	return err
}

// readMessage reads a message with a length prefix
func readMessage(reader *bufio.Reader) (string, error) {
	lengthBytes := make([]byte, 4)
	_, err := io.ReadFull(reader, lengthBytes)
	if err != nil {
		return "", err
	}
	dataLength := binary.BigEndian.Uint32(lengthBytes) //convert the received message length into uint32
	data := make([]byte, dataLength)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
