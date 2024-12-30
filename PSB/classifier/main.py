from concurrent import futures
import grpc
import classifier_pb2
import classifier_pb2_grpc
import review_classifier

class ParserServiceServicer(classifier_pb2_grpc.ProtoServiceServicer):
    def Predict(self, request, context):
        text = request.original_text
        classifier = review_classifier.ReviewClassifier('keywords.json')
        category = classifier.classify_review(text)
        return classifier_pb2.Response(category)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    classifier_pb2_grpc.add_ProtoServiceServicer_to_server(ParserServiceServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started on port 50051")
    server.wait_for_termination()

if __name__ == "__main__":
    serve()
