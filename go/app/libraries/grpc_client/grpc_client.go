package grpc_client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"microservices/app/libraries/logging"
	proto2 "microservices/app/libraries/proto"
	"mime/multipart"
	"time"
)

const (
	contentType = "image/png"
)

var (
	CannotGetLogger = errors.New("can`t cast interface from context to logging.Logger")
)

type GetAdviceRequest struct {
	Data []*multipart.FileHeader
}

func TellAboutPhoto(ctx context.Context, files *GetAdviceRequest) (*proto2.TellResponse, error) {
	logger, ok := ctx.Value("logger").(*logging.Logger)
	if !ok {
		return nil, CannotGetLogger
	}

	grpcConn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    30 * time.Second,
			Timeout: 20 * time.Second,
		}))

	if err != nil {
		logger.Fatal("error while dial localhost:50051")
		return nil, err
	}
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logger.Error("memory leak (error while closing grpcConn)")
		}
	}(grpcConn)

	tellerAboutPhoto := proto2.NewTellAboutPhotoClient(grpcConn)

	mediaArr := make([]*proto2.MediaRequest, 0, len(files.Data))
	for i, file := range files.Data {
		f, err := file.Open()
		if err != nil {
			logger.Errorf("error while trying to open fileHeader: %s", err.Error())
			return nil, err
		}
		defer f.Close()

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Errorf("error while trying to read bytes from file: %s", err.Error())
			return nil, err
		}

		//fmt.Println(bytes)
		media := &proto2.MediaRequest{
			Data:        bytes,
			ContentType: fmt.Sprintf("%s/%d", contentType, i),
		}

		mediaArr = append(mediaArr, media)
	}

	req := &proto2.TellRequest{
		MediaTellReqArr: mediaArr,
	}

	resp, err := tellerAboutPhoto.GetInfo(ctx, req)
	if err != nil {
		logger.Errorf("error while sending request to neyro service: %s", err.Error())
		return nil, err
	}
	return resp, nil
}
