AI Microservices

This repository contains two microservices written in Go and Python for an AI-based application. The main Go service is developed using the Echo framework, while the Python service utilizes the default library and a worker pool with a fixed number of worker threads.

The interaction between the services is established through gRPC. When a request is received by the main backend (Go), it processes the request and sends a gRPC request to the Python service. The Python service is responsible for image processing tasks, such as detecting the presence of humans in the photo and determining their emotions. The emotion detection model is trained using open datasets and various Python libraries. The project also utilizes the OpenCV (cv2) library to detect the presence of humans in the photo.

The Python service sends back the detected emotion and the presence of a person, and Go leverages the OpenAI API (Chat GPT-3 Turbo) to receive advice on how to handle negative emotions and suggestions on how to maintain positive emotions. On the frontend, JavaScript is used to receive responses from Go and display them beautifully on the website.
Features

    Two microservices: Go and Python
    Go service built with the Echo framework
    Python service utilizing the default library and a worker pool
    Communication between services through gRPC
    Image processing in Python to detect human presence and emotions
    Utilization of open datasets and Python libraries for emotion detection
    OpenCV (cv2) integration for human presence detection
    OpenAI API (Chat GPT-3 Turbo) for emotion handling and guidance
    JavaScript frontend for receiving and displaying responses

Technologies Used

    Go
    Echo framework
    Python
    gRPC
    OpenCV (cv2)
    OpenAI API (Chat GPT-3 Turbo)
    JavaScript

Acknowledgements

I would like to express my gratitude to the open-source community for their contributions and the developers of the Echo framework, OpenCV, and OpenAI API for their invaluable tools and libraries that have contributed to this project.

Enjoy exploring the AI Microservices and their AI-driven functionalities!