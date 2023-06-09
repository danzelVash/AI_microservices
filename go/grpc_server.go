package microservices

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	proto2 "microservices/app/libraries/proto"
	"net"
)

type TellerServer struct {
	proto2.UnimplementedTellAboutPhotoServer
}

func NewTellerServer() *TellerServer {
	return &TellerServer{}
}

func (ts *TellerServer) GetInfo(ctx context.Context, req *proto2.TellRequest) (*proto2.TellResponse, error) {
	logrus.Infof("request was given")
	var mediaRespArr []*proto2.MediaResponse
	for i, mediaReq := range req.MediaTellReqArr {
		mediaResp := proto2.MediaResponse{
			Data:        mediaReq.Data,
			ContentType: mediaReq.ContentType,
			Description: fmt.Sprintf("Вы отправили фотографию %d", i),
		}

		mediaRespArr = append(mediaRespArr, &mediaResp)
	}

	resp := &proto2.TellResponse{
		MediaTellRespArr: mediaRespArr,
		DescriptionAll:   "Какой же вы молодец что вот так заботитесь о себе! Надо меньше стрессовать",
	}

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatal("error listen port")
	}
	srv := grpc.NewServer()

	proto2.RegisterTellAboutPhotoServer(srv, NewTellerServer())

	fmt.Println("starting service on port 50051...")
	srv.Serve(listen)
}
