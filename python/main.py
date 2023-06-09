from concurrent import futures
import grpc
from main_pb2_grpc import TellAboutPhotoServicer, add_TellAboutPhotoServicer_to_server
from main_pb2 import TellRequest, TellResponse, MediaResponse
# from define_emotion_neyro.test_neyro import gen_rand_description
from libraries.content_util.loader import load_to_ssd, delete_from_ssd
from define_emotion.model_main import detect_faces, describe_photo


WITHOUT_PEOPLE = "Вероятнее всего на этом фото нет людей. Если я ошибся, " \
                 "попробуйте еще раз или отправьте другую фотографию."


class TellAboutPhotoServer(TellAboutPhotoServicer):
    def GetInfo(self, request: TellRequest, context) -> TellResponse:
        filePaths = []
        mediaRespArr = []
        for i in range(len(request.mediaTellReqArr)):
            file_path = load_to_ssd(request.mediaTellReqArr[i].data, str(i))
            if file_path == "":
                mediaResp = MediaResponse(
                    data=bytes(),
                    contentType=request.mediaTellReqArr[i].contentType,
                    description=f"Фотография {i} некорректна или была отправлена в неподдерживаемом формате."
                )
                mediaRespArr.append(mediaResp)
                continue

            filePaths.append(file_path)

            mediaResp = MediaResponse(
                data=bytes(),
                contentType=request.mediaTellReqArr[i].contentType,
                description=WITHOUT_PEOPLE

            )
            if detect_faces(file_path):
                emotion = ""
                try:
                    emotion = describe_photo(file_path)
                except Exception as ex:
                    print(f"unknown error occured while trying to describe photo {ex}")
                mediaResp.description = emotion
                mediaResp.data = request.mediaTellReqArr[i].data

            mediaRespArr.append(mediaResp)
        delete_from_ssd(filePaths)

        return TellResponse(
            mediaTellRespArr=mediaRespArr,
            descriptionAll="Какой же вы молодец что вот так заботитесь о себе! Надо меньше "
                           "стрессовать"
        )


def main():
    srv = grpc.server(thread_pool=futures.ThreadPoolExecutor(max_workers=50))
    add_TellAboutPhotoServicer_to_server(TellAboutPhotoServer(), srv)
    port = 50051
    srv.add_insecure_port(f'[::]:{port}')
    srv.start()
    srv.wait_for_termination()


if __name__ == "__main__":
    print("start server")
    main()
