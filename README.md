## nairalandgoapi - NairalandGo with API support

***nairalandgoapi*** provides an API to the [*nairalandgo* application](https://github.com/DISCRETEboi/nairaland-go). This allows anyone, anywhere to access the services of the application using just HTTP messages. The API accepts an HTTP POST request containing JSON data, and returns a base64 encoding of the PDF resource generated by the application.

The HTTP request is required to be in the following structure and format:

	POST http://etp4ma.octamile.com:9602
	Content-Type: application/json
	 
	{
	    "url": "https://www.nairaland.com/6314385/samsung-official-thread"
	}

Note the following:  
- **POST** request  
- Server address of **etp4ma.octamile.com**, with port number **9602**  
- **JSON** data type  
- Desired thread link replaces the one in the example  

When an empty string is passed as the thread link in the JSON data, a default link [https://www.nairaland.com/6602040/how-successful-people-achieve-career](https://www.nairaland.com/6602040/how-successful-people-achieve-career) will be processsed by the application. Also, when the desired PDF document is not successfully generated by the application, a pre-generated PDF document will be received by the HTTP client. The pre-generated document is also for the thread link [https://www.nairaland.com/6602040/how-successful-people-achieve-career](https://www.nairaland.com/6602040/how-successful-people-achieve-career).  
Reasons why the PDF may not be generated successfully include:  
- the number of posts in the thread too much. Try a thread with a fewer number of posts.  
- the service dealing with the PDF conversion denying access due to the  monthly conversion limit being exceeded.

A base64 encoding of the PDF document content is returned as the response body. The encoded data can be decoded using appropriate functions. For example, in Golang, the function `base64.StdEncoding.DecodeString()` from the `encoding/base64` library, can be used to decode the encoded contents.

... more functionalities to be added in the future ...
